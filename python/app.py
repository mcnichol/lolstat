import requests

august_summoner="batterystaple123"
dad_summoner="Buckethead Wendy"
api_key=""
url_base="https://na1.api.riotgames.com/"
url_summoner_v4_byAccount          ="lol/summoner/v4/summoners/by-account/"
url_summoner_v4_byName             ="lol/summoner/v4/summoners/by-name/"
url_summoner_v4_byPuuid            ="lol/summoner/v4/summoners/by-puuid/"
url_summoner_v4_byEncrypted        ="lol/summoner/v4/summoners/"


def request_summoner_data():
    url = url_base + url_summoner_v4_byName + august_summoner + "?api_key=" + api_key

    response = requests.get(url).json()

    mySummoner = Summoner(response)

    return mySummoner

class Summoner:
    encryptedId         = None
    encryptedAccountId  = None
    level               = None
    name                = None
    puuid               = None
    profileIconId       = None
    revisionDate        = None

    def __init__(self, summonerDTO):
        self.encryptedId         = summonerDTO.get('id')
        self.encryptedAccountId  = summonerDTO.get('accountId')
        self.level               = summonerDTO.get('summonerLevel')
        self.name                = summonerDTO.get('name')
        self.puuid               = summonerDTO.get('puuid')
        self.profileIconId       = summonerDTO.get('profileIconId')
        self.revisionDate        = summonerDTO.get('revisionDate')

    def __str__(self):
        return "Account Id: " + self.encryptedAccountId

#class Match
#class MatchList

def main():
    print(request_summoner_data())

if __name__ == "__main__":
    main()

