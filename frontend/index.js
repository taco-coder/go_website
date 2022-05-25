function getResponse() {
  $.getJSON("/get_response", function (data) {
    $('#responseField').html(data.message);
    console.log(data.message);
  })
}

function send_form() {
  $.post("/post_data", { "first": $('#first').val(), "last": $('#last').val() },
    function (data) { console.log(data.message) },
    "json").fail(function (response) { console.log("error"); console.log(response) })
  return false
}