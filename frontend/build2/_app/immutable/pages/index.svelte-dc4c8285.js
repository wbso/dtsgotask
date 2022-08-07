import{a1 as J,a2 as ee,S as X,i as Y,s as Z,l as _,J as A,a as $,r as U,m as g,n as v,K as L,h as f,c as D,u as j,p as a,q as W,b as P,N as r,U as z,v as F,E as H,G as te,O as se,w as re,a3 as oe,x as ae,L as ne,y as le,f as ie,t as ce,B as de}from"../chunks/index-a1aa545d.js";import{w as ue}from"../chunks/index-6d579add.js";function Q(e){return Object.prototype.toString.call(e)==="[object Date]"}function K(e,s,t,o){if(typeof t=="number"||Q(t)){const n=o-t,l=(t-s)/(e.dt||1/60),c=e.opts.stiffness*n,i=e.opts.damping*l,h=(c-i)*e.inv_mass,u=(l+h)*e.dt;return Math.abs(u)<e.opts.precision&&Math.abs(n)<e.opts.precision?o:(e.settled=!1,Q(t)?new Date(t.getTime()+u):t+u)}else{if(Array.isArray(t))return t.map((n,l)=>K(e,s[l],t[l],o[l]));if(typeof t=="object"){const n={};for(const l in t)n[l]=K(e,s[l],t[l],o[l]);return n}else throw new Error(`Cannot spring ${typeof t} values`)}}function fe(e,s={}){const t=ue(e),{stiffness:o=.15,damping:n=.8,precision:l=.01}=s;let c,i,h,u=e,S=e,T=1,M=0,E=!1;function O(m,p={}){S=m;const k=h={};if(e==null||p.hard||y.stiffness>=1&&y.damping>=1)return E=!0,c=J(),u=m,t.set(e=S),Promise.resolve();if(p.soft){const w=p.soft===!0?.5:+p.soft;M=1/(w*60),T=0}return i||(c=J(),E=!1,i=ee(w=>{if(E)return E=!1,i=null,!1;T=Math.min(T+M,1);const x={inv_mass:T,opts:y,settled:!0,dt:(w-c)*60/1e3},b=K(x,u,e,S);return c=w,u=e,t.set(e=b),x.settled&&(i=null),!x.settled})),new Promise(w=>{i.promise.then(()=>{k===h&&w()})})}const y={set:O,update:(m,p)=>O(m(S,e),p),subscribe:t.subscribe,stiffness:o,damping:n,precision:l};return y}function pe(e){let s,t,o,n,l,c,i,h,u=Math.floor(e[1]+1)+"",S,T,M,E=Math.floor(e[1])+"",O,y,m,p,k,w,x;return{c(){s=_("div"),t=_("button"),o=A("svg"),n=A("path"),l=$(),c=_("div"),i=_("div"),h=_("strong"),S=U(u),T=$(),M=_("strong"),O=U(E),y=$(),m=_("button"),p=A("svg"),k=A("path"),this.h()},l(b){s=g(b,"DIV",{class:!0});var d=v(s);t=g(d,"BUTTON",{"aria-label":!0,class:!0});var N=v(t);o=L(N,"svg",{"aria-hidden":!0,viewBox:!0,class:!0});var B=v(o);n=L(B,"path",{d:!0,class:!0}),v(n).forEach(f),B.forEach(f),N.forEach(f),l=D(d),c=g(d,"DIV",{class:!0});var C=v(c);i=g(C,"DIV",{class:!0,style:!0});var I=v(i);h=g(I,"STRONG",{class:!0,"aria-hidden":!0});var R=v(h);S=j(R,u),R.forEach(f),T=D(I),M=g(I,"STRONG",{class:!0});var G=v(M);O=j(G,E),G.forEach(f),I.forEach(f),C.forEach(f),y=D(d),m=g(d,"BUTTON",{"aria-label":!0,class:!0});var q=v(m);p=L(q,"svg",{"aria-hidden":!0,viewBox:!0,class:!0});var V=v(p);k=L(V,"path",{d:!0,class:!0}),v(k).forEach(f),V.forEach(f),q.forEach(f),d.forEach(f),this.h()},h(){a(n,"d","M0,0.5 L1,0.5"),a(n,"class","svelte-sx9eo4"),a(o,"aria-hidden","true"),a(o,"viewBox","0 0 1 1"),a(o,"class","svelte-sx9eo4"),a(t,"aria-label","Decrease the counter by one"),a(t,"class","svelte-sx9eo4"),a(h,"class","hidden svelte-sx9eo4"),a(h,"aria-hidden","true"),a(M,"class","svelte-sx9eo4"),a(i,"class","counter-digits svelte-sx9eo4"),W(i,"transform","translate(0, "+100*e[2]+"%)"),a(c,"class","counter-viewport svelte-sx9eo4"),a(k,"d","M0,0.5 L1,0.5 M0.5,0 L0.5,1"),a(k,"class","svelte-sx9eo4"),a(p,"aria-hidden","true"),a(p,"viewBox","0 0 1 1"),a(p,"class","svelte-sx9eo4"),a(m,"aria-label","Increase the counter by one"),a(m,"class","svelte-sx9eo4"),a(s,"class","counter svelte-sx9eo4")},m(b,d){P(b,s,d),r(s,t),r(t,o),r(o,n),r(s,l),r(s,c),r(c,i),r(i,h),r(h,S),r(i,T),r(i,M),r(M,O),r(s,y),r(s,m),r(m,p),r(p,k),w||(x=[z(t,"click",e[4]),z(m,"click",e[5])],w=!0)},p(b,[d]){d&2&&u!==(u=Math.floor(b[1]+1)+"")&&F(S,u),d&2&&E!==(E=Math.floor(b[1])+"")&&F(O,E),d&4&&W(i,"transform","translate(0, "+100*b[2]+"%)")},i:H,o:H,d(b){b&&f(s),w=!1,te(x)}}}function he(e,s){return(e%s+s)%s}function me(e,s,t){let o,n,l=0;const c=fe();se(e,c,u=>t(1,n=u));const i=()=>t(0,l-=1),h=()=>t(0,l+=1);return e.$$.update=()=>{e.$$.dirty&1&&c.set(l),e.$$.dirty&2&&t(2,o=he(n,1))},[l,n,o,c,i,h]}class ve extends X{constructor(s){super(),Y(this,s,me,pe,Z,{})}}function _e(e){let s,t,o,n,l,c,i,h,u,S,T,M,E,O,y,m,p,k,w,x,b;return x=new ve({}),{c(){s=_("meta"),t=$(),o=_("section"),n=_("h1"),l=_("span"),c=_("picture"),i=_("source"),h=$(),u=_("img"),T=U(`

		to your new`),M=_("br"),E=U("SvelteKit app"),O=$(),y=_("h2"),m=U("try editing "),p=_("strong"),k=U("src/routes/index.svelte"),w=$(),re(x.$$.fragment),this.h()},l(d){const N=oe('[data-svelte="svelte-t32ptj"]',document.head);s=g(N,"META",{name:!0,content:!0}),N.forEach(f),t=D(d),o=g(d,"SECTION",{class:!0});var B=v(o);n=g(B,"H1",{class:!0});var C=v(n);l=g(C,"SPAN",{class:!0});var I=v(l);c=g(I,"PICTURE",{});var R=v(c);i=g(R,"SOURCE",{srcset:!0,type:!0}),h=D(R),u=g(R,"IMG",{src:!0,alt:!0,class:!0}),R.forEach(f),I.forEach(f),T=j(C,`

		to your new`),M=g(C,"BR",{}),E=j(C,"SvelteKit app"),C.forEach(f),O=D(B),y=g(B,"H2",{});var G=v(y);m=j(G,"try editing "),p=g(G,"STRONG",{});var q=v(p);k=j(q,"src/routes/index.svelte"),q.forEach(f),G.forEach(f),w=D(B),ae(x.$$.fragment,B),B.forEach(f),this.h()},h(){document.title="Home",a(s,"name","description"),a(s,"content","Svelte demo app"),a(i,"srcset","svelte-welcome.webp"),a(i,"type","image/webp"),ne(u.src,S="svelte-welcome.png")||a(u,"src",S),a(u,"alt","Welcome"),a(u,"class","svelte-1egtvge"),a(l,"class","welcome svelte-1egtvge"),a(n,"class","svelte-1egtvge"),a(o,"class","svelte-1egtvge")},m(d,N){r(document.head,s),P(d,t,N),P(d,o,N),r(o,n),r(n,l),r(l,c),r(c,i),r(c,h),r(c,u),r(n,T),r(n,M),r(n,E),r(o,O),r(o,y),r(y,m),r(y,p),r(p,k),r(o,w),le(x,o,null),b=!0},p:H,i(d){b||(ie(x.$$.fragment,d),b=!0)},o(d){ce(x.$$.fragment,d),b=!1},d(d){f(s),d&&f(t),d&&f(o),de(x)}}}class ye extends X{constructor(s){super(),Y(this,s,null,_e,Z,{})}}export{ye as default};
