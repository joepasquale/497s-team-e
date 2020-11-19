import React from 'react';

import { GoogleLogout } from 'react-google-login';

const clientId =
    '350429252210-hq617ss9idkeat0h66hbop59ul53mpnf.apps.googleusercontent.com';

function Logout() {
    const onSuccess = () => {
        console.log('Logout made successfully');
        alert('Logout made successfully ✌');
    };

    return (
        <div>
            <GoogleLogout
                clientId={clientId}
                buttonText="Logout"
                onLogoutSuccess={onSuccess}
            ></GoogleLogout>
        </div>
    );
}

export default Logout;