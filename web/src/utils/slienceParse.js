export const matcherExpParse = (exp,hostID,slienceNameID,id) => {
    var reg = /^["|'](.*)["|']$/g;
    let index = exp.indexOf("=~")
    if ( index !== -1) {
        return {name:exp.substring(0,index), value: exp.substring(index+2,exp.length).replace(reg,"$1"),is_regex: true,is_equal:true,host_id: hostID,slience_name_id:slienceNameID, id: id}
    }
    index = exp.indexOf("!~")
    if ( index !== -1) {
        return {name:exp.substring(0,index), value: exp.substring(index+2,exp.length).replace(reg,"$1"),is_regex: true,is_equal:false,host_id: hostID,slience_name_id:slienceNameID, id: id}
    }
    index = exp.indexOf("!=")
    if ( index !== -1) {
        return {name:exp.substring(0,index), value: exp.substring(index+2,exp.length).replace(reg,"$1"),is_regex: false,is_equal:false,host_id: hostID,slience_name_id:slienceNameID, id: id}
    }
    index = exp.indexOf("=")
    if ( index !== -1) {
        return {name:exp.substring(0,index), value: exp.substring(index+1,exp.length).replace(reg,"$1"),is_regex: false,is_equal:true,host_id: hostID,slience_name_id:slienceNameID, id: id}
    }
}

export const matcherObjToExp = (obj) => {
    if (obj.is_regex) {
        if (obj.is_equal) {
            return {exp: obj.name + "=~" + obj.value, id: obj.id, hostID: obj.host_id, slienceNameID: obj.slience_name_id}
        }else {
            return {exp: obj.name + "!~" + obj.value, id: obj.id, hostID: obj.host_id, slienceNameID: obj.slience_name_id}
        }
    } else {
        if (obj.is_equal) {
            return {exp: obj.name + "=" + obj.value, id: obj.id, hostID: obj.host_id, slienceNameID: obj.slience_name_id}
        }else {
            return {exp: obj.name + "!=" + obj.value, id: obj.id, hostID: obj.host_id, slienceNameID: obj.slience_name_id}
        }
    }
     
}