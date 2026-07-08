<%@Language="JScript" CodePage="65001" EnableSessionState=false%><%
Response.Clear();
Response.Status			= 200;
Response.ContentType	= "text/plain";
Response.CharSet		= "utf-8";
Response.CacheControl	= "max-age=0, no-cache, no-store";

var arrResults	= [];
var arrTests	= [
	// #1:
	 {expected:",,,,",							actual:new Array(5).toString()}
	,{expected:"Apple,Banana,Orange",			actual:new Array('Apple', 'Banana', 'Orange').toString()}
	,{expected:"Apple,Banana,Orange",			actual:['Apple', 'Banana', 'Orange'].toString()}
	,{expected:"",								actual:[].toString()}
	// #5:
	,{expected:"0",								actual:new Number().toString()}
	,{expected:"5",								actual:new Number(5).toString()}
	,{expected:"5",								actual:(5).toString()}
	,{expected:"NaN",							actual:new Number("cars").toString()}
	,{expected:"1",								actual:new Number(true).toString()}
	// #10:
	,{expected:"NaN",							actual:new Number({}).toString()}
	,{expected:"0",								actual:new Number([]).toString()}
	,{expected:"",								actual:new String().toString()}
	,{expected:"Hello, world",					actual:new String("Hello, world").toString()}
	,{expected:"5",								actual:new String(5).toString()}
	// #15:
	,{expected:"true",							actual:new String(true).toString()}
	,{expected:"[object Object]",				actual:new String({}).toString()}
	,{expected:"",								actual:new String([]).toString()}
	,{expected:"",								actual:"".toString()}
	,{expected:"[object Object]",				actual:new Object().toString()}
	// #20:
	,{expected:"5",								actual:new Object(5).toString()}
	,{expected:"test",							actual:new Object("test").toString()}
	,{expected:"[object Object]",				actual:new Object({}).toString()}
	,{expected:"",								actual:new Object([]).toString()}
	,{expected:"true",							actual:new Object(true).toString()}
	// #25:
	,{expected:"[object Object]",				actual:{}.toString()}
	,{expected:"false",							actual:new Boolean().toString()}
	,{expected:"false",							actual:new Boolean(false).toString()}
	,{expected:"true",							actual:new Boolean("cheese").toString()}
	,{expected:"true",							actual:new Boolean(5).toString()}
	// #30:
	,{expected:"true",							actual:new Boolean({}).toString()}
	,{expected:"true",							actual:new Boolean([]).toString()}
	,{expected:"false",							actual:new Boolean(0).toString()}
	,{expected:"false",							actual:(false).toString()}
	,{expected:"function anonymous() {\n\n}",	actual:new Function().toString()}
	// #35:
	,{expected:"function anonymous() {\nreturn 5\n}",	actual:new Function("return 5").toString()}
	,{expected:"function anonymous(intParam1, intParam2) {\nreturn intParam1 + intParam2\n}",	actual:new Function("intParam1", "intParam2", "return intParam1 + intParam2").toString()}
	,{expected:"//",							actual:new RegExp().toString()}
	,{expected:"[object Object]",				actual:new Enumerator().toString()}
	,{expected:"(function hi(){return \"hi\";})",					actual:(function hi(){return "hi";}).toString()}
	// #40:
	,{expected:"(function(strName){return \"hi \" + strName;})",	actual:(function(strName){return "hi " + strName;}).toString()}
];


// Test Loop
var intLoopIndex	= 0;
var intSuccess		= 0;
var intLoopIndexMax	= arrTests.length;

while (intLoopIndex < intLoopIndexMax) {
	var objTest		= arrTests[intLoopIndex];
	var blnResult	= (objTest.expected === objTest.actual);

	if (blnResult) {
		intSuccess++;
	}
	else {
		arrResults.push("#" + (intLoopIndex+1) + ": Expected = " + objTest.expected + "\t\tActual = " + objTest.actual);
	}

	intLoopIndex++;
} // while()


// Output Results
Response.Write("Tests Total: " + intLoopIndexMax + "\r\n");
Response.Write("Tests Passed: " + intSuccess + "\r\n");
Response.Write("Tests Failed: " + (intLoopIndexMax - intSuccess) + "\r\n\r\n");
Response.Write("Failed tests:\r\n");
Response.Write(arrResults.join("\r\n") || "[None; 😁]");



Response.Write("\r\n➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖\r\nAdditional Response.Write implicit casting to string:\r\nThe text at the bottom should look as follows (toml: default_timezone = \"Europe/London\"):\r\n➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖\r\n,,,,5Hello, world[object Object]function anonymous() {\n\n}Tue Jul 7 10:35:27 UTC+0100 2026False\r\n➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖\r\n");

var arrTest = new Array(5);
Response.Write(arrTest);

var intTest = new Number(5);
Response.Write(intTest);

var strTest = new String("Hello, world");
Response.Write(strTest);

var objTest = new Object();
Response.Write(objTest);

var fnTest = new Function();
Response.Write(fnTest);

var dteTest = new Date("Tue Jul 7 10:35:27 UTC+0100 2026");
Response.Write(dteTest);

var blnTest = new Boolean(false);
Response.Write(blnTest);
%>