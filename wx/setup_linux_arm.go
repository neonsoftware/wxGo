package wx

// #cgo CPPFLAGS: -I/root/wxWidgets/wxWidgets-3.1.0/lib -I${SRCDIR}/../wxWidgets/wxWidgets-3.1.0/include -D_FILE_OFFSET_BITS=64 -D__WXGTK__
// #cgo CXXFLAGS: -std=c++11 -D_GLIBCXX_USE_CXX11_ABI=0
// #cgo LDFLAGS: -L/root/wxWidgets/wxWidgets-3.1.0/lib -pthread  -lwx_gtk2u-3.1-arm-linux  -lGL -lGLU -lgthread-2.0 -pthread -lX11 -lXxf86vm -lSM -lgtk-x11-2.0 -lgdk-x11-2.0 -lpangocairo-1.0 -latk-1.0 -lcairo -lpangoft2-1.0 -lpango-1.0 -lfontconfig -lfreetype -lnotify -lgdk_pixbuf-2.0 -lgio-2.0 -lgobject-2.0 -lglib-2.0 -lpng -ljpeg -ltiff -lexpat -lz -ldl -lm
// #cgo LDFLAGS: -s
import "C"
