(window.webpackJsonp=window.webpackJsonp||[]).push([[7],{bUiu:function(n,l,t){"use strict";t.r(l);var a=t("CcnG"),u=function(){function n(){}return Object.defineProperty(n.prototype,"step",{get:function(){return this._step},set:function(n){console.log("@Input step: "+n),this._step=n},enumerable:!0,configurable:!0}),n.prototype.ngOnInit=function(){console.log("step one init")},n.prototype.ngOnDestroy=function(){console.log("step one destroy")},n}(),e=function(){function n(){}return Object.defineProperty(n.prototype,"step",{get:function(){return this._step},set:function(n){console.log("@Input step: "+n),this._step=n},enumerable:!0,configurable:!0}),n.prototype.ngOnInit=function(){console.log("step one init")},n.prototype.ngOnDestroy=function(){console.log("step one destroy")},n}(),o=function(){function n(){}return Object.defineProperty(n.prototype,"step",{get:function(){return this._step},set:function(n){console.log("@Input step: "+n),this._step=n},enumerable:!0,configurable:!0}),n.prototype.ngOnInit=function(){console.log("step one init")},n.prototype.ngOnDestroy=function(){console.log("step one destroy")},n}(),p=function(){function n(n){this.route=n}return n.prototype.ngOnInit=function(){var n=this;this.route.queryParams.subscribe(function(l){var t=l.step||"step1";n.stepComp={step1:{component:u,inputs:{step:t}},step2:{component:e,inputs:{step:t}},step3:{component:o,inputs:{step:t}}}[t]})},n}(),s=function(){},c=t("82da"),r=t("Z24s"),i=function(){function n(n,l){this.vcr=n,this.cfr=l}return Object.defineProperty(n.prototype,"data",{set:function(n){var l=this.cfr.resolveComponentFactory(n.component),t=this.vcr.createComponent(l);if(n.inputs)for(var a in n.inputs)a&&(t.instance[a]=n.inputs[a]);this.destroy(),this.currentComponent=t},enumerable:!0,configurable:!0}),n.prototype.destroy=function(){this.currentComponent&&(this.currentComponent.destroy(),this.currentComponent=null)},n.prototype.ngOnDestroy=function(){this.destroy()},n}(),Z=a.Pa({encapsulation:2,styles:[],data:{}});function d(n){return a.lb(0,[],null,null)}var b=t("ZYCi"),f=a.Pa({encapsulation:0,styles:[[""]],data:{}});function m(n){return a.lb(0,[(n()(),a.Ra(0,0,null,null,1,"h1",[],null,null,null,null,null)),(n()(),a.jb(-1,null,["Angular\u52a8\u6001\u52a0\u8f7d\u7ec4\u4ef6"])),(n()(),a.Ra(2,0,null,null,1,"a",[["href","https://segmentfault.com/a/1190000009582289"]],null,null,null,null,null)),(n()(),a.jb(-1,null,["https://segmentfault.com/a/1190000009582289"])),(n()(),a.Ra(4,16777216,null,null,1,"app-step",[],null,null,null,d,Z)),a.Qa(5,180224,null,0,i,[a.P,a.j],{data:[0,"data"]},null),(n()(),a.Ra(6,0,null,null,1,"a",[["href","#/admin/step?step=step1"]],null,null,null,null,null)),(n()(),a.jb(-1,null,["step1"])),(n()(),a.Ra(8,0,null,null,0,"br",[],null,null,null,null,null)),(n()(),a.Ra(9,0,null,null,1,"a",[["href","#/admin/step?step=step2"]],null,null,null,null,null)),(n()(),a.jb(-1,null,["step2"])),(n()(),a.Ra(11,0,null,null,0,"br",[],null,null,null,null,null)),(n()(),a.Ra(12,0,null,null,1,"a",[["href","#/admin/step?step=step3"]],null,null,null,null,null)),(n()(),a.jb(-1,null,["step3"]))],function(n,l){n(l,5,0,l.component.stepComp)},null)}var g=a.Na("app-index",p,function(n){return a.lb(0,[(n()(),a.Ra(0,0,null,null,1,"app-index",[],null,null,null,m,f)),a.Qa(1,114688,null,0,p,[b.a],null,null)],function(n,l){n(l,1,0)},null)},{},{},[]),y=a.Pa({encapsulation:0,styles:[[""]],data:{}});function h(n){return a.lb(0,[(n()(),a.Ra(0,0,null,null,1,"h2",[],null,null,null,null,null)),(n()(),a.jb(1,null,["Step One Component\uff1aparams value: ",""]))],null,function(n,l){n(l,1,0,l.component.step)})}var v=a.Na("app-user1",u,function(n){return a.lb(0,[(n()(),a.Ra(0,0,null,null,1,"app-user1",[],null,null,null,h,y)),a.Qa(1,245760,null,0,u,[],null,null)],function(n,l){n(l,1,0)},null)},{step:"step"},{},[]),O=a.Pa({encapsulation:0,styles:[[""]],data:{}});function R(n){return a.lb(0,[(n()(),a.Ra(0,0,null,null,1,"h2",[],null,null,null,null,null)),(n()(),a.jb(1,null,["Step One Component\uff1aparams value: ",""]))],null,function(n,l){n(l,1,0,l.component.step)})}var j=a.Na("app-user2",e,function(n){return a.lb(0,[(n()(),a.Ra(0,0,null,null,1,"app-user2",[],null,null,null,R,O)),a.Qa(1,245760,null,0,e,[],null,null)],function(n,l){n(l,1,0)},null)},{step:"step"},{},[]),C=a.Pa({encapsulation:0,styles:[[""]],data:{}});function N(n){return a.lb(0,[(n()(),a.Ra(0,0,null,null,1,"h2",[],null,null,null,null,null)),(n()(),a.jb(1,null,["Step One Component\uff1aparams value: ",""]))],null,function(n,l){n(l,1,0,l.component.step)})}var P=a.Na("app-user3",o,function(n){return a.lb(0,[(n()(),a.Ra(0,0,null,null,1,"app-user3",[],null,null,null,N,C)),a.Qa(1,245760,null,0,o,[],null,null)],function(n,l){n(l,1,0)},null)},{step:"step"},{},[]),q=t("Ip0R"),z=t("gIcY"),A=t("t/Na"),S=t("0+vK"),k=t("5NvZ"),w=t("M2Lx"),I=t("tn8F"),D=t("eDkP"),F=t("Fzqc"),Q=t("4c35"),x=t("dWZg"),J=t("qAlS"),K=t("ADsi"),T=function(){function n(){}return n.withComponents=function(l){return{ngModule:n,providers:[{provide:a.a,useValue:l,multi:!0}]}},n}();t.d(l,"StepModuleNgFactory",function(){return E});var E=a.Oa(s,[],function(n){return a.Ya([a.Za(512,a.j,a.Ca,[[8,[c.M,c.N,c.O,c.P,c.Q,c.R,c.S,r.a,g,v,j,P]],[3,a.j],a.w]),a.Za(4608,q.o,q.n,[a.t,[2,q.A]]),a.Za(4608,z.s,z.s,[]),a.Za(4608,z.d,z.d,[]),a.Za(4608,A.l,A.r,[q.d,a.A,A.p]),a.Za(4608,A.s,A.s,[A.l,A.q]),a.Za(5120,A.a,function(n,l){return[n,new S.a(l)]},[A.s,k.a]),a.Za(4608,A.o,A.o,[]),a.Za(6144,A.m,null,[A.o]),a.Za(4608,A.k,A.k,[A.m]),a.Za(6144,A.b,null,[A.k]),a.Za(4608,A.g,A.n,[A.b,a.q]),a.Za(4608,A.c,A.c,[A.g]),a.Za(4608,w.c,w.c,[]),a.Za(5120,I.Ed,I.Gd,[[3,I.Ed],I.Fd]),a.Za(4608,q.e,q.e,[a.t]),a.Za(5120,I.cc,I.Bc,[[3,I.cc],I.td,I.Ed,q.e]),a.Za(4608,D.d,D.d,[D.k,D.f,a.j,D.i,D.g,a.q,a.y,q.d,F.b]),a.Za(5120,D.l,D.m,[D.d]),a.Za(5120,I.N,I.O,[q.d,[3,I.N]]),a.Za(4608,I.Aa,I.Aa,[]),a.Za(4608,I.Va,I.Va,[]),a.Za(4608,I.cd,I.cd,[D.d,a.q,a.j,a.g]),a.Za(4608,I.id,I.id,[D.d,a.q,a.j,a.g]),a.Za(4608,I.qd,I.qd,[[3,I.qd]]),a.Za(4608,I.sd,I.sd,[D.d,I.Ed,I.qd]),a.Za(1073742336,q.c,q.c,[]),a.Za(1073742336,z.q,z.q,[]),a.Za(1073742336,z.h,z.h,[]),a.Za(1073742336,z.o,z.o,[]),a.Za(1073742336,A.e,A.e,[]),a.Za(1073742336,A.d,A.d,[]),a.Za(1073742336,w.d,w.d,[]),a.Za(1073742336,I.g,I.g,[]),a.Za(1073742336,I.Jd,I.Jd,[]),a.Za(1073742336,I.Id,I.Id,[]),a.Za(1073742336,I.Ld,I.Ld,[]),a.Za(1073742336,F.a,F.a,[]),a.Za(1073742336,Q.c,Q.c,[]),a.Za(1073742336,x.b,x.b,[]),a.Za(1073742336,J.a,J.a,[]),a.Za(1073742336,D.h,D.h,[]),a.Za(1073742336,I.j,I.j,[]),a.Za(1073742336,I.cb,I.cb,[]),a.Za(1073742336,I.t,I.t,[]),a.Za(1073742336,I.y,I.y,[]),a.Za(1073742336,I.A,I.A,[]),a.Za(1073742336,I.J,I.J,[]),a.Za(1073742336,I.Q,I.Q,[]),a.Za(1073742336,I.L,I.L,[]),a.Za(1073742336,I.S,I.S,[]),a.Za(1073742336,I.U,I.U,[]),a.Za(1073742336,I.Ba,I.Ba,[]),a.Za(1073742336,I.Fa,I.Fa,[]),a.Za(1073742336,I.Ha,I.Ha,[]),a.Za(1073742336,I.Ka,I.Ka,[]),a.Za(1073742336,I.Na,I.Na,[]),a.Za(1073742336,I.Ra,I.Ra,[]),a.Za(1073742336,I.ab,I.ab,[]),a.Za(1073742336,I.Ta,I.Ta,[]),a.Za(1073742336,I.eb,I.eb,[]),a.Za(1073742336,I.gb,I.gb,[]),a.Za(1073742336,I.ib,I.ib,[]),a.Za(1073742336,I.kb,I.kb,[]),a.Za(1073742336,I.mb,I.mb,[]),a.Za(1073742336,I.ob,I.ob,[]),a.Za(1073742336,I.vb,I.vb,[]),a.Za(1073742336,I.Ab,I.Ab,[]),a.Za(1073742336,I.Db,I.Db,[]),a.Za(1073742336,I.Gb,I.Gb,[]),a.Za(1073742336,I.Kb,I.Kb,[]),a.Za(1073742336,I.Ob,I.Ob,[]),a.Za(1073742336,I.Qb,I.Qb,[]),a.Za(1073742336,I.Tb,I.Tb,[]),a.Za(1073742336,I.dc,I.dc,[]),a.Za(1073742336,I.bc,I.bc,[]),a.Za(1073742336,I.xc,I.xc,[]),a.Za(1073742336,I.zc,I.zc,[]),a.Za(1073742336,I.Jc,I.Jc,[]),a.Za(1073742336,I.Nc,I.Nc,[]),a.Za(1073742336,I.Rc,I.Rc,[]),a.Za(1073742336,I.Vc,I.Vc,[]),a.Za(1073742336,I.Xc,I.Xc,[]),a.Za(1073742336,I.dd,I.dd,[]),a.Za(1073742336,I.jd,I.jd,[]),a.Za(1073742336,I.ld,I.ld,[]),a.Za(1073742336,I.od,I.od,[]),a.Za(1073742336,I.ud,I.ud,[]),a.Za(1073742336,I.wd,I.wd,[]),a.Za(1073742336,I.yd,I.yd,[]),a.Za(1073742336,I.Cd,I.Cd,[]),a.Za(1073742336,I.b,I.b,[]),a.Za(1073742336,K.a,K.a,[]),a.Za(1073742336,b.o,b.o,[[2,b.t],[2,b.l]]),a.Za(1073742336,T,T,[]),a.Za(1073742336,s,s,[]),a.Za(256,A.p,"XSRF-TOKEN",[]),a.Za(256,A.q,"X-XSRF-TOKEN",[]),a.Za(256,I.Fd,!1,[]),a.Za(256,I.td,void 0,[]),a.Za(256,I.Zc,{nzDuration:3e3,nzAnimate:!0,nzPauseOnHover:!0,nzMaxStack:7},[]),a.Za(256,I.gd,{nzTop:"24px",nzBottom:"24px",nzPlacement:"topRight",nzDuration:4500,nzMaxStack:7,nzPauseOnHover:!0,nzAnimate:!0},[]),a.Za(1024,b.j,function(){return[[{path:"",component:p,data:{name:"step"}}]]},[])])})}}]);