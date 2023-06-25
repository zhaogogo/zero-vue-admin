
export const Unix = (sec) => {
    var t = new Date(sec)
    
    var year = t.getFullYear();
    
    var month = t.getMonth();
    month = month < 10 ? "0"+month:month;
    
    var day = t.getDate();  
    day = day < 10 ? "0"+day:day;
    
    var hour = t.getHours(); 
    hour = hour < 10 ? "0"+hour:hour;

    var minute = t.getMinutes(); 
    minute = minute < 10 ? "0"+minute:minute;

    var second = t.getSeconds(); 
    second = second < 10 ? "0"+second:second;

    return year+"-"+month+"-"+day+" "+hour+":"+minute+":"+second;
}