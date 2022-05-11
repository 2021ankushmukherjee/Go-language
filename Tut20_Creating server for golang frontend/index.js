const express = require("express");
const app = express();
port = 4000;


app.use(express.json())
app.use(express.urlencoded({extended: true}))

app.get("/", (req,res) =>{

    res.status(200).json({

        message: "Welcome to my server"
    })
})

app.post("/", (req,res) =>{

    const name = req.body.name;
    
    res.status(200).send(name)
});


app.post("/postform", (req,res)=>{

    res.status(200).send(req.body);
});



app.listen(port, ()=>{

    console.log(`server is running on ${port}`);
});