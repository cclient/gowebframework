
var topname="ApApiItem";
var defaulttype="string";
var bson=true; //对应 mongodb 
var json=true; //http response 
var scheme=false; //http request->scheme 
var jdata={
    "_id" : "564d5162e54b3106fb7badea",
    "macs" : [
        "00-21-26-00-C8-B0"
    ],
    "time" : 1447907400,
    "timestr" : "2015-11-19 12:30",
    "shop":{
        "name":"shop1"
    }
};
String.prototype.firstToUpperCase=function(){
    return this[0].toUpperCase()+this.substring(1);
}
var fun=(function(){
    var otherobj=[];
    var goobjstring="";
    function getStruct(data,collectionname){
        goobjstring+="type "+collectionname.firstToUpperCase()+" struct {\n";
        var per="\t";
        for(var key in data){
            var newkey=key.firstToUpperCase();
            goobjstring+=per +newkey+" "+getType(data[key],key);
            if (json||bson||scheme){
                goobjstring+=' `';
                if (json){
                    goobjstring+='json:"'+key+'"';
                }
                if (bson){
                    goobjstring+='bson:"'+key+'"';
                }
                if (scheme){
                    goobjstring+='scheme:"'+key+'"';
                }
                goobjstring+='`';
            }
            goobjstring+="\n";
        }
        goobjstring+="}\n";
        while (otherobj.length>0){
            var subobj=otherobj.pop();
            getStruct(subobj.obj,subobj.key)
        }
        return goobjstring
    }
    function getType(obj,key){
        var type=defaulttype;
        if(obj){
            switch(obj.constructor)
            {
                case Array:
                    type="[]"+getType(obj[0]||"",key.firstToUpperCase()) ;
                    break;
                case Object:
                    otherobj.push({key:key,obj:obj});
                    type=key.firstToUpperCase()
                    break;
                case String:
                    type="string"
                    break;
                case Number:
                    type="int"
                    break;
                case Boolean:
                    type="bool"
                    break;
                default :
            }
        }
        return type;
    }
    return getStruct
})()

console.log(fun(jdata,topname))
