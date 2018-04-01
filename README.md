# kibup

a simple and smart way to back up kibana objects

## Instalation

Install the client package from [GitHub](https://github.com/leocomelli/kibup):

```
go get github.com/leocomelli/kibup
```

or install it from a source code checkout:

```
git clone git@github.com:leocomelli/kibup.git
make
./dist/kibup
```

## Usage

```shell
./kibup 
Error: required flag(s) "repo", "token" not set
Usage:
  kibup [flags]

Flags:
      --author-email string   github author email (default "kibup")
      --author-name string    github author name (default "kibup@kibup.com")
      --file string           backup filename (default "kibana.json")
      --github string         github api url (default "https://api.github.com/")
  -h, --help                  help for kibup
      --host string           elasticsearch host:port (default "http://127.0.0.1:9200")
      --index string          kibana index name (default ".kibana")
      --local                 create file locally
      --repo string           repository name (owner/repo)
      --size int              elastisearch query result size (default 10000)
      --sort string           field to sort the elasticsearch query (default "_type")
      --token string          github personal access token
      --types stringArray     kibana object types (default [dashboard,visualization,search])

required flag(s) "repo", "token" not set
```

## Example

```shell
./kibup --token s3cr3t \
        --host http://192.168.1.120:9200 \
        --repo leocomelli/kibana \        
        --file mybackup.json \
        --local
```

## Compatibility

* GitHub API v3
* Elasticsearch 5.x

## References

* [GitHub API v3](https://developer.github.com/v3/)
* [ GitHub Personal Access Token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/)
* [Elasticsearch 5.x](https://www.elastic.co/guide/en/elasticsearch/reference/5.0/index.html) 

## Contributing

All contributions are welcome: ideas, patches, documentation, bug reports, complaints, and even something you drew up on a napkin.

Programming is not a required skill. Whatever you've seen about open source and maintainers or community members saying "send patches or die" - you will not see that here.

It is more important to the community that you are able to contribute.