(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{bUiu:function(n,l,t){"use strict";t.r(l);var e=t("CcnG"),u=function(){function n(){}return Object.defineProperty(n.prototype,"step",{get:function(){return this._step},set:function(n){console.log("@Input step: "+n),this._step=n},enumerable:!0,configurable:!0}),n.prototype.ngOnInit=function(){console.log("step one init")},n.prototype.ngOnDestroy=function(){console.log("step one destroy")},n}(),b=function(){function n(){}return Object.defineProperty(n.prototype,"step",{get:function(){return this._step},set:function(n){console.log("@Input step: "+n),this._step=n},enumerable:!0,configurable:!0}),n.prototype.ngOnInit=function(){console.log("step one init")},n.prototype.ngOnDestroy=function(){console.log("step one destroy")},n}(),o=function(){function n(){}return Object.defineProperty(n.prototype,"step",{get:function(){return this._step},set:function(n){console.log("@Input step: "+n),this._step=n},enumerable:!0,configurable:!0}),n.prototype.ngOnInit=function(){console.log("step one init")},n.prototype.ngOnDestroy=function(){console.log("step one destroy")},n}(),p=function(){function n(n){this.route=n}return n.prototype.ngOnInit=function(){var n=this;this.route.queryParams.subscribe(function(l){var t=l.step||"step1";n.stepComp={step1:{component:u,inputs:{step:t}},step2:{component:b,inputs:{step:t}},step3:{component:o,inputs:{step:t}}}[t]})},n}(),s=function(){return function(){}}(),r=t("ebDo"),c=t("Z24s"),i=t("pMnS"),C=function(){function n(n,l){this.vcr=n,this.cfr=l}return Object.defineProperty(n.prototype,"data",{set:function(n){var l=this.cfr.resolveComponentFactory(n.component),t=this.vcr.createComponent(l);if(n.inputs)for(var e in n.inputs)e&&(t.instance[e]=n.inputs[e]);this.destroy(),this.currentComponent=t},enumerable:!0,configurable:!0}),n.prototype.destroy=function(){this.currentComponent&&(this.currentComponent.destroy(),this.currentComponent=null)},n.prototype.ngOnDestroy=function(){this.destroy()},n}(),a=e.sb({encapsulation:2,styles:[],data:{}});function d(n){return e.Ob(0,[],null,null)}var f=t("ZYCi"),m=e.sb({encapsulation:0,styles:[[""]],data:{}});function g(n){return e.Ob(0,[(n()(),e.ub(0,0,null,null,1,"h1",[],null,null,null,null,null)),(n()(),e.Mb(-1,null,["Angular\u52a8\u6001\u52a0\u8f7d\u7ec4\u4ef6"])),(n()(),e.ub(2,0,null,null,1,"a",[["href","https://segmentfault.com/a/1190000009582289"]],null,null,null,null,null)),(n()(),e.Mb(-1,null,["https://segmentfault.com/a/1190000009582289"])),(n()(),e.ub(4,16777216,null,null,1,"app-step",[],null,null,null,d,a)),e.tb(5,180224,null,0,C,[e.S,e.j],{data:[0,"data"]},null),(n()(),e.ub(6,0,null,null,1,"a",[["href","#/admin/step?step=step1"]],null,null,null,null,null)),(n()(),e.Mb(-1,null,["step1"])),(n()(),e.ub(8,0,null,null,0,"br",[],null,null,null,null,null)),(n()(),e.ub(9,0,null,null,1,"a",[["href","#/admin/step?step=step2"]],null,null,null,null,null)),(n()(),e.Mb(-1,null,["step2"])),(n()(),e.ub(11,0,null,null,0,"br",[],null,null,null,null,null)),(n()(),e.ub(12,0,null,null,1,"a",[["href","#/admin/step?step=step3"]],null,null,null,null,null)),(n()(),e.Mb(-1,null,["step3"]))],function(n,l){n(l,5,0,l.component.stepComp)},null)}function y(n){return e.Ob(0,[(n()(),e.ub(0,0,null,null,1,"app-index",[],null,null,null,g,m)),e.tb(1,114688,null,0,p,[f.a],null,null)],function(n,l){n(l,1,0)},null)}var h=e.qb("app-index",p,y,{},{},[]),O=e.sb({encapsulation:0,styles:[[""]],data:{}});function v(n){return e.Ob(0,[(n()(),e.ub(0,0,null,null,1,"h2",[],null,null,null,null,null)),(n()(),e.Mb(1,null,["Step One Component\uff1aparams value: ",""]))],null,function(n,l){n(l,1,0,l.component.step)})}function k(n){return e.Ob(0,[(n()(),e.ub(0,0,null,null,1,"app-user1",[],null,null,null,v,O)),e.tb(1,245760,null,0,u,[],null,null)],function(n,l){n(l,1,0)},null)}var D=e.qb("app-user1",u,k,{step:"step"},{},[]),M=e.sb({encapsulation:0,styles:[[""]],data:{}});function S(n){return e.Ob(0,[(n()(),e.ub(0,0,null,null,1,"h2",[],null,null,null,null,null)),(n()(),e.Mb(1,null,["Step One Component\uff1aparams value: ",""]))],null,function(n,l){n(l,1,0,l.component.step)})}function P(n){return e.Ob(0,[(n()(),e.ub(0,0,null,null,1,"app-user2",[],null,null,null,S,M)),e.tb(1,245760,null,0,b,[],null,null)],function(n,l){n(l,1,0)},null)}var z=e.qb("app-user2",b,P,{step:"step"},{},[]),j=e.sb({encapsulation:0,styles:[[""]],data:{}});function w(n){return e.Ob(0,[(n()(),e.ub(0,0,null,null,1,"h2",[],null,null,null,null,null)),(n()(),e.Mb(1,null,["Step One Component\uff1aparams value: ",""]))],null,function(n,l){n(l,1,0,l.component.step)})}function q(n){return e.Ob(0,[(n()(),e.ub(0,0,null,null,1,"app-user3",[],null,null,null,w,j)),e.tb(1,245760,null,0,o,[],null,null)],function(n,l){n(l,1,0)},null)}var x=e.qb("app-user3",o,q,{step:"step"},{},[]),F=t("Ip0R"),I=t("gIcY"),A=t("t/Na"),N=t("0+vK"),B=t("5NvZ"),R=t("M2Lx"),T=t("6Cds"),J=t("eDkP"),Z=t("Fzqc"),_=t("dWZg"),V=t("4c35"),X=t("qAlS"),E=t("ADsi"),H=function(){function n(){}return n.withComponents=function(l){return{ngModule:n,providers:[{provide:e.a,useValue:l,multi:!0}]}},n}();t.d(l,"StepModuleNgFactory",function(){return L});var L=e.rb(s,[],function(n){return e.Bb([e.Cb(512,e.j,e.fb,[[8,[r.M,r.N,r.O,r.P,r.Q,r.R,r.S,r.T,c.a,i.a,h,D,z,x]],[3,e.j],e.y]),e.Cb(4608,F.o,F.n,[e.v,[2,F.B]]),e.Cb(4608,I.s,I.s,[]),e.Cb(4608,I.d,I.d,[]),e.Cb(4608,A.l,A.r,[F.d,e.C,A.p]),e.Cb(4608,A.s,A.s,[A.l,A.q]),e.Cb(5120,A.a,function(n,l){return[n,new N.a(l)]},[A.s,B.a]),e.Cb(4608,A.o,A.o,[]),e.Cb(6144,A.m,null,[A.o]),e.Cb(4608,A.k,A.k,[A.m]),e.Cb(6144,A.b,null,[A.k]),e.Cb(4608,A.g,A.n,[A.b,e.r]),e.Cb(4608,A.c,A.c,[A.g]),e.Cb(4608,R.c,R.c,[]),e.Cb(5120,T.xe,T.ze,[[3,T.xe],T.ye]),e.Cb(4608,F.e,F.e,[e.v]),e.Cb(5120,T.ue,T.ve,[[3,T.ue],T.we,T.xe,F.e]),e.Cb(4608,J.d,J.d,[J.k,J.f,e.j,J.i,J.g,e.r,e.A,F.d,Z.b]),e.Cb(5120,J.l,J.m,[J.d]),e.Cb(5120,T.S,T.T,[F.d,[3,T.S]]),e.Cb(4608,T.gb,T.gb,[]),e.Cb(4608,T.Ab,T.Ab,[]),e.Cb(4608,T.kd,T.kd,[J.d]),e.Cb(4608,T.Od,T.Od,[J.d,e.r,e.j,e.g]),e.Cb(4608,T.Ud,T.Ud,[J.d,e.r,e.j,e.g]),e.Cb(4608,T.ee,T.ee,[[3,T.ee]]),e.Cb(4608,T.ge,T.ge,[J.d,T.xe,T.ee]),e.Cb(1073742336,F.c,F.c,[]),e.Cb(1073742336,I.q,I.q,[]),e.Cb(1073742336,I.h,I.h,[]),e.Cb(1073742336,I.o,I.o,[]),e.Cb(1073742336,A.e,A.e,[]),e.Cb(1073742336,A.d,A.d,[]),e.Cb(1073742336,R.d,R.d,[]),e.Cb(1073742336,_.b,_.b,[]),e.Cb(1073742336,T.Cc,T.Cc,[]),e.Cb(1073742336,T.Bd,T.Bd,[]),e.Cb(1073742336,T.g,T.g,[]),e.Cb(1073742336,T.i,T.i,[]),e.Cb(1073742336,T.Be,T.Be,[]),e.Cb(1073742336,T.k,T.k,[]),e.Cb(1073742336,Z.a,Z.a,[]),e.Cb(1073742336,V.e,V.e,[]),e.Cb(1073742336,X.a,X.a,[]),e.Cb(1073742336,J.h,J.h,[]),e.Cb(1073742336,T.o,T.o,[]),e.Cb(1073742336,T.Vd,T.Vd,[]),e.Cb(1073742336,T.y,T.y,[]),e.Cb(1073742336,T.D,T.D,[]),e.Cb(1073742336,T.F,T.F,[]),e.Cb(1073742336,T.O,T.O,[]),e.Cb(1073742336,T.V,T.V,[]),e.Cb(1073742336,T.Q,T.Q,[]),e.Cb(1073742336,T.X,T.X,[]),e.Cb(1073742336,T.Z,T.Z,[]),e.Cb(1073742336,T.hb,T.hb,[]),e.Cb(1073742336,T.kb,T.kb,[]),e.Cb(1073742336,T.mb,T.mb,[]),e.Cb(1073742336,T.pb,T.pb,[]),e.Cb(1073742336,T.sb,T.sb,[]),e.Cb(1073742336,T.wb,T.wb,[]),e.Cb(1073742336,T.Gb,T.Gb,[]),e.Cb(1073742336,T.yb,T.yb,[]),e.Cb(1073742336,T.Jb,T.Jb,[]),e.Cb(1073742336,T.Lb,T.Lb,[]),e.Cb(1073742336,T.Nb,T.Nb,[]),e.Cb(1073742336,T.Pb,T.Pb,[]),e.Cb(1073742336,T.Rb,T.Rb,[]),e.Cb(1073742336,T.Tb,T.Tb,[]),e.Cb(1073742336,T.ac,T.ac,[]),e.Cb(1073742336,T.gc,T.gc,[]),e.Cb(1073742336,T.ic,T.ic,[]),e.Cb(1073742336,T.lc,T.lc,[]),e.Cb(1073742336,T.pc,T.pc,[]),e.Cb(1073742336,T.rc,T.rc,[]),e.Cb(1073742336,T.uc,T.uc,[]),e.Cb(1073742336,T.Fc,T.Fc,[]),e.Cb(1073742336,T.Ec,T.Ec,[]),e.Cb(1073742336,T.Dc,T.Dc,[]),e.Cb(1073742336,T.fd,T.fd,[]),e.Cb(1073742336,T.hd,T.hd,[]),e.Cb(1073742336,T.ld,T.ld,[]),e.Cb(1073742336,T.ud,T.ud,[]),e.Cb(1073742336,T.yd,T.yd,[]),e.Cb(1073742336,T.Dd,T.Dd,[]),e.Cb(1073742336,T.Hd,T.Hd,[]),e.Cb(1073742336,T.Jd,T.Jd,[]),e.Cb(1073742336,T.Pd,T.Pd,[]),e.Cb(1073742336,T.Wd,T.Wd,[]),e.Cb(1073742336,T.Yd,T.Yd,[]),e.Cb(1073742336,T.be,T.be,[]),e.Cb(1073742336,T.he,T.he,[]),e.Cb(1073742336,T.je,T.je,[]),e.Cb(1073742336,T.le,T.le,[]),e.Cb(1073742336,T.pe,T.pe,[]),e.Cb(1073742336,T.se,T.se,[]),e.Cb(1073742336,T.b,T.b,[]),e.Cb(1073742336,E.a,E.a,[]),e.Cb(1073742336,f.p,f.p,[[2,f.v],[2,f.m]]),e.Cb(1073742336,H,H,[]),e.Cb(1073742336,s,s,[]),e.Cb(256,A.p,"XSRF-TOKEN",[]),e.Cb(256,A.q,"X-XSRF-TOKEN",[]),e.Cb(256,T.ye,!1,[]),e.Cb(256,T.we,void 0,[]),e.Cb(256,T.Ld,{nzDuration:3e3,nzAnimate:!0,nzPauseOnHover:!0,nzMaxStack:7},[]),e.Cb(256,T.Sd,{nzTop:"24px",nzBottom:"24px",nzPlacement:"topRight",nzDuration:4500,nzMaxStack:7,nzPauseOnHover:!0,nzAnimate:!0},[]),e.Cb(1024,f.k,function(){return[[{path:"",component:p,data:{name:"step"}}]]},[])])})}}]);