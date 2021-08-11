(self.webpackChunkdoc_ops=self.webpackChunkdoc_ops||[]).push([[5664],{3905:function(t,e,n){"use strict";n.d(e,{Zo:function(){return u},kt:function(){return d}});var r=n(7294);function a(t,e,n){return e in t?Object.defineProperty(t,e,{value:n,enumerable:!0,configurable:!0,writable:!0}):t[e]=n,t}function o(t,e){var n=Object.keys(t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(t);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),n.push.apply(n,r)}return n}function c(t){for(var e=1;e<arguments.length;e++){var n=null!=arguments[e]?arguments[e]:{};e%2?o(Object(n),!0).forEach((function(e){a(t,e,n[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(n,e))}))}return t}function i(t,e){if(null==t)return{};var n,r,a=function(t,e){if(null==t)return{};var n,r,a={},o=Object.keys(t);for(r=0;r<o.length;r++)n=o[r],e.indexOf(n)>=0||(a[n]=t[n]);return a}(t,e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(t);for(r=0;r<o.length;r++)n=o[r],e.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(t,n)&&(a[n]=t[n])}return a}var s=r.createContext({}),l=function(t){var e=r.useContext(s),n=e;return t&&(n="function"==typeof t?t(e):c(c({},e),t)),n},u=function(t){var e=l(t.components);return r.createElement(s.Provider,{value:e},t.children)},p={inlineCode:"code",wrapper:function(t){var e=t.children;return r.createElement(r.Fragment,{},e)}},m=r.forwardRef((function(t,e){var n=t.components,a=t.mdxType,o=t.originalType,s=t.parentName,u=i(t,["components","mdxType","originalType","parentName"]),m=l(n),d=a,h=m["".concat(s,".").concat(d)]||m[d]||p[d]||o;return n?r.createElement(h,c(c({ref:e},u),{},{components:n})):r.createElement(h,c({ref:e},u))}));function d(t,e){var n=arguments,a=e&&e.mdxType;if("string"==typeof t||a){var o=n.length,c=new Array(o);c[0]=m;var i={};for(var s in e)hasOwnProperty.call(e,s)&&(i[s]=e[s]);i.originalType=t,i.mdxType="string"==typeof t?t:a,c[1]=i;for(var l=2;l<o;l++)c[l]=n[l];return r.createElement.apply(null,c)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},2079:function(t,e,n){"use strict";n.r(e),n.d(e,{frontMatter:function(){return i},contentTitle:function(){return s},metadata:function(){return l},toc:function(){return u},default:function(){return m}});var r=n(2122),a=n(9756),o=(n(7294),n(3905)),c=["components"],i={},s="Smart Contracts",l={unversionedId:"guide/core_concepts/smart-contracts",id:"guide/core_concepts/smart-contracts",isDocsHomePage:!1,title:"Smart Contracts",description:"What are smart contracts?",source:"@site/docs/guide/core_concepts/smart-contracts.md",sourceDirName:"guide/core_concepts",slug:"/guide/core_concepts/smart-contracts",permalink:"/docs/guide/core_concepts/smart-contracts",editUrl:"https://github.com/iotaledger/chronicle.rs/tree/main/docs/docs/guide/core_concepts/smart-contracts.md",version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"ISCP Documentation",permalink:"/docs/guide/readme"},next:{title:"ISCP",permalink:"/docs/guide/core_concepts/iscp"}},u=[{value:"What are smart contracts?",id:"what-are-smart-contracts",children:[{value:"Applications You Can Trust",id:"applications-you-can-trust",children:[]},{value:"Scalable Smart Contracts",id:"scalable-smart-contracts",children:[]}]}],p={toc:u};function m(t){var e=t.components,n=(0,a.Z)(t,c);return(0,o.kt)("wrapper",(0,r.Z)({},p,n,{components:e,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"smart-contracts"},"Smart Contracts"),(0,o.kt)("h2",{id:"what-are-smart-contracts"},"What are smart contracts?"),(0,o.kt)("p",null,"Smart contracts are software applications that run on a distributed network with multiple validators executing and validating the same code.  This ensures the application behaves as expected, and that there is no tampering in the execution of the program. "),(0,o.kt)("h3",{id:"applications-you-can-trust"},"Applications You Can Trust"),(0,o.kt)("p",null,"As you can be certain that the code being executed is always the same (and will not change), this results in applications you can trust. This allows you to use smart contracts for applications where there was a trust issue. The rules of the smart contract define what the contract can and can not do, making it a decentralized and a predictable decision-maker."),(0,o.kt)("p",null,"Smart contracts are used for all kinds of purposes.  A recurring reason to use a smart contract is to automate certain actions without needing a centralized entity to enforce this specific action. A good example of this is a smart contract that can exchange a certain amount of IOTA tokens for a certain amount of land ownership. The smart contract will accept both the IOTA tokens and the land ownership, and will predictably exchange them between both parties without the risk of one of the parties not delivering on their promise. ",(0,o.kt)("strong",{parentName:"p"},"With a smart contract, code is law"),"."),(0,o.kt)("h3",{id:"scalable-smart-contracts"},"Scalable Smart Contracts"),(0,o.kt)("p",null,"On a public blockchain, anyone willing to pay the fees for deploying a smart contract can deploy one. Once your smart contract has been deployed to the chain you no longer have the option to change it, and you can be assured that your smart contract application will be there as long as that blockchain exists. Smart contracts can communicate with one another, and you can invoke programmed public functions on a smart contract to trigger actions on a smart contract, or address the state of a smart contract."),(0,o.kt)("p",null,"Because smart contracts do not run on a single computer, but on many validators, a network of smart contracts can only process so many smart contracts at once, even if the software has been written optimally. This means smart contracts are expensive to execute, and do not scale well on a single blockchain, often resulting in congested networks and expensive fees for executing functions on smart contracts. ",(0,o.kt)("strong",{parentName:"p"},"By allowing many blockchains executing smart contracts to run in parallel"),", and communicate with one another, the ",(0,o.kt)("strong",{parentName:"p"},"ISCP will solve this scalability problem.")))}m.isMDXComponent=!0}}]);