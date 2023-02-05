function postAction(action) {
  $.ajax(
    {
      type: "POST",
      async: false,
      url: "/action",
      data: {
        action: action
      },
      success: function(result) {
        console.log(result)
      },
      error: function(XMLHttpRequest, textStatus, errorThrown){
        console.log(errorThrown)
      },
      dataType: "json"
    }
  );
}

$('#led-on').click(function(){postAction("ON")});

$('#led-off').click(function(){postAction("OFF")});