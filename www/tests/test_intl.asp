<script runat="server" language="JScript">
    function report(name, ok, actual) {
        Response.Write(name + ":" + (ok ? "PASS" : "FAIL"));
        if (!ok && actual !== undefined) {
            Response.Write(":" + actual);
        }
        Response.Write("\n");
    }

    var dateValue = new Date(Date.UTC(2026, 0, 2, 3, 4, 5));
    var enDate = new Intl.DateTimeFormat("en-US", { dateStyle: "short" }).format(dateValue);
    var ptDate = new Intl.DateTimeFormat("pt-BR", { dateStyle: "short" }).format(dateValue);
    var deDate = new Intl.DateTimeFormat("de-DE", { dateStyle: "short" }).format(dateValue);
    report("DateTimeFormatEnUS", enDate === "1/2/2026", enDate);
    report("DateTimeFormatPtBR", ptDate === "02/01/2026", ptDate);
    report("DateTimeFormatDeDE", deDate === "02.01.2026", deDate);

    var numberValue = 1234567.89;
    var enNumber = new Intl.NumberFormat("en-US", { style: "decimal", maximumFractionDigits: 2 }).format(numberValue);
    var ptNumber = new Intl.NumberFormat("pt-BR", { style: "decimal", maximumFractionDigits: 2 }).format(numberValue);
    var deCurrency = new Intl.NumberFormat("de-DE", { style: "currency", currency: "EUR", maximumFractionDigits: 2 }).format(numberValue);
    report("NumberFormatEnUS", enNumber === "1,234,567.89", enNumber);
    report("NumberFormatPtBR", ptNumber === "1.234.567,89", ptNumber);
    report("NumberFormatDeDECurrency", deCurrency === "€ 1.234.567,89", deCurrency);
</script>