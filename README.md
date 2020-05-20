# utilsgo

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

utilsgo is a general purpose utility package for golang projects

  - Import the package
  - Use the package functions
  - Save time

# Features!

  - zap logger
  - success failure printing functions
  - pretty print (print struct in json format on console)
  - generate UUID
  - Extract seconds from time string

# Quickstart
    utilsgo.Bootstrap("info.log","err.log")
    
# Funtions
    //generate UUID
    uuid:=utilsgo.UUID()
    
    //printint success
    utilsgo.PrintSuccess("This is success")

License
----

MIT

