package gdk

/*
#include <stdlib.h>
#include <stdint.h>

typedef size_t GType;

extern void _Gdk_go_callback_cleanup(void *gofunc);
static void _c_callback_cleanup(void *userdata)
{
	_Gdk_go_callback_cleanup(userdata);
}
typedef struct _cairoContext cairoContext;
struct _cairoContext {};
typedef struct _cairoSurface cairoSurface;
struct _cairoSurface {};
typedef struct _cairoMatrix cairoMatrix;
struct _cairoMatrix {};
typedef struct _cairoPattern cairoPattern;
struct _cairoPattern {};
typedef struct _cairoRegion cairoRegion;
struct _cairoRegion {};
typedef uint32_t cairoContent;
typedef struct _cairoFontOptions cairoFontOptions;
struct _cairoFontOptions {};
typedef struct _cairoFontType cairoFontType;
struct _cairoFontType {};
typedef struct _cairoFontFace cairoFontFace;
struct _cairoFontFace {};
typedef struct _cairoScaledFont cairoScaledFont;
struct _cairoScaledFont {};
typedef struct _cairoPath cairoPath;
struct _cairoPath {};
typedef struct _cairoRectangleInt cairoRectangleInt;
struct _cairoRectangleInt { uint8_t _data[16]; };
typedef uint32_t PangoAlignment;
typedef struct _PangoAnalysis PangoAnalysis;
struct _PangoAnalysis { uint8_t _data[48]; };
typedef struct _PangoAttrClass PangoAttrClass;
struct _PangoAttrClass { uint8_t _data[32]; };
typedef struct _PangoAttrColor PangoAttrColor;
struct _PangoAttrColor { uint8_t _data[24]; };
typedef void* PangoAttrFilterFunc;
extern void _PangoAttrFilterFunc_c_wrapper();
extern void _PangoAttrFilterFunc_c_wrapper_once();
typedef struct _PangoAttrFloat PangoAttrFloat;
struct _PangoAttrFloat { uint8_t _data[24]; };
typedef struct _PangoAttrFontDesc PangoAttrFontDesc;
struct _PangoAttrFontDesc { uint8_t _data[24]; };
typedef struct _PangoAttrInt PangoAttrInt;
struct _PangoAttrInt { uint8_t _data[24]; };
typedef struct _PangoAttrIterator PangoAttrIterator;
struct _PangoAttrIterator {};
typedef struct _PangoAttrLanguage PangoAttrLanguage;
struct _PangoAttrLanguage { uint8_t _data[24]; };
typedef struct _PangoAttrList PangoAttrList;
struct _PangoAttrList {};
typedef struct _PangoAttrShape PangoAttrShape;
struct _PangoAttrShape { uint8_t _data[72]; };
typedef struct _PangoAttrSize PangoAttrSize;
struct _PangoAttrSize { uint8_t _data[24]; };
typedef struct _PangoAttrString PangoAttrString;
struct _PangoAttrString { uint8_t _data[24]; };
typedef uint32_t PangoAttrType;
typedef struct _PangoAttribute PangoAttribute;
struct _PangoAttribute { uint8_t _data[16]; };
typedef uint32_t PangoBidiType;
typedef struct _PangoColor PangoColor;
struct _PangoColor { uint8_t _data[6]; };
typedef struct _PangoContext PangoContext;
typedef struct _PangoContextClass PangoContextClass;
struct _PangoContextClass {};
typedef struct _PangoCoverage PangoCoverage;
struct _PangoCoverage {};
typedef uint32_t PangoCoverageLevel;
typedef uint32_t PangoDirection;
typedef uint32_t PangoEllipsizeMode;
typedef struct _PangoEngineLang PangoEngineLang;
struct _PangoEngineLang {};
typedef struct _PangoEngineShape PangoEngineShape;
struct _PangoEngineShape {};
typedef struct _PangoFont PangoFont;
typedef struct _PangoFontDescription PangoFontDescription;
struct _PangoFontDescription {};
typedef struct _PangoFontFace PangoFontFace;
typedef struct _PangoFontFamily PangoFontFamily;
typedef struct _PangoFontMap PangoFontMap;
typedef uint32_t PangoFontMask;
typedef struct _PangoFontMetrics PangoFontMetrics;
struct _PangoFontMetrics {};
typedef struct _PangoFontset PangoFontset;
typedef void* PangoFontsetForeachFunc;
extern void _PangoFontsetForeachFunc_c_wrapper();
extern void _PangoFontsetForeachFunc_c_wrapper_once();
typedef struct _PangoGlyphGeometry PangoGlyphGeometry;
struct _PangoGlyphGeometry { uint8_t _data[12]; };
typedef struct _PangoGlyphInfo PangoGlyphInfo;
struct _PangoGlyphInfo { uint8_t _data[20]; };
typedef struct _PangoGlyphItem PangoGlyphItem;
struct _PangoGlyphItem { uint8_t _data[16]; };
typedef struct _PangoGlyphItemIter PangoGlyphItemIter;
struct _PangoGlyphItemIter { uint8_t _data[40]; };
typedef struct _PangoGlyphString PangoGlyphString;
struct _PangoGlyphString { uint8_t _data[32]; };
typedef struct _PangoGlyphVisAttr PangoGlyphVisAttr;
struct _PangoGlyphVisAttr { uint8_t _data[4]; };
typedef uint32_t PangoGravity;
typedef uint32_t PangoGravityHint;
typedef struct _PangoItem PangoItem;
struct _PangoItem { uint8_t _data[64]; };
typedef struct _PangoLanguage PangoLanguage;
struct _PangoLanguage {};
typedef struct _PangoLayout PangoLayout;
typedef struct _PangoLayoutClass PangoLayoutClass;
struct _PangoLayoutClass {};
typedef struct _PangoLayoutIter PangoLayoutIter;
struct _PangoLayoutIter {};
typedef struct _PangoLayoutLine PangoLayoutLine;
struct _PangoLayoutLine { uint8_t _data[32]; };
typedef struct _PangoLogAttr PangoLogAttr;
struct _PangoLogAttr { uint8_t _data[52]; };
typedef struct _PangoMatrix PangoMatrix;
struct _PangoMatrix { uint8_t _data[48]; };
typedef struct _PangoRectangle PangoRectangle;
struct _PangoRectangle { uint8_t _data[16]; };
typedef uint32_t PangoRenderPart;
typedef struct _PangoRenderer PangoRenderer;
typedef struct _PangoRendererClass PangoRendererClass;
struct _PangoRendererClass { uint8_t _data[248]; };
typedef struct _PangoRendererPrivate PangoRendererPrivate;
struct _PangoRendererPrivate {};
typedef int32_t PangoScript;
typedef struct _PangoScriptIter PangoScriptIter;
struct _PangoScriptIter {};
typedef uint32_t PangoStretch;
typedef uint32_t PangoStyle;
typedef uint32_t PangoTabAlign;
typedef struct _PangoTabArray PangoTabArray;
struct _PangoTabArray {};
typedef uint32_t PangoUnderline;
typedef uint32_t PangoVariant;
typedef uint32_t PangoWeight;
typedef uint32_t PangoWrapMode;
typedef struct _Pango_ScriptForLang Pango_ScriptForLang;
struct _Pango_ScriptForLang { uint8_t _data[20]; };
typedef struct _GAction GAction;
typedef struct _GActionEntry GActionEntry;
struct _GActionEntry { uint8_t _data[64]; };
typedef struct _GActionGroup GActionGroup;
typedef struct _GActionGroupInterface GActionGroupInterface;
struct _GActionGroupInterface { uint8_t _data[120]; };
typedef struct _GActionInterface GActionInterface;
struct _GActionInterface { uint8_t _data[80]; };
typedef struct _GAppInfo GAppInfo;
typedef uint32_t GAppInfoCreateFlags;
typedef struct _GAppInfoIface GAppInfoIface;
struct _GAppInfoIface { uint8_t _data[192]; };
typedef struct _GAppLaunchContext GAppLaunchContext;
typedef struct _GAppLaunchContextClass GAppLaunchContextClass;
struct _GAppLaunchContextClass { uint8_t _data[200]; };
typedef struct _GAppLaunchContextPrivate GAppLaunchContextPrivate;
struct _GAppLaunchContextPrivate {};
typedef struct _GApplication GApplication;
typedef struct _GApplicationClass GApplicationClass;
struct _GApplicationClass { uint8_t _data[312]; };
typedef struct _GApplicationCommandLine GApplicationCommandLine;
typedef struct _GApplicationCommandLineClass GApplicationCommandLineClass;
struct _GApplicationCommandLineClass { uint8_t _data[248]; };
typedef struct _GApplicationCommandLinePrivate GApplicationCommandLinePrivate;
struct _GApplicationCommandLinePrivate {};
typedef uint32_t GApplicationFlags;
typedef struct _GApplicationPrivate GApplicationPrivate;
struct _GApplicationPrivate {};
typedef uint32_t GAskPasswordFlags;
typedef struct _GAsyncInitable GAsyncInitable;
typedef struct _GAsyncInitableIface GAsyncInitableIface;
struct _GAsyncInitableIface { uint8_t _data[32]; };
typedef void* GAsyncReadyCallback;
extern void _GAsyncReadyCallback_c_wrapper();
extern void _GAsyncReadyCallback_c_wrapper_once();
typedef struct _GAsyncResult GAsyncResult;
typedef struct _GAsyncResultIface GAsyncResultIface;
struct _GAsyncResultIface { uint8_t _data[32]; };
typedef struct _GBufferedInputStream GBufferedInputStream;
typedef struct _GBufferedInputStreamClass GBufferedInputStreamClass;
struct _GBufferedInputStreamClass { uint8_t _data[336]; };
typedef struct _GBufferedInputStreamPrivate GBufferedInputStreamPrivate;
struct _GBufferedInputStreamPrivate {};
typedef struct _GBufferedOutputStream GBufferedOutputStream;
typedef struct _GBufferedOutputStreamClass GBufferedOutputStreamClass;
struct _GBufferedOutputStreamClass { uint8_t _data[336]; };
typedef struct _GBufferedOutputStreamPrivate GBufferedOutputStreamPrivate;
struct _GBufferedOutputStreamPrivate {};
typedef void* GBusAcquiredCallback;
extern void _GBusAcquiredCallback_c_wrapper();
extern void _GBusAcquiredCallback_c_wrapper_once();
typedef void* GBusNameAcquiredCallback;
extern void _GBusNameAcquiredCallback_c_wrapper();
extern void _GBusNameAcquiredCallback_c_wrapper_once();
typedef void* GBusNameAppearedCallback;
extern void _GBusNameAppearedCallback_c_wrapper();
extern void _GBusNameAppearedCallback_c_wrapper_once();
typedef void* GBusNameLostCallback;
extern void _GBusNameLostCallback_c_wrapper();
extern void _GBusNameLostCallback_c_wrapper_once();
typedef uint32_t GBusNameOwnerFlags;
typedef void* GBusNameVanishedCallback;
extern void _GBusNameVanishedCallback_c_wrapper();
extern void _GBusNameVanishedCallback_c_wrapper_once();
typedef uint32_t GBusNameWatcherFlags;
typedef int32_t GBusType;
typedef struct _GCancellable GCancellable;
typedef struct _GCancellableClass GCancellableClass;
struct _GCancellableClass { uint8_t _data[184]; };
typedef struct _GCancellablePrivate GCancellablePrivate;
struct _GCancellablePrivate {};
typedef void* GCancellableSourceFunc;
extern void _GCancellableSourceFunc_c_wrapper();
extern void _GCancellableSourceFunc_c_wrapper_once();
typedef struct _GCharsetConverter GCharsetConverter;
typedef struct _GCharsetConverterClass GCharsetConverterClass;
struct _GCharsetConverterClass { uint8_t _data[136]; };
typedef struct _GConverter GConverter;
typedef uint32_t GConverterFlags;
typedef struct _GConverterIface GConverterIface;
struct _GConverterIface { uint8_t _data[32]; };
typedef struct _GConverterInputStream GConverterInputStream;
typedef struct _GConverterInputStreamClass GConverterInputStreamClass;
struct _GConverterInputStreamClass { uint8_t _data[312]; };
typedef struct _GConverterInputStreamPrivate GConverterInputStreamPrivate;
struct _GConverterInputStreamPrivate {};
typedef struct _GConverterOutputStream GConverterOutputStream;
typedef struct _GConverterOutputStreamClass GConverterOutputStreamClass;
struct _GConverterOutputStreamClass { uint8_t _data[360]; };
typedef struct _GConverterOutputStreamPrivate GConverterOutputStreamPrivate;
struct _GConverterOutputStreamPrivate {};
typedef uint32_t GConverterResult;
typedef struct _GCredentials GCredentials;
typedef struct _GCredentialsClass GCredentialsClass;
struct _GCredentialsClass {};
typedef uint32_t GCredentialsType;
typedef struct _GDBusAnnotationInfo GDBusAnnotationInfo;
struct _GDBusAnnotationInfo { uint8_t _data[32]; };
typedef struct _GDBusArgInfo GDBusArgInfo;
struct _GDBusArgInfo { uint8_t _data[32]; };
typedef struct _GDBusAuthObserver GDBusAuthObserver;
typedef uint32_t GDBusCallFlags;
typedef uint32_t GDBusCapabilityFlags;
typedef struct _GDBusConnection GDBusConnection;
typedef uint32_t GDBusConnectionFlags;
typedef uint32_t GDBusError;
typedef struct _GDBusErrorEntry GDBusErrorEntry;
struct _GDBusErrorEntry { uint8_t _data[16]; };
typedef struct _GDBusInterface GDBusInterface;
typedef void* GDBusInterfaceGetPropertyFunc;
extern void _GDBusInterfaceGetPropertyFunc_c_wrapper();
extern void _GDBusInterfaceGetPropertyFunc_c_wrapper_once();
typedef struct _GDBusInterfaceIface GDBusInterfaceIface;
struct _GDBusInterfaceIface { uint8_t _data[40]; };
typedef struct _GDBusInterfaceInfo GDBusInterfaceInfo;
struct _GDBusInterfaceInfo { uint8_t _data[48]; };
typedef void* GDBusInterfaceMethodCallFunc;
extern void _GDBusInterfaceMethodCallFunc_c_wrapper();
extern void _GDBusInterfaceMethodCallFunc_c_wrapper_once();
typedef void* GDBusInterfaceSetPropertyFunc;
extern void _GDBusInterfaceSetPropertyFunc_c_wrapper();
extern void _GDBusInterfaceSetPropertyFunc_c_wrapper_once();
typedef struct _GDBusInterfaceSkeleton GDBusInterfaceSkeleton;
typedef struct _GDBusInterfaceSkeletonClass GDBusInterfaceSkeletonClass;
struct _GDBusInterfaceSkeletonClass { uint8_t _data[304]; };
typedef uint32_t GDBusInterfaceSkeletonFlags;
typedef struct _GDBusInterfaceSkeletonPrivate GDBusInterfaceSkeletonPrivate;
struct _GDBusInterfaceSkeletonPrivate {};
typedef struct _GDBusInterfaceVTable GDBusInterfaceVTable;
struct _GDBusInterfaceVTable { uint8_t _data[88]; };
typedef struct _GDBusMessage GDBusMessage;
typedef uint32_t GDBusMessageByteOrder;
typedef void* GDBusMessageFilterFunction;
extern void _GDBusMessageFilterFunction_c_wrapper();
extern void _GDBusMessageFilterFunction_c_wrapper_once();
typedef uint32_t GDBusMessageFlags;
typedef uint32_t GDBusMessageHeaderField;
typedef uint32_t GDBusMessageType;
typedef struct _GDBusMethodInfo GDBusMethodInfo;
struct _GDBusMethodInfo { uint8_t _data[40]; };
typedef struct _GDBusMethodInvocation GDBusMethodInvocation;
typedef struct _GDBusNodeInfo GDBusNodeInfo;
struct _GDBusNodeInfo { uint8_t _data[40]; };
typedef struct _GDBusObject GDBusObject;
typedef struct _GDBusObjectIface GDBusObjectIface;
struct _GDBusObjectIface { uint8_t _data[56]; };
typedef struct _GDBusObjectManager GDBusObjectManager;
typedef struct _GDBusObjectManagerClient GDBusObjectManagerClient;
typedef struct _GDBusObjectManagerClientClass GDBusObjectManagerClientClass;
struct _GDBusObjectManagerClientClass { uint8_t _data[216]; };
typedef uint32_t GDBusObjectManagerClientFlags;
typedef struct _GDBusObjectManagerClientPrivate GDBusObjectManagerClientPrivate;
struct _GDBusObjectManagerClientPrivate {};
typedef struct _GDBusObjectManagerIface GDBusObjectManagerIface;
struct _GDBusObjectManagerIface { uint8_t _data[80]; };
typedef struct _GDBusObjectManagerServer GDBusObjectManagerServer;
typedef struct _GDBusObjectManagerServerClass GDBusObjectManagerServerClass;
struct _GDBusObjectManagerServerClass { uint8_t _data[200]; };
typedef struct _GDBusObjectManagerServerPrivate GDBusObjectManagerServerPrivate;
struct _GDBusObjectManagerServerPrivate {};
typedef struct _GDBusObjectProxy GDBusObjectProxy;
typedef struct _GDBusObjectProxyClass GDBusObjectProxyClass;
struct _GDBusObjectProxyClass { uint8_t _data[200]; };
typedef struct _GDBusObjectProxyPrivate GDBusObjectProxyPrivate;
struct _GDBusObjectProxyPrivate {};
typedef struct _GDBusObjectSkeleton GDBusObjectSkeleton;
typedef struct _GDBusObjectSkeletonClass GDBusObjectSkeletonClass;
struct _GDBusObjectSkeletonClass { uint8_t _data[208]; };
typedef struct _GDBusObjectSkeletonPrivate GDBusObjectSkeletonPrivate;
struct _GDBusObjectSkeletonPrivate {};
typedef struct _GDBusPropertyInfo GDBusPropertyInfo;
struct _GDBusPropertyInfo { uint8_t _data[40]; };
typedef uint32_t GDBusPropertyInfoFlags;
typedef struct _GDBusProxy GDBusProxy;
typedef struct _GDBusProxyClass GDBusProxyClass;
struct _GDBusProxyClass { uint8_t _data[408]; };
typedef uint32_t GDBusProxyFlags;
typedef struct _GDBusProxyPrivate GDBusProxyPrivate;
struct _GDBusProxyPrivate {};
typedef void* GDBusProxyTypeFunc;
extern void _GDBusProxyTypeFunc_c_wrapper();
extern void _GDBusProxyTypeFunc_c_wrapper_once();
typedef uint32_t GDBusSendMessageFlags;
typedef struct _GDBusServer GDBusServer;
typedef uint32_t GDBusServerFlags;
typedef void* GDBusSignalCallback;
extern void _GDBusSignalCallback_c_wrapper();
extern void _GDBusSignalCallback_c_wrapper_once();
typedef uint32_t GDBusSignalFlags;
typedef struct _GDBusSignalInfo GDBusSignalInfo;
struct _GDBusSignalInfo { uint8_t _data[32]; };
typedef void* GDBusSubtreeDispatchFunc;
extern void _GDBusSubtreeDispatchFunc_c_wrapper();
extern void _GDBusSubtreeDispatchFunc_c_wrapper_once();
typedef uint32_t GDBusSubtreeFlags;
typedef void* GDBusSubtreeIntrospectFunc;
extern void _GDBusSubtreeIntrospectFunc_c_wrapper();
extern void _GDBusSubtreeIntrospectFunc_c_wrapper_once();
typedef struct _GDBusSubtreeVTable GDBusSubtreeVTable;
struct _GDBusSubtreeVTable { uint8_t _data[88]; };
typedef struct _GDataInputStream GDataInputStream;
typedef struct _GDataInputStreamClass GDataInputStreamClass;
struct _GDataInputStreamClass { uint8_t _data[376]; };
typedef struct _GDataInputStreamPrivate GDataInputStreamPrivate;
struct _GDataInputStreamPrivate {};
typedef struct _GDataOutputStream GDataOutputStream;
typedef struct _GDataOutputStreamClass GDataOutputStreamClass;
struct _GDataOutputStreamClass { uint8_t _data[360]; };
typedef struct _GDataOutputStreamPrivate GDataOutputStreamPrivate;
struct _GDataOutputStreamPrivate {};
typedef uint32_t GDataStreamByteOrder;
typedef uint32_t GDataStreamNewlineType;
typedef struct _GDesktopAppInfo GDesktopAppInfo;
typedef struct _GDesktopAppInfoClass GDesktopAppInfoClass;
struct _GDesktopAppInfoClass { uint8_t _data[136]; };
typedef struct _GDesktopAppInfoLookup GDesktopAppInfoLookup;
typedef struct _GDesktopAppInfoLookupIface GDesktopAppInfoLookupIface;
struct _GDesktopAppInfoLookupIface { uint8_t _data[24]; };
typedef void* GDesktopAppLaunchCallback;
extern void _GDesktopAppLaunchCallback_c_wrapper();
extern void _GDesktopAppLaunchCallback_c_wrapper_once();
typedef struct _GDrive GDrive;
typedef struct _GDriveIface GDriveIface;
struct _GDriveIface { uint8_t _data[248]; };
typedef uint32_t GDriveStartFlags;
typedef uint32_t GDriveStartStopType;
typedef struct _GEmblem GEmblem;
typedef struct _GEmblemClass GEmblemClass;
struct _GEmblemClass {};
typedef uint32_t GEmblemOrigin;
typedef struct _GEmblemedIcon GEmblemedIcon;
typedef struct _GEmblemedIconClass GEmblemedIconClass;
struct _GEmblemedIconClass { uint8_t _data[136]; };
typedef struct _GEmblemedIconPrivate GEmblemedIconPrivate;
struct _GEmblemedIconPrivate {};
typedef struct _GFile GFile;
typedef struct _GFileAttributeInfo GFileAttributeInfo;
struct _GFileAttributeInfo { uint8_t _data[16]; };
typedef uint32_t GFileAttributeInfoFlags;
typedef struct _GFileAttributeInfoList GFileAttributeInfoList;
struct _GFileAttributeInfoList { uint8_t _data[16]; };
typedef struct _GFileAttributeMatcher GFileAttributeMatcher;
struct _GFileAttributeMatcher {};
typedef uint32_t GFileAttributeStatus;
typedef uint32_t GFileAttributeType;
typedef uint32_t GFileCopyFlags;
typedef uint32_t GFileCreateFlags;
typedef struct _GFileDescriptorBased GFileDescriptorBased;
typedef struct _GFileDescriptorBasedIface GFileDescriptorBasedIface;
struct _GFileDescriptorBasedIface { uint8_t _data[24]; };
typedef struct _GFileEnumerator GFileEnumerator;
typedef struct _GFileEnumeratorClass GFileEnumeratorClass;
struct _GFileEnumeratorClass { uint8_t _data[240]; };
typedef struct _GFileEnumeratorPrivate GFileEnumeratorPrivate;
struct _GFileEnumeratorPrivate {};
typedef struct _GFileIOStream GFileIOStream;
typedef struct _GFileIOStreamClass GFileIOStreamClass;
struct _GFileIOStreamClass { uint8_t _data[368]; };
typedef struct _GFileIOStreamPrivate GFileIOStreamPrivate;
struct _GFileIOStreamPrivate {};
typedef struct _GFileIcon GFileIcon;
typedef struct _GFileIconClass GFileIconClass;
struct _GFileIconClass {};
typedef struct _GFileIface GFileIface;
struct _GFileIface { uint8_t _data[816]; };
typedef struct _GFileInfo GFileInfo;
typedef struct _GFileInfoClass GFileInfoClass;
struct _GFileInfoClass {};
typedef struct _GFileInputStream GFileInputStream;
typedef struct _GFileInputStreamClass GFileInputStreamClass;
struct _GFileInputStreamClass { uint8_t _data[336]; };
typedef struct _GFileInputStreamPrivate GFileInputStreamPrivate;
struct _GFileInputStreamPrivate {};
typedef struct _GFileMonitor GFileMonitor;
typedef struct _GFileMonitorClass GFileMonitorClass;
struct _GFileMonitorClass { uint8_t _data[192]; };
typedef uint32_t GFileMonitorEvent;
typedef uint32_t GFileMonitorFlags;
typedef struct _GFileMonitorPrivate GFileMonitorPrivate;
struct _GFileMonitorPrivate {};
typedef struct _GFileOutputStream GFileOutputStream;
typedef struct _GFileOutputStreamClass GFileOutputStreamClass;
struct _GFileOutputStreamClass { uint8_t _data[408]; };
typedef struct _GFileOutputStreamPrivate GFileOutputStreamPrivate;
struct _GFileOutputStreamPrivate {};
typedef void* GFileProgressCallback;
extern void _GFileProgressCallback_c_wrapper();
extern void _GFileProgressCallback_c_wrapper_once();
typedef uint32_t GFileQueryInfoFlags;
typedef void* GFileReadMoreCallback;
extern void _GFileReadMoreCallback_c_wrapper();
extern void _GFileReadMoreCallback_c_wrapper_once();
typedef uint32_t GFileType;
typedef struct _GFilenameCompleter GFilenameCompleter;
typedef struct _GFilenameCompleterClass GFilenameCompleterClass;
struct _GFilenameCompleterClass { uint8_t _data[168]; };
typedef uint32_t GFilesystemPreviewType;
typedef struct _GFilterInputStream GFilterInputStream;
typedef struct _GFilterInputStreamClass GFilterInputStreamClass;
struct _GFilterInputStreamClass { uint8_t _data[272]; };
typedef struct _GFilterOutputStream GFilterOutputStream;
typedef struct _GFilterOutputStreamClass GFilterOutputStreamClass;
struct _GFilterOutputStreamClass { uint8_t _data[320]; };
typedef uint32_t GIOErrorEnum;
typedef struct _GIOExtension GIOExtension;
struct _GIOExtension {};
typedef struct _GIOExtensionPoint GIOExtensionPoint;
struct _GIOExtensionPoint {};
typedef struct _GIOModule GIOModule;
typedef struct _GIOModuleClass GIOModuleClass;
struct _GIOModuleClass {};
typedef struct _GIOModuleScope GIOModuleScope;
struct _GIOModuleScope {};
typedef uint32_t GIOModuleScopeFlags;
typedef struct _GIOSchedulerJob GIOSchedulerJob;
struct _GIOSchedulerJob {};
typedef void* GIOSchedulerJobFunc;
extern void _GIOSchedulerJobFunc_c_wrapper();
extern void _GIOSchedulerJobFunc_c_wrapper_once();
typedef struct _GIOStream GIOStream;
typedef struct _GIOStreamAdapter GIOStreamAdapter;
struct _GIOStreamAdapter {};
typedef struct _GIOStreamClass GIOStreamClass;
struct _GIOStreamClass { uint8_t _data[256]; };
typedef struct _GIOStreamPrivate GIOStreamPrivate;
struct _GIOStreamPrivate {};
typedef uint32_t GIOStreamSpliceFlags;
typedef struct _GIcon GIcon;
typedef struct _GIconIface GIconIface;
struct _GIconIface { uint8_t _data[48]; };
typedef struct _GInetAddress GInetAddress;
typedef struct _GInetAddressClass GInetAddressClass;
struct _GInetAddressClass { uint8_t _data[152]; };
typedef struct _GInetAddressPrivate GInetAddressPrivate;
struct _GInetAddressPrivate {};
typedef struct _GInetSocketAddress GInetSocketAddress;
typedef struct _GInetSocketAddressClass GInetSocketAddressClass;
struct _GInetSocketAddressClass { uint8_t _data[160]; };
typedef struct _GInetSocketAddressPrivate GInetSocketAddressPrivate;
struct _GInetSocketAddressPrivate {};
typedef struct _GInitable GInitable;
typedef struct _GInitableIface GInitableIface;
struct _GInitableIface { uint8_t _data[24]; };
typedef struct _GInputStream GInputStream;
typedef struct _GInputStreamClass GInputStreamClass;
struct _GInputStreamClass { uint8_t _data[248]; };
typedef struct _GInputStreamPrivate GInputStreamPrivate;
struct _GInputStreamPrivate {};
typedef struct _GInputVector GInputVector;
struct _GInputVector { uint8_t _data[16]; };
typedef struct _GLoadableIcon GLoadableIcon;
typedef struct _GLoadableIconIface GLoadableIconIface;
struct _GLoadableIconIface { uint8_t _data[40]; };
typedef struct _GMemoryInputStream GMemoryInputStream;
typedef struct _GMemoryInputStreamClass GMemoryInputStreamClass;
struct _GMemoryInputStreamClass { uint8_t _data[288]; };
typedef struct _GMemoryInputStreamPrivate GMemoryInputStreamPrivate;
struct _GMemoryInputStreamPrivate {};
typedef struct _GMemoryOutputStream GMemoryOutputStream;
typedef struct _GMemoryOutputStreamClass GMemoryOutputStreamClass;
struct _GMemoryOutputStreamClass { uint8_t _data[336]; };
typedef struct _GMemoryOutputStreamPrivate GMemoryOutputStreamPrivate;
struct _GMemoryOutputStreamPrivate {};
typedef struct _GMount GMount;
typedef struct _GMountIface GMountIface;
struct _GMountIface { uint8_t _data[216]; };
typedef uint32_t GMountMountFlags;
typedef struct _GMountOperation GMountOperation;
typedef struct _GMountOperationClass GMountOperationClass;
struct _GMountOperationClass { uint8_t _data[256]; };
typedef struct _GMountOperationPrivate GMountOperationPrivate;
struct _GMountOperationPrivate {};
typedef uint32_t GMountOperationResult;
typedef uint32_t GMountUnmountFlags;
typedef struct _GNativeVolumeMonitor GNativeVolumeMonitor;
typedef struct _GNativeVolumeMonitorClass GNativeVolumeMonitorClass;
struct _GNativeVolumeMonitorClass { uint8_t _data[344]; };
typedef struct _GNetworkAddress GNetworkAddress;
typedef struct _GNetworkAddressClass GNetworkAddressClass;
struct _GNetworkAddressClass { uint8_t _data[136]; };
typedef struct _GNetworkAddressPrivate GNetworkAddressPrivate;
struct _GNetworkAddressPrivate {};
typedef struct _GNetworkService GNetworkService;
typedef struct _GNetworkServiceClass GNetworkServiceClass;
struct _GNetworkServiceClass { uint8_t _data[136]; };
typedef struct _GNetworkServicePrivate GNetworkServicePrivate;
struct _GNetworkServicePrivate {};
typedef struct _GOutputStream GOutputStream;
typedef struct _GOutputStreamClass GOutputStreamClass;
struct _GOutputStreamClass { uint8_t _data[296]; };
typedef struct _GOutputStreamPrivate GOutputStreamPrivate;
struct _GOutputStreamPrivate {};
typedef uint32_t GOutputStreamSpliceFlags;
typedef struct _GOutputVector GOutputVector;
struct _GOutputVector { uint8_t _data[16]; };
typedef uint32_t GPasswordSave;
typedef struct _GPermission GPermission;
typedef struct _GPermissionClass GPermissionClass;
struct _GPermissionClass { uint8_t _data[312]; };
typedef struct _GPermissionPrivate GPermissionPrivate;
struct _GPermissionPrivate {};
typedef struct _GPollableInputStream GPollableInputStream;
typedef struct _GPollableInputStreamInterface GPollableInputStreamInterface;
struct _GPollableInputStreamInterface { uint8_t _data[48]; };
typedef struct _GPollableOutputStream GPollableOutputStream;
typedef struct _GPollableOutputStreamInterface GPollableOutputStreamInterface;
struct _GPollableOutputStreamInterface { uint8_t _data[48]; };
typedef void* GPollableSourceFunc;
extern void _GPollableSourceFunc_c_wrapper();
extern void _GPollableSourceFunc_c_wrapper_once();
typedef struct _GProxy GProxy;
typedef struct _GProxyAddress GProxyAddress;
typedef struct _GProxyAddressClass GProxyAddressClass;
struct _GProxyAddressClass { uint8_t _data[160]; };
typedef struct _GProxyAddressEnumerator GProxyAddressEnumerator;
typedef struct _GProxyAddressEnumeratorClass GProxyAddressEnumeratorClass;
struct _GProxyAddressEnumeratorClass { uint8_t _data[216]; };
typedef struct _GProxyAddressEnumeratorPrivate GProxyAddressEnumeratorPrivate;
struct _GProxyAddressEnumeratorPrivate {};
typedef struct _GProxyAddressPrivate GProxyAddressPrivate;
struct _GProxyAddressPrivate {};
typedef struct _GProxyInterface GProxyInterface;
struct _GProxyInterface { uint8_t _data[48]; };
typedef struct _GProxyResolver GProxyResolver;
typedef struct _GProxyResolverInterface GProxyResolverInterface;
struct _GProxyResolverInterface { uint8_t _data[48]; };
typedef struct _GResolver GResolver;
typedef struct _GResolverClass GResolverClass;
struct _GResolverClass { uint8_t _data[264]; };
typedef uint32_t GResolverError;
typedef struct _GResolverPrivate GResolverPrivate;
struct _GResolverPrivate {};
typedef struct _GSeekable GSeekable;
typedef struct _GSeekableIface GSeekableIface;
struct _GSeekableIface { uint8_t _data[56]; };
typedef struct _GSettings GSettings;
typedef struct _GSettingsBackend GSettingsBackend;
struct _GSettingsBackend {};
typedef uint32_t GSettingsBindFlags;
typedef void* GSettingsBindGetMapping;
extern void _GSettingsBindGetMapping_c_wrapper();
extern void _GSettingsBindGetMapping_c_wrapper_once();
typedef void* GSettingsBindSetMapping;
extern void _GSettingsBindSetMapping_c_wrapper();
extern void _GSettingsBindSetMapping_c_wrapper_once();
typedef struct _GSettingsClass GSettingsClass;
struct _GSettingsClass { uint8_t _data[328]; };
typedef void* GSettingsGetMapping;
extern void _GSettingsGetMapping_c_wrapper();
extern void _GSettingsGetMapping_c_wrapper_once();
typedef struct _GSettingsPrivate GSettingsPrivate;
struct _GSettingsPrivate {};
typedef struct _GSimpleAction GSimpleAction;
typedef struct _GSimpleActionGroup GSimpleActionGroup;
typedef struct _GSimpleActionGroupClass GSimpleActionGroupClass;
struct _GSimpleActionGroupClass { uint8_t _data[232]; };
typedef struct _GSimpleActionGroupPrivate GSimpleActionGroupPrivate;
struct _GSimpleActionGroupPrivate {};
typedef struct _GSimpleAsyncResult GSimpleAsyncResult;
typedef struct _GSimpleAsyncResultClass GSimpleAsyncResultClass;
struct _GSimpleAsyncResultClass {};
typedef void* GSimpleAsyncThreadFunc;
extern void _GSimpleAsyncThreadFunc_c_wrapper();
extern void _GSimpleAsyncThreadFunc_c_wrapper_once();
typedef struct _GSimplePermission GSimplePermission;
typedef struct _GSocket GSocket;
typedef struct _GSocketAddress GSocketAddress;
typedef struct _GSocketAddressClass GSocketAddressClass;
struct _GSocketAddressClass { uint8_t _data[160]; };
typedef struct _GSocketAddressEnumerator GSocketAddressEnumerator;
typedef struct _GSocketAddressEnumeratorClass GSocketAddressEnumeratorClass;
struct _GSocketAddressEnumeratorClass { uint8_t _data[160]; };
typedef struct _GSocketClass GSocketClass;
struct _GSocketClass { uint8_t _data[216]; };
typedef struct _GSocketClient GSocketClient;
typedef struct _GSocketClientClass GSocketClientClass;
struct _GSocketClientClass { uint8_t _data[176]; };
typedef struct _GSocketClientPrivate GSocketClientPrivate;
struct _GSocketClientPrivate {};
typedef struct _GSocketConnectable GSocketConnectable;
typedef struct _GSocketConnectableIface GSocketConnectableIface;
struct _GSocketConnectableIface { uint8_t _data[32]; };
typedef struct _GSocketConnection GSocketConnection;
typedef struct _GSocketConnectionClass GSocketConnectionClass;
struct _GSocketConnectionClass { uint8_t _data[304]; };
typedef struct _GSocketConnectionPrivate GSocketConnectionPrivate;
struct _GSocketConnectionPrivate {};
typedef struct _GSocketControlMessage GSocketControlMessage;
typedef struct _GSocketControlMessageClass GSocketControlMessageClass;
struct _GSocketControlMessageClass { uint8_t _data[216]; };
typedef struct _GSocketControlMessagePrivate GSocketControlMessagePrivate;
struct _GSocketControlMessagePrivate {};
typedef uint32_t GSocketFamily;
typedef struct _GSocketListener GSocketListener;
typedef struct _GSocketListenerClass GSocketListenerClass;
struct _GSocketListenerClass { uint8_t _data[192]; };
typedef struct _GSocketListenerPrivate GSocketListenerPrivate;
struct _GSocketListenerPrivate {};
typedef uint32_t GSocketMsgFlags;
typedef struct _GSocketPrivate GSocketPrivate;
struct _GSocketPrivate {};
typedef int32_t GSocketProtocol;
typedef struct _GSocketService GSocketService;
typedef struct _GSocketServiceClass GSocketServiceClass;
struct _GSocketServiceClass { uint8_t _data[248]; };
typedef struct _GSocketServicePrivate GSocketServicePrivate;
struct _GSocketServicePrivate {};
typedef void* GSocketSourceFunc;
extern void _GSocketSourceFunc_c_wrapper();
extern void _GSocketSourceFunc_c_wrapper_once();
typedef uint32_t GSocketType;
typedef struct _GSrvTarget GSrvTarget;
struct _GSrvTarget {};
typedef struct _GTcpConnection GTcpConnection;
typedef struct _GTcpConnectionClass GTcpConnectionClass;
struct _GTcpConnectionClass { uint8_t _data[304]; };
typedef struct _GTcpConnectionPrivate GTcpConnectionPrivate;
struct _GTcpConnectionPrivate {};
typedef struct _GTcpWrapperConnection GTcpWrapperConnection;
typedef struct _GTcpWrapperConnectionClass GTcpWrapperConnectionClass;
struct _GTcpWrapperConnectionClass { uint8_t _data[304]; };
typedef struct _GTcpWrapperConnectionPrivate GTcpWrapperConnectionPrivate;
struct _GTcpWrapperConnectionPrivate {};
typedef struct _GThemedIcon GThemedIcon;
typedef struct _GThemedIconClass GThemedIconClass;
struct _GThemedIconClass {};
typedef struct _GThreadedSocketService GThreadedSocketService;
typedef struct _GThreadedSocketServiceClass GThreadedSocketServiceClass;
struct _GThreadedSocketServiceClass { uint8_t _data[296]; };
typedef struct _GThreadedSocketServicePrivate GThreadedSocketServicePrivate;
struct _GThreadedSocketServicePrivate {};
typedef uint32_t GTlsAuthenticationMode;
typedef struct _GTlsBackend GTlsBackend;
typedef struct _GTlsBackendInterface GTlsBackendInterface;
struct _GTlsBackendInterface { uint8_t _data[64]; };
typedef struct _GTlsCertificate GTlsCertificate;
typedef struct _GTlsCertificateClass GTlsCertificateClass;
struct _GTlsCertificateClass { uint8_t _data[208]; };
typedef uint32_t GTlsCertificateFlags;
typedef struct _GTlsCertificatePrivate GTlsCertificatePrivate;
struct _GTlsCertificatePrivate {};
typedef struct _GTlsClientConnection GTlsClientConnection;
typedef struct _GTlsClientConnectionInterface GTlsClientConnectionInterface;
struct _GTlsClientConnectionInterface { uint8_t _data[16]; };
typedef struct _GTlsConnection GTlsConnection;
typedef struct _GTlsConnectionClass GTlsConnectionClass;
struct _GTlsConnectionClass { uint8_t _data[352]; };
typedef struct _GTlsConnectionPrivate GTlsConnectionPrivate;
struct _GTlsConnectionPrivate {};
typedef struct _GTlsDatabase GTlsDatabase;
typedef struct _GTlsDatabaseClass GTlsDatabaseClass;
struct _GTlsDatabaseClass { uint8_t _data[368]; };
typedef uint32_t GTlsDatabaseLookupFlags;
typedef struct _GTlsDatabasePrivate GTlsDatabasePrivate;
struct _GTlsDatabasePrivate {};
typedef uint32_t GTlsDatabaseVerifyFlags;
typedef uint32_t GTlsError;
typedef struct _GTlsFileDatabase GTlsFileDatabase;
typedef struct _GTlsFileDatabaseInterface GTlsFileDatabaseInterface;
struct _GTlsFileDatabaseInterface { uint8_t _data[80]; };
typedef struct _GTlsInteraction GTlsInteraction;
typedef struct _GTlsInteractionClass GTlsInteractionClass;
struct _GTlsInteractionClass { uint8_t _data[352]; };
typedef struct _GTlsInteractionPrivate GTlsInteractionPrivate;
struct _GTlsInteractionPrivate {};
typedef uint32_t GTlsInteractionResult;
typedef struct _GTlsPassword GTlsPassword;
typedef struct _GTlsPasswordClass GTlsPasswordClass;
struct _GTlsPasswordClass { uint8_t _data[192]; };
typedef uint32_t GTlsPasswordFlags;
typedef struct _GTlsPasswordPrivate GTlsPasswordPrivate;
struct _GTlsPasswordPrivate {};
typedef uint32_t GTlsRehandshakeMode;
typedef struct _GTlsServerConnection GTlsServerConnection;
typedef struct _GTlsServerConnectionInterface GTlsServerConnectionInterface;
struct _GTlsServerConnectionInterface { uint8_t _data[16]; };
typedef struct _GUnixConnection GUnixConnection;
typedef struct _GUnixConnectionClass GUnixConnectionClass;
struct _GUnixConnectionClass { uint8_t _data[304]; };
typedef struct _GUnixConnectionPrivate GUnixConnectionPrivate;
struct _GUnixConnectionPrivate {};
typedef struct _GUnixCredentialsMessage GUnixCredentialsMessage;
typedef struct _GUnixCredentialsMessageClass GUnixCredentialsMessageClass;
struct _GUnixCredentialsMessageClass { uint8_t _data[232]; };
typedef struct _GUnixCredentialsMessagePrivate GUnixCredentialsMessagePrivate;
struct _GUnixCredentialsMessagePrivate {};
typedef struct _GUnixFDList GUnixFDList;
typedef struct _GUnixFDListClass GUnixFDListClass;
struct _GUnixFDListClass { uint8_t _data[176]; };
typedef struct _GUnixFDListPrivate GUnixFDListPrivate;
struct _GUnixFDListPrivate {};
typedef struct _GUnixFDMessage GUnixFDMessage;
typedef struct _GUnixFDMessageClass GUnixFDMessageClass;
struct _GUnixFDMessageClass { uint8_t _data[232]; };
typedef struct _GUnixFDMessagePrivate GUnixFDMessagePrivate;
struct _GUnixFDMessagePrivate {};
typedef struct _GUnixInputStream GUnixInputStream;
typedef struct _GUnixInputStreamClass GUnixInputStreamClass;
struct _GUnixInputStreamClass { uint8_t _data[288]; };
typedef struct _GUnixInputStreamPrivate GUnixInputStreamPrivate;
struct _GUnixInputStreamPrivate {};
typedef struct _GUnixMountEntry GUnixMountEntry;
struct _GUnixMountEntry {};
typedef struct _GUnixMountMonitor GUnixMountMonitor;
typedef struct _GUnixMountMonitorClass GUnixMountMonitorClass;
struct _GUnixMountMonitorClass {};
typedef struct _GUnixMountPoint GUnixMountPoint;
struct _GUnixMountPoint {};
typedef struct _GUnixOutputStream GUnixOutputStream;
typedef struct _GUnixOutputStreamClass GUnixOutputStreamClass;
struct _GUnixOutputStreamClass { uint8_t _data[336]; };
typedef struct _GUnixOutputStreamPrivate GUnixOutputStreamPrivate;
struct _GUnixOutputStreamPrivate {};
typedef struct _GUnixSocketAddress GUnixSocketAddress;
typedef struct _GUnixSocketAddressClass GUnixSocketAddressClass;
struct _GUnixSocketAddressClass { uint8_t _data[160]; };
typedef struct _GUnixSocketAddressPrivate GUnixSocketAddressPrivate;
struct _GUnixSocketAddressPrivate {};
typedef uint32_t GUnixSocketAddressType;
typedef struct _GVfs GVfs;
typedef struct _GVfsClass GVfsClass;
struct _GVfsClass { uint8_t _data[272]; };
typedef struct _GVolume GVolume;
typedef struct _GVolumeIface GVolumeIface;
struct _GVolumeIface { uint8_t _data[168]; };
typedef struct _GVolumeMonitor GVolumeMonitor;
typedef struct _GVolumeMonitorClass GVolumeMonitorClass;
struct _GVolumeMonitorClass { uint8_t _data[336]; };
typedef struct _GZlibCompressor GZlibCompressor;
typedef struct _GZlibCompressorClass GZlibCompressorClass;
struct _GZlibCompressorClass { uint8_t _data[136]; };
typedef uint32_t GZlibCompressorFormat;
typedef struct _GZlibDecompressor GZlibDecompressor;
typedef struct _GZlibDecompressorClass GZlibDecompressorClass;
struct _GZlibDecompressorClass { uint8_t _data[136]; };
typedef uint32_t GdkColorspace;
typedef uint32_t GdkInterpType;
typedef struct _GdkPixbuf GdkPixbuf;
typedef uint32_t GdkPixbufAlphaMode;
typedef struct _GdkPixbufAnimation GdkPixbufAnimation;
typedef struct _GdkPixbufAnimationIter GdkPixbufAnimationIter;
typedef void* GdkPixbufDestroyNotify;
extern void _GdkPixbufDestroyNotify_c_wrapper();
extern void _GdkPixbufDestroyNotify_c_wrapper_once();
typedef uint32_t GdkPixbufError;
typedef struct _GdkPixbufFormat GdkPixbufFormat;
struct _GdkPixbufFormat {};
typedef struct _GdkPixbufLoader GdkPixbufLoader;
typedef struct _GdkPixbufLoaderClass GdkPixbufLoaderClass;
struct _GdkPixbufLoaderClass { uint8_t _data[168]; };
typedef uint32_t GdkPixbufRotation;
typedef void* GdkPixbufSaveFunc;
extern void _GdkPixbufSaveFunc_c_wrapper();
extern void _GdkPixbufSaveFunc_c_wrapper_once();
typedef struct _GdkPixbufSimpleAnim GdkPixbufSimpleAnim;
typedef struct _GdkPixbufSimpleAnimClass GdkPixbufSimpleAnimClass;
struct _GdkPixbufSimpleAnimClass {};
typedef struct _GdkPixbufSimpleAnimIter GdkPixbufSimpleAnimIter;
typedef struct _GdkPixdata GdkPixdata;
struct _GdkPixdata { uint8_t _data[32]; };
typedef uint32_t GdkPixdataDumpType;
typedef uint32_t GdkPixdataType;
typedef void* GBaseFinalizeFunc;
extern void _GBaseFinalizeFunc_c_wrapper();
extern void _GBaseFinalizeFunc_c_wrapper_once();
typedef void* GBaseInitFunc;
extern void _GBaseInitFunc_c_wrapper();
extern void _GBaseInitFunc_c_wrapper_once();
typedef struct _GBinding GBinding;
typedef uint32_t GBindingFlags;
typedef void* GBindingTransformFunc;
extern void _GBindingTransformFunc_c_wrapper();
extern void _GBindingTransformFunc_c_wrapper_once();
typedef void* GBoxedFreeFunc;
extern void _GBoxedFreeFunc_c_wrapper();
extern void _GBoxedFreeFunc_c_wrapper_once();
typedef struct _GCClosure GCClosure;
struct _GCClosure { uint8_t _data[72]; };
typedef void* GCallback;
extern void _GCallback_c_wrapper();
extern void _GCallback_c_wrapper_once();
typedef void* GClassFinalizeFunc;
extern void _GClassFinalizeFunc_c_wrapper();
extern void _GClassFinalizeFunc_c_wrapper_once();
typedef void* GClassInitFunc;
extern void _GClassInitFunc_c_wrapper();
extern void _GClassInitFunc_c_wrapper_once();
typedef struct _GClosure GClosure;
struct _GClosure { uint8_t _data[64]; };
typedef void* GClosureMarshal;
extern void _GClosureMarshal_c_wrapper();
extern void _GClosureMarshal_c_wrapper_once();
typedef void* GClosureNotify;
extern void _GClosureNotify_c_wrapper();
extern void _GClosureNotify_c_wrapper_once();
typedef struct _GClosureNotifyData GClosureNotifyData;
struct _GClosureNotifyData { uint8_t _data[16]; };
typedef uint32_t GConnectFlags;
typedef struct _GEnumClass GEnumClass;
struct _GEnumClass { uint8_t _data[32]; };
typedef struct _GEnumValue GEnumValue;
struct _GEnumValue { uint8_t _data[24]; };
typedef struct _GFlagsClass GFlagsClass;
struct _GFlagsClass { uint8_t _data[24]; };
typedef struct _GFlagsValue GFlagsValue;
struct _GFlagsValue { uint8_t _data[24]; };
typedef struct _GInitiallyUnowned GInitiallyUnowned;
typedef struct _GInitiallyUnownedClass GInitiallyUnownedClass;
struct _GInitiallyUnownedClass { uint8_t _data[136]; };
typedef void* GInstanceInitFunc;
extern void _GInstanceInitFunc_c_wrapper();
extern void _GInstanceInitFunc_c_wrapper_once();
typedef void* GInterfaceFinalizeFunc;
extern void _GInterfaceFinalizeFunc_c_wrapper();
extern void _GInterfaceFinalizeFunc_c_wrapper_once();
typedef struct _GInterfaceInfo GInterfaceInfo;
struct _GInterfaceInfo { uint8_t _data[24]; };
typedef void* GInterfaceInitFunc;
extern void _GInterfaceInitFunc_c_wrapper();
extern void _GInterfaceInitFunc_c_wrapper_once();
typedef struct _GObject GObject;
typedef struct _GObjectClass GObjectClass;
struct _GObjectClass { uint8_t _data[136]; };
typedef struct _GObjectConstructParam GObjectConstructParam;
struct _GObjectConstructParam { uint8_t _data[16]; };
typedef void* GObjectFinalizeFunc;
extern void _GObjectFinalizeFunc_c_wrapper();
extern void _GObjectFinalizeFunc_c_wrapper_once();
typedef void* GObjectGetPropertyFunc;
extern void _GObjectGetPropertyFunc_c_wrapper();
extern void _GObjectGetPropertyFunc_c_wrapper_once();
typedef void* GObjectSetPropertyFunc;
extern void _GObjectSetPropertyFunc_c_wrapper();
extern void _GObjectSetPropertyFunc_c_wrapper_once();
typedef uint32_t GParamFlags;
typedef struct _GParamSpec GParamSpec;
typedef struct _GParamSpecBoolean GParamSpecBoolean;
typedef struct _GParamSpecBoxed GParamSpecBoxed;
typedef struct _GParamSpecChar GParamSpecChar;
typedef struct _GParamSpecClass GParamSpecClass;
struct _GParamSpecClass { uint8_t _data[80]; };
typedef struct _GParamSpecDouble GParamSpecDouble;
typedef struct _GParamSpecEnum GParamSpecEnum;
typedef struct _GParamSpecFlags GParamSpecFlags;
typedef struct _GParamSpecFloat GParamSpecFloat;
typedef struct _GParamSpecGType GParamSpecGType;
typedef struct _GParamSpecInt GParamSpecInt;
typedef struct _GParamSpecInt64 GParamSpecInt64;
typedef struct _GParamSpecLong GParamSpecLong;
typedef struct _GParamSpecObject GParamSpecObject;
typedef struct _GParamSpecOverride GParamSpecOverride;
typedef struct _GParamSpecParam GParamSpecParam;
typedef struct _GParamSpecPointer GParamSpecPointer;
typedef struct _GParamSpecPool GParamSpecPool;
struct _GParamSpecPool {};
typedef struct _GParamSpecString GParamSpecString;
typedef struct _GParamSpecTypeInfo GParamSpecTypeInfo;
struct _GParamSpecTypeInfo { uint8_t _data[56]; };
typedef struct _GParamSpecUChar GParamSpecUChar;
typedef struct _GParamSpecUInt GParamSpecUInt;
typedef struct _GParamSpecUInt64 GParamSpecUInt64;
typedef struct _GParamSpecULong GParamSpecULong;
typedef struct _GParamSpecUnichar GParamSpecUnichar;
typedef struct _GParamSpecValueArray GParamSpecValueArray;
typedef struct _GParamSpecVariant GParamSpecVariant;
typedef struct _GParameter GParameter;
struct _GParameter { uint8_t _data[32]; };
typedef void* GSignalAccumulator;
extern void _GSignalAccumulator_c_wrapper();
extern void _GSignalAccumulator_c_wrapper_once();
typedef void* GSignalEmissionHook;
extern void _GSignalEmissionHook_c_wrapper();
extern void _GSignalEmissionHook_c_wrapper_once();
typedef uint32_t GSignalFlags;
typedef struct _GSignalInvocationHint GSignalInvocationHint;
struct _GSignalInvocationHint { uint8_t _data[12]; };
typedef uint32_t GSignalMatchType;
typedef struct _GSignalQuery GSignalQuery;
struct _GSignalQuery { uint8_t _data[56]; };
typedef void* GToggleNotify;
extern void _GToggleNotify_c_wrapper();
extern void _GToggleNotify_c_wrapper_once();
typedef struct _GTypeCValue GTypeCValue;
struct _GTypeCValue { uint8_t _data[8]; };
typedef struct _GTypeClass GTypeClass;
struct _GTypeClass { uint8_t _data[8]; };
typedef void* GTypeClassCacheFunc;
extern void _GTypeClassCacheFunc_c_wrapper();
extern void _GTypeClassCacheFunc_c_wrapper_once();
typedef uint32_t GTypeDebugFlags;
typedef uint32_t GTypeFlags;
typedef uint32_t GTypeFundamentalFlags;
typedef struct _GTypeFundamentalInfo GTypeFundamentalInfo;
struct _GTypeFundamentalInfo { uint8_t _data[4]; };
typedef struct _GTypeInfo GTypeInfo;
struct _GTypeInfo { uint8_t _data[72]; };
typedef struct _GTypeInstance GTypeInstance;
struct _GTypeInstance { uint8_t _data[8]; };
typedef struct _GTypeInterface GTypeInterface;
struct _GTypeInterface { uint8_t _data[16]; };
typedef void* GTypeInterfaceCheckFunc;
extern void _GTypeInterfaceCheckFunc_c_wrapper();
extern void _GTypeInterfaceCheckFunc_c_wrapper_once();
typedef struct _GTypeModule GTypeModule;
typedef struct _GTypeModuleClass GTypeModuleClass;
struct _GTypeModuleClass { uint8_t _data[184]; };
typedef struct _GTypePlugin GTypePlugin;
typedef struct _GTypePluginClass GTypePluginClass;
struct _GTypePluginClass { uint8_t _data[48]; };
typedef void* GTypePluginCompleteInterfaceInfo;
extern void _GTypePluginCompleteInterfaceInfo_c_wrapper();
extern void _GTypePluginCompleteInterfaceInfo_c_wrapper_once();
typedef void* GTypePluginCompleteTypeInfo;
extern void _GTypePluginCompleteTypeInfo_c_wrapper();
extern void _GTypePluginCompleteTypeInfo_c_wrapper_once();
typedef void* GTypePluginUnuse;
extern void _GTypePluginUnuse_c_wrapper();
extern void _GTypePluginUnuse_c_wrapper_once();
typedef void* GTypePluginUse;
extern void _GTypePluginUse_c_wrapper();
extern void _GTypePluginUse_c_wrapper_once();
typedef struct _GTypeQuery GTypeQuery;
struct _GTypeQuery { uint8_t _data[24]; };
typedef struct _GTypeValueTable GTypeValueTable;
struct _GTypeValueTable { uint8_t _data[64]; };
typedef struct _GValue GValue;
struct _GValue { uint8_t _data[24]; };
typedef struct _GValueArray GValueArray;
struct _GValueArray { uint8_t _data[24]; };
typedef void* GValueTransform;
extern void _GValueTransform_c_wrapper();
extern void _GValueTransform_c_wrapper_once();
typedef void* GWeakNotify;
extern void _GWeakNotify_c_wrapper();
extern void _GWeakNotify_c_wrapper_once();
typedef struct _G_Value__data__union G_Value__data__union;
struct _G_Value__data__union { uint8_t _data[8]; };
typedef struct _GModule GModule;
struct _GModule {};
typedef void* GModuleCheckInit;
extern void _GModuleCheckInit_c_wrapper();
extern void _GModuleCheckInit_c_wrapper_once();
typedef uint32_t GModuleFlags;
typedef void* GModuleUnload;
extern void _GModuleUnload_c_wrapper();
extern void _GModuleUnload_c_wrapper_once();
typedef struct _GAllocator GAllocator;
struct _GAllocator {};
typedef struct _GArray GArray;
struct _GArray { uint8_t _data[16]; };
typedef uint32_t GAsciiType;
typedef struct _GAsyncQueue GAsyncQueue;
struct _GAsyncQueue {};
typedef struct _GBookmarkFile GBookmarkFile;
struct _GBookmarkFile {};
typedef uint32_t GBookmarkFileError;
typedef struct _GByteArray GByteArray;
struct _GByteArray { uint8_t _data[16]; };
typedef struct _GCache GCache;
struct _GCache {};
typedef void* GCacheDestroyFunc;
extern void _GCacheDestroyFunc_c_wrapper();
extern void _GCacheDestroyFunc_c_wrapper_once();
typedef struct _GChecksum GChecksum;
struct _GChecksum {};
typedef uint32_t GChecksumType;
typedef void* GChildWatchFunc;
extern void _GChildWatchFunc_c_wrapper();
extern void _GChildWatchFunc_c_wrapper_once();
typedef void* GCompareDataFunc;
extern void _GCompareDataFunc_c_wrapper();
extern void _GCompareDataFunc_c_wrapper_once();
typedef void* GCompareFunc;
extern void _GCompareFunc_c_wrapper();
extern void _GCompareFunc_c_wrapper_once();
typedef struct _GCompletion GCompletion;
struct _GCompletion { uint8_t _data[40]; };
typedef void* GCompletionFunc;
extern void _GCompletionFunc_c_wrapper();
extern void _GCompletionFunc_c_wrapper_once();
typedef void* GCompletionStrncmpFunc;
extern void _GCompletionStrncmpFunc_c_wrapper();
extern void _GCompletionStrncmpFunc_c_wrapper_once();
typedef struct _GCond GCond;
struct _GCond {};
typedef uint32_t GConvertError;
typedef struct _GData GData;
struct _GData {};
typedef void* GDataForeachFunc;
extern void _GDataForeachFunc_c_wrapper();
extern void _GDataForeachFunc_c_wrapper_once();
typedef struct _GDate GDate;
struct _GDate { uint8_t _data[24]; };
typedef uint32_t GDateDMY;
typedef uint32_t GDateMonth;
typedef struct _GDateTime GDateTime;
struct _GDateTime {};
typedef uint32_t GDateWeekday;
typedef struct _GDebugKey GDebugKey;
struct _GDebugKey { uint8_t _data[16]; };
typedef void* GDestroyNotify;
extern void _GDestroyNotify_c_wrapper();
extern void _GDestroyNotify_c_wrapper_once();
typedef struct _GDir GDir;
struct _GDir {};
typedef struct _GDoubleIEEE754 GDoubleIEEE754;
struct _GDoubleIEEE754 { uint8_t _data[8]; };
typedef void* GEqualFunc;
extern void _GEqualFunc_c_wrapper();
extern void _GEqualFunc_c_wrapper_once();
typedef struct _GError GError;
struct _GError { uint8_t _data[16]; };
typedef uint32_t GErrorType;
typedef uint32_t GFileError;
typedef uint32_t GFileTest;
typedef struct _GFloatIEEE754 GFloatIEEE754;
struct _GFloatIEEE754 { uint8_t _data[4]; };
typedef uint32_t GFormatSizeFlags;
typedef void* GFreeFunc;
extern void _GFreeFunc_c_wrapper();
extern void _GFreeFunc_c_wrapper_once();
typedef void* GFunc;
extern void _GFunc_c_wrapper();
extern void _GFunc_c_wrapper_once();
typedef void* GHFunc;
extern void _GHFunc_c_wrapper();
extern void _GHFunc_c_wrapper_once();
typedef void* GHRFunc;
extern void _GHRFunc_c_wrapper();
extern void _GHRFunc_c_wrapper_once();
typedef void* GHashFunc;
extern void _GHashFunc_c_wrapper();
extern void _GHashFunc_c_wrapper_once();
typedef struct _GHashTable GHashTable;
struct _GHashTable {};
typedef struct _GHashTableIter GHashTableIter;
struct _GHashTableIter { uint8_t _data[40]; };
typedef struct _GHmac GHmac;
struct _GHmac {};
typedef struct _GHook GHook;
struct _GHook { uint8_t _data[64]; };
typedef void* GHookCheckFunc;
extern void _GHookCheckFunc_c_wrapper();
extern void _GHookCheckFunc_c_wrapper_once();
typedef void* GHookCheckMarshaller;
extern void _GHookCheckMarshaller_c_wrapper();
extern void _GHookCheckMarshaller_c_wrapper_once();
typedef void* GHookCompareFunc;
extern void _GHookCompareFunc_c_wrapper();
extern void _GHookCompareFunc_c_wrapper_once();
typedef void* GHookFinalizeFunc;
extern void _GHookFinalizeFunc_c_wrapper();
extern void _GHookFinalizeFunc_c_wrapper_once();
typedef void* GHookFindFunc;
extern void _GHookFindFunc_c_wrapper();
extern void _GHookFindFunc_c_wrapper_once();
typedef uint32_t GHookFlagMask;
typedef void* GHookFunc;
extern void _GHookFunc_c_wrapper();
extern void _GHookFunc_c_wrapper_once();
typedef struct _GHookList GHookList;
struct _GHookList { uint8_t _data[56]; };
typedef void* GHookMarshaller;
extern void _GHookMarshaller_c_wrapper();
extern void _GHookMarshaller_c_wrapper_once();
typedef struct _GIConv GIConv;
struct _GIConv {};
typedef struct _GIOChannel GIOChannel;
struct _GIOChannel { uint8_t _data[136]; };
typedef uint32_t GIOChannelError;
typedef uint32_t GIOCondition;
typedef uint32_t GIOError;
typedef uint32_t GIOFlags;
typedef void* GIOFunc;
extern void _GIOFunc_c_wrapper();
extern void _GIOFunc_c_wrapper_once();
typedef struct _GIOFuncs GIOFuncs;
struct _GIOFuncs { uint8_t _data[64]; };
typedef uint32_t GIOStatus;
typedef struct _GKeyFile GKeyFile;
struct _GKeyFile {};
typedef uint32_t GKeyFileError;
typedef uint32_t GKeyFileFlags;
typedef struct _GList GList;
struct _GList { uint8_t _data[24]; };
typedef void* GLogFunc;
extern void _GLogFunc_c_wrapper();
extern void _GLogFunc_c_wrapper_once();
typedef int32_t GLogLevelFlags;
typedef struct _GMainContext GMainContext;
struct _GMainContext {};
typedef struct _GMainLoop GMainLoop;
struct _GMainLoop {};
typedef struct _GMappedFile GMappedFile;
struct _GMappedFile {};
typedef uint32_t GMarkupCollectType;
typedef uint32_t GMarkupError;
typedef struct _GMarkupParseContext GMarkupParseContext;
struct _GMarkupParseContext {};
typedef uint32_t GMarkupParseFlags;
typedef struct _GMarkupParser GMarkupParser;
struct _GMarkupParser { uint8_t _data[40]; };
typedef struct _GMatchInfo GMatchInfo;
struct _GMatchInfo {};
typedef struct _GMemChunk GMemChunk;
struct _GMemChunk {};
typedef struct _GMemVTable GMemVTable;
struct _GMemVTable { uint8_t _data[48]; };
typedef struct _GMutex GMutex;
struct _GMutex {};
typedef struct _GNode GNode;
struct _GNode { uint8_t _data[40]; };
typedef void* GNodeForeachFunc;
extern void _GNodeForeachFunc_c_wrapper();
extern void _GNodeForeachFunc_c_wrapper_once();
typedef void* GNodeTraverseFunc;
extern void _GNodeTraverseFunc_c_wrapper();
extern void _GNodeTraverseFunc_c_wrapper_once();
typedef uint32_t GNormalizeMode;
typedef struct _GOnce GOnce;
struct _GOnce { uint8_t _data[16]; };
typedef uint32_t GOnceStatus;
typedef uint32_t GOptionArg;
typedef void* GOptionArgFunc;
extern void _GOptionArgFunc_c_wrapper();
extern void _GOptionArgFunc_c_wrapper_once();
typedef struct _GOptionContext GOptionContext;
struct _GOptionContext {};
typedef struct _GOptionEntry GOptionEntry;
struct _GOptionEntry { uint8_t _data[48]; };
typedef uint32_t GOptionError;
typedef void* GOptionErrorFunc;
extern void _GOptionErrorFunc_c_wrapper();
extern void _GOptionErrorFunc_c_wrapper_once();
typedef uint32_t GOptionFlags;
typedef struct _GOptionGroup GOptionGroup;
struct _GOptionGroup {};
typedef void* GOptionParseFunc;
extern void _GOptionParseFunc_c_wrapper();
extern void _GOptionParseFunc_c_wrapper_once();
typedef struct _GPatternSpec GPatternSpec;
struct _GPatternSpec {};
typedef struct _GPollFD GPollFD;
struct _GPollFD { uint8_t _data[8]; };
typedef void* GPollFunc;
extern void _GPollFunc_c_wrapper();
extern void _GPollFunc_c_wrapper_once();
typedef void* GPrintFunc;
extern void _GPrintFunc_c_wrapper();
extern void _GPrintFunc_c_wrapper_once();
typedef struct _GPrivate GPrivate;
struct _GPrivate {};
typedef struct _GPtrArray GPtrArray;
struct _GPtrArray { uint8_t _data[16]; };
typedef struct _GQueue GQueue;
struct _GQueue { uint8_t _data[24]; };
typedef struct _GRand GRand;
struct _GRand {};
typedef struct _GRegex GRegex;
struct _GRegex {};
typedef uint32_t GRegexCompileFlags;
typedef uint32_t GRegexError;
typedef void* GRegexEvalCallback;
extern void _GRegexEvalCallback_c_wrapper();
extern void _GRegexEvalCallback_c_wrapper_once();
typedef uint32_t GRegexMatchFlags;
typedef struct _GRelation GRelation;
struct _GRelation {};
typedef struct _GSList GSList;
struct _GSList { uint8_t _data[16]; };
typedef struct _GScanner GScanner;
struct _GScanner { uint8_t _data[144]; };
typedef struct _GScannerConfig GScannerConfig;
struct _GScannerConfig { uint8_t _data[128]; };
typedef void* GScannerMsgFunc;
extern void _GScannerMsgFunc_c_wrapper();
extern void _GScannerMsgFunc_c_wrapper_once();
typedef uint32_t GSeekType;
typedef struct _GSequence GSequence;
struct _GSequence {};
typedef struct _GSequenceIter GSequenceIter;
struct _GSequenceIter {};
typedef void* GSequenceIterCompareFunc;
extern void _GSequenceIterCompareFunc_c_wrapper();
extern void _GSequenceIterCompareFunc_c_wrapper_once();
typedef uint32_t GShellError;
typedef uint32_t GSliceConfig;
typedef struct _GSource GSource;
struct _GSource { uint8_t _data[96]; };
typedef struct _GSourceCallbackFuncs GSourceCallbackFuncs;
struct _GSourceCallbackFuncs { uint8_t _data[24]; };
typedef void* GSourceDummyMarshal;
extern void _GSourceDummyMarshal_c_wrapper();
extern void _GSourceDummyMarshal_c_wrapper_once();
typedef void* GSourceFunc;
extern void _GSourceFunc_c_wrapper();
extern void _GSourceFunc_c_wrapper_once();
typedef struct _GSourceFuncs GSourceFuncs;
struct _GSourceFuncs { uint8_t _data[48]; };
typedef struct _GSourcePrivate GSourcePrivate;
struct _GSourcePrivate {};
typedef void* GSpawnChildSetupFunc;
extern void _GSpawnChildSetupFunc_c_wrapper();
extern void _GSpawnChildSetupFunc_c_wrapper_once();
typedef uint32_t GSpawnError;
typedef uint32_t GSpawnFlags;
typedef struct _GStatBuf GStatBuf;
struct _GStatBuf {};
typedef struct _GStaticMutex GStaticMutex;
struct _GStaticMutex { uint8_t _data[8]; };
typedef struct _GStaticPrivate GStaticPrivate;
struct _GStaticPrivate { uint8_t _data[4]; };
typedef struct _GStaticRWLock GStaticRWLock;
struct _GStaticRWLock { uint8_t _data[40]; };
typedef struct _GStaticRecMutex GStaticRecMutex;
struct _GStaticRecMutex { uint8_t _data[24]; };
typedef struct _GString GString;
struct _GString { uint8_t _data[24]; };
typedef struct _GStringChunk GStringChunk;
struct _GStringChunk {};
typedef struct _GSystemThread GSystemThread;
struct _GSystemThread { uint8_t _data[8]; };
typedef struct _GTestCase GTestCase;
struct _GTestCase {};
typedef struct _GTestConfig GTestConfig;
struct _GTestConfig { uint8_t _data[20]; };
typedef void* GTestDataFunc;
extern void _GTestDataFunc_c_wrapper();
extern void _GTestDataFunc_c_wrapper_once();
typedef void* GTestFixtureFunc;
extern void _GTestFixtureFunc_c_wrapper();
extern void _GTestFixtureFunc_c_wrapper_once();
typedef void* GTestFunc;
extern void _GTestFunc_c_wrapper();
extern void _GTestFunc_c_wrapper_once();
typedef struct _GTestLogBuffer GTestLogBuffer;
struct _GTestLogBuffer { uint8_t _data[16]; };
typedef void* GTestLogFatalFunc;
extern void _GTestLogFatalFunc_c_wrapper();
extern void _GTestLogFatalFunc_c_wrapper_once();
typedef struct _GTestLogMsg GTestLogMsg;
struct _GTestLogMsg { uint8_t _data[32]; };
typedef uint32_t GTestLogType;
typedef struct _GTestSuite GTestSuite;
struct _GTestSuite {};
typedef uint32_t GTestTrapFlags;
typedef struct _GThread GThread;
struct _GThread { uint8_t _data[24]; };
typedef uint32_t GThreadError;
typedef struct _GThreadFunctions GThreadFunctions;
struct _GThreadFunctions { uint8_t _data[168]; };
typedef struct _GThreadPool GThreadPool;
struct _GThreadPool { uint8_t _data[24]; };
typedef uint32_t GThreadPriority;
typedef uint32_t GTimeType;
typedef struct _GTimeVal GTimeVal;
struct _GTimeVal { uint8_t _data[16]; };
typedef struct _GTimeZone GTimeZone;
struct _GTimeZone {};
typedef struct _GTimer GTimer;
struct _GTimer {};
typedef uint32_t GTokenType;
typedef struct _GTokenValue GTokenValue;
struct _GTokenValue { uint8_t _data[8]; };
typedef void* GTranslateFunc;
extern void _GTranslateFunc_c_wrapper();
extern void _GTranslateFunc_c_wrapper_once();
typedef struct _GTrashStack GTrashStack;
struct _GTrashStack { uint8_t _data[8]; };
typedef uint32_t GTraverseFlags;
typedef void* GTraverseFunc;
extern void _GTraverseFunc_c_wrapper();
extern void _GTraverseFunc_c_wrapper_once();
typedef uint32_t GTraverseType;
typedef struct _GTree GTree;
struct _GTree {};
typedef struct _GTuples GTuples;
struct _GTuples { uint8_t _data[4]; };
typedef uint32_t GUnicodeBreakType;
typedef int32_t GUnicodeScript;
typedef uint32_t GUnicodeType;
typedef uint32_t GUserDirectory;
typedef struct _GVariant GVariant;
struct _GVariant {};
typedef struct _GVariantBuilder GVariantBuilder;
struct _GVariantBuilder { uint8_t _data[128]; };
typedef uint32_t GVariantClass;
typedef uint32_t GVariantParseError;
typedef struct _GVariantType GVariantType;
struct _GVariantType {};
typedef void* GVoidFunc;
extern void _GVoidFunc_c_wrapper();
extern void _GVoidFunc_c_wrapper_once();
typedef struct _G_StaticAssert_347 G_StaticAssert_347;
struct _G_StaticAssert_347 { uint8_t _data[1]; };
typedef struct _GdkAppLaunchContext GdkAppLaunchContext;
typedef void *GdkAtom;
typedef uint32_t GdkAxisUse;
typedef uint32_t GdkByteOrder;
typedef struct _GdkColor GdkColor;
typedef uint32_t GdkCrossingMode;
typedef struct _GdkCursor GdkCursor;
typedef int32_t GdkCursorType;
typedef struct _GdkDevice GdkDevice;
typedef struct _GdkDeviceManager GdkDeviceManager;
typedef uint32_t GdkDeviceType;
typedef struct _GdkDisplay GdkDisplay;
typedef struct _GdkDisplayManager GdkDisplayManager;
typedef uint32_t GdkDragAction;
typedef struct _GdkDragContext GdkDragContext;
typedef uint32_t GdkDragProtocol;
typedef struct _GdkEvent GdkEvent;
typedef struct _GdkEventAny GdkEventAny;
typedef struct _GdkEventButton GdkEventButton;
typedef struct _GdkEventConfigure GdkEventConfigure;
typedef struct _GdkEventCrossing GdkEventCrossing;
typedef struct _GdkEventDND GdkEventDND;
typedef struct _GdkEventExpose GdkEventExpose;
typedef struct _GdkEventFocus GdkEventFocus;
typedef void* GdkEventFunc;
extern void _GdkEventFunc_c_wrapper();
extern void _GdkEventFunc_c_wrapper_once();
typedef struct _GdkEventGrabBroken GdkEventGrabBroken;
typedef struct _GdkEventKey GdkEventKey;
typedef uint32_t GdkEventMask;
typedef struct _GdkEventMotion GdkEventMotion;
typedef struct _GdkEventOwnerChange GdkEventOwnerChange;
typedef struct _GdkEventProperty GdkEventProperty;
typedef struct _GdkEventProximity GdkEventProximity;
typedef struct _GdkEventScroll GdkEventScroll;
typedef struct _GdkEventSelection GdkEventSelection;
typedef struct _GdkEventSetting GdkEventSetting;
typedef int32_t GdkEventType;
typedef struct _GdkEventVisibility GdkEventVisibility;
typedef struct _GdkEventWindowState GdkEventWindowState;
typedef uint32_t GdkExtensionMode;
typedef void* GdkFilterFunc;
extern void _GdkFilterFunc_c_wrapper();
extern void _GdkFilterFunc_c_wrapper_once();
typedef uint32_t GdkFilterReturn;
typedef struct _GdkGeometry GdkGeometry;
typedef uint32_t GdkGrabOwnership;
typedef uint32_t GdkGrabStatus;
typedef uint32_t GdkGravity;
typedef uint32_t GdkInputMode;
typedef uint32_t GdkInputSource;
typedef struct _GdkKeymap GdkKeymap;
typedef struct _GdkKeymapKey GdkKeymapKey;
typedef uint32_t GdkModifierType;
typedef uint32_t GdkNotifyType;
typedef uint32_t GdkOwnerChange;
typedef struct _GdkPoint GdkPoint;
typedef uint32_t GdkPropMode;
typedef uint32_t GdkPropertyState;
typedef struct _GdkRGBA GdkRGBA;
typedef struct _GdkScreen GdkScreen;
typedef uint32_t GdkScrollDirection;
typedef uint32_t GdkSettingAction;
typedef int32_t GdkStatus;
typedef struct _GdkTimeCoord GdkTimeCoord;
typedef uint32_t GdkVisibilityState;
typedef struct _GdkVisual GdkVisual;
typedef uint32_t GdkVisualType;
typedef uint32_t GdkWMDecoration;
typedef uint32_t GdkWMFunction;
typedef struct _GdkWindow GdkWindow;
typedef struct _GdkWindowAttr GdkWindowAttr;
typedef uint32_t GdkWindowAttributesType;
typedef void* GdkWindowChildFunc;
extern void _GdkWindowChildFunc_c_wrapper();
extern void _GdkWindowChildFunc_c_wrapper_once();
typedef struct _GdkWindowClass GdkWindowClass;
typedef uint32_t GdkWindowEdge;
typedef uint32_t GdkWindowHints;
typedef struct _GdkWindowRedirect GdkWindowRedirect;
typedef uint32_t GdkWindowState;
typedef uint32_t GdkWindowType;
typedef uint32_t GdkWindowTypeHint;
typedef uint32_t GdkWindowWindowClass;
extern GdkAppLaunchContext* gdk_app_launch_context_new();
extern void gdk_app_launch_context_set_desktop(GdkAppLaunchContext*, int32_t);
extern void gdk_app_launch_context_set_display(GdkAppLaunchContext*, GdkDisplay*);
extern void gdk_app_launch_context_set_icon(GdkAppLaunchContext*, GIcon*);
extern void gdk_app_launch_context_set_icon_name(GdkAppLaunchContext*, char*);
extern void gdk_app_launch_context_set_screen(GdkAppLaunchContext*, GdkScreen*);
extern void gdk_app_launch_context_set_timestamp(GdkAppLaunchContext*, uint32_t);
extern GType gdk_app_launch_context_get_type();
extern char* gdk_atom_name(GdkAtom);
extern GdkAtom gdk_atom_intern(char*, int);
extern GdkAtom gdk_atom_intern_static_string(char*);
extern GdkColor* gdk_color_copy(GdkColor*);
extern int gdk_color_equal(GdkColor*, GdkColor*);
extern void gdk_color_free(GdkColor*);
extern uint32_t gdk_color_hash(GdkColor*);
extern char* gdk_color_to_string(GdkColor*);
extern int gdk_color_parse(char*, GdkColor*);
extern GdkCursor* gdk_cursor_new(GdkCursorType);
extern GdkCursor* gdk_cursor_new_for_display(GdkDisplay*, GdkCursorType);
extern GdkCursor* gdk_cursor_new_from_name(GdkDisplay*, char*);
extern GdkCursor* gdk_cursor_new_from_pixbuf(GdkDisplay*, GdkPixbuf*, int32_t, int32_t);
extern GdkCursorType gdk_cursor_get_cursor_type(GdkCursor*);
extern GdkDisplay* gdk_cursor_get_display(GdkCursor*);
extern GdkPixbuf* gdk_cursor_get_image(GdkCursor*);
extern GdkCursor* gdk_cursor_ref(GdkCursor*);
extern void gdk_cursor_unref(GdkCursor*);
extern GType gdk_cursor_get_type();
extern int gdk_device_grab_info_libgtk_only(GdkDisplay*, GdkDevice*, GdkWindow**, int*);
extern GdkDevice* gdk_device_get_associated_device(GdkDevice*);
extern GdkAxisUse gdk_device_get_axis_use(GdkDevice*, uint32_t);
extern GdkDeviceType gdk_device_get_device_type(GdkDevice*);
extern GdkDisplay* gdk_device_get_display(GdkDevice*);
extern int gdk_device_get_has_cursor(GdkDevice*);
extern int gdk_device_get_key(GdkDevice*, uint32_t, uint32_t*, GdkModifierType*);
extern GdkInputMode gdk_device_get_mode(GdkDevice*);
extern int32_t gdk_device_get_n_axes(GdkDevice*);
extern int32_t gdk_device_get_n_keys(GdkDevice*);
extern char* gdk_device_get_name(GdkDevice*);
extern void gdk_device_get_position(GdkDevice*, GdkScreen**, int32_t*, int32_t*);
extern GdkInputSource gdk_device_get_source(GdkDevice*);
extern GdkWindow* gdk_device_get_window_at_position(GdkDevice*, int32_t*, int32_t*);
extern GdkGrabStatus gdk_device_grab(GdkDevice*, GdkWindow*, GdkGrabOwnership, int, GdkEventMask, GdkCursor*, uint32_t);
extern GList* gdk_device_list_axes(GdkDevice*);
extern GList* gdk_device_list_slave_devices(GdkDevice*);
extern void gdk_device_set_axis_use(GdkDevice*, uint32_t, GdkAxisUse);
extern void gdk_device_set_key(GdkDevice*, uint32_t, uint32_t, GdkModifierType);
extern int gdk_device_set_mode(GdkDevice*, GdkInputMode);
extern void gdk_device_ungrab(GdkDevice*, uint32_t);
extern void gdk_device_warp(GdkDevice*, GdkScreen*, int32_t, int32_t);
extern GType gdk_device_get_type();
extern GdkDevice* gdk_device_manager_get_client_pointer(GdkDeviceManager*);
extern GdkDisplay* gdk_device_manager_get_display(GdkDeviceManager*);
extern GList* gdk_device_manager_list_devices(GdkDeviceManager*, GdkDeviceType);
extern GType gdk_device_manager_get_type();
extern GdkDisplay* gdk_display_get_default();
extern GdkDisplay* gdk_display_open(char*);
extern GdkDisplay* gdk_display_open_default_libgtk_only();
extern void gdk_display_beep(GdkDisplay*);
extern void gdk_display_close(GdkDisplay*);
extern int gdk_display_device_is_grabbed(GdkDisplay*, GdkDevice*);
extern void gdk_display_flush(GdkDisplay*);
extern GdkAppLaunchContext* gdk_display_get_app_launch_context(GdkDisplay*);
extern uint32_t gdk_display_get_default_cursor_size(GdkDisplay*);
extern GdkWindow* gdk_display_get_default_group(GdkDisplay*);
extern GdkScreen* gdk_display_get_default_screen(GdkDisplay*);
extern GdkDeviceManager* gdk_display_get_device_manager(GdkDisplay*);
extern GdkEvent* gdk_display_get_event(GdkDisplay*);
extern void gdk_display_get_maximal_cursor_size(GdkDisplay*, uint32_t*, uint32_t*);
extern int32_t gdk_display_get_n_screens(GdkDisplay*);
extern char* gdk_display_get_name(GdkDisplay*);
extern void gdk_display_get_pointer(GdkDisplay*, GdkScreen**, int32_t*, int32_t*, GdkModifierType*);
extern GdkScreen* gdk_display_get_screen(GdkDisplay*, int32_t);
extern GdkWindow* gdk_display_get_window_at_pointer(GdkDisplay*, int32_t*, int32_t*);
extern int gdk_display_has_pending(GdkDisplay*);
extern int gdk_display_is_closed(GdkDisplay*);
extern void gdk_display_keyboard_ungrab(GdkDisplay*, uint32_t);
extern GList* gdk_display_list_devices(GdkDisplay*);
extern void gdk_display_notify_startup_complete(GdkDisplay*, char*);
extern GdkEvent* gdk_display_peek_event(GdkDisplay*);
extern int gdk_display_pointer_is_grabbed(GdkDisplay*);
extern void gdk_display_pointer_ungrab(GdkDisplay*, uint32_t);
extern void gdk_display_put_event(GdkDisplay*, GdkEvent*);
extern int gdk_display_request_selection_notification(GdkDisplay*, GdkAtom);
extern void gdk_display_set_double_click_distance(GdkDisplay*, uint32_t);
extern void gdk_display_set_double_click_time(GdkDisplay*, uint32_t);
extern void gdk_display_store_clipboard(GdkDisplay*, GdkWindow*, uint32_t, GdkAtom*, int32_t);
extern int gdk_display_supports_clipboard_persistence(GdkDisplay*);
extern int gdk_display_supports_composite(GdkDisplay*);
extern int gdk_display_supports_cursor_alpha(GdkDisplay*);
extern int gdk_display_supports_cursor_color(GdkDisplay*);
extern int gdk_display_supports_input_shapes(GdkDisplay*);
extern int gdk_display_supports_selection_notification(GdkDisplay*);
extern int gdk_display_supports_shapes(GdkDisplay*);
extern void gdk_display_sync(GdkDisplay*);
extern void gdk_display_warp_pointer(GdkDisplay*, GdkScreen*, int32_t, int32_t);
extern GType gdk_display_get_type();
extern GdkDisplayManager* gdk_display_manager_get();
extern GdkDisplay* gdk_display_manager_get_default_display(GdkDisplayManager*);
extern GSList* gdk_display_manager_list_displays(GdkDisplayManager*);
extern GdkDisplay* gdk_display_manager_open_display(GdkDisplayManager*, char*);
extern void gdk_display_manager_set_default_display(GdkDisplayManager*, GdkDisplay*);
extern GType gdk_display_manager_get_type();
extern GdkDragAction gdk_drag_context_get_actions(GdkDragContext*);
extern GdkWindow* gdk_drag_context_get_dest_window(GdkDragContext*);
extern GdkDevice* gdk_drag_context_get_device(GdkDragContext*);
extern GdkDragProtocol gdk_drag_context_get_protocol(GdkDragContext*);
extern GdkDragAction gdk_drag_context_get_selected_action(GdkDragContext*);
extern GdkWindow* gdk_drag_context_get_source_window(GdkDragContext*);
extern GdkDragAction gdk_drag_context_get_suggested_action(GdkDragContext*);
extern GList* gdk_drag_context_list_targets(GdkDragContext*);
extern void gdk_drag_context_set_device(GdkDragContext*, GdkDevice*);
extern GType gdk_drag_context_get_type();
extern GdkEvent* gdk_event_new(GdkEventType);
extern int gdk_events_get_angle(GdkEvent*, GdkEvent*, double*);
extern int gdk_events_get_center(GdkEvent*, GdkEvent*, double*, double*);
extern int gdk_events_get_distance(GdkEvent*, GdkEvent*, double*);
extern GdkEvent* gdk_event_copy(GdkEvent*);
extern void gdk_event_free(GdkEvent*);
extern int gdk_event_get_axis(GdkEvent*, GdkAxisUse, double*);
extern int gdk_event_get_button(GdkEvent*, uint32_t*);
extern int gdk_event_get_click_count(GdkEvent*, uint32_t*);
extern int gdk_event_get_coords(GdkEvent*, double*, double*);
extern GdkDevice* gdk_event_get_device(GdkEvent*);
extern int gdk_event_get_keycode(GdkEvent*, uint16_t*);
extern int gdk_event_get_keyval(GdkEvent*, uint32_t*);
extern int gdk_event_get_root_coords(GdkEvent*, double*, double*);
extern GdkScreen* gdk_event_get_screen(GdkEvent*);
extern int gdk_event_get_scroll_direction(GdkEvent*, GdkScrollDirection*);
extern GdkDevice* gdk_event_get_source_device(GdkEvent*);
extern int gdk_event_get_state(GdkEvent*, GdkModifierType*);
extern uint32_t gdk_event_get_time(GdkEvent*);
extern void gdk_event_put(GdkEvent*);
extern void gdk_event_set_device(GdkEvent*, GdkDevice*);
extern void gdk_event_set_screen(GdkEvent*, GdkScreen*);
extern void gdk_event_set_source_device(GdkEvent*, GdkDevice*);
extern GdkEvent* gdk_event_get();
extern void gdk_event_handler_set(GdkEventFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _gdk_event_handler_set(void* gofunc) {
	if (gofunc) {
		gdk_event_handler_set(_GdkEventFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		gdk_event_handler_set(0, 0, 0);
	}
}
extern GdkEvent* gdk_event_peek();
extern void gdk_event_request_motions(GdkEventMotion*);
extern GdkKeymap* gdk_keymap_get_default();
extern GdkKeymap* gdk_keymap_get_for_display(GdkDisplay*);
extern void gdk_keymap_add_virtual_modifiers(GdkKeymap*, GdkModifierType*);
extern int gdk_keymap_get_caps_lock_state(GdkKeymap*);
extern PangoDirection gdk_keymap_get_direction(GdkKeymap*);
extern int gdk_keymap_get_entries_for_keycode(GdkKeymap*, uint32_t, GdkKeymapKey**, uint32_t**, int32_t*);
extern int gdk_keymap_get_entries_for_keyval(GdkKeymap*, uint32_t, GdkKeymapKey**, int32_t*);
extern int gdk_keymap_get_num_lock_state(GdkKeymap*);
extern int gdk_keymap_have_bidi_layouts(GdkKeymap*);
extern uint32_t gdk_keymap_lookup_key(GdkKeymap*, GdkKeymapKey*);
extern int gdk_keymap_map_virtual_modifiers(GdkKeymap*, GdkModifierType*);
extern int gdk_keymap_translate_keyboard_state(GdkKeymap*, uint32_t, GdkModifierType, int32_t, uint32_t*, int32_t*, int32_t*, GdkModifierType*);
extern GType gdk_keymap_get_type();
extern GdkRGBA* gdk_rgba_copy(GdkRGBA*);
extern int gdk_rgba_equal(GdkRGBA*, GdkRGBA*);
extern void gdk_rgba_free(GdkRGBA*);
extern uint32_t gdk_rgba_hash(GdkRGBA*);
extern int gdk_rgba_parse(GdkRGBA*, char*);
extern char* gdk_rgba_to_string(GdkRGBA*);
extern GdkScreen* gdk_screen_get_default();
extern int32_t gdk_screen_height();
extern int32_t gdk_screen_height_mm();
extern int32_t gdk_screen_width();
extern int32_t gdk_screen_width_mm();
extern GdkWindow* gdk_screen_get_active_window(GdkScreen*);
extern GdkDisplay* gdk_screen_get_display(GdkScreen*);
extern cairoFontOptions* gdk_screen_get_font_options(GdkScreen*);
extern int32_t gdk_screen_get_height(GdkScreen*);
extern int32_t gdk_screen_get_height_mm(GdkScreen*);
extern int32_t gdk_screen_get_monitor_at_point(GdkScreen*, int32_t, int32_t);
extern int32_t gdk_screen_get_monitor_at_window(GdkScreen*, GdkWindow*);
extern void gdk_screen_get_monitor_geometry(GdkScreen*, int32_t, cairoRectangleInt*);
extern int32_t gdk_screen_get_monitor_height_mm(GdkScreen*, int32_t);
extern char* gdk_screen_get_monitor_plug_name(GdkScreen*, int32_t);
extern int32_t gdk_screen_get_monitor_width_mm(GdkScreen*, int32_t);
extern int32_t gdk_screen_get_n_monitors(GdkScreen*);
extern int32_t gdk_screen_get_number(GdkScreen*);
extern int32_t gdk_screen_get_primary_monitor(GdkScreen*);
extern double gdk_screen_get_resolution(GdkScreen*);
extern GdkVisual* gdk_screen_get_rgba_visual(GdkScreen*);
extern GdkWindow* gdk_screen_get_root_window(GdkScreen*);
extern int gdk_screen_get_setting(GdkScreen*, char*, GValue*);
extern GdkVisual* gdk_screen_get_system_visual(GdkScreen*);
extern GList* gdk_screen_get_toplevel_windows(GdkScreen*);
extern int32_t gdk_screen_get_width(GdkScreen*);
extern int32_t gdk_screen_get_width_mm(GdkScreen*);
extern GList* gdk_screen_get_window_stack(GdkScreen*);
extern int gdk_screen_is_composited(GdkScreen*);
extern GList* gdk_screen_list_visuals(GdkScreen*);
extern char* gdk_screen_make_display_name(GdkScreen*);
extern void gdk_screen_set_font_options(GdkScreen*, cairoFontOptions*);
extern void gdk_screen_set_resolution(GdkScreen*, double);
extern GType gdk_screen_get_type();
extern GdkVisual* gdk_visual_get_best();
extern int32_t gdk_visual_get_best_depth();
extern GdkVisualType gdk_visual_get_best_type();
extern GdkVisual* gdk_visual_get_best_with_both(int32_t, GdkVisualType);
extern GdkVisual* gdk_visual_get_best_with_depth(int32_t);
extern GdkVisual* gdk_visual_get_best_with_type(GdkVisualType);
extern GdkVisual* gdk_visual_get_system();
extern int32_t gdk_visual_get_bits_per_rgb(GdkVisual*);
extern void gdk_visual_get_blue_pixel_details(GdkVisual*, uint32_t*, int32_t*, int32_t*);
extern GdkByteOrder gdk_visual_get_byte_order(GdkVisual*);
extern int32_t gdk_visual_get_colormap_size(GdkVisual*);
extern int32_t gdk_visual_get_depth(GdkVisual*);
extern void gdk_visual_get_green_pixel_details(GdkVisual*, uint32_t*, int32_t*, int32_t*);
extern void gdk_visual_get_red_pixel_details(GdkVisual*, uint32_t*, int32_t*, int32_t*);
extern GdkScreen* gdk_visual_get_screen(GdkVisual*);
extern GdkVisualType gdk_visual_get_visual_type(GdkVisual*);
extern GType gdk_visual_get_type();
extern GdkWindow* gdk_window_new(GdkWindow*, GdkWindowAttr*, int32_t);
extern GdkWindow* gdk_window_at_pointer(int32_t*, int32_t*);
extern void gdk_window_constrain_size(GdkGeometry*, uint32_t, int32_t, int32_t, int32_t*, int32_t*);
extern void gdk_window_process_all_updates();
extern void gdk_window_set_debug_updates(int);
extern void gdk_window_beep(GdkWindow*);
extern void gdk_window_begin_move_drag(GdkWindow*, int32_t, int32_t, int32_t, uint32_t);
extern void gdk_window_begin_paint_rect(GdkWindow*, cairoRectangleInt*);
extern void gdk_window_begin_paint_region(GdkWindow*, cairoRegion*);
extern void gdk_window_begin_resize_drag(GdkWindow*, GdkWindowEdge, int32_t, int32_t, int32_t, uint32_t);
extern void gdk_window_configure_finished(GdkWindow*);
extern void gdk_window_coords_from_parent(GdkWindow*, double, double, double*, double*);
extern void gdk_window_coords_to_parent(GdkWindow*, double, double, double*, double*);
extern cairoSurface* gdk_window_create_similar_surface(GdkWindow*, cairoContent, int32_t, int32_t);
extern void gdk_window_deiconify(GdkWindow*);
extern void gdk_window_destroy(GdkWindow*);
extern void gdk_window_destroy_notify(GdkWindow*);
extern void gdk_window_enable_synchronized_configure(GdkWindow*);
extern void gdk_window_end_paint(GdkWindow*);
extern int gdk_window_ensure_native(GdkWindow*);
extern void gdk_window_flush(GdkWindow*);
extern void gdk_window_focus(GdkWindow*, uint32_t);
extern void gdk_window_freeze_toplevel_updates_libgtk_only(GdkWindow*);
extern void gdk_window_freeze_updates(GdkWindow*);
extern void gdk_window_fullscreen(GdkWindow*);
extern void gdk_window_geometry_changed(GdkWindow*);
extern int gdk_window_get_accept_focus(GdkWindow*);
extern cairoPattern* gdk_window_get_background_pattern(GdkWindow*);
extern GList* gdk_window_get_children(GdkWindow*);
extern cairoRegion* gdk_window_get_clip_region(GdkWindow*);
extern int gdk_window_get_composited(GdkWindow*);
extern GdkCursor* gdk_window_get_cursor(GdkWindow*);
extern int gdk_window_get_decorations(GdkWindow*, GdkWMDecoration*);
extern GdkCursor* gdk_window_get_device_cursor(GdkWindow*, GdkDevice*);
extern GdkEventMask gdk_window_get_device_events(GdkWindow*, GdkDevice*);
extern GdkWindow* gdk_window_get_device_position(GdkWindow*, GdkDevice*, int32_t*, int32_t*, GdkModifierType*);
extern GdkDisplay* gdk_window_get_display(GdkWindow*);
extern GdkDragProtocol gdk_window_get_drag_protocol(GdkWindow*, GdkWindow**);
extern GdkWindow* gdk_window_get_effective_parent(GdkWindow*);
extern GdkWindow* gdk_window_get_effective_toplevel(GdkWindow*);
extern GdkEventMask gdk_window_get_events(GdkWindow*);
extern int gdk_window_get_focus_on_map(GdkWindow*);
extern void gdk_window_get_frame_extents(GdkWindow*, cairoRectangleInt*);
extern void gdk_window_get_geometry(GdkWindow*, int32_t*, int32_t*, int32_t*, int32_t*);
extern GdkWindow* gdk_window_get_group(GdkWindow*);
extern int32_t gdk_window_get_height(GdkWindow*);
extern int gdk_window_get_modal_hint(GdkWindow*);
extern int32_t gdk_window_get_origin(GdkWindow*, int32_t*, int32_t*);
extern GdkWindow* gdk_window_get_parent(GdkWindow*);
extern GdkWindow* gdk_window_get_pointer(GdkWindow*, int32_t*, int32_t*, GdkModifierType*);
extern void gdk_window_get_position(GdkWindow*, int32_t*, int32_t*);
extern void gdk_window_get_root_coords(GdkWindow*, int32_t, int32_t, int32_t*, int32_t*);
extern void gdk_window_get_root_origin(GdkWindow*, int32_t*, int32_t*);
extern GdkScreen* gdk_window_get_screen(GdkWindow*);
extern GdkEventMask gdk_window_get_source_events(GdkWindow*, GdkInputSource);
extern GdkWindowState gdk_window_get_state(GdkWindow*);
extern int gdk_window_get_support_multidevice(GdkWindow*);
extern GdkWindow* gdk_window_get_toplevel(GdkWindow*);
extern GdkWindowTypeHint gdk_window_get_type_hint(GdkWindow*);
extern cairoRegion* gdk_window_get_update_area(GdkWindow*);
extern void gdk_window_get_user_data(GdkWindow*, void**);
extern cairoRegion* gdk_window_get_visible_region(GdkWindow*);
extern GdkVisual* gdk_window_get_visual(GdkWindow*);
extern int32_t gdk_window_get_width(GdkWindow*);
extern GdkWindowType gdk_window_get_window_type(GdkWindow*);
extern int gdk_window_has_native(GdkWindow*);
extern void gdk_window_hide(GdkWindow*);
extern void gdk_window_iconify(GdkWindow*);
extern void gdk_window_input_shape_combine_region(GdkWindow*, cairoRegion*, int32_t, int32_t);
extern void gdk_window_invalidate_maybe_recurse(GdkWindow*, cairoRegion*, GdkWindowChildFunc, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _gdk_window_invalidate_maybe_recurse(GdkWindow* this, cairoRegion* arg0, void* gofunc) {
	if (gofunc) {
		gdk_window_invalidate_maybe_recurse(this, arg0, _GdkWindowChildFunc_c_wrapper, gofunc);
	} else {
		gdk_window_invalidate_maybe_recurse(this, arg0, 0, 0);
	}
}
extern void gdk_window_invalidate_rect(GdkWindow*, cairoRectangleInt*, int);
extern void gdk_window_invalidate_region(GdkWindow*, cairoRegion*, int);
extern int gdk_window_is_destroyed(GdkWindow*);
extern int gdk_window_is_input_only(GdkWindow*);
extern int gdk_window_is_shaped(GdkWindow*);
extern int gdk_window_is_viewable(GdkWindow*);
extern int gdk_window_is_visible(GdkWindow*);
extern void gdk_window_lower(GdkWindow*);
extern void gdk_window_maximize(GdkWindow*);
extern void gdk_window_merge_child_input_shapes(GdkWindow*);
extern void gdk_window_merge_child_shapes(GdkWindow*);
extern void gdk_window_move(GdkWindow*, int32_t, int32_t);
extern void gdk_window_move_region(GdkWindow*, cairoRegion*, int32_t, int32_t);
extern void gdk_window_move_resize(GdkWindow*, int32_t, int32_t, int32_t, int32_t);
extern GList* gdk_window_peek_children(GdkWindow*);
extern void gdk_window_process_updates(GdkWindow*, int);
extern void gdk_window_raise(GdkWindow*);
extern void gdk_window_register_dnd(GdkWindow*);
extern void gdk_window_reparent(GdkWindow*, GdkWindow*, int32_t, int32_t);
extern void gdk_window_resize(GdkWindow*, int32_t, int32_t);
extern void gdk_window_restack(GdkWindow*, GdkWindow*, int);
extern void gdk_window_scroll(GdkWindow*, int32_t, int32_t);
extern void gdk_window_set_accept_focus(GdkWindow*, int);
extern void gdk_window_set_background(GdkWindow*, GdkColor*);
extern void gdk_window_set_background_pattern(GdkWindow*, cairoPattern*);
extern void gdk_window_set_background_rgba(GdkWindow*, GdkRGBA*);
extern void gdk_window_set_child_input_shapes(GdkWindow*);
extern void gdk_window_set_child_shapes(GdkWindow*);
extern void gdk_window_set_composited(GdkWindow*, int);
extern void gdk_window_set_cursor(GdkWindow*, GdkCursor*);
extern void gdk_window_set_decorations(GdkWindow*, GdkWMDecoration);
extern void gdk_window_set_device_cursor(GdkWindow*, GdkDevice*, GdkCursor*);
extern void gdk_window_set_device_events(GdkWindow*, GdkDevice*, GdkEventMask);
extern void gdk_window_set_events(GdkWindow*, GdkEventMask);
extern void gdk_window_set_focus_on_map(GdkWindow*, int);
extern void gdk_window_set_functions(GdkWindow*, GdkWMFunction);
extern void gdk_window_set_geometry_hints(GdkWindow*, GdkGeometry*, GdkWindowHints);
extern void gdk_window_set_group(GdkWindow*, GdkWindow*);
extern void gdk_window_set_icon_list(GdkWindow*, GList*);
extern void gdk_window_set_icon_name(GdkWindow*, char*);
extern void gdk_window_set_keep_above(GdkWindow*, int);
extern void gdk_window_set_keep_below(GdkWindow*, int);
extern void gdk_window_set_modal_hint(GdkWindow*, int);
extern void gdk_window_set_opacity(GdkWindow*, double);
extern void gdk_window_set_override_redirect(GdkWindow*, int);
extern void gdk_window_set_role(GdkWindow*, char*);
extern void gdk_window_set_skip_pager_hint(GdkWindow*, int);
extern void gdk_window_set_skip_taskbar_hint(GdkWindow*, int);
extern void gdk_window_set_source_events(GdkWindow*, GdkInputSource, GdkEventMask);
extern void gdk_window_set_startup_id(GdkWindow*, char*);
extern int gdk_window_set_static_gravities(GdkWindow*, int);
extern void gdk_window_set_support_multidevice(GdkWindow*, int);
extern void gdk_window_set_title(GdkWindow*, char*);
extern void gdk_window_set_transient_for(GdkWindow*, GdkWindow*);
extern void gdk_window_set_type_hint(GdkWindow*, GdkWindowTypeHint);
extern void gdk_window_set_urgency_hint(GdkWindow*, int);
extern void gdk_window_set_user_data(GdkWindow*, GObject*);
extern void gdk_window_shape_combine_region(GdkWindow*, cairoRegion*, int32_t, int32_t);
extern void gdk_window_show(GdkWindow*);
extern void gdk_window_show_unraised(GdkWindow*);
extern void gdk_window_stick(GdkWindow*);
extern void gdk_window_thaw_toplevel_updates_libgtk_only(GdkWindow*);
extern void gdk_window_thaw_updates(GdkWindow*);
extern void gdk_window_unfullscreen(GdkWindow*);
extern void gdk_window_unmaximize(GdkWindow*);
extern void gdk_window_unstick(GdkWindow*);
extern void gdk_window_withdraw(GdkWindow*);
extern GType gdk_window_get_type();
extern void gdk_add_option_entries_libgtk_only(GOptionGroup*);
extern void gdk_beep();
extern cairoContext* gdk_cairo_create(GdkWindow*);
extern int gdk_cairo_get_clip_rectangle(cairoContext*, cairoRectangleInt*);
extern void gdk_cairo_rectangle(cairoContext*, cairoRectangleInt*);
extern void gdk_cairo_region(cairoContext*, cairoRegion*);
extern cairoRegion* gdk_cairo_region_create_from_surface(cairoSurface*);
extern void gdk_cairo_set_source_color(cairoContext*, GdkColor*);
extern void gdk_cairo_set_source_pixbuf(cairoContext*, GdkPixbuf*, double, double);
extern void gdk_cairo_set_source_rgba(cairoContext*, GdkRGBA*);
extern void gdk_cairo_set_source_window(cairoContext*, GdkWindow*, double, double);
extern void gdk_disable_multidevice();
extern void gdk_drag_abort(GdkDragContext*, uint32_t);
extern GdkDragContext* gdk_drag_begin(GdkWindow*, GList*);
extern GdkDragContext* gdk_drag_begin_for_device(GdkWindow*, GdkDevice*, GList*);
extern void gdk_drag_drop(GdkDragContext*, uint32_t);
extern int gdk_drag_drop_succeeded(GdkDragContext*);
extern void gdk_drag_find_window_for_screen(GdkDragContext*, GdkWindow*, GdkScreen*, int32_t, int32_t, GdkWindow**, GdkDragProtocol*);
extern GdkAtom gdk_drag_get_selection(GdkDragContext*);
extern int gdk_drag_motion(GdkDragContext*, GdkWindow*, GdkDragProtocol, int32_t, int32_t, GdkDragAction, GdkDragAction, uint32_t);
extern void gdk_drag_status(GdkDragContext*, GdkDragAction, uint32_t);
extern void gdk_drop_finish(GdkDragContext*, int, uint32_t);
extern void gdk_drop_reply(GdkDragContext*, int, uint32_t);
extern int32_t gdk_error_trap_pop();
extern void gdk_error_trap_pop_ignored();
extern void gdk_error_trap_push();
extern int gdk_events_pending();
extern void gdk_flush();
extern GdkWindow* gdk_get_default_root_window();
extern char* gdk_get_display();
extern char* gdk_get_display_arg_name();
extern char* gdk_get_program_class();
extern int gdk_get_show_events();
extern void gdk_init(int32_t*, char***);
extern int gdk_init_check(int32_t*, char***);
extern GdkGrabStatus gdk_keyboard_grab(GdkWindow*, int, uint32_t);
extern void gdk_keyboard_ungrab(uint32_t);
extern void gdk_keyval_convert_case(uint32_t, uint32_t*, uint32_t*);
extern uint32_t gdk_keyval_from_name(char*);
extern int gdk_keyval_is_lower(uint32_t);
extern int gdk_keyval_is_upper(uint32_t);
extern char* gdk_keyval_name(uint32_t);
extern uint32_t gdk_keyval_to_lower(uint32_t);
extern uint32_t gdk_keyval_to_unicode(uint32_t);
extern uint32_t gdk_keyval_to_upper(uint32_t);
extern GList* gdk_list_visuals();
extern void gdk_notify_startup_complete();
extern void gdk_notify_startup_complete_with_id(char*);
extern GdkWindow* gdk_offscreen_window_get_embedder(GdkWindow*);
extern cairoSurface* gdk_offscreen_window_get_surface(GdkWindow*);
extern void gdk_offscreen_window_set_embedder(GdkWindow*, GdkWindow*);
extern PangoContext* gdk_pango_context_get();
extern PangoContext* gdk_pango_context_get_for_screen(GdkScreen*);
extern void gdk_parse_args(int32_t*, char***);
extern GdkPixbuf* gdk_pixbuf_get_from_surface(cairoSurface*, int32_t, int32_t, int32_t, int32_t);
extern GdkPixbuf* gdk_pixbuf_get_from_window(GdkWindow*, int32_t, int32_t, int32_t, int32_t);
extern GdkGrabStatus gdk_pointer_grab(GdkWindow*, int, GdkEventMask, GdkWindow*, GdkCursor*, uint32_t);
extern int gdk_pointer_is_grabbed();
extern void gdk_pointer_ungrab(uint32_t);
extern void gdk_pre_parse_libgtk_only();
extern void gdk_property_delete(GdkWindow*, GdkAtom);
extern int gdk_property_get(GdkWindow*, GdkAtom, GdkAtom, uint64_t, uint64_t, int32_t, GdkAtom*, int32_t*, int32_t*, uint8_t**);
extern void gdk_query_depths(int32_t**, int32_t*);
extern void gdk_query_visual_types(GdkVisualType**, int32_t*);
extern GType gdk_rectangle_get_type();
extern int gdk_rectangle_intersect(cairoRectangleInt*, cairoRectangleInt*, cairoRectangleInt*);
extern void gdk_rectangle_union(cairoRectangleInt*, cairoRectangleInt*, cairoRectangleInt*);
extern void gdk_selection_convert(GdkWindow*, GdkAtom, GdkAtom, uint32_t);
extern GdkWindow* gdk_selection_owner_get(GdkAtom);
extern GdkWindow* gdk_selection_owner_get_for_display(GdkDisplay*, GdkAtom);
extern int gdk_selection_owner_set(GdkWindow*, GdkAtom, uint32_t, int);
extern int gdk_selection_owner_set_for_display(GdkDisplay*, GdkWindow*, GdkAtom, uint32_t, int);
extern void gdk_selection_send_notify(GdkWindow*, GdkAtom, GdkAtom, GdkAtom, uint32_t);
extern void gdk_selection_send_notify_for_display(GdkDisplay*, GdkWindow*, GdkAtom, GdkAtom, GdkAtom, uint32_t);
extern void gdk_set_double_click_time(uint32_t);
extern void gdk_set_program_class(char*);
extern void gdk_set_show_events(int);
extern int gdk_setting_get(char*, GValue*);
extern void gdk_synthesize_window_state(GdkWindow*, GdkWindowState, GdkWindowState);
extern void gdk_test_render_sync(GdkWindow*);
extern int gdk_test_simulate_button(GdkWindow*, int32_t, int32_t, uint32_t, GdkModifierType, GdkEventType);
extern int gdk_test_simulate_key(GdkWindow*, int32_t, int32_t, uint32_t, GdkModifierType, GdkEventType);
extern int32_t gdk_text_property_to_utf8_list_for_display(GdkDisplay*, GdkAtom, int32_t, uint8_t*, int32_t, char***);
extern uint32_t gdk_threads_add_idle_full(int32_t, GSourceFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint32_t _gdk_threads_add_idle_full(int32_t arg0, void* gofunc) {
	if (gofunc) {
		return gdk_threads_add_idle_full(arg0, _GSourceFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return gdk_threads_add_idle_full(arg0, 0, 0, 0);
	}
}
extern uint32_t gdk_threads_add_timeout_full(int32_t, uint32_t, GSourceFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint32_t _gdk_threads_add_timeout_full(int32_t arg0, uint32_t arg1, void* gofunc) {
	if (gofunc) {
		return gdk_threads_add_timeout_full(arg0, arg1, _GSourceFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return gdk_threads_add_timeout_full(arg0, arg1, 0, 0, 0);
	}
}
extern uint32_t gdk_threads_add_timeout_seconds_full(int32_t, uint32_t, GSourceFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint32_t _gdk_threads_add_timeout_seconds_full(int32_t arg0, uint32_t arg1, void* gofunc) {
	if (gofunc) {
		return gdk_threads_add_timeout_seconds_full(arg0, arg1, _GSourceFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return gdk_threads_add_timeout_seconds_full(arg0, arg1, 0, 0, 0);
	}
}
extern void gdk_threads_enter();
extern void gdk_threads_init();
extern void gdk_threads_leave();
extern uint32_t gdk_unicode_to_keyval(uint32_t);
extern char* gdk_utf8_to_string_target(char*);
struct _GdkColor { uint8_t _data[12]; };
struct _GdkEvent { uint8_t _data[88]; };
struct _GdkEventAny { uint8_t _data[24]; };
struct _GdkEventButton { uint8_t _data[80]; };
struct _GdkEventConfigure { uint8_t _data[40]; };
struct _GdkEventCrossing { uint8_t _data[88]; };
struct _GdkEventDND { uint8_t _data[40]; };
struct _GdkEventExpose { uint8_t _data[56]; };
struct _GdkEventFocus { uint8_t _data[24]; };
struct _GdkEventGrabBroken { uint8_t _data[40]; };
struct _GdkEventKey { uint8_t _data[56]; };
struct _GdkEventMotion { uint8_t _data[80]; };
struct _GdkEventOwnerChange { uint8_t _data[56]; };
struct _GdkEventProperty { uint8_t _data[40]; };
struct _GdkEventProximity { uint8_t _data[32]; };
struct _GdkEventScroll { uint8_t _data[72]; };
struct _GdkEventSelection { uint8_t _data[64]; };
struct _GdkEventSetting { uint8_t _data[32]; };
struct _GdkEventVisibility { uint8_t _data[24]; };
struct _GdkEventWindowState { uint8_t _data[32]; };
struct _GdkGeometry { uint8_t _data[56]; };
struct _GdkKeymapKey { uint8_t _data[12]; };
struct _GdkPoint { uint8_t _data[8]; };
struct _GdkRGBA { uint8_t _data[32]; };
struct _GdkTimeCoord { uint8_t _data[1032]; };
struct _GdkWindowAttr { uint8_t _data[80]; };
struct _GdkWindowClass { uint8_t _data[232]; };
struct _GdkWindowRedirect {};


extern GObject *g_object_ref_sink(GObject*);
extern void g_object_unref(GObject*);
extern void g_free(void*);

#cgo pkg-config: gdk-3.0
*/
import "C"
import "unsafe"

import (
	"github.com/bytbox/gogobject/gobject-2.0"
	"github.com/bytbox/gogobject/cairo-1.0"
	"github.com/bytbox/gogobject/gdkpixbuf-2.0"
	"github.com/bytbox/gogobject/pango-1.0"
)

const alot = 999999

type _GSList struct {
	data unsafe.Pointer
	next *_GSList
}

type _GList struct {
	data unsafe.Pointer
	next *_GList
	prev *_GList
}

type _GError struct {
	domain uint32
	code int32
	message *C.char
}

func _GoStringToGString(x string) *C.char {
	if x == "\x00" {
		return nil
	}
	return C.CString(x)
}

func _GoBoolToCBool(x bool) C.int {
	if x { return 1 }
	return 0
}

func _CInterfaceToGoInterface(iface [2]unsafe.Pointer) interface{} {
	return *(*interface{})(unsafe.Pointer(&iface))
}

func _GoInterfaceToCInterface(iface interface{}) *unsafe.Pointer {
	return (*unsafe.Pointer)(unsafe.Pointer(&iface))
}


//export _Gdk_go_callback_cleanup
func _Gdk_go_callback_cleanup(gofunc unsafe.Pointer) {
	gobject.Holder.Release(gofunc)
}


// blacklisted: AppLaunchContext (object)
type Atom struct { Pointer unsafe.Pointer }
func (this0 Atom) Name() string {
	var this1 C.GdkAtom
	this1 = *(*C.GdkAtom)(unsafe.Pointer(&this0))
	ret1 := C.gdk_atom_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func AtomIntern(atom_name0 string, only_if_exists0 bool) Atom {
	var atom_name1 *C.char
	var only_if_exists1 C.int
	atom_name1 = _GoStringToGString(atom_name0)
	defer C.free(unsafe.Pointer(atom_name1))
	only_if_exists1 = _GoBoolToCBool(only_if_exists0)
	ret1 := C.gdk_atom_intern(atom_name1, only_if_exists1)
	var ret2 Atom
	ret2 = Atom{unsafe.Pointer(ret1)}
	return ret2
}
func AtomInternStaticString(atom_name0 string) Atom {
	var atom_name1 *C.char
	atom_name1 = _GoStringToGString(atom_name0)
	defer C.free(unsafe.Pointer(atom_name1))
	ret1 := C.gdk_atom_intern_static_string(atom_name1)
	var ret2 Atom
	ret2 = Atom{unsafe.Pointer(ret1)}
	return ret2
}
type AxisUse C.uint32_t
const (
	AxisUseIgnore AxisUse = 0
	AxisUseX AxisUse = 1
	AxisUseY AxisUse = 2
	AxisUsePressure AxisUse = 3
	AxisUseXtilt AxisUse = 4
	AxisUseYtilt AxisUse = 5
	AxisUseWheel AxisUse = 6
	AxisUseLast AxisUse = 7
)
type ByteOrder C.uint32_t
const (
	ByteOrderLsbFirst ByteOrder = 0
	ByteOrderMsbFirst ByteOrder = 1
)
const CurrentTime = 0
type Color struct {
	Pixel uint32
	Red uint16
	Green uint16
	Blue uint16
	_ [2]byte
}
func (this0 *Color) Copy() *Color {
	var this1 *C.GdkColor
	this1 = (*C.GdkColor)(unsafe.Pointer(this0))
	ret1 := C.gdk_color_copy(this1)
	var ret2 *Color
	ret2 = (*Color)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Color) Equal(colorb0 *Color) bool {
	var this1 *C.GdkColor
	var colorb1 *C.GdkColor
	this1 = (*C.GdkColor)(unsafe.Pointer(this0))
	colorb1 = (*C.GdkColor)(unsafe.Pointer(colorb0))
	ret1 := C.gdk_color_equal(this1, colorb1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Color) Free() {
	var this1 *C.GdkColor
	this1 = (*C.GdkColor)(unsafe.Pointer(this0))
	C.gdk_color_free(this1)
}
func (this0 *Color) Hash() int {
	var this1 *C.GdkColor
	this1 = (*C.GdkColor)(unsafe.Pointer(this0))
	ret1 := C.gdk_color_hash(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Color) ToString() string {
	var this1 *C.GdkColor
	this1 = (*C.GdkColor)(unsafe.Pointer(this0))
	ret1 := C.gdk_color_to_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func ColorParse(spec0 string) (Color, bool) {
	var spec1 *C.char
	var color1 C.GdkColor
	spec1 = _GoStringToGString(spec0)
	defer C.free(unsafe.Pointer(spec1))
	ret1 := C.gdk_color_parse(spec1, &color1)
	var color2 Color
	var ret2 bool
	color2 = *(*Color)(unsafe.Pointer(&color1))
	ret2 = ret1 != 0
	return color2, ret2
}
type CrossingMode C.uint32_t
const (
	CrossingModeNormal CrossingMode = 0
	CrossingModeGrab CrossingMode = 1
	CrossingModeUngrab CrossingMode = 2
	CrossingModeGtkGrab CrossingMode = 3
	CrossingModeGtkUngrab CrossingMode = 4
	CrossingModeStateChanged CrossingMode = 5
)
type CursorLike interface {
	gobject.ObjectLike
	InheritedFromGdkCursor() *C.GdkCursor
}

type Cursor struct {
	gobject.Object
	
}

func ToCursor(objlike gobject.ObjectLike) *Cursor {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Cursor)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Cursor)(obj)
	}
	panic("cannot cast to Cursor")
}

func (this0 *Cursor) InheritedFromGdkCursor() *C.GdkCursor {
	if this0 == nil {
		return nil
	}
	return (*C.GdkCursor)(this0.C)
}

func (this0 *Cursor) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_cursor_get_type())
}

func CursorGetType() gobject.Type {
	return (*Cursor)(nil).GetStaticType()
}
func NewCursor(cursor_type0 CursorType) *Cursor {
	var cursor_type1 C.GdkCursorType
	cursor_type1 = C.GdkCursorType(cursor_type0)
	ret1 := C.gdk_cursor_new(cursor_type1)
	var ret2 *Cursor
	ret2 = (*Cursor)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewCursorForDisplay(display0 DisplayLike, cursor_type0 CursorType) *Cursor {
	var display1 *C.GdkDisplay
	var cursor_type1 C.GdkCursorType
	if display0 != nil {
		display1 = display0.InheritedFromGdkDisplay()
	}
	cursor_type1 = C.GdkCursorType(cursor_type0)
	ret1 := C.gdk_cursor_new_for_display(display1, cursor_type1)
	var ret2 *Cursor
	ret2 = (*Cursor)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewCursorFromName(display0 DisplayLike, name0 string) *Cursor {
	var display1 *C.GdkDisplay
	var name1 *C.char
	if display0 != nil {
		display1 = display0.InheritedFromGdkDisplay()
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.gdk_cursor_new_from_name(display1, name1)
	var ret2 *Cursor
	ret2 = (*Cursor)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewCursorFromPixbuf(display0 DisplayLike, pixbuf0 gdkpixbuf.PixbufLike, x0 int, y0 int) *Cursor {
	var display1 *C.GdkDisplay
	var pixbuf1 *C.GdkPixbuf
	var x1 C.int32_t
	var y1 C.int32_t
	if display0 != nil {
		display1 = display0.InheritedFromGdkDisplay()
	}
	if pixbuf0 != nil {
		pixbuf1 = pixbuf0.InheritedFromGdkPixbuf()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	ret1 := C.gdk_cursor_new_from_pixbuf(display1, pixbuf1, x1, y1)
	var ret2 *Cursor
	ret2 = (*Cursor)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Cursor) GetCursorType() CursorType {
	var this1 *C.GdkCursor
	if this0 != nil {
		this1 = this0.InheritedFromGdkCursor()
	}
	ret1 := C.gdk_cursor_get_cursor_type(this1)
	var ret2 CursorType
	ret2 = CursorType(ret1)
	return ret2
}
func (this0 *Cursor) GetDisplay() *Display {
	var this1 *C.GdkCursor
	if this0 != nil {
		this1 = this0.InheritedFromGdkCursor()
	}
	ret1 := C.gdk_cursor_get_display(this1)
	var ret2 *Display
	ret2 = (*Display)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Cursor) GetImage() *gdkpixbuf.Pixbuf {
	var this1 *C.GdkCursor
	if this0 != nil {
		this1 = this0.InheritedFromGdkCursor()
	}
	ret1 := C.gdk_cursor_get_image(this1)
	var ret2 *gdkpixbuf.Pixbuf
	ret2 = (*gdkpixbuf.Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
// blacklisted: Cursor.ref (method)
// blacklisted: Cursor.unref (method)
type CursorType C.int32_t
const (
	CursorTypeXCursor CursorType = 0
	CursorTypeArrow CursorType = 2
	CursorTypeBasedArrowDown CursorType = 4
	CursorTypeBasedArrowUp CursorType = 6
	CursorTypeBoat CursorType = 8
	CursorTypeBogosity CursorType = 10
	CursorTypeBottomLeftCorner CursorType = 12
	CursorTypeBottomRightCorner CursorType = 14
	CursorTypeBottomSide CursorType = 16
	CursorTypeBottomTee CursorType = 18
	CursorTypeBoxSpiral CursorType = 20
	CursorTypeCenterPtr CursorType = 22
	CursorTypeCircle CursorType = 24
	CursorTypeClock CursorType = 26
	CursorTypeCoffeeMug CursorType = 28
	CursorTypeCross CursorType = 30
	CursorTypeCrossReverse CursorType = 32
	CursorTypeCrosshair CursorType = 34
	CursorTypeDiamondCross CursorType = 36
	CursorTypeDot CursorType = 38
	CursorTypeDotbox CursorType = 40
	CursorTypeDoubleArrow CursorType = 42
	CursorTypeDraftLarge CursorType = 44
	CursorTypeDraftSmall CursorType = 46
	CursorTypeDrapedBox CursorType = 48
	CursorTypeExchange CursorType = 50
	CursorTypeFleur CursorType = 52
	CursorTypeGobbler CursorType = 54
	CursorTypeGumby CursorType = 56
	CursorTypeHand1 CursorType = 58
	CursorTypeHand2 CursorType = 60
	CursorTypeHeart CursorType = 62
	CursorTypeIcon CursorType = 64
	CursorTypeIronCross CursorType = 66
	CursorTypeLeftPtr CursorType = 68
	CursorTypeLeftSide CursorType = 70
	CursorTypeLeftTee CursorType = 72
	CursorTypeLeftbutton CursorType = 74
	CursorTypeLlAngle CursorType = 76
	CursorTypeLrAngle CursorType = 78
	CursorTypeMan CursorType = 80
	CursorTypeMiddlebutton CursorType = 82
	CursorTypeMouse CursorType = 84
	CursorTypePencil CursorType = 86
	CursorTypePirate CursorType = 88
	CursorTypePlus CursorType = 90
	CursorTypeQuestionArrow CursorType = 92
	CursorTypeRightPtr CursorType = 94
	CursorTypeRightSide CursorType = 96
	CursorTypeRightTee CursorType = 98
	CursorTypeRightbutton CursorType = 100
	CursorTypeRtlLogo CursorType = 102
	CursorTypeSailboat CursorType = 104
	CursorTypeSbDownArrow CursorType = 106
	CursorTypeSbHDoubleArrow CursorType = 108
	CursorTypeSbLeftArrow CursorType = 110
	CursorTypeSbRightArrow CursorType = 112
	CursorTypeSbUpArrow CursorType = 114
	CursorTypeSbVDoubleArrow CursorType = 116
	CursorTypeShuttle CursorType = 118
	CursorTypeSizing CursorType = 120
	CursorTypeSpider CursorType = 122
	CursorTypeSpraycan CursorType = 124
	CursorTypeStar CursorType = 126
	CursorTypeTarget CursorType = 128
	CursorTypeTcross CursorType = 130
	CursorTypeTopLeftArrow CursorType = 132
	CursorTypeTopLeftCorner CursorType = 134
	CursorTypeTopRightCorner CursorType = 136
	CursorTypeTopSide CursorType = 138
	CursorTypeTopTee CursorType = 140
	CursorTypeTrek CursorType = 142
	CursorTypeUlAngle CursorType = 144
	CursorTypeUmbrella CursorType = 146
	CursorTypeUrAngle CursorType = 148
	CursorTypeWatch CursorType = 150
	CursorTypeXterm CursorType = 152
	CursorTypeLastCursor CursorType = 153
	CursorTypeBlankCursor CursorType = -2
	CursorTypeCursorIsPixmap CursorType = -1
)
type DeviceLike interface {
	gobject.ObjectLike
	InheritedFromGdkDevice() *C.GdkDevice
}

type Device struct {
	gobject.Object
	
}

func ToDevice(objlike gobject.ObjectLike) *Device {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Device)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Device)(obj)
	}
	panic("cannot cast to Device")
}

func (this0 *Device) InheritedFromGdkDevice() *C.GdkDevice {
	if this0 == nil {
		return nil
	}
	return (*C.GdkDevice)(this0.C)
}

func (this0 *Device) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_device_get_type())
}

func DeviceGetType() gobject.Type {
	return (*Device)(nil).GetStaticType()
}
func DeviceGrabInfoLibgtkOnly(display0 DisplayLike, device0 DeviceLike) (*Window, bool, bool) {
	var display1 *C.GdkDisplay
	var device1 *C.GdkDevice
	var grab_window1 *C.GdkWindow
	var owner_events1 C.int
	if display0 != nil {
		display1 = display0.InheritedFromGdkDisplay()
	}
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_grab_info_libgtk_only(display1, device1, &grab_window1, &owner_events1)
	var grab_window2 *Window
	var owner_events2 bool
	var ret2 bool
	grab_window2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(grab_window1), true))
	owner_events2 = owner_events1 != 0
	ret2 = ret1 != 0
	return grab_window2, owner_events2, ret2
}
func (this0 *Device) GetAssociatedDevice() *Device {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_get_associated_device(this1)
	var ret2 *Device
	ret2 = (*Device)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Device) GetAxisUse(index_0 int) AxisUse {
	var this1 *C.GdkDevice
	var index_1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	index_1 = C.uint32_t(index_0)
	ret1 := C.gdk_device_get_axis_use(this1, index_1)
	var ret2 AxisUse
	ret2 = AxisUse(ret1)
	return ret2
}
func (this0 *Device) GetDeviceType() DeviceType {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_get_device_type(this1)
	var ret2 DeviceType
	ret2 = DeviceType(ret1)
	return ret2
}
func (this0 *Device) GetDisplay() *Display {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_get_display(this1)
	var ret2 *Display
	ret2 = (*Display)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Device) GetHasCursor() bool {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_get_has_cursor(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Device) GetKey(index_0 int) (int, ModifierType, bool) {
	var this1 *C.GdkDevice
	var index_1 C.uint32_t
	var keyval1 C.uint32_t
	var modifiers1 C.GdkModifierType
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	index_1 = C.uint32_t(index_0)
	ret1 := C.gdk_device_get_key(this1, index_1, &keyval1, &modifiers1)
	var keyval2 int
	var modifiers2 ModifierType
	var ret2 bool
	keyval2 = int(keyval1)
	modifiers2 = ModifierType(modifiers1)
	ret2 = ret1 != 0
	return keyval2, modifiers2, ret2
}
func (this0 *Device) GetMode() InputMode {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_get_mode(this1)
	var ret2 InputMode
	ret2 = InputMode(ret1)
	return ret2
}
func (this0 *Device) GetNAxes() int {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_get_n_axes(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Device) GetNKeys() int {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_get_n_keys(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Device) GetName() string {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Device) GetPosition() (*Screen, int, int) {
	var this1 *C.GdkDevice
	var screen1 *C.GdkScreen
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	C.gdk_device_get_position(this1, &screen1, &x1, &y1)
	var screen2 *Screen
	var x2 int
	var y2 int
	screen2 = (*Screen)(gobject.ObjectWrap(unsafe.Pointer(screen1), true))
	x2 = int(x1)
	y2 = int(y1)
	return screen2, x2, y2
}
func (this0 *Device) GetSource() InputSource {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_get_source(this1)
	var ret2 InputSource
	ret2 = InputSource(ret1)
	return ret2
}
func (this0 *Device) GetWindowAtPosition() (int, int, *Window) {
	var this1 *C.GdkDevice
	var win_x1 C.int32_t
	var win_y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_get_window_at_position(this1, &win_x1, &win_y1)
	var win_x2 int
	var win_y2 int
	var ret2 *Window
	win_x2 = int(win_x1)
	win_y2 = int(win_y1)
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return win_x2, win_y2, ret2
}
func (this0 *Device) Grab(window0 WindowLike, grab_ownership0 GrabOwnership, owner_events0 bool, event_mask0 EventMask, cursor0 CursorLike, time_0 int) GrabStatus {
	var this1 *C.GdkDevice
	var window1 *C.GdkWindow
	var grab_ownership1 C.GdkGrabOwnership
	var owner_events1 C.int
	var event_mask1 C.GdkEventMask
	var cursor1 *C.GdkCursor
	var time_1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	grab_ownership1 = C.GdkGrabOwnership(grab_ownership0)
	owner_events1 = _GoBoolToCBool(owner_events0)
	event_mask1 = C.GdkEventMask(event_mask0)
	if cursor0 != nil {
		cursor1 = cursor0.InheritedFromGdkCursor()
	}
	time_1 = C.uint32_t(time_0)
	ret1 := C.gdk_device_grab(this1, window1, grab_ownership1, owner_events1, event_mask1, cursor1, time_1)
	var ret2 GrabStatus
	ret2 = GrabStatus(ret1)
	return ret2
}
func (this0 *Device) ListAxes() []Atom {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_list_axes(this1)
	var ret2 []Atom
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt Atom
		elt = *(*Atom)(unsafe.Pointer((*C.GdkAtom)(iter.data)))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Device) ListSlaveDevices() []*Device {
	var this1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_device_list_slave_devices(this1)
	var ret2 []*Device
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Device
		elt = (*Device)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkDevice)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Device) SetAxisUse(index_0 int, use0 AxisUse) {
	var this1 *C.GdkDevice
	var index_1 C.uint32_t
	var use1 C.GdkAxisUse
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	index_1 = C.uint32_t(index_0)
	use1 = C.GdkAxisUse(use0)
	C.gdk_device_set_axis_use(this1, index_1, use1)
}
func (this0 *Device) SetKey(index_0 int, keyval0 int, modifiers0 ModifierType) {
	var this1 *C.GdkDevice
	var index_1 C.uint32_t
	var keyval1 C.uint32_t
	var modifiers1 C.GdkModifierType
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	index_1 = C.uint32_t(index_0)
	keyval1 = C.uint32_t(keyval0)
	modifiers1 = C.GdkModifierType(modifiers0)
	C.gdk_device_set_key(this1, index_1, keyval1, modifiers1)
}
func (this0 *Device) SetMode(mode0 InputMode) bool {
	var this1 *C.GdkDevice
	var mode1 C.GdkInputMode
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	mode1 = C.GdkInputMode(mode0)
	ret1 := C.gdk_device_set_mode(this1, mode1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Device) Ungrab(time_0 int) {
	var this1 *C.GdkDevice
	var time_1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	time_1 = C.uint32_t(time_0)
	C.gdk_device_ungrab(this1, time_1)
}
func (this0 *Device) Warp(screen0 ScreenLike, x0 int, y0 int) {
	var this1 *C.GdkDevice
	var screen1 *C.GdkScreen
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDevice()
	}
	if screen0 != nil {
		screen1 = screen0.InheritedFromGdkScreen()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	C.gdk_device_warp(this1, screen1, x1, y1)
}
type DeviceManagerLike interface {
	gobject.ObjectLike
	InheritedFromGdkDeviceManager() *C.GdkDeviceManager
}

type DeviceManager struct {
	gobject.Object
	
}

func ToDeviceManager(objlike gobject.ObjectLike) *DeviceManager {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*DeviceManager)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*DeviceManager)(obj)
	}
	panic("cannot cast to DeviceManager")
}

func (this0 *DeviceManager) InheritedFromGdkDeviceManager() *C.GdkDeviceManager {
	if this0 == nil {
		return nil
	}
	return (*C.GdkDeviceManager)(this0.C)
}

func (this0 *DeviceManager) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_device_manager_get_type())
}

func DeviceManagerGetType() gobject.Type {
	return (*DeviceManager)(nil).GetStaticType()
}
func (this0 *DeviceManager) GetClientPointer() *Device {
	var this1 *C.GdkDeviceManager
	if this0 != nil {
		this1 = this0.InheritedFromGdkDeviceManager()
	}
	ret1 := C.gdk_device_manager_get_client_pointer(this1)
	var ret2 *Device
	ret2 = (*Device)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *DeviceManager) GetDisplay() *Display {
	var this1 *C.GdkDeviceManager
	if this0 != nil {
		this1 = this0.InheritedFromGdkDeviceManager()
	}
	ret1 := C.gdk_device_manager_get_display(this1)
	var ret2 *Display
	ret2 = (*Display)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *DeviceManager) ListDevices(type0 DeviceType) []*Device {
	var this1 *C.GdkDeviceManager
	var type1 C.GdkDeviceType
	if this0 != nil {
		this1 = this0.InheritedFromGdkDeviceManager()
	}
	type1 = C.GdkDeviceType(type0)
	ret1 := C.gdk_device_manager_list_devices(this1, type1)
	var ret2 []*Device
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Device
		elt = (*Device)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkDevice)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
type DeviceType C.uint32_t
const (
	DeviceTypeMaster DeviceType = 0
	DeviceTypeSlave DeviceType = 1
	DeviceTypeFloating DeviceType = 2
)
type DisplayLike interface {
	gobject.ObjectLike
	InheritedFromGdkDisplay() *C.GdkDisplay
}

type Display struct {
	gobject.Object
	
}

func ToDisplay(objlike gobject.ObjectLike) *Display {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Display)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Display)(obj)
	}
	panic("cannot cast to Display")
}

func (this0 *Display) InheritedFromGdkDisplay() *C.GdkDisplay {
	if this0 == nil {
		return nil
	}
	return (*C.GdkDisplay)(this0.C)
}

func (this0 *Display) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_display_get_type())
}

func DisplayGetType() gobject.Type {
	return (*Display)(nil).GetStaticType()
}
func DisplayGetDefault() *Display {
	ret1 := C.gdk_display_get_default()
	var ret2 *Display
	ret2 = (*Display)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func DisplayOpen(display_name0 string) *Display {
	var display_name1 *C.char
	display_name1 = _GoStringToGString(display_name0)
	defer C.free(unsafe.Pointer(display_name1))
	ret1 := C.gdk_display_open(display_name1)
	var ret2 *Display
	ret2 = (*Display)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func DisplayOpenDefaultLibgtkOnly() *Display {
	ret1 := C.gdk_display_open_default_libgtk_only()
	var ret2 *Display
	ret2 = (*Display)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Display) Beep() {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	C.gdk_display_beep(this1)
}
func (this0 *Display) Close() {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	C.gdk_display_close(this1)
}
func (this0 *Display) DeviceIsGrabbed(device0 DeviceLike) bool {
	var this1 *C.GdkDisplay
	var device1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_display_device_is_grabbed(this1, device1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) Flush() {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	C.gdk_display_flush(this1)
}
// blacklisted: Display.get_app_launch_context (method)
func (this0 *Display) GetDefaultCursorSize() int {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_get_default_cursor_size(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Display) GetDefaultGroup() *Window {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_get_default_group(this1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Display) GetDefaultScreen() *Screen {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_get_default_screen(this1)
	var ret2 *Screen
	ret2 = (*Screen)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Display) GetDeviceManager() *DeviceManager {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_get_device_manager(this1)
	var ret2 *DeviceManager
	ret2 = (*DeviceManager)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Display) GetEvent() *Event {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_get_event(this1)
	var ret2 *Event
	ret2 = (*Event)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Display) GetMaximalCursorSize() (int, int) {
	var this1 *C.GdkDisplay
	var width1 C.uint32_t
	var height1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	C.gdk_display_get_maximal_cursor_size(this1, &width1, &height1)
	var width2 int
	var height2 int
	width2 = int(width1)
	height2 = int(height1)
	return width2, height2
}
func (this0 *Display) GetNScreens() int {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_get_n_screens(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Display) GetName() string {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Display) GetPointer() (*Screen, int, int, ModifierType) {
	var this1 *C.GdkDisplay
	var screen1 *C.GdkScreen
	var x1 C.int32_t
	var y1 C.int32_t
	var mask1 C.GdkModifierType
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	C.gdk_display_get_pointer(this1, &screen1, &x1, &y1, &mask1)
	var screen2 *Screen
	var x2 int
	var y2 int
	var mask2 ModifierType
	screen2 = (*Screen)(gobject.ObjectWrap(unsafe.Pointer(screen1), true))
	x2 = int(x1)
	y2 = int(y1)
	mask2 = ModifierType(mask1)
	return screen2, x2, y2, mask2
}
func (this0 *Display) GetScreen(screen_num0 int) *Screen {
	var this1 *C.GdkDisplay
	var screen_num1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	screen_num1 = C.int32_t(screen_num0)
	ret1 := C.gdk_display_get_screen(this1, screen_num1)
	var ret2 *Screen
	ret2 = (*Screen)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Display) GetWindowAtPointer() (int, int, *Window) {
	var this1 *C.GdkDisplay
	var win_x1 C.int32_t
	var win_y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_get_window_at_pointer(this1, &win_x1, &win_y1)
	var win_x2 int
	var win_y2 int
	var ret2 *Window
	win_x2 = int(win_x1)
	win_y2 = int(win_y1)
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return win_x2, win_y2, ret2
}
func (this0 *Display) HasPending() bool {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_has_pending(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) IsClosed() bool {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_is_closed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) KeyboardUngrab(time_0 int) {
	var this1 *C.GdkDisplay
	var time_1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	time_1 = C.uint32_t(time_0)
	C.gdk_display_keyboard_ungrab(this1, time_1)
}
func (this0 *Display) ListDevices() []*Device {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_list_devices(this1)
	var ret2 []*Device
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Device
		elt = (*Device)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkDevice)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Display) NotifyStartupComplete(startup_id0 string) {
	var this1 *C.GdkDisplay
	var startup_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	startup_id1 = _GoStringToGString(startup_id0)
	defer C.free(unsafe.Pointer(startup_id1))
	C.gdk_display_notify_startup_complete(this1, startup_id1)
}
func (this0 *Display) PeekEvent() *Event {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_peek_event(this1)
	var ret2 *Event
	ret2 = (*Event)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Display) PointerIsGrabbed() bool {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_pointer_is_grabbed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) PointerUngrab(time_0 int) {
	var this1 *C.GdkDisplay
	var time_1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	time_1 = C.uint32_t(time_0)
	C.gdk_display_pointer_ungrab(this1, time_1)
}
func (this0 *Display) PutEvent(event0 *Event) {
	var this1 *C.GdkDisplay
	var event1 *C.GdkEvent
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	C.gdk_display_put_event(this1, event1)
}
func (this0 *Display) RequestSelectionNotification(selection0 Atom) bool {
	var this1 *C.GdkDisplay
	var selection1 C.GdkAtom
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	selection1 = *(*C.GdkAtom)(unsafe.Pointer(&selection0))
	ret1 := C.gdk_display_request_selection_notification(this1, selection1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) SetDoubleClickDistance(distance0 int) {
	var this1 *C.GdkDisplay
	var distance1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	distance1 = C.uint32_t(distance0)
	C.gdk_display_set_double_click_distance(this1, distance1)
}
func (this0 *Display) SetDoubleClickTime(msec0 int) {
	var this1 *C.GdkDisplay
	var msec1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	msec1 = C.uint32_t(msec0)
	C.gdk_display_set_double_click_time(this1, msec1)
}
func (this0 *Display) StoreClipboard(clipboard_window0 WindowLike, time_0 int, targets0 []Atom) {
	var this1 *C.GdkDisplay
	var clipboard_window1 *C.GdkWindow
	var time_1 C.uint32_t
	var targets1 *C.GdkAtom
	var n_targets1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	if clipboard_window0 != nil {
		clipboard_window1 = clipboard_window0.InheritedFromGdkWindow()
	}
	time_1 = C.uint32_t(time_0)
	targets1 = (*C.GdkAtom)(C.malloc(C.size_t(int(unsafe.Sizeof(*targets1)) * len(targets0))))
	defer C.free(unsafe.Pointer(targets1))
	for i, e := range targets0 {
		(*(*[999999]C.GdkAtom)(unsafe.Pointer(targets1)))[i] = *(*C.GdkAtom)(unsafe.Pointer(&e))
	}
	n_targets1 = C.int32_t(len(targets0))
	C.gdk_display_store_clipboard(this1, clipboard_window1, time_1, targets1, n_targets1)
}
func (this0 *Display) SupportsClipboardPersistence() bool {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_supports_clipboard_persistence(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) SupportsComposite() bool {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_supports_composite(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) SupportsCursorAlpha() bool {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_supports_cursor_alpha(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) SupportsCursorColor() bool {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_supports_cursor_color(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) SupportsInputShapes() bool {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_supports_input_shapes(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) SupportsSelectionNotification() bool {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_supports_selection_notification(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) SupportsShapes() bool {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_display_supports_shapes(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Display) Sync() {
	var this1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	C.gdk_display_sync(this1)
}
func (this0 *Display) WarpPointer(screen0 ScreenLike, x0 int, y0 int) {
	var this1 *C.GdkDisplay
	var screen1 *C.GdkScreen
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplay()
	}
	if screen0 != nil {
		screen1 = screen0.InheritedFromGdkScreen()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	C.gdk_display_warp_pointer(this1, screen1, x1, y1)
}
type DisplayManagerLike interface {
	gobject.ObjectLike
	InheritedFromGdkDisplayManager() *C.GdkDisplayManager
}

type DisplayManager struct {
	gobject.Object
	
}

func ToDisplayManager(objlike gobject.ObjectLike) *DisplayManager {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*DisplayManager)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*DisplayManager)(obj)
	}
	panic("cannot cast to DisplayManager")
}

func (this0 *DisplayManager) InheritedFromGdkDisplayManager() *C.GdkDisplayManager {
	if this0 == nil {
		return nil
	}
	return (*C.GdkDisplayManager)(this0.C)
}

func (this0 *DisplayManager) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_display_manager_get_type())
}

func DisplayManagerGetType() gobject.Type {
	return (*DisplayManager)(nil).GetStaticType()
}
func DisplayManagerGet() *DisplayManager {
	ret1 := C.gdk_display_manager_get()
	var ret2 *DisplayManager
	ret2 = (*DisplayManager)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *DisplayManager) GetDefaultDisplay() *Display {
	var this1 *C.GdkDisplayManager
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplayManager()
	}
	ret1 := C.gdk_display_manager_get_default_display(this1)
	var ret2 *Display
	ret2 = (*Display)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *DisplayManager) ListDisplays() []*Display {
	var this1 *C.GdkDisplayManager
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplayManager()
	}
	ret1 := C.gdk_display_manager_list_displays(this1)
	var ret2 []*Display
	for iter := (*_GSList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Display
		elt = (*Display)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkDisplay)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *DisplayManager) OpenDisplay(name0 string) *Display {
	var this1 *C.GdkDisplayManager
	var name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplayManager()
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.gdk_display_manager_open_display(this1, name1)
	var ret2 *Display
	ret2 = (*Display)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *DisplayManager) SetDefaultDisplay(display0 DisplayLike) {
	var this1 *C.GdkDisplayManager
	var display1 *C.GdkDisplay
	if this0 != nil {
		this1 = this0.InheritedFromGdkDisplayManager()
	}
	if display0 != nil {
		display1 = display0.InheritedFromGdkDisplay()
	}
	C.gdk_display_manager_set_default_display(this1, display1)
}
type DragAction C.uint32_t
const (
	DragActionDefault DragAction = 1
	DragActionCopy DragAction = 2
	DragActionMove DragAction = 4
	DragActionLink DragAction = 8
	DragActionPrivate DragAction = 16
	DragActionAsk DragAction = 32
)
type DragContextLike interface {
	gobject.ObjectLike
	InheritedFromGdkDragContext() *C.GdkDragContext
}

type DragContext struct {
	gobject.Object
	
}

func ToDragContext(objlike gobject.ObjectLike) *DragContext {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*DragContext)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*DragContext)(obj)
	}
	panic("cannot cast to DragContext")
}

func (this0 *DragContext) InheritedFromGdkDragContext() *C.GdkDragContext {
	if this0 == nil {
		return nil
	}
	return (*C.GdkDragContext)(this0.C)
}

func (this0 *DragContext) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_drag_context_get_type())
}

func DragContextGetType() gobject.Type {
	return (*DragContext)(nil).GetStaticType()
}
func (this0 *DragContext) GetActions() DragAction {
	var this1 *C.GdkDragContext
	if this0 != nil {
		this1 = this0.InheritedFromGdkDragContext()
	}
	ret1 := C.gdk_drag_context_get_actions(this1)
	var ret2 DragAction
	ret2 = DragAction(ret1)
	return ret2
}
func (this0 *DragContext) GetDestWindow() *Window {
	var this1 *C.GdkDragContext
	if this0 != nil {
		this1 = this0.InheritedFromGdkDragContext()
	}
	ret1 := C.gdk_drag_context_get_dest_window(this1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *DragContext) GetDevice() *Device {
	var this1 *C.GdkDragContext
	if this0 != nil {
		this1 = this0.InheritedFromGdkDragContext()
	}
	ret1 := C.gdk_drag_context_get_device(this1)
	var ret2 *Device
	ret2 = (*Device)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *DragContext) GetProtocol() DragProtocol {
	var this1 *C.GdkDragContext
	if this0 != nil {
		this1 = this0.InheritedFromGdkDragContext()
	}
	ret1 := C.gdk_drag_context_get_protocol(this1)
	var ret2 DragProtocol
	ret2 = DragProtocol(ret1)
	return ret2
}
func (this0 *DragContext) GetSelectedAction() DragAction {
	var this1 *C.GdkDragContext
	if this0 != nil {
		this1 = this0.InheritedFromGdkDragContext()
	}
	ret1 := C.gdk_drag_context_get_selected_action(this1)
	var ret2 DragAction
	ret2 = DragAction(ret1)
	return ret2
}
func (this0 *DragContext) GetSourceWindow() *Window {
	var this1 *C.GdkDragContext
	if this0 != nil {
		this1 = this0.InheritedFromGdkDragContext()
	}
	ret1 := C.gdk_drag_context_get_source_window(this1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *DragContext) GetSuggestedAction() DragAction {
	var this1 *C.GdkDragContext
	if this0 != nil {
		this1 = this0.InheritedFromGdkDragContext()
	}
	ret1 := C.gdk_drag_context_get_suggested_action(this1)
	var ret2 DragAction
	ret2 = DragAction(ret1)
	return ret2
}
func (this0 *DragContext) ListTargets() []Atom {
	var this1 *C.GdkDragContext
	if this0 != nil {
		this1 = this0.InheritedFromGdkDragContext()
	}
	ret1 := C.gdk_drag_context_list_targets(this1)
	var ret2 []Atom
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt Atom
		elt = *(*Atom)(unsafe.Pointer((*C.GdkAtom)(iter.data)))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *DragContext) SetDevice(device0 DeviceLike) {
	var this1 *C.GdkDragContext
	var device1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkDragContext()
	}
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	C.gdk_drag_context_set_device(this1, device1)
}
type DragProtocol C.uint32_t
const (
	DragProtocolNone DragProtocol = 0
	DragProtocolMotif DragProtocol = 1
	DragProtocolXdnd DragProtocol = 2
	DragProtocolRootwin DragProtocol = 3
	DragProtocolWin32Dropfiles DragProtocol = 4
	DragProtocolOle2 DragProtocol = 5
	DragProtocolLocal DragProtocol = 6
)
type Event struct {
	_data [88]byte
}
func NewEvent(type0 EventType) *Event {
	var type1 C.GdkEventType
	type1 = C.GdkEventType(type0)
	ret1 := C.gdk_event_new(type1)
	var ret2 *Event
	ret2 = (*Event)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Event) _GetAngle(event20 *Event) (float64, bool) {
	var this1 *C.GdkEvent
	var event21 *C.GdkEvent
	var angle1 C.double
	ret1 := C.gdk_events_get_angle(this1, event21, &angle1)
	var angle2 float64
	var ret2 bool
	angle2 = float64(angle1)
	ret2 = ret1 != 0
	return angle2, ret2
}
func (this0 *Event) _GetCenter(event20 *Event) (float64, float64, bool) {
	var this1 *C.GdkEvent
	var event21 *C.GdkEvent
	var x1 C.double
	var y1 C.double
	ret1 := C.gdk_events_get_center(this1, event21, &x1, &y1)
	var x2 float64
	var y2 float64
	var ret2 bool
	x2 = float64(x1)
	y2 = float64(y1)
	ret2 = ret1 != 0
	return x2, y2, ret2
}
func (this0 *Event) _GetDistance(event20 *Event) (float64, bool) {
	var this1 *C.GdkEvent
	var event21 *C.GdkEvent
	var distance1 C.double
	ret1 := C.gdk_events_get_distance(this1, event21, &distance1)
	var distance2 float64
	var ret2 bool
	distance2 = float64(distance1)
	ret2 = ret1 != 0
	return distance2, ret2
}
func (this0 *Event) Copy() *Event {
	var this1 *C.GdkEvent
	ret1 := C.gdk_event_copy(this1)
	var ret2 *Event
	ret2 = (*Event)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Event) Free() {
	var this1 *C.GdkEvent
	C.gdk_event_free(this1)
}
func (this0 *Event) GetAxis(axis_use0 AxisUse) (float64, bool) {
	var this1 *C.GdkEvent
	var axis_use1 C.GdkAxisUse
	var value1 C.double
	axis_use1 = C.GdkAxisUse(axis_use0)
	ret1 := C.gdk_event_get_axis(this1, axis_use1, &value1)
	var value2 float64
	var ret2 bool
	value2 = float64(value1)
	ret2 = ret1 != 0
	return value2, ret2
}
func (this0 *Event) GetButton() (int, bool) {
	var this1 *C.GdkEvent
	var button1 C.uint32_t
	ret1 := C.gdk_event_get_button(this1, &button1)
	var button2 int
	var ret2 bool
	button2 = int(button1)
	ret2 = ret1 != 0
	return button2, ret2
}
func (this0 *Event) GetClickCount() (int, bool) {
	var this1 *C.GdkEvent
	var click_count1 C.uint32_t
	ret1 := C.gdk_event_get_click_count(this1, &click_count1)
	var click_count2 int
	var ret2 bool
	click_count2 = int(click_count1)
	ret2 = ret1 != 0
	return click_count2, ret2
}
func (this0 *Event) GetCoords() (float64, float64, bool) {
	var this1 *C.GdkEvent
	var x_win1 C.double
	var y_win1 C.double
	ret1 := C.gdk_event_get_coords(this1, &x_win1, &y_win1)
	var x_win2 float64
	var y_win2 float64
	var ret2 bool
	x_win2 = float64(x_win1)
	y_win2 = float64(y_win1)
	ret2 = ret1 != 0
	return x_win2, y_win2, ret2
}
func (this0 *Event) GetDevice() *Device {
	var this1 *C.GdkEvent
	ret1 := C.gdk_event_get_device(this1)
	var ret2 *Device
	ret2 = (*Device)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Event) GetKeycode() (int, bool) {
	var this1 *C.GdkEvent
	var keycode1 C.uint16_t
	ret1 := C.gdk_event_get_keycode(this1, &keycode1)
	var keycode2 int
	var ret2 bool
	keycode2 = int(keycode1)
	ret2 = ret1 != 0
	return keycode2, ret2
}
func (this0 *Event) GetKeyval() (int, bool) {
	var this1 *C.GdkEvent
	var keyval1 C.uint32_t
	ret1 := C.gdk_event_get_keyval(this1, &keyval1)
	var keyval2 int
	var ret2 bool
	keyval2 = int(keyval1)
	ret2 = ret1 != 0
	return keyval2, ret2
}
func (this0 *Event) GetRootCoords() (float64, float64, bool) {
	var this1 *C.GdkEvent
	var x_root1 C.double
	var y_root1 C.double
	ret1 := C.gdk_event_get_root_coords(this1, &x_root1, &y_root1)
	var x_root2 float64
	var y_root2 float64
	var ret2 bool
	x_root2 = float64(x_root1)
	y_root2 = float64(y_root1)
	ret2 = ret1 != 0
	return x_root2, y_root2, ret2
}
func (this0 *Event) GetScreen() *Screen {
	var this1 *C.GdkEvent
	ret1 := C.gdk_event_get_screen(this1)
	var ret2 *Screen
	ret2 = (*Screen)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Event) GetScrollDirection() (ScrollDirection, bool) {
	var this1 *C.GdkEvent
	var direction1 C.GdkScrollDirection
	ret1 := C.gdk_event_get_scroll_direction(this1, &direction1)
	var direction2 ScrollDirection
	var ret2 bool
	direction2 = ScrollDirection(direction1)
	ret2 = ret1 != 0
	return direction2, ret2
}
func (this0 *Event) GetSourceDevice() *Device {
	var this1 *C.GdkEvent
	ret1 := C.gdk_event_get_source_device(this1)
	var ret2 *Device
	ret2 = (*Device)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Event) GetState() (ModifierType, bool) {
	var this1 *C.GdkEvent
	var state1 C.GdkModifierType
	ret1 := C.gdk_event_get_state(this1, &state1)
	var state2 ModifierType
	var ret2 bool
	state2 = ModifierType(state1)
	ret2 = ret1 != 0
	return state2, ret2
}
func (this0 *Event) GetTime() int {
	var this1 *C.GdkEvent
	ret1 := C.gdk_event_get_time(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Event) Put() {
	var this1 *C.GdkEvent
	C.gdk_event_put(this1)
}
func (this0 *Event) SetDevice(device0 DeviceLike) {
	var this1 *C.GdkEvent
	var device1 *C.GdkDevice
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	C.gdk_event_set_device(this1, device1)
}
func (this0 *Event) SetScreen(screen0 ScreenLike) {
	var this1 *C.GdkEvent
	var screen1 *C.GdkScreen
	if screen0 != nil {
		screen1 = screen0.InheritedFromGdkScreen()
	}
	C.gdk_event_set_screen(this1, screen1)
}
func (this0 *Event) SetSourceDevice(device0 DeviceLike) {
	var this1 *C.GdkEvent
	var device1 *C.GdkDevice
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	C.gdk_event_set_source_device(this1, device1)
}
func EventGet() *Event {
	ret1 := C.gdk_event_get()
	var ret2 *Event
	ret2 = (*Event)(unsafe.Pointer(ret1))
	return ret2
}
func EventHandlerSet(func0 EventFunc) {
	var func1 unsafe.Pointer
	if func0 != nil {
		func1 = unsafe.Pointer(&func0)}
	gobject.Holder.Grab(func1)
	C._gdk_event_handler_set(func1)
}
func EventPeek() *Event {
	ret1 := C.gdk_event_peek()
	var ret2 *Event
	ret2 = (*Event)(unsafe.Pointer(ret1))
	return ret2
}
func EventRequestMotions(event0 *EventMotion) {
	var event1 *C.GdkEventMotion
	event1 = (*C.GdkEventMotion)(unsafe.Pointer(event0))
	C.gdk_event_request_motions(event1)
}
type EventAny struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [7]byte
}
func (this0 *EventAny) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
type EventButton struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	Time uint32
	X float64
	Y float64
	Axes *float64
	State ModifierType
	Button uint32
	device0 *C.GdkDevice
	XRoot float64
	YRoot float64
}
func (this0 *EventButton) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventButton) Device() *Device {
	var device1 *Device
	device1 = (*Device)(gobject.ObjectWrap(unsafe.Pointer(this0.device0), true))
	return device1
}
type EventConfigure struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	X int32
	Y int32
	Width int32
	Height int32
	_ [4]byte
}
func (this0 *EventConfigure) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
type EventCrossing struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [7]byte
	subwindow0 *C.GdkWindow
	Time uint32
	_ [4]byte
	X float64
	Y float64
	XRoot float64
	YRoot float64
	Mode CrossingMode
	Detail NotifyType
	Focus int32
	State ModifierType
}
func (this0 *EventCrossing) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventCrossing) Subwindow() *Window {
	var subwindow1 *Window
	subwindow1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.subwindow0), true))
	return subwindow1
}
type EventDND struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [7]byte
	context0 *C.GdkDragContext
	Time uint32
	XRoot int16
	YRoot int16
}
func (this0 *EventDND) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventDND) Context() *DragContext {
	var context1 *DragContext
	context1 = (*DragContext)(gobject.ObjectWrap(unsafe.Pointer(this0.context0), true))
	return context1
}
type EventExpose struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	Area cairo.RectangleInt
	_ [4]byte
	Region *cairo.Region
	Count int32
	_ [4]byte
}
func (this0 *EventExpose) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
type EventFocus struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [1]byte
	In int16
	_ [4]byte
}
func (this0 *EventFocus) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
type EventFunc func(event *Event)
//export _GdkEventFunc_c_wrapper
func _GdkEventFunc_c_wrapper(event0 unsafe.Pointer, data0 unsafe.Pointer) {
	var event1 *Event
	var data1 EventFunc
	event1 = (*Event)(unsafe.Pointer((*C.GdkEvent)(event0)))
	data1 = *(*EventFunc)(data0)
	data1(event1)
}
//export _GdkEventFunc_c_wrapper_once
func _GdkEventFunc_c_wrapper_once(event0 unsafe.Pointer, data0 unsafe.Pointer) {
	_GdkEventFunc_c_wrapper(event0, data0)
	gobject.Holder.Release(data0)
}
type EventGrabBroken struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	Keyboard int32
	Implicit int32
	_ [4]byte
	grab_window0 *C.GdkWindow
}
func (this0 *EventGrabBroken) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventGrabBroken) GrabWindow() *Window {
	var grab_window1 *Window
	grab_window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.grab_window0), true))
	return grab_window1
}
type EventKey struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	Time uint32
	State ModifierType
	Keyval uint32
	Length int32
	_ [4]byte
	string0 *C.char
	HardwareKeycode uint16
	Group uint8
	_ [1]byte
	IsModifier uint32
}
func (this0 *EventKey) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventKey) String() string {
	var string1 string
	string1 = C.GoString(this0.string0)
	return string1
}
type EventMask C.uint32_t
const (
	EventMaskExposureMask EventMask = 2
	EventMaskPointerMotionMask EventMask = 4
	EventMaskPointerMotionHintMask EventMask = 8
	EventMaskButtonMotionMask EventMask = 16
	EventMaskButton1MotionMask EventMask = 32
	EventMaskButton2MotionMask EventMask = 64
	EventMaskButton3MotionMask EventMask = 128
	EventMaskButtonPressMask EventMask = 256
	EventMaskButtonReleaseMask EventMask = 512
	EventMaskKeyPressMask EventMask = 1024
	EventMaskKeyReleaseMask EventMask = 2048
	EventMaskEnterNotifyMask EventMask = 4096
	EventMaskLeaveNotifyMask EventMask = 8192
	EventMaskFocusChangeMask EventMask = 16384
	EventMaskStructureMask EventMask = 32768
	EventMaskPropertyChangeMask EventMask = 65536
	EventMaskVisibilityNotifyMask EventMask = 131072
	EventMaskProximityInMask EventMask = 262144
	EventMaskProximityOutMask EventMask = 524288
	EventMaskSubstructureMask EventMask = 1048576
	EventMaskScrollMask EventMask = 2097152
	EventMaskAllEventsMask EventMask = 4194302
)
type EventMotion struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	Time uint32
	X float64
	Y float64
	Axes *float64
	State ModifierType
	IsHint int16
	_ [2]byte
	device0 *C.GdkDevice
	XRoot float64
	YRoot float64
}
func (this0 *EventMotion) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventMotion) Device() *Device {
	var device1 *Device
	device1 = (*Device)(gobject.ObjectWrap(unsafe.Pointer(this0.device0), true))
	return device1
}
type EventOwnerChange struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [7]byte
	owner0 *C.GdkWindow
	Reason OwnerChange
	_ [4]byte
	Selection Atom
	Time uint32
	SelectionTime uint32
}
func (this0 *EventOwnerChange) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventOwnerChange) Owner() *Window {
	var owner1 *Window
	owner1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.owner0), true))
	return owner1
}
type EventProperty struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [7]byte
	Atom Atom
	Time uint32
	State uint32
}
func (this0 *EventProperty) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
type EventProximity struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	Time uint32
	device0 *C.GdkDevice
}
func (this0 *EventProximity) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventProximity) Device() *Device {
	var device1 *Device
	device1 = (*Device)(gobject.ObjectWrap(unsafe.Pointer(this0.device0), true))
	return device1
}
type EventScroll struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	Time uint32
	X float64
	Y float64
	State ModifierType
	Direction ScrollDirection
	device0 *C.GdkDevice
	XRoot float64
	YRoot float64
}
func (this0 *EventScroll) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventScroll) Device() *Device {
	var device1 *Device
	device1 = (*Device)(gobject.ObjectWrap(unsafe.Pointer(this0.device0), true))
	return device1
}
type EventSelection struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [7]byte
	Selection Atom
	Target Atom
	Property Atom
	Time uint32
	_ [4]byte
	requestor0 *C.GdkWindow
}
func (this0 *EventSelection) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventSelection) Requestor() *Window {
	var requestor1 *Window
	requestor1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.requestor0), true))
	return requestor1
}
type EventSetting struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	Action SettingAction
	name0 *C.char
}
func (this0 *EventSetting) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
func (this0 *EventSetting) Name() string {
	var name1 string
	name1 = C.GoString(this0.name0)
	return name1
}
type EventType C.int32_t
const (
	EventTypeNothing EventType = -1
	EventTypeDelete EventType = 0
	EventTypeDestroy EventType = 1
	EventTypeExpose EventType = 2
	EventTypeMotionNotify EventType = 3
	EventTypeButtonPress EventType = 4
	EventType2buttonPress EventType = 5
	EventType3buttonPress EventType = 6
	EventTypeButtonRelease EventType = 7
	EventTypeKeyPress EventType = 8
	EventTypeKeyRelease EventType = 9
	EventTypeEnterNotify EventType = 10
	EventTypeLeaveNotify EventType = 11
	EventTypeFocusChange EventType = 12
	EventTypeConfigure EventType = 13
	EventTypeMap EventType = 14
	EventTypeUnmap EventType = 15
	EventTypePropertyNotify EventType = 16
	EventTypeSelectionClear EventType = 17
	EventTypeSelectionRequest EventType = 18
	EventTypeSelectionNotify EventType = 19
	EventTypeProximityIn EventType = 20
	EventTypeProximityOut EventType = 21
	EventTypeDragEnter EventType = 22
	EventTypeDragLeave EventType = 23
	EventTypeDragMotion EventType = 24
	EventTypeDragStatus EventType = 25
	EventTypeDropStart EventType = 26
	EventTypeDropFinished EventType = 27
	EventTypeClientEvent EventType = 28
	EventTypeVisibilityNotify EventType = 29
	EventTypeScroll EventType = 31
	EventTypeWindowState EventType = 32
	EventTypeSetting EventType = 33
	EventTypeOwnerChange EventType = 34
	EventTypeGrabBroken EventType = 35
	EventTypeDamage EventType = 36
	EventTypeEventLast EventType = 37
)
type EventVisibility struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	State VisibilityState
}
func (this0 *EventVisibility) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
type EventWindowState struct {
	Type EventType
	_ [4]byte
	window0 *C.GdkWindow
	SendEvent int8
	_ [3]byte
	ChangedMask WindowState
	NewWindowState WindowState
	_ [4]byte
}
func (this0 *EventWindowState) Window() *Window {
	var window1 *Window
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(this0.window0), true))
	return window1
}
type ExtensionMode C.uint32_t
const (
	ExtensionModeNone ExtensionMode = 0
	ExtensionModeAll ExtensionMode = 1
	ExtensionModeCursor ExtensionMode = 2
)
type FilterFunc func(event *Event, data unsafe.Pointer) FilterReturn
//export _GdkFilterFunc_c_wrapper
func _GdkFilterFunc_c_wrapper(xevent0 unsafe.Pointer, event0 unsafe.Pointer, data0 unsafe.Pointer) uint32 {
	var xevent1 FilterFunc
	var event1 *Event
	var data1 unsafe.Pointer
	xevent1 = *(*FilterFunc)(xevent0)
	event1 = (*Event)(unsafe.Pointer((*C.GdkEvent)(event0)))
	data1 = (unsafe.Pointer)(data0)
	ret1 := xevent1(event1, data1)
	var ret2 C.GdkFilterReturn
	ret2 = C.GdkFilterReturn(ret1)
	return (uint32)(ret2)
}
//export _GdkFilterFunc_c_wrapper_once
func _GdkFilterFunc_c_wrapper_once(xevent0 unsafe.Pointer, event0 unsafe.Pointer, data0 unsafe.Pointer) uint32 {
	ret := _GdkFilterFunc_c_wrapper(xevent0, event0, data0)
	gobject.Holder.Release(xevent0)
	return ret
}
type FilterReturn C.uint32_t
const (
	FilterReturnContinue FilterReturn = 0
	FilterReturnTranslate FilterReturn = 1
	FilterReturnRemove FilterReturn = 2
)
type Geometry struct {
	MinWidth int32
	MinHeight int32
	MaxWidth int32
	MaxHeight int32
	BaseWidth int32
	BaseHeight int32
	WidthInc int32
	HeightInc int32
	MinAspect float64
	MaxAspect float64
	WinGravity Gravity
	_ [4]byte
}
type GrabOwnership C.uint32_t
const (
	GrabOwnershipNone GrabOwnership = 0
	GrabOwnershipWindow GrabOwnership = 1
	GrabOwnershipApplication GrabOwnership = 2
)
type GrabStatus C.uint32_t
const (
	GrabStatusSuccess GrabStatus = 0
	GrabStatusAlreadyGrabbed GrabStatus = 1
	GrabStatusInvalidTime GrabStatus = 2
	GrabStatusNotViewable GrabStatus = 3
	GrabStatusFrozen GrabStatus = 4
)
type Gravity C.uint32_t
const (
	GravityNorthWest Gravity = 1
	GravityNorth Gravity = 2
	GravityNorthEast Gravity = 3
	GravityWest Gravity = 4
	GravityCenter Gravity = 5
	GravityEast Gravity = 6
	GravitySouthWest Gravity = 7
	GravitySouth Gravity = 8
	GravitySouthEast Gravity = 9
	GravityStatic Gravity = 10
)
type InputMode C.uint32_t
const (
	InputModeDisabled InputMode = 0
	InputModeScreen InputMode = 1
	InputModeWindow InputMode = 2
)
type InputSource C.uint32_t
const (
	InputSourceMouse InputSource = 0
	InputSourcePen InputSource = 1
	InputSourceEraser InputSource = 2
	InputSourceCursor InputSource = 3
	InputSourceKeyboard InputSource = 4
)
const Key_0 = 48
const Key_1 = 49
const Key_2 = 50
const Key_3 = 51
const Key_3270_AltCursor = 64784
const Key_3270_Attn = 64782
const Key_3270_BackTab = 64773
const Key_3270_ChangeScreen = 64793
const Key_3270_Copy = 64789
const Key_3270_CursorBlink = 64783
const Key_3270_CursorSelect = 64796
const Key_3270_DeleteWord = 64794
const Key_3270_Duplicate = 64769
const Key_3270_Enter = 64798
const Key_3270_EraseEOF = 64774
const Key_3270_EraseInput = 64775
const Key_3270_ExSelect = 64795
const Key_3270_FieldMark = 64770
const Key_3270_Ident = 64787
const Key_3270_Jump = 64786
const Key_3270_KeyClick = 64785
const Key_3270_Left2 = 64772
const Key_3270_PA1 = 64778
const Key_3270_PA2 = 64779
const Key_3270_PA3 = 64780
const Key_3270_Play = 64790
const Key_3270_PrintScreen = 64797
const Key_3270_Quit = 64777
const Key_3270_Record = 64792
const Key_3270_Reset = 64776
const Key_3270_Right2 = 64771
const Key_3270_Rule = 64788
const Key_3270_Setup = 64791
const Key_3270_Test = 64781
const Key_4 = 52
const Key_5 = 53
const Key_6 = 54
const Key_7 = 55
const Key_8 = 56
const Key_9 = 57
const Key_A = 65
const Key_AE = 198
const Key_Aacute = 193
const Key_Abelowdot = 16785056
const Key_Abreve = 451
const Key_Abreveacute = 16785070
const Key_Abrevebelowdot = 16785078
const Key_Abrevegrave = 16785072
const Key_Abrevehook = 16785074
const Key_Abrevetilde = 16785076
const Key_AccessX_Enable = 65136
const Key_AccessX_Feedback_Enable = 65137
const Key_Acircumflex = 194
const Key_Acircumflexacute = 16785060
const Key_Acircumflexbelowdot = 16785068
const Key_Acircumflexgrave = 16785062
const Key_Acircumflexhook = 16785064
const Key_Acircumflextilde = 16785066
const Key_AddFavorite = 269025081
const Key_Adiaeresis = 196
const Key_Agrave = 192
const Key_Ahook = 16785058
const Key_Alt_L = 65513
const Key_Alt_R = 65514
const Key_Amacron = 960
const Key_Aogonek = 417
const Key_ApplicationLeft = 269025104
const Key_ApplicationRight = 269025105
const Key_Arabic_0 = 16778848
const Key_Arabic_1 = 16778849
const Key_Arabic_2 = 16778850
const Key_Arabic_3 = 16778851
const Key_Arabic_4 = 16778852
const Key_Arabic_5 = 16778853
const Key_Arabic_6 = 16778854
const Key_Arabic_7 = 16778855
const Key_Arabic_8 = 16778856
const Key_Arabic_9 = 16778857
const Key_Arabic_ain = 1497
const Key_Arabic_alef = 1479
const Key_Arabic_alefmaksura = 1513
const Key_Arabic_beh = 1480
const Key_Arabic_comma = 1452
const Key_Arabic_dad = 1494
const Key_Arabic_dal = 1487
const Key_Arabic_damma = 1519
const Key_Arabic_dammatan = 1516
const Key_Arabic_ddal = 16778888
const Key_Arabic_farsi_yeh = 16778956
const Key_Arabic_fatha = 1518
const Key_Arabic_fathatan = 1515
const Key_Arabic_feh = 1505
const Key_Arabic_fullstop = 16778964
const Key_Arabic_gaf = 16778927
const Key_Arabic_ghain = 1498
const Key_Arabic_ha = 1511
const Key_Arabic_hah = 1485
const Key_Arabic_hamza = 1473
const Key_Arabic_hamza_above = 16778836
const Key_Arabic_hamza_below = 16778837
const Key_Arabic_hamzaonalef = 1475
const Key_Arabic_hamzaonwaw = 1476
const Key_Arabic_hamzaonyeh = 1478
const Key_Arabic_hamzaunderalef = 1477
const Key_Arabic_heh = 1511
const Key_Arabic_heh_doachashmee = 16778942
const Key_Arabic_heh_goal = 16778945
const Key_Arabic_jeem = 1484
const Key_Arabic_jeh = 16778904
const Key_Arabic_kaf = 1507
const Key_Arabic_kasra = 1520
const Key_Arabic_kasratan = 1517
const Key_Arabic_keheh = 16778921
const Key_Arabic_khah = 1486
const Key_Arabic_lam = 1508
const Key_Arabic_madda_above = 16778835
const Key_Arabic_maddaonalef = 1474
const Key_Arabic_meem = 1509
const Key_Arabic_noon = 1510
const Key_Arabic_noon_ghunna = 16778938
const Key_Arabic_peh = 16778878
const Key_Arabic_percent = 16778858
const Key_Arabic_qaf = 1506
const Key_Arabic_question_mark = 1471
const Key_Arabic_ra = 1489
const Key_Arabic_rreh = 16778897
const Key_Arabic_sad = 1493
const Key_Arabic_seen = 1491
const Key_Arabic_semicolon = 1467
const Key_Arabic_shadda = 1521
const Key_Arabic_sheen = 1492
const Key_Arabic_sukun = 1522
const Key_Arabic_superscript_alef = 16778864
const Key_Arabic_switch = 65406
const Key_Arabic_tah = 1495
const Key_Arabic_tatweel = 1504
const Key_Arabic_tcheh = 16778886
const Key_Arabic_teh = 1482
const Key_Arabic_tehmarbuta = 1481
const Key_Arabic_thal = 1488
const Key_Arabic_theh = 1483
const Key_Arabic_tteh = 16778873
const Key_Arabic_veh = 16778916
const Key_Arabic_waw = 1512
const Key_Arabic_yeh = 1514
const Key_Arabic_yeh_baree = 16778962
const Key_Arabic_zah = 1496
const Key_Arabic_zain = 1490
const Key_Aring = 197
const Key_Armenian_AT = 16778552
const Key_Armenian_AYB = 16778545
const Key_Armenian_BEN = 16778546
const Key_Armenian_CHA = 16778569
const Key_Armenian_DA = 16778548
const Key_Armenian_DZA = 16778561
const Key_Armenian_E = 16778551
const Key_Armenian_FE = 16778582
const Key_Armenian_GHAT = 16778562
const Key_Armenian_GIM = 16778547
const Key_Armenian_HI = 16778565
const Key_Armenian_HO = 16778560
const Key_Armenian_INI = 16778555
const Key_Armenian_JE = 16778571
const Key_Armenian_KE = 16778580
const Key_Armenian_KEN = 16778559
const Key_Armenian_KHE = 16778557
const Key_Armenian_LYUN = 16778556
const Key_Armenian_MEN = 16778564
const Key_Armenian_NU = 16778566
const Key_Armenian_O = 16778581
const Key_Armenian_PE = 16778570
const Key_Armenian_PYUR = 16778579
const Key_Armenian_RA = 16778572
const Key_Armenian_RE = 16778576
const Key_Armenian_SE = 16778573
const Key_Armenian_SHA = 16778567
const Key_Armenian_TCHE = 16778563
const Key_Armenian_TO = 16778553
const Key_Armenian_TSA = 16778558
const Key_Armenian_TSO = 16778577
const Key_Armenian_TYUN = 16778575
const Key_Armenian_VEV = 16778574
const Key_Armenian_VO = 16778568
const Key_Armenian_VYUN = 16778578
const Key_Armenian_YECH = 16778549
const Key_Armenian_ZA = 16778550
const Key_Armenian_ZHE = 16778554
const Key_Armenian_accent = 16778587
const Key_Armenian_amanak = 16778588
const Key_Armenian_apostrophe = 16778586
const Key_Armenian_at = 16778600
const Key_Armenian_ayb = 16778593
const Key_Armenian_ben = 16778594
const Key_Armenian_but = 16778589
const Key_Armenian_cha = 16778617
const Key_Armenian_da = 16778596
const Key_Armenian_dza = 16778609
const Key_Armenian_e = 16778599
const Key_Armenian_exclam = 16778588
const Key_Armenian_fe = 16778630
const Key_Armenian_full_stop = 16778633
const Key_Armenian_ghat = 16778610
const Key_Armenian_gim = 16778595
const Key_Armenian_hi = 16778613
const Key_Armenian_ho = 16778608
const Key_Armenian_hyphen = 16778634
const Key_Armenian_ini = 16778603
const Key_Armenian_je = 16778619
const Key_Armenian_ke = 16778628
const Key_Armenian_ken = 16778607
const Key_Armenian_khe = 16778605
const Key_Armenian_ligature_ew = 16778631
const Key_Armenian_lyun = 16778604
const Key_Armenian_men = 16778612
const Key_Armenian_nu = 16778614
const Key_Armenian_o = 16778629
const Key_Armenian_paruyk = 16778590
const Key_Armenian_pe = 16778618
const Key_Armenian_pyur = 16778627
const Key_Armenian_question = 16778590
const Key_Armenian_ra = 16778620
const Key_Armenian_re = 16778624
const Key_Armenian_se = 16778621
const Key_Armenian_separation_mark = 16778589
const Key_Armenian_sha = 16778615
const Key_Armenian_shesht = 16778587
const Key_Armenian_tche = 16778611
const Key_Armenian_to = 16778601
const Key_Armenian_tsa = 16778606
const Key_Armenian_tso = 16778625
const Key_Armenian_tyun = 16778623
const Key_Armenian_verjaket = 16778633
const Key_Armenian_vev = 16778622
const Key_Armenian_vo = 16778616
const Key_Armenian_vyun = 16778626
const Key_Armenian_yech = 16778597
const Key_Armenian_yentamna = 16778634
const Key_Armenian_za = 16778598
const Key_Armenian_zhe = 16778602
const Key_Atilde = 195
const Key_AudibleBell_Enable = 65146
const Key_AudioCycleTrack = 269025179
const Key_AudioForward = 269025175
const Key_AudioLowerVolume = 269025041
const Key_AudioMedia = 269025074
const Key_AudioMute = 269025042
const Key_AudioNext = 269025047
const Key_AudioPause = 269025073
const Key_AudioPlay = 269025044
const Key_AudioPrev = 269025046
const Key_AudioRaiseVolume = 269025043
const Key_AudioRandomPlay = 269025177
const Key_AudioRecord = 269025052
const Key_AudioRepeat = 269025176
const Key_AudioRewind = 269025086
const Key_AudioStop = 269025045
const Key_Away = 269025165
const Key_B = 66
const Key_Babovedot = 16784898
const Key_Back = 269025062
const Key_BackForward = 269025087
const Key_BackSpace = 65288
const Key_Battery = 269025171
const Key_Begin = 65368
const Key_Blue = 269025190
const Key_Bluetooth = 269025172
const Key_Book = 269025106
const Key_BounceKeys_Enable = 65140
const Key_Break = 65387
const Key_BrightnessAdjust = 269025083
const Key_Byelorussian_SHORTU = 1726
const Key_Byelorussian_shortu = 1710
const Key_C = 67
const Key_CD = 269025107
const Key_Cabovedot = 709
const Key_Cacute = 454
const Key_Calculator = 269025053
const Key_Calendar = 269025056
const Key_Cancel = 65385
const Key_Caps_Lock = 65509
const Key_Ccaron = 456
const Key_Ccedilla = 199
const Key_Ccircumflex = 710
const Key_Clear = 65291
const Key_ClearGrab = 269024801
const Key_Close = 269025110
const Key_Codeinput = 65335
const Key_ColonSign = 16785569
const Key_Community = 269025085
const Key_ContrastAdjust = 269025058
const Key_Control_L = 65507
const Key_Control_R = 65508
const Key_Copy = 269025111
const Key_CruzeiroSign = 16785570
const Key_Cut = 269025112
const Key_CycleAngle = 269025180
const Key_Cyrillic_A = 1761
const Key_Cyrillic_BE = 1762
const Key_Cyrillic_CHE = 1790
const Key_Cyrillic_CHE_descender = 16778422
const Key_Cyrillic_CHE_vertstroke = 16778424
const Key_Cyrillic_DE = 1764
const Key_Cyrillic_DZHE = 1727
const Key_Cyrillic_E = 1788
const Key_Cyrillic_EF = 1766
const Key_Cyrillic_EL = 1772
const Key_Cyrillic_EM = 1773
const Key_Cyrillic_EN = 1774
const Key_Cyrillic_EN_descender = 16778402
const Key_Cyrillic_ER = 1778
const Key_Cyrillic_ES = 1779
const Key_Cyrillic_GHE = 1767
const Key_Cyrillic_GHE_bar = 16778386
const Key_Cyrillic_HA = 1768
const Key_Cyrillic_HARDSIGN = 1791
const Key_Cyrillic_HA_descender = 16778418
const Key_Cyrillic_I = 1769
const Key_Cyrillic_IE = 1765
const Key_Cyrillic_IO = 1715
const Key_Cyrillic_I_macron = 16778466
const Key_Cyrillic_JE = 1720
const Key_Cyrillic_KA = 1771
const Key_Cyrillic_KA_descender = 16778394
const Key_Cyrillic_KA_vertstroke = 16778396
const Key_Cyrillic_LJE = 1721
const Key_Cyrillic_NJE = 1722
const Key_Cyrillic_O = 1775
const Key_Cyrillic_O_bar = 16778472
const Key_Cyrillic_PE = 1776
const Key_Cyrillic_SCHWA = 16778456
const Key_Cyrillic_SHA = 1787
const Key_Cyrillic_SHCHA = 1789
const Key_Cyrillic_SHHA = 16778426
const Key_Cyrillic_SHORTI = 1770
const Key_Cyrillic_SOFTSIGN = 1784
const Key_Cyrillic_TE = 1780
const Key_Cyrillic_TSE = 1763
const Key_Cyrillic_U = 1781
const Key_Cyrillic_U_macron = 16778478
const Key_Cyrillic_U_straight = 16778414
const Key_Cyrillic_U_straight_bar = 16778416
const Key_Cyrillic_VE = 1783
const Key_Cyrillic_YA = 1777
const Key_Cyrillic_YERU = 1785
const Key_Cyrillic_YU = 1760
const Key_Cyrillic_ZE = 1786
const Key_Cyrillic_ZHE = 1782
const Key_Cyrillic_ZHE_descender = 16778390
const Key_Cyrillic_a = 1729
const Key_Cyrillic_be = 1730
const Key_Cyrillic_che = 1758
const Key_Cyrillic_che_descender = 16778423
const Key_Cyrillic_che_vertstroke = 16778425
const Key_Cyrillic_de = 1732
const Key_Cyrillic_dzhe = 1711
const Key_Cyrillic_e = 1756
const Key_Cyrillic_ef = 1734
const Key_Cyrillic_el = 1740
const Key_Cyrillic_em = 1741
const Key_Cyrillic_en = 1742
const Key_Cyrillic_en_descender = 16778403
const Key_Cyrillic_er = 1746
const Key_Cyrillic_es = 1747
const Key_Cyrillic_ghe = 1735
const Key_Cyrillic_ghe_bar = 16778387
const Key_Cyrillic_ha = 1736
const Key_Cyrillic_ha_descender = 16778419
const Key_Cyrillic_hardsign = 1759
const Key_Cyrillic_i = 1737
const Key_Cyrillic_i_macron = 16778467
const Key_Cyrillic_ie = 1733
const Key_Cyrillic_io = 1699
const Key_Cyrillic_je = 1704
const Key_Cyrillic_ka = 1739
const Key_Cyrillic_ka_descender = 16778395
const Key_Cyrillic_ka_vertstroke = 16778397
const Key_Cyrillic_lje = 1705
const Key_Cyrillic_nje = 1706
const Key_Cyrillic_o = 1743
const Key_Cyrillic_o_bar = 16778473
const Key_Cyrillic_pe = 1744
const Key_Cyrillic_schwa = 16778457
const Key_Cyrillic_sha = 1755
const Key_Cyrillic_shcha = 1757
const Key_Cyrillic_shha = 16778427
const Key_Cyrillic_shorti = 1738
const Key_Cyrillic_softsign = 1752
const Key_Cyrillic_te = 1748
const Key_Cyrillic_tse = 1731
const Key_Cyrillic_u = 1749
const Key_Cyrillic_u_macron = 16778479
const Key_Cyrillic_u_straight = 16778415
const Key_Cyrillic_u_straight_bar = 16778417
const Key_Cyrillic_ve = 1751
const Key_Cyrillic_ya = 1745
const Key_Cyrillic_yeru = 1753
const Key_Cyrillic_yu = 1728
const Key_Cyrillic_ze = 1754
const Key_Cyrillic_zhe = 1750
const Key_Cyrillic_zhe_descender = 16778391
const Key_D = 68
const Key_DOS = 269025114
const Key_Dabovedot = 16784906
const Key_Dcaron = 463
const Key_Delete = 65535
const Key_Display = 269025113
const Key_Documents = 269025115
const Key_DongSign = 16785579
const Key_Down = 65364
const Key_Dstroke = 464
const Key_E = 69
const Key_ENG = 957
const Key_ETH = 208
const Key_Eabovedot = 972
const Key_Eacute = 201
const Key_Ebelowdot = 16785080
const Key_Ecaron = 460
const Key_Ecircumflex = 202
const Key_Ecircumflexacute = 16785086
const Key_Ecircumflexbelowdot = 16785094
const Key_Ecircumflexgrave = 16785088
const Key_Ecircumflexhook = 16785090
const Key_Ecircumflextilde = 16785092
const Key_EcuSign = 16785568
const Key_Ediaeresis = 203
const Key_Egrave = 200
const Key_Ehook = 16785082
const Key_Eisu_Shift = 65327
const Key_Eisu_toggle = 65328
const Key_Eject = 269025068
const Key_Emacron = 938
const Key_End = 65367
const Key_Eogonek = 458
const Key_Escape = 65307
const Key_Eth = 208
const Key_Etilde = 16785084
const Key_EuroSign = 8364
const Key_Excel = 269025116
const Key_Execute = 65378
const Key_Explorer = 269025117
const Key_F = 70
const Key_F1 = 65470
const Key_F10 = 65479
const Key_F11 = 65480
const Key_F12 = 65481
const Key_F13 = 65482
const Key_F14 = 65483
const Key_F15 = 65484
const Key_F16 = 65485
const Key_F17 = 65486
const Key_F18 = 65487
const Key_F19 = 65488
const Key_F2 = 65471
const Key_F20 = 65489
const Key_F21 = 65490
const Key_F22 = 65491
const Key_F23 = 65492
const Key_F24 = 65493
const Key_F25 = 65494
const Key_F26 = 65495
const Key_F27 = 65496
const Key_F28 = 65497
const Key_F29 = 65498
const Key_F3 = 65472
const Key_F30 = 65499
const Key_F31 = 65500
const Key_F32 = 65501
const Key_F33 = 65502
const Key_F34 = 65503
const Key_F35 = 65504
const Key_F4 = 65473
const Key_F5 = 65474
const Key_F6 = 65475
const Key_F7 = 65476
const Key_F8 = 65477
const Key_F9 = 65478
const Key_FFrancSign = 16785571
const Key_Fabovedot = 16784926
const Key_Farsi_0 = 16778992
const Key_Farsi_1 = 16778993
const Key_Farsi_2 = 16778994
const Key_Farsi_3 = 16778995
const Key_Farsi_4 = 16778996
const Key_Farsi_5 = 16778997
const Key_Farsi_6 = 16778998
const Key_Farsi_7 = 16778999
const Key_Farsi_8 = 16779000
const Key_Farsi_9 = 16779001
const Key_Farsi_yeh = 16778956
const Key_Favorites = 269025072
const Key_Finance = 269025084
const Key_Find = 65384
const Key_First_Virtual_Screen = 65232
const Key_Forward = 269025063
const Key_FrameBack = 269025181
const Key_FrameForward = 269025182
const Key_G = 71
const Key_Gabovedot = 725
const Key_Game = 269025118
const Key_Gbreve = 683
const Key_Gcaron = 16777702
const Key_Gcedilla = 939
const Key_Gcircumflex = 728
const Key_Georgian_an = 16781520
const Key_Georgian_ban = 16781521
const Key_Georgian_can = 16781546
const Key_Georgian_char = 16781549
const Key_Georgian_chin = 16781545
const Key_Georgian_cil = 16781548
const Key_Georgian_don = 16781523
const Key_Georgian_en = 16781524
const Key_Georgian_fi = 16781558
const Key_Georgian_gan = 16781522
const Key_Georgian_ghan = 16781542
const Key_Georgian_hae = 16781552
const Key_Georgian_har = 16781556
const Key_Georgian_he = 16781553
const Key_Georgian_hie = 16781554
const Key_Georgian_hoe = 16781557
const Key_Georgian_in = 16781528
const Key_Georgian_jhan = 16781551
const Key_Georgian_jil = 16781547
const Key_Georgian_kan = 16781529
const Key_Georgian_khar = 16781541
const Key_Georgian_las = 16781530
const Key_Georgian_man = 16781531
const Key_Georgian_nar = 16781532
const Key_Georgian_on = 16781533
const Key_Georgian_par = 16781534
const Key_Georgian_phar = 16781540
const Key_Georgian_qar = 16781543
const Key_Georgian_rae = 16781536
const Key_Georgian_san = 16781537
const Key_Georgian_shin = 16781544
const Key_Georgian_tan = 16781527
const Key_Georgian_tar = 16781538
const Key_Georgian_un = 16781539
const Key_Georgian_vin = 16781525
const Key_Georgian_we = 16781555
const Key_Georgian_xan = 16781550
const Key_Georgian_zen = 16781526
const Key_Georgian_zhar = 16781535
const Key_Go = 269025119
const Key_Greek_ALPHA = 1985
const Key_Greek_ALPHAaccent = 1953
const Key_Greek_BETA = 1986
const Key_Greek_CHI = 2007
const Key_Greek_DELTA = 1988
const Key_Greek_EPSILON = 1989
const Key_Greek_EPSILONaccent = 1954
const Key_Greek_ETA = 1991
const Key_Greek_ETAaccent = 1955
const Key_Greek_GAMMA = 1987
const Key_Greek_IOTA = 1993
const Key_Greek_IOTAaccent = 1956
const Key_Greek_IOTAdiaeresis = 1957
const Key_Greek_IOTAdieresis = 1957
const Key_Greek_KAPPA = 1994
const Key_Greek_LAMBDA = 1995
const Key_Greek_LAMDA = 1995
const Key_Greek_MU = 1996
const Key_Greek_NU = 1997
const Key_Greek_OMEGA = 2009
const Key_Greek_OMEGAaccent = 1963
const Key_Greek_OMICRON = 1999
const Key_Greek_OMICRONaccent = 1959
const Key_Greek_PHI = 2006
const Key_Greek_PI = 2000
const Key_Greek_PSI = 2008
const Key_Greek_RHO = 2001
const Key_Greek_SIGMA = 2002
const Key_Greek_TAU = 2004
const Key_Greek_THETA = 1992
const Key_Greek_UPSILON = 2005
const Key_Greek_UPSILONaccent = 1960
const Key_Greek_UPSILONdieresis = 1961
const Key_Greek_XI = 1998
const Key_Greek_ZETA = 1990
const Key_Greek_accentdieresis = 1966
const Key_Greek_alpha = 2017
const Key_Greek_alphaaccent = 1969
const Key_Greek_beta = 2018
const Key_Greek_chi = 2039
const Key_Greek_delta = 2020
const Key_Greek_epsilon = 2021
const Key_Greek_epsilonaccent = 1970
const Key_Greek_eta = 2023
const Key_Greek_etaaccent = 1971
const Key_Greek_finalsmallsigma = 2035
const Key_Greek_gamma = 2019
const Key_Greek_horizbar = 1967
const Key_Greek_iota = 2025
const Key_Greek_iotaaccent = 1972
const Key_Greek_iotaaccentdieresis = 1974
const Key_Greek_iotadieresis = 1973
const Key_Greek_kappa = 2026
const Key_Greek_lambda = 2027
const Key_Greek_lamda = 2027
const Key_Greek_mu = 2028
const Key_Greek_nu = 2029
const Key_Greek_omega = 2041
const Key_Greek_omegaaccent = 1979
const Key_Greek_omicron = 2031
const Key_Greek_omicronaccent = 1975
const Key_Greek_phi = 2038
const Key_Greek_pi = 2032
const Key_Greek_psi = 2040
const Key_Greek_rho = 2033
const Key_Greek_sigma = 2034
const Key_Greek_switch = 65406
const Key_Greek_tau = 2036
const Key_Greek_theta = 2024
const Key_Greek_upsilon = 2037
const Key_Greek_upsilonaccent = 1976
const Key_Greek_upsilonaccentdieresis = 1978
const Key_Greek_upsilondieresis = 1977
const Key_Greek_xi = 2030
const Key_Greek_zeta = 2022
const Key_Green = 269025188
const Key_H = 72
const Key_Hangul = 65329
const Key_Hangul_A = 3775
const Key_Hangul_AE = 3776
const Key_Hangul_AraeA = 3830
const Key_Hangul_AraeAE = 3831
const Key_Hangul_Banja = 65337
const Key_Hangul_Cieuc = 3770
const Key_Hangul_Codeinput = 65335
const Key_Hangul_Dikeud = 3751
const Key_Hangul_E = 3780
const Key_Hangul_EO = 3779
const Key_Hangul_EU = 3793
const Key_Hangul_End = 65331
const Key_Hangul_Hanja = 65332
const Key_Hangul_Hieuh = 3774
const Key_Hangul_I = 3795
const Key_Hangul_Ieung = 3767
const Key_Hangul_J_Cieuc = 3818
const Key_Hangul_J_Dikeud = 3802
const Key_Hangul_J_Hieuh = 3822
const Key_Hangul_J_Ieung = 3816
const Key_Hangul_J_Jieuj = 3817
const Key_Hangul_J_Khieuq = 3819
const Key_Hangul_J_Kiyeog = 3796
const Key_Hangul_J_KiyeogSios = 3798
const Key_Hangul_J_KkogjiDalrinIeung = 3833
const Key_Hangul_J_Mieum = 3811
const Key_Hangul_J_Nieun = 3799
const Key_Hangul_J_NieunHieuh = 3801
const Key_Hangul_J_NieunJieuj = 3800
const Key_Hangul_J_PanSios = 3832
const Key_Hangul_J_Phieuf = 3821
const Key_Hangul_J_Pieub = 3812
const Key_Hangul_J_PieubSios = 3813
const Key_Hangul_J_Rieul = 3803
const Key_Hangul_J_RieulHieuh = 3810
const Key_Hangul_J_RieulKiyeog = 3804
const Key_Hangul_J_RieulMieum = 3805
const Key_Hangul_J_RieulPhieuf = 3809
const Key_Hangul_J_RieulPieub = 3806
const Key_Hangul_J_RieulSios = 3807
const Key_Hangul_J_RieulTieut = 3808
const Key_Hangul_J_Sios = 3814
const Key_Hangul_J_SsangKiyeog = 3797
const Key_Hangul_J_SsangSios = 3815
const Key_Hangul_J_Tieut = 3820
const Key_Hangul_J_YeorinHieuh = 3834
const Key_Hangul_Jamo = 65333
const Key_Hangul_Jeonja = 65336
const Key_Hangul_Jieuj = 3768
const Key_Hangul_Khieuq = 3771
const Key_Hangul_Kiyeog = 3745
const Key_Hangul_KiyeogSios = 3747
const Key_Hangul_KkogjiDalrinIeung = 3827
const Key_Hangul_Mieum = 3761
const Key_Hangul_MultipleCandidate = 65341
const Key_Hangul_Nieun = 3748
const Key_Hangul_NieunHieuh = 3750
const Key_Hangul_NieunJieuj = 3749
const Key_Hangul_O = 3783
const Key_Hangul_OE = 3786
const Key_Hangul_PanSios = 3826
const Key_Hangul_Phieuf = 3773
const Key_Hangul_Pieub = 3762
const Key_Hangul_PieubSios = 3764
const Key_Hangul_PostHanja = 65339
const Key_Hangul_PreHanja = 65338
const Key_Hangul_PreviousCandidate = 65342
const Key_Hangul_Rieul = 3753
const Key_Hangul_RieulHieuh = 3760
const Key_Hangul_RieulKiyeog = 3754
const Key_Hangul_RieulMieum = 3755
const Key_Hangul_RieulPhieuf = 3759
const Key_Hangul_RieulPieub = 3756
const Key_Hangul_RieulSios = 3757
const Key_Hangul_RieulTieut = 3758
const Key_Hangul_RieulYeorinHieuh = 3823
const Key_Hangul_Romaja = 65334
const Key_Hangul_SingleCandidate = 65340
const Key_Hangul_Sios = 3765
const Key_Hangul_Special = 65343
const Key_Hangul_SsangDikeud = 3752
const Key_Hangul_SsangJieuj = 3769
const Key_Hangul_SsangKiyeog = 3746
const Key_Hangul_SsangPieub = 3763
const Key_Hangul_SsangSios = 3766
const Key_Hangul_Start = 65330
const Key_Hangul_SunkyeongeumMieum = 3824
const Key_Hangul_SunkyeongeumPhieuf = 3828
const Key_Hangul_SunkyeongeumPieub = 3825
const Key_Hangul_Tieut = 3772
const Key_Hangul_U = 3788
const Key_Hangul_WA = 3784
const Key_Hangul_WAE = 3785
const Key_Hangul_WE = 3790
const Key_Hangul_WEO = 3789
const Key_Hangul_WI = 3791
const Key_Hangul_YA = 3777
const Key_Hangul_YAE = 3778
const Key_Hangul_YE = 3782
const Key_Hangul_YEO = 3781
const Key_Hangul_YI = 3794
const Key_Hangul_YO = 3787
const Key_Hangul_YU = 3792
const Key_Hangul_YeorinHieuh = 3829
const Key_Hangul_switch = 65406
const Key_Hankaku = 65321
const Key_Hcircumflex = 678
const Key_Hebrew_switch = 65406
const Key_Help = 65386
const Key_Henkan = 65315
const Key_Henkan_Mode = 65315
const Key_Hibernate = 269025192
const Key_Hiragana = 65317
const Key_Hiragana_Katakana = 65319
const Key_History = 269025079
const Key_Home = 65360
const Key_HomePage = 269025048
const Key_HotLinks = 269025082
const Key_Hstroke = 673
const Key_Hyper_L = 65517
const Key_Hyper_R = 65518
const Key_I = 73
const Key_ISO_Center_Object = 65075
const Key_ISO_Continuous_Underline = 65072
const Key_ISO_Discontinuous_Underline = 65073
const Key_ISO_Emphasize = 65074
const Key_ISO_Enter = 65076
const Key_ISO_Fast_Cursor_Down = 65071
const Key_ISO_Fast_Cursor_Left = 65068
const Key_ISO_Fast_Cursor_Right = 65069
const Key_ISO_Fast_Cursor_Up = 65070
const Key_ISO_First_Group = 65036
const Key_ISO_First_Group_Lock = 65037
const Key_ISO_Group_Latch = 65030
const Key_ISO_Group_Lock = 65031
const Key_ISO_Group_Shift = 65406
const Key_ISO_Last_Group = 65038
const Key_ISO_Last_Group_Lock = 65039
const Key_ISO_Left_Tab = 65056
const Key_ISO_Level2_Latch = 65026
const Key_ISO_Level3_Latch = 65028
const Key_ISO_Level3_Lock = 65029
const Key_ISO_Level3_Shift = 65027
const Key_ISO_Level5_Latch = 65042
const Key_ISO_Level5_Lock = 65043
const Key_ISO_Level5_Shift = 65041
const Key_ISO_Lock = 65025
const Key_ISO_Move_Line_Down = 65058
const Key_ISO_Move_Line_Up = 65057
const Key_ISO_Next_Group = 65032
const Key_ISO_Next_Group_Lock = 65033
const Key_ISO_Partial_Line_Down = 65060
const Key_ISO_Partial_Line_Up = 65059
const Key_ISO_Partial_Space_Left = 65061
const Key_ISO_Partial_Space_Right = 65062
const Key_ISO_Prev_Group = 65034
const Key_ISO_Prev_Group_Lock = 65035
const Key_ISO_Release_Both_Margins = 65067
const Key_ISO_Release_Margin_Left = 65065
const Key_ISO_Release_Margin_Right = 65066
const Key_ISO_Set_Margin_Left = 65063
const Key_ISO_Set_Margin_Right = 65064
const Key_Iabovedot = 681
const Key_Iacute = 205
const Key_Ibelowdot = 16785098
const Key_Ibreve = 16777516
const Key_Icircumflex = 206
const Key_Idiaeresis = 207
const Key_Igrave = 204
const Key_Ihook = 16785096
const Key_Imacron = 975
const Key_Insert = 65379
const Key_Iogonek = 967
const Key_Itilde = 933
const Key_J = 74
const Key_Jcircumflex = 684
const Key_K = 75
const Key_KP_0 = 65456
const Key_KP_1 = 65457
const Key_KP_2 = 65458
const Key_KP_3 = 65459
const Key_KP_4 = 65460
const Key_KP_5 = 65461
const Key_KP_6 = 65462
const Key_KP_7 = 65463
const Key_KP_8 = 65464
const Key_KP_9 = 65465
const Key_KP_Add = 65451
const Key_KP_Begin = 65437
const Key_KP_Decimal = 65454
const Key_KP_Delete = 65439
const Key_KP_Divide = 65455
const Key_KP_Down = 65433
const Key_KP_End = 65436
const Key_KP_Enter = 65421
const Key_KP_Equal = 65469
const Key_KP_F1 = 65425
const Key_KP_F2 = 65426
const Key_KP_F3 = 65427
const Key_KP_F4 = 65428
const Key_KP_Home = 65429
const Key_KP_Insert = 65438
const Key_KP_Left = 65430
const Key_KP_Multiply = 65450
const Key_KP_Next = 65435
const Key_KP_Page_Down = 65435
const Key_KP_Page_Up = 65434
const Key_KP_Prior = 65434
const Key_KP_Right = 65432
const Key_KP_Separator = 65452
const Key_KP_Space = 65408
const Key_KP_Subtract = 65453
const Key_KP_Tab = 65417
const Key_KP_Up = 65431
const Key_Kana_Lock = 65325
const Key_Kana_Shift = 65326
const Key_Kanji = 65313
const Key_Kanji_Bangou = 65335
const Key_Katakana = 65318
const Key_KbdBrightnessDown = 269025030
const Key_KbdBrightnessUp = 269025029
const Key_KbdLightOnOff = 269025028
const Key_Kcedilla = 979
const Key_Korean_Won = 3839
const Key_L = 76
const Key_L1 = 65480
const Key_L10 = 65489
const Key_L2 = 65481
const Key_L3 = 65482
const Key_L4 = 65483
const Key_L5 = 65484
const Key_L6 = 65485
const Key_L7 = 65486
const Key_L8 = 65487
const Key_L9 = 65488
const Key_Lacute = 453
const Key_Last_Virtual_Screen = 65236
const Key_Launch0 = 269025088
const Key_Launch1 = 269025089
const Key_Launch2 = 269025090
const Key_Launch3 = 269025091
const Key_Launch4 = 269025092
const Key_Launch5 = 269025093
const Key_Launch6 = 269025094
const Key_Launch7 = 269025095
const Key_Launch8 = 269025096
const Key_Launch9 = 269025097
const Key_LaunchA = 269025098
const Key_LaunchB = 269025099
const Key_LaunchC = 269025100
const Key_LaunchD = 269025101
const Key_LaunchE = 269025102
const Key_LaunchF = 269025103
const Key_Lbelowdot = 16784950
const Key_Lcaron = 421
const Key_Lcedilla = 934
const Key_Left = 65361
const Key_LightBulb = 269025077
const Key_Linefeed = 65290
const Key_LiraSign = 16785572
const Key_LogOff = 269025121
const Key_Lstroke = 419
const Key_M = 77
const Key_Mabovedot = 16784960
const Key_Macedonia_DSE = 1717
const Key_Macedonia_GJE = 1714
const Key_Macedonia_KJE = 1724
const Key_Macedonia_dse = 1701
const Key_Macedonia_gje = 1698
const Key_Macedonia_kje = 1708
const Key_Mae_Koho = 65342
const Key_Mail = 269025049
const Key_MailForward = 269025168
const Key_Market = 269025122
const Key_Massyo = 65324
const Key_Meeting = 269025123
const Key_Memo = 269025054
const Key_Menu = 65383
const Key_MenuKB = 269025125
const Key_MenuPB = 269025126
const Key_Messenger = 269025166
const Key_Meta_L = 65511
const Key_Meta_R = 65512
const Key_MillSign = 16785573
const Key_ModeLock = 269025025
const Key_Mode_switch = 65406
const Key_MonBrightnessDown = 269025027
const Key_MonBrightnessUp = 269025026
const Key_MouseKeys_Accel_Enable = 65143
const Key_MouseKeys_Enable = 65142
const Key_Muhenkan = 65314
const Key_Multi_key = 65312
const Key_MultipleCandidate = 65341
const Key_Music = 269025170
const Key_MyComputer = 269025075
const Key_MySites = 269025127
const Key_N = 78
const Key_Nacute = 465
const Key_NairaSign = 16785574
const Key_Ncaron = 466
const Key_Ncedilla = 977
const Key_New = 269025128
const Key_NewSheqelSign = 16785578
const Key_News = 269025129
const Key_Next = 65366
const Key_Next_VMode = 269024802
const Key_Next_Virtual_Screen = 65234
const Key_Ntilde = 209
const Key_Num_Lock = 65407
const Key_O = 79
const Key_OE = 5052
const Key_Oacute = 211
const Key_Obarred = 16777631
const Key_Obelowdot = 16785100
const Key_Ocaron = 16777681
const Key_Ocircumflex = 212
const Key_Ocircumflexacute = 16785104
const Key_Ocircumflexbelowdot = 16785112
const Key_Ocircumflexgrave = 16785106
const Key_Ocircumflexhook = 16785108
const Key_Ocircumflextilde = 16785110
const Key_Odiaeresis = 214
const Key_Odoubleacute = 469
const Key_OfficeHome = 269025130
const Key_Ograve = 210
const Key_Ohook = 16785102
const Key_Ohorn = 16777632
const Key_Ohornacute = 16785114
const Key_Ohornbelowdot = 16785122
const Key_Ohorngrave = 16785116
const Key_Ohornhook = 16785118
const Key_Ohorntilde = 16785120
const Key_Omacron = 978
const Key_Ooblique = 216
const Key_Open = 269025131
const Key_OpenURL = 269025080
const Key_Option = 269025132
const Key_Oslash = 216
const Key_Otilde = 213
const Key_Overlay1_Enable = 65144
const Key_Overlay2_Enable = 65145
const Key_P = 80
const Key_Pabovedot = 16784982
const Key_Page_Down = 65366
const Key_Page_Up = 65365
const Key_Paste = 269025133
const Key_Pause = 65299
const Key_PesetaSign = 16785575
const Key_Phone = 269025134
const Key_Pictures = 269025169
const Key_Pointer_Accelerate = 65274
const Key_Pointer_Button1 = 65257
const Key_Pointer_Button2 = 65258
const Key_Pointer_Button3 = 65259
const Key_Pointer_Button4 = 65260
const Key_Pointer_Button5 = 65261
const Key_Pointer_Button_Dflt = 65256
const Key_Pointer_DblClick1 = 65263
const Key_Pointer_DblClick2 = 65264
const Key_Pointer_DblClick3 = 65265
const Key_Pointer_DblClick4 = 65266
const Key_Pointer_DblClick5 = 65267
const Key_Pointer_DblClick_Dflt = 65262
const Key_Pointer_DfltBtnNext = 65275
const Key_Pointer_DfltBtnPrev = 65276
const Key_Pointer_Down = 65251
const Key_Pointer_DownLeft = 65254
const Key_Pointer_DownRight = 65255
const Key_Pointer_Drag1 = 65269
const Key_Pointer_Drag2 = 65270
const Key_Pointer_Drag3 = 65271
const Key_Pointer_Drag4 = 65272
const Key_Pointer_Drag5 = 65277
const Key_Pointer_Drag_Dflt = 65268
const Key_Pointer_EnableKeys = 65273
const Key_Pointer_Left = 65248
const Key_Pointer_Right = 65249
const Key_Pointer_Up = 65250
const Key_Pointer_UpLeft = 65252
const Key_Pointer_UpRight = 65253
const Key_PowerDown = 269025057
const Key_PowerOff = 269025066
const Key_Prev_VMode = 269024803
const Key_Prev_Virtual_Screen = 65233
const Key_PreviousCandidate = 65342
const Key_Print = 65377
const Key_Prior = 65365
const Key_Q = 81
const Key_R = 82
const Key_R1 = 65490
const Key_R10 = 65499
const Key_R11 = 65500
const Key_R12 = 65501
const Key_R13 = 65502
const Key_R14 = 65503
const Key_R15 = 65504
const Key_R2 = 65491
const Key_R3 = 65492
const Key_R4 = 65493
const Key_R5 = 65494
const Key_R6 = 65495
const Key_R7 = 65496
const Key_R8 = 65497
const Key_R9 = 65498
const Key_Racute = 448
const Key_Rcaron = 472
const Key_Rcedilla = 931
const Key_Red = 269025187
const Key_Redo = 65382
const Key_Refresh = 269025065
const Key_Reload = 269025139
const Key_RepeatKeys_Enable = 65138
const Key_Reply = 269025138
const Key_Return = 65293
const Key_Right = 65363
const Key_RockerDown = 269025060
const Key_RockerEnter = 269025061
const Key_RockerUp = 269025059
const Key_Romaji = 65316
const Key_RotateWindows = 269025140
const Key_RotationKB = 269025142
const Key_RotationPB = 269025141
const Key_RupeeSign = 16785576
const Key_S = 83
const Key_SCHWA = 16777615
const Key_Sabovedot = 16784992
const Key_Sacute = 422
const Key_Save = 269025143
const Key_Scaron = 425
const Key_Scedilla = 426
const Key_Scircumflex = 734
const Key_ScreenSaver = 269025069
const Key_ScrollClick = 269025146
const Key_ScrollDown = 269025145
const Key_ScrollUp = 269025144
const Key_Scroll_Lock = 65300
const Key_Search = 269025051
const Key_Select = 65376
const Key_SelectButton = 269025184
const Key_Send = 269025147
const Key_Serbian_DJE = 1713
const Key_Serbian_DZE = 1727
const Key_Serbian_JE = 1720
const Key_Serbian_LJE = 1721
const Key_Serbian_NJE = 1722
const Key_Serbian_TSHE = 1723
const Key_Serbian_dje = 1697
const Key_Serbian_dze = 1711
const Key_Serbian_je = 1704
const Key_Serbian_lje = 1705
const Key_Serbian_nje = 1706
const Key_Serbian_tshe = 1707
const Key_Shift_L = 65505
const Key_Shift_Lock = 65510
const Key_Shift_R = 65506
const Key_Shop = 269025078
const Key_SingleCandidate = 65340
const Key_Sinh_a = 16780677
const Key_Sinh_aa = 16780678
const Key_Sinh_aa2 = 16780751
const Key_Sinh_ae = 16780679
const Key_Sinh_ae2 = 16780752
const Key_Sinh_aee = 16780680
const Key_Sinh_aee2 = 16780753
const Key_Sinh_ai = 16780691
const Key_Sinh_ai2 = 16780763
const Key_Sinh_al = 16780746
const Key_Sinh_au = 16780694
const Key_Sinh_au2 = 16780766
const Key_Sinh_ba = 16780726
const Key_Sinh_bha = 16780727
const Key_Sinh_ca = 16780704
const Key_Sinh_cha = 16780705
const Key_Sinh_dda = 16780713
const Key_Sinh_ddha = 16780714
const Key_Sinh_dha = 16780719
const Key_Sinh_dhha = 16780720
const Key_Sinh_e = 16780689
const Key_Sinh_e2 = 16780761
const Key_Sinh_ee = 16780690
const Key_Sinh_ee2 = 16780762
const Key_Sinh_fa = 16780742
const Key_Sinh_ga = 16780700
const Key_Sinh_gha = 16780701
const Key_Sinh_h2 = 16780675
const Key_Sinh_ha = 16780740
const Key_Sinh_i = 16780681
const Key_Sinh_i2 = 16780754
const Key_Sinh_ii = 16780682
const Key_Sinh_ii2 = 16780755
const Key_Sinh_ja = 16780706
const Key_Sinh_jha = 16780707
const Key_Sinh_jnya = 16780709
const Key_Sinh_ka = 16780698
const Key_Sinh_kha = 16780699
const Key_Sinh_kunddaliya = 16780788
const Key_Sinh_la = 16780733
const Key_Sinh_lla = 16780741
const Key_Sinh_lu = 16780687
const Key_Sinh_lu2 = 16780767
const Key_Sinh_luu = 16780688
const Key_Sinh_luu2 = 16780787
const Key_Sinh_ma = 16780728
const Key_Sinh_mba = 16780729
const Key_Sinh_na = 16780721
const Key_Sinh_ndda = 16780716
const Key_Sinh_ndha = 16780723
const Key_Sinh_ng = 16780674
const Key_Sinh_ng2 = 16780702
const Key_Sinh_nga = 16780703
const Key_Sinh_nja = 16780710
const Key_Sinh_nna = 16780715
const Key_Sinh_nya = 16780708
const Key_Sinh_o = 16780692
const Key_Sinh_o2 = 16780764
const Key_Sinh_oo = 16780693
const Key_Sinh_oo2 = 16780765
const Key_Sinh_pa = 16780724
const Key_Sinh_pha = 16780725
const Key_Sinh_ra = 16780731
const Key_Sinh_ri = 16780685
const Key_Sinh_rii = 16780686
const Key_Sinh_ru2 = 16780760
const Key_Sinh_ruu2 = 16780786
const Key_Sinh_sa = 16780739
const Key_Sinh_sha = 16780737
const Key_Sinh_ssha = 16780738
const Key_Sinh_tha = 16780717
const Key_Sinh_thha = 16780718
const Key_Sinh_tta = 16780711
const Key_Sinh_ttha = 16780712
const Key_Sinh_u = 16780683
const Key_Sinh_u2 = 16780756
const Key_Sinh_uu = 16780684
const Key_Sinh_uu2 = 16780758
const Key_Sinh_va = 16780736
const Key_Sinh_ya = 16780730
const Key_Sleep = 269025071
const Key_SlowKeys_Enable = 65139
const Key_Spell = 269025148
const Key_SplitScreen = 269025149
const Key_Standby = 269025040
const Key_Start = 269025050
const Key_StickyKeys_Enable = 65141
const Key_Stop = 269025064
const Key_Subtitle = 269025178
const Key_Super_L = 65515
const Key_Super_R = 65516
const Key_Support = 269025150
const Key_Suspend = 269025191
const Key_Switch_VT_1 = 269024769
const Key_Switch_VT_10 = 269024778
const Key_Switch_VT_11 = 269024779
const Key_Switch_VT_12 = 269024780
const Key_Switch_VT_2 = 269024770
const Key_Switch_VT_3 = 269024771
const Key_Switch_VT_4 = 269024772
const Key_Switch_VT_5 = 269024773
const Key_Switch_VT_6 = 269024774
const Key_Switch_VT_7 = 269024775
const Key_Switch_VT_8 = 269024776
const Key_Switch_VT_9 = 269024777
const Key_Sys_Req = 65301
const Key_T = 84
const Key_THORN = 222
const Key_Tab = 65289
const Key_Tabovedot = 16785002
const Key_TaskPane = 269025151
const Key_Tcaron = 427
const Key_Tcedilla = 478
const Key_Terminal = 269025152
const Key_Terminate_Server = 65237
const Key_Thai_baht = 3551
const Key_Thai_bobaimai = 3514
const Key_Thai_chochan = 3496
const Key_Thai_chochang = 3498
const Key_Thai_choching = 3497
const Key_Thai_chochoe = 3500
const Key_Thai_dochada = 3502
const Key_Thai_dodek = 3508
const Key_Thai_fofa = 3517
const Key_Thai_fofan = 3519
const Key_Thai_hohip = 3531
const Key_Thai_honokhuk = 3534
const Key_Thai_khokhai = 3490
const Key_Thai_khokhon = 3493
const Key_Thai_khokhuat = 3491
const Key_Thai_khokhwai = 3492
const Key_Thai_khorakhang = 3494
const Key_Thai_kokai = 3489
const Key_Thai_lakkhangyao = 3557
const Key_Thai_lekchet = 3575
const Key_Thai_lekha = 3573
const Key_Thai_lekhok = 3574
const Key_Thai_lekkao = 3577
const Key_Thai_leknung = 3569
const Key_Thai_lekpaet = 3576
const Key_Thai_leksam = 3571
const Key_Thai_leksi = 3572
const Key_Thai_leksong = 3570
const Key_Thai_leksun = 3568
const Key_Thai_lochula = 3532
const Key_Thai_loling = 3525
const Key_Thai_lu = 3526
const Key_Thai_maichattawa = 3563
const Key_Thai_maiek = 3560
const Key_Thai_maihanakat = 3537
const Key_Thai_maihanakat_maitho = 3550
const Key_Thai_maitaikhu = 3559
const Key_Thai_maitho = 3561
const Key_Thai_maitri = 3562
const Key_Thai_maiyamok = 3558
const Key_Thai_moma = 3521
const Key_Thai_ngongu = 3495
const Key_Thai_nikhahit = 3565
const Key_Thai_nonen = 3507
const Key_Thai_nonu = 3513
const Key_Thai_oang = 3533
const Key_Thai_paiyannoi = 3535
const Key_Thai_phinthu = 3546
const Key_Thai_phophan = 3518
const Key_Thai_phophung = 3516
const Key_Thai_phosamphao = 3520
const Key_Thai_popla = 3515
const Key_Thai_rorua = 3523
const Key_Thai_ru = 3524
const Key_Thai_saraa = 3536
const Key_Thai_saraaa = 3538
const Key_Thai_saraae = 3553
const Key_Thai_saraaimaimalai = 3556
const Key_Thai_saraaimaimuan = 3555
const Key_Thai_saraam = 3539
const Key_Thai_sarae = 3552
const Key_Thai_sarai = 3540
const Key_Thai_saraii = 3541
const Key_Thai_sarao = 3554
const Key_Thai_sarau = 3544
const Key_Thai_saraue = 3542
const Key_Thai_sarauee = 3543
const Key_Thai_sarauu = 3545
const Key_Thai_sorusi = 3529
const Key_Thai_sosala = 3528
const Key_Thai_soso = 3499
const Key_Thai_sosua = 3530
const Key_Thai_thanthakhat = 3564
const Key_Thai_thonangmontho = 3505
const Key_Thai_thophuthao = 3506
const Key_Thai_thothahan = 3511
const Key_Thai_thothan = 3504
const Key_Thai_thothong = 3512
const Key_Thai_thothung = 3510
const Key_Thai_topatak = 3503
const Key_Thai_totao = 3509
const Key_Thai_wowaen = 3527
const Key_Thai_yoyak = 3522
const Key_Thai_yoying = 3501
const Key_Thorn = 222
const Key_Time = 269025183
const Key_ToDoList = 269025055
const Key_Tools = 269025153
const Key_TopMenu = 269025186
const Key_TouchpadOff = 269025201
const Key_TouchpadOn = 269025200
const Key_TouchpadToggle = 269025193
const Key_Touroku = 65323
const Key_Travel = 269025154
const Key_Tslash = 940
const Key_U = 85
const Key_UWB = 269025174
const Key_Uacute = 218
const Key_Ubelowdot = 16785124
const Key_Ubreve = 733
const Key_Ucircumflex = 219
const Key_Udiaeresis = 220
const Key_Udoubleacute = 475
const Key_Ugrave = 217
const Key_Uhook = 16785126
const Key_Uhorn = 16777647
const Key_Uhornacute = 16785128
const Key_Uhornbelowdot = 16785136
const Key_Uhorngrave = 16785130
const Key_Uhornhook = 16785132
const Key_Uhorntilde = 16785134
const Key_Ukrainian_GHE_WITH_UPTURN = 1725
const Key_Ukrainian_I = 1718
const Key_Ukrainian_IE = 1716
const Key_Ukrainian_YI = 1719
const Key_Ukrainian_ghe_with_upturn = 1709
const Key_Ukrainian_i = 1702
const Key_Ukrainian_ie = 1700
const Key_Ukrainian_yi = 1703
const Key_Ukranian_I = 1718
const Key_Ukranian_JE = 1716
const Key_Ukranian_YI = 1719
const Key_Ukranian_i = 1702
const Key_Ukranian_je = 1700
const Key_Ukranian_yi = 1703
const Key_Umacron = 990
const Key_Undo = 65381
const Key_Ungrab = 269024800
const Key_Uogonek = 985
const Key_Up = 65362
const Key_Uring = 473
const Key_User1KB = 269025157
const Key_User2KB = 269025158
const Key_UserPB = 269025156
const Key_Utilde = 989
const Key_V = 86
const Key_VendorHome = 269025076
const Key_Video = 269025159
const Key_View = 269025185
const Key_VoidSymbol = 16777215
const Key_W = 87
const Key_WLAN = 269025173
const Key_WWW = 269025070
const Key_Wacute = 16785026
const Key_WakeUp = 269025067
const Key_Wcircumflex = 16777588
const Key_Wdiaeresis = 16785028
const Key_WebCam = 269025167
const Key_Wgrave = 16785024
const Key_WheelButton = 269025160
const Key_WindowClear = 269025109
const Key_WonSign = 16785577
const Key_Word = 269025161
const Key_X = 88
const Key_Xabovedot = 16785034
const Key_Xfer = 269025162
const Key_Y = 89
const Key_Yacute = 221
const Key_Ybelowdot = 16785140
const Key_Ycircumflex = 16777590
const Key_Ydiaeresis = 5054
const Key_Yellow = 269025189
const Key_Ygrave = 16785138
const Key_Yhook = 16785142
const Key_Ytilde = 16785144
const Key_Z = 90
const Key_Zabovedot = 431
const Key_Zacute = 428
const Key_Zcaron = 430
const Key_Zen_Koho = 65341
const Key_Zenkaku = 65320
const Key_Zenkaku_Hankaku = 65322
const Key_ZoomIn = 269025163
const Key_ZoomOut = 269025164
const Key_Zstroke = 16777653
const Key_a = 97
const Key_aacute = 225
const Key_abelowdot = 16785057
const Key_abovedot = 511
const Key_abreve = 483
const Key_abreveacute = 16785071
const Key_abrevebelowdot = 16785079
const Key_abrevegrave = 16785073
const Key_abrevehook = 16785075
const Key_abrevetilde = 16785077
const Key_acircumflex = 226
const Key_acircumflexacute = 16785061
const Key_acircumflexbelowdot = 16785069
const Key_acircumflexgrave = 16785063
const Key_acircumflexhook = 16785065
const Key_acircumflextilde = 16785067
const Key_acute = 180
const Key_adiaeresis = 228
const Key_ae = 230
const Key_agrave = 224
const Key_ahook = 16785059
const Key_amacron = 992
const Key_ampersand = 38
const Key_aogonek = 433
const Key_apostrophe = 39
const Key_approxeq = 16785992
const Key_approximate = 2248
const Key_aring = 229
const Key_asciicircum = 94
const Key_asciitilde = 126
const Key_asterisk = 42
const Key_at = 64
const Key_atilde = 227
const Key_b = 98
const Key_babovedot = 16784899
const Key_backslash = 92
const Key_ballotcross = 2804
const Key_bar = 124
const Key_because = 16785973
const Key_blank = 2527
const Key_botintegral = 2213
const Key_botleftparens = 2220
const Key_botleftsqbracket = 2216
const Key_botleftsummation = 2226
const Key_botrightparens = 2222
const Key_botrightsqbracket = 2218
const Key_botrightsummation = 2230
const Key_bott = 2550
const Key_botvertsummationconnector = 2228
const Key_braceleft = 123
const Key_braceright = 125
const Key_bracketleft = 91
const Key_bracketright = 93
const Key_braille_blank = 16787456
const Key_braille_dot_1 = 65521
const Key_braille_dot_10 = 65530
const Key_braille_dot_2 = 65522
const Key_braille_dot_3 = 65523
const Key_braille_dot_4 = 65524
const Key_braille_dot_5 = 65525
const Key_braille_dot_6 = 65526
const Key_braille_dot_7 = 65527
const Key_braille_dot_8 = 65528
const Key_braille_dot_9 = 65529
const Key_braille_dots_1 = 16787457
const Key_braille_dots_12 = 16787459
const Key_braille_dots_123 = 16787463
const Key_braille_dots_1234 = 16787471
const Key_braille_dots_12345 = 16787487
const Key_braille_dots_123456 = 16787519
const Key_braille_dots_1234567 = 16787583
const Key_braille_dots_12345678 = 16787711
const Key_braille_dots_1234568 = 16787647
const Key_braille_dots_123457 = 16787551
const Key_braille_dots_1234578 = 16787679
const Key_braille_dots_123458 = 16787615
const Key_braille_dots_12346 = 16787503
const Key_braille_dots_123467 = 16787567
const Key_braille_dots_1234678 = 16787695
const Key_braille_dots_123468 = 16787631
const Key_braille_dots_12347 = 16787535
const Key_braille_dots_123478 = 16787663
const Key_braille_dots_12348 = 16787599
const Key_braille_dots_1235 = 16787479
const Key_braille_dots_12356 = 16787511
const Key_braille_dots_123567 = 16787575
const Key_braille_dots_1235678 = 16787703
const Key_braille_dots_123568 = 16787639
const Key_braille_dots_12357 = 16787543
const Key_braille_dots_123578 = 16787671
const Key_braille_dots_12358 = 16787607
const Key_braille_dots_1236 = 16787495
const Key_braille_dots_12367 = 16787559
const Key_braille_dots_123678 = 16787687
const Key_braille_dots_12368 = 16787623
const Key_braille_dots_1237 = 16787527
const Key_braille_dots_12378 = 16787655
const Key_braille_dots_1238 = 16787591
const Key_braille_dots_124 = 16787467
const Key_braille_dots_1245 = 16787483
const Key_braille_dots_12456 = 16787515
const Key_braille_dots_124567 = 16787579
const Key_braille_dots_1245678 = 16787707
const Key_braille_dots_124568 = 16787643
const Key_braille_dots_12457 = 16787547
const Key_braille_dots_124578 = 16787675
const Key_braille_dots_12458 = 16787611
const Key_braille_dots_1246 = 16787499
const Key_braille_dots_12467 = 16787563
const Key_braille_dots_124678 = 16787691
const Key_braille_dots_12468 = 16787627
const Key_braille_dots_1247 = 16787531
const Key_braille_dots_12478 = 16787659
const Key_braille_dots_1248 = 16787595
const Key_braille_dots_125 = 16787475
const Key_braille_dots_1256 = 16787507
const Key_braille_dots_12567 = 16787571
const Key_braille_dots_125678 = 16787699
const Key_braille_dots_12568 = 16787635
const Key_braille_dots_1257 = 16787539
const Key_braille_dots_12578 = 16787667
const Key_braille_dots_1258 = 16787603
const Key_braille_dots_126 = 16787491
const Key_braille_dots_1267 = 16787555
const Key_braille_dots_12678 = 16787683
const Key_braille_dots_1268 = 16787619
const Key_braille_dots_127 = 16787523
const Key_braille_dots_1278 = 16787651
const Key_braille_dots_128 = 16787587
const Key_braille_dots_13 = 16787461
const Key_braille_dots_134 = 16787469
const Key_braille_dots_1345 = 16787485
const Key_braille_dots_13456 = 16787517
const Key_braille_dots_134567 = 16787581
const Key_braille_dots_1345678 = 16787709
const Key_braille_dots_134568 = 16787645
const Key_braille_dots_13457 = 16787549
const Key_braille_dots_134578 = 16787677
const Key_braille_dots_13458 = 16787613
const Key_braille_dots_1346 = 16787501
const Key_braille_dots_13467 = 16787565
const Key_braille_dots_134678 = 16787693
const Key_braille_dots_13468 = 16787629
const Key_braille_dots_1347 = 16787533
const Key_braille_dots_13478 = 16787661
const Key_braille_dots_1348 = 16787597
const Key_braille_dots_135 = 16787477
const Key_braille_dots_1356 = 16787509
const Key_braille_dots_13567 = 16787573
const Key_braille_dots_135678 = 16787701
const Key_braille_dots_13568 = 16787637
const Key_braille_dots_1357 = 16787541
const Key_braille_dots_13578 = 16787669
const Key_braille_dots_1358 = 16787605
const Key_braille_dots_136 = 16787493
const Key_braille_dots_1367 = 16787557
const Key_braille_dots_13678 = 16787685
const Key_braille_dots_1368 = 16787621
const Key_braille_dots_137 = 16787525
const Key_braille_dots_1378 = 16787653
const Key_braille_dots_138 = 16787589
const Key_braille_dots_14 = 16787465
const Key_braille_dots_145 = 16787481
const Key_braille_dots_1456 = 16787513
const Key_braille_dots_14567 = 16787577
const Key_braille_dots_145678 = 16787705
const Key_braille_dots_14568 = 16787641
const Key_braille_dots_1457 = 16787545
const Key_braille_dots_14578 = 16787673
const Key_braille_dots_1458 = 16787609
const Key_braille_dots_146 = 16787497
const Key_braille_dots_1467 = 16787561
const Key_braille_dots_14678 = 16787689
const Key_braille_dots_1468 = 16787625
const Key_braille_dots_147 = 16787529
const Key_braille_dots_1478 = 16787657
const Key_braille_dots_148 = 16787593
const Key_braille_dots_15 = 16787473
const Key_braille_dots_156 = 16787505
const Key_braille_dots_1567 = 16787569
const Key_braille_dots_15678 = 16787697
const Key_braille_dots_1568 = 16787633
const Key_braille_dots_157 = 16787537
const Key_braille_dots_1578 = 16787665
const Key_braille_dots_158 = 16787601
const Key_braille_dots_16 = 16787489
const Key_braille_dots_167 = 16787553
const Key_braille_dots_1678 = 16787681
const Key_braille_dots_168 = 16787617
const Key_braille_dots_17 = 16787521
const Key_braille_dots_178 = 16787649
const Key_braille_dots_18 = 16787585
const Key_braille_dots_2 = 16787458
const Key_braille_dots_23 = 16787462
const Key_braille_dots_234 = 16787470
const Key_braille_dots_2345 = 16787486
const Key_braille_dots_23456 = 16787518
const Key_braille_dots_234567 = 16787582
const Key_braille_dots_2345678 = 16787710
const Key_braille_dots_234568 = 16787646
const Key_braille_dots_23457 = 16787550
const Key_braille_dots_234578 = 16787678
const Key_braille_dots_23458 = 16787614
const Key_braille_dots_2346 = 16787502
const Key_braille_dots_23467 = 16787566
const Key_braille_dots_234678 = 16787694
const Key_braille_dots_23468 = 16787630
const Key_braille_dots_2347 = 16787534
const Key_braille_dots_23478 = 16787662
const Key_braille_dots_2348 = 16787598
const Key_braille_dots_235 = 16787478
const Key_braille_dots_2356 = 16787510
const Key_braille_dots_23567 = 16787574
const Key_braille_dots_235678 = 16787702
const Key_braille_dots_23568 = 16787638
const Key_braille_dots_2357 = 16787542
const Key_braille_dots_23578 = 16787670
const Key_braille_dots_2358 = 16787606
const Key_braille_dots_236 = 16787494
const Key_braille_dots_2367 = 16787558
const Key_braille_dots_23678 = 16787686
const Key_braille_dots_2368 = 16787622
const Key_braille_dots_237 = 16787526
const Key_braille_dots_2378 = 16787654
const Key_braille_dots_238 = 16787590
const Key_braille_dots_24 = 16787466
const Key_braille_dots_245 = 16787482
const Key_braille_dots_2456 = 16787514
const Key_braille_dots_24567 = 16787578
const Key_braille_dots_245678 = 16787706
const Key_braille_dots_24568 = 16787642
const Key_braille_dots_2457 = 16787546
const Key_braille_dots_24578 = 16787674
const Key_braille_dots_2458 = 16787610
const Key_braille_dots_246 = 16787498
const Key_braille_dots_2467 = 16787562
const Key_braille_dots_24678 = 16787690
const Key_braille_dots_2468 = 16787626
const Key_braille_dots_247 = 16787530
const Key_braille_dots_2478 = 16787658
const Key_braille_dots_248 = 16787594
const Key_braille_dots_25 = 16787474
const Key_braille_dots_256 = 16787506
const Key_braille_dots_2567 = 16787570
const Key_braille_dots_25678 = 16787698
const Key_braille_dots_2568 = 16787634
const Key_braille_dots_257 = 16787538
const Key_braille_dots_2578 = 16787666
const Key_braille_dots_258 = 16787602
const Key_braille_dots_26 = 16787490
const Key_braille_dots_267 = 16787554
const Key_braille_dots_2678 = 16787682
const Key_braille_dots_268 = 16787618
const Key_braille_dots_27 = 16787522
const Key_braille_dots_278 = 16787650
const Key_braille_dots_28 = 16787586
const Key_braille_dots_3 = 16787460
const Key_braille_dots_34 = 16787468
const Key_braille_dots_345 = 16787484
const Key_braille_dots_3456 = 16787516
const Key_braille_dots_34567 = 16787580
const Key_braille_dots_345678 = 16787708
const Key_braille_dots_34568 = 16787644
const Key_braille_dots_3457 = 16787548
const Key_braille_dots_34578 = 16787676
const Key_braille_dots_3458 = 16787612
const Key_braille_dots_346 = 16787500
const Key_braille_dots_3467 = 16787564
const Key_braille_dots_34678 = 16787692
const Key_braille_dots_3468 = 16787628
const Key_braille_dots_347 = 16787532
const Key_braille_dots_3478 = 16787660
const Key_braille_dots_348 = 16787596
const Key_braille_dots_35 = 16787476
const Key_braille_dots_356 = 16787508
const Key_braille_dots_3567 = 16787572
const Key_braille_dots_35678 = 16787700
const Key_braille_dots_3568 = 16787636
const Key_braille_dots_357 = 16787540
const Key_braille_dots_3578 = 16787668
const Key_braille_dots_358 = 16787604
const Key_braille_dots_36 = 16787492
const Key_braille_dots_367 = 16787556
const Key_braille_dots_3678 = 16787684
const Key_braille_dots_368 = 16787620
const Key_braille_dots_37 = 16787524
const Key_braille_dots_378 = 16787652
const Key_braille_dots_38 = 16787588
const Key_braille_dots_4 = 16787464
const Key_braille_dots_45 = 16787480
const Key_braille_dots_456 = 16787512
const Key_braille_dots_4567 = 16787576
const Key_braille_dots_45678 = 16787704
const Key_braille_dots_4568 = 16787640
const Key_braille_dots_457 = 16787544
const Key_braille_dots_4578 = 16787672
const Key_braille_dots_458 = 16787608
const Key_braille_dots_46 = 16787496
const Key_braille_dots_467 = 16787560
const Key_braille_dots_4678 = 16787688
const Key_braille_dots_468 = 16787624
const Key_braille_dots_47 = 16787528
const Key_braille_dots_478 = 16787656
const Key_braille_dots_48 = 16787592
const Key_braille_dots_5 = 16787472
const Key_braille_dots_56 = 16787504
const Key_braille_dots_567 = 16787568
const Key_braille_dots_5678 = 16787696
const Key_braille_dots_568 = 16787632
const Key_braille_dots_57 = 16787536
const Key_braille_dots_578 = 16787664
const Key_braille_dots_58 = 16787600
const Key_braille_dots_6 = 16787488
const Key_braille_dots_67 = 16787552
const Key_braille_dots_678 = 16787680
const Key_braille_dots_68 = 16787616
const Key_braille_dots_7 = 16787520
const Key_braille_dots_78 = 16787648
const Key_braille_dots_8 = 16787584
const Key_breve = 418
const Key_brokenbar = 166
const Key_c = 99
const Key_cabovedot = 741
const Key_cacute = 486
const Key_careof = 2744
const Key_caret = 2812
const Key_caron = 439
const Key_ccaron = 488
const Key_ccedilla = 231
const Key_ccircumflex = 742
const Key_cedilla = 184
const Key_cent = 162
const Key_checkerboard = 2529
const Key_checkmark = 2803
const Key_circle = 3023
const Key_club = 2796
const Key_colon = 58
const Key_comma = 44
const Key_containsas = 16785931
const Key_copyright = 169
const Key_cr = 2532
const Key_crossinglines = 2542
const Key_cuberoot = 16785947
const Key_currency = 164
const Key_cursor = 2815
const Key_d = 100
const Key_dabovedot = 16784907
const Key_dagger = 2801
const Key_dcaron = 495
const Key_dead_A = 65153
const Key_dead_E = 65155
const Key_dead_I = 65157
const Key_dead_O = 65159
const Key_dead_U = 65161
const Key_dead_a = 65152
const Key_dead_abovecomma = 65124
const Key_dead_abovedot = 65110
const Key_dead_abovereversedcomma = 65125
const Key_dead_abovering = 65112
const Key_dead_acute = 65105
const Key_dead_belowbreve = 65131
const Key_dead_belowcircumflex = 65129
const Key_dead_belowcomma = 65134
const Key_dead_belowdiaeresis = 65132
const Key_dead_belowdot = 65120
const Key_dead_belowmacron = 65128
const Key_dead_belowring = 65127
const Key_dead_belowtilde = 65130
const Key_dead_breve = 65109
const Key_dead_capital_schwa = 65163
const Key_dead_caron = 65114
const Key_dead_cedilla = 65115
const Key_dead_circumflex = 65106
const Key_dead_currency = 65135
const Key_dead_dasia = 65125
const Key_dead_diaeresis = 65111
const Key_dead_doubleacute = 65113
const Key_dead_doublegrave = 65126
const Key_dead_e = 65154
const Key_dead_grave = 65104
const Key_dead_hook = 65121
const Key_dead_horn = 65122
const Key_dead_i = 65156
const Key_dead_invertedbreve = 65133
const Key_dead_iota = 65117
const Key_dead_macron = 65108
const Key_dead_o = 65158
const Key_dead_ogonek = 65116
const Key_dead_perispomeni = 65107
const Key_dead_psili = 65124
const Key_dead_semivoiced_sound = 65119
const Key_dead_small_schwa = 65162
const Key_dead_stroke = 65123
const Key_dead_tilde = 65107
const Key_dead_u = 65160
const Key_dead_voiced_sound = 65118
const Key_decimalpoint = 2749
const Key_degree = 176
const Key_diaeresis = 168
const Key_diamond = 2797
const Key_digitspace = 2725
const Key_dintegral = 16785964
const Key_division = 247
const Key_dollar = 36
const Key_doubbaselinedot = 2735
const Key_doubleacute = 445
const Key_doubledagger = 2802
const Key_doublelowquotemark = 2814
const Key_downarrow = 2302
const Key_downcaret = 2984
const Key_downshoe = 3030
const Key_downstile = 3012
const Key_downtack = 3010
const Key_dstroke = 496
const Key_e = 101
const Key_eabovedot = 1004
const Key_eacute = 233
const Key_ebelowdot = 16785081
const Key_ecaron = 492
const Key_ecircumflex = 234
const Key_ecircumflexacute = 16785087
const Key_ecircumflexbelowdot = 16785095
const Key_ecircumflexgrave = 16785089
const Key_ecircumflexhook = 16785091
const Key_ecircumflextilde = 16785093
const Key_ediaeresis = 235
const Key_egrave = 232
const Key_ehook = 16785083
const Key_eightsubscript = 16785544
const Key_eightsuperior = 16785528
const Key_elementof = 16785928
const Key_ellipsis = 2734
const Key_em3space = 2723
const Key_em4space = 2724
const Key_emacron = 954
const Key_emdash = 2729
const Key_emfilledcircle = 2782
const Key_emfilledrect = 2783
const Key_emopencircle = 2766
const Key_emopenrectangle = 2767
const Key_emptyset = 16785925
const Key_emspace = 2721
const Key_endash = 2730
const Key_enfilledcircbullet = 2790
const Key_enfilledsqbullet = 2791
const Key_eng = 959
const Key_enopencircbullet = 2784
const Key_enopensquarebullet = 2785
const Key_enspace = 2722
const Key_eogonek = 490
const Key_equal = 61
const Key_eth = 240
const Key_etilde = 16785085
const Key_exclam = 33
const Key_exclamdown = 161
const Key_f = 102
const Key_fabovedot = 16784927
const Key_femalesymbol = 2808
const Key_ff = 2531
const Key_figdash = 2747
const Key_filledlefttribullet = 2780
const Key_filledrectbullet = 2779
const Key_filledrighttribullet = 2781
const Key_filledtribulletdown = 2793
const Key_filledtribulletup = 2792
const Key_fiveeighths = 2757
const Key_fivesixths = 2743
const Key_fivesubscript = 16785541
const Key_fivesuperior = 16785525
const Key_fourfifths = 2741
const Key_foursubscript = 16785540
const Key_foursuperior = 16785524
const Key_fourthroot = 16785948
const Key_function = 2294
const Key_g = 103
const Key_gabovedot = 757
const Key_gbreve = 699
const Key_gcaron = 16777703
const Key_gcedilla = 955
const Key_gcircumflex = 760
const Key_grave = 96
const Key_greater = 62
const Key_greaterthanequal = 2238
const Key_guillemotleft = 171
const Key_guillemotright = 187
const Key_h = 104
const Key_hairspace = 2728
const Key_hcircumflex = 694
const Key_heart = 2798
const Key_hebrew_aleph = 3296
const Key_hebrew_ayin = 3314
const Key_hebrew_bet = 3297
const Key_hebrew_beth = 3297
const Key_hebrew_chet = 3303
const Key_hebrew_dalet = 3299
const Key_hebrew_daleth = 3299
const Key_hebrew_doublelowline = 3295
const Key_hebrew_finalkaph = 3306
const Key_hebrew_finalmem = 3309
const Key_hebrew_finalnun = 3311
const Key_hebrew_finalpe = 3315
const Key_hebrew_finalzade = 3317
const Key_hebrew_finalzadi = 3317
const Key_hebrew_gimel = 3298
const Key_hebrew_gimmel = 3298
const Key_hebrew_he = 3300
const Key_hebrew_het = 3303
const Key_hebrew_kaph = 3307
const Key_hebrew_kuf = 3319
const Key_hebrew_lamed = 3308
const Key_hebrew_mem = 3310
const Key_hebrew_nun = 3312
const Key_hebrew_pe = 3316
const Key_hebrew_qoph = 3319
const Key_hebrew_resh = 3320
const Key_hebrew_samech = 3313
const Key_hebrew_samekh = 3313
const Key_hebrew_shin = 3321
const Key_hebrew_taf = 3322
const Key_hebrew_taw = 3322
const Key_hebrew_tet = 3304
const Key_hebrew_teth = 3304
const Key_hebrew_waw = 3301
const Key_hebrew_yod = 3305
const Key_hebrew_zade = 3318
const Key_hebrew_zadi = 3318
const Key_hebrew_zain = 3302
const Key_hebrew_zayin = 3302
const Key_hexagram = 2778
const Key_horizconnector = 2211
const Key_horizlinescan1 = 2543
const Key_horizlinescan3 = 2544
const Key_horizlinescan5 = 2545
const Key_horizlinescan7 = 2546
const Key_horizlinescan9 = 2547
const Key_hstroke = 689
const Key_ht = 2530
const Key_hyphen = 173
const Key_i = 105
const Key_iTouch = 269025120
const Key_iacute = 237
const Key_ibelowdot = 16785099
const Key_ibreve = 16777517
const Key_icircumflex = 238
const Key_identical = 2255
const Key_idiaeresis = 239
const Key_idotless = 697
const Key_ifonlyif = 2253
const Key_igrave = 236
const Key_ihook = 16785097
const Key_imacron = 1007
const Key_implies = 2254
const Key_includedin = 2266
const Key_includes = 2267
const Key_infinity = 2242
const Key_integral = 2239
const Key_intersection = 2268
const Key_iogonek = 999
const Key_itilde = 949
const Key_j = 106
const Key_jcircumflex = 700
const Key_jot = 3018
const Key_k = 107
const Key_kana_A = 1201
const Key_kana_CHI = 1217
const Key_kana_E = 1204
const Key_kana_FU = 1228
const Key_kana_HA = 1226
const Key_kana_HE = 1229
const Key_kana_HI = 1227
const Key_kana_HO = 1230
const Key_kana_HU = 1228
const Key_kana_I = 1202
const Key_kana_KA = 1206
const Key_kana_KE = 1209
const Key_kana_KI = 1207
const Key_kana_KO = 1210
const Key_kana_KU = 1208
const Key_kana_MA = 1231
const Key_kana_ME = 1234
const Key_kana_MI = 1232
const Key_kana_MO = 1235
const Key_kana_MU = 1233
const Key_kana_N = 1245
const Key_kana_NA = 1221
const Key_kana_NE = 1224
const Key_kana_NI = 1222
const Key_kana_NO = 1225
const Key_kana_NU = 1223
const Key_kana_O = 1205
const Key_kana_RA = 1239
const Key_kana_RE = 1242
const Key_kana_RI = 1240
const Key_kana_RO = 1243
const Key_kana_RU = 1241
const Key_kana_SA = 1211
const Key_kana_SE = 1214
const Key_kana_SHI = 1212
const Key_kana_SO = 1215
const Key_kana_SU = 1213
const Key_kana_TA = 1216
const Key_kana_TE = 1219
const Key_kana_TI = 1217
const Key_kana_TO = 1220
const Key_kana_TSU = 1218
const Key_kana_TU = 1218
const Key_kana_U = 1203
const Key_kana_WA = 1244
const Key_kana_WO = 1190
const Key_kana_YA = 1236
const Key_kana_YO = 1238
const Key_kana_YU = 1237
const Key_kana_a = 1191
const Key_kana_closingbracket = 1187
const Key_kana_comma = 1188
const Key_kana_conjunctive = 1189
const Key_kana_e = 1194
const Key_kana_fullstop = 1185
const Key_kana_i = 1192
const Key_kana_middledot = 1189
const Key_kana_o = 1195
const Key_kana_openingbracket = 1186
const Key_kana_switch = 65406
const Key_kana_tsu = 1199
const Key_kana_tu = 1199
const Key_kana_u = 1193
const Key_kana_ya = 1196
const Key_kana_yo = 1198
const Key_kana_yu = 1197
const Key_kappa = 930
const Key_kcedilla = 1011
const Key_kra = 930
const Key_l = 108
const Key_lacute = 485
const Key_latincross = 2777
const Key_lbelowdot = 16784951
const Key_lcaron = 437
const Key_lcedilla = 950
const Key_leftanglebracket = 2748
const Key_leftarrow = 2299
const Key_leftcaret = 2979
const Key_leftdoublequotemark = 2770
const Key_leftmiddlecurlybrace = 2223
const Key_leftopentriangle = 2764
const Key_leftpointer = 2794
const Key_leftradical = 2209
const Key_leftshoe = 3034
const Key_leftsinglequotemark = 2768
const Key_leftt = 2548
const Key_lefttack = 3036
const Key_less = 60
const Key_lessthanequal = 2236
const Key_lf = 2533
const Key_logicaland = 2270
const Key_logicalor = 2271
const Key_lowleftcorner = 2541
const Key_lowrightcorner = 2538
const Key_lstroke = 435
const Key_m = 109
const Key_mabovedot = 16784961
const Key_macron = 175
const Key_malesymbol = 2807
const Key_maltesecross = 2800
const Key_marker = 2751
const Key_masculine = 186
const Key_minus = 45
const Key_minutes = 2774
const Key_mu = 181
const Key_multiply = 215
const Key_musicalflat = 2806
const Key_musicalsharp = 2805
const Key_n = 110
const Key_nabla = 2245
const Key_nacute = 497
const Key_ncaron = 498
const Key_ncedilla = 1009
const Key_ninesubscript = 16785545
const Key_ninesuperior = 16785529
const Key_nl = 2536
const Key_nobreakspace = 160
const Key_notapproxeq = 16785991
const Key_notelementof = 16785929
const Key_notequal = 2237
const Key_notidentical = 16786018
const Key_notsign = 172
const Key_ntilde = 241
const Key_numbersign = 35
const Key_numerosign = 1712
const Key_o = 111
const Key_oacute = 243
const Key_obarred = 16777845
const Key_obelowdot = 16785101
const Key_ocaron = 16777682
const Key_ocircumflex = 244
const Key_ocircumflexacute = 16785105
const Key_ocircumflexbelowdot = 16785113
const Key_ocircumflexgrave = 16785107
const Key_ocircumflexhook = 16785109
const Key_ocircumflextilde = 16785111
const Key_odiaeresis = 246
const Key_odoubleacute = 501
const Key_oe = 5053
const Key_ogonek = 434
const Key_ograve = 242
const Key_ohook = 16785103
const Key_ohorn = 16777633
const Key_ohornacute = 16785115
const Key_ohornbelowdot = 16785123
const Key_ohorngrave = 16785117
const Key_ohornhook = 16785119
const Key_ohorntilde = 16785121
const Key_omacron = 1010
const Key_oneeighth = 2755
const Key_onefifth = 2738
const Key_onehalf = 189
const Key_onequarter = 188
const Key_onesixth = 2742
const Key_onesubscript = 16785537
const Key_onesuperior = 185
const Key_onethird = 2736
const Key_ooblique = 248
const Key_openrectbullet = 2786
const Key_openstar = 2789
const Key_opentribulletdown = 2788
const Key_opentribulletup = 2787
const Key_ordfeminine = 170
const Key_oslash = 248
const Key_otilde = 245
const Key_overbar = 3008
const Key_overline = 1150
const Key_p = 112
const Key_pabovedot = 16784983
const Key_paragraph = 182
const Key_parenleft = 40
const Key_parenright = 41
const Key_partdifferential = 16785922
const Key_partialderivative = 2287
const Key_percent = 37
const Key_period = 46
const Key_periodcentered = 183
const Key_phonographcopyright = 2811
const Key_plus = 43
const Key_plusminus = 177
const Key_prescription = 2772
const Key_prolongedsound = 1200
const Key_punctspace = 2726
const Key_q = 113
const Key_quad = 3020
const Key_question = 63
const Key_questiondown = 191
const Key_quotedbl = 34
const Key_quoteleft = 96
const Key_quoteright = 39
const Key_r = 114
const Key_racute = 480
const Key_radical = 2262
const Key_rcaron = 504
const Key_rcedilla = 947
const Key_registered = 174
const Key_rightanglebracket = 2750
const Key_rightarrow = 2301
const Key_rightcaret = 2982
const Key_rightdoublequotemark = 2771
const Key_rightmiddlecurlybrace = 2224
const Key_rightmiddlesummation = 2231
const Key_rightopentriangle = 2765
const Key_rightpointer = 2795
const Key_rightshoe = 3032
const Key_rightsinglequotemark = 2769
const Key_rightt = 2549
const Key_righttack = 3068
const Key_s = 115
const Key_sabovedot = 16784993
const Key_sacute = 438
const Key_scaron = 441
const Key_scedilla = 442
const Key_schwa = 16777817
const Key_scircumflex = 766
const Key_script_switch = 65406
const Key_seconds = 2775
const Key_section = 167
const Key_semicolon = 59
const Key_semivoicedsound = 1247
const Key_seveneighths = 2758
const Key_sevensubscript = 16785543
const Key_sevensuperior = 16785527
const Key_signaturemark = 2762
const Key_signifblank = 2732
const Key_similarequal = 2249
const Key_singlelowquotemark = 2813
const Key_sixsubscript = 16785542
const Key_sixsuperior = 16785526
const Key_slash = 47
const Key_soliddiamond = 2528
const Key_space = 32
const Key_squareroot = 16785946
const Key_ssharp = 223
const Key_sterling = 163
const Key_stricteq = 16786019
const Key_t = 116
const Key_tabovedot = 16785003
const Key_tcaron = 443
const Key_tcedilla = 510
const Key_telephone = 2809
const Key_telephonerecorder = 2810
const Key_therefore = 2240
const Key_thinspace = 2727
const Key_thorn = 254
const Key_threeeighths = 2756
const Key_threefifths = 2740
const Key_threequarters = 190
const Key_threesubscript = 16785539
const Key_threesuperior = 179
const Key_tintegral = 16785965
const Key_topintegral = 2212
const Key_topleftparens = 2219
const Key_topleftradical = 2210
const Key_topleftsqbracket = 2215
const Key_topleftsummation = 2225
const Key_toprightparens = 2221
const Key_toprightsqbracket = 2217
const Key_toprightsummation = 2229
const Key_topt = 2551
const Key_topvertsummationconnector = 2227
const Key_trademark = 2761
const Key_trademarkincircle = 2763
const Key_tslash = 956
const Key_twofifths = 2739
const Key_twosubscript = 16785538
const Key_twosuperior = 178
const Key_twothirds = 2737
const Key_u = 117
const Key_uacute = 250
const Key_ubelowdot = 16785125
const Key_ubreve = 765
const Key_ucircumflex = 251
const Key_udiaeresis = 252
const Key_udoubleacute = 507
const Key_ugrave = 249
const Key_uhook = 16785127
const Key_uhorn = 16777648
const Key_uhornacute = 16785129
const Key_uhornbelowdot = 16785137
const Key_uhorngrave = 16785131
const Key_uhornhook = 16785133
const Key_uhorntilde = 16785135
const Key_umacron = 1022
const Key_underbar = 3014
const Key_underscore = 95
const Key_union = 2269
const Key_uogonek = 1017
const Key_uparrow = 2300
const Key_upcaret = 2985
const Key_upleftcorner = 2540
const Key_uprightcorner = 2539
const Key_upshoe = 3011
const Key_upstile = 3027
const Key_uptack = 3022
const Key_uring = 505
const Key_utilde = 1021
const Key_v = 118
const Key_variation = 2241
const Key_vertbar = 2552
const Key_vertconnector = 2214
const Key_voicedsound = 1246
const Key_vt = 2537
const Key_w = 119
const Key_wacute = 16785027
const Key_wcircumflex = 16777589
const Key_wdiaeresis = 16785029
const Key_wgrave = 16785025
const Key_x = 120
const Key_xabovedot = 16785035
const Key_y = 121
const Key_yacute = 253
const Key_ybelowdot = 16785141
const Key_ycircumflex = 16777591
const Key_ydiaeresis = 255
const Key_yen = 165
const Key_ygrave = 16785139
const Key_yhook = 16785143
const Key_ytilde = 16785145
const Key_z = 122
const Key_zabovedot = 447
const Key_zacute = 444
const Key_zcaron = 446
const Key_zerosubscript = 16785536
const Key_zerosuperior = 16785520
const Key_zstroke = 16777654
type KeymapLike interface {
	gobject.ObjectLike
	InheritedFromGdkKeymap() *C.GdkKeymap
}

type Keymap struct {
	gobject.Object
	
}

func ToKeymap(objlike gobject.ObjectLike) *Keymap {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Keymap)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Keymap)(obj)
	}
	panic("cannot cast to Keymap")
}

func (this0 *Keymap) InheritedFromGdkKeymap() *C.GdkKeymap {
	if this0 == nil {
		return nil
	}
	return (*C.GdkKeymap)(this0.C)
}

func (this0 *Keymap) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_keymap_get_type())
}

func KeymapGetType() gobject.Type {
	return (*Keymap)(nil).GetStaticType()
}
func KeymapGetDefault() *Keymap {
	ret1 := C.gdk_keymap_get_default()
	var ret2 *Keymap
	ret2 = (*Keymap)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func KeymapGetForDisplay(display0 DisplayLike) *Keymap {
	var display1 *C.GdkDisplay
	if display0 != nil {
		display1 = display0.InheritedFromGdkDisplay()
	}
	ret1 := C.gdk_keymap_get_for_display(display1)
	var ret2 *Keymap
	ret2 = (*Keymap)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Keymap) AddVirtualModifiers() ModifierType {
	var this1 *C.GdkKeymap
	var state1 C.GdkModifierType
	if this0 != nil {
		this1 = this0.InheritedFromGdkKeymap()
	}
	C.gdk_keymap_add_virtual_modifiers(this1, &state1)
	var state2 ModifierType
	state2 = ModifierType(state1)
	return state2
}
func (this0 *Keymap) GetCapsLockState() bool {
	var this1 *C.GdkKeymap
	if this0 != nil {
		this1 = this0.InheritedFromGdkKeymap()
	}
	ret1 := C.gdk_keymap_get_caps_lock_state(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Keymap) GetDirection() pango.Direction {
	var this1 *C.GdkKeymap
	if this0 != nil {
		this1 = this0.InheritedFromGdkKeymap()
	}
	ret1 := C.gdk_keymap_get_direction(this1)
	var ret2 pango.Direction
	ret2 = pango.Direction(ret1)
	return ret2
}
// blacklisted: Keymap.get_entries_for_keycode (method)
func (this0 *Keymap) GetEntriesForKeyval(keyval0 int) ([]KeymapKey, bool) {
	var this1 *C.GdkKeymap
	var keyval1 C.uint32_t
	var keys1 *C.GdkKeymapKey
	var n_keys1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkKeymap()
	}
	keyval1 = C.uint32_t(keyval0)
	ret1 := C.gdk_keymap_get_entries_for_keyval(this1, keyval1, &keys1, &n_keys1)
	var keys2 []KeymapKey
	var ret2 bool
	keys2 = make([]KeymapKey, n_keys1)
	for i := range keys2 {
		keys2[i] = *(*KeymapKey)(unsafe.Pointer(&(*(*[999999]C.GdkKeymapKey)(unsafe.Pointer(keys1)))[i]))
	}
	ret2 = ret1 != 0
	return keys2, ret2
}
func (this0 *Keymap) GetNumLockState() bool {
	var this1 *C.GdkKeymap
	if this0 != nil {
		this1 = this0.InheritedFromGdkKeymap()
	}
	ret1 := C.gdk_keymap_get_num_lock_state(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Keymap) HaveBidiLayouts() bool {
	var this1 *C.GdkKeymap
	if this0 != nil {
		this1 = this0.InheritedFromGdkKeymap()
	}
	ret1 := C.gdk_keymap_have_bidi_layouts(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Keymap) LookupKey(key0 *KeymapKey) int {
	var this1 *C.GdkKeymap
	var key1 *C.GdkKeymapKey
	if this0 != nil {
		this1 = this0.InheritedFromGdkKeymap()
	}
	key1 = (*C.GdkKeymapKey)(unsafe.Pointer(key0))
	ret1 := C.gdk_keymap_lookup_key(this1, key1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Keymap) MapVirtualModifiers() (ModifierType, bool) {
	var this1 *C.GdkKeymap
	var state1 C.GdkModifierType
	if this0 != nil {
		this1 = this0.InheritedFromGdkKeymap()
	}
	ret1 := C.gdk_keymap_map_virtual_modifiers(this1, &state1)
	var state2 ModifierType
	var ret2 bool
	state2 = ModifierType(state1)
	ret2 = ret1 != 0
	return state2, ret2
}
func (this0 *Keymap) TranslateKeyboardState(hardware_keycode0 int, state0 ModifierType, group0 int) (int, int, int, ModifierType, bool) {
	var this1 *C.GdkKeymap
	var hardware_keycode1 C.uint32_t
	var state1 C.GdkModifierType
	var group1 C.int32_t
	var keyval1 C.uint32_t
	var effective_group1 C.int32_t
	var level1 C.int32_t
	var consumed_modifiers1 C.GdkModifierType
	if this0 != nil {
		this1 = this0.InheritedFromGdkKeymap()
	}
	hardware_keycode1 = C.uint32_t(hardware_keycode0)
	state1 = C.GdkModifierType(state0)
	group1 = C.int32_t(group0)
	ret1 := C.gdk_keymap_translate_keyboard_state(this1, hardware_keycode1, state1, group1, &keyval1, &effective_group1, &level1, &consumed_modifiers1)
	var keyval2 int
	var effective_group2 int
	var level2 int
	var consumed_modifiers2 ModifierType
	var ret2 bool
	keyval2 = int(keyval1)
	effective_group2 = int(effective_group1)
	level2 = int(level1)
	consumed_modifiers2 = ModifierType(consumed_modifiers1)
	ret2 = ret1 != 0
	return keyval2, effective_group2, level2, consumed_modifiers2, ret2
}
type KeymapKey struct {
	Keycode uint32
	Group int32
	Level int32
}
const MaxTimecoordAxes = 128
type ModifierType C.uint32_t
const (
	ModifierTypeShiftMask ModifierType = 1
	ModifierTypeLockMask ModifierType = 2
	ModifierTypeControlMask ModifierType = 4
	ModifierTypeMod1Mask ModifierType = 8
	ModifierTypeMod2Mask ModifierType = 16
	ModifierTypeMod3Mask ModifierType = 32
	ModifierTypeMod4Mask ModifierType = 64
	ModifierTypeMod5Mask ModifierType = 128
	ModifierTypeButton1Mask ModifierType = 256
	ModifierTypeButton2Mask ModifierType = 512
	ModifierTypeButton3Mask ModifierType = 1024
	ModifierTypeButton4Mask ModifierType = 2048
	ModifierTypeButton5Mask ModifierType = 4096
	ModifierTypeModifierReserved13Mask ModifierType = 8192
	ModifierTypeModifierReserved14Mask ModifierType = 16384
	ModifierTypeModifierReserved15Mask ModifierType = 32768
	ModifierTypeModifierReserved16Mask ModifierType = 65536
	ModifierTypeModifierReserved17Mask ModifierType = 131072
	ModifierTypeModifierReserved18Mask ModifierType = 262144
	ModifierTypeModifierReserved19Mask ModifierType = 524288
	ModifierTypeModifierReserved20Mask ModifierType = 1048576
	ModifierTypeModifierReserved21Mask ModifierType = 2097152
	ModifierTypeModifierReserved22Mask ModifierType = 4194304
	ModifierTypeModifierReserved23Mask ModifierType = 8388608
	ModifierTypeModifierReserved24Mask ModifierType = 16777216
	ModifierTypeModifierReserved25Mask ModifierType = 33554432
	ModifierTypeSuperMask ModifierType = 67108864
	ModifierTypeHyperMask ModifierType = 134217728
	ModifierTypeMetaMask ModifierType = 268435456
	ModifierTypeModifierReserved29Mask ModifierType = 536870912
	ModifierTypeReleaseMask ModifierType = 1073741824
	ModifierTypeModifierMask ModifierType = 1543512063
)
type NotifyType C.uint32_t
const (
	NotifyTypeAncestor NotifyType = 0
	NotifyTypeVirtual NotifyType = 1
	NotifyTypeInferior NotifyType = 2
	NotifyTypeNonlinear NotifyType = 3
	NotifyTypeNonlinearVirtual NotifyType = 4
	NotifyTypeUnknown NotifyType = 5
)
type OwnerChange C.uint32_t
const (
	OwnerChangeNewOwner OwnerChange = 0
	OwnerChangeDestroy OwnerChange = 1
	OwnerChangeClose OwnerChange = 2
)
const ParentRelative = 1
const PriorityRedraw = 20
type Point struct {
	X int32
	Y int32
}
type PropMode C.uint32_t
const (
	PropModeReplace PropMode = 0
	PropModePrepend PropMode = 1
	PropModeAppend PropMode = 2
)
type PropertyState C.uint32_t
const (
	PropertyStateNewValue PropertyState = 0
	PropertyStateDelete PropertyState = 1
)
type RGBA struct {
	Red float64
	Green float64
	Blue float64
	Alpha float64
}
func (this0 *RGBA) Copy() *RGBA {
	var this1 *C.GdkRGBA
	this1 = (*C.GdkRGBA)(unsafe.Pointer(this0))
	ret1 := C.gdk_rgba_copy(this1)
	var ret2 *RGBA
	ret2 = (*RGBA)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *RGBA) Equal(p20 *RGBA) bool {
	var this1 *C.GdkRGBA
	var p21 *C.GdkRGBA
	this1 = (*C.GdkRGBA)(unsafe.Pointer(this0))
	p21 = (*C.GdkRGBA)(unsafe.Pointer(p20))
	ret1 := C.gdk_rgba_equal(this1, p21)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *RGBA) Free() {
	var this1 *C.GdkRGBA
	this1 = (*C.GdkRGBA)(unsafe.Pointer(this0))
	C.gdk_rgba_free(this1)
}
func (this0 *RGBA) Hash() int {
	var this1 *C.GdkRGBA
	this1 = (*C.GdkRGBA)(unsafe.Pointer(this0))
	ret1 := C.gdk_rgba_hash(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *RGBA) Parse(spec0 string) bool {
	var this1 *C.GdkRGBA
	var spec1 *C.char
	this1 = (*C.GdkRGBA)(unsafe.Pointer(this0))
	spec1 = _GoStringToGString(spec0)
	defer C.free(unsafe.Pointer(spec1))
	ret1 := C.gdk_rgba_parse(this1, spec1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *RGBA) ToString() string {
	var this1 *C.GdkRGBA
	this1 = (*C.GdkRGBA)(unsafe.Pointer(this0))
	ret1 := C.gdk_rgba_to_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
type ScreenLike interface {
	gobject.ObjectLike
	InheritedFromGdkScreen() *C.GdkScreen
}

type Screen struct {
	gobject.Object
	
}

func ToScreen(objlike gobject.ObjectLike) *Screen {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Screen)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Screen)(obj)
	}
	panic("cannot cast to Screen")
}

func (this0 *Screen) InheritedFromGdkScreen() *C.GdkScreen {
	if this0 == nil {
		return nil
	}
	return (*C.GdkScreen)(this0.C)
}

func (this0 *Screen) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_screen_get_type())
}

func ScreenGetType() gobject.Type {
	return (*Screen)(nil).GetStaticType()
}
func ScreenGetDefault() *Screen {
	ret1 := C.gdk_screen_get_default()
	var ret2 *Screen
	ret2 = (*Screen)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func ScreenHeight() int {
	ret1 := C.gdk_screen_height()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func ScreenHeightMm() int {
	ret1 := C.gdk_screen_height_mm()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func ScreenWidth() int {
	ret1 := C.gdk_screen_width()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func ScreenWidthMm() int {
	ret1 := C.gdk_screen_width_mm()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetActiveWindow() *Window {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_active_window(this1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Screen) GetDisplay() *Display {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_display(this1)
	var ret2 *Display
	ret2 = (*Display)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Screen) GetFontOptions() *cairo.FontOptions {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_font_options(this1)
	var ret2 *cairo.FontOptions
	ret2 = (*cairo.FontOptions)(cairo.FontOptionsWrap(unsafe.Pointer(ret1)))
	return ret2
}
func (this0 *Screen) GetHeight() int {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_height(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetHeightMm() int {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_height_mm(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetMonitorAtPoint(x0 int, y0 int) int {
	var this1 *C.GdkScreen
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	ret1 := C.gdk_screen_get_monitor_at_point(this1, x1, y1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetMonitorAtWindow(window0 WindowLike) int {
	var this1 *C.GdkScreen
	var window1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_screen_get_monitor_at_window(this1, window1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetMonitorGeometry(monitor_num0 int) cairo.RectangleInt {
	var this1 *C.GdkScreen
	var monitor_num1 C.int32_t
	var dest1 C.cairoRectangleInt
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	monitor_num1 = C.int32_t(monitor_num0)
	C.gdk_screen_get_monitor_geometry(this1, monitor_num1, &dest1)
	var dest2 cairo.RectangleInt
	dest2 = *(*cairo.RectangleInt)(unsafe.Pointer(&dest1))
	return dest2
}
func (this0 *Screen) GetMonitorHeightMm(monitor_num0 int) int {
	var this1 *C.GdkScreen
	var monitor_num1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	monitor_num1 = C.int32_t(monitor_num0)
	ret1 := C.gdk_screen_get_monitor_height_mm(this1, monitor_num1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetMonitorPlugName(monitor_num0 int) string {
	var this1 *C.GdkScreen
	var monitor_num1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	monitor_num1 = C.int32_t(monitor_num0)
	ret1 := C.gdk_screen_get_monitor_plug_name(this1, monitor_num1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Screen) GetMonitorWidthMm(monitor_num0 int) int {
	var this1 *C.GdkScreen
	var monitor_num1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	monitor_num1 = C.int32_t(monitor_num0)
	ret1 := C.gdk_screen_get_monitor_width_mm(this1, monitor_num1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetNMonitors() int {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_n_monitors(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetNumber() int {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_number(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetPrimaryMonitor() int {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_primary_monitor(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetResolution() float64 {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_resolution(this1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *Screen) GetRGBAVisual() *Visual {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_rgba_visual(this1)
	var ret2 *Visual
	ret2 = (*Visual)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Screen) GetRootWindow() *Window {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_root_window(this1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Screen) GetSetting(name0 string, value0 *gobject.Value) bool {
	var this1 *C.GdkScreen
	var name1 *C.char
	var value1 *C.GValue
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	ret1 := C.gdk_screen_get_setting(this1, name1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Screen) GetSystemVisual() *Visual {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_system_visual(this1)
	var ret2 *Visual
	ret2 = (*Visual)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Screen) GetToplevelWindows() []*Window {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_toplevel_windows(this1)
	var ret2 []*Window
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Window
		elt = (*Window)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkWindow)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Screen) GetWidth() int {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetWidthMm() int {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_width_mm(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Screen) GetWindowStack() []*Window {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_get_window_stack(this1)
	var ret2 []*Window
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Window
		elt = (*Window)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkWindow)(iter.data)), false))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Screen) IsComposited() bool {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_is_composited(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Screen) ListVisuals() []*Visual {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_list_visuals(this1)
	var ret2 []*Visual
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Visual
		elt = (*Visual)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkVisual)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Screen) MakeDisplayName() string {
	var this1 *C.GdkScreen
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	ret1 := C.gdk_screen_make_display_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Screen) SetFontOptions(options0 *cairo.FontOptions) {
	var this1 *C.GdkScreen
	var options1 *C.cairoFontOptions
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	if options0 != nil {
		options1 = (*C.cairoFontOptions)(options0.C)
	}
	C.gdk_screen_set_font_options(this1, options1)
}
func (this0 *Screen) SetResolution(dpi0 float64) {
	var this1 *C.GdkScreen
	var dpi1 C.double
	if this0 != nil {
		this1 = this0.InheritedFromGdkScreen()
	}
	dpi1 = C.double(dpi0)
	C.gdk_screen_set_resolution(this1, dpi1)
}
type ScrollDirection C.uint32_t
const (
	ScrollDirectionUp ScrollDirection = 0
	ScrollDirectionDown ScrollDirection = 1
	ScrollDirectionLeft ScrollDirection = 2
	ScrollDirectionRight ScrollDirection = 3
)
type SettingAction C.uint32_t
const (
	SettingActionNew SettingAction = 0
	SettingActionChanged SettingAction = 1
	SettingActionDeleted SettingAction = 2
)
type Status C.int32_t
const (
	StatusOk Status = 0
	StatusError Status = -1
	StatusErrorParam Status = -2
	StatusErrorFile Status = -3
	StatusErrorMem Status = -4
)
type TimeCoord struct {
	Time uint32
	_ [4]byte
	Axes [128]float64
}
type VisibilityState C.uint32_t
const (
	VisibilityStateUnobscured VisibilityState = 0
	VisibilityStatePartial VisibilityState = 1
	VisibilityStateFullyObscured VisibilityState = 2
)
type VisualLike interface {
	gobject.ObjectLike
	InheritedFromGdkVisual() *C.GdkVisual
}

type Visual struct {
	gobject.Object
	
}

func ToVisual(objlike gobject.ObjectLike) *Visual {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Visual)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Visual)(obj)
	}
	panic("cannot cast to Visual")
}

func (this0 *Visual) InheritedFromGdkVisual() *C.GdkVisual {
	if this0 == nil {
		return nil
	}
	return (*C.GdkVisual)(this0.C)
}

func (this0 *Visual) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_visual_get_type())
}

func VisualGetType() gobject.Type {
	return (*Visual)(nil).GetStaticType()
}
func VisualGetBest() *Visual {
	ret1 := C.gdk_visual_get_best()
	var ret2 *Visual
	ret2 = (*Visual)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func VisualGetBestDepth() int {
	ret1 := C.gdk_visual_get_best_depth()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func VisualGetBestType() VisualType {
	ret1 := C.gdk_visual_get_best_type()
	var ret2 VisualType
	ret2 = VisualType(ret1)
	return ret2
}
func VisualGetBestWithBoth(depth0 int, visual_type0 VisualType) *Visual {
	var depth1 C.int32_t
	var visual_type1 C.GdkVisualType
	depth1 = C.int32_t(depth0)
	visual_type1 = C.GdkVisualType(visual_type0)
	ret1 := C.gdk_visual_get_best_with_both(depth1, visual_type1)
	var ret2 *Visual
	ret2 = (*Visual)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func VisualGetBestWithDepth(depth0 int) *Visual {
	var depth1 C.int32_t
	depth1 = C.int32_t(depth0)
	ret1 := C.gdk_visual_get_best_with_depth(depth1)
	var ret2 *Visual
	ret2 = (*Visual)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func VisualGetBestWithType(visual_type0 VisualType) *Visual {
	var visual_type1 C.GdkVisualType
	visual_type1 = C.GdkVisualType(visual_type0)
	ret1 := C.gdk_visual_get_best_with_type(visual_type1)
	var ret2 *Visual
	ret2 = (*Visual)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func VisualGetSystem() *Visual {
	ret1 := C.gdk_visual_get_system()
	var ret2 *Visual
	ret2 = (*Visual)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Visual) GetBitsPerRGB() int {
	var this1 *C.GdkVisual
	if this0 != nil {
		this1 = this0.InheritedFromGdkVisual()
	}
	ret1 := C.gdk_visual_get_bits_per_rgb(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Visual) GetBluePixelDetails() (int, int, int) {
	var this1 *C.GdkVisual
	var mask1 C.uint32_t
	var shift1 C.int32_t
	var precision1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkVisual()
	}
	C.gdk_visual_get_blue_pixel_details(this1, &mask1, &shift1, &precision1)
	var mask2 int
	var shift2 int
	var precision2 int
	mask2 = int(mask1)
	shift2 = int(shift1)
	precision2 = int(precision1)
	return mask2, shift2, precision2
}
func (this0 *Visual) GetByteOrder() ByteOrder {
	var this1 *C.GdkVisual
	if this0 != nil {
		this1 = this0.InheritedFromGdkVisual()
	}
	ret1 := C.gdk_visual_get_byte_order(this1)
	var ret2 ByteOrder
	ret2 = ByteOrder(ret1)
	return ret2
}
func (this0 *Visual) GetColormapSize() int {
	var this1 *C.GdkVisual
	if this0 != nil {
		this1 = this0.InheritedFromGdkVisual()
	}
	ret1 := C.gdk_visual_get_colormap_size(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Visual) GetDepth() int {
	var this1 *C.GdkVisual
	if this0 != nil {
		this1 = this0.InheritedFromGdkVisual()
	}
	ret1 := C.gdk_visual_get_depth(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Visual) GetGreenPixelDetails() (int, int, int) {
	var this1 *C.GdkVisual
	var mask1 C.uint32_t
	var shift1 C.int32_t
	var precision1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkVisual()
	}
	C.gdk_visual_get_green_pixel_details(this1, &mask1, &shift1, &precision1)
	var mask2 int
	var shift2 int
	var precision2 int
	mask2 = int(mask1)
	shift2 = int(shift1)
	precision2 = int(precision1)
	return mask2, shift2, precision2
}
func (this0 *Visual) GetRedPixelDetails() (int, int, int) {
	var this1 *C.GdkVisual
	var mask1 C.uint32_t
	var shift1 C.int32_t
	var precision1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkVisual()
	}
	C.gdk_visual_get_red_pixel_details(this1, &mask1, &shift1, &precision1)
	var mask2 int
	var shift2 int
	var precision2 int
	mask2 = int(mask1)
	shift2 = int(shift1)
	precision2 = int(precision1)
	return mask2, shift2, precision2
}
func (this0 *Visual) GetScreen() *Screen {
	var this1 *C.GdkVisual
	if this0 != nil {
		this1 = this0.InheritedFromGdkVisual()
	}
	ret1 := C.gdk_visual_get_screen(this1)
	var ret2 *Screen
	ret2 = (*Screen)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Visual) GetVisualType() VisualType {
	var this1 *C.GdkVisual
	if this0 != nil {
		this1 = this0.InheritedFromGdkVisual()
	}
	ret1 := C.gdk_visual_get_visual_type(this1)
	var ret2 VisualType
	ret2 = VisualType(ret1)
	return ret2
}
type VisualType C.uint32_t
const (
	VisualTypeStaticGray VisualType = 0
	VisualTypeGrayscale VisualType = 1
	VisualTypeStaticColor VisualType = 2
	VisualTypePseudoColor VisualType = 3
	VisualTypeTrueColor VisualType = 4
	VisualTypeDirectColor VisualType = 5
)
type WMDecoration C.uint32_t
const (
	WMDecorationAll WMDecoration = 1
	WMDecorationBorder WMDecoration = 2
	WMDecorationResizeh WMDecoration = 4
	WMDecorationTitle WMDecoration = 8
	WMDecorationMenu WMDecoration = 16
	WMDecorationMinimize WMDecoration = 32
	WMDecorationMaximize WMDecoration = 64
)
type WMFunction C.uint32_t
const (
	WMFunctionAll WMFunction = 1
	WMFunctionResize WMFunction = 2
	WMFunctionMove WMFunction = 4
	WMFunctionMinimize WMFunction = 8
	WMFunctionMaximize WMFunction = 16
	WMFunctionClose WMFunction = 32
)
type WindowLike interface {
	gobject.ObjectLike
	InheritedFromGdkWindow() *C.GdkWindow
}

type Window struct {
	gobject.Object
	
}

func ToWindow(objlike gobject.ObjectLike) *Window {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Window)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Window)(obj)
	}
	panic("cannot cast to Window")
}

func (this0 *Window) InheritedFromGdkWindow() *C.GdkWindow {
	if this0 == nil {
		return nil
	}
	return (*C.GdkWindow)(this0.C)
}

func (this0 *Window) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_window_get_type())
}

func WindowGetType() gobject.Type {
	return (*Window)(nil).GetStaticType()
}
func NewWindow(parent0 WindowLike, attributes0 *WindowAttr, attributes_mask0 int) *Window {
	var parent1 *C.GdkWindow
	var attributes1 *C.GdkWindowAttr
	var attributes_mask1 C.int32_t
	if parent0 != nil {
		parent1 = parent0.InheritedFromGdkWindow()
	}
	attributes1 = (*C.GdkWindowAttr)(unsafe.Pointer(attributes0))
	attributes_mask1 = C.int32_t(attributes_mask0)
	ret1 := C.gdk_window_new(parent1, attributes1, attributes_mask1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func WindowAtPointer() (int, int, *Window) {
	var win_x1 C.int32_t
	var win_y1 C.int32_t
	ret1 := C.gdk_window_at_pointer(&win_x1, &win_y1)
	var win_x2 int
	var win_y2 int
	var ret2 *Window
	win_x2 = int(win_x1)
	win_y2 = int(win_y1)
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return win_x2, win_y2, ret2
}
func WindowConstrainSize(geometry0 *Geometry, flags0 int, width0 int, height0 int) (int, int) {
	var geometry1 *C.GdkGeometry
	var flags1 C.uint32_t
	var width1 C.int32_t
	var height1 C.int32_t
	var new_width1 C.int32_t
	var new_height1 C.int32_t
	geometry1 = (*C.GdkGeometry)(unsafe.Pointer(geometry0))
	flags1 = C.uint32_t(flags0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	C.gdk_window_constrain_size(geometry1, flags1, width1, height1, &new_width1, &new_height1)
	var new_width2 int
	var new_height2 int
	new_width2 = int(new_width1)
	new_height2 = int(new_height1)
	return new_width2, new_height2
}
func WindowProcessAllUpdates() {
	C.gdk_window_process_all_updates()
}
func WindowSetDebugUpdates(setting0 bool) {
	var setting1 C.int
	setting1 = _GoBoolToCBool(setting0)
	C.gdk_window_set_debug_updates(setting1)
}
func (this0 *Window) Beep() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_beep(this1)
}
func (this0 *Window) BeginMoveDrag(button0 int, root_x0 int, root_y0 int, timestamp0 int) {
	var this1 *C.GdkWindow
	var button1 C.int32_t
	var root_x1 C.int32_t
	var root_y1 C.int32_t
	var timestamp1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	button1 = C.int32_t(button0)
	root_x1 = C.int32_t(root_x0)
	root_y1 = C.int32_t(root_y0)
	timestamp1 = C.uint32_t(timestamp0)
	C.gdk_window_begin_move_drag(this1, button1, root_x1, root_y1, timestamp1)
}
func (this0 *Window) BeginPaintRect(rectangle0 *cairo.RectangleInt) {
	var this1 *C.GdkWindow
	var rectangle1 *C.cairoRectangleInt
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	rectangle1 = (*C.cairoRectangleInt)(unsafe.Pointer(rectangle0))
	C.gdk_window_begin_paint_rect(this1, rectangle1)
}
func (this0 *Window) BeginPaintRegion(region0 *cairo.Region) {
	var this1 *C.GdkWindow
	var region1 *C.cairoRegion
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if region0 != nil {
		region1 = (*C.cairoRegion)(region0.C)
	}
	C.gdk_window_begin_paint_region(this1, region1)
}
func (this0 *Window) BeginResizeDrag(edge0 WindowEdge, button0 int, root_x0 int, root_y0 int, timestamp0 int) {
	var this1 *C.GdkWindow
	var edge1 C.GdkWindowEdge
	var button1 C.int32_t
	var root_x1 C.int32_t
	var root_y1 C.int32_t
	var timestamp1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	edge1 = C.GdkWindowEdge(edge0)
	button1 = C.int32_t(button0)
	root_x1 = C.int32_t(root_x0)
	root_y1 = C.int32_t(root_y0)
	timestamp1 = C.uint32_t(timestamp0)
	C.gdk_window_begin_resize_drag(this1, edge1, button1, root_x1, root_y1, timestamp1)
}
func (this0 *Window) ConfigureFinished() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_configure_finished(this1)
}
func (this0 *Window) CoordsFromParent(parent_x0 float64, parent_y0 float64) (float64, float64) {
	var this1 *C.GdkWindow
	var parent_x1 C.double
	var parent_y1 C.double
	var x1 C.double
	var y1 C.double
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	parent_x1 = C.double(parent_x0)
	parent_y1 = C.double(parent_y0)
	C.gdk_window_coords_from_parent(this1, parent_x1, parent_y1, &x1, &y1)
	var x2 float64
	var y2 float64
	x2 = float64(x1)
	y2 = float64(y1)
	return x2, y2
}
func (this0 *Window) CoordsToParent(x0 float64, y0 float64) (float64, float64) {
	var this1 *C.GdkWindow
	var x1 C.double
	var y1 C.double
	var parent_x1 C.double
	var parent_y1 C.double
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	x1 = C.double(x0)
	y1 = C.double(y0)
	C.gdk_window_coords_to_parent(this1, x1, y1, &parent_x1, &parent_y1)
	var parent_x2 float64
	var parent_y2 float64
	parent_x2 = float64(parent_x1)
	parent_y2 = float64(parent_y1)
	return parent_x2, parent_y2
}
func (this0 *Window) CreateSimilarSurface(content0 cairo.Content, width0 int, height0 int) *cairo.Surface {
	var this1 *C.GdkWindow
	var content1 C.cairoContent
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	content1 = C.cairoContent(content0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	ret1 := C.gdk_window_create_similar_surface(this1, content1, width1, height1)
	var ret2 *cairo.Surface
	ret2 = (*cairo.Surface)(cairo.SurfaceWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Window) Deiconify() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_deiconify(this1)
}
func (this0 *Window) Destroy() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_destroy(this1)
}
func (this0 *Window) DestroyNotify() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_destroy_notify(this1)
}
func (this0 *Window) EnableSynchronizedConfigure() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_enable_synchronized_configure(this1)
}
func (this0 *Window) EndPaint() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_end_paint(this1)
}
func (this0 *Window) EnsureNative() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_ensure_native(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) Flush() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_flush(this1)
}
func (this0 *Window) Focus(timestamp0 int) {
	var this1 *C.GdkWindow
	var timestamp1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	timestamp1 = C.uint32_t(timestamp0)
	C.gdk_window_focus(this1, timestamp1)
}
func (this0 *Window) FreezeToplevelUpdatesLibgtkOnly() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_freeze_toplevel_updates_libgtk_only(this1)
}
func (this0 *Window) FreezeUpdates() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_freeze_updates(this1)
}
func (this0 *Window) Fullscreen() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_fullscreen(this1)
}
func (this0 *Window) GeometryChanged() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_geometry_changed(this1)
}
func (this0 *Window) GetAcceptFocus() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_accept_focus(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) GetBackgroundPattern() *cairo.Pattern {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_background_pattern(this1)
	var ret2 *cairo.Pattern
	ret2 = (*cairo.Pattern)(cairo.PatternWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetChildren() []*Window {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_children(this1)
	var ret2 []*Window
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Window
		elt = (*Window)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkWindow)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Window) GetClipRegion() *cairo.Region {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_clip_region(this1)
	var ret2 *cairo.Region
	ret2 = (*cairo.Region)(cairo.RegionWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Window) GetComposited() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_composited(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) GetCursor() *Cursor {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_cursor(this1)
	var ret2 *Cursor
	ret2 = (*Cursor)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetDecorations() (WMDecoration, bool) {
	var this1 *C.GdkWindow
	var decorations1 C.GdkWMDecoration
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_decorations(this1, &decorations1)
	var decorations2 WMDecoration
	var ret2 bool
	decorations2 = WMDecoration(decorations1)
	ret2 = ret1 != 0
	return decorations2, ret2
}
func (this0 *Window) GetDeviceCursor(device0 DeviceLike) *Cursor {
	var this1 *C.GdkWindow
	var device1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_window_get_device_cursor(this1, device1)
	var ret2 *Cursor
	ret2 = (*Cursor)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetDeviceEvents(device0 DeviceLike) EventMask {
	var this1 *C.GdkWindow
	var device1 *C.GdkDevice
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_window_get_device_events(this1, device1)
	var ret2 EventMask
	ret2 = EventMask(ret1)
	return ret2
}
func (this0 *Window) GetDevicePosition(device0 DeviceLike) (int, int, ModifierType, *Window) {
	var this1 *C.GdkWindow
	var device1 *C.GdkDevice
	var x1 C.int32_t
	var y1 C.int32_t
	var mask1 C.GdkModifierType
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_window_get_device_position(this1, device1, &x1, &y1, &mask1)
	var x2 int
	var y2 int
	var mask2 ModifierType
	var ret2 *Window
	x2 = int(x1)
	y2 = int(y1)
	mask2 = ModifierType(mask1)
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return x2, y2, mask2, ret2
}
func (this0 *Window) GetDisplay() *Display {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_display(this1)
	var ret2 *Display
	ret2 = (*Display)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetDragProtocol() (*Window, DragProtocol) {
	var this1 *C.GdkWindow
	var target1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_drag_protocol(this1, &target1)
	var target2 *Window
	var ret2 DragProtocol
	target2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(target1), false))
	ret2 = DragProtocol(ret1)
	return target2, ret2
}
func (this0 *Window) GetEffectiveParent() *Window {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_effective_parent(this1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetEffectiveToplevel() *Window {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_effective_toplevel(this1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetEvents() EventMask {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_events(this1)
	var ret2 EventMask
	ret2 = EventMask(ret1)
	return ret2
}
func (this0 *Window) GetFocusOnMap() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_focus_on_map(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) GetFrameExtents(rect0 *cairo.RectangleInt) {
	var this1 *C.GdkWindow
	var rect1 *C.cairoRectangleInt
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	rect1 = (*C.cairoRectangleInt)(unsafe.Pointer(rect0))
	C.gdk_window_get_frame_extents(this1, rect1)
}
func (this0 *Window) GetGeometry() (int, int, int, int) {
	var this1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_get_geometry(this1, &x1, &y1, &width1, &height1)
	var x2 int
	var y2 int
	var width2 int
	var height2 int
	x2 = int(x1)
	y2 = int(y1)
	width2 = int(width1)
	height2 = int(height1)
	return x2, y2, width2, height2
}
func (this0 *Window) GetGroup() *Window {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_group(this1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetHeight() int {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_height(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Window) GetModalHint() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_modal_hint(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) GetOrigin() (int, int, int) {
	var this1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_origin(this1, &x1, &y1)
	var x2 int
	var y2 int
	var ret2 int
	x2 = int(x1)
	y2 = int(y1)
	ret2 = int(ret1)
	return x2, y2, ret2
}
func (this0 *Window) GetParent() *Window {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_parent(this1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetPointer() (int, int, ModifierType, *Window) {
	var this1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	var mask1 C.GdkModifierType
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_pointer(this1, &x1, &y1, &mask1)
	var x2 int
	var y2 int
	var mask2 ModifierType
	var ret2 *Window
	x2 = int(x1)
	y2 = int(y1)
	mask2 = ModifierType(mask1)
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return x2, y2, mask2, ret2
}
func (this0 *Window) GetPosition() (int, int) {
	var this1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_get_position(this1, &x1, &y1)
	var x2 int
	var y2 int
	x2 = int(x1)
	y2 = int(y1)
	return x2, y2
}
func (this0 *Window) GetRootCoords(x0 int, y0 int) (int, int) {
	var this1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	var root_x1 C.int32_t
	var root_y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	C.gdk_window_get_root_coords(this1, x1, y1, &root_x1, &root_y1)
	var root_x2 int
	var root_y2 int
	root_x2 = int(root_x1)
	root_y2 = int(root_y1)
	return root_x2, root_y2
}
func (this0 *Window) GetRootOrigin() (int, int) {
	var this1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_get_root_origin(this1, &x1, &y1)
	var x2 int
	var y2 int
	x2 = int(x1)
	y2 = int(y1)
	return x2, y2
}
func (this0 *Window) GetScreen() *Screen {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_screen(this1)
	var ret2 *Screen
	ret2 = (*Screen)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetSourceEvents(source0 InputSource) EventMask {
	var this1 *C.GdkWindow
	var source1 C.GdkInputSource
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	source1 = C.GdkInputSource(source0)
	ret1 := C.gdk_window_get_source_events(this1, source1)
	var ret2 EventMask
	ret2 = EventMask(ret1)
	return ret2
}
func (this0 *Window) GetState() WindowState {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_state(this1)
	var ret2 WindowState
	ret2 = WindowState(ret1)
	return ret2
}
func (this0 *Window) GetSupportMultidevice() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_support_multidevice(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) GetToplevel() *Window {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_toplevel(this1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetTypeHint() WindowTypeHint {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_type_hint(this1)
	var ret2 WindowTypeHint
	ret2 = WindowTypeHint(ret1)
	return ret2
}
func (this0 *Window) GetUpdateArea() *cairo.Region {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_update_area(this1)
	var ret2 *cairo.Region
	ret2 = (*cairo.Region)(cairo.RegionWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Window) GetUserData() unsafe.Pointer {
	var this1 *C.GdkWindow
	var data1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_get_user_data(this1, &data1)
	var data2 unsafe.Pointer
	data2 = data1
	return data2
}
func (this0 *Window) GetVisibleRegion() *cairo.Region {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_visible_region(this1)
	var ret2 *cairo.Region
	ret2 = (*cairo.Region)(cairo.RegionWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Window) GetVisual() *Visual {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_visual(this1)
	var ret2 *Visual
	ret2 = (*Visual)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Window) GetWidth() int {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Window) GetWindowType() WindowType {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_get_window_type(this1)
	var ret2 WindowType
	ret2 = WindowType(ret1)
	return ret2
}
func (this0 *Window) HasNative() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_has_native(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) Hide() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_hide(this1)
}
func (this0 *Window) Iconify() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_iconify(this1)
}
func (this0 *Window) InputShapeCombineRegion(shape_region0 *cairo.Region, offset_x0 int, offset_y0 int) {
	var this1 *C.GdkWindow
	var shape_region1 *C.cairoRegion
	var offset_x1 C.int32_t
	var offset_y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if shape_region0 != nil {
		shape_region1 = (*C.cairoRegion)(shape_region0.C)
	}
	offset_x1 = C.int32_t(offset_x0)
	offset_y1 = C.int32_t(offset_y0)
	C.gdk_window_input_shape_combine_region(this1, shape_region1, offset_x1, offset_y1)
}
func (this0 *Window) InvalidateMaybeRecurse(region0 *cairo.Region, child_func0 WindowChildFunc) {
	var this1 *C.GdkWindow
	var region1 *C.cairoRegion
	var child_func1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if region0 != nil {
		region1 = (*C.cairoRegion)(region0.C)
	}
	if child_func0 != nil {
		child_func1 = unsafe.Pointer(&child_func0)}
	C._gdk_window_invalidate_maybe_recurse(this1, region1, child_func1)
}
func (this0 *Window) InvalidateRect(rect0 *cairo.RectangleInt, invalidate_children0 bool) {
	var this1 *C.GdkWindow
	var rect1 *C.cairoRectangleInt
	var invalidate_children1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	rect1 = (*C.cairoRectangleInt)(unsafe.Pointer(rect0))
	invalidate_children1 = _GoBoolToCBool(invalidate_children0)
	C.gdk_window_invalidate_rect(this1, rect1, invalidate_children1)
}
func (this0 *Window) InvalidateRegion(region0 *cairo.Region, invalidate_children0 bool) {
	var this1 *C.GdkWindow
	var region1 *C.cairoRegion
	var invalidate_children1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if region0 != nil {
		region1 = (*C.cairoRegion)(region0.C)
	}
	invalidate_children1 = _GoBoolToCBool(invalidate_children0)
	C.gdk_window_invalidate_region(this1, region1, invalidate_children1)
}
func (this0 *Window) IsDestroyed() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_is_destroyed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) IsInputOnly() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_is_input_only(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) IsShaped() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_is_shaped(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) IsViewable() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_is_viewable(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) IsVisible() bool {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_is_visible(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) Lower() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_lower(this1)
}
func (this0 *Window) Maximize() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_maximize(this1)
}
func (this0 *Window) MergeChildInputShapes() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_merge_child_input_shapes(this1)
}
func (this0 *Window) MergeChildShapes() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_merge_child_shapes(this1)
}
func (this0 *Window) Move(x0 int, y0 int) {
	var this1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	C.gdk_window_move(this1, x1, y1)
}
func (this0 *Window) MoveRegion(region0 *cairo.Region, dx0 int, dy0 int) {
	var this1 *C.GdkWindow
	var region1 *C.cairoRegion
	var dx1 C.int32_t
	var dy1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if region0 != nil {
		region1 = (*C.cairoRegion)(region0.C)
	}
	dx1 = C.int32_t(dx0)
	dy1 = C.int32_t(dy0)
	C.gdk_window_move_region(this1, region1, dx1, dy1)
}
func (this0 *Window) MoveResize(x0 int, y0 int, width0 int, height0 int) {
	var this1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	C.gdk_window_move_resize(this1, x1, y1, width1, height1)
}
func (this0 *Window) PeekChildren() []*Window {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_window_peek_children(this1)
	var ret2 []*Window
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Window
		elt = (*Window)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkWindow)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Window) ProcessUpdates(update_children0 bool) {
	var this1 *C.GdkWindow
	var update_children1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	update_children1 = _GoBoolToCBool(update_children0)
	C.gdk_window_process_updates(this1, update_children1)
}
func (this0 *Window) Raise() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_raise(this1)
}
func (this0 *Window) RegisterDnd() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_register_dnd(this1)
}
func (this0 *Window) Reparent(new_parent0 WindowLike, x0 int, y0 int) {
	var this1 *C.GdkWindow
	var new_parent1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if new_parent0 != nil {
		new_parent1 = new_parent0.InheritedFromGdkWindow()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	C.gdk_window_reparent(this1, new_parent1, x1, y1)
}
func (this0 *Window) Resize(width0 int, height0 int) {
	var this1 *C.GdkWindow
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	C.gdk_window_resize(this1, width1, height1)
}
func (this0 *Window) Restack(sibling0 WindowLike, above0 bool) {
	var this1 *C.GdkWindow
	var sibling1 *C.GdkWindow
	var above1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if sibling0 != nil {
		sibling1 = sibling0.InheritedFromGdkWindow()
	}
	above1 = _GoBoolToCBool(above0)
	C.gdk_window_restack(this1, sibling1, above1)
}
func (this0 *Window) Scroll(dx0 int, dy0 int) {
	var this1 *C.GdkWindow
	var dx1 C.int32_t
	var dy1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	dx1 = C.int32_t(dx0)
	dy1 = C.int32_t(dy0)
	C.gdk_window_scroll(this1, dx1, dy1)
}
func (this0 *Window) SetAcceptFocus(accept_focus0 bool) {
	var this1 *C.GdkWindow
	var accept_focus1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	accept_focus1 = _GoBoolToCBool(accept_focus0)
	C.gdk_window_set_accept_focus(this1, accept_focus1)
}
func (this0 *Window) SetBackground(color0 *Color) {
	var this1 *C.GdkWindow
	var color1 *C.GdkColor
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	color1 = (*C.GdkColor)(unsafe.Pointer(color0))
	C.gdk_window_set_background(this1, color1)
}
func (this0 *Window) SetBackgroundPattern(pattern0 cairo.PatternLike) {
	var this1 *C.GdkWindow
	var pattern1 *C.cairoPattern
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if pattern0 != nil {
		pattern1 = (*C.cairoPattern)(pattern0.InheritedFromCairoPattern().C)
	}
	C.gdk_window_set_background_pattern(this1, pattern1)
}
func (this0 *Window) SetBackgroundRGBA(rgba0 *RGBA) {
	var this1 *C.GdkWindow
	var rgba1 *C.GdkRGBA
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	rgba1 = (*C.GdkRGBA)(unsafe.Pointer(rgba0))
	C.gdk_window_set_background_rgba(this1, rgba1)
}
func (this0 *Window) SetChildInputShapes() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_set_child_input_shapes(this1)
}
func (this0 *Window) SetChildShapes() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_set_child_shapes(this1)
}
func (this0 *Window) SetComposited(composited0 bool) {
	var this1 *C.GdkWindow
	var composited1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	composited1 = _GoBoolToCBool(composited0)
	C.gdk_window_set_composited(this1, composited1)
}
func (this0 *Window) SetCursor(cursor0 CursorLike) {
	var this1 *C.GdkWindow
	var cursor1 *C.GdkCursor
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if cursor0 != nil {
		cursor1 = cursor0.InheritedFromGdkCursor()
	}
	C.gdk_window_set_cursor(this1, cursor1)
}
func (this0 *Window) SetDecorations(decorations0 WMDecoration) {
	var this1 *C.GdkWindow
	var decorations1 C.GdkWMDecoration
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	decorations1 = C.GdkWMDecoration(decorations0)
	C.gdk_window_set_decorations(this1, decorations1)
}
func (this0 *Window) SetDeviceCursor(device0 DeviceLike, cursor0 CursorLike) {
	var this1 *C.GdkWindow
	var device1 *C.GdkDevice
	var cursor1 *C.GdkCursor
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	if cursor0 != nil {
		cursor1 = cursor0.InheritedFromGdkCursor()
	}
	C.gdk_window_set_device_cursor(this1, device1, cursor1)
}
func (this0 *Window) SetDeviceEvents(device0 DeviceLike, event_mask0 EventMask) {
	var this1 *C.GdkWindow
	var device1 *C.GdkDevice
	var event_mask1 C.GdkEventMask
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	event_mask1 = C.GdkEventMask(event_mask0)
	C.gdk_window_set_device_events(this1, device1, event_mask1)
}
func (this0 *Window) SetEvents(event_mask0 EventMask) {
	var this1 *C.GdkWindow
	var event_mask1 C.GdkEventMask
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	event_mask1 = C.GdkEventMask(event_mask0)
	C.gdk_window_set_events(this1, event_mask1)
}
func (this0 *Window) SetFocusOnMap(focus_on_map0 bool) {
	var this1 *C.GdkWindow
	var focus_on_map1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	focus_on_map1 = _GoBoolToCBool(focus_on_map0)
	C.gdk_window_set_focus_on_map(this1, focus_on_map1)
}
func (this0 *Window) SetFunctions(functions0 WMFunction) {
	var this1 *C.GdkWindow
	var functions1 C.GdkWMFunction
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	functions1 = C.GdkWMFunction(functions0)
	C.gdk_window_set_functions(this1, functions1)
}
func (this0 *Window) SetGeometryHints(geometry0 *Geometry, geom_mask0 WindowHints) {
	var this1 *C.GdkWindow
	var geometry1 *C.GdkGeometry
	var geom_mask1 C.GdkWindowHints
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	geometry1 = (*C.GdkGeometry)(unsafe.Pointer(geometry0))
	geom_mask1 = C.GdkWindowHints(geom_mask0)
	C.gdk_window_set_geometry_hints(this1, geometry1, geom_mask1)
}
func (this0 *Window) SetGroup(leader0 WindowLike) {
	var this1 *C.GdkWindow
	var leader1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if leader0 != nil {
		leader1 = leader0.InheritedFromGdkWindow()
	}
	C.gdk_window_set_group(this1, leader1)
}
func (this0 *Window) SetIconList(pixbufs0 []*gdkpixbuf.Pixbuf) {
	var this1 *C.GdkWindow
	var pixbufs1 *C.GList
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_set_icon_list(this1, pixbufs1)
}
func (this0 *Window) SetIconName(name0 string) {
	var this1 *C.GdkWindow
	var name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	C.gdk_window_set_icon_name(this1, name1)
}
func (this0 *Window) SetKeepAbove(setting0 bool) {
	var this1 *C.GdkWindow
	var setting1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	setting1 = _GoBoolToCBool(setting0)
	C.gdk_window_set_keep_above(this1, setting1)
}
func (this0 *Window) SetKeepBelow(setting0 bool) {
	var this1 *C.GdkWindow
	var setting1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	setting1 = _GoBoolToCBool(setting0)
	C.gdk_window_set_keep_below(this1, setting1)
}
func (this0 *Window) SetModalHint(modal0 bool) {
	var this1 *C.GdkWindow
	var modal1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	modal1 = _GoBoolToCBool(modal0)
	C.gdk_window_set_modal_hint(this1, modal1)
}
func (this0 *Window) SetOpacity(opacity0 float64) {
	var this1 *C.GdkWindow
	var opacity1 C.double
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	opacity1 = C.double(opacity0)
	C.gdk_window_set_opacity(this1, opacity1)
}
func (this0 *Window) SetOverrideRedirect(override_redirect0 bool) {
	var this1 *C.GdkWindow
	var override_redirect1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	override_redirect1 = _GoBoolToCBool(override_redirect0)
	C.gdk_window_set_override_redirect(this1, override_redirect1)
}
func (this0 *Window) SetRole(role0 string) {
	var this1 *C.GdkWindow
	var role1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	role1 = _GoStringToGString(role0)
	defer C.free(unsafe.Pointer(role1))
	C.gdk_window_set_role(this1, role1)
}
func (this0 *Window) SetSkipPagerHint(skips_pager0 bool) {
	var this1 *C.GdkWindow
	var skips_pager1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	skips_pager1 = _GoBoolToCBool(skips_pager0)
	C.gdk_window_set_skip_pager_hint(this1, skips_pager1)
}
func (this0 *Window) SetSkipTaskbarHint(skips_taskbar0 bool) {
	var this1 *C.GdkWindow
	var skips_taskbar1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	skips_taskbar1 = _GoBoolToCBool(skips_taskbar0)
	C.gdk_window_set_skip_taskbar_hint(this1, skips_taskbar1)
}
func (this0 *Window) SetSourceEvents(source0 InputSource, event_mask0 EventMask) {
	var this1 *C.GdkWindow
	var source1 C.GdkInputSource
	var event_mask1 C.GdkEventMask
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	source1 = C.GdkInputSource(source0)
	event_mask1 = C.GdkEventMask(event_mask0)
	C.gdk_window_set_source_events(this1, source1, event_mask1)
}
func (this0 *Window) SetStartupID(startup_id0 string) {
	var this1 *C.GdkWindow
	var startup_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	startup_id1 = _GoStringToGString(startup_id0)
	defer C.free(unsafe.Pointer(startup_id1))
	C.gdk_window_set_startup_id(this1, startup_id1)
}
func (this0 *Window) SetStaticGravities(use_static0 bool) bool {
	var this1 *C.GdkWindow
	var use_static1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	use_static1 = _GoBoolToCBool(use_static0)
	ret1 := C.gdk_window_set_static_gravities(this1, use_static1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Window) SetSupportMultidevice(support_multidevice0 bool) {
	var this1 *C.GdkWindow
	var support_multidevice1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	support_multidevice1 = _GoBoolToCBool(support_multidevice0)
	C.gdk_window_set_support_multidevice(this1, support_multidevice1)
}
func (this0 *Window) SetTitle(title0 string) {
	var this1 *C.GdkWindow
	var title1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	title1 = _GoStringToGString(title0)
	defer C.free(unsafe.Pointer(title1))
	C.gdk_window_set_title(this1, title1)
}
func (this0 *Window) SetTransientFor(parent0 WindowLike) {
	var this1 *C.GdkWindow
	var parent1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if parent0 != nil {
		parent1 = parent0.InheritedFromGdkWindow()
	}
	C.gdk_window_set_transient_for(this1, parent1)
}
func (this0 *Window) SetTypeHint(hint0 WindowTypeHint) {
	var this1 *C.GdkWindow
	var hint1 C.GdkWindowTypeHint
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	hint1 = C.GdkWindowTypeHint(hint0)
	C.gdk_window_set_type_hint(this1, hint1)
}
func (this0 *Window) SetUrgencyHint(urgent0 bool) {
	var this1 *C.GdkWindow
	var urgent1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	urgent1 = _GoBoolToCBool(urgent0)
	C.gdk_window_set_urgency_hint(this1, urgent1)
}
func (this0 *Window) SetUserData(user_data0 gobject.ObjectLike) {
	var this1 *C.GdkWindow
	var user_data1 *C.GObject
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if user_data0 != nil {
		user_data1 = user_data0.InheritedFromGObject()
	}
	C.gdk_window_set_user_data(this1, user_data1)
}
func (this0 *Window) ShapeCombineRegion(shape_region0 *cairo.Region, offset_x0 int, offset_y0 int) {
	var this1 *C.GdkWindow
	var shape_region1 *C.cairoRegion
	var offset_x1 C.int32_t
	var offset_y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	if shape_region0 != nil {
		shape_region1 = (*C.cairoRegion)(shape_region0.C)
	}
	offset_x1 = C.int32_t(offset_x0)
	offset_y1 = C.int32_t(offset_y0)
	C.gdk_window_shape_combine_region(this1, shape_region1, offset_x1, offset_y1)
}
func (this0 *Window) Show() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_show(this1)
}
func (this0 *Window) ShowUnraised() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_show_unraised(this1)
}
func (this0 *Window) Stick() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_stick(this1)
}
func (this0 *Window) ThawToplevelUpdatesLibgtkOnly() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_thaw_toplevel_updates_libgtk_only(this1)
}
func (this0 *Window) ThawUpdates() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_thaw_updates(this1)
}
func (this0 *Window) Unfullscreen() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_unfullscreen(this1)
}
func (this0 *Window) Unmaximize() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_unmaximize(this1)
}
func (this0 *Window) Unstick() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_unstick(this1)
}
func (this0 *Window) Withdraw() {
	var this1 *C.GdkWindow
	if this0 != nil {
		this1 = this0.InheritedFromGdkWindow()
	}
	C.gdk_window_withdraw(this1)
}
type WindowAttr struct {
	title0 *C.char
	EventMask int32
	X int32
	Y int32
	Width int32
	Height int32
	Wclass WindowWindowClass
	visual0 *C.GdkVisual
	WindowType WindowType
	_ [4]byte
	cursor0 *C.GdkCursor
	wmclass_name0 *C.char
	wmclass_class0 *C.char
	OverrideRedirect int32
	TypeHint WindowTypeHint
}
func (this0 *WindowAttr) Title() string {
	var title1 string
	title1 = C.GoString(this0.title0)
	return title1
}
func (this0 *WindowAttr) Visual() *Visual {
	var visual1 *Visual
	visual1 = (*Visual)(gobject.ObjectWrap(unsafe.Pointer(this0.visual0), true))
	return visual1
}
func (this0 *WindowAttr) Cursor() *Cursor {
	var cursor1 *Cursor
	cursor1 = (*Cursor)(gobject.ObjectWrap(unsafe.Pointer(this0.cursor0), true))
	return cursor1
}
func (this0 *WindowAttr) WMClassName() string {
	var wmclass_name1 string
	wmclass_name1 = C.GoString(this0.wmclass_name0)
	return wmclass_name1
}
func (this0 *WindowAttr) WMClassClass() string {
	var wmclass_class1 string
	wmclass_class1 = C.GoString(this0.wmclass_class0)
	return wmclass_class1
}
type WindowAttributesType C.uint32_t
const (
	WindowAttributesTypeTitle WindowAttributesType = 2
	WindowAttributesTypeX WindowAttributesType = 4
	WindowAttributesTypeY WindowAttributesType = 8
	WindowAttributesTypeCursor WindowAttributesType = 16
	WindowAttributesTypeVisual WindowAttributesType = 32
	WindowAttributesTypeWMClass WindowAttributesType = 64
	WindowAttributesTypeNoredir WindowAttributesType = 128
	WindowAttributesTypeTypeHint WindowAttributesType = 256
)
type WindowChildFunc func(window *Window) bool
//export _GdkWindowChildFunc_c_wrapper
func _GdkWindowChildFunc_c_wrapper(window0 unsafe.Pointer, user_data0 unsafe.Pointer) int32 {
	var window1 *Window
	var user_data1 WindowChildFunc
	window1 = (*Window)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkWindow)(window0)), true))
	user_data1 = *(*WindowChildFunc)(user_data0)
	ret1 := user_data1(window1)
	var ret2 C.int
	ret2 = _GoBoolToCBool(ret1)
	return (int32)(ret2)
}
//export _GdkWindowChildFunc_c_wrapper_once
func _GdkWindowChildFunc_c_wrapper_once(window0 unsafe.Pointer, user_data0 unsafe.Pointer) int32 {
	ret := _GdkWindowChildFunc_c_wrapper(window0, user_data0)
	gobject.Holder.Release(user_data0)
	return ret
}
type WindowEdge C.uint32_t
const (
	WindowEdgeNorthWest WindowEdge = 0
	WindowEdgeNorth WindowEdge = 1
	WindowEdgeNorthEast WindowEdge = 2
	WindowEdgeWest WindowEdge = 3
	WindowEdgeEast WindowEdge = 4
	WindowEdgeSouthWest WindowEdge = 5
	WindowEdgeSouth WindowEdge = 6
	WindowEdgeSouthEast WindowEdge = 7
)
type WindowHints C.uint32_t
const (
	WindowHintsPos WindowHints = 1
	WindowHintsMinSize WindowHints = 2
	WindowHintsMaxSize WindowHints = 4
	WindowHintsBaseSize WindowHints = 8
	WindowHintsAspect WindowHints = 16
	WindowHintsResizeInc WindowHints = 32
	WindowHintsWinGravity WindowHints = 64
	WindowHintsUserPos WindowHints = 128
	WindowHintsUserSize WindowHints = 256
)
type WindowRedirect struct {}
type WindowState C.uint32_t
const (
	WindowStateWithdrawn WindowState = 1
	WindowStateIconified WindowState = 2
	WindowStateMaximized WindowState = 4
	WindowStateSticky WindowState = 8
	WindowStateFullscreen WindowState = 16
	WindowStateAbove WindowState = 32
	WindowStateBelow WindowState = 64
)
type WindowType C.uint32_t
const (
	WindowTypeRoot WindowType = 0
	WindowTypeToplevel WindowType = 1
	WindowTypeChild WindowType = 2
	WindowTypeTemp WindowType = 3
	WindowTypeForeign WindowType = 4
	WindowTypeOffscreen WindowType = 5
)
type WindowTypeHint C.uint32_t
const (
	WindowTypeHintNormal WindowTypeHint = 0
	WindowTypeHintDialog WindowTypeHint = 1
	WindowTypeHintMenu WindowTypeHint = 2
	WindowTypeHintToolbar WindowTypeHint = 3
	WindowTypeHintSplashscreen WindowTypeHint = 4
	WindowTypeHintUtility WindowTypeHint = 5
	WindowTypeHintDock WindowTypeHint = 6
	WindowTypeHintDesktop WindowTypeHint = 7
	WindowTypeHintDropdownMenu WindowTypeHint = 8
	WindowTypeHintPopupMenu WindowTypeHint = 9
	WindowTypeHintTooltip WindowTypeHint = 10
	WindowTypeHintNotification WindowTypeHint = 11
	WindowTypeHintCombo WindowTypeHint = 12
	WindowTypeHintDnd WindowTypeHint = 13
)
type WindowWindowClass C.uint32_t
const (
	WindowWindowClassOutput WindowWindowClass = 0
	WindowWindowClassOnly WindowWindowClass = 1
)
// blacklisted: add_option_entries_libgtk_only (function)
// blacklisted: atom_intern (function)
// blacklisted: atom_intern_static_string (function)
func Beep() {
	C.gdk_beep()
}
func CairoCreate(window0 WindowLike) *cairo.Context {
	var window1 *C.GdkWindow
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_cairo_create(window1)
	var ret2 *cairo.Context
	ret2 = (*cairo.Context)(cairo.ContextWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func CairoGetClipRectangle(cr0 *cairo.Context) (cairo.RectangleInt, bool) {
	var cr1 *C.cairoContext
	var rect1 C.cairoRectangleInt
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	ret1 := C.gdk_cairo_get_clip_rectangle(cr1, &rect1)
	var rect2 cairo.RectangleInt
	var ret2 bool
	rect2 = *(*cairo.RectangleInt)(unsafe.Pointer(&rect1))
	ret2 = ret1 != 0
	return rect2, ret2
}
func CairoRectangle(cr0 *cairo.Context, rectangle0 *cairo.RectangleInt) {
	var cr1 *C.cairoContext
	var rectangle1 *C.cairoRectangleInt
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	rectangle1 = (*C.cairoRectangleInt)(unsafe.Pointer(rectangle0))
	C.gdk_cairo_rectangle(cr1, rectangle1)
}
func CairoRegion(cr0 *cairo.Context, region0 *cairo.Region) {
	var cr1 *C.cairoContext
	var region1 *C.cairoRegion
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	if region0 != nil {
		region1 = (*C.cairoRegion)(region0.C)
	}
	C.gdk_cairo_region(cr1, region1)
}
func CairoRegionCreateFromSurface(surface0 cairo.SurfaceLike) *cairo.Region {
	var surface1 *C.cairoSurface
	if surface0 != nil {
		surface1 = (*C.cairoSurface)(surface0.InheritedFromCairoSurface().C)
	}
	ret1 := C.gdk_cairo_region_create_from_surface(surface1)
	var ret2 *cairo.Region
	ret2 = (*cairo.Region)(cairo.RegionWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func CairoSetSourceColor(cr0 *cairo.Context, color0 *Color) {
	var cr1 *C.cairoContext
	var color1 *C.GdkColor
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	color1 = (*C.GdkColor)(unsafe.Pointer(color0))
	C.gdk_cairo_set_source_color(cr1, color1)
}
func CairoSetSourcePixbuf(cr0 *cairo.Context, pixbuf0 gdkpixbuf.PixbufLike, pixbuf_x0 float64, pixbuf_y0 float64) {
	var cr1 *C.cairoContext
	var pixbuf1 *C.GdkPixbuf
	var pixbuf_x1 C.double
	var pixbuf_y1 C.double
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	if pixbuf0 != nil {
		pixbuf1 = pixbuf0.InheritedFromGdkPixbuf()
	}
	pixbuf_x1 = C.double(pixbuf_x0)
	pixbuf_y1 = C.double(pixbuf_y0)
	C.gdk_cairo_set_source_pixbuf(cr1, pixbuf1, pixbuf_x1, pixbuf_y1)
}
func CairoSetSourceRGBA(cr0 *cairo.Context, rgba0 *RGBA) {
	var cr1 *C.cairoContext
	var rgba1 *C.GdkRGBA
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	rgba1 = (*C.GdkRGBA)(unsafe.Pointer(rgba0))
	C.gdk_cairo_set_source_rgba(cr1, rgba1)
}
func CairoSetSourceWindow(cr0 *cairo.Context, window0 WindowLike, x0 float64, y0 float64) {
	var cr1 *C.cairoContext
	var window1 *C.GdkWindow
	var x1 C.double
	var y1 C.double
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	x1 = C.double(x0)
	y1 = C.double(y0)
	C.gdk_cairo_set_source_window(cr1, window1, x1, y1)
}
// blacklisted: color_parse (function)
func DisableMultidevice() {
	C.gdk_disable_multidevice()
}
func DragAbort(context0 DragContextLike, time_0 int) {
	var context1 *C.GdkDragContext
	var time_1 C.uint32_t
	if context0 != nil {
		context1 = context0.InheritedFromGdkDragContext()
	}
	time_1 = C.uint32_t(time_0)
	C.gdk_drag_abort(context1, time_1)
}
func DragBegin(window0 WindowLike, targets0 []Atom) *DragContext {
	var window1 *C.GdkWindow
	var targets1 *C.GList
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_drag_begin(window1, targets1)
	var ret2 *DragContext
	ret2 = (*DragContext)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func DragBeginForDevice(window0 WindowLike, device0 DeviceLike, targets0 []Atom) *DragContext {
	var window1 *C.GdkWindow
	var device1 *C.GdkDevice
	var targets1 *C.GList
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	if device0 != nil {
		device1 = device0.InheritedFromGdkDevice()
	}
	ret1 := C.gdk_drag_begin_for_device(window1, device1, targets1)
	var ret2 *DragContext
	ret2 = (*DragContext)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func DragDrop(context0 DragContextLike, time_0 int) {
	var context1 *C.GdkDragContext
	var time_1 C.uint32_t
	if context0 != nil {
		context1 = context0.InheritedFromGdkDragContext()
	}
	time_1 = C.uint32_t(time_0)
	C.gdk_drag_drop(context1, time_1)
}
func DragDropSucceeded(context0 DragContextLike) bool {
	var context1 *C.GdkDragContext
	if context0 != nil {
		context1 = context0.InheritedFromGdkDragContext()
	}
	ret1 := C.gdk_drag_drop_succeeded(context1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func DragFindWindowForScreen(context0 DragContextLike, drag_window0 WindowLike, screen0 ScreenLike, x_root0 int, y_root0 int) (*Window, DragProtocol) {
	var context1 *C.GdkDragContext
	var drag_window1 *C.GdkWindow
	var screen1 *C.GdkScreen
	var x_root1 C.int32_t
	var y_root1 C.int32_t
	var dest_window1 *C.GdkWindow
	var protocol1 C.GdkDragProtocol
	if context0 != nil {
		context1 = context0.InheritedFromGdkDragContext()
	}
	if drag_window0 != nil {
		drag_window1 = drag_window0.InheritedFromGdkWindow()
	}
	if screen0 != nil {
		screen1 = screen0.InheritedFromGdkScreen()
	}
	x_root1 = C.int32_t(x_root0)
	y_root1 = C.int32_t(y_root0)
	C.gdk_drag_find_window_for_screen(context1, drag_window1, screen1, x_root1, y_root1, &dest_window1, &protocol1)
	var dest_window2 *Window
	var protocol2 DragProtocol
	dest_window2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(dest_window1), false))
	protocol2 = DragProtocol(protocol1)
	return dest_window2, protocol2
}
func DragGetSelection(context0 DragContextLike) Atom {
	var context1 *C.GdkDragContext
	if context0 != nil {
		context1 = context0.InheritedFromGdkDragContext()
	}
	ret1 := C.gdk_drag_get_selection(context1)
	var ret2 Atom
	ret2 = Atom{unsafe.Pointer(ret1)}
	return ret2
}
func DragMotion(context0 DragContextLike, dest_window0 WindowLike, protocol0 DragProtocol, x_root0 int, y_root0 int, suggested_action0 DragAction, possible_actions0 DragAction, time_0 int) bool {
	var context1 *C.GdkDragContext
	var dest_window1 *C.GdkWindow
	var protocol1 C.GdkDragProtocol
	var x_root1 C.int32_t
	var y_root1 C.int32_t
	var suggested_action1 C.GdkDragAction
	var possible_actions1 C.GdkDragAction
	var time_1 C.uint32_t
	if context0 != nil {
		context1 = context0.InheritedFromGdkDragContext()
	}
	if dest_window0 != nil {
		dest_window1 = dest_window0.InheritedFromGdkWindow()
	}
	protocol1 = C.GdkDragProtocol(protocol0)
	x_root1 = C.int32_t(x_root0)
	y_root1 = C.int32_t(y_root0)
	suggested_action1 = C.GdkDragAction(suggested_action0)
	possible_actions1 = C.GdkDragAction(possible_actions0)
	time_1 = C.uint32_t(time_0)
	ret1 := C.gdk_drag_motion(context1, dest_window1, protocol1, x_root1, y_root1, suggested_action1, possible_actions1, time_1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func DragStatus(context0 DragContextLike, action0 DragAction, time_0 int) {
	var context1 *C.GdkDragContext
	var action1 C.GdkDragAction
	var time_1 C.uint32_t
	if context0 != nil {
		context1 = context0.InheritedFromGdkDragContext()
	}
	action1 = C.GdkDragAction(action0)
	time_1 = C.uint32_t(time_0)
	C.gdk_drag_status(context1, action1, time_1)
}
func DropFinish(context0 DragContextLike, success0 bool, time_0 int) {
	var context1 *C.GdkDragContext
	var success1 C.int
	var time_1 C.uint32_t
	if context0 != nil {
		context1 = context0.InheritedFromGdkDragContext()
	}
	success1 = _GoBoolToCBool(success0)
	time_1 = C.uint32_t(time_0)
	C.gdk_drop_finish(context1, success1, time_1)
}
func DropReply(context0 DragContextLike, accepted0 bool, time_0 int) {
	var context1 *C.GdkDragContext
	var accepted1 C.int
	var time_1 C.uint32_t
	if context0 != nil {
		context1 = context0.InheritedFromGdkDragContext()
	}
	accepted1 = _GoBoolToCBool(accepted0)
	time_1 = C.uint32_t(time_0)
	C.gdk_drop_reply(context1, accepted1, time_1)
}
func ErrorTrapPop() int {
	ret1 := C.gdk_error_trap_pop()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func ErrorTrapPopIgnored() {
	C.gdk_error_trap_pop_ignored()
}
func ErrorTrapPush() {
	C.gdk_error_trap_push()
}
// blacklisted: event_get (function)
// blacklisted: event_handler_set (function)
// blacklisted: event_peek (function)
// blacklisted: event_request_motions (function)
func EventsPending() bool {
	ret1 := C.gdk_events_pending()
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func Flush() {
	C.gdk_flush()
}
func GetDefaultRootWindow() *Window {
	ret1 := C.gdk_get_default_root_window()
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func GetDisplay() string {
	ret1 := C.gdk_get_display()
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func GetDisplayArgName() string {
	ret1 := C.gdk_get_display_arg_name()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func GetProgramClass() string {
	ret1 := C.gdk_get_program_class()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func GetShowEvents() bool {
	ret1 := C.gdk_get_show_events()
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func Init(argv0 []string) []string {
	var argv1 **C.char
	var argc1 C.int32_t
	argv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*argv1)) * len(argv0))))
	defer C.free(unsafe.Pointer(argv1))
	for i, e := range argv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i]))
	}
	argc1 = C.int32_t(len(argv0))
	C.gdk_init(&argc1, &argv1)
	var argv2 []string
	argv2 = make([]string, argc1)
	for i := range argv2 {
		argv2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i])
	}
	return argv2
}
func InitCheck(argv0 []string) ([]string, bool) {
	var argv1 **C.char
	var argc1 C.int32_t
	argv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*argv1)) * len(argv0))))
	defer C.free(unsafe.Pointer(argv1))
	for i, e := range argv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i]))
	}
	argc1 = C.int32_t(len(argv0))
	ret1 := C.gdk_init_check(&argc1, &argv1)
	var argv2 []string
	var ret2 bool
	argv2 = make([]string, argc1)
	for i := range argv2 {
		argv2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i])
	}
	ret2 = ret1 != 0
	return argv2, ret2
}
func KeyboardGrab(window0 WindowLike, owner_events0 bool, time_0 int) GrabStatus {
	var window1 *C.GdkWindow
	var owner_events1 C.int
	var time_1 C.uint32_t
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	owner_events1 = _GoBoolToCBool(owner_events0)
	time_1 = C.uint32_t(time_0)
	ret1 := C.gdk_keyboard_grab(window1, owner_events1, time_1)
	var ret2 GrabStatus
	ret2 = GrabStatus(ret1)
	return ret2
}
func KeyboardUngrab(time_0 int) {
	var time_1 C.uint32_t
	time_1 = C.uint32_t(time_0)
	C.gdk_keyboard_ungrab(time_1)
}
func KeyvalConvertCase(symbol0 int) (int, int) {
	var symbol1 C.uint32_t
	var lower1 C.uint32_t
	var upper1 C.uint32_t
	symbol1 = C.uint32_t(symbol0)
	C.gdk_keyval_convert_case(symbol1, &lower1, &upper1)
	var lower2 int
	var upper2 int
	lower2 = int(lower1)
	upper2 = int(upper1)
	return lower2, upper2
}
func KeyvalFromName(keyval_name0 string) int {
	var keyval_name1 *C.char
	keyval_name1 = _GoStringToGString(keyval_name0)
	defer C.free(unsafe.Pointer(keyval_name1))
	ret1 := C.gdk_keyval_from_name(keyval_name1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func KeyvalIsLower(keyval0 int) bool {
	var keyval1 C.uint32_t
	keyval1 = C.uint32_t(keyval0)
	ret1 := C.gdk_keyval_is_lower(keyval1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func KeyvalIsUpper(keyval0 int) bool {
	var keyval1 C.uint32_t
	keyval1 = C.uint32_t(keyval0)
	ret1 := C.gdk_keyval_is_upper(keyval1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func KeyvalName(keyval0 int) string {
	var keyval1 C.uint32_t
	keyval1 = C.uint32_t(keyval0)
	ret1 := C.gdk_keyval_name(keyval1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func KeyvalToLower(keyval0 int) int {
	var keyval1 C.uint32_t
	keyval1 = C.uint32_t(keyval0)
	ret1 := C.gdk_keyval_to_lower(keyval1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func KeyvalToUnicode(keyval0 int) int {
	var keyval1 C.uint32_t
	keyval1 = C.uint32_t(keyval0)
	ret1 := C.gdk_keyval_to_unicode(keyval1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func KeyvalToUpper(keyval0 int) int {
	var keyval1 C.uint32_t
	keyval1 = C.uint32_t(keyval0)
	ret1 := C.gdk_keyval_to_upper(keyval1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func ListVisuals() []*Visual {
	ret1 := C.gdk_list_visuals()
	var ret2 []*Visual
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Visual
		elt = (*Visual)(gobject.ObjectWrap(unsafe.Pointer((*C.GdkVisual)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func NotifyStartupComplete() {
	C.gdk_notify_startup_complete()
}
func NotifyStartupCompleteWithID(startup_id0 string) {
	var startup_id1 *C.char
	startup_id1 = _GoStringToGString(startup_id0)
	defer C.free(unsafe.Pointer(startup_id1))
	C.gdk_notify_startup_complete_with_id(startup_id1)
}
func OffscreenWindowGetEmbedder(window0 WindowLike) *Window {
	var window1 *C.GdkWindow
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_offscreen_window_get_embedder(window1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func OffscreenWindowGetSurface(window0 WindowLike) *cairo.Surface {
	var window1 *C.GdkWindow
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	ret1 := C.gdk_offscreen_window_get_surface(window1)
	var ret2 *cairo.Surface
	ret2 = (*cairo.Surface)(cairo.SurfaceWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func OffscreenWindowSetEmbedder(window0 WindowLike, embedder0 WindowLike) {
	var window1 *C.GdkWindow
	var embedder1 *C.GdkWindow
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	if embedder0 != nil {
		embedder1 = embedder0.InheritedFromGdkWindow()
	}
	C.gdk_offscreen_window_set_embedder(window1, embedder1)
}
// blacklisted: pango_context_get (function)
// blacklisted: pango_context_get_for_screen (function)
func ParseArgs(argv0 []string) []string {
	var argv1 **C.char
	var argc1 C.int32_t
	argv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*argv1)) * len(argv0))))
	defer C.free(unsafe.Pointer(argv1))
	for i, e := range argv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i]))
	}
	argc1 = C.int32_t(len(argv0))
	C.gdk_parse_args(&argc1, &argv1)
	var argv2 []string
	argv2 = make([]string, argc1)
	for i := range argv2 {
		argv2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i])
	}
	return argv2
}
func PixbufGetFromSurface(surface0 cairo.SurfaceLike, src_x0 int, src_y0 int, width0 int, height0 int) *gdkpixbuf.Pixbuf {
	var surface1 *C.cairoSurface
	var src_x1 C.int32_t
	var src_y1 C.int32_t
	var width1 C.int32_t
	var height1 C.int32_t
	if surface0 != nil {
		surface1 = (*C.cairoSurface)(surface0.InheritedFromCairoSurface().C)
	}
	src_x1 = C.int32_t(src_x0)
	src_y1 = C.int32_t(src_y0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	ret1 := C.gdk_pixbuf_get_from_surface(surface1, src_x1, src_y1, width1, height1)
	var ret2 *gdkpixbuf.Pixbuf
	ret2 = (*gdkpixbuf.Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func PixbufGetFromWindow(window0 WindowLike, src_x0 int, src_y0 int, width0 int, height0 int) *gdkpixbuf.Pixbuf {
	var window1 *C.GdkWindow
	var src_x1 C.int32_t
	var src_y1 C.int32_t
	var width1 C.int32_t
	var height1 C.int32_t
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	src_x1 = C.int32_t(src_x0)
	src_y1 = C.int32_t(src_y0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	ret1 := C.gdk_pixbuf_get_from_window(window1, src_x1, src_y1, width1, height1)
	var ret2 *gdkpixbuf.Pixbuf
	ret2 = (*gdkpixbuf.Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func PointerGrab(window0 WindowLike, owner_events0 bool, event_mask0 EventMask, confine_to0 WindowLike, cursor0 CursorLike, time_0 int) GrabStatus {
	var window1 *C.GdkWindow
	var owner_events1 C.int
	var event_mask1 C.GdkEventMask
	var confine_to1 *C.GdkWindow
	var cursor1 *C.GdkCursor
	var time_1 C.uint32_t
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	owner_events1 = _GoBoolToCBool(owner_events0)
	event_mask1 = C.GdkEventMask(event_mask0)
	if confine_to0 != nil {
		confine_to1 = confine_to0.InheritedFromGdkWindow()
	}
	if cursor0 != nil {
		cursor1 = cursor0.InheritedFromGdkCursor()
	}
	time_1 = C.uint32_t(time_0)
	ret1 := C.gdk_pointer_grab(window1, owner_events1, event_mask1, confine_to1, cursor1, time_1)
	var ret2 GrabStatus
	ret2 = GrabStatus(ret1)
	return ret2
}
func PointerIsGrabbed() bool {
	ret1 := C.gdk_pointer_is_grabbed()
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func PointerUngrab(time_0 int) {
	var time_1 C.uint32_t
	time_1 = C.uint32_t(time_0)
	C.gdk_pointer_ungrab(time_1)
}
func PreParseLibgtkOnly() {
	C.gdk_pre_parse_libgtk_only()
}
func PropertyDelete(window0 WindowLike, property0 Atom) {
	var window1 *C.GdkWindow
	var property1 C.GdkAtom
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	property1 = *(*C.GdkAtom)(unsafe.Pointer(&property0))
	C.gdk_property_delete(window1, property1)
}
func PropertyGet(window0 WindowLike, property0 Atom, type0 Atom, offset0 uint64, length0 uint64, pdelete0 int) (Atom, int, []int, bool) {
	var window1 *C.GdkWindow
	var property1 C.GdkAtom
	var type1 C.GdkAtom
	var offset1 C.uint64_t
	var length1 C.uint64_t
	var pdelete1 C.int32_t
	var actual_property_type1 C.GdkAtom
	var actual_format1 C.int32_t
	var data1 *C.uint8_t
	var actual_length1 C.int32_t
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	property1 = *(*C.GdkAtom)(unsafe.Pointer(&property0))
	type1 = *(*C.GdkAtom)(unsafe.Pointer(&type0))
	offset1 = C.uint64_t(offset0)
	length1 = C.uint64_t(length0)
	pdelete1 = C.int32_t(pdelete0)
	ret1 := C.gdk_property_get(window1, property1, type1, offset1, length1, pdelete1, &actual_property_type1, &actual_format1, &actual_length1, &data1)
	var actual_property_type2 Atom
	var actual_format2 int
	var data2 []int
	var ret2 bool
	actual_property_type2 = Atom{unsafe.Pointer(actual_property_type1)}
	actual_format2 = int(actual_format1)
	data2 = make([]int, actual_length1)
	for i := range data2 {
		data2[i] = int((*(*[999999]C.uint8_t)(unsafe.Pointer(data1)))[i])
	}
	ret2 = ret1 != 0
	return actual_property_type2, actual_format2, data2, ret2
}
func QueryDepths() []int {
	var depths1 *C.int32_t
	var count1 C.int32_t
	C.gdk_query_depths(&depths1, &count1)
	var depths2 []int
	depths2 = make([]int, count1)
	for i := range depths2 {
		depths2[i] = int((*(*[999999]C.int32_t)(unsafe.Pointer(depths1)))[i])
	}
	return depths2
}
func QueryVisualTypes() []VisualType {
	var visual_types1 *C.GdkVisualType
	var count1 C.int32_t
	C.gdk_query_visual_types(&visual_types1, &count1)
	var visual_types2 []VisualType
	visual_types2 = make([]VisualType, count1)
	for i := range visual_types2 {
		visual_types2[i] = VisualType((*(*[999999]C.GdkVisualType)(unsafe.Pointer(visual_types1)))[i])
	}
	return visual_types2
}
func RectangleGetType() gobject.Type {
	ret1 := C.gdk_rectangle_get_type()
	var ret2 gobject.Type
	ret2 = gobject.Type(ret1)
	return ret2
}
func RectangleIntersect(src10 *cairo.RectangleInt, src20 *cairo.RectangleInt) (cairo.RectangleInt, bool) {
	var src11 *C.cairoRectangleInt
	var src21 *C.cairoRectangleInt
	var dest1 C.cairoRectangleInt
	src11 = (*C.cairoRectangleInt)(unsafe.Pointer(src10))
	src21 = (*C.cairoRectangleInt)(unsafe.Pointer(src20))
	ret1 := C.gdk_rectangle_intersect(src11, src21, &dest1)
	var dest2 cairo.RectangleInt
	var ret2 bool
	dest2 = *(*cairo.RectangleInt)(unsafe.Pointer(&dest1))
	ret2 = ret1 != 0
	return dest2, ret2
}
func RectangleUnion(src10 *cairo.RectangleInt, src20 *cairo.RectangleInt) cairo.RectangleInt {
	var src11 *C.cairoRectangleInt
	var src21 *C.cairoRectangleInt
	var dest1 C.cairoRectangleInt
	src11 = (*C.cairoRectangleInt)(unsafe.Pointer(src10))
	src21 = (*C.cairoRectangleInt)(unsafe.Pointer(src20))
	C.gdk_rectangle_union(src11, src21, &dest1)
	var dest2 cairo.RectangleInt
	dest2 = *(*cairo.RectangleInt)(unsafe.Pointer(&dest1))
	return dest2
}
func SelectionConvert(requestor0 WindowLike, selection0 Atom, target0 Atom, time_0 int) {
	var requestor1 *C.GdkWindow
	var selection1 C.GdkAtom
	var target1 C.GdkAtom
	var time_1 C.uint32_t
	if requestor0 != nil {
		requestor1 = requestor0.InheritedFromGdkWindow()
	}
	selection1 = *(*C.GdkAtom)(unsafe.Pointer(&selection0))
	target1 = *(*C.GdkAtom)(unsafe.Pointer(&target0))
	time_1 = C.uint32_t(time_0)
	C.gdk_selection_convert(requestor1, selection1, target1, time_1)
}
func SelectionOwnerGet(selection0 Atom) *Window {
	var selection1 C.GdkAtom
	selection1 = *(*C.GdkAtom)(unsafe.Pointer(&selection0))
	ret1 := C.gdk_selection_owner_get(selection1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func SelectionOwnerGetForDisplay(display0 DisplayLike, selection0 Atom) *Window {
	var display1 *C.GdkDisplay
	var selection1 C.GdkAtom
	if display0 != nil {
		display1 = display0.InheritedFromGdkDisplay()
	}
	selection1 = *(*C.GdkAtom)(unsafe.Pointer(&selection0))
	ret1 := C.gdk_selection_owner_get_for_display(display1, selection1)
	var ret2 *Window
	ret2 = (*Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func SelectionOwnerSet(owner0 WindowLike, selection0 Atom, time_0 int, send_event0 bool) bool {
	var owner1 *C.GdkWindow
	var selection1 C.GdkAtom
	var time_1 C.uint32_t
	var send_event1 C.int
	if owner0 != nil {
		owner1 = owner0.InheritedFromGdkWindow()
	}
	selection1 = *(*C.GdkAtom)(unsafe.Pointer(&selection0))
	time_1 = C.uint32_t(time_0)
	send_event1 = _GoBoolToCBool(send_event0)
	ret1 := C.gdk_selection_owner_set(owner1, selection1, time_1, send_event1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func SelectionOwnerSetForDisplay(display0 DisplayLike, owner0 WindowLike, selection0 Atom, time_0 int, send_event0 bool) bool {
	var display1 *C.GdkDisplay
	var owner1 *C.GdkWindow
	var selection1 C.GdkAtom
	var time_1 C.uint32_t
	var send_event1 C.int
	if display0 != nil {
		display1 = display0.InheritedFromGdkDisplay()
	}
	if owner0 != nil {
		owner1 = owner0.InheritedFromGdkWindow()
	}
	selection1 = *(*C.GdkAtom)(unsafe.Pointer(&selection0))
	time_1 = C.uint32_t(time_0)
	send_event1 = _GoBoolToCBool(send_event0)
	ret1 := C.gdk_selection_owner_set_for_display(display1, owner1, selection1, time_1, send_event1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func SelectionSendNotify(requestor0 WindowLike, selection0 Atom, target0 Atom, property0 Atom, time_0 int) {
	var requestor1 *C.GdkWindow
	var selection1 C.GdkAtom
	var target1 C.GdkAtom
	var property1 C.GdkAtom
	var time_1 C.uint32_t
	if requestor0 != nil {
		requestor1 = requestor0.InheritedFromGdkWindow()
	}
	selection1 = *(*C.GdkAtom)(unsafe.Pointer(&selection0))
	target1 = *(*C.GdkAtom)(unsafe.Pointer(&target0))
	property1 = *(*C.GdkAtom)(unsafe.Pointer(&property0))
	time_1 = C.uint32_t(time_0)
	C.gdk_selection_send_notify(requestor1, selection1, target1, property1, time_1)
}
func SelectionSendNotifyForDisplay(display0 DisplayLike, requestor0 WindowLike, selection0 Atom, target0 Atom, property0 Atom, time_0 int) {
	var display1 *C.GdkDisplay
	var requestor1 *C.GdkWindow
	var selection1 C.GdkAtom
	var target1 C.GdkAtom
	var property1 C.GdkAtom
	var time_1 C.uint32_t
	if display0 != nil {
		display1 = display0.InheritedFromGdkDisplay()
	}
	if requestor0 != nil {
		requestor1 = requestor0.InheritedFromGdkWindow()
	}
	selection1 = *(*C.GdkAtom)(unsafe.Pointer(&selection0))
	target1 = *(*C.GdkAtom)(unsafe.Pointer(&target0))
	property1 = *(*C.GdkAtom)(unsafe.Pointer(&property0))
	time_1 = C.uint32_t(time_0)
	C.gdk_selection_send_notify_for_display(display1, requestor1, selection1, target1, property1, time_1)
}
func SetDoubleClickTime(msec0 int) {
	var msec1 C.uint32_t
	msec1 = C.uint32_t(msec0)
	C.gdk_set_double_click_time(msec1)
}
func SetProgramClass(program_class0 string) {
	var program_class1 *C.char
	program_class1 = _GoStringToGString(program_class0)
	defer C.free(unsafe.Pointer(program_class1))
	C.gdk_set_program_class(program_class1)
}
func SetShowEvents(show_events0 bool) {
	var show_events1 C.int
	show_events1 = _GoBoolToCBool(show_events0)
	C.gdk_set_show_events(show_events1)
}
func SettingGet(name0 string, value0 *gobject.Value) bool {
	var name1 *C.char
	var value1 *C.GValue
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	ret1 := C.gdk_setting_get(name1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func SynthesizeWindowState(window0 WindowLike, unset_flags0 WindowState, set_flags0 WindowState) {
	var window1 *C.GdkWindow
	var unset_flags1 C.GdkWindowState
	var set_flags1 C.GdkWindowState
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	unset_flags1 = C.GdkWindowState(unset_flags0)
	set_flags1 = C.GdkWindowState(set_flags0)
	C.gdk_synthesize_window_state(window1, unset_flags1, set_flags1)
}
func TestRenderSync(window0 WindowLike) {
	var window1 *C.GdkWindow
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	C.gdk_test_render_sync(window1)
}
func TestSimulateButton(window0 WindowLike, x0 int, y0 int, button0 int, modifiers0 ModifierType, button_pressrelease0 EventType) bool {
	var window1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	var button1 C.uint32_t
	var modifiers1 C.GdkModifierType
	var button_pressrelease1 C.GdkEventType
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	button1 = C.uint32_t(button0)
	modifiers1 = C.GdkModifierType(modifiers0)
	button_pressrelease1 = C.GdkEventType(button_pressrelease0)
	ret1 := C.gdk_test_simulate_button(window1, x1, y1, button1, modifiers1, button_pressrelease1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func TestSimulateKey(window0 WindowLike, x0 int, y0 int, keyval0 int, modifiers0 ModifierType, key_pressrelease0 EventType) bool {
	var window1 *C.GdkWindow
	var x1 C.int32_t
	var y1 C.int32_t
	var keyval1 C.uint32_t
	var modifiers1 C.GdkModifierType
	var key_pressrelease1 C.GdkEventType
	if window0 != nil {
		window1 = window0.InheritedFromGdkWindow()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	keyval1 = C.uint32_t(keyval0)
	modifiers1 = C.GdkModifierType(modifiers0)
	key_pressrelease1 = C.GdkEventType(key_pressrelease0)
	ret1 := C.gdk_test_simulate_key(window1, x1, y1, keyval1, modifiers1, key_pressrelease1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func TextPropertyToUtf8ListForDisplay(display0 DisplayLike, encoding0 Atom, format0 int, text0 []int) ([]string, int) {
	var display1 *C.GdkDisplay
	var encoding1 C.GdkAtom
	var format1 C.int32_t
	var text1 *C.uint8_t
	var length1 C.int32_t
	var list1 **C.char
	if display0 != nil {
		display1 = display0.InheritedFromGdkDisplay()
	}
	encoding1 = *(*C.GdkAtom)(unsafe.Pointer(&encoding0))
	format1 = C.int32_t(format0)
	text1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*text1)) * len(text0))))
	defer C.free(unsafe.Pointer(text1))
	for i, e := range text0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(text1)))[i] = C.uint8_t(e)
	}
	length1 = C.int32_t(len(text0))
	ret1 := C.gdk_text_property_to_utf8_list_for_display(display1, encoding1, format1, text1, length1, &list1)
	var list2 []string
	var ret2 int
	for i := range list2 {
		list2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(list1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(list1)))[i]))
	}
	ret2 = int(ret1)
	return list2, ret2
}
// blacklisted: threads_add_idle (function)
// blacklisted: threads_add_timeout (function)
// blacklisted: threads_add_timeout_seconds (function)
func ThreadsEnter() {
	C.gdk_threads_enter()
}
func ThreadsInit() {
	C.gdk_threads_init()
}
func ThreadsLeave() {
	C.gdk_threads_leave()
}
func UnicodeToKeyval(wc0 int) int {
	var wc1 C.uint32_t
	wc1 = C.uint32_t(wc0)
	ret1 := C.gdk_unicode_to_keyval(wc1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func Utf8ToStringTarget(str0 string) string {
	var str1 *C.char
	str1 = _GoStringToGString(str0)
	defer C.free(unsafe.Pointer(str1))
	ret1 := C.gdk_utf8_to_string_target(str1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
