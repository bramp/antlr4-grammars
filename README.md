# antlr4test-go
by Andrew Brampton

Precompiled Go versions of most of the grammars on github.com/antlr/grammars-v4.

```go
Example:
TODO
```

## Supported

[ OK ] abnf           
[ OK ] agc            
[ OK ] arithmetic     
[ OK ] atl            
[ OK ] bnf            
[ OK ] brainfuck      
[ OK ] c              
[ OK ] clf            
[ OK ] clif           
[ OK ] cobol85        
[ OK ] cookie         
[ OK ] cool           
[ OK ] creole         
[ OK ] csv            
[ OK ] datetime       
[ OK ] dcm            
[ OK ] dice           
[ OK ] emailaddress   
[ OK ] fasta          
[ OK ] fol            
[ OK ] gml            
[ OK ] gtin           
[ OK ] idl            
[ OK ] iri            
[ OK ] istc           
[ OK ] java           
[ OK ] json           
[ OK ] lambda         
[ OK ] lcc            
[ OK ] mdx            
[ OK ] memcached_protocol
[ OK ] metric         
[ OK ] modelica       
[ OK ] molecule       
[ OK ] morsecode      
[ OK ] mps            
[ OK ] mumath         
[ OK ] mumps          
[ OK ] muparser       
[ OK ] p              
[ OK ] pcre           
[ OK ] peoplecode     
[ OK ] postalcode     
[ OK ] powerbuilder   
[ OK ] propcalc       
[ OK ] r              
[ OK ] rcs            
[ OK ] robotwars      
[ OK ] romannumerals  
[ OK ] rpn            
[ OK ] ruby           
[ OK ] scss           
[ OK ] sexpression    
[ OK ] snobol         
[ OK ] suokif         
[ OK ] telephone      
[ OK ] tiny           
[ OK ] tinyc          
[ OK ] tnt            
[ OK ] useragent      
[ OK ] vhdl           
[ OK ] wavefront      
[ OK ] xpath          

## Update
To pull down the latest grammars and compile them:

```bash
git submodule init
git submodule update

# TODO Instructions to download antlr
# Update the Makefile
go run makemake.go

# Now build all the grammars
make all -k -j 8 2> /dev/null

# Update the table at the top with new successes or failures
```
