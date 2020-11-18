const express = require('express');

const app = express();

app.get('/reversed',async function(req,res){
    let string = req.query.string;
    let rerversedstring = "";
    for (let i= string.length-1;i>=0;i--)
    {
        var str = string.charAt(i)
    rerversedstring =rerversedstring + str;
    }
    res.json(rerversedstring);
});
app.listen(8080, () =>{
    console.log('listening on port 8080');
});
app.get('/palindrome',async function(req,res){
    let string =req.query.string;
    let reversedstring ="";
    for (let i= string.length-1;i>=0;i--)
    {
        var str = string.charAt(i)
        reversedstring = reversedstring + str;
    } 
   if (reversedstring == string) {
    res.json(true)
  }else {

    res.json(false)
  }
});
app.listen(8081, () =>{
    console.log('listening on port 8081');
});
app.get('/star',async function(req,res){
    let string = req.query.string;
    if(string == "left"){
        let output = "";
        for(let i = 1; i <= 5; i++ ){
            for(let j = 1; j <= i; j++) {
             output = output + "*";

            }
            output = output + "\n";
        }

        res.json(output)
    }
});
app.listen(8082, () =>{
    console.log('listening on port 8082');
});
