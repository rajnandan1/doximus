import{S as U,i as z,s as G,e as p,j,k as E,c as b,a as w,m as V,d as m,n as k,b as J,f as B,H as v,o as P,x as A,u as F,v as q}from"../../chunks/vendor-d32ee1dd.js";import{e as C}from"../../chunks/variables-abb00116.js";import{S as K,F as L,a as Q,P as R}from"../../chunks/footer-c0b2c89f.js";import"../../chunks/singletons-12a22614.js";function T(r){let a,c,l,O,o,f,t,s,n,h,S,u,y,D,d,I;return l=new K({props:{searchJSON:r[3],pageID:r[4],siteJSON:r[2],contentJSON:r[1]}}),o=new L({}),u=new Q({props:{searchJSON:r[3]}}),d=new R({props:{mainJSON:r[5],curlJSON:r[0],contentJSON:r[1]}}),{c(){a=p("nav"),c=p("div"),j(l.$$.fragment),O=E(),j(o.$$.fragment),f=E(),t=p("main"),s=p("div"),n=E(),h=p("div"),S=p("div"),j(u.$$.fragment),y=E(),D=p("div"),j(d.$$.fragment),this.h()},l(e){a=b(e,"NAV",{id:!0,class:!0});var i=w(a);c=b(i,"DIV",{class:!0});var g=w(c);V(l.$$.fragment,g),g.forEach(m),O=k(i),V(o.$$.fragment,i),i.forEach(m),f=k(e),t=b(e,"MAIN",{class:!0});var N=w(t);s=b(N,"DIV",{id:!0,class:!0}),w(s).forEach(m),n=k(N),h=b(N,"DIV",{class:!0});var _=w(h);S=b(_,"DIV",{class:!0});var H=w(S);V(u.$$.fragment,H),H.forEach(m),_.forEach(m),y=k(N),D=b(N,"DIV",{class:!0});var M=w(D);V(d.$$.fragment,M),M.forEach(m),N.forEach(m),this.h()},h(){J(c,"class","sidebar-content"),J(a,"id","sidebar"),J(a,"class","sidebar-wrapper"),J(s,"id","overlay"),J(s,"class","overlay"),J(S,"class","row d-flex align-items-center p-3"),J(h,"class","container-fluid"),J(D,"class","row p-lg-4"),J(t,"class","page-content")},m(e,i){B(e,a,i),v(a,c),P(l,c,null),v(a,O),P(o,a,null),B(e,f,i),B(e,t,i),v(t,s),v(t,n),v(t,h),v(h,S),P(u,S,null),v(t,y),v(t,D),P(d,D,null),I=!0},p(e,[i]){const g={};i&8&&(g.searchJSON=e[3]),i&16&&(g.pageID=e[4]),i&4&&(g.siteJSON=e[2]),i&2&&(g.contentJSON=e[1]),l.$set(g);const N={};i&8&&(N.searchJSON=e[3]),u.$set(N);const _={};i&32&&(_.mainJSON=e[5]),i&1&&(_.curlJSON=e[0]),i&2&&(_.contentJSON=e[1]),d.$set(_)},i(e){I||(A(l.$$.fragment,e),A(o.$$.fragment,e),A(u.$$.fragment,e),A(d.$$.fragment,e),I=!0)},o(e){F(l.$$.fragment,e),F(o.$$.fragment,e),F(u.$$.fragment,e),F(d.$$.fragment,e),I=!1},d(e){e&&m(a),q(l),q(o),e&&m(f),e&&m(t),q(u),q(d)}}}async function x({page:r,fetch:a,session:c,stuff:l}){let O=l.siteJSON,o=l.searchJSON,f="docs";const t=r.path;r.path!=""&&r.path!="/"&&(f=t.replace(/\//g,".").substr(1));let s=await a(C.api+"files/"+f+".json");const n=await s.json();s=await a(C.api+"files/"+f+".curl.json");const h=await s.json();s=await a(C.api+"files/"+f+".content.txt");let S="";try{S=await s.json()}catch{}return{props:{siteJSON:O,searchJSON:o,pageID:r.path,mainJSON:n,curlJSON:h,contentJSON:S}}}function W(r,a,c){let{siteJSON:l}=a,{searchJSON:O}=a,{pageID:o}=a,{mainJSON:f}=a,{curlJSON:t}=a,{contentJSON:s}=a;return t||(t={apis:[]}),typeof s!="string"&&(s=""),r.$$set=n=>{"siteJSON"in n&&c(2,l=n.siteJSON),"searchJSON"in n&&c(3,O=n.searchJSON),"pageID"in n&&c(4,o=n.pageID),"mainJSON"in n&&c(5,f=n.mainJSON),"curlJSON"in n&&c(0,t=n.curlJSON),"contentJSON"in n&&c(1,s=n.contentJSON)},[t,s,l,O,o,f]}class ee extends U{constructor(a){super();z(this,a,W,T,G,{siteJSON:2,searchJSON:3,pageID:4,mainJSON:5,curlJSON:0,contentJSON:1})}}export{ee as default,x as load};
