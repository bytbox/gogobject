package gdkpixbuf

/*
#include <stdlib.h>
#include <stdint.h>

typedef size_t GType;

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
typedef struct _GdkPixbufLoader GdkPixbufLoader;
typedef struct _GdkPixbufLoaderClass GdkPixbufLoaderClass;
typedef uint32_t GdkPixbufRotation;
typedef void* GdkPixbufSaveFunc;
extern void _GdkPixbufSaveFunc_c_wrapper();
extern void _GdkPixbufSaveFunc_c_wrapper_once();
typedef struct _GdkPixbufSimpleAnim GdkPixbufSimpleAnim;
typedef struct _GdkPixbufSimpleAnimClass GdkPixbufSimpleAnimClass;
typedef struct _GdkPixbufSimpleAnimIter GdkPixbufSimpleAnimIter;
typedef struct _GdkPixdata GdkPixdata;
typedef uint32_t GdkPixdataDumpType;
typedef uint32_t GdkPixdataType;
extern GdkPixbuf* gdk_pixbuf_new(GdkColorspace, int, int32_t, int32_t, int32_t);
extern GdkPixbuf* gdk_pixbuf_new_from_data(uint8_t*, GdkColorspace, int, int32_t, int32_t, int32_t, int32_t, GdkPixbufDestroyNotify, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static GdkPixbuf* _gdk_pixbuf_new_from_data(uint8_t* arg0, GdkColorspace arg1, int arg2, int32_t arg3, int32_t arg4, int32_t arg5, int32_t arg6, void* gofunc) {
	if (gofunc) {
		return gdk_pixbuf_new_from_data(arg0, arg1, arg2, arg3, arg4, arg5, arg6, _GdkPixbufDestroyNotify_c_wrapper_once, gofunc);
	} else {
		return gdk_pixbuf_new_from_data(arg0, arg1, arg2, arg3, arg4, arg5, arg6, 0, 0);
	}
}
extern GdkPixbuf* gdk_pixbuf_new_from_file(char*, GError**);
extern GdkPixbuf* gdk_pixbuf_new_from_file_at_scale(char*, int32_t, int32_t, int, GError**);
extern GdkPixbuf* gdk_pixbuf_new_from_file_at_size(char*, int32_t, int32_t, GError**);
extern GdkPixbuf* gdk_pixbuf_new_from_inline(int32_t, uint8_t*, int, GError**);
extern GdkPixbuf* gdk_pixbuf_new_from_stream(GInputStream*, GCancellable*, GError**);
extern GdkPixbuf* gdk_pixbuf_new_from_stream_at_scale(GInputStream*, int32_t, int32_t, int, GCancellable*, GError**);
extern GdkPixbuf* gdk_pixbuf_new_from_stream_finish(GAsyncResult*, GError**);
extern GdkPixbuf* gdk_pixbuf_new_from_xpm_data(char*);
extern GdkPixbufFormat* gdk_pixbuf_get_file_info(char*, int32_t*, int32_t*);
extern GSList* gdk_pixbuf_get_formats();
extern char* gdk_pixbuf_gettext(char*);
extern void gdk_pixbuf_new_from_stream_async(GInputStream*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _gdk_pixbuf_new_from_stream_async(GInputStream* arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		gdk_pixbuf_new_from_stream_async(arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		gdk_pixbuf_new_from_stream_async(arg0, arg1, 0, 0);
	}
}
extern void gdk_pixbuf_new_from_stream_at_scale_async(GInputStream*, int32_t, int32_t, int, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _gdk_pixbuf_new_from_stream_at_scale_async(GInputStream* arg0, int32_t arg1, int32_t arg2, int arg3, GCancellable* arg4, void* gofunc) {
	if (gofunc) {
		gdk_pixbuf_new_from_stream_at_scale_async(arg0, arg1, arg2, arg3, arg4, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		gdk_pixbuf_new_from_stream_at_scale_async(arg0, arg1, arg2, arg3, arg4, 0, 0);
	}
}
extern int gdk_pixbuf_save_to_stream_finish(GAsyncResult*, GError**);
extern GdkPixbuf* gdk_pixbuf_add_alpha(GdkPixbuf*, int, uint8_t, uint8_t, uint8_t);
extern GdkPixbuf* gdk_pixbuf_apply_embedded_orientation(GdkPixbuf*);
extern void gdk_pixbuf_composite(GdkPixbuf*, GdkPixbuf*, int32_t, int32_t, int32_t, int32_t, double, double, double, double, GdkInterpType, int32_t);
extern void gdk_pixbuf_composite_color(GdkPixbuf*, GdkPixbuf*, int32_t, int32_t, int32_t, int32_t, double, double, double, double, GdkInterpType, int32_t, int32_t, int32_t, int32_t, uint32_t, uint32_t);
extern GdkPixbuf* gdk_pixbuf_composite_color_simple(GdkPixbuf*, int32_t, int32_t, GdkInterpType, int32_t, int32_t, uint32_t, uint32_t);
extern GdkPixbuf* gdk_pixbuf_copy(GdkPixbuf*);
extern void gdk_pixbuf_copy_area(GdkPixbuf*, int32_t, int32_t, int32_t, int32_t, GdkPixbuf*, int32_t, int32_t);
extern void gdk_pixbuf_fill(GdkPixbuf*, uint32_t);
extern GdkPixbuf* gdk_pixbuf_flip(GdkPixbuf*, int);
extern int32_t gdk_pixbuf_get_bits_per_sample(GdkPixbuf*);
extern GdkColorspace gdk_pixbuf_get_colorspace(GdkPixbuf*);
extern int gdk_pixbuf_get_has_alpha(GdkPixbuf*);
extern int32_t gdk_pixbuf_get_height(GdkPixbuf*);
extern int32_t gdk_pixbuf_get_n_channels(GdkPixbuf*);
extern char* gdk_pixbuf_get_option(GdkPixbuf*, char*);
extern uint8_t* gdk_pixbuf_get_pixels(GdkPixbuf*);
extern int32_t gdk_pixbuf_get_rowstride(GdkPixbuf*);
extern int32_t gdk_pixbuf_get_width(GdkPixbuf*);
extern GdkPixbuf* gdk_pixbuf_new_subpixbuf(GdkPixbuf*, int32_t, int32_t, int32_t, int32_t);
extern GdkPixbuf* gdk_pixbuf_rotate_simple(GdkPixbuf*, GdkPixbufRotation);
extern void gdk_pixbuf_saturate_and_pixelate(GdkPixbuf*, GdkPixbuf*, float, int);
extern int gdk_pixbuf_save_to_bufferv(GdkPixbuf*, char*, uint64_t*, char*, char**, char**, GError**);
extern int gdk_pixbuf_save_to_callbackv(GdkPixbuf*, GdkPixbufSaveFunc, void*, char*, char**, char**, GError**);
#pragma GCC diagnostic ignored "-Wunused-function"
static int _gdk_pixbuf_save_to_callbackv(GdkPixbuf* this, void* gofunc, char* arg2, char** arg3, char** arg4, GError** arg5) {
	if (gofunc) {
		return gdk_pixbuf_save_to_callbackv(this, _GdkPixbufSaveFunc_c_wrapper, gofunc, arg2, arg3, arg4, arg5);
	} else {
		return gdk_pixbuf_save_to_callbackv(this, 0, 0, arg2, arg3, arg4, arg5);
	}
}
extern int gdk_pixbuf_savev(GdkPixbuf*, char*, char*, char**, char**, GError**);
extern void gdk_pixbuf_scale(GdkPixbuf*, GdkPixbuf*, int32_t, int32_t, int32_t, int32_t, double, double, double, double, GdkInterpType);
extern GdkPixbuf* gdk_pixbuf_scale_simple(GdkPixbuf*, int32_t, int32_t, GdkInterpType);
extern GType gdk_pixbuf_get_type();
extern GdkPixbufAnimation* gdk_pixbuf_animation_new_from_file(char*, GError**);
extern int32_t gdk_pixbuf_animation_get_height(GdkPixbufAnimation*);
extern GdkPixbufAnimationIter* gdk_pixbuf_animation_get_iter(GdkPixbufAnimation*, GTimeVal*);
extern GdkPixbuf* gdk_pixbuf_animation_get_static_image(GdkPixbufAnimation*);
extern int32_t gdk_pixbuf_animation_get_width(GdkPixbufAnimation*);
extern int gdk_pixbuf_animation_is_static_image(GdkPixbufAnimation*);
extern GType gdk_pixbuf_animation_get_type();
extern int gdk_pixbuf_animation_iter_advance(GdkPixbufAnimationIter*, GTimeVal*);
extern int32_t gdk_pixbuf_animation_iter_get_delay_time(GdkPixbufAnimationIter*);
extern GdkPixbuf* gdk_pixbuf_animation_iter_get_pixbuf(GdkPixbufAnimationIter*);
extern int gdk_pixbuf_animation_iter_on_currently_loading_frame(GdkPixbufAnimationIter*);
extern GType gdk_pixbuf_animation_iter_get_type();
extern GdkPixbufFormat* gdk_pixbuf_format_copy(GdkPixbufFormat*);
extern void gdk_pixbuf_format_free(GdkPixbufFormat*);
extern char* gdk_pixbuf_format_get_description(GdkPixbufFormat*);
extern char** gdk_pixbuf_format_get_extensions(GdkPixbufFormat*);
extern char* gdk_pixbuf_format_get_license(GdkPixbufFormat*);
extern char** gdk_pixbuf_format_get_mime_types(GdkPixbufFormat*);
extern char* gdk_pixbuf_format_get_name(GdkPixbufFormat*);
extern int gdk_pixbuf_format_is_disabled(GdkPixbufFormat*);
extern int gdk_pixbuf_format_is_scalable(GdkPixbufFormat*);
extern int gdk_pixbuf_format_is_writable(GdkPixbufFormat*);
extern void gdk_pixbuf_format_set_disabled(GdkPixbufFormat*, int);
extern GdkPixbufLoader* gdk_pixbuf_loader_new();
extern GdkPixbufLoader* gdk_pixbuf_loader_new_with_mime_type(char*, GError**);
extern GdkPixbufLoader* gdk_pixbuf_loader_new_with_type(char*, GError**);
extern int gdk_pixbuf_loader_close(GdkPixbufLoader*, GError**);
extern GdkPixbufAnimation* gdk_pixbuf_loader_get_animation(GdkPixbufLoader*);
extern GdkPixbufFormat* gdk_pixbuf_loader_get_format(GdkPixbufLoader*);
extern GdkPixbuf* gdk_pixbuf_loader_get_pixbuf(GdkPixbufLoader*);
extern void gdk_pixbuf_loader_set_size(GdkPixbufLoader*, int32_t, int32_t);
extern int gdk_pixbuf_loader_write(GdkPixbufLoader*, uint8_t*, uint64_t, GError**);
extern GType gdk_pixbuf_loader_get_type();
extern GdkPixbufSimpleAnim* gdk_pixbuf_simple_anim_new(int32_t, int32_t, float);
extern void gdk_pixbuf_simple_anim_add_frame(GdkPixbufSimpleAnim*, GdkPixbuf*);
extern int gdk_pixbuf_simple_anim_get_loop(GdkPixbufSimpleAnim*);
extern void gdk_pixbuf_simple_anim_set_loop(GdkPixbufSimpleAnim*, int);
extern GType gdk_pixbuf_simple_anim_get_type();
extern GType gdk_pixbuf_simple_anim_iter_get_type();
extern int gdk_pixdata_deserialize(GdkPixdata*, uint32_t, uint8_t*, GError**);
extern uint8_t* gdk_pixdata_serialize(GdkPixdata*, uint32_t*);
extern GString* gdk_pixdata_to_csource(GdkPixdata*, char*, GdkPixdataDumpType);
extern uint32_t gdk_pixbuf_error_quark();
struct _GdkPixbufFormat {};
struct _GdkPixbufLoaderClass { uint8_t _data[168]; };
struct _GdkPixbufSimpleAnimClass {};
struct _GdkPixdata { uint8_t _data[32]; };


extern GObject *g_object_ref_sink(GObject*);
extern void g_object_unref(GObject*);
extern void g_error_free(GError*);
extern void g_free(void*);

#cgo pkg-config: gdk-pixbuf-2.0
*/
import "C"
import "unsafe"
import "errors"

// package dependencies
import (
	"github.com/bytbox/gogobject/gobject-2.0"
	"github.com/bytbox/gogobject/gio-2.0"
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


//export _GdkPixbuf_go_callback_cleanup
func _GdkPixbuf_go_callback_cleanup(gofunc unsafe.Pointer) {
	gobject.Holder.Release(gofunc)
}


type Colorspace C.uint32_t
const (
	ColorspaceRGB Colorspace = 0
)
type InterpType C.uint32_t
const (
	InterpTypeNearest InterpType = 0
	InterpTypeTiles InterpType = 1
	InterpTypeBilinear InterpType = 2
	InterpTypeHyper InterpType = 3
)
const PixbufFeaturesH = 1
const PixbufMagicNumber = 1197763408
const PixbufMajor = 2
const PixbufMicro = 1
const PixbufMinor = 24
const PixbufVersion = "2.24.1"
const PixdataHeaderLength = 24
type PixbufLike interface {
	gobject.ObjectLike
	InheritedFromGdkPixbuf() *C.GdkPixbuf
}

type Pixbuf struct {
	gobject.Object
	gio.IconImpl
}

func ToPixbuf(objlike gobject.ObjectLike) *Pixbuf {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Pixbuf)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Pixbuf)(obj)
	}
	panic("cannot cast to Pixbuf")
}

func (this0 *Pixbuf) InheritedFromGdkPixbuf() *C.GdkPixbuf {
	if this0 == nil {
		return nil
	}
	return (*C.GdkPixbuf)(this0.C)
}

func (this0 *Pixbuf) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_pixbuf_get_type())
}

func PixbufGetType() gobject.Type {
	return (*Pixbuf)(nil).GetStaticType()
}
func NewPixbuf(colorspace0 Colorspace, has_alpha0 bool, bits_per_sample0 int, width0 int, height0 int) *Pixbuf {
	var colorspace1 C.GdkColorspace
	var has_alpha1 C.int
	var bits_per_sample1 C.int32_t
	var width1 C.int32_t
	var height1 C.int32_t
	colorspace1 = C.GdkColorspace(colorspace0)
	has_alpha1 = _GoBoolToCBool(has_alpha0)
	bits_per_sample1 = C.int32_t(bits_per_sample0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	ret1 := C.gdk_pixbuf_new(colorspace1, has_alpha1, bits_per_sample1, width1, height1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
// blacklisted: Pixbuf.new_from_data (method)
func NewPixbufFromFile(filename0 string) (*Pixbuf, error) {
	var filename1 *C.char
	var err1 *C.GError
	filename1 = _GoStringToGString(filename0)
	defer C.free(unsafe.Pointer(filename1))
	ret1 := C.gdk_pixbuf_new_from_file(filename1, &err1)
	var ret2 *Pixbuf
	var err2 error
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func NewPixbufFromFileAtScale(filename0 string, width0 int, height0 int, preserve_aspect_ratio0 bool) (*Pixbuf, error) {
	var filename1 *C.char
	var width1 C.int32_t
	var height1 C.int32_t
	var preserve_aspect_ratio1 C.int
	var err1 *C.GError
	filename1 = _GoStringToGString(filename0)
	defer C.free(unsafe.Pointer(filename1))
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	preserve_aspect_ratio1 = _GoBoolToCBool(preserve_aspect_ratio0)
	ret1 := C.gdk_pixbuf_new_from_file_at_scale(filename1, width1, height1, preserve_aspect_ratio1, &err1)
	var ret2 *Pixbuf
	var err2 error
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func NewPixbufFromFileAtSize(filename0 string, width0 int, height0 int) (*Pixbuf, error) {
	var filename1 *C.char
	var width1 C.int32_t
	var height1 C.int32_t
	var err1 *C.GError
	filename1 = _GoStringToGString(filename0)
	defer C.free(unsafe.Pointer(filename1))
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	ret1 := C.gdk_pixbuf_new_from_file_at_size(filename1, width1, height1, &err1)
	var ret2 *Pixbuf
	var err2 error
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: Pixbuf.new_from_inline (method)
// blacklisted: Pixbuf.new_from_stream (method)
// blacklisted: Pixbuf.new_from_stream_at_scale (method)
// blacklisted: Pixbuf.new_from_stream_finish (method)
func NewPixbufFromXpmData(data0 string) *Pixbuf {
	var data1 *C.char
	data1 = _GoStringToGString(data0)
	defer C.free(unsafe.Pointer(data1))
	ret1 := C.gdk_pixbuf_new_from_xpm_data(data1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
// blacklisted: Pixbuf.get_file_info (method)
func PixbufGetFormats() []PixbufFormat {
	ret1 := C.gdk_pixbuf_get_formats()
	var ret2 []PixbufFormat
	for iter := (*_GSList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt PixbufFormat
		elt = *(*PixbufFormat)(unsafe.Pointer((*C.GdkPixbufFormat)(iter.data)))
		ret2 = append(ret2, elt)
	}
	return ret2
}
// blacklisted: Pixbuf.gettext (method)
// blacklisted: Pixbuf.new_from_stream_async (method)
// blacklisted: Pixbuf.new_from_stream_at_scale_async (method)
// blacklisted: Pixbuf.save_to_stream_finish (method)
func (this0 *Pixbuf) AddAlpha(substitute_color0 bool, r0 int, g0 int, b0 int) *Pixbuf {
	var this1 *C.GdkPixbuf
	var substitute_color1 C.int
	var r1 C.uint8_t
	var g1 C.uint8_t
	var b1 C.uint8_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	substitute_color1 = _GoBoolToCBool(substitute_color0)
	r1 = C.uint8_t(r0)
	g1 = C.uint8_t(g0)
	b1 = C.uint8_t(b0)
	ret1 := C.gdk_pixbuf_add_alpha(this1, substitute_color1, r1, g1, b1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Pixbuf) ApplyEmbeddedOrientation() *Pixbuf {
	var this1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	ret1 := C.gdk_pixbuf_apply_embedded_orientation(this1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Pixbuf) Composite(dest0 PixbufLike, dest_x0 int, dest_y0 int, dest_width0 int, dest_height0 int, offset_x0 float64, offset_y0 float64, scale_x0 float64, scale_y0 float64, interp_type0 InterpType, overall_alpha0 int) {
	var this1 *C.GdkPixbuf
	var dest1 *C.GdkPixbuf
	var dest_x1 C.int32_t
	var dest_y1 C.int32_t
	var dest_width1 C.int32_t
	var dest_height1 C.int32_t
	var offset_x1 C.double
	var offset_y1 C.double
	var scale_x1 C.double
	var scale_y1 C.double
	var interp_type1 C.GdkInterpType
	var overall_alpha1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	if dest0 != nil {
		dest1 = dest0.InheritedFromGdkPixbuf()
	}
	dest_x1 = C.int32_t(dest_x0)
	dest_y1 = C.int32_t(dest_y0)
	dest_width1 = C.int32_t(dest_width0)
	dest_height1 = C.int32_t(dest_height0)
	offset_x1 = C.double(offset_x0)
	offset_y1 = C.double(offset_y0)
	scale_x1 = C.double(scale_x0)
	scale_y1 = C.double(scale_y0)
	interp_type1 = C.GdkInterpType(interp_type0)
	overall_alpha1 = C.int32_t(overall_alpha0)
	C.gdk_pixbuf_composite(this1, dest1, dest_x1, dest_y1, dest_width1, dest_height1, offset_x1, offset_y1, scale_x1, scale_y1, interp_type1, overall_alpha1)
}
func (this0 *Pixbuf) CompositeColor(dest0 PixbufLike, dest_x0 int, dest_y0 int, dest_width0 int, dest_height0 int, offset_x0 float64, offset_y0 float64, scale_x0 float64, scale_y0 float64, interp_type0 InterpType, overall_alpha0 int, check_x0 int, check_y0 int, check_size0 int, color10 int, color20 int) {
	var this1 *C.GdkPixbuf
	var dest1 *C.GdkPixbuf
	var dest_x1 C.int32_t
	var dest_y1 C.int32_t
	var dest_width1 C.int32_t
	var dest_height1 C.int32_t
	var offset_x1 C.double
	var offset_y1 C.double
	var scale_x1 C.double
	var scale_y1 C.double
	var interp_type1 C.GdkInterpType
	var overall_alpha1 C.int32_t
	var check_x1 C.int32_t
	var check_y1 C.int32_t
	var check_size1 C.int32_t
	var color11 C.uint32_t
	var color21 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	if dest0 != nil {
		dest1 = dest0.InheritedFromGdkPixbuf()
	}
	dest_x1 = C.int32_t(dest_x0)
	dest_y1 = C.int32_t(dest_y0)
	dest_width1 = C.int32_t(dest_width0)
	dest_height1 = C.int32_t(dest_height0)
	offset_x1 = C.double(offset_x0)
	offset_y1 = C.double(offset_y0)
	scale_x1 = C.double(scale_x0)
	scale_y1 = C.double(scale_y0)
	interp_type1 = C.GdkInterpType(interp_type0)
	overall_alpha1 = C.int32_t(overall_alpha0)
	check_x1 = C.int32_t(check_x0)
	check_y1 = C.int32_t(check_y0)
	check_size1 = C.int32_t(check_size0)
	color11 = C.uint32_t(color10)
	color21 = C.uint32_t(color20)
	C.gdk_pixbuf_composite_color(this1, dest1, dest_x1, dest_y1, dest_width1, dest_height1, offset_x1, offset_y1, scale_x1, scale_y1, interp_type1, overall_alpha1, check_x1, check_y1, check_size1, color11, color21)
}
func (this0 *Pixbuf) CompositeColorSimple(dest_width0 int, dest_height0 int, interp_type0 InterpType, overall_alpha0 int, check_size0 int, color10 int, color20 int) *Pixbuf {
	var this1 *C.GdkPixbuf
	var dest_width1 C.int32_t
	var dest_height1 C.int32_t
	var interp_type1 C.GdkInterpType
	var overall_alpha1 C.int32_t
	var check_size1 C.int32_t
	var color11 C.uint32_t
	var color21 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	dest_width1 = C.int32_t(dest_width0)
	dest_height1 = C.int32_t(dest_height0)
	interp_type1 = C.GdkInterpType(interp_type0)
	overall_alpha1 = C.int32_t(overall_alpha0)
	check_size1 = C.int32_t(check_size0)
	color11 = C.uint32_t(color10)
	color21 = C.uint32_t(color20)
	ret1 := C.gdk_pixbuf_composite_color_simple(this1, dest_width1, dest_height1, interp_type1, overall_alpha1, check_size1, color11, color21)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Pixbuf) Copy() *Pixbuf {
	var this1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	ret1 := C.gdk_pixbuf_copy(this1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Pixbuf) CopyArea(src_x0 int, src_y0 int, width0 int, height0 int, dest_pixbuf0 PixbufLike, dest_x0 int, dest_y0 int) {
	var this1 *C.GdkPixbuf
	var src_x1 C.int32_t
	var src_y1 C.int32_t
	var width1 C.int32_t
	var height1 C.int32_t
	var dest_pixbuf1 *C.GdkPixbuf
	var dest_x1 C.int32_t
	var dest_y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	src_x1 = C.int32_t(src_x0)
	src_y1 = C.int32_t(src_y0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	if dest_pixbuf0 != nil {
		dest_pixbuf1 = dest_pixbuf0.InheritedFromGdkPixbuf()
	}
	dest_x1 = C.int32_t(dest_x0)
	dest_y1 = C.int32_t(dest_y0)
	C.gdk_pixbuf_copy_area(this1, src_x1, src_y1, width1, height1, dest_pixbuf1, dest_x1, dest_y1)
}
func (this0 *Pixbuf) Fill(pixel0 int) {
	var this1 *C.GdkPixbuf
	var pixel1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	pixel1 = C.uint32_t(pixel0)
	C.gdk_pixbuf_fill(this1, pixel1)
}
func (this0 *Pixbuf) Flip(horizontal0 bool) *Pixbuf {
	var this1 *C.GdkPixbuf
	var horizontal1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	horizontal1 = _GoBoolToCBool(horizontal0)
	ret1 := C.gdk_pixbuf_flip(this1, horizontal1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Pixbuf) GetBitsPerSample() int {
	var this1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	ret1 := C.gdk_pixbuf_get_bits_per_sample(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Pixbuf) GetColorspace() Colorspace {
	var this1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	ret1 := C.gdk_pixbuf_get_colorspace(this1)
	var ret2 Colorspace
	ret2 = Colorspace(ret1)
	return ret2
}
func (this0 *Pixbuf) GetHasAlpha() bool {
	var this1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	ret1 := C.gdk_pixbuf_get_has_alpha(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Pixbuf) GetHeight() int {
	var this1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	ret1 := C.gdk_pixbuf_get_height(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Pixbuf) GetNChannels() int {
	var this1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	ret1 := C.gdk_pixbuf_get_n_channels(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Pixbuf) GetOption(key0 string) string {
	var this1 *C.GdkPixbuf
	var key1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	key1 = _GoStringToGString(key0)
	defer C.free(unsafe.Pointer(key1))
	ret1 := C.gdk_pixbuf_get_option(this1, key1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
// blacklisted: Pixbuf.get_pixels (method)
func (this0 *Pixbuf) GetRowstride() int {
	var this1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	ret1 := C.gdk_pixbuf_get_rowstride(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Pixbuf) GetWidth() int {
	var this1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	ret1 := C.gdk_pixbuf_get_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Pixbuf) NewSubpixbuf(src_x0 int, src_y0 int, width0 int, height0 int) *Pixbuf {
	var this1 *C.GdkPixbuf
	var src_x1 C.int32_t
	var src_y1 C.int32_t
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	src_x1 = C.int32_t(src_x0)
	src_y1 = C.int32_t(src_y0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	ret1 := C.gdk_pixbuf_new_subpixbuf(this1, src_x1, src_y1, width1, height1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Pixbuf) RotateSimple(angle0 PixbufRotation) *Pixbuf {
	var this1 *C.GdkPixbuf
	var angle1 C.GdkPixbufRotation
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	angle1 = C.GdkPixbufRotation(angle0)
	ret1 := C.gdk_pixbuf_rotate_simple(this1, angle1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Pixbuf) SaturateAndPixelate(dest0 PixbufLike, saturation0 float64, pixelate0 bool) {
	var this1 *C.GdkPixbuf
	var dest1 *C.GdkPixbuf
	var saturation1 C.float
	var pixelate1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	if dest0 != nil {
		dest1 = dest0.InheritedFromGdkPixbuf()
	}
	saturation1 = C.float(saturation0)
	pixelate1 = _GoBoolToCBool(pixelate0)
	C.gdk_pixbuf_saturate_and_pixelate(this1, dest1, saturation1, pixelate1)
}
// blacklisted: Pixbuf.save_to_bufferv (method)
// blacklisted: Pixbuf.save_to_callbackv (method)
// blacklisted: Pixbuf.savev (method)
func (this0 *Pixbuf) Scale(dest0 PixbufLike, dest_x0 int, dest_y0 int, dest_width0 int, dest_height0 int, offset_x0 float64, offset_y0 float64, scale_x0 float64, scale_y0 float64, interp_type0 InterpType) {
	var this1 *C.GdkPixbuf
	var dest1 *C.GdkPixbuf
	var dest_x1 C.int32_t
	var dest_y1 C.int32_t
	var dest_width1 C.int32_t
	var dest_height1 C.int32_t
	var offset_x1 C.double
	var offset_y1 C.double
	var scale_x1 C.double
	var scale_y1 C.double
	var interp_type1 C.GdkInterpType
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	if dest0 != nil {
		dest1 = dest0.InheritedFromGdkPixbuf()
	}
	dest_x1 = C.int32_t(dest_x0)
	dest_y1 = C.int32_t(dest_y0)
	dest_width1 = C.int32_t(dest_width0)
	dest_height1 = C.int32_t(dest_height0)
	offset_x1 = C.double(offset_x0)
	offset_y1 = C.double(offset_y0)
	scale_x1 = C.double(scale_x0)
	scale_y1 = C.double(scale_y0)
	interp_type1 = C.GdkInterpType(interp_type0)
	C.gdk_pixbuf_scale(this1, dest1, dest_x1, dest_y1, dest_width1, dest_height1, offset_x1, offset_y1, scale_x1, scale_y1, interp_type1)
}
func (this0 *Pixbuf) ScaleSimple(dest_width0 int, dest_height0 int, interp_type0 InterpType) *Pixbuf {
	var this1 *C.GdkPixbuf
	var dest_width1 C.int32_t
	var dest_height1 C.int32_t
	var interp_type1 C.GdkInterpType
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbuf()
	}
	dest_width1 = C.int32_t(dest_width0)
	dest_height1 = C.int32_t(dest_height0)
	interp_type1 = C.GdkInterpType(interp_type0)
	ret1 := C.gdk_pixbuf_scale_simple(this1, dest_width1, dest_height1, interp_type1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
type PixbufAlphaMode C.uint32_t
const (
	PixbufAlphaModeBilevel PixbufAlphaMode = 0
	PixbufAlphaModeFull PixbufAlphaMode = 1
)
type PixbufAnimationLike interface {
	gobject.ObjectLike
	InheritedFromGdkPixbufAnimation() *C.GdkPixbufAnimation
}

type PixbufAnimation struct {
	gobject.Object
	
}

func ToPixbufAnimation(objlike gobject.ObjectLike) *PixbufAnimation {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*PixbufAnimation)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*PixbufAnimation)(obj)
	}
	panic("cannot cast to PixbufAnimation")
}

func (this0 *PixbufAnimation) InheritedFromGdkPixbufAnimation() *C.GdkPixbufAnimation {
	if this0 == nil {
		return nil
	}
	return (*C.GdkPixbufAnimation)(this0.C)
}

func (this0 *PixbufAnimation) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_pixbuf_animation_get_type())
}

func PixbufAnimationGetType() gobject.Type {
	return (*PixbufAnimation)(nil).GetStaticType()
}
func NewPixbufAnimationFromFile(filename0 string) (*PixbufAnimation, error) {
	var filename1 *C.char
	var err1 *C.GError
	filename1 = _GoStringToGString(filename0)
	defer C.free(unsafe.Pointer(filename1))
	ret1 := C.gdk_pixbuf_animation_new_from_file(filename1, &err1)
	var ret2 *PixbufAnimation
	var err2 error
	ret2 = (*PixbufAnimation)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *PixbufAnimation) GetHeight() int {
	var this1 *C.GdkPixbufAnimation
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufAnimation()
	}
	ret1 := C.gdk_pixbuf_animation_get_height(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
// blacklisted: PixbufAnimation.get_iter (method)
func (this0 *PixbufAnimation) GetStaticImage() *Pixbuf {
	var this1 *C.GdkPixbufAnimation
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufAnimation()
	}
	ret1 := C.gdk_pixbuf_animation_get_static_image(this1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *PixbufAnimation) GetWidth() int {
	var this1 *C.GdkPixbufAnimation
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufAnimation()
	}
	ret1 := C.gdk_pixbuf_animation_get_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *PixbufAnimation) IsStaticImage() bool {
	var this1 *C.GdkPixbufAnimation
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufAnimation()
	}
	ret1 := C.gdk_pixbuf_animation_is_static_image(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type PixbufAnimationIterLike interface {
	gobject.ObjectLike
	InheritedFromGdkPixbufAnimationIter() *C.GdkPixbufAnimationIter
}

type PixbufAnimationIter struct {
	gobject.Object
	
}

func ToPixbufAnimationIter(objlike gobject.ObjectLike) *PixbufAnimationIter {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*PixbufAnimationIter)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*PixbufAnimationIter)(obj)
	}
	panic("cannot cast to PixbufAnimationIter")
}

func (this0 *PixbufAnimationIter) InheritedFromGdkPixbufAnimationIter() *C.GdkPixbufAnimationIter {
	if this0 == nil {
		return nil
	}
	return (*C.GdkPixbufAnimationIter)(this0.C)
}

func (this0 *PixbufAnimationIter) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_pixbuf_animation_iter_get_type())
}

func PixbufAnimationIterGetType() gobject.Type {
	return (*PixbufAnimationIter)(nil).GetStaticType()
}
// blacklisted: PixbufAnimationIter.advance (method)
func (this0 *PixbufAnimationIter) GetDelayTime() int {
	var this1 *C.GdkPixbufAnimationIter
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufAnimationIter()
	}
	ret1 := C.gdk_pixbuf_animation_iter_get_delay_time(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *PixbufAnimationIter) GetPixbuf() *Pixbuf {
	var this1 *C.GdkPixbufAnimationIter
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufAnimationIter()
	}
	ret1 := C.gdk_pixbuf_animation_iter_get_pixbuf(this1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *PixbufAnimationIter) OnCurrentlyLoadingFrame() bool {
	var this1 *C.GdkPixbufAnimationIter
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufAnimationIter()
	}
	ret1 := C.gdk_pixbuf_animation_iter_on_currently_loading_frame(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: PixbufDestroyNotify (callback)
type PixbufError C.uint32_t
const (
	PixbufErrorCorruptImage PixbufError = 0
	PixbufErrorInsufficientMemory PixbufError = 1
	PixbufErrorBadOption PixbufError = 2
	PixbufErrorUnknownType PixbufError = 3
	PixbufErrorUnsupportedOperation PixbufError = 4
	PixbufErrorFailed PixbufError = 5
)
type PixbufFormat struct {}
func (this0 *PixbufFormat) Copy() *PixbufFormat {
	var this1 *C.GdkPixbufFormat
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	ret1 := C.gdk_pixbuf_format_copy(this1)
	var ret2 *PixbufFormat
	ret2 = (*PixbufFormat)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *PixbufFormat) Free() {
	var this1 *C.GdkPixbufFormat
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	C.gdk_pixbuf_format_free(this1)
}
func (this0 *PixbufFormat) GetDescription() string {
	var this1 *C.GdkPixbufFormat
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	ret1 := C.gdk_pixbuf_format_get_description(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *PixbufFormat) GetExtensions() []string {
	var this1 *C.GdkPixbufFormat
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	ret1 := C.gdk_pixbuf_format_get_extensions(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *PixbufFormat) GetLicense() string {
	var this1 *C.GdkPixbufFormat
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	ret1 := C.gdk_pixbuf_format_get_license(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *PixbufFormat) GetMIMETypes() []string {
	var this1 *C.GdkPixbufFormat
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	ret1 := C.gdk_pixbuf_format_get_mime_types(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *PixbufFormat) GetName() string {
	var this1 *C.GdkPixbufFormat
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	ret1 := C.gdk_pixbuf_format_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *PixbufFormat) IsDisabled() bool {
	var this1 *C.GdkPixbufFormat
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	ret1 := C.gdk_pixbuf_format_is_disabled(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *PixbufFormat) IsScalable() bool {
	var this1 *C.GdkPixbufFormat
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	ret1 := C.gdk_pixbuf_format_is_scalable(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *PixbufFormat) IsWritable() bool {
	var this1 *C.GdkPixbufFormat
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	ret1 := C.gdk_pixbuf_format_is_writable(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *PixbufFormat) SetDisabled(disabled0 bool) {
	var this1 *C.GdkPixbufFormat
	var disabled1 C.int
	this1 = (*C.GdkPixbufFormat)(unsafe.Pointer(this0))
	disabled1 = _GoBoolToCBool(disabled0)
	C.gdk_pixbuf_format_set_disabled(this1, disabled1)
}
type PixbufLoaderLike interface {
	gobject.ObjectLike
	InheritedFromGdkPixbufLoader() *C.GdkPixbufLoader
}

type PixbufLoader struct {
	gobject.Object
	
}

func ToPixbufLoader(objlike gobject.ObjectLike) *PixbufLoader {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*PixbufLoader)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*PixbufLoader)(obj)
	}
	panic("cannot cast to PixbufLoader")
}

func (this0 *PixbufLoader) InheritedFromGdkPixbufLoader() *C.GdkPixbufLoader {
	if this0 == nil {
		return nil
	}
	return (*C.GdkPixbufLoader)(this0.C)
}

func (this0 *PixbufLoader) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_pixbuf_loader_get_type())
}

func PixbufLoaderGetType() gobject.Type {
	return (*PixbufLoader)(nil).GetStaticType()
}
func NewPixbufLoader() *PixbufLoader {
	ret1 := C.gdk_pixbuf_loader_new()
	var ret2 *PixbufLoader
	ret2 = (*PixbufLoader)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewPixbufLoaderWithMIMEType(mime_type0 string) (*PixbufLoader, error) {
	var mime_type1 *C.char
	var err1 *C.GError
	mime_type1 = _GoStringToGString(mime_type0)
	defer C.free(unsafe.Pointer(mime_type1))
	ret1 := C.gdk_pixbuf_loader_new_with_mime_type(mime_type1, &err1)
	var ret2 *PixbufLoader
	var err2 error
	ret2 = (*PixbufLoader)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func NewPixbufLoaderWithType(image_type0 string) (*PixbufLoader, error) {
	var image_type1 *C.char
	var err1 *C.GError
	image_type1 = _GoStringToGString(image_type0)
	defer C.free(unsafe.Pointer(image_type1))
	ret1 := C.gdk_pixbuf_loader_new_with_type(image_type1, &err1)
	var ret2 *PixbufLoader
	var err2 error
	ret2 = (*PixbufLoader)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *PixbufLoader) Close() (bool, error) {
	var this1 *C.GdkPixbufLoader
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufLoader()
	}
	ret1 := C.gdk_pixbuf_loader_close(this1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *PixbufLoader) GetAnimation() *PixbufAnimation {
	var this1 *C.GdkPixbufLoader
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufLoader()
	}
	ret1 := C.gdk_pixbuf_loader_get_animation(this1)
	var ret2 *PixbufAnimation
	ret2 = (*PixbufAnimation)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *PixbufLoader) GetFormat() *PixbufFormat {
	var this1 *C.GdkPixbufLoader
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufLoader()
	}
	ret1 := C.gdk_pixbuf_loader_get_format(this1)
	var ret2 *PixbufFormat
	ret2 = (*PixbufFormat)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *PixbufLoader) GetPixbuf() *Pixbuf {
	var this1 *C.GdkPixbufLoader
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufLoader()
	}
	ret1 := C.gdk_pixbuf_loader_get_pixbuf(this1)
	var ret2 *Pixbuf
	ret2 = (*Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *PixbufLoader) SetSize(width0 int, height0 int) {
	var this1 *C.GdkPixbufLoader
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufLoader()
	}
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	C.gdk_pixbuf_loader_set_size(this1, width1, height1)
}
func (this0 *PixbufLoader) Write(buf0 []int) (bool, error) {
	var this1 *C.GdkPixbufLoader
	var buf1 *C.uint8_t
	var count1 C.uint64_t
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufLoader()
	}
	buf1 = (*C.uint8_t)(C.malloc(C.size_t(int(unsafe.Sizeof(*buf1)) * len(buf0))))
	defer C.free(unsafe.Pointer(buf1))
	for i, e := range buf0 {
		(*(*[999999]C.uint8_t)(unsafe.Pointer(buf1)))[i] = C.uint8_t(e)
	}
	count1 = C.uint64_t(len(buf0))
	ret1 := C.gdk_pixbuf_loader_write(this1, buf1, count1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
type PixbufRotation C.uint32_t
const (
	PixbufRotationNone PixbufRotation = 0
	PixbufRotationCounterclockwise PixbufRotation = 90
	PixbufRotationUpsidedown PixbufRotation = 180
	PixbufRotationClockwise PixbufRotation = 270
)
// blacklisted: PixbufSaveFunc (callback)
type PixbufSimpleAnimLike interface {
	PixbufAnimationLike
	InheritedFromGdkPixbufSimpleAnim() *C.GdkPixbufSimpleAnim
}

type PixbufSimpleAnim struct {
	PixbufAnimation
	
}

func ToPixbufSimpleAnim(objlike gobject.ObjectLike) *PixbufSimpleAnim {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*PixbufSimpleAnim)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*PixbufSimpleAnim)(obj)
	}
	panic("cannot cast to PixbufSimpleAnim")
}

func (this0 *PixbufSimpleAnim) InheritedFromGdkPixbufSimpleAnim() *C.GdkPixbufSimpleAnim {
	if this0 == nil {
		return nil
	}
	return (*C.GdkPixbufSimpleAnim)(this0.C)
}

func (this0 *PixbufSimpleAnim) GetStaticType() gobject.Type {
	return gobject.Type(C.gdk_pixbuf_simple_anim_get_type())
}

func PixbufSimpleAnimGetType() gobject.Type {
	return (*PixbufSimpleAnim)(nil).GetStaticType()
}
func NewPixbufSimpleAnim(width0 int, height0 int, rate0 float64) *PixbufSimpleAnim {
	var width1 C.int32_t
	var height1 C.int32_t
	var rate1 C.float
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	rate1 = C.float(rate0)
	ret1 := C.gdk_pixbuf_simple_anim_new(width1, height1, rate1)
	var ret2 *PixbufSimpleAnim
	ret2 = (*PixbufSimpleAnim)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *PixbufSimpleAnim) AddFrame(pixbuf0 PixbufLike) {
	var this1 *C.GdkPixbufSimpleAnim
	var pixbuf1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufSimpleAnim()
	}
	if pixbuf0 != nil {
		pixbuf1 = pixbuf0.InheritedFromGdkPixbuf()
	}
	C.gdk_pixbuf_simple_anim_add_frame(this1, pixbuf1)
}
func (this0 *PixbufSimpleAnim) GetLoop() bool {
	var this1 *C.GdkPixbufSimpleAnim
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufSimpleAnim()
	}
	ret1 := C.gdk_pixbuf_simple_anim_get_loop(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *PixbufSimpleAnim) SetLoop(loop0 bool) {
	var this1 *C.GdkPixbufSimpleAnim
	var loop1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGdkPixbufSimpleAnim()
	}
	loop1 = _GoBoolToCBool(loop0)
	C.gdk_pixbuf_simple_anim_set_loop(this1, loop1)
}
// blacklisted: PixbufSimpleAnimIter (object)
type Pixdata struct {
	Magic uint32
	Length int32
	PixdataType uint32
	Rowstride uint32
	Width uint32
	Height uint32
	PixelData *uint8
}
type PixdataDumpType C.uint32_t
const (
	PixdataDumpTypePixdataStream PixdataDumpType = 0
	PixdataDumpTypePixdataStruct PixdataDumpType = 1
	PixdataDumpTypeMacros PixdataDumpType = 2
	PixdataDumpTypeGtypes PixdataDumpType = 0
	PixdataDumpTypeCtypes PixdataDumpType = 256
	PixdataDumpTypeStatic PixdataDumpType = 512
	PixdataDumpTypeConst PixdataDumpType = 1024
	PixdataDumpTypeRleDecoder PixdataDumpType = 65536
)
type PixdataType C.uint32_t
const (
	PixdataTypeColorTypeRGB PixdataType = 1
	PixdataTypeColorTypeRGBA PixdataType = 2
	PixdataTypeColorTypeMask PixdataType = 255
	PixdataTypeSampleWidth8 PixdataType = 65536
	PixdataTypeSampleWidthMask PixdataType = 983040
	PixdataTypeEncodingRaw PixdataType = 16777216
	PixdataTypeEncodingRle PixdataType = 33554432
	PixdataTypeEncodingMask PixdataType = 251658240
)
func PixbufErrorQuark() int {
	ret1 := C.gdk_pixbuf_error_quark()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}


func (this *Pixbuf) Save(filename0, type0 string, args ...string) error {
	if len(args) & 1 != 0 {
		panic("the number of arguments should be even (key/value pairs)")
	}

	nkeyvalues := len(args) / 2 + 1

	keys := make([]*C.char, 0, nkeyvalues)
	values := make([]*C.char, 0, nkeyvalues)
	for i := range args {
		cstr := C.CString(args[i])
		defer C.free(unsafe.Pointer(cstr))
		if i & 1 == 0 {
			keys = append(keys, cstr)
		} else {
			values = append(values, cstr)
		}
	}

	keys = append(keys, nil)
	values = append(values, nil)

	filename1 := C.CString(filename0)
	defer C.free(unsafe.Pointer(filename1))
	type1 := C.CString(type0)
	defer C.free(unsafe.Pointer(type1))

	var err1 *C.GError

	C.gdk_pixbuf_savev(this.InheritedFromGdkPixbuf(), filename1, type1,
		&keys[0], &values[0], &err1)
	var err2 error
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}

	return err2
}