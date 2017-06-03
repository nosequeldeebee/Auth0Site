$(document).ready(function() {
    var lock = new Auth0Lock(AUTH0_CLIENT_ID, AUTH0_DOMAIN, { auth: {
        redirectUrl: AUTH0_CALLBACK_URL
      }, theme : { logo : 'https://cdn4.iconfinder.com/data/icons/medical-14/512/7-128.png' }, languageDictionary : { title : "Sign In" }});

    $('.btn-login').click(function(e) {
      e.preventDefault();
      lock.show();
    });
});
