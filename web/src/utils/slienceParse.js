export const matcherExpParse = (exp,hostID,slienceNameID) => {
    let index = exp.indexOf("=~")
    if ( index !== -1) {
        return {name:exp.substring(0,index), value: exp.substring(index+2,exp.length),is_regex: true,is_equal:true,host_id: hostID,slience_name_id:slienceNameID}
    }
    index = exp.indexOf("!~")
    if ( index !== -1) {
        return {name:exp.substring(0,index), value: exp.substring(index+2,exp.length),is_regex: true,is_equal:false,host_id: hostID,slience_name_id:slienceNameID}
    }
    index = exp.indexOf("!=")
    if ( index !== -1) {
        return {name:exp.substring(0,index), value: exp.substring(index+2,exp.length),is_regex: false,is_equal:false,host_id: hostID,slience_name_id:slienceNameID}
    }
    index = exp.indexOf("=")
    if ( index !== -1) {
        return {name:exp.substring(0,index), value: exp.substring(index+1,exp.length),is_regex: false,is_equal:true,host_id: hostID,slience_name_id:slienceNameID}
    }
}