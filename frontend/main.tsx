/// <reference path="../typings/main.d.ts" />

import * as React from 'react';
import * as ReactDOM from 'react-dom';

interface BlogPost {
    /*
    _id: string;
    legacyIDs: Array<number>;
    welcomeTemplate: string;
    account: string;
    list: string;
    */    
}
class BlogPosts extends React.Component<BlogPost, any> {
    constructor(props: BlogPost) {
        super(props);
    }
    
    render() {
        return <div>Hello</div>;
    }
}

ReactDOM.render(
    <BlogPosts />,
    document.getElementById('content')
)