# League Stats
Get statistics on your peeps in league

## Python
August is doing most of the work here.
Python based Rest Client against Riot Dev Portal

### Usage
```shell script
git clone https://github.com/mcnichol/lolstat.git && cd lolstat/python
python3 -m venv venv && source venv/bin/activate && pip install -r requirements.txt
python3 python/app.py
```

## Golang
Dad doing the work here

Term-UI based HTOP style view of League Stats

### Usage
```shell script
git clone https://github.com/mcnichol/lolstat.git && cd lolstat/golang
echo "$RIOT_API_KEY" | config/riot-api.key  //Set RIOT_API_KEY to your key from the Dev Portal
go build -o bin/lolstat && bin/lolstat
```