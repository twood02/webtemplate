
function checkCP(event) {
if (window.location.port != 4000) {
    alert("No copy paste for you!");
    event.preventDefault();
  }
}
