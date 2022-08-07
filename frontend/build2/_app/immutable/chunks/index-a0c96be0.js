import{H as w}from"./index-a1aa545d.js";function m(a){const t=a-1;return t*t*t+1}function C(a,{from:t,to:s},f={}){const c=getComputedStyle(a),l=c.transform==="none"?"":c.transform,[o,n]=c.transformOrigin.split(" ").map(parseFloat),r=t.left+t.width*o/s.width-(s.left+o),e=t.top+t.height*n/s.height-(s.top+n),{delay:y=0,duration:i=d=>Math.sqrt(d)*120,easing:p=m}=f;return{delay:y,duration:w(i)?i(Math.sqrt(r*r+e*e)):i,easing:p,css:(d,h)=>{const u=h*r,$=h*e,g=d+h*t.width/s.width,x=d+h*t.height/s.height;return`transform: ${l} translate(${u}px, ${$}px) scale(${g}, ${x});`}}}function S(a,{delay:t=0,duration:s=400,easing:f=m,x:c=0,y:l=0,opacity:o=0}={}){const n=getComputedStyle(a),r=+n.opacity,e=n.transform==="none"?"":n.transform,y=r*(1-o);return{delay:t,duration:s,easing:f,css:(i,p)=>`
			transform: ${e} translate(${(1-i)*c}px, ${(1-i)*l}px);
			opacity: ${r-y*p}`}}function q(a,{delay:t=0,duration:s=400,easing:f=m,start:c=0,opacity:l=0}={}){const o=getComputedStyle(a),n=+o.opacity,r=o.transform==="none"?"":o.transform,e=1-c,y=n*(1-l);return{delay:t,duration:s,easing:f,css:(i,p)=>`
			transform: ${r} scale(${1-e*p});
			opacity: ${n-y*p}
		`}}export{S as a,C as f,q as s};
