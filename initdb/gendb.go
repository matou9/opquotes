package initdb


func CreateTable(){
	account:=new(Account)
	account.CreateTable()
	agent:=new(Agent)
	agent.CreateTable()
	applyfund:=new(Applyfund)
	applyfund.Createtable()
	comm:=new(Commission)
	comm.CreateTable()

	dailya:=new(Dailyaccount)
	dailya.CreateTable()
	dailyp:=new(Dailyposition)
	dailyp.CreateTable()
	day:=new(Day)
	day.CreateTable()
	min1:=new(Min1)
	min1.CreateTable()
	min5:=new(Min5)
	min5.CreateTable()
	min15:=new(Min15)
	min15.CreateTable()
	min30:=new(Min30)
	min30.CreateTable()
	min60:=new(Min60)
	min60.CreateTable()
	min120:=new(Min120)
	min120.CreateTable()

	md:=new(Mindata)
	md.CreateTable()
	op:=new(Option)
	op.CreateTable()
	opt:=new(Optiontick)
	opt.CreateTable()
	order:=new(Orders)
	order.CreateTable()
	paya:=new(Payaccount)
	paya.CreateTable()
	p:=new(Position)
	p.CreateTable()
	sa:=new(Subaccount)
	sa.CreateTable()
	so:=new(Suborders)
	so.CreateTable()
	sp:=new(Subposition)
	sp.CreateTable()

	st:=new(Subtrade)
	st.CreateTable()
	tr:=new(Trade)
	tr.CreateTable()
	tapi:=new(Tradeapi)
	tapi.CreateTable()
	td:=new(Tradedate)
	td.CreateTable()
	tt:=new(Tradetime)
	tt.CreateTable()
	ub:=new(Userbank)
	ub.CreateTable()
	ui:=new(Userinfo)
	ui.CreateTable()


}
