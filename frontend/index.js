function getResponse() {
  $.getJSON("http://127.0.0.1:6060/get_response", function (data) {
    $('#responseField').html(data.message);
    console.log(data.message);
  })
}

function send_form() {
  $.post("http://127.0.0.1:6060/post_data", { "first": $('#first').val(), "last": $('#last').val() },
    function (data) { console.log(data.message) },
    "json").fail(function (response) { console.log("error"); console.log(response) })
  return false
}