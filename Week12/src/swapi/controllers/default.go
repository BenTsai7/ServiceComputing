package controllers

import (
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"bytes"
    "io"
    "net/http"
    "time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

type APIController struct {
	beego.Controller
}

func (c *APIController) Films() {
	db, err := bolt.Open("swapi.db", 0600, nil)
    if err != nil {
        panic(err)
	}
	defer db.Close()
	params := c.Ctx.Input.Params()
	if (len(params)!=2){
		c.Abort("404")
		return
	}
	first := true
	for k, v := range params {
		if (first){
			first = false
			continue
		}
		beego.Notice(k, v)
		//query
		err = db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Film"))
        //往表里面存储数据
        if b != nil {
			data := b.Get([]byte(v))
			if data!=nil{
				beego.Notice("Found in DataBase")
				c.Data["json"] = string(data)
				c.ServeJSON()
				return nil
			}else{
				url := "https://swapi.co/api/films/"+v
				client := http.Client{Timeout: 5 * time.Second}
				resp, err := client.Get(url)
				if (resp.StatusCode==404){
					c.Abort("404")
					return nil
				}
    			if err != nil {
        			panic(err)
    			}
    			defer resp.Body.Close()
    			var buffer [512]byte
    			result := bytes.NewBuffer(nil)
    			for {
        			n, err := resp.Body.Read(buffer[0:])
        			result.Write(buffer[0:n])
        			if err != nil && err == io.EOF {
            			break
     				} else if err != nil {
            			c.Abort("404")
        			}
				}
				err = b.Put([]byte(v),[]byte(result.String()))
            	if err != nil {
					panic(err)
            	}
				beego.Notice(result.String())
				c.Data["json"] = result.String()
				c.ServeJSON()
				return nil
			}
        }
		return nil
   	 	})
    	if err != nil {
        	panic(err)
		}
	}
}

func (c *APIController) People() {
	db, err := bolt.Open("swapi.db", 0600, nil)
    if err != nil {
        panic(err)
	}
	defer db.Close()
	params := c.Ctx.Input.Params()
	if (len(params)!=2){
		c.Abort("404")
		return
	}
	first := true
	for k, v := range params {
		if (first){
			first = false
			continue
		}
		beego.Notice(k, v)
		//query
		err = db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("People"))
        //往表里面存储数据
        if b != nil {
			data := b.Get([]byte(v))
			if data!=nil{
				beego.Notice("Found in DataBase")
				c.Data["json"] = string(data)
				c.ServeJSON()
				return nil
			}else{
				url := "https://swapi.co/api/people/"+v
				client := http.Client{Timeout: 5 * time.Second}
				resp, err := client.Get(url)
				if (resp.StatusCode==404){
					c.Abort("404")
					return nil
				}
    			if err != nil {
        			panic(err)
    			}
    			defer resp.Body.Close()
    			var buffer [512]byte
    			result := bytes.NewBuffer(nil)
    			for {
        			n, err := resp.Body.Read(buffer[0:])
        			result.Write(buffer[0:n])
        			if err != nil && err == io.EOF {
            			break
     				} else if err != nil {
            			c.Abort("404")
        			}
				}
				err = b.Put([]byte(v),[]byte(result.String()))
            	if err != nil {
					panic(err)
            	}
				beego.Notice(result.String())
				c.Data["json"] = result.String()
				c.ServeJSON()
				return nil
			}
        }
		return nil
   	 	})
    	if err != nil {
        	panic(err)
		}
	}
}

func (c *APIController) Planets() {
	db, err := bolt.Open("swapi.db", 0600, nil)
    if err != nil {
        panic(err)
	}
	defer db.Close()
	params := c.Ctx.Input.Params()
	if (len(params)!=2){
		c.Abort("404")
		return
	}
	first := true
	for k, v := range params {
		if (first){
			first = false
			continue
		}
		beego.Notice(k, v)
		//query
		err = db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Planet"))
        //往表里面存储数据
        if b != nil {
			data := b.Get([]byte(v))
			if data!=nil{
				beego.Notice("Found in DataBase")
				c.Data["json"] = string(data)
				c.ServeJSON()
				return nil
			}else{
				url := "https://swapi.co/api/planets/"+v
				client := http.Client{Timeout: 5 * time.Second}
				resp, err := client.Get(url)
				if (resp.StatusCode==404){
					c.Abort("404")
					return nil
				}
    			if err != nil {
        			panic(err)
    			}
    			defer resp.Body.Close()
    			var buffer [512]byte
    			result := bytes.NewBuffer(nil)
    			for {
        			n, err := resp.Body.Read(buffer[0:])
        			result.Write(buffer[0:n])
        			if err != nil && err == io.EOF {
            			break
     				} else if err != nil {
            			c.Abort("404")
        			}
				}
				err = b.Put([]byte(v),[]byte(result.String()))
            	if err != nil {
					panic(err)
            	}
				beego.Notice(result.String())
				c.Data["json"] = result.String()
				c.ServeJSON()
				return nil
			}
        }
		return nil
   	 	})
    	if err != nil {
        	panic(err)
		}
	}
}

func (c *APIController) Species() {
	db, err := bolt.Open("swapi.db", 0600, nil)
    if err != nil {
        panic(err)
	}
	defer db.Close()
	params := c.Ctx.Input.Params()
	if (len(params)!=2){
		c.Abort("404")
		return
	}
	first := true
	for k, v := range params {
		if (first){
			first = false
			continue
		}
		beego.Notice(k, v)
		//query
		err = db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Specie"))
        //往表里面存储数据
        if b != nil {
			data := b.Get([]byte(v))
			if data!=nil{
				beego.Notice("Found in DataBase")
				c.Data["json"] = string(data)
				c.ServeJSON()
				return nil
			}else{
				url := "https://swapi.co/api/species/"+v
				client := http.Client{Timeout: 5 * time.Second}
				resp, err := client.Get(url)
				if (resp.StatusCode==404){
					c.Abort("404")
					return nil
				}
    			if err != nil {
        			panic(err)
    			}
    			defer resp.Body.Close()
    			var buffer [512]byte
    			result := bytes.NewBuffer(nil)
    			for {
        			n, err := resp.Body.Read(buffer[0:])
        			result.Write(buffer[0:n])
        			if err != nil && err == io.EOF {
            			break
     				} else if err != nil {
            			c.Abort("404")
        			}
				}
				err = b.Put([]byte(v),[]byte(result.String()))
            	if err != nil {
					panic(err)
            	}
				beego.Notice(result.String())
				c.Data["json"] = result.String()
				c.ServeJSON()
				return nil
			}
        }
		return nil
   	 	})
    	if err != nil {
        	panic(err)
		}
	}
}

func (c *APIController) Starships() {
	db, err := bolt.Open("swapi.db", 0600, nil)
    if err != nil {
        panic(err)
	}
	defer db.Close()
	params := c.Ctx.Input.Params()
	if (len(params)!=2){
		c.Abort("404")
		return
	}
	first := true
	for k, v := range params {
		if (first){
			first = false
			continue
		}
		beego.Notice(k, v)
		//query
		err = db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Starship"))
        //往表里面存储数据
        if b != nil {
			data := b.Get([]byte(v))
			if data!=nil{
				beego.Notice("Found in DataBase")
				c.Data["json"] = string(data)
				c.ServeJSON()
				return nil
			}else{
				url := "https://swapi.co/api/starships/"+v
				client := http.Client{Timeout: 5 * time.Second}
				resp, err := client.Get(url)
				if (resp.StatusCode==404){
					c.Abort("404")
					return nil
				}
    			if err != nil {
        			panic(err)
    			}
    			defer resp.Body.Close()
    			var buffer [512]byte
    			result := bytes.NewBuffer(nil)
    			for {
        			n, err := resp.Body.Read(buffer[0:])
        			result.Write(buffer[0:n])
        			if err != nil && err == io.EOF {
            			break
     				} else if err != nil {
            			c.Abort("404")
        			}
				}
				err = b.Put([]byte(v),[]byte(result.String()))
            	if err != nil {
					panic(err)
            	}
				beego.Notice(result.String())
				c.Data["json"] = result.String()
				c.ServeJSON()
				return nil
			}
        }
		return nil
   	 	})
    	if err != nil {
        	panic(err)
		}
	}
}

func (c *APIController) Vehicles() {
	db, err := bolt.Open("swapi.db", 0600, nil)
    if err != nil {
        panic(err)
	}
	defer db.Close()
	params := c.Ctx.Input.Params()
	if (len(params)!=2){
		c.Abort("404")
		return
	}
	first := true
	for k, v := range params {
		if (first){
			first = false
			continue
		}
		beego.Notice(k, v)
		//query
		err = db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Vehicle"))
        //往表里面存储数据
        if b != nil {
			data := b.Get([]byte(v))
			if data!=nil{
				beego.Notice("Found in DataBase")
				c.Data["json"] = string(data)
				c.ServeJSON()
				return nil
			}else{
				url := "https://swapi.co/api/vehicles/"+v
				client := http.Client{Timeout: 5 * time.Second}
				resp, err := client.Get(url)
				if (resp.StatusCode==404){
					c.Abort("404")
					return nil
				}
    			if err != nil {
        			panic(err)
    			}
    			defer resp.Body.Close()
    			var buffer [512]byte
    			result := bytes.NewBuffer(nil)
    			for {
        			n, err := resp.Body.Read(buffer[0:])
        			result.Write(buffer[0:n])
        			if err != nil && err == io.EOF {
            			break
     				} else if err != nil {
            			c.Abort("404")
        			}
				}
				err = b.Put([]byte(v),[]byte(result.String()))
            	if err != nil {
					panic(err)
            	}
				beego.Notice(result.String())
				c.Data["json"] = result.String()
				c.ServeJSON()
				return nil
			}
        }
		return nil
   	 	})
    	if err != nil {
        	panic(err)
		}
	}
}