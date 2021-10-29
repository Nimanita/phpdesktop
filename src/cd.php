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
