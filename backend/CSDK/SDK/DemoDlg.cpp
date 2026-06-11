
// DemoDlg.cpp : 实现文件
//
#include "stdafx.h"
#include "Demo.h"
#include "DemoDlg.h"
#include "afxdialogex.h"
#include <new> 
#ifdef _DEBUG
#define new DEBUG_NEW
#endif

#define FP_AUTO_CONNECT
//#define FP_MULTI_PARALLEL_MODE

#define FPDTOTALNUM 3   //total number of FPD
CString editLog;

CHAR FPSn[3][32] = {'\0'};
CHAR FPSnOpened[32] = {0};
CHAR* pPicBuff=NULL;   
//根据平板尺寸配备合适的空间,如下，或者动态申请对应型号的空间.
//CHAR pPicBuff[3072*3072*2] = {0};
CHAR* pFpMultiPicBuff[FPDTOTALNUM] = { NULL };
// 用于应用程序“关于”菜单项的 CAboutDlg 对话框

class CAboutDlg : public CDialogEx
{
public:
	CAboutDlg();

// 对话框数据
	enum { IDD = IDD_ABOUTBOX };

	protected:
	virtual void DoDataExchange(CDataExchange* pDX);    // DDX/DDV 支持

// 实现
protected:
	DECLARE_MESSAGE_MAP()
};

CAboutDlg::CAboutDlg() : CDialogEx(CAboutDlg::IDD)
{
}

void CAboutDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
}

BEGIN_MESSAGE_MAP(CAboutDlg, CDialogEx)
END_MESSAGE_MAP()

// CDemoDlg 对话框

CDemoDlg::CDemoDlg(CWnd* pParent /*=NULL*/)
	: CDialogEx(CDemoDlg::IDD, pParent)
{
	m_hIcon = AfxGetApp()->LoadIcon(IDR_MAINFRAME);
}

void CDemoDlg::DoDataExchange(CDataExchange* pDX)
{
	CDialogEx::DoDataExchange(pDX);
}

BEGIN_MESSAGE_MAP(CDemoDlg, CDialogEx)
	ON_WM_SYSCOMMAND()
	ON_WM_PAINT()
	ON_WM_QUERYDRAGICON()
	ON_BN_CLICKED(IDC_BUTTONINIT, &CDemoDlg::OnBnClickedButtoninit)
	ON_BN_CLICKED(IDC_BUTTONOPEN, &CDemoDlg::OnBnClickedButtonopen)
	ON_BN_CLICKED(IDC_BUTTON_HST, &CDemoDlg::OnBnClickedButtonHst)
	ON_BN_CLICKED(IDC_BUTTON_AED, &CDemoDlg::OnBnClickedButtonAed)
	ON_BN_CLICKED(IDC_BUTTON_STOP, &CDemoDlg::OnBnClickedButtonStop)
	ON_BN_CLICKED(IDC_BUTTONCLOSE, &CDemoDlg::OnBnClickedButtonclose)
	ON_CBN_SELCHANGE(IDC_COMBO_FPLIST, &CDemoDlg::OnCbnSelchangeComboFplist)
	ON_CBN_DROPDOWN(IDC_COMBO_FPLIST, &CDemoDlg::OnCbnDropdownComboFplist)
	ON_BN_CLICKED(IDC_BUTTON_HST2, &CDemoDlg::OnBnClickedButtonTrigger)
	ON_BN_CLICKED(IDC_BUTTON_AED2, &CDemoDlg::OnBnClickedButtonAedTrigger)
	ON_BN_CLICKED(IDC_BUTTONNETSTOP, &CDemoDlg::OnBnClickedButtonnetstop)
	ON_BN_CLICKED(IDC_BUTTONNETSTART, &CDemoDlg::OnBnClickedButtonnetstart)

	ON_BN_CLICKED(IDC_BUTTON_AED_MULTI, &CDemoDlg::OnBnClickedButtonAedMulti)
	ON_BN_CLICKED(IDC_BTN_HST_MULTI, &CDemoDlg::OnBnClickedBtnHstMulti)
	ON_BN_CLICKED(IDC_BTN_HSTTRIGGET_MULTI, &CDemoDlg::OnBnClickedBtnHsttriggetMulti)
	ON_BN_CLICKED(IDC_BTN_DST_MULTI, &CDemoDlg::OnBnClickedBtnDstMulti)
	ON_BN_CLICKED(IDC_BTN_Dacq_MULTI, &CDemoDlg::OnBnClickedBtnDacqMulti)
	ON_BN_CLICKED(IDC_BTN_Dexit_MULTI, &CDemoDlg::OnBnClickedBtnDexitMulti)
	ON_BN_CLICKED(IDC_BTN_STOP_MULTI, &CDemoDlg::OnBnClickedBtnStopMulti)
END_MESSAGE_MAP()


// CDemoDlg 消息处理程序

BOOL CDemoDlg::OnInitDialog()
{
	CDialogEx::OnInitDialog();

	// 将“关于...”菜单项添加到系统菜单中。

	// IDM_ABOUTBOX 必须在系统命令范围内。
	ASSERT((IDM_ABOUTBOX & 0xFFF0) == IDM_ABOUTBOX);
	ASSERT(IDM_ABOUTBOX < 0xF000);

	CMenu* pSysMenu = GetSystemMenu(FALSE);
	if (pSysMenu != NULL)
	{
		BOOL bNameValid;
		CString strAboutMenu;
		bNameValid = strAboutMenu.LoadString(IDS_ABOUTBOX);
		ASSERT(bNameValid);
		if (!strAboutMenu.IsEmpty())
		{
			pSysMenu->AppendMenu(MF_SEPARATOR);
			pSysMenu->AppendMenu(MF_STRING, IDM_ABOUTBOX, strAboutMenu);
		}
	}

	// 设置此对话框的图标。  当应用程序主窗口不是对话框时，框架将自动
	//  执行此操作
	SetIcon(m_hIcon, TRUE);			// 设置大图标
	SetIcon(m_hIcon, FALSE);		// 设置小图标

	ShowWindow(SW_SHOW);

	// TODO:  在此添加额外的初始化代码

	SetDlgItemText(IDC_STATIC_SN, NULL);
	SetDlgItemText(IDC_STATIC_STAT, NULL);
	SetDlgItemText(IDC_STATIC_WIFISIG, NULL);
	SetDlgItemText(IDC_STATIC_TEMPHUM, NULL);
	SetDlgItemText(IDC_STATIC_BAT, NULL);

	SetDlgItemText(IDC_STATIC_SN2, NULL);
	SetDlgItemText(IDC_STATIC_STAT2, NULL);
	SetDlgItemText(IDC_STATIC_WIFISIG2, NULL);
	SetDlgItemText(IDC_STATIC_TEMPHUM2, NULL);
	SetDlgItemText(IDC_STATIC_BAT2, NULL);

	SetDlgItemText(IDC_STATIC_SN3, NULL);
	SetDlgItemText(IDC_STATIC_STAT3, NULL);
	SetDlgItemText(IDC_STATIC_WIFISIG3, NULL);
	SetDlgItemText(IDC_STATIC_TEMPHUM3, NULL);
	SetDlgItemText(IDC_STATIC_BAT3, NULL);

	GetDlgItem(IDC_BUTTON_HST)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTON_HST2)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTON_AED)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTON_AED2)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTON_STOP)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTONCLOSE)->EnableWindow(FALSE);

	return TRUE;  // 除非将焦点设置到控件，否则返回 TRUE
}

void CDemoDlg::OnSysCommand(UINT nID, LPARAM lParam)
{
	if ((nID & 0xFFF0) == IDM_ABOUTBOX)
	{
		CAboutDlg dlgAbout;
		dlgAbout.DoModal();
	}
	else
	{
		CDialogEx::OnSysCommand(nID, lParam);
	}
}

// 如果向对话框添加最小化按钮，则需要下面的代码
//  来绘制该图标。  对于使用文档/视图模型的 MFC 应用程序，
//  这将由框架自动完成。

void CDemoDlg::OnPaint()
{
	if (IsIconic())
	{
		CPaintDC dc(this); // 用于绘制的设备上下文

		SendMessage(WM_ICONERASEBKGND, reinterpret_cast<WPARAM>(dc.GetSafeHdc()), 0);

		// 使图标在工作区矩形中居中
		int cxIcon = GetSystemMetrics(SM_CXICON);
		int cyIcon = GetSystemMetrics(SM_CYICON);
		CRect rect;
		GetClientRect(&rect);
		int x = (rect.Width() - cxIcon + 1) / 2;
		int y = (rect.Height() - cyIcon + 1) / 2;

		// 绘制图标
		dc.DrawIcon(x, y, m_hIcon);
	}
	else
	{
		CDialogEx::OnPaint();
	}
}

//当用户拖动最小化窗口时系统调用此函数取得光标
//显示。
HCURSOR CDemoDlg::OnQueryDragIcon()
{
	return static_cast<HCURSOR>(m_hIcon);
}

BOOL WINAPI FuncImageCallBack(char nEvent)
{
	CString auctmp;
	CDemoDlg *pDlg = (CDemoDlg*)theApp.m_pMainWnd;

	TImageMode tImageMode;
	COM_GetImageMode(&tImageMode);
	UINT16 u16ImgRow = tImageMode.usRow;
	UINT16 u16ImgCol = tImageMode.usCol;
	auctmp.Format(_T("Get image demo.raw\r\n row:%d,col:%d,16 bit unsigned/pixel,Little-endian \r\n"), u16ImgRow, u16ImgCol);
	editLog.Append(auctmp);
	pDlg->SetDlgItemText(IDC_EDIT_LOG, editLog);

	COM_GetImage(pPicBuff);

	CString filename;
	filename = "demo.raw";
	CFile file(filename, CFile::modeWrite | CFile::modeCreate | CFile::typeBinary);
	file.Write(pPicBuff, 2 * u16ImgRow*u16ImgCol);
	file.Close();

	return TRUE;
}

BOOL WINAPI FuncBreakCallBack(char nEvent)
{
	CDemoDlg *pDlg = (CDemoDlg*)theApp.m_pMainWnd;

	editLog.Append(_T("FP Link down\r\n"));
	pDlg->SetDlgItemText(IDC_EDIT_LOG, editLog);
	return 0;
}

BOOL WINAPI FuncLinkCallBack(char nEvent)
{
	CDemoDlg *pDlg = (CDemoDlg*)theApp.m_pMainWnd;

	COM_GetFPsn(FPSnOpened);
	editLog.Append(_T("FP Link up\r\n"));
	pDlg->SetDlgItemText(IDC_EDIT_LOG, editLog);

	TComFpList tComFpList = {0};
	COM_List(&tComFpList);

	for (CHAR i = 0; i < FPDTOTALNUM; i++)
	{
		if (0 == strncmp(FPSnOpened, tComFpList.tFpNode[i].FPPsn, 32))
			pDlg->SetDlgItemText(IDC_STATIC_OPEN + i, _T("opened"));
		else
			pDlg->SetDlgItemText(IDC_STATIC_OPEN + i, _T("connected"));
	}
	
#ifdef FP_AUTO_CONNECT
	TImageMode tImageMode;
	COM_GetImageMode(&tImageMode);
	UINT16 u16ImgRow = tImageMode.usRow;
	UINT16 u16ImgCol = tImageMode.usCol;

	//Make sure that image buffer has been allocated. Its just for single fp.
	try
	{
		pPicBuff = new CHAR[u16ImgRow * u16ImgCol * sizeof(UINT16)];
	}
	catch (const std::bad_alloc &ex)
	{
		editLog.Append(_T("ERROR:New image buffer."));
		pDlg->SetDlgItemText(IDC_EDIT_LOG, editLog);
		return 1;
	}
#endif

	pDlg->GetDlgItem(IDC_BUTTON_HST)->EnableWindow(TRUE);
	pDlg->GetDlgItem(IDC_BUTTON_HST2)->EnableWindow(TRUE);
	pDlg->GetDlgItem(IDC_BUTTON_AED)->EnableWindow(TRUE);
	pDlg->GetDlgItem(IDC_BUTTON_AED2)->EnableWindow(TRUE);
	pDlg->GetDlgItem(IDC_BUTTON_STOP)->EnableWindow(TRUE);
	pDlg->GetDlgItem(IDC_BUTTONCLOSE)->EnableWindow(TRUE);
	return 0;
}

BOOL WINAPI FuncHeartBeatCallBack(char nEvent)
{
	CDemoDlg *pDlg = (CDemoDlg*)theApp.m_pMainWnd;
	CHAR acSnTmp[32] = {0};
#ifdef FP_AUTO_CONNECT
	COM_GetFPsn(acSnTmp);
	CHAR cFpCurStat = COM_GetFPCurStatus();
	TFPStat tFPStat = { 0 };
	COM_GetFPStatus(&tFPStat);	
	pDlg->ShowHeartBeatInfo(0, cFpCurStat, acSnTmp, &tFPStat);	
#endif
	return 0;
}

BOOL WINAPI FuncReadyCallBack(char nEvent)
{
	CDemoDlg *pDlg = (CDemoDlg*)theApp.m_pMainWnd;

	editLog.Append(_T("FP Ready，It's time to fire x-ray.\r\n"));
	pDlg->SetDlgItemText(IDC_EDIT_LOG, editLog);
	return 0;
}

BOOL WINAPI FuncErrorCallBack(char nEvent)
{
	INT32 errNo= COM_GetErrNo();
	//e.g. COM_NO_TPL
	return 0;
}

BOOL WINAPI FuncExposeCallBack(char nEvent)
{
	CDemoDlg *pDlg = (CDemoDlg*)theApp.m_pMainWnd;

	editLog.Append(_T("FP Expose\r\n"));
	pDlg->SetDlgItemText(IDC_EDIT_LOG, editLog);
	return 0;
}

BOOL CDemoDlg::ShowHeartBeatInfo(char index, CHAR cFpCurStatTmp, CHAR* acSnTmp, TFPStat* ptFPStatTmp)
{
	CString auctmp;
	auctmp.Format(_T("index:%d,Sn:%s"),index, CStringW(acSnTmp));
	SetDlgItemText(IDC_STATIC_SN + index, auctmp);  //IDC_STATIC_SN or IDC_STATIC_SN2 or IDC_STATIC_SN3
	if (STATUS_IDLE == cFpCurStatTmp)
		auctmp = "Stat:IDLE";
	else if (STATUS_HST == cFpCurStatTmp)
		auctmp = "Stat:HST";
	else if (STATUS_AED1 == cFpCurStatTmp)
		auctmp = "Stat:AED1";
	else if (STATUS_AED2 == cFpCurStatTmp)
		auctmp = "Stat:AED2";
	else if (STATUS_RECOVER == cFpCurStatTmp)
		auctmp = "Stat:RECOVER";
	else
		auctmp = "Stat:ERR";

	SetDlgItemText(IDC_STATIC_STAT + index, auctmp); //IDC_STATIC_STAT or IDC_STATIC_STAT2 or IDC_STATIC_STAT3

	auctmp.Format(_T("Wifi signal:%d"), ptFPStatTmp->tWifiStatus.ucSignal_level);
	SetDlgItemText(IDC_STATIC_WIFISIG + index, auctmp);  //IDC_STATIC_WIFISIG or IDC_STATIC_WIFISIG2 or IDC_STATIC_WIFISIG3

	auctmp.Format(_T("Temp:%.1f,Hum:%.1f"), (double(ptFPStatTmp->tFpTempHum.Temp)) / 10, (double(ptFPStatTmp->tFpTempHum.Hum)) / 10);
	SetDlgItemText(IDC_STATIC_TEMPHUM + index, auctmp); //IDC_STATIC_TEMPHUM or IDC_STATIC_TEMPHUM2 or IDC_STATIC_TEMPHUM3

	if (0 != ptFPStatTmp->tBatInfo1.full)
		auctmp.Format(_T("Bat1:%.2f,"), double(ptFPStatTmp->tBatInfo1.Remain) / (ptFPStatTmp->tBatInfo1.full));
	else
		auctmp = "Bat1:NULL,";
	CString auctmp1;
	if (0 != ptFPStatTmp->tBatInfo2.full)
		auctmp1.Format(_T("Bat2:%.2f"), double(ptFPStatTmp->tBatInfo2.Remain) / (ptFPStatTmp->tBatInfo2.full));
	else
		auctmp1 = "Bat2:NULL";
	SetDlgItemText(IDC_STATIC_BAT + index, auctmp + auctmp1);
	return 0;
}

#ifdef FP_MULTI_PARALLEL_MODE
BOOL WINAPI FuncLinkupCallBackEx(INT16 nEvent, char index)
{
	CDemoDlg *pDlg = (CDemoDlg*)theApp.m_pMainWnd;
	CString auctmp;
	CHAR acSnTmp[32] = { '\0' };
	COM_GetFPsnEx(index, acSnTmp);

	strncpy_s(FPSn[index], acSnTmp, 32);
	auctmp.Format(_T("FPsn:%s,index=%d, connected!\r\n"), CStringW(acSnTmp), index);
	editLog.Append(auctmp);
	pDlg->SetDlgItemText(IDC_EDIT_LOG, editLog);

	auctmp = "connected";
	pDlg->SetDlgItemText(IDC_STATIC_OPEN + index, auctmp);

	TImageMode tImageMode = {0};
	COM_GetImageModeEx(&tImageMode,index);

	UINT16 u16ImgRow = tImageMode.usRow;
	UINT16 u16ImgCol = tImageMode.usCol;
	// Make sure memory of image has been allocated for the panel with index "index".
	try
	{
		pFpMultiPicBuff[index] = new CHAR[u16ImgRow * u16ImgCol*sizeof(UINT16)];
	}
	catch (const std::bad_alloc &ex)
	{
		editLog.Append(_T("ERROR:New image buffer."));
		pDlg->SetDlgItemText(IDC_EDIT_LOG, editLog);
		return 1;
	}
	return 0;
}

BOOL WINAPI FuncBreakCallBackEx(INT16 nEvent, char index)
{
	// To handle disconnected events
	CDemoDlg *pDlg = (CDemoDlg*)theApp.m_pMainWnd;
	CString auctmp;
	CHAR acSnTmp[32] = { '\0' };
	COM_GetFPsnEx(index, acSnTmp);
	auctmp.Format(_T("%s,index=%d,disconnected!\r\n"), CStringW(acSnTmp), index);

	editLog.Append(auctmp);
	pDlg->SetDlgItemText(IDC_EDIT_LOG, editLog);

	auctmp = "disconnected";
	pDlg->SetDlgItemText(IDC_STATIC_OPEN + index, auctmp);
	return 0;
}

BOOL WINAPI FuncImageCallBackEx(INT16 nEvent, char index)
{
	CDemoDlg *pDlg = (CDemoDlg*)theApp.m_pMainWnd;

	TImageMode tImageMode;
	COM_GetImageModeEx(&tImageMode, index);
	UINT16 u16ImgRow = tImageMode.usRow;
	UINT16 u16ImgCol = tImageMode.usCol;
	//In the multi-panel parallel acquisition mode, considering concurrency,it is necessary to allocate a separate image buffer for each panel.
	//If the panel will not capture images in parallel, you can only allocate a buffer of appropriate size.
	COM_GetImageEx(pFpMultiPicBuff[index], index);
	// Notify the task to display the image from the panel with index “index”.

	CString auctmp;
	auctmp.Format(_T("FP index=%d,image received!\r\n"),index);
	editLog.Append(auctmp);
	pDlg->SetDlgItemText(IDC_EDIT_LOG, editLog);
	return 0;
}

BOOL WINAPI FuncHeartBeatCallBackEx(INT16 nEvent, char index)
{
	//Each panel triggers a heartbeat callback every 500ms.it is
	//recommended to get the panel status and refresh the display to the interface.
	CDemoDlg *pDlg = (CDemoDlg*)theApp.m_pMainWnd;
	TFPStat tFPStatTmp = { 0 };
	CHAR acSnTmp[32] = { '\0' };
	COM_GetFPsnEx(index, acSnTmp);
	CHAR cFpCurStatTmp = COM_GetFPCurStatusEx(index);
	COM_GetFPStatusEx(&tFPStatTmp, index);
	pDlg->ShowHeartBeatInfo(index, cFpCurStatTmp, acSnTmp, &tFPStatTmp);
	return 0;
}
#endif

void CDemoDlg::OnBnClickedButtoninit()
{
	//These callbacks(no suffix EX) are only valid for opened panel.
	//Only one panel is open at any one time.(COM_Open(CHAR *psn))
	COM_RegisterEvCallBack(EVENT_LINKUP, FuncLinkCallBack);
	COM_RegisterEvCallBack(EVENT_LINKDOWN, FuncBreakCallBack);
	COM_RegisterEvCallBack(EVENT_IMAGEVALID, FuncImageCallBack);
	COM_RegisterEvCallBack(EVENT_HEARTBEAT, FuncHeartBeatCallBack);
	COM_RegisterEvCallBack(EVENT_READY, FuncReadyCallBack);  //Software synchronization mode.
	COM_RegisterEvCallBack(EVENT_TrigErr, FuncErrorCallBack);
//	COM_RegisterEvCallBack(EVENT_EXPOSE, FuncExposeCallBack);

#ifdef FP_MULTI_PARALLEL_MODE
	//These callbacks apply to all connected panels.
	//These registered callback functions exist in parallel with the 
	//callback functions registered above except for COM_RegisterEvCallBack(EVENT_IMAGEVALID, FuncImageCallBack); 
	COM_RegisterEvCallBackEx(EVENT_LINKUP, FuncLinkupCallBackEx);
	COM_RegisterEvCallBackEx(EVENT_LINKDOWN, FuncBreakCallBackEx);
	COM_RegisterEvCallBackEx(EVENT_HEARTBEAT, FuncHeartBeatCallBackEx);
	COM_RegisterEvCallBackEx(EVENT_IMAGEVALID, FuncImageCallBackEx);
	//This callback function "FuncImageCallBackEx" is mutually exclusive with the callback function "FuncImageCallBack" registered by COM_RegisterEvCallBack.
	//When function "FuncImageCallBackEx" is registered by COM_RegisterEvCallBackEx, 
	//the function "FuncImageCallBack" registered by COM_RegisterEvCallBack will not be triggered.
#endif

	COM_Init();

#ifdef FP_AUTO_CONNECT //auto connect at single fp
	COM_Open(NULL);
	GetDlgItem(IDC_COMBO_FPLIST)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTONOPEN)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTONCLOSE)->EnableWindow(FALSE);
#endif

	COM_SetCalibMode(IMG_CALIB_GAIN | IMG_CALIB_DEFECT);  //sdk calibration mode
	/*For the static panel, the panel performs offset correction, and the SDK performs gain and defect correction.
	The static panel is configured with offset correction by default. This configuration will not be lost when power is lost.
	For high frame rate image acquisition on dynamic panels, offset, gain and defect corrections are generally performed on panels. 
	The SDK does not perform any corrections.*/
	editLog.Append(_T("FP init done\r\n"));
	SetDlgItemText(IDC_EDIT_LOG, editLog);
	
	GetDlgItem(IDC_BUTTONINIT)->EnableWindow(FALSE);
}

void CDemoDlg::OnBnClickedButtonopen()
{
	CComboBox* cmbFpList = (CComboBox*)GetDlgItem(IDC_COMBO_FPLIST);
	int nCur = cmbFpList->GetCurSel();
	CString csTmp;
	CString snSelect;
	if (nCur >= 0)
	{
		((CComboBox*)GetDlgItem(IDC_COMBO_FPLIST))->GetLBText(nCur, snSelect);
	}
	USES_CONVERSION;
	CHAR* psn = T2A(snSelect);
	COM_Open(psn);
	csTmp.Format(_T("FP:%s opened\r\n"), snSelect);
	editLog.Append(csTmp);
	SetDlgItemText(IDC_EDIT_LOG, editLog);

	GetDlgItem(IDC_BUTTONOPEN)->EnableWindow(FALSE);
}

void CDemoDlg::OnBnClickedButtonHst()
{
	COM_HstAcq();

	editLog.Append(_T("FP Hst Mode\r\n"));
	SetDlgItemText(IDC_EDIT_LOG, editLog);
}

void CDemoDlg::OnBnClickedButtonAed()
{
	COM_AedAcq();
	editLog.Append(_T("FP Aed Mode\r\n"));
	SetDlgItemText(IDC_EDIT_LOG, editLog);
}

void CDemoDlg::OnBnClickedButtonStop()
{
	COM_Stop();

	editLog.Append(_T("FP Stop\r\n"));
	SetDlgItemText(IDC_EDIT_LOG, editLog);
}

void CDemoDlg::OnBnClickedButtonclose()
{
	COM_Close();

	editLog.Append(_T("FP Close\r\n"));
	SetDlgItemText(IDC_EDIT_LOG, editLog);

	GetDlgItem(IDC_BUTTONOPEN)->EnableWindow(TRUE);

	GetDlgItem(IDC_BUTTON_HST)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTON_HST2)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTON_AED)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTON_AED2)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTON_STOP)->EnableWindow(FALSE);
	GetDlgItem(IDC_BUTTONCLOSE)->EnableWindow(FALSE);
}

void CDemoDlg::OnCbnSelchangeComboFplist()
{

}

void CDemoDlg::OnCbnDropdownComboFplist()
{
	// ==========================================
	// 调试代码开始：向界面的 LOG 控件追加一条调试日志
	editLog.Append(_T("[Debug] 用户点击了下拉菜单，正在刷新设备列表...\r\n"));
	SetDlgItemText(IDC_EDIT_LOG, editLog);


	CComboBox* cmbFpList = (CComboBox*)GetDlgItem(IDC_COMBO_FPLIST);
	cmbFpList->ResetContent();

	TComFpList tComFpList;

	COM_List(&tComFpList);
	printf("list\n"); // 这一行只在控制台程序或有 Console 的调试下可见

	char cnt = tComFpList.ncount;

	// ==========================================
	// 调试代码：打印找到的设备数量
	CString csCount;
	csCount.Format(_T("[Debug] 发现 FPD 设备数量: %d 个\r\n"), cnt);
	editLog.Append(csCount);
	SetDlgItemText(IDC_EDIT_LOG, editLog);
	// ==========================================

	for (char i = 0; i < cnt; i++)
	{
		cmbFpList->AddString(CString(tComFpList.tFpNode[i].FPPsn));
	}
}

void CDemoDlg::OnBnClickedButtonTrigger()
{
	//Software synchronization mode in HST mode.
	COM_ExposeReq();
	editLog.Append(_T("Trigger\r\n"));
	SetDlgItemText(IDC_EDIT_LOG, editLog);
}

void CDemoDlg::OnBnClickedButtonAedTrigger()
{
	COM_AedTrigger();
	editLog.Append(_T("FP Aed Trigger\r\n"));
	SetDlgItemText(IDC_EDIT_LOG, editLog);
}

void CDemoDlg::OnBnClickedButtonnetstop()
{
	COM_StopNet();
	editLog.Append(_T("Stop Net\r\n"));
	SetDlgItemText(IDC_EDIT_LOG, editLog);
}

void CDemoDlg::OnBnClickedButtonnetstart()
{
	COM_StartNet();
	editLog.Append(_T("Start Net\r\n"));
	SetDlgItemText(IDC_EDIT_LOG, editLog);
}

void CDemoDlg::OnBnClickedButtonAedMulti()
{
	if(BST_CHECKED== ((CButton*)GetDlgItem(IDC_CHK_FP0))->GetCheck())
		COM_AedAcqEx(0);
	if(BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP1))->GetCheck())
		COM_AedAcqEx(1);
}

void CDemoDlg::OnBnClickedBtnHstMulti()
{
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP0))->GetCheck())
		COM_HstAcqEx(0);
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP1))->GetCheck())
		COM_HstAcqEx(1);
}

void CDemoDlg::OnBnClickedBtnHsttriggetMulti()
{
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP0))->GetCheck())
		COM_TriggerEx(0);
	//	COM_PrepAcqEx(0);
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP1))->GetCheck())
		COM_TriggerEx(1);
	//	COM_PrepAcqEx(1);
}

void CDemoDlg::OnBnClickedBtnDstMulti()
{
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP0))->GetCheck())
		COM_DstEx(0);
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP1))->GetCheck())
		COM_DstEx(1);
}

void CDemoDlg::OnBnClickedBtnDacqMulti()
{
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP0))->GetCheck())
		COM_DacqEx(0);
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP1))->GetCheck())
		COM_DacqEx(1);
}

void CDemoDlg::OnBnClickedBtnDexitMulti()
{
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP0))->GetCheck())
		COM_DexitEx(0);
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP1))->GetCheck())
		COM_DexitEx(1);
}

void CDemoDlg::OnBnClickedBtnStopMulti()
{
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP0))->GetCheck())
		COM_StopEx(0);
	if (BST_CHECKED == ((CButton*)GetDlgItem(IDC_CHK_FP1))->GetCheck())
		COM_StopEx(1);
}
