project_urls={
        "Documentation": "https://tweepy.readthedocs.io",
        "Issue Tracker": "https://github.com/tweepy/tweepy/issues",
        "Source Code": "https://github.com/tweepy/tweepy",
    },
    download_url="https://pypi.org/project/tweepy/",
    packages=find_packages(exclude=["tests", "examples"]),
   self.api_codes.append(error["code"])
                if "message" in error:
                    self.api_messages.append(error["message"])
                if "code" in error and "message" in error:
                    error_text += f"\n{error['code']} - {error['message']}"
                elif "message" in error:
                    error_text += '\n' + error["message"]
                      for error in errors:
                self.api_errors.append(error)
                if "code" in error:
                    self.api_codes.append(error["code"])
                if "message" in error:
                    self.api_messages.append(error["message"])
                if "code" in error and "message" in error:
 error_text = ""
            for error in errors:
                self.api_errors.append(error)
                if "code" in error:
                    self.api_codes.append(error["code"])
                if "message" in error:
                    self.api_messages.append(error["message"])
# Copyright 2009-2021 Joshua Roesslein
# See LICENSE for details.
'url' => env('APP_URL', 'http://localhost'),
# Appengine users: https://developers.google.com/appengine/docs/python/sockets/#making_httplib_use_sockets

import json
import logging
  References
        ----------
        https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/create-manage-lists/api-reference/get-lists-list
        """
        return self.request(
            'GET', 'lists/list', endpoint_parameters=(
  )

    def get_liking_users(self, id, *, user_auth=False, **params):
        """get_liking_users(id, *, expansions, media_fields, place_fields, \
                            poll_fields, tweet_fields, user_fields)

        Allows you to get information about a Tweet’s liking users.
  "DELETE", route, user_auth=True
        )

    def get_liking_users(self, id, *, user_auth=False, **params):
        """get_liking_users(id, *, expansions, media_fields, place_fields, \
                            poll_fields, tweet_fields, user_fields)
  https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/delete-users-user_id-likes
        """
        id = self.access_token.partition('-')[0]
        route = f"/2/users/{id}/likes/{tweet_id}"

        return self._make_request(
            "DELETE", route, user_auth=True
  References
        ----------
        https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/delete-users-user_id-likes
        """
        id = self.access_token.partition('-')[0]
