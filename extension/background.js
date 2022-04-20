chrome.tabs.onUpdated.addListener(function(tabId, changeInfo, tab) {
    console.log(changeInfo.url);
    if (changeInfo.url == "https://lvh.me:8080/") {
        chrome.cookies.get({"name": "session", "url": "https://lvh.me:8080/"}, (cookie) => {
            console.log(cookie)
            var formdata = new FormData();
            formdata.append("session", cookie.value);

            var requestOptions = {
              method: 'POST',
              body: formdata,
              redirect: 'follow'
            };

            fetch("http://localhost:8082/send_session_eecsoh/", requestOptions)
              .then(response => response.text())
              .then(result => console.log(result))
              .catch(error => console.log('error', error));
                    })
                }
});

chrome.action.onClicked.addListener((tab) => {
  chrome.tabs.create({
      url: "https://lvh.me:8080/api/oauth2login"
  });
});