//attach jQuery to html
const form = $("#user-input");
const list = $("#conversation_list");

form.keypress(function(event){
    if(event.keyCode != 13){ // ENTER
        return;
    }
    
    event.preventDefault(); // stops it from refreshing the page.

    const userText = form.val(); // grabs users input
    form.val(" "); // Clears the text box
    
    // Makes sure text box isn't empty
    list.append("<figure><figcaption>User</figcaption></figure><li  class='list-group-item  text-left list-group-item-primary' id='List'>"+userText + "</li>");

    // GET/POST
    const queryParams = {"user-input" : userText }
    $.get("/chat", queryParams)
    
        .done(function(resp){
            const newItem = "<figure><figcaption>Eliza</figcaption></figure><li  class='list-group-item  text-left list-group-item-danger' id='List'>"+ resp + "</li>";
            setTimeout(function(){
              
                list.append(newItem)
                $("html, body").scrollTop($("body").height());
            }, 2000);//wait to response
        }).fail(function(){
           // display error message if connection fails
            const newItem = "<figure><figcaption>Eliza</figcaption></figure><li  class='list-group-item  text-left list-group-item-danger' id='List'>Currently not working, try again later</li>";
            list.append(newItem);
        });
});