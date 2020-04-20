$(document).ready(function () {
    $("#txtdate").datepicker({
        
        beforeShowDay: $.datepicker.noWeekends,
        minDate: 0,
        maxDate: 16
    });
});
window.onload=function fn(){         
var hours = new Date();
hours=hours.getHours()
var minutes = new Date();
minutes=minutes.getMinutes()
if (hours <10) {
   hours=hours.toString()
   hours="0"+hours
}
if (minutes <10) {
    minutes=minutes.toString()
    minutes="0"+minutes
 }
time= hours+":"+minutes
 
date = new Date()
date=date.getDate()
selecteddate = document.getElementById("txtdate").value
this.console.log("value"+selecteddate)
document.getElementById("timepicker").setAttribute("value", time);
document.getElementById("timepicker").setAttribute("min", time);
document.getElementById("timepicker").setAttribute("max", "19:00");
minutes =new Date()
z=minutes.getMinutes()+30
console.log(z)
if(z>59){
    hours = parseInt(hours)
    console.log(hours)
    hours=  hours+ 1
    hours=hours.toString()
   }
minutes.setMinutes(z)
x=minutes.getMinutes()
console.log(x)

if (x <10) {
    x=x.toString()
    x="0"+x
 }


time= hours +":"+x
document.getElementById("Endpicker").setAttribute("value", time);
document.getElementById("Endpicker").setAttribute("min", time);
document.getElementById("Endpicker").setAttribute("max", "19:00");
}
var i=2
function add(){   
var txtNewlableBox = document.createElement('div');
txtNewlableBox.innerHTML = "<label>Attendee"+i+" "+"EmailId</label>";
document.getElementById("tr").appendChild(txtNewlableBox);
var txtNewInputBox = document.createElement('div');
txtNewInputBox.innerHTML = "<input type='email' id='newInputBox'>";
document.getElementById("tr").appendChild(txtNewInputBox);
i++
}
function fn(){
    selecteddate = document.getElementById("txtdate").value
    d1 = selecteddate[3].toString()
    d2 = selecteddate[4].toString()
    d= d1+d2
    d= parseInt(d)

    date = new Date()
    date=date.getDate()
    console.log(date)
    if(d>date){
        document.getElementById("timepicker").setAttribute("min", "09:00");
        document.getElementById("timepicker").setAttribute("value", "09:00");
        document.getElementById("timepicker").setAttribute("max", "19:00");
    }
    console.log("value in fn"+d)
    console.log(typeof(d))
    end()
}
function end(){
    selecteddate = document.getElementById("txtdate").value
    d1 = selecteddate[3].toString()
    d2 = selecteddate[4].toString()
    d= d1+d2
    d= parseInt(d)
    date1 = new Date()
    date1=date1.getDate()
    console.log(date1)
    if(d>date){
        document.getElementById("Endpicker").setAttribute("min", "09:30");
        document.getElementById("Endpicker").setAttribute("value", "09:30");
        document.getElementById("Endpicker").setAttribute("max", "19:00");
    }
    console.log("value in fn"+d)
    console.log(typeof(d))
}