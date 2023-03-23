# Reference Material

https://docs.google.com/document/d/1MoOk6rNfqYKCIObz4-sGnRMZEHoJStAIZS2Ofw0UFUU/edit


# Reference Material

https://docs.google.com/document/d/1MoOk6rNfqYKCIObz4-sGnRMZEHoJStAIZS2Ofw0UFUU/edit


# Docker compose installation:
     https://docs.docker.com/compose/install/

# Docker installation:
     https://docs.docker.com/engine/install/


# hey installation : 
    GitHub - rakyll/hey: HTTP load generator, ApacheBench (ab) replacement  
    https://github.com/rakyll/hey@latest

    Mac User: brew install hey
    go install github.com/rakyll/hey@latest
    hey -m GET -c 100 -n 1000 http://localhost:3000/readiness/10


# Expvarmon
go install github.com/divan/expvarmon@latest




expvarmon -ports="localhost:4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"
