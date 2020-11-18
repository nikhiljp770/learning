const express = require('express');
const app = express();

app.get('/palindrome',async function(req,res){
    let string =req.query.string;
    let reversedstring ="";
    for (let i= string.length-1;i>=0;i--)
    {
        var str = string.charAt(i)
        reversedstring = reversedstring + str;
   // res.json(reversedstring);  
    } 
   if (reversedstring == string) {
    res.json(true)
  }else {

    res.json(false)
  }
});
app.listen(8080, () =>{
    console.log('listening on port 8080');
});
