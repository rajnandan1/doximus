import{S as M,i as U,s as z,e as v,j as w,k as j,c as _,a as p,m as E,d as m,n as V,b as J,f as k,H as D,o as P,x as A,u as q,v as y}from"../../chunks/vendor-3104456f.js";import{e as B}from"../../chunks/variables-abb00116.js";import{S as F,a as G,P as K}from"../../chunks/page-d5c3e31c.js";import"../../chunks/singletons-12a22614.js";function L(t){let a,c,l,u,i,o,f,n,s,S,I,d,O,b;return l=new F({props:{searchJSON:t[3],pageID:t[4],siteJSON:t[2],contentJSON:t[1]}}),S=new G({props:{searchJSON:t[3]}}),O=new K({props:{mainJSON:t[5],curlJSON:t[0],contentJSON:t[1]}}),{c(){a=v("nav"),c=v("div"),w(l.$$.fragment),u=j(),i=v("main"),o=v("div"),f=j(),n=v("div"),s=v("div"),w(S.$$.fragment),I=j(),d=v("div"),w(O.$$.fragment),this.h()},l(e){a=_(e,"NAV",{id:!0,class:!0});var r=p(a);c=_(r,"DIV",{class:!0});var h=p(c);E(l.$$.fragment,h),h.forEach(m),r.forEach(m),u=V(e),i=_(e,"MAIN",{class:!0});var N=p(i);o=_(N,"DIV",{id:!0,class:!0}),p(o).forEach(m),f=V(N),n=_(N,"DIV",{class:!0});var g=p(n);s=_(g,"DIV",{class:!0});var C=p(s);E(S.$$.fragment,C),C.forEach(m),g.forEach(m),I=V(N),d=_(N,"DIV",{class:!0});var H=p(d);E(O.$$.fragment,H),H.forEach(m),N.forEach(m),this.h()},h(){J(c,"class","sidebar-content"),J(a,"id","sidebar"),J(a,"class","sidebar-wrapper"),J(o,"id","overlay"),J(o,"class","overlay"),J(s,"class","row d-flex align-items-center p-3"),J(n,"class","container-fluid"),J(d,"class","row p-lg-4"),J(i,"class","page-content")},m(e,r){k(e,a,r),D(a,c),P(l,c,null),k(e,u,r),k(e,i,r),D(i,o),D(i,f),D(i,n),D(n,s),P(S,s,null),D(i,I),D(i,d),P(O,d,null),b=!0},p(e,[r]){const h={};r&8&&(h.searchJSON=e[3]),r&16&&(h.pageID=e[4]),r&4&&(h.siteJSON=e[2]),r&2&&(h.contentJSON=e[1]),l.$set(h);const N={};r&8&&(N.searchJSON=e[3]),S.$set(N);const g={};r&32&&(g.mainJSON=e[5]),r&1&&(g.curlJSON=e[0]),r&2&&(g.contentJSON=e[1]),O.$set(g)},i(e){b||(A(l.$$.fragment,e),A(S.$$.fragment,e),A(O.$$.fragment,e),b=!0)},o(e){q(l.$$.fragment,e),q(S.$$.fragment,e),q(O.$$.fragment,e),b=!1},d(e){e&&m(a),y(l),e&&m(u),e&&m(i),y(S),y(O)}}}async function Y({page:t,fetch:a,session:c,stuff:l}){let u=l.siteJSON,i=l.searchJSON,o="docs";const f=t.path;t.path!=""&&t.path!="/"&&(o=f.replace(/\//g,".").substr(1));let n=await a(B.api+"files/"+o+".json");const s=await n.json();n=await a(B.api+"files/"+o+".curl.json");const S=await n.json();n=await a(B.api+"files/"+o+".content.txt");let I="";try{I=await n.json()}catch{}return{props:{siteJSON:u,searchJSON:i,pageID:t.path,mainJSON:s,curlJSON:S,contentJSON:I}}}function Q(t,a,c){let{siteJSON:l}=a,{searchJSON:u}=a,{pageID:i}=a,{mainJSON:o}=a,{curlJSON:f}=a,{contentJSON:n}=a;return f||(f={apis:[]}),typeof n!="string"&&(n=""),t.$$set=s=>{"siteJSON"in s&&c(2,l=s.siteJSON),"searchJSON"in s&&c(3,u=s.searchJSON),"pageID"in s&&c(4,i=s.pageID),"mainJSON"in s&&c(5,o=s.mainJSON),"curlJSON"in s&&c(0,f=s.curlJSON),"contentJSON"in s&&c(1,n=s.contentJSON)},[f,n,l,u,i,o]}class Z extends M{constructor(a){super();U(this,a,Q,L,z,{siteJSON:2,searchJSON:3,pageID:4,mainJSON:5,curlJSON:0,contentJSON:1})}}export{Z as default,Y as load};