<%@ Language="JScript" %>
<%
class Test {
	#x = 10;
	static #y = 20;

	get() {
		return this.#x;
	}
	
	set(val) {
		this.#x = val;
	}

	inc() {
		this.#x++;
		return this.#x;
	}
	
	static getStatic() {
		return Test.#y;
	}

	static setStatic(val) {
		Test.#y = val;
	}
}

var t = new Test();
Response.Write("x: " + t.get() + " | ");
t.set(15);
Response.Write("set x: " + t.get() + " | ");
Response.Write("inc x: " + t.inc() + " | ");
Response.Write("static y: " + Test.getStatic() + " | ");
Test.setStatic(25);
Response.Write("set static y: " + Test.getStatic());
%>