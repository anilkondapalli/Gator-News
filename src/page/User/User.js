import React, { useState, useEffect } from "react";
import { Button, Spinner } from "react-bootstrap";
import BasicLayout from "../../layout/BasicLayout";
import { withRouter } from "react-router-dom";
import { getUserApi } from "../../api/user";
import { toast } from "react-toastify";
import BannerAvatar from "../../components/User/BannerAvatar";
import useAuth from "../../hooks/useAuth";
import InfoUser from "../../components/User/InfoUser"
import { getUserTweetsApi } from "../../api/tweet";
import ListTweets from "../../components/ListTweets";


import "./User.scss";

 function User(props) {
    const { match} = props;
  const [user, setUser] = useState(null);
    const { params } = match;
    const [tweets, setTweets] = useState(null);
    const loggedUser=useAuth();
    console.log(loggedUser);

    console.log(tweets);

    useEffect(() => {
        getUserApi(params.id)
        .then(response => {
            if (!response){ 
                toast.warning("The user you visited does not exist");
                }
                setUser(response);
            })
          .catch(() => {
            toast.warning("The user you visited does not exist");
          });
      }, [params]);

      useEffect(() => {
        getUserTweetsApi(params.id, 1)
          .then((response) => {
            setTweets(response);
          })
          .catch(() => {
            setTweets([]);
          });
      }, [params]);

    return(
        <div className="basicLayout">
        <BasicLayout className="user">
            <BannerAvatar user={user} loggedUser={loggedUser} />  
            <InfoUser user={user}></InfoUser> 
            <div className="user__tweets"> 
              <h3><b>User Messages</b></h3>
              {tweets && <ListTweets tweets={tweets} />}
            </div>
        </BasicLayout>
        </div>
    )
}
export default withRouter(User);
