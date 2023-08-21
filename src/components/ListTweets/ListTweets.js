import React, { useState, useEffect } from "react";
import { Image } from "react-bootstrap";
import { map } from "lodash";
import localization from "moment/locale/es"
import moment from "moment";
import AvatarNoFound from "../../assets/png/alekhya.jpeg";
import { API_HOST } from "../../utils/constant";
import { getUserApi } from "../../api/user";
import { replaceURLWithHTMLLinks } from "../../utils/functions";

import "./ListTweets.scss";

export default function ListTweets(props) {
  const { tweets } = props;

  return (
    <div className="list-tweets">
      {map(tweets, (tweet, index) => (
        <Tweet key={index} tweet={tweet} />
      ))}
    </div>
  );
}

function Tweet(props) {
  const { tweet } = props;
  const [userInfo, setUserInfo] = useState(null);
  const [avatarUrl, setAvatarUrl] = useState(null);

  useEffect(() => {
    getUserApi(tweet.userID).then((response) => {
      setUserInfo(response);
      console.log(response)
      setAvatarUrl(
        response?.avatar
        
          ? `${API_HOST}/getAvatar?id=${response.id}`
          : AvatarNoFound
      );
    });
  }, [tweet]);

  return (
    <div className="tweet">
      <Image className="avatar" src={avatarUrl} roundedCircle />
      <div>
        <div className="name">
          {userInfo?.name} {userInfo?.lastname} 
          <span>  </span>
          <span>{moment(tweet.date)
              .locale("en", localization)
              .format("LL")}</span>
        </div>
        <div
          dangerouslySetInnerHTML={{
            __html: replaceURLWithHTMLLinks(tweet.message),
          }}
        />
      </div>
    </div>
  );
}