/// <reference path="../typings/main.d.ts" />

import * as React from 'react';
import * as ReactDOM from 'react-dom';

interface MailingList {
    /*
    _id: string;
    legacyIDs: Array<number>;
    welcomeTemplate: string;
    account: string;
    list: string;
    */    
}
class MailingLists extends React.Component<MailingList, any> {
    constructor(props: MailingList) {
        super(props);
    }
    
    render() {
        return <div>Hello</div>;
    }
}

ReactDOM.render(
    <MailingLists />,
    document.getElementById('content')
)