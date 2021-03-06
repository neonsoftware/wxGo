%{
    #include "wx/dc.h"
    #include "wx/dcclient.h"
    #include "wx/dcgraph.h"
    #include "wx/dcmemory.h"
    #include "wx/dcbuffer.h"
    #include "wx/dcmirror.h"
    #include "wx/dcprint.h"
    #include "wx/dcps.h"
    #include "wx/dcscreen.h"
    #include "wx/dcsvg.h"
%}

%include "wxGoInterface/dc.h"
%include "wxGoInterface/dcclient.h"
#ifndef __WXMSW__
%include "wxGoInterface/dcgraph.h"
%include "wxGoInterface/dcps.h"
#endif
%include "wxGoInterface/dcmemory.h"
%include "wxGoInterface/dcbuffer.h"
%include "wxGoInterface/dcmirror.h"
%include "wxGoInterface/dcprint.h"
%include "wxGoInterface/dcscreen.h"
%include "wxGoInterface/dcsvg.h"
