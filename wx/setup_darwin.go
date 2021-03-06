package wx

// #cgo CPPFLAGS: -I${SRCDIR}/osx/ -I${SRCDIR}/../wxWidgets/wxWidgets-3.1.0/include -D_FILE_OFFSET_BITS=64 -D__WXOSX__ -D__WXOSX_COCOA__ -D__WXMAC__ -stdlib=libc++
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L${SRCDIR}/osx/lib -framework IOKit -framework Carbon -framework Cocoa -framework AudioToolbox -framework System -framework OpenGL -lwx_osx_cocoau_xrc-3.1 -lwx_osx_cocoau_webview-3.1 -lwx_osx_cocoau_stc-3.1 -lwx_osx_cocoau_richtext-3.1 -lwx_osx_cocoau_ribbon-3.1 -lwx_osx_cocoau_propgrid-3.1 -lwx_osx_cocoau_aui-3.1 -lwx_osx_cocoau_gl-3.1 -lwx_osx_cocoau_qa-3.1 -lwx_baseu_net-3.1 -lwx_osx_cocoau_html-3.1 -lwx_osx_cocoau_adv-3.1 -lwx_osx_cocoau_core-3.1 -lwx_baseu_xml-3.1 -lwx_baseu-3.1 -lwxscintilla-3.1 -framework OpenGL -framework AGL -framework WebKit -lexpat -lwxregexu-3.1 -lwxtiff-3.1 -lwxjpeg-3.1 -lwxpng-3.1 -lz -lpthread -liconv -llzma
// #cgo LDFLAGS: -Wl,-S
import "C"