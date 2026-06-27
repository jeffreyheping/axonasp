<%@ Language="JavaScript"%>
<%
try {
	Response.Clear();
	Response.Status			= 200;
	Response.ContentType	= "text/plain";
	Response.CharSet		= "utf-8";
	Response.CacheControl	= "max-age=0, no-cache, no-store";

	var arrResults	= [];
	var arrTests	= [
		 {expected:1,					actual:1}
		,{expected:-1,					actual:-1}
		,{expected:null,				actual:null}
		,{expected:"a",					actual:"a"}
		,{expected:"😱",				actual:"😱"}

		 // Date
		,{expected:"number",			actual:typeof (new Date()).getDate()}
		,{expected:"function",			actual:typeof Date}
		,{expected:"string",			actual:typeof Date()}
		,{expected:"object",			actual:typeof new Date()}
		,{expected:"number",			actual:typeof (new Date()).getFullYear()}
		,{expected:2000,				actual:(new Date(2000, 0, 2, 3, 4, 5, 6)).getFullYear()}
		,{expected:0,					actual:(new Date(2000, 0, 2, 3, 4, 5, 6)).getMonth()}
		,{expected:2,					actual:(new Date(2000, 0, 2, 3, 4, 5, 6)).getDate()}
		,{expected:3,					actual:(new Date(2000, 0, 2, 3, 4, 5, 6)).getHours()}
		,{expected:4,					actual:(new Date(2000, 0, 2, 3, 4, 5, 6)).getMinutes()}
		,{expected:5,					actual:(new Date(2000, 0, 2, 3, 4, 5, 6)).getSeconds()}
		,{expected:6,					actual:(new Date(2000, 0, 2, 3, 4, 5, 6)).getMilliseconds()}

		 // typeof
		,{expected:"undefined",			actual:typeof void 0}
		,{expected:"object",			actual:typeof null}
		,{expected:"boolean",			actual:typeof true}
		,{expected:"number",			actual:typeof 123}
		,{expected:"number",			actual:typeof NaN}
		,{expected:"number",			actual:typeof Infinity}
		,{expected:"string",			actual:typeof "abc"}
		,{expected:"function",			actual:typeof function () {}}
		,{expected:"object",			actual:typeof {}}
		,{expected:"object",			actual:typeof []}
		,{expected:"object",			actual:typeof /abc/}
		,{expected:"object",			actual:typeof Math}
		,{expected:"function",			actual:typeof Math.max}
		,{expected:"function",			actual:typeof parseInt}
		,{expected:"function",			actual:typeof eval}
		,{expected:"object",			actual:(function () { return typeof arguments; })()}
		,{expected:"number",			actual:(function () { return typeof arguments.length; })()}

		 // instanceof
		,{expected:true,				actual:[] instanceof Array}
		,{expected:true,				actual:[] instanceof Object}
		,{expected:true,				actual:({}) instanceof Object}
		,{expected:true,				actual:(new Date()) instanceof Date}
		,{expected:true,				actual:(new Date()) instanceof Object}
		,{expected:true,				actual:(/abc/) instanceof RegExp}
	,{expected:true,				actual:(function () {}) instanceof Function}
	,{expected:true,				actual:Date instanceof Function}
		,{expected:true,				actual:Date instanceof Object}
		,{expected:true,				actual:(new String("x")) instanceof String}
	,{expected:true,				actual:(new Number(1)) instanceof Number}
	,{expected:true,				actual:(new Boolean(false)) instanceof Boolean}
		,{expected:false,				actual:null instanceof Object}
		,{expected:false,				actual:"x" instanceof String}
	,{expected:false,				actual:1 instanceof Number}
	,{expected:false,				actual:true instanceof Boolean}
		,{expected:false,				actual:(function () { return arguments instanceof Array; })()}

		 // Equality / strict equality oddities
		,{expected:true,				actual:null == void 0}
		,{expected:false,				actual:null === void 0}
		,{expected:false,				actual:null == 0}
		,{expected:true,				actual:null >= 0}
		,{expected:true,				actual:null <= 0}
		,{expected:false,				actual:null > 0}
		,{expected:false,				actual:null < 0}
		,{expected:true,				actual:"" == 0}
		,{expected:false,				actual:"" === 0}
		,{expected:true,				actual:"0" == 0}
		,{expected:false,				actual:"0" === 0}
		,{expected:true,				actual:false == 0}
		,{expected:false,				actual:false === 0}
		,{expected:true,				actual:true == 1}
		,{expected:false,				actual:true === 1}
		,{expected:true,				actual:false == "0"}
		,{expected:false,				actual:false == "false"}
		,{expected:true,				actual:true == "1"}
		,{expected:false,				actual:true == "true"}
		,{expected:true,				actual:[] == false}
	,{expected:true,				actual:[] == 0}
	,{expected:true,				actual:[] == ""}
,{expected:true,				actual:[0] == false}
		,{expected:true,				actual:[1] == true}
		,{expected:false,				actual:[2] == true}
	,{expected:true,				actual:[] == ![]}
		,{expected:false,				actual:{} == false}
		,{expected:false,				actual:NaN == NaN}
		,{expected:true,				actual:NaN != NaN}
		,{expected:true,				actual:NaN !== NaN}
		,{expected:true,				actual:0 == -0}
		,{expected:true,				actual:0 === -0}

		 // Boolean coercion
		,{expected:false,				actual:Boolean(false)}
		,{expected:false,				actual:Boolean(0)}
		,{expected:false,				actual:Boolean(-0)}
		,{expected:false,				actual:Boolean("")}
		,{expected:false,				actual:Boolean(null)}
		,{expected:false,				actual:Boolean(void 0)}
		,{expected:false,				actual:Boolean(NaN)}
		,{expected:true,				actual:Boolean(true)}
		,{expected:true,				actual:Boolean(1)}
		,{expected:true,				actual:Boolean(-1)}
		,{expected:true,				actual:Boolean("false")}
		,{expected:true,				actual:Boolean("0")}
		,{expected:true,				actual:Boolean(" ")}
		,{expected:true,				actual:Boolean([])}
		,{expected:true,				actual:Boolean({})}
		,{expected:true,				actual:Boolean(new Boolean(false))}
		,{expected:true,				actual:(new Boolean(false)) ? true : false}
		,{expected:true,				actual:(new Number(0)) ? true : false}
		,{expected:true,				actual:(new String("")) ? true : false}

		 // Number coercion / arithmetic
		,{expected:0,					actual:Number("")}
		,{expected:0,					actual:Number(" ")}
		,{expected:0,					actual:Number(null)}
		,{expected:0,					actual:Number(false)}
		,{expected:1,					actual:Number(true)}
		,{expected:123,					actual:Number("123")}
		,{expected:true,				actual:isNaN(Number(void 0))}
		,{expected:true,				actual:isNaN(Number("abc"))}
		,{expected:0,					actual:+""}
		,{expected:0,					actual:+" "}
		,{expected:1,					actual:+true}
		,{expected:0,					actual:+false}
		,{expected:1,					actual:+[1]}
		,{expected:0,					actual:+[]}
		,{expected:true,				actual:isNaN(+{})}
		,{expected:2,					actual:true + true}
		,{expected:1,					actual:true + false}
		,{expected:0,					actual:false + false}
		,{expected:1,					actual:false + 1}
		,{expected:1,					actual:null + 1}
		,{expected:true,				actual:isNaN((void 0) + 1)}
		,{expected:"number",			actual:typeof ((void 0) + 1)}
		,{expected:"11",				actual:"1" + 1}
		,{expected:"51",				actual:5 + "1"}
		,{expected:"5true",				actual:"5" + true}
		,{expected:"nullx",				actual:null + "x"}
		,{expected:"undefinedx",		 actual:(void 0) + "x"}
		,{expected:3,					actual:"5" - 2}
		,{expected:10,					actual:"5" * "2"}
		,{expected:2.5,					actual:"5" / "2"}
		,{expected:4,					actual:"5" - true}
		,{expected:true,				actual:isNaN("abc" - 1)}

		 // NaN / Infinity
		,{expected:true,				actual:isNaN(NaN)}
		,{expected:true,				actual:isNaN(0 / 0)}
		,{expected:false,				actual:isNaN("")}
		,{expected:false,				actual:isNaN(" ")}
		,{expected:false,				actual:isNaN("123")}
		,{expected:true,				actual:isNaN("abc")}
		,{expected:false,				actual:isFinite(1 / 0)}
		,{expected:false,				actual:isFinite(-1 / 0)}
		,{expected:true,				actual:isFinite(123)}
		,{expected:true,				actual:1 / 0 === Infinity}
		,{expected:true,				actual:-1 / 0 === -Infinity}
		,{expected:true,				actual:1 / -0 === -Infinity}
		,{expected:0,					actual:1 / Infinity}
		,{expected:Infinity,			actual:Infinity + 1}
		,{expected:true,				actual:isNaN(Infinity - Infinity)}
		,{expected:true,				actual:isNaN(Number.NaN)}
		,{expected:true,				actual:Number.POSITIVE_INFINITY === Infinity}
		,{expected:true,				actual:Number.NEGATIVE_INFINITY === -Infinity}

		 // parseInt / parseFloat
		,{expected:10,					actual:parseInt("10", 10)}
		,{expected:10,					actual:parseInt("10px", 10)}
		,{expected:10,					actual:parseInt("010", 10)}
		,{expected:16,					actual:parseInt("0x10")}
		,{expected:true,				actual:isNaN(parseInt("x", 10))}
		,{expected:1.5,					actual:parseFloat("1.5")}
		,{expected:1.5,					actual:parseFloat("1.5px")}
		,{expected:true,				actual:isNaN(parseFloat("px"))}

		 // String coercion
		,{expected:"null",				actual:String(null)}
		,{expected:"undefined",			actual:String(void 0)}
		,{expected:"true",				actual:String(true)}
		,{expected:"false",				actual:String(false)}
		,{expected:"NaN",				actual:String(NaN)}
		,{expected:"Infinity",			actual:String(Infinity)}
		,{expected:"",					actual:String([])}
		,{expected:"1,2",				actual:String([1, 2])}
		,{expected:"[object Object]",	actual:String({})}
		,{expected:"[object Object]",	actual:[] + {}}
	,{expected:"[object Object]",	actual:({}) + []}
		,{expected:"1,23,4",			actual:[1, 2] + [3, 4]}


		 // Wrapper objects
		,{expected:"object",			actual:typeof new String("x")}
		,{expected:"object",			actual:typeof new Number(1)}
		,{expected:"object",			actual:typeof new Boolean(false)}
//☠		,{expected:true,				actual:new String("x") == "x"}
		,{expected:false,				actual:new String("x") === "x"}
//☠=""		,{expected:true,				actual:new Number(1) == 1}
		,{expected:false,				actual:new Number(1) === 1}
//☠		,{expected:true,				actual:new Boolean(false) == false}
		,{expected:false,				actual:new Boolean(false) === false}
//☠		,{expected:true,				actual:new String("") == false}
//☠		,{expected:true,				actual:new Number(0) == false}

		 // Array
		,{expected:0,					actual:[].length}
		,{expected:3,					actual:[1, 2, 3].length}
		,{expected:3,					actual:(new Array(3)).length}
		,{expected:true,				actual:(new Array(3))[0] === void 0}
		,{expected:"undefined",			actual:typeof (new Array(3))[0]}
		,{expected:"1,2,3",				actual:[1, 2, 3].join(",")}
		,{expected:"1--3",				actual:[1, null, 3].join("-")}
		,{expected:"1--3",				actual:[1, void 0, 3].join("-")}
		,{expected:3,					actual:[1, 2].push(3)}
		,{expected:2,					actual:[1, 2].pop()}
		,{expected:"2,3",				actual:[1, 2, 3].slice(1).join(",")}
		,{expected:"3,2,1",				actual:[1, 2, 3].reverse().join(",")}
		,{expected:"10,2",				actual:[10, 2].sort().join(",")}
		,{expected:true,				actual:[].constructor === Array}

		 // String methods
		,{expected:3,					actual:"abc".length}
		,{expected:"a",					actual:"abc".charAt(0)}
		,{expected:"",					actual:"abc".charAt(99)}
		,{expected:97,					actual:"abc".charCodeAt(0)}
		,{expected:true,				actual:isNaN("abc".charCodeAt(99))}
		,{expected:1,					actual:"abc".indexOf("b")}
		,{expected:-1,					actual:"abc".indexOf("z")}
		,{expected:"bc",				actual:"abc".substr(1)}
		,{expected:"b",					actual:"abc".substr(1, 1)}
		,{expected:"b",					actual:"abc".substring(1, 2)}
		,{expected:"ab",				actual:"abc".substring(2, 0)}
		,{expected:2,					actual:"a,b".split(",").length}
		,{expected:"aXc",				actual:"abc".replace("b", "X")}
		,{expected:"aXc",				actual:"abc".replace(/b/, "X")}
		,{expected:"ABC",				actual:"abc".toUpperCase()}
		,{expected:"abc",				actual:"ABC".toLowerCase()}

		 // RegExp
		,{expected:true,				actual:/b/.test("abc")}
		,{expected:false,				actual:/z/.test("abc")}
		,{expected:"b",					actual:"abc".match(/b/)[0]}
		,{expected:null,				actual:"abc".match(/z/)}
		,{expected:"object",			actual:typeof new RegExp("b")}
		,{expected:true,				actual:(new RegExp("b")).test("abc")}

		 // Math
		,{expected:3,					actual:Math.max(1, 2, 3)}
		,{expected:1,					actual:Math.min(1, 2, 3)}
		,{expected:-Infinity,			actual:Math.max()}
		,{expected:Infinity,			actual:Math.min()}
		,{expected:true,				actual:isNaN(Math.max(1, NaN))}
		,{expected:2,					actual:Math.round(1.5)}
		,{expected:1,					actual:Math.floor(1.9)}
		,{expected:2,					actual:Math.ceil(1.1)}
		,{expected:8,					actual:Math.pow(2, 3)}
		,{expected:5,					actual:Math.abs(-5)}
		,{expected:0,					actual:Math.abs(-0)}

		 // Function / eval
		,{expected:3,					actual:eval("1 + 2")}
//☠		,{expected:7,					actual:Function("return 7")()}
//☠		,{expected:"function",			actual:typeof Function("return 1")}

		 // Object.prototype.toString
		,{expected:"[object Object]",	actual:Object.prototype.toString.call({})}
		,{expected:"[object Array]",	actual:Object.prototype.toString.call([])}
		,{expected:"[object Date]",		actual:Object.prototype.toString.call(new Date())}
		,{expected:"[object RegExp]",	actual:Object.prototype.toString.call(/x/)}
		,{expected:"[object String]",	actual:Object.prototype.toString.call(new String("x"))}
		,{expected:"[object Number]",	actual:Object.prototype.toString.call(new Number(1))}
		,{expected:"[object Boolean]",	actual:Object.prototype.toString.call(new Boolean(false))}
		,{expected:true,				actual:({}).constructor === Object}

		 // More Math
		,{expected:"number",	actual:typeof Math.PI}
		,{expected:"number",	actual:typeof Math.E}
		,{expected:"number",	actual:typeof Math.LN2}
		,{expected:"number",	actual:typeof Math.LN10}
		,{expected:"number",	actual:typeof Math.LOG2E}
		,{expected:"number",	actual:typeof Math.LOG10E}
		,{expected:"number",	actual:typeof Math.SQRT1_2}
		,{expected:"number",	actual:typeof Math.SQRT2}

		,{expected:true,		actual:Math.PI > 3}
		,{expected:true,		actual:Math.PI < 4}
		,{expected:true,		actual:Math.E > 2}
		,{expected:true,		actual:Math.E < 3}
		,{expected:true,		actual:Math.SQRT2 === Math.sqrt(2)}

		,{expected:0,			actual:Math.sin(0)}
		,{expected:1,			actual:Math.cos(0)}
		,{expected:0,			actual:Math.tan(0)}
		,{expected:0,			actual:Math.atan(0)}
		,{expected:0,			actual:Math.atan2(0, 1)}
		,{expected:true,		actual:Math.atan2(1, 0) === Math.PI / 2}
		,{expected:true,		actual:Math.atan2(-1, 0) === -Math.PI / 2}

		,{expected:2,			actual:Math.sqrt(4)}
		,{expected:0,			actual:Math.sqrt(0)}
		,{expected:true,		actual:isNaN(Math.sqrt(-1))}
		,{expected:1,			actual:Math.exp(0)}
		,{expected:0,			actual:Math.log(1)}
		,{expected:true,		actual:isNaN(Math.log(-1))}
		,{expected:-Infinity,	actual:Math.log(0)}

		,{expected:1,			actual:Math.abs(1)}
		,{expected:1,			actual:Math.abs(-1)}
		,{expected:0,			actual:Math.abs(0)}
		,{expected:0,			actual:Math.abs(-0)}
		,{expected:0,			actual:Math.abs(null)}
		,{expected:0,			actual:Math.abs("")}
		,{expected:1,			actual:Math.abs("-1")}
		,{expected:true,		actual:isNaN(Math.abs("x"))}

		,{expected:2,			actual:Math.ceil(1.1)}
		,{expected:2,			actual:Math.ceil(1.9)}
		,{expected:-1,			actual:Math.ceil(-1.1)}
		,{expected:-1,			actual:Math.ceil(-1.9)}
		,{expected:1,			actual:Math.floor(1.1)}
		,{expected:1,			actual:Math.floor(1.9)}
		,{expected:-2,			actual:Math.floor(-1.1)}
		,{expected:-2,			actual:Math.floor(-1.9)}

		,{expected:2,			actual:Math.round(1.5)}
		,{expected:1,			actual:Math.round(1.49)}
		,{expected:-1,			actual:Math.round(-1.5)}
		,{expected:-2,			actual:Math.round(-1.51)}
		,{expected:true,		actual:1 / Math.round(-0.5) === -Infinity}
		,{expected:true,		actual:1 / Math.round(-0.1) === -Infinity}
		,{expected:0,			actual:Math.round(0.1)}

		,{expected:3,			actual:Math.max(1, 2, 3)}
		,{expected:3,			actual:Math.max("1", "2", "3")}
		,{expected:3,			actual:Math.max(null, false, 3)}
		,{expected:0,			actual:Math.max(null, false)}
		,{expected:-Infinity,	actual:Math.max()}
		,{expected:true,		actual:isNaN(Math.max(1, void 0))}
		,{expected:true,		actual:isNaN(Math.max(1, "x"))}

		,{expected:1,			actual:Math.min(1, 2, 3)}
		,{expected:1,			actual:Math.min("1", "2", "3")}
		,{expected:0,			actual:Math.min(null, true, 3)}
		,{expected:0,			actual:Math.min(null, false)}
		,{expected:Infinity,	actual:Math.min()}
		,{expected:true,		actual:isNaN(Math.min(1, void 0))}
		,{expected:true,		actual:isNaN(Math.min(1, "x"))}

		,{expected:8,			actual:Math.pow(2, 3)}
		,{expected:1,			actual:Math.pow(2, 0)}
		,{expected:0.5,			actual:Math.pow(2, -1)}
		,{expected:1,			actual:Math.pow(0, 0)}
		,{expected:0,			actual:Math.pow(0, 1)}
		,{expected:Infinity,	actual:Math.pow(0, -1)}
		,{expected:true,		actual:isNaN(Math.pow(-1, 0.5))}

		,{expected:0,			actual:Math.acos(1)}
		,{expected:true,		actual:Math.acos(-1) === Math.PI}
		,{expected:true,		actual:isNaN(Math.acos(2))}
		,{expected:0,			actual:Math.asin(0)}
		,{expected:true,		actual:Math.asin(1) === Math.PI / 2}
		,{expected:true,		actual:isNaN(Math.asin(2))}

		,{expected:"number",	actual:typeof Math.random()}
		,{expected:true,		actual:Math.random() >= 0}
		,{expected:true,		actual:Math.random() < 1}
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
	Response.Write(arrResults.join("\r\n") || "[None; Lucas Guimarães is amazing 😁]");
}
catch (err) {
	Response.Write("Error: #" + err.number + "\r\n" + err.description);
}
%>