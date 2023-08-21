package bd

import(
	"context"
	"fmt"
	"time"

	"github.com/SaiKumarMalve/Gator-News/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
/* ReadAllUsers reads the users registered in the system, if "R" is received in those who bring only those who are related to me*/
func ReadAllUsers(ID string,page int64,search string,tipo string) ([]*models.User,bool){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("gatorNews")
	col := db.Collection("users")

	var results []*models.User

	findOptions:=options.Find()
	findOptions.SetSkip((page -1)*20)
	findOptions.SetLimit(20)
	query:=bson.M{
		"name":bson.M{"$regex":`(?i)`+ search},
	}
	cur,err := col.Find(ctx,query,findOptions)
	if err != nil{
		fmt.Println(err.Error())
		return results,false
	}
	var found, include bool
	for cur.Next(ctx){
		var s models.User
		err:=cur.Decode(&s)
		if err != nil{
			fmt.Println(err.Error())
			return results,false
		}
		var r models.Relationship
		r.UserID=ID
		r.UserRelationshipID=s.ID.Hex()

		include = false
		found, err = ConsultRelation(r)
		if tipo == "new" && found == false{
			include=true
		}
		if tipo == "follow" && found == true{
			include=true
		}
		if r.UserRelationshipID ==ID{
			include=false
		}
		if include == true{
			s.Password=""
			s.Biography=""
			s.Website=""
			s.Location=""
			s.Banner=""
			s.Email=""

			results = append(results, &s)
		}

	}
	err = cur.Err()
	if err != nil{
		fmt.Println(err.Error())
		return results,false
	}
	cur.Close(ctx)
	return results,true

}

