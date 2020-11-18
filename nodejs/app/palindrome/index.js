const express = require('express');
const app = express();
const axios = require('axios');
const fs = require('fs');


app.get('/palindrome',async function(req,res){
  var palindromresponse = {
    reversedstring : "",
    actualstring : "",
    isPalindrome: false
  }
  axios.get('http://reverse:8080/reverse')
  .then(reverseresponse => {
    palindromresponse.reversedstring = reverseresponse.data;
    fs.readFile("/usr/app/palindrome/reverse.txt", 'utf8', function (err,data) {
      if (err) {
        return "";
      }
      palindromresponse.actualstring = data.trim();
      if(palindromresponse.reversedstring == palindromresponse.actualstring) {
        palindromresponse.isPalindrome = true;
        
      }else {
        palindromresponse.isPalindrome = false;
      }
       res.json(palindromresponse);
    });
   
  })
  .catch(error => {
    console.log(error);
  });

 
  app.get('/error',async function(req,res){
    process.exit(1);
  });

});
app.listen(8080, () =>{
    console.log('listening on port 8080');
});








