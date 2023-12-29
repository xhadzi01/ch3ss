function IsReadyToProceed()
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", "/", false ); // false for synchronous request
    xmlHttp.send( null );
    return xmlHttp.responseText;
}