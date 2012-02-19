package gio

/*
#include <stdlib.h>
#include <stdint.h>

typedef size_t GType;

extern void _Gio_go_callback_cleanup(void *gofunc);
static void _c_callback_cleanup(void *userdata)
{
	_Gio_go_callback_cleanup(userdata);
}
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
typedef struct _GAction GAction;
typedef struct _GActionEntry GActionEntry;
typedef struct _GActionGroup GActionGroup;
typedef struct _GActionGroupInterface GActionGroupInterface;
typedef struct _GActionInterface GActionInterface;
typedef struct _GAppInfo GAppInfo;
typedef uint32_t GAppInfoCreateFlags;
typedef struct _GAppInfoIface GAppInfoIface;
typedef struct _GAppLaunchContext GAppLaunchContext;
typedef struct _GAppLaunchContextClass GAppLaunchContextClass;
typedef struct _GAppLaunchContextPrivate GAppLaunchContextPrivate;
typedef struct _GApplication GApplication;
typedef struct _GApplicationClass GApplicationClass;
typedef struct _GApplicationCommandLine GApplicationCommandLine;
typedef struct _GApplicationCommandLineClass GApplicationCommandLineClass;
typedef struct _GApplicationCommandLinePrivate GApplicationCommandLinePrivate;
typedef uint32_t GApplicationFlags;
typedef struct _GApplicationPrivate GApplicationPrivate;
typedef uint32_t GAskPasswordFlags;
typedef struct _GAsyncInitable GAsyncInitable;
typedef struct _GAsyncInitableIface GAsyncInitableIface;
typedef void* GAsyncReadyCallback;
extern void _GAsyncReadyCallback_c_wrapper();
extern void _GAsyncReadyCallback_c_wrapper_once();
typedef struct _GAsyncResult GAsyncResult;
typedef struct _GAsyncResultIface GAsyncResultIface;
typedef struct _GBufferedInputStream GBufferedInputStream;
typedef struct _GBufferedInputStreamClass GBufferedInputStreamClass;
typedef struct _GBufferedInputStreamPrivate GBufferedInputStreamPrivate;
typedef struct _GBufferedOutputStream GBufferedOutputStream;
typedef struct _GBufferedOutputStreamClass GBufferedOutputStreamClass;
typedef struct _GBufferedOutputStreamPrivate GBufferedOutputStreamPrivate;
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
typedef struct _GCancellablePrivate GCancellablePrivate;
typedef void* GCancellableSourceFunc;
extern void _GCancellableSourceFunc_c_wrapper();
extern void _GCancellableSourceFunc_c_wrapper_once();
typedef struct _GCharsetConverter GCharsetConverter;
typedef struct _GCharsetConverterClass GCharsetConverterClass;
typedef struct _GConverter GConverter;
typedef uint32_t GConverterFlags;
typedef struct _GConverterIface GConverterIface;
typedef struct _GConverterInputStream GConverterInputStream;
typedef struct _GConverterInputStreamClass GConverterInputStreamClass;
typedef struct _GConverterInputStreamPrivate GConverterInputStreamPrivate;
typedef struct _GConverterOutputStream GConverterOutputStream;
typedef struct _GConverterOutputStreamClass GConverterOutputStreamClass;
typedef struct _GConverterOutputStreamPrivate GConverterOutputStreamPrivate;
typedef uint32_t GConverterResult;
typedef struct _GCredentials GCredentials;
typedef struct _GCredentialsClass GCredentialsClass;
typedef uint32_t GCredentialsType;
typedef struct _GDBusAnnotationInfo GDBusAnnotationInfo;
typedef struct _GDBusArgInfo GDBusArgInfo;
typedef struct _GDBusAuthObserver GDBusAuthObserver;
typedef uint32_t GDBusCallFlags;
typedef uint32_t GDBusCapabilityFlags;
typedef struct _GDBusConnection GDBusConnection;
typedef uint32_t GDBusConnectionFlags;
typedef uint32_t GDBusError;
typedef struct _GDBusErrorEntry GDBusErrorEntry;
typedef struct _GDBusInterface GDBusInterface;
typedef void* GDBusInterfaceGetPropertyFunc;
extern void _GDBusInterfaceGetPropertyFunc_c_wrapper();
extern void _GDBusInterfaceGetPropertyFunc_c_wrapper_once();
typedef struct _GDBusInterfaceIface GDBusInterfaceIface;
typedef struct _GDBusInterfaceInfo GDBusInterfaceInfo;
typedef void* GDBusInterfaceMethodCallFunc;
extern void _GDBusInterfaceMethodCallFunc_c_wrapper();
extern void _GDBusInterfaceMethodCallFunc_c_wrapper_once();
typedef void* GDBusInterfaceSetPropertyFunc;
extern void _GDBusInterfaceSetPropertyFunc_c_wrapper();
extern void _GDBusInterfaceSetPropertyFunc_c_wrapper_once();
typedef struct _GDBusInterfaceSkeleton GDBusInterfaceSkeleton;
typedef struct _GDBusInterfaceSkeletonClass GDBusInterfaceSkeletonClass;
typedef uint32_t GDBusInterfaceSkeletonFlags;
typedef struct _GDBusInterfaceSkeletonPrivate GDBusInterfaceSkeletonPrivate;
typedef struct _GDBusInterfaceVTable GDBusInterfaceVTable;
typedef struct _GDBusMessage GDBusMessage;
typedef uint32_t GDBusMessageByteOrder;
typedef void* GDBusMessageFilterFunction;
extern void _GDBusMessageFilterFunction_c_wrapper();
extern void _GDBusMessageFilterFunction_c_wrapper_once();
typedef uint32_t GDBusMessageFlags;
typedef uint32_t GDBusMessageHeaderField;
typedef uint32_t GDBusMessageType;
typedef struct _GDBusMethodInfo GDBusMethodInfo;
typedef struct _GDBusMethodInvocation GDBusMethodInvocation;
typedef struct _GDBusNodeInfo GDBusNodeInfo;
typedef struct _GDBusObject GDBusObject;
typedef struct _GDBusObjectIface GDBusObjectIface;
typedef struct _GDBusObjectManager GDBusObjectManager;
typedef struct _GDBusObjectManagerClient GDBusObjectManagerClient;
typedef struct _GDBusObjectManagerClientClass GDBusObjectManagerClientClass;
typedef uint32_t GDBusObjectManagerClientFlags;
typedef struct _GDBusObjectManagerClientPrivate GDBusObjectManagerClientPrivate;
typedef struct _GDBusObjectManagerIface GDBusObjectManagerIface;
typedef struct _GDBusObjectManagerServer GDBusObjectManagerServer;
typedef struct _GDBusObjectManagerServerClass GDBusObjectManagerServerClass;
typedef struct _GDBusObjectManagerServerPrivate GDBusObjectManagerServerPrivate;
typedef struct _GDBusObjectProxy GDBusObjectProxy;
typedef struct _GDBusObjectProxyClass GDBusObjectProxyClass;
typedef struct _GDBusObjectProxyPrivate GDBusObjectProxyPrivate;
typedef struct _GDBusObjectSkeleton GDBusObjectSkeleton;
typedef struct _GDBusObjectSkeletonClass GDBusObjectSkeletonClass;
typedef struct _GDBusObjectSkeletonPrivate GDBusObjectSkeletonPrivate;
typedef struct _GDBusPropertyInfo GDBusPropertyInfo;
typedef uint32_t GDBusPropertyInfoFlags;
typedef struct _GDBusProxy GDBusProxy;
typedef struct _GDBusProxyClass GDBusProxyClass;
typedef uint32_t GDBusProxyFlags;
typedef struct _GDBusProxyPrivate GDBusProxyPrivate;
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
typedef void* GDBusSubtreeDispatchFunc;
extern void _GDBusSubtreeDispatchFunc_c_wrapper();
extern void _GDBusSubtreeDispatchFunc_c_wrapper_once();
typedef uint32_t GDBusSubtreeFlags;
typedef void* GDBusSubtreeIntrospectFunc;
extern void _GDBusSubtreeIntrospectFunc_c_wrapper();
extern void _GDBusSubtreeIntrospectFunc_c_wrapper_once();
typedef struct _GDBusSubtreeVTable GDBusSubtreeVTable;
typedef struct _GDataInputStream GDataInputStream;
typedef struct _GDataInputStreamClass GDataInputStreamClass;
typedef struct _GDataInputStreamPrivate GDataInputStreamPrivate;
typedef struct _GDataOutputStream GDataOutputStream;
typedef struct _GDataOutputStreamClass GDataOutputStreamClass;
typedef struct _GDataOutputStreamPrivate GDataOutputStreamPrivate;
typedef uint32_t GDataStreamByteOrder;
typedef uint32_t GDataStreamNewlineType;
typedef struct _GDesktopAppInfo GDesktopAppInfo;
typedef struct _GDesktopAppInfoClass GDesktopAppInfoClass;
typedef struct _GDesktopAppInfoLookup GDesktopAppInfoLookup;
typedef struct _GDesktopAppInfoLookupIface GDesktopAppInfoLookupIface;
typedef void* GDesktopAppLaunchCallback;
extern void _GDesktopAppLaunchCallback_c_wrapper();
extern void _GDesktopAppLaunchCallback_c_wrapper_once();
typedef struct _GDrive GDrive;
typedef struct _GDriveIface GDriveIface;
typedef uint32_t GDriveStartFlags;
typedef uint32_t GDriveStartStopType;
typedef struct _GEmblem GEmblem;
typedef struct _GEmblemClass GEmblemClass;
typedef uint32_t GEmblemOrigin;
typedef struct _GEmblemedIcon GEmblemedIcon;
typedef struct _GEmblemedIconClass GEmblemedIconClass;
typedef struct _GEmblemedIconPrivate GEmblemedIconPrivate;
typedef struct _GFile GFile;
typedef struct _GFileAttributeInfo GFileAttributeInfo;
typedef uint32_t GFileAttributeInfoFlags;
typedef struct _GFileAttributeInfoList GFileAttributeInfoList;
typedef struct _GFileAttributeMatcher GFileAttributeMatcher;
typedef uint32_t GFileAttributeStatus;
typedef uint32_t GFileAttributeType;
typedef uint32_t GFileCopyFlags;
typedef uint32_t GFileCreateFlags;
typedef struct _GFileDescriptorBased GFileDescriptorBased;
typedef struct _GFileDescriptorBasedIface GFileDescriptorBasedIface;
typedef struct _GFileEnumerator GFileEnumerator;
typedef struct _GFileEnumeratorClass GFileEnumeratorClass;
typedef struct _GFileEnumeratorPrivate GFileEnumeratorPrivate;
typedef struct _GFileIOStream GFileIOStream;
typedef struct _GFileIOStreamClass GFileIOStreamClass;
typedef struct _GFileIOStreamPrivate GFileIOStreamPrivate;
typedef struct _GFileIcon GFileIcon;
typedef struct _GFileIconClass GFileIconClass;
typedef struct _GFileIface GFileIface;
typedef struct _GFileInfo GFileInfo;
typedef struct _GFileInfoClass GFileInfoClass;
typedef struct _GFileInputStream GFileInputStream;
typedef struct _GFileInputStreamClass GFileInputStreamClass;
typedef struct _GFileInputStreamPrivate GFileInputStreamPrivate;
typedef struct _GFileMonitor GFileMonitor;
typedef struct _GFileMonitorClass GFileMonitorClass;
typedef uint32_t GFileMonitorEvent;
typedef uint32_t GFileMonitorFlags;
typedef struct _GFileMonitorPrivate GFileMonitorPrivate;
typedef struct _GFileOutputStream GFileOutputStream;
typedef struct _GFileOutputStreamClass GFileOutputStreamClass;
typedef struct _GFileOutputStreamPrivate GFileOutputStreamPrivate;
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
typedef uint32_t GFilesystemPreviewType;
typedef struct _GFilterInputStream GFilterInputStream;
typedef struct _GFilterInputStreamClass GFilterInputStreamClass;
typedef struct _GFilterOutputStream GFilterOutputStream;
typedef struct _GFilterOutputStreamClass GFilterOutputStreamClass;
typedef uint32_t GIOErrorEnum;
typedef struct _GIOExtension GIOExtension;
typedef struct _GIOExtensionPoint GIOExtensionPoint;
typedef struct _GIOModule GIOModule;
typedef struct _GIOModuleClass GIOModuleClass;
typedef struct _GIOModuleScope GIOModuleScope;
typedef uint32_t GIOModuleScopeFlags;
typedef struct _GIOSchedulerJob GIOSchedulerJob;
typedef void* GIOSchedulerJobFunc;
extern void _GIOSchedulerJobFunc_c_wrapper();
extern void _GIOSchedulerJobFunc_c_wrapper_once();
typedef struct _GIOStream GIOStream;
typedef struct _GIOStreamAdapter GIOStreamAdapter;
typedef struct _GIOStreamClass GIOStreamClass;
typedef struct _GIOStreamPrivate GIOStreamPrivate;
typedef uint32_t GIOStreamSpliceFlags;
typedef struct _GIcon GIcon;
typedef struct _GIconIface GIconIface;
typedef struct _GInetAddress GInetAddress;
typedef struct _GInetAddressClass GInetAddressClass;
typedef struct _GInetAddressPrivate GInetAddressPrivate;
typedef struct _GInetSocketAddress GInetSocketAddress;
typedef struct _GInetSocketAddressClass GInetSocketAddressClass;
typedef struct _GInetSocketAddressPrivate GInetSocketAddressPrivate;
typedef struct _GInitable GInitable;
typedef struct _GInitableIface GInitableIface;
typedef struct _GInputStream GInputStream;
typedef struct _GInputStreamClass GInputStreamClass;
typedef struct _GInputStreamPrivate GInputStreamPrivate;
typedef struct _GInputVector GInputVector;
typedef struct _GLoadableIcon GLoadableIcon;
typedef struct _GLoadableIconIface GLoadableIconIface;
typedef struct _GMemoryInputStream GMemoryInputStream;
typedef struct _GMemoryInputStreamClass GMemoryInputStreamClass;
typedef struct _GMemoryInputStreamPrivate GMemoryInputStreamPrivate;
typedef struct _GMemoryOutputStream GMemoryOutputStream;
typedef struct _GMemoryOutputStreamClass GMemoryOutputStreamClass;
typedef struct _GMemoryOutputStreamPrivate GMemoryOutputStreamPrivate;
typedef struct _GMount GMount;
typedef struct _GMountIface GMountIface;
typedef uint32_t GMountMountFlags;
typedef struct _GMountOperation GMountOperation;
typedef struct _GMountOperationClass GMountOperationClass;
typedef struct _GMountOperationPrivate GMountOperationPrivate;
typedef uint32_t GMountOperationResult;
typedef uint32_t GMountUnmountFlags;
typedef struct _GNativeVolumeMonitor GNativeVolumeMonitor;
typedef struct _GNativeVolumeMonitorClass GNativeVolumeMonitorClass;
typedef struct _GNetworkAddress GNetworkAddress;
typedef struct _GNetworkAddressClass GNetworkAddressClass;
typedef struct _GNetworkAddressPrivate GNetworkAddressPrivate;
typedef struct _GNetworkService GNetworkService;
typedef struct _GNetworkServiceClass GNetworkServiceClass;
typedef struct _GNetworkServicePrivate GNetworkServicePrivate;
typedef struct _GOutputStream GOutputStream;
typedef struct _GOutputStreamClass GOutputStreamClass;
typedef struct _GOutputStreamPrivate GOutputStreamPrivate;
typedef uint32_t GOutputStreamSpliceFlags;
typedef struct _GOutputVector GOutputVector;
typedef uint32_t GPasswordSave;
typedef struct _GPermission GPermission;
typedef struct _GPermissionClass GPermissionClass;
typedef struct _GPermissionPrivate GPermissionPrivate;
typedef struct _GPollableInputStream GPollableInputStream;
typedef struct _GPollableInputStreamInterface GPollableInputStreamInterface;
typedef struct _GPollableOutputStream GPollableOutputStream;
typedef struct _GPollableOutputStreamInterface GPollableOutputStreamInterface;
typedef void* GPollableSourceFunc;
extern void _GPollableSourceFunc_c_wrapper();
extern void _GPollableSourceFunc_c_wrapper_once();
typedef struct _GProxy GProxy;
typedef struct _GProxyAddress GProxyAddress;
typedef struct _GProxyAddressClass GProxyAddressClass;
typedef struct _GProxyAddressEnumerator GProxyAddressEnumerator;
typedef struct _GProxyAddressEnumeratorClass GProxyAddressEnumeratorClass;
typedef struct _GProxyAddressEnumeratorPrivate GProxyAddressEnumeratorPrivate;
typedef struct _GProxyAddressPrivate GProxyAddressPrivate;
typedef struct _GProxyInterface GProxyInterface;
typedef struct _GProxyResolver GProxyResolver;
typedef struct _GProxyResolverInterface GProxyResolverInterface;
typedef struct _GResolver GResolver;
typedef struct _GResolverClass GResolverClass;
typedef uint32_t GResolverError;
typedef struct _GResolverPrivate GResolverPrivate;
typedef struct _GSeekable GSeekable;
typedef struct _GSeekableIface GSeekableIface;
typedef struct _GSettings GSettings;
typedef struct _GSettingsBackend GSettingsBackend;
typedef uint32_t GSettingsBindFlags;
typedef void* GSettingsBindGetMapping;
extern void _GSettingsBindGetMapping_c_wrapper();
extern void _GSettingsBindGetMapping_c_wrapper_once();
typedef void* GSettingsBindSetMapping;
extern void _GSettingsBindSetMapping_c_wrapper();
extern void _GSettingsBindSetMapping_c_wrapper_once();
typedef struct _GSettingsClass GSettingsClass;
typedef void* GSettingsGetMapping;
extern void _GSettingsGetMapping_c_wrapper();
extern void _GSettingsGetMapping_c_wrapper_once();
typedef struct _GSettingsPrivate GSettingsPrivate;
typedef struct _GSimpleAction GSimpleAction;
typedef struct _GSimpleActionGroup GSimpleActionGroup;
typedef struct _GSimpleActionGroupClass GSimpleActionGroupClass;
typedef struct _GSimpleActionGroupPrivate GSimpleActionGroupPrivate;
typedef struct _GSimpleAsyncResult GSimpleAsyncResult;
typedef struct _GSimpleAsyncResultClass GSimpleAsyncResultClass;
typedef void* GSimpleAsyncThreadFunc;
extern void _GSimpleAsyncThreadFunc_c_wrapper();
extern void _GSimpleAsyncThreadFunc_c_wrapper_once();
typedef struct _GSimplePermission GSimplePermission;
typedef struct _GSocket GSocket;
typedef struct _GSocketAddress GSocketAddress;
typedef struct _GSocketAddressClass GSocketAddressClass;
typedef struct _GSocketAddressEnumerator GSocketAddressEnumerator;
typedef struct _GSocketAddressEnumeratorClass GSocketAddressEnumeratorClass;
typedef struct _GSocketClass GSocketClass;
typedef struct _GSocketClient GSocketClient;
typedef struct _GSocketClientClass GSocketClientClass;
typedef struct _GSocketClientPrivate GSocketClientPrivate;
typedef struct _GSocketConnectable GSocketConnectable;
typedef struct _GSocketConnectableIface GSocketConnectableIface;
typedef struct _GSocketConnection GSocketConnection;
typedef struct _GSocketConnectionClass GSocketConnectionClass;
typedef struct _GSocketConnectionPrivate GSocketConnectionPrivate;
typedef struct _GSocketControlMessage GSocketControlMessage;
typedef struct _GSocketControlMessageClass GSocketControlMessageClass;
typedef struct _GSocketControlMessagePrivate GSocketControlMessagePrivate;
typedef uint32_t GSocketFamily;
typedef struct _GSocketListener GSocketListener;
typedef struct _GSocketListenerClass GSocketListenerClass;
typedef struct _GSocketListenerPrivate GSocketListenerPrivate;
typedef uint32_t GSocketMsgFlags;
typedef struct _GSocketPrivate GSocketPrivate;
typedef int32_t GSocketProtocol;
typedef struct _GSocketService GSocketService;
typedef struct _GSocketServiceClass GSocketServiceClass;
typedef struct _GSocketServicePrivate GSocketServicePrivate;
typedef void* GSocketSourceFunc;
extern void _GSocketSourceFunc_c_wrapper();
extern void _GSocketSourceFunc_c_wrapper_once();
typedef uint32_t GSocketType;
typedef struct _GSrvTarget GSrvTarget;
typedef struct _GTcpConnection GTcpConnection;
typedef struct _GTcpConnectionClass GTcpConnectionClass;
typedef struct _GTcpConnectionPrivate GTcpConnectionPrivate;
typedef struct _GTcpWrapperConnection GTcpWrapperConnection;
typedef struct _GTcpWrapperConnectionClass GTcpWrapperConnectionClass;
typedef struct _GTcpWrapperConnectionPrivate GTcpWrapperConnectionPrivate;
typedef struct _GThemedIcon GThemedIcon;
typedef struct _GThemedIconClass GThemedIconClass;
typedef struct _GThreadedSocketService GThreadedSocketService;
typedef struct _GThreadedSocketServiceClass GThreadedSocketServiceClass;
typedef struct _GThreadedSocketServicePrivate GThreadedSocketServicePrivate;
typedef uint32_t GTlsAuthenticationMode;
typedef struct _GTlsBackend GTlsBackend;
typedef struct _GTlsBackendInterface GTlsBackendInterface;
typedef struct _GTlsCertificate GTlsCertificate;
typedef struct _GTlsCertificateClass GTlsCertificateClass;
typedef uint32_t GTlsCertificateFlags;
typedef struct _GTlsCertificatePrivate GTlsCertificatePrivate;
typedef struct _GTlsClientConnection GTlsClientConnection;
typedef struct _GTlsClientConnectionInterface GTlsClientConnectionInterface;
typedef struct _GTlsConnection GTlsConnection;
typedef struct _GTlsConnectionClass GTlsConnectionClass;
typedef struct _GTlsConnectionPrivate GTlsConnectionPrivate;
typedef struct _GTlsDatabase GTlsDatabase;
typedef struct _GTlsDatabaseClass GTlsDatabaseClass;
typedef uint32_t GTlsDatabaseLookupFlags;
typedef struct _GTlsDatabasePrivate GTlsDatabasePrivate;
typedef uint32_t GTlsDatabaseVerifyFlags;
typedef uint32_t GTlsError;
typedef struct _GTlsFileDatabase GTlsFileDatabase;
typedef struct _GTlsFileDatabaseInterface GTlsFileDatabaseInterface;
typedef struct _GTlsInteraction GTlsInteraction;
typedef struct _GTlsInteractionClass GTlsInteractionClass;
typedef struct _GTlsInteractionPrivate GTlsInteractionPrivate;
typedef uint32_t GTlsInteractionResult;
typedef struct _GTlsPassword GTlsPassword;
typedef struct _GTlsPasswordClass GTlsPasswordClass;
typedef uint32_t GTlsPasswordFlags;
typedef struct _GTlsPasswordPrivate GTlsPasswordPrivate;
typedef uint32_t GTlsRehandshakeMode;
typedef struct _GTlsServerConnection GTlsServerConnection;
typedef struct _GTlsServerConnectionInterface GTlsServerConnectionInterface;
typedef struct _GUnixConnection GUnixConnection;
typedef struct _GUnixConnectionClass GUnixConnectionClass;
typedef struct _GUnixConnectionPrivate GUnixConnectionPrivate;
typedef struct _GUnixCredentialsMessage GUnixCredentialsMessage;
typedef struct _GUnixCredentialsMessageClass GUnixCredentialsMessageClass;
typedef struct _GUnixCredentialsMessagePrivate GUnixCredentialsMessagePrivate;
typedef struct _GUnixFDList GUnixFDList;
typedef struct _GUnixFDListClass GUnixFDListClass;
typedef struct _GUnixFDListPrivate GUnixFDListPrivate;
typedef struct _GUnixFDMessage GUnixFDMessage;
typedef struct _GUnixFDMessageClass GUnixFDMessageClass;
typedef struct _GUnixFDMessagePrivate GUnixFDMessagePrivate;
typedef struct _GUnixInputStream GUnixInputStream;
typedef struct _GUnixInputStreamClass GUnixInputStreamClass;
typedef struct _GUnixInputStreamPrivate GUnixInputStreamPrivate;
typedef struct _GUnixMountEntry GUnixMountEntry;
typedef struct _GUnixMountMonitor GUnixMountMonitor;
typedef struct _GUnixMountMonitorClass GUnixMountMonitorClass;
typedef struct _GUnixMountPoint GUnixMountPoint;
typedef struct _GUnixOutputStream GUnixOutputStream;
typedef struct _GUnixOutputStreamClass GUnixOutputStreamClass;
typedef struct _GUnixOutputStreamPrivate GUnixOutputStreamPrivate;
typedef struct _GUnixSocketAddress GUnixSocketAddress;
typedef struct _GUnixSocketAddressClass GUnixSocketAddressClass;
typedef struct _GUnixSocketAddressPrivate GUnixSocketAddressPrivate;
typedef uint32_t GUnixSocketAddressType;
typedef struct _GVfs GVfs;
typedef struct _GVfsClass GVfsClass;
typedef struct _GVolume GVolume;
typedef struct _GVolumeIface GVolumeIface;
typedef struct _GVolumeMonitor GVolumeMonitor;
typedef struct _GVolumeMonitorClass GVolumeMonitorClass;
typedef struct _GZlibCompressor GZlibCompressor;
typedef struct _GZlibCompressorClass GZlibCompressorClass;
typedef uint32_t GZlibCompressorFormat;
typedef struct _GZlibDecompressor GZlibDecompressor;
typedef struct _GZlibDecompressorClass GZlibDecompressorClass;
extern void g_action_activate(GAction*, GVariant*);
extern void g_action_change_state(GAction*, GVariant*);
extern int g_action_get_enabled(GAction*);
extern char* g_action_get_name(GAction*);
extern GVariantType* g_action_get_parameter_type(GAction*);
extern GVariant* g_action_get_state(GAction*);
extern GVariant* g_action_get_state_hint(GAction*);
extern GVariantType* g_action_get_state_type(GAction*);
extern GType g_action_get_type();
extern void g_action_group_action_added(GActionGroup*, char*);
extern void g_action_group_action_enabled_changed(GActionGroup*, char*, int);
extern void g_action_group_action_removed(GActionGroup*, char*);
extern void g_action_group_action_state_changed(GActionGroup*, char*, GVariant*);
extern void g_action_group_activate_action(GActionGroup*, char*, GVariant*);
extern void g_action_group_change_action_state(GActionGroup*, char*, GVariant*);
extern int g_action_group_get_action_enabled(GActionGroup*, char*);
extern GVariantType* g_action_group_get_action_parameter_type(GActionGroup*, char*);
extern GVariant* g_action_group_get_action_state(GActionGroup*, char*);
extern GVariant* g_action_group_get_action_state_hint(GActionGroup*, char*);
extern GVariantType* g_action_group_get_action_state_type(GActionGroup*, char*);
extern int g_action_group_has_action(GActionGroup*, char*);
extern char** g_action_group_list_actions(GActionGroup*);
extern GType g_action_group_get_type();
extern GAppInfo* g_app_info_create_from_commandline(char*, char*, GAppInfoCreateFlags, GError**);
extern GList* g_app_info_get_all();
extern GList* g_app_info_get_all_for_type(char*);
extern GAppInfo* g_app_info_get_default_for_type(char*, int);
extern GAppInfo* g_app_info_get_default_for_uri_scheme(char*);
extern GList* g_app_info_get_fallback_for_type(char*);
extern GList* g_app_info_get_recommended_for_type(char*);
extern int g_app_info_launch_default_for_uri(char*, GAppLaunchContext*, GError**);
extern void g_app_info_reset_type_associations(char*);
extern int g_app_info_add_supports_type(GAppInfo*, char*, GError**);
extern int g_app_info_can_delete(GAppInfo*);
extern int g_app_info_can_remove_supports_type(GAppInfo*);
extern int g_app_info_delete(GAppInfo*);
extern GAppInfo* g_app_info_dup(GAppInfo*);
extern int g_app_info_equal(GAppInfo*, GAppInfo*);
extern char* g_app_info_get_commandline(GAppInfo*);
extern char* g_app_info_get_description(GAppInfo*);
extern char* g_app_info_get_display_name(GAppInfo*);
extern char* g_app_info_get_executable(GAppInfo*);
extern GIcon* g_app_info_get_icon(GAppInfo*);
extern char* g_app_info_get_id(GAppInfo*);
extern char* g_app_info_get_name(GAppInfo*);
extern int g_app_info_launch(GAppInfo*, GList*, GAppLaunchContext*, GError**);
extern int g_app_info_launch_uris(GAppInfo*, GList*, GAppLaunchContext*, GError**);
extern int g_app_info_remove_supports_type(GAppInfo*, char*, GError**);
extern int g_app_info_set_as_default_for_extension(GAppInfo*, char*, GError**);
extern int g_app_info_set_as_default_for_type(GAppInfo*, char*, GError**);
extern int g_app_info_set_as_last_used_for_type(GAppInfo*, char*, GError**);
extern int g_app_info_should_show(GAppInfo*);
extern int g_app_info_supports_files(GAppInfo*);
extern int g_app_info_supports_uris(GAppInfo*);
extern GType g_app_info_get_type();
extern GAppLaunchContext* g_app_launch_context_new();
extern char* g_app_launch_context_get_display(GAppLaunchContext*, GAppInfo*, GList*);
extern char* g_app_launch_context_get_startup_notify_id(GAppLaunchContext*, GAppInfo*, GList*);
extern void g_app_launch_context_launch_failed(GAppLaunchContext*, char*);
extern GType g_app_launch_context_get_type();
extern GApplication* g_application_new(char*, GApplicationFlags);
extern int g_application_id_is_valid(char*);
extern void g_application_activate(GApplication*);
extern char* g_application_get_application_id(GApplication*);
extern GApplicationFlags g_application_get_flags(GApplication*);
extern uint32_t g_application_get_inactivity_timeout(GApplication*);
extern int g_application_get_is_registered(GApplication*);
extern int g_application_get_is_remote(GApplication*);
extern void g_application_hold(GApplication*);
extern void g_application_open(GApplication*, GFile**, int32_t, char*);
extern int g_application_register(GApplication*, GCancellable*, GError**);
extern void g_application_release(GApplication*);
extern int32_t g_application_run(GApplication*, int32_t, char**);
extern void g_application_set_action_group(GApplication*, GActionGroup*);
extern void g_application_set_application_id(GApplication*, char*);
extern void g_application_set_flags(GApplication*, GApplicationFlags);
extern void g_application_set_inactivity_timeout(GApplication*, uint32_t);
extern GType g_application_get_type();
extern char** g_application_command_line_get_arguments(GApplicationCommandLine*, int32_t*);
extern char* g_application_command_line_get_cwd(GApplicationCommandLine*);
extern char** g_application_command_line_get_environ(GApplicationCommandLine*);
extern int32_t g_application_command_line_get_exit_status(GApplicationCommandLine*);
extern int g_application_command_line_get_is_remote(GApplicationCommandLine*);
extern GVariant* g_application_command_line_get_platform_data(GApplicationCommandLine*);
extern char* g_application_command_line_getenv(GApplicationCommandLine*, char*);
extern void g_application_command_line_set_exit_status(GApplicationCommandLine*, int32_t);
extern GType g_application_command_line_get_type();
extern void g_async_initable_newv_async(GType, uint32_t, GParameter*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_async_initable_newv_async(GType arg0, uint32_t arg1, GParameter* arg2, int32_t arg3, GCancellable* arg4, void* gofunc) {
	if (gofunc) {
		g_async_initable_newv_async(arg0, arg1, arg2, arg3, arg4, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_async_initable_newv_async(arg0, arg1, arg2, arg3, arg4, 0, 0);
	}
}
extern void g_async_initable_init_async(GAsyncInitable*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_async_initable_init_async(GAsyncInitable* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_async_initable_init_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_async_initable_init_async(this, arg0, arg1, 0, 0);
	}
}
extern int g_async_initable_init_finish(GAsyncInitable*, GAsyncResult*, GError**);
extern GObject* g_async_initable_new_finish(GAsyncInitable*, GAsyncResult*, GError**);
extern GType g_async_initable_get_type();
extern GObject* g_async_result_get_source_object(GAsyncResult*);
extern void* g_async_result_get_user_data(GAsyncResult*);
extern GType g_async_result_get_type();
extern GInputStream* g_buffered_input_stream_new(GInputStream*);
extern GInputStream* g_buffered_input_stream_new_sized(GInputStream*, uint64_t);
extern int64_t g_buffered_input_stream_fill(GBufferedInputStream*, int64_t, GCancellable*, GError**);
extern void g_buffered_input_stream_fill_async(GBufferedInputStream*, int64_t, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_buffered_input_stream_fill_async(GBufferedInputStream* this, int64_t arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_buffered_input_stream_fill_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_buffered_input_stream_fill_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int64_t g_buffered_input_stream_fill_finish(GBufferedInputStream*, GAsyncResult*, GError**);
extern uint64_t g_buffered_input_stream_get_available(GBufferedInputStream*);
extern uint64_t g_buffered_input_stream_get_buffer_size(GBufferedInputStream*);
extern uint64_t g_buffered_input_stream_peek(GBufferedInputStream*, void*, uint64_t, uint64_t);
extern uint8_t* g_buffered_input_stream_peek_buffer(GBufferedInputStream*, uint64_t*);
extern int32_t g_buffered_input_stream_read_byte(GBufferedInputStream*, GCancellable*, GError**);
extern void g_buffered_input_stream_set_buffer_size(GBufferedInputStream*, uint64_t);
extern GType g_buffered_input_stream_get_type();
extern GOutputStream* g_buffered_output_stream_new(GOutputStream*);
extern GOutputStream* g_buffered_output_stream_new_sized(GOutputStream*, uint64_t);
extern int g_buffered_output_stream_get_auto_grow(GBufferedOutputStream*);
extern uint64_t g_buffered_output_stream_get_buffer_size(GBufferedOutputStream*);
extern void g_buffered_output_stream_set_auto_grow(GBufferedOutputStream*, int);
extern void g_buffered_output_stream_set_buffer_size(GBufferedOutputStream*, uint64_t);
extern GType g_buffered_output_stream_get_type();
extern GCancellable* g_cancellable_new();
extern GCancellable* g_cancellable_get_current();
extern void g_cancellable_cancel(GCancellable*);
extern uint64_t g_cancellable_connect(GCancellable*, GCallback, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint64_t _g_cancellable_connect(GCancellable* this, void* gofunc) {
	if (gofunc) {
		return g_cancellable_connect(this, _GCallback_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return g_cancellable_connect(this, 0, 0, 0);
	}
}
extern void g_cancellable_disconnect(GCancellable*, uint64_t);
extern int32_t g_cancellable_get_fd(GCancellable*);
extern int g_cancellable_is_cancelled(GCancellable*);
extern int g_cancellable_make_pollfd(GCancellable*, GPollFD*);
extern void g_cancellable_pop_current(GCancellable*);
extern void g_cancellable_push_current(GCancellable*);
extern void g_cancellable_release_fd(GCancellable*);
extern void g_cancellable_reset(GCancellable*);
extern int g_cancellable_set_error_if_cancelled(GCancellable*, GError**);
extern GType g_cancellable_get_type();
extern GCharsetConverter* g_charset_converter_new(char*, char*, GError**);
extern uint32_t g_charset_converter_get_num_fallbacks(GCharsetConverter*);
extern int g_charset_converter_get_use_fallback(GCharsetConverter*);
extern void g_charset_converter_set_use_fallback(GCharsetConverter*, int);
extern GType g_charset_converter_get_type();
extern GConverterResult g_converter_convert(GConverter*, uint8_t*, uint64_t, void*, uint64_t, GConverterFlags, uint64_t*, uint64_t*, GError**);
extern void g_converter_reset(GConverter*);
extern GType g_converter_get_type();
extern GInputStream* g_converter_input_stream_new(GInputStream*, GConverter*);
extern GConverter* g_converter_input_stream_get_converter(GConverterInputStream*);
extern GType g_converter_input_stream_get_type();
extern GOutputStream* g_converter_output_stream_new(GOutputStream*, GConverter*);
extern GConverter* g_converter_output_stream_get_converter(GConverterOutputStream*);
extern GType g_converter_output_stream_get_type();
extern GCredentials* g_credentials_new();
extern uint32_t g_credentials_get_unix_user(GCredentials*, GError**);
extern int g_credentials_is_same_user(GCredentials*, GCredentials*, GError**);
extern void g_credentials_set_native(GCredentials*, GCredentialsType, void*);
extern int g_credentials_set_unix_user(GCredentials*, uint32_t, GError**);
extern char* g_credentials_to_string(GCredentials*);
extern GType g_credentials_get_type();
extern GDBusAnnotationInfo* g_dbus_annotation_info_ref(GDBusAnnotationInfo*);
extern void g_dbus_annotation_info_unref(GDBusAnnotationInfo*);
extern char* g_dbus_annotation_info_lookup(GDBusAnnotationInfo*, char*);
extern GDBusArgInfo* g_dbus_arg_info_ref(GDBusArgInfo*);
extern void g_dbus_arg_info_unref(GDBusArgInfo*);
extern GDBusAuthObserver* g_dbus_auth_observer_new();
extern int g_dbus_auth_observer_authorize_authenticated_peer(GDBusAuthObserver*, GIOStream*, GCredentials*);
extern GType g_dbus_auth_observer_get_type();
extern GDBusConnection* g_dbus_connection_new_finish(GAsyncResult*, GError**);
extern GDBusConnection* g_dbus_connection_new_for_address_finish(GAsyncResult*, GError**);
extern GDBusConnection* g_dbus_connection_new_for_address_sync(char*, GDBusConnectionFlags, GDBusAuthObserver*, GCancellable*, GError**);
extern GDBusConnection* g_dbus_connection_new_sync(GIOStream*, char*, GDBusConnectionFlags, GDBusAuthObserver*, GCancellable*, GError**);
extern void g_dbus_connection_new(GIOStream*, char*, GDBusConnectionFlags, GDBusAuthObserver*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_connection_new(GIOStream* arg0, char* arg1, GDBusConnectionFlags arg2, GDBusAuthObserver* arg3, GCancellable* arg4, void* gofunc) {
	if (gofunc) {
		g_dbus_connection_new(arg0, arg1, arg2, arg3, arg4, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_connection_new(arg0, arg1, arg2, arg3, arg4, 0, 0);
	}
}
extern void g_dbus_connection_new_for_address(char*, GDBusConnectionFlags, GDBusAuthObserver*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_connection_new_for_address(char* arg0, GDBusConnectionFlags arg1, GDBusAuthObserver* arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_dbus_connection_new_for_address(arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_connection_new_for_address(arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern uint32_t g_dbus_connection_add_filter(GDBusConnection*, GDBusMessageFilterFunction, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint32_t _g_dbus_connection_add_filter(GDBusConnection* this, void* gofunc) {
	if (gofunc) {
		return g_dbus_connection_add_filter(this, _GDBusMessageFilterFunction_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return g_dbus_connection_add_filter(this, 0, 0, 0);
	}
}
extern void g_dbus_connection_call(GDBusConnection*, char*, char*, char*, char*, GVariant*, GVariantType*, GDBusCallFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_connection_call(GDBusConnection* this, char* arg0, char* arg1, char* arg2, char* arg3, GVariant* arg4, GVariantType* arg5, GDBusCallFlags arg6, int32_t arg7, GCancellable* arg8, void* gofunc) {
	if (gofunc) {
		g_dbus_connection_call(this, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_connection_call(this, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, 0, 0);
	}
}
extern GVariant* g_dbus_connection_call_finish(GDBusConnection*, GAsyncResult*, GError**);
extern GVariant* g_dbus_connection_call_sync(GDBusConnection*, char*, char*, char*, char*, GVariant*, GVariantType*, GDBusCallFlags, int32_t, GCancellable*, GError**);
extern void g_dbus_connection_call_with_unix_fd_list(GDBusConnection*, char*, char*, char*, char*, GVariant*, GVariantType*, GDBusCallFlags, int32_t, GUnixFDList*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_connection_call_with_unix_fd_list(GDBusConnection* this, char* arg0, char* arg1, char* arg2, char* arg3, GVariant* arg4, GVariantType* arg5, GDBusCallFlags arg6, int32_t arg7, GUnixFDList* arg8, GCancellable* arg9, void* gofunc) {
	if (gofunc) {
		g_dbus_connection_call_with_unix_fd_list(this, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_connection_call_with_unix_fd_list(this, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, 0, 0);
	}
}
extern GVariant* g_dbus_connection_call_with_unix_fd_list_finish(GDBusConnection*, GUnixFDList**, GAsyncResult*, GError**);
extern GVariant* g_dbus_connection_call_with_unix_fd_list_sync(GDBusConnection*, char*, char*, char*, char*, GVariant*, GVariantType*, GDBusCallFlags, int32_t, GUnixFDList*, GUnixFDList**, GCancellable*, GError**);
extern void g_dbus_connection_close(GDBusConnection*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_connection_close(GDBusConnection* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_dbus_connection_close(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_connection_close(this, arg0, 0, 0);
	}
}
extern int g_dbus_connection_close_finish(GDBusConnection*, GAsyncResult*, GError**);
extern int g_dbus_connection_close_sync(GDBusConnection*, GCancellable*, GError**);
extern int g_dbus_connection_emit_signal(GDBusConnection*, char*, char*, char*, char*, GVariant*, GError**);
extern void g_dbus_connection_flush(GDBusConnection*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_connection_flush(GDBusConnection* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_dbus_connection_flush(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_connection_flush(this, arg0, 0, 0);
	}
}
extern int g_dbus_connection_flush_finish(GDBusConnection*, GAsyncResult*, GError**);
extern int g_dbus_connection_flush_sync(GDBusConnection*, GCancellable*, GError**);
extern GDBusCapabilityFlags g_dbus_connection_get_capabilities(GDBusConnection*);
extern int g_dbus_connection_get_exit_on_close(GDBusConnection*);
extern char* g_dbus_connection_get_guid(GDBusConnection*);
extern GCredentials* g_dbus_connection_get_peer_credentials(GDBusConnection*);
extern GIOStream* g_dbus_connection_get_stream(GDBusConnection*);
extern char* g_dbus_connection_get_unique_name(GDBusConnection*);
extern int g_dbus_connection_is_closed(GDBusConnection*);
extern uint32_t g_dbus_connection_register_object(GDBusConnection*, char*, GDBusInterfaceInfo*, GDBusInterfaceVTable*, void*, GDestroyNotify, GError**);
extern uint32_t g_dbus_connection_register_subtree(GDBusConnection*, char*, GDBusSubtreeVTable*, GDBusSubtreeFlags, void*, GDestroyNotify, GError**);
extern void g_dbus_connection_remove_filter(GDBusConnection*, uint32_t);
extern int g_dbus_connection_send_message(GDBusConnection*, GDBusMessage*, GDBusSendMessageFlags, uint32_t*, GError**);
extern void g_dbus_connection_send_message_with_reply(GDBusConnection*, GDBusMessage*, GDBusSendMessageFlags, int32_t, uint32_t*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_connection_send_message_with_reply(GDBusConnection* this, GDBusMessage* arg0, GDBusSendMessageFlags arg1, int32_t arg2, uint32_t* arg3, GCancellable* arg4, void* gofunc) {
	if (gofunc) {
		g_dbus_connection_send_message_with_reply(this, arg0, arg1, arg2, arg3, arg4, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_connection_send_message_with_reply(this, arg0, arg1, arg2, arg3, arg4, 0, 0);
	}
}
extern GDBusMessage* g_dbus_connection_send_message_with_reply_finish(GDBusConnection*, GAsyncResult*, GError**);
extern GDBusMessage* g_dbus_connection_send_message_with_reply_sync(GDBusConnection*, GDBusMessage*, GDBusSendMessageFlags, int32_t, uint32_t*, GCancellable*, GError**);
extern void g_dbus_connection_set_exit_on_close(GDBusConnection*, int);
extern uint32_t g_dbus_connection_signal_subscribe(GDBusConnection*, char*, char*, char*, char*, char*, GDBusSignalFlags, GDBusSignalCallback, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint32_t _g_dbus_connection_signal_subscribe(GDBusConnection* this, char* arg0, char* arg1, char* arg2, char* arg3, char* arg4, GDBusSignalFlags arg5, void* gofunc) {
	if (gofunc) {
		return g_dbus_connection_signal_subscribe(this, arg0, arg1, arg2, arg3, arg4, arg5, _GDBusSignalCallback_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return g_dbus_connection_signal_subscribe(this, arg0, arg1, arg2, arg3, arg4, arg5, 0, 0, 0);
	}
}
extern void g_dbus_connection_signal_unsubscribe(GDBusConnection*, uint32_t);
extern void g_dbus_connection_start_message_processing(GDBusConnection*);
extern int g_dbus_connection_unregister_object(GDBusConnection*, uint32_t);
extern int g_dbus_connection_unregister_subtree(GDBusConnection*, uint32_t);
extern GType g_dbus_connection_get_type();
extern GDBusInterfaceInfo* g_dbus_interface_get_info(GDBusInterface*);
extern GDBusObject* g_dbus_interface_get_object(GDBusInterface*);
extern void g_dbus_interface_set_object(GDBusInterface*, GDBusObject*);
extern GType g_dbus_interface_get_type();
extern void g_dbus_interface_info_cache_build(GDBusInterfaceInfo*);
extern void g_dbus_interface_info_cache_release(GDBusInterfaceInfo*);
extern void g_dbus_interface_info_generate_xml(GDBusInterfaceInfo*, uint32_t, GString*);
extern GDBusMethodInfo* g_dbus_interface_info_lookup_method(GDBusInterfaceInfo*, char*);
extern GDBusPropertyInfo* g_dbus_interface_info_lookup_property(GDBusInterfaceInfo*, char*);
extern GDBusSignalInfo* g_dbus_interface_info_lookup_signal(GDBusInterfaceInfo*, char*);
extern GDBusInterfaceInfo* g_dbus_interface_info_ref(GDBusInterfaceInfo*);
extern void g_dbus_interface_info_unref(GDBusInterfaceInfo*);
extern int g_dbus_interface_skeleton_export(GDBusInterfaceSkeleton*, GDBusConnection*, char*, GError**);
extern void g_dbus_interface_skeleton_flush(GDBusInterfaceSkeleton*);
extern GDBusConnection* g_dbus_interface_skeleton_get_connection(GDBusInterfaceSkeleton*);
extern GDBusInterfaceSkeletonFlags g_dbus_interface_skeleton_get_flags(GDBusInterfaceSkeleton*);
extern GDBusInterfaceInfo* g_dbus_interface_skeleton_get_info(GDBusInterfaceSkeleton*);
extern char* g_dbus_interface_skeleton_get_object_path(GDBusInterfaceSkeleton*);
extern GVariant* g_dbus_interface_skeleton_get_properties(GDBusInterfaceSkeleton*);
extern void g_dbus_interface_skeleton_set_flags(GDBusInterfaceSkeleton*, GDBusInterfaceSkeletonFlags);
extern void g_dbus_interface_skeleton_unexport(GDBusInterfaceSkeleton*);
extern GType g_dbus_interface_skeleton_get_type();
extern GDBusMessage* g_dbus_message_new();
extern GDBusMessage* g_dbus_message_new_from_blob(uint8_t*, uint64_t, GDBusCapabilityFlags, GError**);
extern GDBusMessage* g_dbus_message_new_method_call(char*, char*, char*, char*);
extern GDBusMessage* g_dbus_message_new_signal(char*, char*, char*);
extern int64_t g_dbus_message_bytes_needed(uint8_t*, uint64_t, GError**);
extern GDBusMessage* g_dbus_message_copy(GDBusMessage*, GError**);
extern char* g_dbus_message_get_arg0(GDBusMessage*);
extern GVariant* g_dbus_message_get_body(GDBusMessage*);
extern GDBusMessageByteOrder g_dbus_message_get_byte_order(GDBusMessage*);
extern char* g_dbus_message_get_destination(GDBusMessage*);
extern char* g_dbus_message_get_error_name(GDBusMessage*);
extern GDBusMessageFlags g_dbus_message_get_flags(GDBusMessage*);
extern GVariant* g_dbus_message_get_header(GDBusMessage*, GDBusMessageHeaderField);
extern uint8_t* g_dbus_message_get_header_fields(GDBusMessage*);
extern char* g_dbus_message_get_interface(GDBusMessage*);
extern int g_dbus_message_get_locked(GDBusMessage*);
extern char* g_dbus_message_get_member(GDBusMessage*);
extern GDBusMessageType g_dbus_message_get_message_type(GDBusMessage*);
extern uint32_t g_dbus_message_get_num_unix_fds(GDBusMessage*);
extern char* g_dbus_message_get_path(GDBusMessage*);
extern uint32_t g_dbus_message_get_reply_serial(GDBusMessage*);
extern char* g_dbus_message_get_sender(GDBusMessage*);
extern uint32_t g_dbus_message_get_serial(GDBusMessage*);
extern char* g_dbus_message_get_signature(GDBusMessage*);
extern GUnixFDList* g_dbus_message_get_unix_fd_list(GDBusMessage*);
extern void g_dbus_message_lock(GDBusMessage*);
extern GDBusMessage* g_dbus_message_new_method_error_literal(GDBusMessage*, char*, char*);
extern GDBusMessage* g_dbus_message_new_method_reply(GDBusMessage*);
extern char* g_dbus_message_print(GDBusMessage*, uint32_t);
extern void g_dbus_message_set_body(GDBusMessage*, GVariant*);
extern void g_dbus_message_set_byte_order(GDBusMessage*, GDBusMessageByteOrder);
extern void g_dbus_message_set_destination(GDBusMessage*, char*);
extern void g_dbus_message_set_error_name(GDBusMessage*, char*);
extern void g_dbus_message_set_flags(GDBusMessage*, GDBusMessageFlags);
extern void g_dbus_message_set_header(GDBusMessage*, GDBusMessageHeaderField, GVariant*);
extern void g_dbus_message_set_interface(GDBusMessage*, char*);
extern void g_dbus_message_set_member(GDBusMessage*, char*);
extern void g_dbus_message_set_message_type(GDBusMessage*, GDBusMessageType);
extern void g_dbus_message_set_num_unix_fds(GDBusMessage*, uint32_t);
extern void g_dbus_message_set_path(GDBusMessage*, char*);
extern void g_dbus_message_set_reply_serial(GDBusMessage*, uint32_t);
extern void g_dbus_message_set_sender(GDBusMessage*, char*);
extern void g_dbus_message_set_serial(GDBusMessage*, uint32_t);
extern void g_dbus_message_set_signature(GDBusMessage*, char*);
extern void g_dbus_message_set_unix_fd_list(GDBusMessage*, GUnixFDList*);
extern uint8_t* g_dbus_message_to_blob(GDBusMessage*, uint64_t*, GDBusCapabilityFlags, GError**);
extern int g_dbus_message_to_gerror(GDBusMessage*, GError**);
extern GType g_dbus_message_get_type();
extern GDBusMethodInfo* g_dbus_method_info_ref(GDBusMethodInfo*);
extern void g_dbus_method_info_unref(GDBusMethodInfo*);
extern GDBusConnection* g_dbus_method_invocation_get_connection(GDBusMethodInvocation*);
extern char* g_dbus_method_invocation_get_interface_name(GDBusMethodInvocation*);
extern GDBusMessage* g_dbus_method_invocation_get_message(GDBusMethodInvocation*);
extern GDBusMethodInfo* g_dbus_method_invocation_get_method_info(GDBusMethodInvocation*);
extern char* g_dbus_method_invocation_get_method_name(GDBusMethodInvocation*);
extern char* g_dbus_method_invocation_get_object_path(GDBusMethodInvocation*);
extern GVariant* g_dbus_method_invocation_get_parameters(GDBusMethodInvocation*);
extern char* g_dbus_method_invocation_get_sender(GDBusMethodInvocation*);
extern void g_dbus_method_invocation_return_dbus_error(GDBusMethodInvocation*, char*, char*);
extern void g_dbus_method_invocation_return_error_literal(GDBusMethodInvocation*, uint32_t, int32_t, char*);
extern void g_dbus_method_invocation_return_gerror(GDBusMethodInvocation*, GError*);
extern void g_dbus_method_invocation_return_value(GDBusMethodInvocation*, GVariant*);
extern void g_dbus_method_invocation_return_value_with_unix_fd_list(GDBusMethodInvocation*, GVariant*, GUnixFDList*);
extern GType g_dbus_method_invocation_get_type();
extern GDBusNodeInfo* g_dbus_node_info_new_for_xml(char*, GError**);
extern void g_dbus_node_info_generate_xml(GDBusNodeInfo*, uint32_t, GString*);
extern GDBusInterfaceInfo* g_dbus_node_info_lookup_interface(GDBusNodeInfo*, char*);
extern GDBusNodeInfo* g_dbus_node_info_ref(GDBusNodeInfo*);
extern void g_dbus_node_info_unref(GDBusNodeInfo*);
extern GDBusInterface* g_dbus_object_get_interface(GDBusObject*, char*);
extern GList* g_dbus_object_get_interfaces(GDBusObject*);
extern char* g_dbus_object_get_object_path(GDBusObject*);
extern GType g_dbus_object_get_type();
extern GDBusInterface* g_dbus_object_manager_get_interface(GDBusObjectManager*, char*, char*);
extern GDBusObject* g_dbus_object_manager_get_object(GDBusObjectManager*, char*);
extern char* g_dbus_object_manager_get_object_path(GDBusObjectManager*);
extern GList* g_dbus_object_manager_get_objects(GDBusObjectManager*);
extern GType g_dbus_object_manager_get_type();
extern GDBusObjectManagerClient* g_dbus_object_manager_client_new_finish(GAsyncResult*, GError**);
extern GDBusObjectManagerClient* g_dbus_object_manager_client_new_for_bus_finish(GAsyncResult*, GError**);
extern GDBusObjectManagerClient* g_dbus_object_manager_client_new_for_bus_sync(GBusType, GDBusObjectManagerClientFlags, char*, char*, GDBusProxyTypeFunc, void*, GDestroyNotify, GCancellable*, GError**);
#pragma GCC diagnostic ignored "-Wunused-function"
static GDBusObjectManagerClient* _g_dbus_object_manager_client_new_for_bus_sync(GBusType arg0, GDBusObjectManagerClientFlags arg1, char* arg2, char* arg3, void* gofunc, GCancellable* arg7, GError** arg8) {
	if (gofunc) {
		return g_dbus_object_manager_client_new_for_bus_sync(arg0, arg1, arg2, arg3, _GDBusProxyTypeFunc_c_wrapper, gofunc, _c_callback_cleanup, arg7, arg8);
	} else {
		return g_dbus_object_manager_client_new_for_bus_sync(arg0, arg1, arg2, arg3, 0, 0, 0, arg7, arg8);
	}
}
extern GDBusObjectManagerClient* g_dbus_object_manager_client_new_sync(GDBusConnection*, GDBusObjectManagerClientFlags, char*, char*, GDBusProxyTypeFunc, void*, GDestroyNotify, GCancellable*, GError**);
#pragma GCC diagnostic ignored "-Wunused-function"
static GDBusObjectManagerClient* _g_dbus_object_manager_client_new_sync(GDBusConnection* arg0, GDBusObjectManagerClientFlags arg1, char* arg2, char* arg3, void* gofunc, GCancellable* arg7, GError** arg8) {
	if (gofunc) {
		return g_dbus_object_manager_client_new_sync(arg0, arg1, arg2, arg3, _GDBusProxyTypeFunc_c_wrapper, gofunc, _c_callback_cleanup, arg7, arg8);
	} else {
		return g_dbus_object_manager_client_new_sync(arg0, arg1, arg2, arg3, 0, 0, 0, arg7, arg8);
	}
}
extern void g_dbus_object_manager_client_new(GDBusConnection*, GDBusObjectManagerClientFlags, char*, char*, GDBusProxyTypeFunc, void*, GDestroyNotify, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_object_manager_client_new(GDBusConnection* arg0, GDBusObjectManagerClientFlags arg1, char* arg2, char* arg3, GDBusProxyTypeFunc arg4, void* arg5, GCancellable* arg7, void* gofunc) {
	if (gofunc) {
		g_dbus_object_manager_client_new(arg0, arg1, arg2, arg3, arg4, arg5, _c_callback_cleanup, arg7, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_object_manager_client_new(arg0, arg1, arg2, arg3, arg4, arg5, 0, arg7, 0, 0);
	}
}
extern void g_dbus_object_manager_client_new_for_bus(GBusType, GDBusObjectManagerClientFlags, char*, char*, GDBusProxyTypeFunc, void*, GDestroyNotify, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_object_manager_client_new_for_bus(GBusType arg0, GDBusObjectManagerClientFlags arg1, char* arg2, char* arg3, GDBusProxyTypeFunc arg4, void* arg5, GCancellable* arg7, void* gofunc) {
	if (gofunc) {
		g_dbus_object_manager_client_new_for_bus(arg0, arg1, arg2, arg3, arg4, arg5, _c_callback_cleanup, arg7, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_object_manager_client_new_for_bus(arg0, arg1, arg2, arg3, arg4, arg5, 0, arg7, 0, 0);
	}
}
extern GDBusConnection* g_dbus_object_manager_client_get_connection(GDBusObjectManagerClient*);
extern GDBusObjectManagerClientFlags g_dbus_object_manager_client_get_flags(GDBusObjectManagerClient*);
extern char* g_dbus_object_manager_client_get_name(GDBusObjectManagerClient*);
extern char* g_dbus_object_manager_client_get_name_owner(GDBusObjectManagerClient*);
extern GType g_dbus_object_manager_client_get_type();
extern GDBusObjectManagerServer* g_dbus_object_manager_server_new(char*);
extern void g_dbus_object_manager_server_export(GDBusObjectManagerServer*, GDBusObjectSkeleton*);
extern void g_dbus_object_manager_server_export_uniquely(GDBusObjectManagerServer*, GDBusObjectSkeleton*);
extern GDBusConnection* g_dbus_object_manager_server_get_connection(GDBusObjectManagerServer*);
extern void g_dbus_object_manager_server_set_connection(GDBusObjectManagerServer*, GDBusConnection*);
extern int g_dbus_object_manager_server_unexport(GDBusObjectManagerServer*, char*);
extern GType g_dbus_object_manager_server_get_type();
extern GDBusObjectProxy* g_dbus_object_proxy_new(GDBusConnection*, char*);
extern GDBusConnection* g_dbus_object_proxy_get_connection(GDBusObjectProxy*);
extern GType g_dbus_object_proxy_get_type();
extern GDBusObjectSkeleton* g_dbus_object_skeleton_new(char*);
extern void g_dbus_object_skeleton_add_interface(GDBusObjectSkeleton*, GDBusInterfaceSkeleton*);
extern void g_dbus_object_skeleton_flush(GDBusObjectSkeleton*);
extern void g_dbus_object_skeleton_remove_interface(GDBusObjectSkeleton*, GDBusInterfaceSkeleton*);
extern void g_dbus_object_skeleton_remove_interface_by_name(GDBusObjectSkeleton*, char*);
extern void g_dbus_object_skeleton_set_object_path(GDBusObjectSkeleton*, char*);
extern GType g_dbus_object_skeleton_get_type();
extern GDBusPropertyInfo* g_dbus_property_info_ref(GDBusPropertyInfo*);
extern void g_dbus_property_info_unref(GDBusPropertyInfo*);
extern GDBusProxy* g_dbus_proxy_new_finish(GAsyncResult*, GError**);
extern GDBusProxy* g_dbus_proxy_new_for_bus_finish(GAsyncResult*, GError**);
extern GDBusProxy* g_dbus_proxy_new_for_bus_sync(GBusType, GDBusProxyFlags, GDBusInterfaceInfo*, char*, char*, char*, GCancellable*, GError**);
extern GDBusProxy* g_dbus_proxy_new_sync(GDBusConnection*, GDBusProxyFlags, GDBusInterfaceInfo*, char*, char*, char*, GCancellable*, GError**);
extern void g_dbus_proxy_new(GDBusConnection*, GDBusProxyFlags, GDBusInterfaceInfo*, char*, char*, char*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_proxy_new(GDBusConnection* arg0, GDBusProxyFlags arg1, GDBusInterfaceInfo* arg2, char* arg3, char* arg4, char* arg5, GCancellable* arg6, void* gofunc) {
	if (gofunc) {
		g_dbus_proxy_new(arg0, arg1, arg2, arg3, arg4, arg5, arg6, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_proxy_new(arg0, arg1, arg2, arg3, arg4, arg5, arg6, 0, 0);
	}
}
extern void g_dbus_proxy_new_for_bus(GBusType, GDBusProxyFlags, GDBusInterfaceInfo*, char*, char*, char*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_proxy_new_for_bus(GBusType arg0, GDBusProxyFlags arg1, GDBusInterfaceInfo* arg2, char* arg3, char* arg4, char* arg5, GCancellable* arg6, void* gofunc) {
	if (gofunc) {
		g_dbus_proxy_new_for_bus(arg0, arg1, arg2, arg3, arg4, arg5, arg6, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_proxy_new_for_bus(arg0, arg1, arg2, arg3, arg4, arg5, arg6, 0, 0);
	}
}
extern void g_dbus_proxy_call(GDBusProxy*, char*, GVariant*, GDBusCallFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_proxy_call(GDBusProxy* this, char* arg0, GVariant* arg1, GDBusCallFlags arg2, int32_t arg3, GCancellable* arg4, void* gofunc) {
	if (gofunc) {
		g_dbus_proxy_call(this, arg0, arg1, arg2, arg3, arg4, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_proxy_call(this, arg0, arg1, arg2, arg3, arg4, 0, 0);
	}
}
extern GVariant* g_dbus_proxy_call_finish(GDBusProxy*, GAsyncResult*, GError**);
extern GVariant* g_dbus_proxy_call_sync(GDBusProxy*, char*, GVariant*, GDBusCallFlags, int32_t, GCancellable*, GError**);
extern void g_dbus_proxy_call_with_unix_fd_list(GDBusProxy*, char*, GVariant*, GDBusCallFlags, int32_t, GUnixFDList*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_proxy_call_with_unix_fd_list(GDBusProxy* this, char* arg0, GVariant* arg1, GDBusCallFlags arg2, int32_t arg3, GUnixFDList* arg4, GCancellable* arg5, void* gofunc) {
	if (gofunc) {
		g_dbus_proxy_call_with_unix_fd_list(this, arg0, arg1, arg2, arg3, arg4, arg5, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_proxy_call_with_unix_fd_list(this, arg0, arg1, arg2, arg3, arg4, arg5, 0, 0);
	}
}
extern GVariant* g_dbus_proxy_call_with_unix_fd_list_finish(GDBusProxy*, GUnixFDList**, GAsyncResult*, GError**);
extern GVariant* g_dbus_proxy_call_with_unix_fd_list_sync(GDBusProxy*, char*, GVariant*, GDBusCallFlags, int32_t, GUnixFDList*, GUnixFDList**, GCancellable*, GError**);
extern GVariant* g_dbus_proxy_get_cached_property(GDBusProxy*, char*);
extern char** g_dbus_proxy_get_cached_property_names(GDBusProxy*);
extern GDBusConnection* g_dbus_proxy_get_connection(GDBusProxy*);
extern int32_t g_dbus_proxy_get_default_timeout(GDBusProxy*);
extern GDBusProxyFlags g_dbus_proxy_get_flags(GDBusProxy*);
extern GDBusInterfaceInfo* g_dbus_proxy_get_interface_info(GDBusProxy*);
extern char* g_dbus_proxy_get_interface_name(GDBusProxy*);
extern char* g_dbus_proxy_get_name(GDBusProxy*);
extern char* g_dbus_proxy_get_name_owner(GDBusProxy*);
extern char* g_dbus_proxy_get_object_path(GDBusProxy*);
extern void g_dbus_proxy_set_cached_property(GDBusProxy*, char*, GVariant*);
extern void g_dbus_proxy_set_default_timeout(GDBusProxy*, int32_t);
extern void g_dbus_proxy_set_interface_info(GDBusProxy*, GDBusInterfaceInfo*);
extern GType g_dbus_proxy_get_type();
extern GDBusServer* g_dbus_server_new_sync(char*, GDBusServerFlags, char*, GDBusAuthObserver*, GCancellable*, GError**);
extern char* g_dbus_server_get_client_address(GDBusServer*);
extern GDBusServerFlags g_dbus_server_get_flags(GDBusServer*);
extern char* g_dbus_server_get_guid(GDBusServer*);
extern int g_dbus_server_is_active(GDBusServer*);
extern void g_dbus_server_start(GDBusServer*);
extern void g_dbus_server_stop(GDBusServer*);
extern GType g_dbus_server_get_type();
extern GDBusSignalInfo* g_dbus_signal_info_ref(GDBusSignalInfo*);
extern void g_dbus_signal_info_unref(GDBusSignalInfo*);
extern GDataInputStream* g_data_input_stream_new(GInputStream*);
extern GDataStreamByteOrder g_data_input_stream_get_byte_order(GDataInputStream*);
extern GDataStreamNewlineType g_data_input_stream_get_newline_type(GDataInputStream*);
extern uint8_t g_data_input_stream_read_byte(GDataInputStream*, GCancellable*, GError**);
extern int16_t g_data_input_stream_read_int16(GDataInputStream*, GCancellable*, GError**);
extern int32_t g_data_input_stream_read_int32(GDataInputStream*, GCancellable*, GError**);
extern int64_t g_data_input_stream_read_int64(GDataInputStream*, GCancellable*, GError**);
extern uint8_t* g_data_input_stream_read_line(GDataInputStream*, uint64_t*, GCancellable*, GError**);
extern void g_data_input_stream_read_line_async(GDataInputStream*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_data_input_stream_read_line_async(GDataInputStream* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_data_input_stream_read_line_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_data_input_stream_read_line_async(this, arg0, arg1, 0, 0);
	}
}
extern uint8_t* g_data_input_stream_read_line_finish(GDataInputStream*, GAsyncResult*, uint64_t*, GError**);
extern char* g_data_input_stream_read_line_finish_utf8(GDataInputStream*, GAsyncResult*, uint64_t*, GError**);
extern char* g_data_input_stream_read_line_utf8(GDataInputStream*, uint64_t*, GCancellable*, GError**);
extern uint16_t g_data_input_stream_read_uint16(GDataInputStream*, GCancellable*, GError**);
extern uint32_t g_data_input_stream_read_uint32(GDataInputStream*, GCancellable*, GError**);
extern uint64_t g_data_input_stream_read_uint64(GDataInputStream*, GCancellable*, GError**);
extern char* g_data_input_stream_read_until(GDataInputStream*, char*, uint64_t*, GCancellable*, GError**);
extern void g_data_input_stream_read_until_async(GDataInputStream*, char*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_data_input_stream_read_until_async(GDataInputStream* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_data_input_stream_read_until_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_data_input_stream_read_until_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern char* g_data_input_stream_read_until_finish(GDataInputStream*, GAsyncResult*, uint64_t*, GError**);
extern char* g_data_input_stream_read_upto(GDataInputStream*, char*, int64_t, uint64_t*, GCancellable*, GError**);
extern void g_data_input_stream_read_upto_async(GDataInputStream*, char*, int64_t, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_data_input_stream_read_upto_async(GDataInputStream* this, char* arg0, int64_t arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_data_input_stream_read_upto_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_data_input_stream_read_upto_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern char* g_data_input_stream_read_upto_finish(GDataInputStream*, GAsyncResult*, uint64_t*, GError**);
extern void g_data_input_stream_set_byte_order(GDataInputStream*, GDataStreamByteOrder);
extern void g_data_input_stream_set_newline_type(GDataInputStream*, GDataStreamNewlineType);
extern GType g_data_input_stream_get_type();
extern GDataOutputStream* g_data_output_stream_new(GOutputStream*);
extern GDataStreamByteOrder g_data_output_stream_get_byte_order(GDataOutputStream*);
extern int g_data_output_stream_put_byte(GDataOutputStream*, uint8_t, GCancellable*, GError**);
extern int g_data_output_stream_put_int16(GDataOutputStream*, int16_t, GCancellable*, GError**);
extern int g_data_output_stream_put_int32(GDataOutputStream*, int32_t, GCancellable*, GError**);
extern int g_data_output_stream_put_int64(GDataOutputStream*, int64_t, GCancellable*, GError**);
extern int g_data_output_stream_put_string(GDataOutputStream*, char*, GCancellable*, GError**);
extern int g_data_output_stream_put_uint16(GDataOutputStream*, uint16_t, GCancellable*, GError**);
extern int g_data_output_stream_put_uint32(GDataOutputStream*, uint32_t, GCancellable*, GError**);
extern int g_data_output_stream_put_uint64(GDataOutputStream*, uint64_t, GCancellable*, GError**);
extern void g_data_output_stream_set_byte_order(GDataOutputStream*, GDataStreamByteOrder);
extern GType g_data_output_stream_get_type();
extern GDesktopAppInfo* g_desktop_app_info_new(char*);
extern GDesktopAppInfo* g_desktop_app_info_new_from_filename(char*);
extern GDesktopAppInfo* g_desktop_app_info_new_from_keyfile(GKeyFile*);
extern void g_desktop_app_info_set_desktop_env(char*);
extern char* g_desktop_app_info_get_categories(GDesktopAppInfo*);
extern char* g_desktop_app_info_get_filename(GDesktopAppInfo*);
extern char* g_desktop_app_info_get_generic_name(GDesktopAppInfo*);
extern int g_desktop_app_info_get_is_hidden(GDesktopAppInfo*);
extern int g_desktop_app_info_get_nodisplay(GDesktopAppInfo*);
extern int g_desktop_app_info_get_show_in(GDesktopAppInfo*, char*);
extern int g_desktop_app_info_launch_uris_as_manager(GDesktopAppInfo*, GList*, GAppLaunchContext*, GSpawnFlags, GSpawnChildSetupFunc, void*, GDesktopAppLaunchCallback, void*, GError**);
#pragma GCC diagnostic ignored "-Wunused-function"
static int _g_desktop_app_info_launch_uris_as_manager(GDesktopAppInfo* this, GList* arg0, GAppLaunchContext* arg1, GSpawnFlags arg2, GSpawnChildSetupFunc arg3, void* arg4, void* gofunc, GError** arg7) {
	if (gofunc) {
		return g_desktop_app_info_launch_uris_as_manager(this, arg0, arg1, arg2, arg3, arg4, _GDesktopAppLaunchCallback_c_wrapper, gofunc, arg7);
	} else {
		return g_desktop_app_info_launch_uris_as_manager(this, arg0, arg1, arg2, arg3, arg4, 0, 0, arg7);
	}
}
extern GType g_desktop_app_info_get_type();
extern GAppInfo* g_desktop_app_info_lookup_get_default_for_uri_scheme(GDesktopAppInfoLookup*, char*);
extern GType g_desktop_app_info_lookup_get_type();
extern int g_drive_can_eject(GDrive*);
extern int g_drive_can_poll_for_media(GDrive*);
extern int g_drive_can_start(GDrive*);
extern int g_drive_can_start_degraded(GDrive*);
extern int g_drive_can_stop(GDrive*);
extern void g_drive_eject(GDrive*, GMountUnmountFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_drive_eject(GDrive* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_drive_eject(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_drive_eject(this, arg0, arg1, 0, 0);
	}
}
extern int g_drive_eject_finish(GDrive*, GAsyncResult*, GError**);
extern void g_drive_eject_with_operation(GDrive*, GMountUnmountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_drive_eject_with_operation(GDrive* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_drive_eject_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_drive_eject_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_drive_eject_with_operation_finish(GDrive*, GAsyncResult*, GError**);
extern char** g_drive_enumerate_identifiers(GDrive*);
extern GIcon* g_drive_get_icon(GDrive*);
extern char* g_drive_get_identifier(GDrive*, char*);
extern char* g_drive_get_name(GDrive*);
extern GDriveStartStopType g_drive_get_start_stop_type(GDrive*);
extern GList* g_drive_get_volumes(GDrive*);
extern int g_drive_has_media(GDrive*);
extern int g_drive_has_volumes(GDrive*);
extern int g_drive_is_media_check_automatic(GDrive*);
extern int g_drive_is_media_removable(GDrive*);
extern void g_drive_poll_for_media(GDrive*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_drive_poll_for_media(GDrive* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_drive_poll_for_media(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_drive_poll_for_media(this, arg0, 0, 0);
	}
}
extern int g_drive_poll_for_media_finish(GDrive*, GAsyncResult*, GError**);
extern void g_drive_start(GDrive*, GDriveStartFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_drive_start(GDrive* this, GDriveStartFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_drive_start(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_drive_start(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_drive_start_finish(GDrive*, GAsyncResult*, GError**);
extern void g_drive_stop(GDrive*, GMountUnmountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_drive_stop(GDrive* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_drive_stop(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_drive_stop(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_drive_stop_finish(GDrive*, GAsyncResult*, GError**);
extern GType g_drive_get_type();
extern GEmblem* g_emblem_new(GIcon*);
extern GEmblem* g_emblem_new_with_origin(GIcon*, GEmblemOrigin);
extern GIcon* g_emblem_get_icon(GEmblem*);
extern GEmblemOrigin g_emblem_get_origin(GEmblem*);
extern GType g_emblem_get_type();
extern GIcon* g_emblemed_icon_new(GIcon*, GEmblem*);
extern void g_emblemed_icon_add_emblem(GEmblemedIcon*, GEmblem*);
extern void g_emblemed_icon_clear_emblems(GEmblemedIcon*);
extern GList* g_emblemed_icon_get_emblems(GEmblemedIcon*);
extern GIcon* g_emblemed_icon_get_icon(GEmblemedIcon*);
extern GType g_emblemed_icon_get_type();
extern uint32_t g_file_hash(void*);
extern GFile* g_file_new_for_commandline_arg(char*);
extern GFile* g_file_new_for_path(char*);
extern GFile* g_file_new_for_uri(char*);
extern GFile* g_file_parse_name(char*);
extern GFileOutputStream* g_file_append_to(GFile*, GFileCreateFlags, GCancellable*, GError**);
extern void g_file_append_to_async(GFile*, GFileCreateFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_append_to_async(GFile* this, GFileCreateFlags arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_append_to_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_append_to_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GFileOutputStream* g_file_append_to_finish(GFile*, GAsyncResult*, GError**);
extern int g_file_copy(GFile*, GFile*, GFileCopyFlags, GCancellable*, GFileProgressCallback, void*, GError**);
#pragma GCC diagnostic ignored "-Wunused-function"
static int _g_file_copy(GFile* this, GFile* arg0, GFileCopyFlags arg1, GCancellable* arg2, void* gofunc, GError** arg5) {
	if (gofunc) {
		return g_file_copy(this, arg0, arg1, arg2, _GFileProgressCallback_c_wrapper, gofunc, arg5);
	} else {
		return g_file_copy(this, arg0, arg1, arg2, 0, 0, arg5);
	}
}
extern int g_file_copy_attributes(GFile*, GFile*, GFileCopyFlags, GCancellable*, GError**);
extern int g_file_copy_finish(GFile*, GAsyncResult*, GError**);
extern GFileOutputStream* g_file_create(GFile*, GFileCreateFlags, GCancellable*, GError**);
extern void g_file_create_async(GFile*, GFileCreateFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_create_async(GFile* this, GFileCreateFlags arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_create_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_create_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GFileOutputStream* g_file_create_finish(GFile*, GAsyncResult*, GError**);
extern GFileIOStream* g_file_create_readwrite(GFile*, GFileCreateFlags, GCancellable*, GError**);
extern void g_file_create_readwrite_async(GFile*, GFileCreateFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_create_readwrite_async(GFile* this, GFileCreateFlags arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_create_readwrite_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_create_readwrite_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GFileIOStream* g_file_create_readwrite_finish(GFile*, GAsyncResult*, GError**);
extern int g_file_delete(GFile*, GCancellable*, GError**);
extern GFile* g_file_dup(GFile*);
extern void g_file_eject_mountable(GFile*, GMountUnmountFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_eject_mountable(GFile* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_eject_mountable(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_eject_mountable(this, arg0, arg1, 0, 0);
	}
}
extern int g_file_eject_mountable_finish(GFile*, GAsyncResult*, GError**);
extern void g_file_eject_mountable_with_operation(GFile*, GMountUnmountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_eject_mountable_with_operation(GFile* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_eject_mountable_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_eject_mountable_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_file_eject_mountable_with_operation_finish(GFile*, GAsyncResult*, GError**);
extern GFileEnumerator* g_file_enumerate_children(GFile*, char*, GFileQueryInfoFlags, GCancellable*, GError**);
extern void g_file_enumerate_children_async(GFile*, char*, GFileQueryInfoFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_enumerate_children_async(GFile* this, char* arg0, GFileQueryInfoFlags arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_file_enumerate_children_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_enumerate_children_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern GFileEnumerator* g_file_enumerate_children_finish(GFile*, GAsyncResult*, GError**);
extern int g_file_equal(GFile*, GFile*);
extern GMount* g_file_find_enclosing_mount(GFile*, GCancellable*, GError**);
extern void g_file_find_enclosing_mount_async(GFile*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_find_enclosing_mount_async(GFile* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_find_enclosing_mount_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_find_enclosing_mount_async(this, arg0, arg1, 0, 0);
	}
}
extern GMount* g_file_find_enclosing_mount_finish(GFile*, GAsyncResult*, GError**);
extern char* g_file_get_basename(GFile*);
extern GFile* g_file_get_child(GFile*, char*);
extern GFile* g_file_get_child_for_display_name(GFile*, char*, GError**);
extern GFile* g_file_get_parent(GFile*);
extern char* g_file_get_parse_name(GFile*);
extern char* g_file_get_path(GFile*);
extern char* g_file_get_relative_path(GFile*, GFile*);
extern char* g_file_get_uri(GFile*);
extern char* g_file_get_uri_scheme(GFile*);
extern int g_file_has_parent(GFile*, GFile*);
extern int g_file_has_prefix(GFile*, GFile*);
extern int g_file_has_uri_scheme(GFile*, char*);
extern GIcon* g_file_icon_new(GFile*);
extern int g_file_is_native(GFile*);
extern int g_file_load_contents(GFile*, GCancellable*, uint8_t**, uint64_t*, char**, GError**);
extern void g_file_load_contents_async(GFile*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_load_contents_async(GFile* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_file_load_contents_async(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_load_contents_async(this, arg0, 0, 0);
	}
}
extern int g_file_load_contents_finish(GFile*, GAsyncResult*, uint8_t**, uint64_t*, char**, GError**);
extern int g_file_load_partial_contents_finish(GFile*, GAsyncResult*, uint8_t**, uint64_t*, char**, GError**);
extern int g_file_make_directory(GFile*, GCancellable*, GError**);
extern int g_file_make_directory_with_parents(GFile*, GCancellable*, GError**);
extern int g_file_make_symbolic_link(GFile*, char*, GCancellable*, GError**);
extern GFileMonitor* g_file_monitor(GFile*, GFileMonitorFlags, GCancellable*, GError**);
extern GFileMonitor* g_file_monitor_directory(GFile*, GFileMonitorFlags, GCancellable*, GError**);
extern GFileMonitor* g_file_monitor_file(GFile*, GFileMonitorFlags, GCancellable*, GError**);
extern void g_file_mount_enclosing_volume(GFile*, GMountMountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_mount_enclosing_volume(GFile* this, GMountMountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_mount_enclosing_volume(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_mount_enclosing_volume(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_file_mount_enclosing_volume_finish(GFile*, GAsyncResult*, GError**);
extern void g_file_mount_mountable(GFile*, GMountMountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_mount_mountable(GFile* this, GMountMountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_mount_mountable(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_mount_mountable(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GFile* g_file_mount_mountable_finish(GFile*, GAsyncResult*, GError**);
extern int g_file_move(GFile*, GFile*, GFileCopyFlags, GCancellable*, GFileProgressCallback, void*, GError**);
#pragma GCC diagnostic ignored "-Wunused-function"
static int _g_file_move(GFile* this, GFile* arg0, GFileCopyFlags arg1, GCancellable* arg2, void* gofunc, GError** arg5) {
	if (gofunc) {
		return g_file_move(this, arg0, arg1, arg2, _GFileProgressCallback_c_wrapper, gofunc, arg5);
	} else {
		return g_file_move(this, arg0, arg1, arg2, 0, 0, arg5);
	}
}
extern GFileIOStream* g_file_open_readwrite(GFile*, GCancellable*, GError**);
extern void g_file_open_readwrite_async(GFile*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_open_readwrite_async(GFile* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_open_readwrite_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_open_readwrite_async(this, arg0, arg1, 0, 0);
	}
}
extern GFileIOStream* g_file_open_readwrite_finish(GFile*, GAsyncResult*, GError**);
extern void g_file_poll_mountable(GFile*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_poll_mountable(GFile* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_file_poll_mountable(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_poll_mountable(this, arg0, 0, 0);
	}
}
extern int g_file_poll_mountable_finish(GFile*, GAsyncResult*, GError**);
extern GAppInfo* g_file_query_default_handler(GFile*, GCancellable*, GError**);
extern int g_file_query_exists(GFile*, GCancellable*);
extern GFileType g_file_query_file_type(GFile*, GFileQueryInfoFlags, GCancellable*);
extern GFileInfo* g_file_query_filesystem_info(GFile*, char*, GCancellable*, GError**);
extern void g_file_query_filesystem_info_async(GFile*, char*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_query_filesystem_info_async(GFile* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_query_filesystem_info_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_query_filesystem_info_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GFileInfo* g_file_query_filesystem_info_finish(GFile*, GAsyncResult*, GError**);
extern GFileInfo* g_file_query_info(GFile*, char*, GFileQueryInfoFlags, GCancellable*, GError**);
extern void g_file_query_info_async(GFile*, char*, GFileQueryInfoFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_query_info_async(GFile* this, char* arg0, GFileQueryInfoFlags arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_file_query_info_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_query_info_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern GFileInfo* g_file_query_info_finish(GFile*, GAsyncResult*, GError**);
extern GFileAttributeInfoList* g_file_query_settable_attributes(GFile*, GCancellable*, GError**);
extern GFileAttributeInfoList* g_file_query_writable_namespaces(GFile*, GCancellable*, GError**);
extern GFileInputStream* g_file_read(GFile*, GCancellable*, GError**);
extern void g_file_read_async(GFile*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_read_async(GFile* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_read_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_read_async(this, arg0, arg1, 0, 0);
	}
}
extern GFileInputStream* g_file_read_finish(GFile*, GAsyncResult*, GError**);
extern GFileOutputStream* g_file_replace(GFile*, char*, int, GFileCreateFlags, GCancellable*, GError**);
extern void g_file_replace_async(GFile*, char*, int, GFileCreateFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_replace_async(GFile* this, char* arg0, int arg1, GFileCreateFlags arg2, int32_t arg3, GCancellable* arg4, void* gofunc) {
	if (gofunc) {
		g_file_replace_async(this, arg0, arg1, arg2, arg3, arg4, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_replace_async(this, arg0, arg1, arg2, arg3, arg4, 0, 0);
	}
}
extern int g_file_replace_contents(GFile*, uint8_t*, uint64_t, char*, int, GFileCreateFlags, char**, GCancellable*, GError**);
extern void g_file_replace_contents_async(GFile*, uint8_t*, uint64_t, char*, int, GFileCreateFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_replace_contents_async(GFile* this, uint8_t* arg0, uint64_t arg1, char* arg2, int arg3, GFileCreateFlags arg4, GCancellable* arg5, void* gofunc) {
	if (gofunc) {
		g_file_replace_contents_async(this, arg0, arg1, arg2, arg3, arg4, arg5, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_replace_contents_async(this, arg0, arg1, arg2, arg3, arg4, arg5, 0, 0);
	}
}
extern int g_file_replace_contents_finish(GFile*, GAsyncResult*, char**, GError**);
extern GFileOutputStream* g_file_replace_finish(GFile*, GAsyncResult*, GError**);
extern GFileIOStream* g_file_replace_readwrite(GFile*, char*, int, GFileCreateFlags, GCancellable*, GError**);
extern void g_file_replace_readwrite_async(GFile*, char*, int, GFileCreateFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_replace_readwrite_async(GFile* this, char* arg0, int arg1, GFileCreateFlags arg2, int32_t arg3, GCancellable* arg4, void* gofunc) {
	if (gofunc) {
		g_file_replace_readwrite_async(this, arg0, arg1, arg2, arg3, arg4, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_replace_readwrite_async(this, arg0, arg1, arg2, arg3, arg4, 0, 0);
	}
}
extern GFileIOStream* g_file_replace_readwrite_finish(GFile*, GAsyncResult*, GError**);
extern GFile* g_file_resolve_relative_path(GFile*, char*);
extern int g_file_set_attribute(GFile*, char*, GFileAttributeType, void*, GFileQueryInfoFlags, GCancellable*, GError**);
extern int g_file_set_attribute_byte_string(GFile*, char*, char*, GFileQueryInfoFlags, GCancellable*, GError**);
extern int g_file_set_attribute_int32(GFile*, char*, int32_t, GFileQueryInfoFlags, GCancellable*, GError**);
extern int g_file_set_attribute_int64(GFile*, char*, int64_t, GFileQueryInfoFlags, GCancellable*, GError**);
extern int g_file_set_attribute_string(GFile*, char*, char*, GFileQueryInfoFlags, GCancellable*, GError**);
extern int g_file_set_attribute_uint32(GFile*, char*, uint32_t, GFileQueryInfoFlags, GCancellable*, GError**);
extern int g_file_set_attribute_uint64(GFile*, char*, uint64_t, GFileQueryInfoFlags, GCancellable*, GError**);
extern void g_file_set_attributes_async(GFile*, GFileInfo*, GFileQueryInfoFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_set_attributes_async(GFile* this, GFileInfo* arg0, GFileQueryInfoFlags arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_file_set_attributes_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_set_attributes_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern int g_file_set_attributes_finish(GFile*, GAsyncResult*, GFileInfo**, GError**);
extern int g_file_set_attributes_from_info(GFile*, GFileInfo*, GFileQueryInfoFlags, GCancellable*, GError**);
extern GFile* g_file_set_display_name(GFile*, char*, GCancellable*, GError**);
extern void g_file_set_display_name_async(GFile*, char*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_set_display_name_async(GFile* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_set_display_name_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_set_display_name_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GFile* g_file_set_display_name_finish(GFile*, GAsyncResult*, GError**);
extern void g_file_start_mountable(GFile*, GDriveStartFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_start_mountable(GFile* this, GDriveStartFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_start_mountable(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_start_mountable(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_file_start_mountable_finish(GFile*, GAsyncResult*, GError**);
extern void g_file_stop_mountable(GFile*, GMountUnmountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_stop_mountable(GFile* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_stop_mountable(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_stop_mountable(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_file_stop_mountable_finish(GFile*, GAsyncResult*, GError**);
extern int g_file_supports_thread_contexts(GFile*);
extern int g_file_trash(GFile*, GCancellable*, GError**);
extern void g_file_unmount_mountable(GFile*, GMountUnmountFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_unmount_mountable(GFile* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_unmount_mountable(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_unmount_mountable(this, arg0, arg1, 0, 0);
	}
}
extern int g_file_unmount_mountable_finish(GFile*, GAsyncResult*, GError**);
extern void g_file_unmount_mountable_with_operation(GFile*, GMountUnmountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_unmount_mountable_with_operation(GFile* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_unmount_mountable_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_unmount_mountable_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_file_unmount_mountable_with_operation_finish(GFile*, GAsyncResult*, GError**);
extern GType g_file_get_type();
extern GFileAttributeInfoList* g_file_attribute_info_list_new();
extern void g_file_attribute_info_list_add(GFileAttributeInfoList*, char*, GFileAttributeType, GFileAttributeInfoFlags);
extern GFileAttributeInfoList* g_file_attribute_info_list_dup(GFileAttributeInfoList*);
extern GFileAttributeInfo* g_file_attribute_info_list_lookup(GFileAttributeInfoList*, char*);
extern GFileAttributeInfoList* g_file_attribute_info_list_ref(GFileAttributeInfoList*);
extern void g_file_attribute_info_list_unref(GFileAttributeInfoList*);
extern GFileAttributeMatcher* g_file_attribute_matcher_new(char*);
extern int g_file_attribute_matcher_enumerate_namespace(GFileAttributeMatcher*, char*);
extern char* g_file_attribute_matcher_enumerate_next(GFileAttributeMatcher*);
extern int g_file_attribute_matcher_matches(GFileAttributeMatcher*, char*);
extern int g_file_attribute_matcher_matches_only(GFileAttributeMatcher*, char*);
extern GFileAttributeMatcher* g_file_attribute_matcher_ref(GFileAttributeMatcher*);
extern void g_file_attribute_matcher_unref(GFileAttributeMatcher*);
extern int32_t g_file_descriptor_based_get_fd(GFileDescriptorBased*);
extern GType g_file_descriptor_based_get_type();
extern int g_file_enumerator_close(GFileEnumerator*, GCancellable*, GError**);
extern void g_file_enumerator_close_async(GFileEnumerator*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_enumerator_close_async(GFileEnumerator* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_file_enumerator_close_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_enumerator_close_async(this, arg0, arg1, 0, 0);
	}
}
extern int g_file_enumerator_close_finish(GFileEnumerator*, GAsyncResult*, GError**);
extern GFile* g_file_enumerator_get_container(GFileEnumerator*);
extern int g_file_enumerator_has_pending(GFileEnumerator*);
extern int g_file_enumerator_is_closed(GFileEnumerator*);
extern GFileInfo* g_file_enumerator_next_file(GFileEnumerator*, GCancellable*, GError**);
extern void g_file_enumerator_next_files_async(GFileEnumerator*, int32_t, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_enumerator_next_files_async(GFileEnumerator* this, int32_t arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_enumerator_next_files_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_enumerator_next_files_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GList* g_file_enumerator_next_files_finish(GFileEnumerator*, GAsyncResult*, GError**);
extern void g_file_enumerator_set_pending(GFileEnumerator*, int);
extern GType g_file_enumerator_get_type();
extern char* g_file_io_stream_get_etag(GFileIOStream*);
extern GFileInfo* g_file_io_stream_query_info(GFileIOStream*, char*, GCancellable*, GError**);
extern void g_file_io_stream_query_info_async(GFileIOStream*, char*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_io_stream_query_info_async(GFileIOStream* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_io_stream_query_info_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_io_stream_query_info_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GFileInfo* g_file_io_stream_query_info_finish(GFileIOStream*, GAsyncResult*, GError**);
extern GType g_file_io_stream_get_type();
extern GFile* g_file_icon_get_file(GFileIcon*);
extern GType g_file_icon_get_type();
extern GFileInfo* g_file_info_new();
extern void g_file_info_clear_status(GFileInfo*);
extern void g_file_info_copy_into(GFileInfo*, GFileInfo*);
extern GFileInfo* g_file_info_dup(GFileInfo*);
extern char* g_file_info_get_attribute_as_string(GFileInfo*, char*);
extern int g_file_info_get_attribute_boolean(GFileInfo*, char*);
extern char* g_file_info_get_attribute_byte_string(GFileInfo*, char*);
extern int g_file_info_get_attribute_data(GFileInfo*, char*, GFileAttributeType*, void**, GFileAttributeStatus*);
extern int32_t g_file_info_get_attribute_int32(GFileInfo*, char*);
extern int64_t g_file_info_get_attribute_int64(GFileInfo*, char*);
extern GObject* g_file_info_get_attribute_object(GFileInfo*, char*);
extern GFileAttributeStatus g_file_info_get_attribute_status(GFileInfo*, char*);
extern char* g_file_info_get_attribute_string(GFileInfo*, char*);
extern char** g_file_info_get_attribute_stringv(GFileInfo*, char*);
extern GFileAttributeType g_file_info_get_attribute_type(GFileInfo*, char*);
extern uint32_t g_file_info_get_attribute_uint32(GFileInfo*, char*);
extern uint64_t g_file_info_get_attribute_uint64(GFileInfo*, char*);
extern char* g_file_info_get_content_type(GFileInfo*);
extern char* g_file_info_get_display_name(GFileInfo*);
extern char* g_file_info_get_edit_name(GFileInfo*);
extern char* g_file_info_get_etag(GFileInfo*);
extern GFileType g_file_info_get_file_type(GFileInfo*);
extern GIcon* g_file_info_get_icon(GFileInfo*);
extern int g_file_info_get_is_backup(GFileInfo*);
extern int g_file_info_get_is_hidden(GFileInfo*);
extern int g_file_info_get_is_symlink(GFileInfo*);
extern void g_file_info_get_modification_time(GFileInfo*, GTimeVal*);
extern char* g_file_info_get_name(GFileInfo*);
extern int64_t g_file_info_get_size(GFileInfo*);
extern int32_t g_file_info_get_sort_order(GFileInfo*);
extern char* g_file_info_get_symlink_target(GFileInfo*);
extern int g_file_info_has_attribute(GFileInfo*, char*);
extern int g_file_info_has_namespace(GFileInfo*, char*);
extern char** g_file_info_list_attributes(GFileInfo*, char*);
extern void g_file_info_remove_attribute(GFileInfo*, char*);
extern void g_file_info_set_attribute(GFileInfo*, char*, GFileAttributeType, void*);
extern void g_file_info_set_attribute_boolean(GFileInfo*, char*, int);
extern void g_file_info_set_attribute_byte_string(GFileInfo*, char*, char*);
extern void g_file_info_set_attribute_int32(GFileInfo*, char*, int32_t);
extern void g_file_info_set_attribute_int64(GFileInfo*, char*, int64_t);
extern void g_file_info_set_attribute_mask(GFileInfo*, GFileAttributeMatcher*);
extern void g_file_info_set_attribute_object(GFileInfo*, char*, GObject*);
extern int g_file_info_set_attribute_status(GFileInfo*, char*, GFileAttributeStatus);
extern void g_file_info_set_attribute_string(GFileInfo*, char*, char*);
extern void g_file_info_set_attribute_stringv(GFileInfo*, char*, char*);
extern void g_file_info_set_attribute_uint32(GFileInfo*, char*, uint32_t);
extern void g_file_info_set_attribute_uint64(GFileInfo*, char*, uint64_t);
extern void g_file_info_set_content_type(GFileInfo*, char*);
extern void g_file_info_set_display_name(GFileInfo*, char*);
extern void g_file_info_set_edit_name(GFileInfo*, char*);
extern void g_file_info_set_file_type(GFileInfo*, GFileType);
extern void g_file_info_set_icon(GFileInfo*, GIcon*);
extern void g_file_info_set_is_hidden(GFileInfo*, int);
extern void g_file_info_set_is_symlink(GFileInfo*, int);
extern void g_file_info_set_modification_time(GFileInfo*, GTimeVal*);
extern void g_file_info_set_name(GFileInfo*, char*);
extern void g_file_info_set_size(GFileInfo*, int64_t);
extern void g_file_info_set_sort_order(GFileInfo*, int32_t);
extern void g_file_info_set_symlink_target(GFileInfo*, char*);
extern void g_file_info_unset_attribute_mask(GFileInfo*);
extern GType g_file_info_get_type();
extern GFileInfo* g_file_input_stream_query_info(GFileInputStream*, char*, GCancellable*, GError**);
extern void g_file_input_stream_query_info_async(GFileInputStream*, char*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_input_stream_query_info_async(GFileInputStream* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_input_stream_query_info_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_input_stream_query_info_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GFileInfo* g_file_input_stream_query_info_finish(GFileInputStream*, GAsyncResult*, GError**);
extern GType g_file_input_stream_get_type();
extern int g_file_monitor_cancel(GFileMonitor*);
extern void g_file_monitor_emit_event(GFileMonitor*, GFile*, GFile*, GFileMonitorEvent);
extern int g_file_monitor_is_cancelled(GFileMonitor*);
extern void g_file_monitor_set_rate_limit(GFileMonitor*, int32_t);
extern GType g_file_monitor_get_type();
extern char* g_file_output_stream_get_etag(GFileOutputStream*);
extern GFileInfo* g_file_output_stream_query_info(GFileOutputStream*, char*, GCancellable*, GError**);
extern void g_file_output_stream_query_info_async(GFileOutputStream*, char*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_file_output_stream_query_info_async(GFileOutputStream* this, char* arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_file_output_stream_query_info_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_file_output_stream_query_info_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GFileInfo* g_file_output_stream_query_info_finish(GFileOutputStream*, GAsyncResult*, GError**);
extern GType g_file_output_stream_get_type();
extern GFilenameCompleter* g_filename_completer_new();
extern char* g_filename_completer_get_completion_suffix(GFilenameCompleter*, char*);
extern char** g_filename_completer_get_completions(GFilenameCompleter*, char*);
extern void g_filename_completer_set_dirs_only(GFilenameCompleter*, int);
extern GType g_filename_completer_get_type();
extern GInputStream* g_filter_input_stream_get_base_stream(GFilterInputStream*);
extern int g_filter_input_stream_get_close_base_stream(GFilterInputStream*);
extern void g_filter_input_stream_set_close_base_stream(GFilterInputStream*, int);
extern GType g_filter_input_stream_get_type();
extern GOutputStream* g_filter_output_stream_get_base_stream(GFilterOutputStream*);
extern int g_filter_output_stream_get_close_base_stream(GFilterOutputStream*);
extern void g_filter_output_stream_set_close_base_stream(GFilterOutputStream*, int);
extern GType g_filter_output_stream_get_type();
extern char* g_io_extension_get_name(GIOExtension*);
extern int32_t g_io_extension_get_priority(GIOExtension*);
extern GIOExtension* g_io_extension_point_get_extension_by_name(GIOExtensionPoint*, char*);
extern GList* g_io_extension_point_get_extensions(GIOExtensionPoint*);
extern GType g_io_extension_point_get_required_type(GIOExtensionPoint*);
extern void g_io_extension_point_set_required_type(GIOExtensionPoint*, GType);
extern GIOExtension* g_io_extension_point_implement(char*, GType, char*, int32_t);
extern GIOExtensionPoint* g_io_extension_point_lookup(char*);
extern GIOExtensionPoint* g_io_extension_point_register(char*);
extern GIOModule* g_io_module_new(char*);
extern char** g_io_module_query();
extern void g_io_module_load(GIOModule*);
extern void g_io_module_unload(GIOModule*);
extern GType g_io_module_get_type();
extern void g_io_module_scope_block(GIOModuleScope*, char*);
extern void g_io_module_scope_free(GIOModuleScope*);
extern int g_io_scheduler_job_send_to_mainloop(GIOSchedulerJob*, GSourceFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static int _g_io_scheduler_job_send_to_mainloop(GIOSchedulerJob* this, void* gofunc) {
	if (gofunc) {
		return g_io_scheduler_job_send_to_mainloop(this, _GSourceFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return g_io_scheduler_job_send_to_mainloop(this, 0, 0, 0);
	}
}
extern void g_io_scheduler_job_send_to_mainloop_async(GIOSchedulerJob*, GSourceFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_io_scheduler_job_send_to_mainloop_async(GIOSchedulerJob* this, void* gofunc) {
	if (gofunc) {
		g_io_scheduler_job_send_to_mainloop_async(this, _GSourceFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		g_io_scheduler_job_send_to_mainloop_async(this, 0, 0, 0);
	}
}
extern int g_io_stream_splice_finish(GAsyncResult*, GError**);
extern void g_io_stream_clear_pending(GIOStream*);
extern int g_io_stream_close(GIOStream*, GCancellable*, GError**);
extern void g_io_stream_close_async(GIOStream*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_io_stream_close_async(GIOStream* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_io_stream_close_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_io_stream_close_async(this, arg0, arg1, 0, 0);
	}
}
extern int g_io_stream_close_finish(GIOStream*, GAsyncResult*, GError**);
extern GInputStream* g_io_stream_get_input_stream(GIOStream*);
extern GOutputStream* g_io_stream_get_output_stream(GIOStream*);
extern int g_io_stream_has_pending(GIOStream*);
extern int g_io_stream_is_closed(GIOStream*);
extern int g_io_stream_set_pending(GIOStream*, GError**);
extern void g_io_stream_splice_async(GIOStream*, GIOStream*, GIOStreamSpliceFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_io_stream_splice_async(GIOStream* this, GIOStream* arg0, GIOStreamSpliceFlags arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_io_stream_splice_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_io_stream_splice_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern GType g_io_stream_get_type();
extern uint32_t g_icon_hash(void*);
extern GIcon* g_icon_new_for_string(char*, GError**);
extern int g_icon_equal(GIcon*, GIcon*);
extern char* g_icon_to_string(GIcon*);
extern GType g_icon_get_type();
extern GInetAddress* g_inet_address_new_any(GSocketFamily);
extern GInetAddress* g_inet_address_new_from_bytes(uint8_t*, GSocketFamily);
extern GInetAddress* g_inet_address_new_from_string(char*);
extern GInetAddress* g_inet_address_new_loopback(GSocketFamily);
extern int g_inet_address_equal(GInetAddress*, GInetAddress*);
extern GSocketFamily g_inet_address_get_family(GInetAddress*);
extern int g_inet_address_get_is_any(GInetAddress*);
extern int g_inet_address_get_is_link_local(GInetAddress*);
extern int g_inet_address_get_is_loopback(GInetAddress*);
extern int g_inet_address_get_is_mc_global(GInetAddress*);
extern int g_inet_address_get_is_mc_link_local(GInetAddress*);
extern int g_inet_address_get_is_mc_node_local(GInetAddress*);
extern int g_inet_address_get_is_mc_org_local(GInetAddress*);
extern int g_inet_address_get_is_mc_site_local(GInetAddress*);
extern int g_inet_address_get_is_multicast(GInetAddress*);
extern int g_inet_address_get_is_site_local(GInetAddress*);
extern uint64_t g_inet_address_get_native_size(GInetAddress*);
extern char* g_inet_address_to_string(GInetAddress*);
extern GType g_inet_address_get_type();
extern GSocketAddress* g_inet_socket_address_new(GInetAddress*, uint16_t);
extern GInetAddress* g_inet_socket_address_get_address(GInetSocketAddress*);
extern uint16_t g_inet_socket_address_get_port(GInetSocketAddress*);
extern GType g_inet_socket_address_get_type();
extern void* g_initable_newv(GType, uint32_t, GParameter*, GCancellable*, GError**);
extern int g_initable_init(GInitable*, GCancellable*, GError**);
extern GType g_initable_get_type();
extern void g_input_stream_clear_pending(GInputStream*);
extern int g_input_stream_close(GInputStream*, GCancellable*, GError**);
extern void g_input_stream_close_async(GInputStream*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_input_stream_close_async(GInputStream* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_input_stream_close_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_input_stream_close_async(this, arg0, arg1, 0, 0);
	}
}
extern int g_input_stream_close_finish(GInputStream*, GAsyncResult*, GError**);
extern int g_input_stream_has_pending(GInputStream*);
extern int g_input_stream_is_closed(GInputStream*);
extern int64_t g_input_stream_read(GInputStream*, void*, uint64_t, GCancellable*, GError**);
extern int g_input_stream_read_all(GInputStream*, void*, uint64_t, uint64_t*, GCancellable*, GError**);
extern void g_input_stream_read_async(GInputStream*, void*, uint64_t, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_input_stream_read_async(GInputStream* this, void* arg0, uint64_t arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_input_stream_read_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_input_stream_read_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern int64_t g_input_stream_read_finish(GInputStream*, GAsyncResult*, GError**);
extern int g_input_stream_set_pending(GInputStream*, GError**);
extern int64_t g_input_stream_skip(GInputStream*, uint64_t, GCancellable*, GError**);
extern void g_input_stream_skip_async(GInputStream*, uint64_t, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_input_stream_skip_async(GInputStream* this, uint64_t arg0, int32_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_input_stream_skip_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_input_stream_skip_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int64_t g_input_stream_skip_finish(GInputStream*, GAsyncResult*, GError**);
extern GType g_input_stream_get_type();
extern GInputStream* g_loadable_icon_load(GLoadableIcon*, int32_t, char**, GCancellable*, GError**);
extern void g_loadable_icon_load_async(GLoadableIcon*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_loadable_icon_load_async(GLoadableIcon* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_loadable_icon_load_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_loadable_icon_load_async(this, arg0, arg1, 0, 0);
	}
}
extern GInputStream* g_loadable_icon_load_finish(GLoadableIcon*, GAsyncResult*, char*, GError**);
extern GType g_loadable_icon_get_type();
extern GInputStream* g_memory_input_stream_new();
extern GInputStream* g_memory_input_stream_new_from_data(uint8_t*, int64_t, GDestroyNotify);
extern void g_memory_input_stream_add_data(GMemoryInputStream*, uint8_t*, int64_t, GDestroyNotify);
extern GType g_memory_input_stream_get_type();
extern void* g_memory_output_stream_get_data(GMemoryOutputStream*);
extern uint64_t g_memory_output_stream_get_data_size(GMemoryOutputStream*);
extern uint64_t g_memory_output_stream_get_size(GMemoryOutputStream*);
extern void* g_memory_output_stream_steal_data(GMemoryOutputStream*);
extern GType g_memory_output_stream_get_type();
extern int g_mount_can_eject(GMount*);
extern int g_mount_can_unmount(GMount*);
extern void g_mount_eject(GMount*, GMountUnmountFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_mount_eject(GMount* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_mount_eject(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_eject(this, arg0, arg1, 0, 0);
	}
}
extern int g_mount_eject_finish(GMount*, GAsyncResult*, GError**);
extern void g_mount_eject_with_operation(GMount*, GMountUnmountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_mount_eject_with_operation(GMount* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_mount_eject_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_eject_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_mount_eject_with_operation_finish(GMount*, GAsyncResult*, GError**);
extern GFile* g_mount_get_default_location(GMount*);
extern GDrive* g_mount_get_drive(GMount*);
extern GIcon* g_mount_get_icon(GMount*);
extern char* g_mount_get_name(GMount*);
extern GFile* g_mount_get_root(GMount*);
extern char* g_mount_get_uuid(GMount*);
extern GVolume* g_mount_get_volume(GMount*);
extern void g_mount_guess_content_type(GMount*, int, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_mount_guess_content_type(GMount* this, int arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_mount_guess_content_type(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_guess_content_type(this, arg0, arg1, 0, 0);
	}
}
extern char** g_mount_guess_content_type_finish(GMount*, GAsyncResult*, GError**);
extern char** g_mount_guess_content_type_sync(GMount*, int, GCancellable*, GError**);
extern int g_mount_is_shadowed(GMount*);
extern void g_mount_remount(GMount*, GMountMountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_mount_remount(GMount* this, GMountMountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_mount_remount(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_remount(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_mount_remount_finish(GMount*, GAsyncResult*, GError**);
extern void g_mount_shadow(GMount*);
extern void g_mount_unmount(GMount*, GMountUnmountFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_mount_unmount(GMount* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_mount_unmount(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_unmount(this, arg0, arg1, 0, 0);
	}
}
extern int g_mount_unmount_finish(GMount*, GAsyncResult*, GError**);
extern void g_mount_unmount_with_operation(GMount*, GMountUnmountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_mount_unmount_with_operation(GMount* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_mount_unmount_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_mount_unmount_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_mount_unmount_with_operation_finish(GMount*, GAsyncResult*, GError**);
extern void g_mount_unshadow(GMount*);
extern GType g_mount_get_type();
extern GMountOperation* g_mount_operation_new();
extern int g_mount_operation_get_anonymous(GMountOperation*);
extern int32_t g_mount_operation_get_choice(GMountOperation*);
extern char* g_mount_operation_get_domain(GMountOperation*);
extern char* g_mount_operation_get_password(GMountOperation*);
extern GPasswordSave g_mount_operation_get_password_save(GMountOperation*);
extern char* g_mount_operation_get_username(GMountOperation*);
extern void g_mount_operation_reply(GMountOperation*, GMountOperationResult);
extern void g_mount_operation_set_anonymous(GMountOperation*, int);
extern void g_mount_operation_set_choice(GMountOperation*, int32_t);
extern void g_mount_operation_set_domain(GMountOperation*, char*);
extern void g_mount_operation_set_password(GMountOperation*, char*);
extern void g_mount_operation_set_password_save(GMountOperation*, GPasswordSave);
extern void g_mount_operation_set_username(GMountOperation*, char*);
extern GType g_mount_operation_get_type();
extern GType g_native_volume_monitor_get_type();
extern GSocketConnectable* g_network_address_new(char*, uint16_t);
extern GSocketConnectable* g_network_address_parse(char*, uint16_t, GError**);
extern GSocketConnectable* g_network_address_parse_uri(char*, uint16_t, GError**);
extern char* g_network_address_get_hostname(GNetworkAddress*);
extern uint16_t g_network_address_get_port(GNetworkAddress*);
extern char* g_network_address_get_scheme(GNetworkAddress*);
extern GType g_network_address_get_type();
extern GSocketConnectable* g_network_service_new(char*, char*, char*);
extern char* g_network_service_get_domain(GNetworkService*);
extern char* g_network_service_get_protocol(GNetworkService*);
extern char* g_network_service_get_scheme(GNetworkService*);
extern char* g_network_service_get_service(GNetworkService*);
extern void g_network_service_set_scheme(GNetworkService*, char*);
extern GType g_network_service_get_type();
extern void g_output_stream_clear_pending(GOutputStream*);
extern int g_output_stream_close(GOutputStream*, GCancellable*, GError**);
extern void g_output_stream_close_async(GOutputStream*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_output_stream_close_async(GOutputStream* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_output_stream_close_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_output_stream_close_async(this, arg0, arg1, 0, 0);
	}
}
extern int g_output_stream_close_finish(GOutputStream*, GAsyncResult*, GError**);
extern int g_output_stream_flush(GOutputStream*, GCancellable*, GError**);
extern void g_output_stream_flush_async(GOutputStream*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_output_stream_flush_async(GOutputStream* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_output_stream_flush_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_output_stream_flush_async(this, arg0, arg1, 0, 0);
	}
}
extern int g_output_stream_flush_finish(GOutputStream*, GAsyncResult*, GError**);
extern int g_output_stream_has_pending(GOutputStream*);
extern int g_output_stream_is_closed(GOutputStream*);
extern int g_output_stream_is_closing(GOutputStream*);
extern int g_output_stream_set_pending(GOutputStream*, GError**);
extern int64_t g_output_stream_splice(GOutputStream*, GInputStream*, GOutputStreamSpliceFlags, GCancellable*, GError**);
extern void g_output_stream_splice_async(GOutputStream*, GInputStream*, GOutputStreamSpliceFlags, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_output_stream_splice_async(GOutputStream* this, GInputStream* arg0, GOutputStreamSpliceFlags arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_output_stream_splice_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_output_stream_splice_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern int64_t g_output_stream_splice_finish(GOutputStream*, GAsyncResult*, GError**);
extern int64_t g_output_stream_write(GOutputStream*, uint8_t*, uint64_t, GCancellable*, GError**);
extern int g_output_stream_write_all(GOutputStream*, uint8_t*, uint64_t, uint64_t*, GCancellable*, GError**);
extern void g_output_stream_write_async(GOutputStream*, uint8_t*, uint64_t, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_output_stream_write_async(GOutputStream* this, uint8_t* arg0, uint64_t arg1, int32_t arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_output_stream_write_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_output_stream_write_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern int64_t g_output_stream_write_finish(GOutputStream*, GAsyncResult*, GError**);
extern GType g_output_stream_get_type();
extern int g_permission_acquire(GPermission*, GCancellable*, GError**);
extern void g_permission_acquire_async(GPermission*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_permission_acquire_async(GPermission* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_permission_acquire_async(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_permission_acquire_async(this, arg0, 0, 0);
	}
}
extern int g_permission_acquire_finish(GPermission*, GAsyncResult*, GError**);
extern int g_permission_get_allowed(GPermission*);
extern int g_permission_get_can_acquire(GPermission*);
extern int g_permission_get_can_release(GPermission*);
extern void g_permission_impl_update(GPermission*, int, int, int);
extern int g_permission_release(GPermission*, GCancellable*, GError**);
extern void g_permission_release_async(GPermission*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_permission_release_async(GPermission* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_permission_release_async(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_permission_release_async(this, arg0, 0, 0);
	}
}
extern int g_permission_release_finish(GPermission*, GAsyncResult*, GError**);
extern GType g_permission_get_type();
extern int g_pollable_input_stream_can_poll(GPollableInputStream*);
extern int g_pollable_input_stream_is_readable(GPollableInputStream*);
extern int64_t g_pollable_input_stream_read_nonblocking(GPollableInputStream*, void*, uint64_t, GCancellable*, GError**);
extern GType g_pollable_input_stream_get_type();
extern int g_pollable_output_stream_can_poll(GPollableOutputStream*);
extern int g_pollable_output_stream_is_writable(GPollableOutputStream*);
extern int64_t g_pollable_output_stream_write_nonblocking(GPollableOutputStream*, uint8_t*, uint64_t, GCancellable*, GError**);
extern GType g_pollable_output_stream_get_type();
extern GProxy* g_proxy_get_default_for_protocol(char*);
extern GIOStream* g_proxy_connect(GProxy*, GIOStream*, GProxyAddress*, GCancellable*, GError**);
extern void g_proxy_connect_async(GProxy*, GIOStream*, GProxyAddress*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_proxy_connect_async(GProxy* this, GIOStream* arg0, GProxyAddress* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_proxy_connect_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_proxy_connect_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GIOStream* g_proxy_connect_finish(GProxy*, GAsyncResult*, GError**);
extern int g_proxy_supports_hostname(GProxy*);
extern GType g_proxy_get_type();
extern GSocketAddress* g_proxy_address_new(GInetAddress*, uint16_t, char*, char*, uint16_t, char*, char*);
extern char* g_proxy_address_get_destination_hostname(GProxyAddress*);
extern uint16_t g_proxy_address_get_destination_port(GProxyAddress*);
extern char* g_proxy_address_get_password(GProxyAddress*);
extern char* g_proxy_address_get_protocol(GProxyAddress*);
extern char* g_proxy_address_get_username(GProxyAddress*);
extern GType g_proxy_address_get_type();
extern GType g_proxy_address_enumerator_get_type();
extern GProxyResolver* g_proxy_resolver_get_default();
extern int g_proxy_resolver_is_supported(GProxyResolver*);
extern char** g_proxy_resolver_lookup(GProxyResolver*, char*, GCancellable*, GError**);
extern void g_proxy_resolver_lookup_async(GProxyResolver*, char*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_proxy_resolver_lookup_async(GProxyResolver* this, char* arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_proxy_resolver_lookup_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_proxy_resolver_lookup_async(this, arg0, arg1, 0, 0);
	}
}
extern char** g_proxy_resolver_lookup_finish(GProxyResolver*, GAsyncResult*, GError**);
extern GType g_proxy_resolver_get_type();
extern GResolver* g_resolver_get_default();
extern char* g_resolver_lookup_by_address(GResolver*, GInetAddress*, GCancellable*, GError**);
extern void g_resolver_lookup_by_address_async(GResolver*, GInetAddress*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_resolver_lookup_by_address_async(GResolver* this, GInetAddress* arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_resolver_lookup_by_address_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_resolver_lookup_by_address_async(this, arg0, arg1, 0, 0);
	}
}
extern char* g_resolver_lookup_by_address_finish(GResolver*, GAsyncResult*, GError**);
extern GList* g_resolver_lookup_by_name(GResolver*, char*, GCancellable*, GError**);
extern void g_resolver_lookup_by_name_async(GResolver*, char*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_resolver_lookup_by_name_async(GResolver* this, char* arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_resolver_lookup_by_name_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_resolver_lookup_by_name_async(this, arg0, arg1, 0, 0);
	}
}
extern GList* g_resolver_lookup_by_name_finish(GResolver*, GAsyncResult*, GError**);
extern GList* g_resolver_lookup_service(GResolver*, char*, char*, char*, GCancellable*, GError**);
extern void g_resolver_lookup_service_async(GResolver*, char*, char*, char*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_resolver_lookup_service_async(GResolver* this, char* arg0, char* arg1, char* arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_resolver_lookup_service_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_resolver_lookup_service_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern GList* g_resolver_lookup_service_finish(GResolver*, GAsyncResult*, GError**);
extern void g_resolver_set_default(GResolver*);
extern GType g_resolver_get_type();
extern int g_seekable_can_seek(GSeekable*);
extern int g_seekable_can_truncate(GSeekable*);
extern int g_seekable_seek(GSeekable*, int64_t, GSeekType, GCancellable*, GError**);
extern int64_t g_seekable_tell(GSeekable*);
extern int g_seekable_truncate(GSeekable*, int64_t, GCancellable*, GError**);
extern GType g_seekable_get_type();
extern GSettings* g_settings_new(char*);
extern GSettings* g_settings_new_with_backend(char*, GSettingsBackend*);
extern GSettings* g_settings_new_with_backend_and_path(char*, GSettingsBackend*, char*);
extern GSettings* g_settings_new_with_path(char*, char*);
extern char** g_settings_list_relocatable_schemas();
extern char** g_settings_list_schemas();
extern void g_settings_sync();
extern void g_settings_unbind(void*, char*);
extern void g_settings_apply(GSettings*);
extern void g_settings_bind(GSettings*, char*, GObject*, char*, GSettingsBindFlags);
extern void g_settings_bind_writable(GSettings*, char*, GObject*, char*, int);
extern void g_settings_delay(GSettings*);
extern int g_settings_get_boolean(GSettings*, char*);
extern GSettings* g_settings_get_child(GSettings*, char*);
extern double g_settings_get_double(GSettings*, char*);
extern int32_t g_settings_get_enum(GSettings*, char*);
extern uint32_t g_settings_get_flags(GSettings*, char*);
extern int g_settings_get_has_unapplied(GSettings*);
extern int32_t g_settings_get_int(GSettings*, char*);
extern void* g_settings_get_mapped(GSettings*, char*, GSettingsGetMapping, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void* _g_settings_get_mapped(GSettings* this, char* arg0, void* gofunc) {
	if (gofunc) {
		return g_settings_get_mapped(this, arg0, _GSettingsGetMapping_c_wrapper, gofunc);
	} else {
		return g_settings_get_mapped(this, arg0, 0, 0);
	}
}
extern GVariant* g_settings_get_range(GSettings*, char*);
extern char* g_settings_get_string(GSettings*, char*);
extern char** g_settings_get_strv(GSettings*, char*);
extern uint32_t g_settings_get_uint(GSettings*, char*);
extern GVariant* g_settings_get_value(GSettings*, char*);
extern int g_settings_is_writable(GSettings*, char*);
extern char** g_settings_list_children(GSettings*);
extern char** g_settings_list_keys(GSettings*);
extern int g_settings_range_check(GSettings*, char*, GVariant*);
extern void g_settings_reset(GSettings*, char*);
extern void g_settings_revert(GSettings*);
extern int g_settings_set_boolean(GSettings*, char*, int);
extern int g_settings_set_double(GSettings*, char*, double);
extern int g_settings_set_enum(GSettings*, char*, int32_t);
extern int g_settings_set_flags(GSettings*, char*, uint32_t);
extern int g_settings_set_int(GSettings*, char*, int32_t);
extern int g_settings_set_string(GSettings*, char*, char*);
extern int g_settings_set_strv(GSettings*, char*, char**);
extern int g_settings_set_uint(GSettings*, char*, uint32_t);
extern int g_settings_set_value(GSettings*, char*, GVariant*);
extern GType g_settings_get_type();
extern GSimpleAction* g_simple_action_new(char*, GVariantType*);
extern GSimpleAction* g_simple_action_new_stateful(char*, GVariantType*, GVariant*);
extern void g_simple_action_set_enabled(GSimpleAction*, int);
extern void g_simple_action_set_state(GSimpleAction*, GVariant*);
extern GType g_simple_action_get_type();
extern GSimpleActionGroup* g_simple_action_group_new();
extern void g_simple_action_group_add_entries(GSimpleActionGroup*, GActionEntry*, int32_t, void*);
extern void g_simple_action_group_insert(GSimpleActionGroup*, GAction*);
extern GAction* g_simple_action_group_lookup(GSimpleActionGroup*, char*);
extern void g_simple_action_group_remove(GSimpleActionGroup*, char*);
extern GType g_simple_action_group_get_type();
extern GSimpleAsyncResult* g_simple_async_result_new(GObject*, GAsyncReadyCallback, void*, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static GSimpleAsyncResult* _g_simple_async_result_new(GObject* arg0, void* gofunc, void* arg3) {
	if (gofunc) {
		return g_simple_async_result_new(arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc, arg3);
	} else {
		return g_simple_async_result_new(arg0, 0, 0, arg3);
	}
}
extern GSimpleAsyncResult* g_simple_async_result_new_from_error(GObject*, GAsyncReadyCallback, void*, GError*);
#pragma GCC diagnostic ignored "-Wunused-function"
static GSimpleAsyncResult* _g_simple_async_result_new_from_error(GObject* arg0, void* gofunc, GError* arg3) {
	if (gofunc) {
		return g_simple_async_result_new_from_error(arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc, arg3);
	} else {
		return g_simple_async_result_new_from_error(arg0, 0, 0, arg3);
	}
}
extern int g_simple_async_result_is_valid(GAsyncResult*, GObject*, void*);
extern void g_simple_async_result_complete(GSimpleAsyncResult*);
extern void g_simple_async_result_complete_in_idle(GSimpleAsyncResult*);
extern int g_simple_async_result_get_op_res_gboolean(GSimpleAsyncResult*);
extern int64_t g_simple_async_result_get_op_res_gssize(GSimpleAsyncResult*);
extern int g_simple_async_result_propagate_error(GSimpleAsyncResult*, GError**);
extern void g_simple_async_result_set_from_error(GSimpleAsyncResult*, GError*);
extern void g_simple_async_result_set_handle_cancellation(GSimpleAsyncResult*, int);
extern void g_simple_async_result_set_op_res_gboolean(GSimpleAsyncResult*, int);
extern void g_simple_async_result_set_op_res_gssize(GSimpleAsyncResult*, int64_t);
extern GType g_simple_async_result_get_type();
extern GPermission* g_simple_permission_new(int);
extern GType g_simple_permission_get_type();
extern GSocket* g_socket_new(GSocketFamily, GSocketType, GSocketProtocol, GError**);
extern GSocket* g_socket_new_from_fd(int32_t, GError**);
extern GSocket* g_socket_accept(GSocket*, GCancellable*, GError**);
extern int g_socket_bind(GSocket*, GSocketAddress*, int, GError**);
extern int g_socket_check_connect_result(GSocket*, GError**);
extern int g_socket_close(GSocket*, GError**);
extern GIOCondition g_socket_condition_check(GSocket*, GIOCondition);
extern int g_socket_condition_wait(GSocket*, GIOCondition, GCancellable*, GError**);
extern int g_socket_connect(GSocket*, GSocketAddress*, GCancellable*, GError**);
extern GSocketConnection* g_socket_connection_factory_create_connection(GSocket*);
extern int g_socket_get_blocking(GSocket*);
extern GCredentials* g_socket_get_credentials(GSocket*, GError**);
extern GSocketFamily g_socket_get_family(GSocket*);
extern int32_t g_socket_get_fd(GSocket*);
extern int g_socket_get_keepalive(GSocket*);
extern int32_t g_socket_get_listen_backlog(GSocket*);
extern GSocketAddress* g_socket_get_local_address(GSocket*, GError**);
extern GSocketProtocol g_socket_get_protocol(GSocket*);
extern GSocketAddress* g_socket_get_remote_address(GSocket*, GError**);
extern GSocketType g_socket_get_socket_type(GSocket*);
extern uint32_t g_socket_get_timeout(GSocket*);
extern int g_socket_is_closed(GSocket*);
extern int g_socket_is_connected(GSocket*);
extern int g_socket_listen(GSocket*, GError**);
extern int64_t g_socket_receive(GSocket*, char*, uint64_t, GCancellable*, GError**);
extern int64_t g_socket_receive_from(GSocket*, GSocketAddress*, char*, uint64_t, GCancellable*, GError**);
extern int64_t g_socket_receive_message(GSocket*, GSocketAddress*, GInputVector*, int32_t, GSocketControlMessage**, int32_t*, int32_t*, GCancellable*, GError**);
extern int64_t g_socket_receive_with_blocking(GSocket*, char*, uint64_t, int, GCancellable*, GError**);
extern int64_t g_socket_send(GSocket*, char**, uint64_t, GCancellable*, GError**);
extern int64_t g_socket_send_message(GSocket*, GSocketAddress*, GOutputVector*, int32_t, GSocketControlMessage**, int32_t, int32_t, GCancellable*, GError**);
extern int64_t g_socket_send_to(GSocket*, GSocketAddress*, char**, uint64_t, GCancellable*, GError**);
extern int64_t g_socket_send_with_blocking(GSocket*, char**, uint64_t, int, GCancellable*, GError**);
extern void g_socket_set_blocking(GSocket*, int);
extern void g_socket_set_keepalive(GSocket*, int);
extern void g_socket_set_listen_backlog(GSocket*, int32_t);
extern void g_socket_set_timeout(GSocket*, uint32_t);
extern int g_socket_shutdown(GSocket*, int, int, GError**);
extern int g_socket_speaks_ipv4(GSocket*);
extern GType g_socket_get_type();
extern GSocketAddress* g_socket_address_new_from_native(void*, uint64_t);
extern GSocketFamily g_socket_address_get_family(GSocketAddress*);
extern int64_t g_socket_address_get_native_size(GSocketAddress*);
extern int g_socket_address_to_native(GSocketAddress*, void*, uint64_t, GError**);
extern GType g_socket_address_get_type();
extern GSocketAddress* g_socket_address_enumerator_next(GSocketAddressEnumerator*, GCancellable*, GError**);
extern void g_socket_address_enumerator_next_async(GSocketAddressEnumerator*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_socket_address_enumerator_next_async(GSocketAddressEnumerator* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_socket_address_enumerator_next_async(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_socket_address_enumerator_next_async(this, arg0, 0, 0);
	}
}
extern GSocketAddress* g_socket_address_enumerator_next_finish(GSocketAddressEnumerator*, GAsyncResult*, GError**);
extern GType g_socket_address_enumerator_get_type();
extern GSocketClient* g_socket_client_new();
extern void g_socket_client_add_application_proxy(GSocketClient*, char*);
extern GSocketConnection* g_socket_client_connect(GSocketClient*, GSocketConnectable*, GCancellable*, GError**);
extern void g_socket_client_connect_async(GSocketClient*, GSocketConnectable*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_socket_client_connect_async(GSocketClient* this, GSocketConnectable* arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_socket_client_connect_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_socket_client_connect_async(this, arg0, arg1, 0, 0);
	}
}
extern GSocketConnection* g_socket_client_connect_finish(GSocketClient*, GAsyncResult*, GError**);
extern GSocketConnection* g_socket_client_connect_to_host(GSocketClient*, char*, uint16_t, GCancellable*, GError**);
extern void g_socket_client_connect_to_host_async(GSocketClient*, char*, uint16_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_socket_client_connect_to_host_async(GSocketClient* this, char* arg0, uint16_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_socket_client_connect_to_host_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_socket_client_connect_to_host_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GSocketConnection* g_socket_client_connect_to_host_finish(GSocketClient*, GAsyncResult*, GError**);
extern GSocketConnection* g_socket_client_connect_to_service(GSocketClient*, char*, char*, GCancellable*, GError**);
extern void g_socket_client_connect_to_service_async(GSocketClient*, char*, char*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_socket_client_connect_to_service_async(GSocketClient* this, char* arg0, char* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_socket_client_connect_to_service_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_socket_client_connect_to_service_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GSocketConnection* g_socket_client_connect_to_service_finish(GSocketClient*, GAsyncResult*, GError**);
extern GSocketConnection* g_socket_client_connect_to_uri(GSocketClient*, char*, uint16_t, GCancellable*, GError**);
extern void g_socket_client_connect_to_uri_async(GSocketClient*, char*, uint16_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_socket_client_connect_to_uri_async(GSocketClient* this, char* arg0, uint16_t arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_socket_client_connect_to_uri_async(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_socket_client_connect_to_uri_async(this, arg0, arg1, arg2, 0, 0);
	}
}
extern GSocketConnection* g_socket_client_connect_to_uri_finish(GSocketClient*, GAsyncResult*, GError**);
extern int g_socket_client_get_enable_proxy(GSocketClient*);
extern GSocketFamily g_socket_client_get_family(GSocketClient*);
extern GSocketAddress* g_socket_client_get_local_address(GSocketClient*);
extern GSocketProtocol g_socket_client_get_protocol(GSocketClient*);
extern GSocketType g_socket_client_get_socket_type(GSocketClient*);
extern uint32_t g_socket_client_get_timeout(GSocketClient*);
extern int g_socket_client_get_tls(GSocketClient*);
extern GTlsCertificateFlags g_socket_client_get_tls_validation_flags(GSocketClient*);
extern void g_socket_client_set_enable_proxy(GSocketClient*, int);
extern void g_socket_client_set_family(GSocketClient*, GSocketFamily);
extern void g_socket_client_set_local_address(GSocketClient*, GSocketAddress*);
extern void g_socket_client_set_protocol(GSocketClient*, GSocketProtocol);
extern void g_socket_client_set_socket_type(GSocketClient*, GSocketType);
extern void g_socket_client_set_timeout(GSocketClient*, uint32_t);
extern void g_socket_client_set_tls(GSocketClient*, int);
extern void g_socket_client_set_tls_validation_flags(GSocketClient*, GTlsCertificateFlags);
extern GType g_socket_client_get_type();
extern GSocketAddressEnumerator* g_socket_connectable_enumerate(GSocketConnectable*);
extern GSocketAddressEnumerator* g_socket_connectable_proxy_enumerate(GSocketConnectable*);
extern GType g_socket_connectable_get_type();
extern GType g_socket_connection_factory_lookup_type(GSocketFamily, GSocketType, int32_t);
extern void g_socket_connection_factory_register_type(GType, GSocketFamily, GSocketType, int32_t);
extern GSocketAddress* g_socket_connection_get_local_address(GSocketConnection*, GError**);
extern GSocketAddress* g_socket_connection_get_remote_address(GSocketConnection*, GError**);
extern GSocket* g_socket_connection_get_socket(GSocketConnection*);
extern GType g_socket_connection_get_type();
extern GSocketControlMessage* g_socket_control_message_deserialize(int32_t, int32_t, uint64_t, uint8_t*);
extern int32_t g_socket_control_message_get_level(GSocketControlMessage*);
extern int32_t g_socket_control_message_get_msg_type(GSocketControlMessage*);
extern uint64_t g_socket_control_message_get_size(GSocketControlMessage*);
extern void g_socket_control_message_serialize(GSocketControlMessage*, void*);
extern GType g_socket_control_message_get_type();
extern GSocketListener* g_socket_listener_new();
extern GSocketConnection* g_socket_listener_accept(GSocketListener*, GObject**, GCancellable*, GError**);
extern void g_socket_listener_accept_async(GSocketListener*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_socket_listener_accept_async(GSocketListener* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_socket_listener_accept_async(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_socket_listener_accept_async(this, arg0, 0, 0);
	}
}
extern GSocketConnection* g_socket_listener_accept_finish(GSocketListener*, GAsyncResult*, GObject**, GError**);
extern GSocket* g_socket_listener_accept_socket(GSocketListener*, GObject**, GCancellable*, GError**);
extern void g_socket_listener_accept_socket_async(GSocketListener*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_socket_listener_accept_socket_async(GSocketListener* this, GCancellable* arg0, void* gofunc) {
	if (gofunc) {
		g_socket_listener_accept_socket_async(this, arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_socket_listener_accept_socket_async(this, arg0, 0, 0);
	}
}
extern GSocket* g_socket_listener_accept_socket_finish(GSocketListener*, GAsyncResult*, GObject**, GError**);
extern int g_socket_listener_add_address(GSocketListener*, GSocketAddress*, GSocketType, GSocketProtocol, GObject*, GSocketAddress**, GError**);
extern uint16_t g_socket_listener_add_any_inet_port(GSocketListener*, GObject*, GError**);
extern int g_socket_listener_add_inet_port(GSocketListener*, uint16_t, GObject*, GError**);
extern int g_socket_listener_add_socket(GSocketListener*, GSocket*, GObject*, GError**);
extern void g_socket_listener_close(GSocketListener*);
extern void g_socket_listener_set_backlog(GSocketListener*, int32_t);
extern GType g_socket_listener_get_type();
extern GSocketService* g_socket_service_new();
extern int g_socket_service_is_active(GSocketService*);
extern void g_socket_service_start(GSocketService*);
extern void g_socket_service_stop(GSocketService*);
extern GType g_socket_service_get_type();
extern GSrvTarget* g_srv_target_new(char*, uint16_t, uint16_t, uint16_t);
extern GSrvTarget* g_srv_target_copy(GSrvTarget*);
extern void g_srv_target_free(GSrvTarget*);
extern char* g_srv_target_get_hostname(GSrvTarget*);
extern uint16_t g_srv_target_get_port(GSrvTarget*);
extern uint16_t g_srv_target_get_priority(GSrvTarget*);
extern uint16_t g_srv_target_get_weight(GSrvTarget*);
extern int g_tcp_connection_get_graceful_disconnect(GTcpConnection*);
extern void g_tcp_connection_set_graceful_disconnect(GTcpConnection*, int);
extern GType g_tcp_connection_get_type();
extern GSocketConnection* g_tcp_wrapper_connection_new(GIOStream*, GSocket*);
extern GIOStream* g_tcp_wrapper_connection_get_base_io_stream(GTcpWrapperConnection*);
extern GType g_tcp_wrapper_connection_get_type();
extern GIcon* g_themed_icon_new(char*);
extern GIcon* g_themed_icon_new_from_names(char**, int32_t);
extern GIcon* g_themed_icon_new_with_default_fallbacks(char*);
extern void g_themed_icon_append_name(GThemedIcon*, char*);
extern char** g_themed_icon_get_names(GThemedIcon*);
extern void g_themed_icon_prepend_name(GThemedIcon*, char*);
extern GType g_themed_icon_get_type();
extern GSocketService* g_threaded_socket_service_new(int32_t);
extern GType g_threaded_socket_service_get_type();
extern GTlsBackend* g_tls_backend_get_default();
extern GType g_tls_backend_get_certificate_type(GTlsBackend*);
extern GType g_tls_backend_get_client_connection_type(GTlsBackend*);
extern GTlsDatabase* g_tls_backend_get_default_database(GTlsBackend*);
extern GType g_tls_backend_get_file_database_type(GTlsBackend*);
extern GType g_tls_backend_get_server_connection_type(GTlsBackend*);
extern int g_tls_backend_supports_tls(GTlsBackend*);
extern GType g_tls_backend_get_type();
extern GTlsCertificate* g_tls_certificate_new_from_file(char*, GError**);
extern GTlsCertificate* g_tls_certificate_new_from_files(char*, char*, GError**);
extern GTlsCertificate* g_tls_certificate_new_from_pem(char*, int64_t, GError**);
extern GList* g_tls_certificate_list_new_from_file(char*, GError**);
extern GTlsCertificate* g_tls_certificate_get_issuer(GTlsCertificate*);
extern GTlsCertificateFlags g_tls_certificate_verify(GTlsCertificate*, GSocketConnectable*, GTlsCertificate*);
extern GType g_tls_certificate_get_type();
extern GIOStream* g_tls_client_connection_new(GIOStream*, GSocketConnectable*, GError**);
extern GList* g_tls_client_connection_get_accepted_cas(GTlsClientConnection*);
extern GSocketConnectable* g_tls_client_connection_get_server_identity(GTlsClientConnection*);
extern int g_tls_client_connection_get_use_ssl3(GTlsClientConnection*);
extern GTlsCertificateFlags g_tls_client_connection_get_validation_flags(GTlsClientConnection*);
extern void g_tls_client_connection_set_server_identity(GTlsClientConnection*, GSocketConnectable*);
extern void g_tls_client_connection_set_use_ssl3(GTlsClientConnection*, int);
extern void g_tls_client_connection_set_validation_flags(GTlsClientConnection*, GTlsCertificateFlags);
extern GType g_tls_client_connection_get_type();
extern int g_tls_connection_emit_accept_certificate(GTlsConnection*, GTlsCertificate*, GTlsCertificateFlags);
extern GTlsCertificate* g_tls_connection_get_certificate(GTlsConnection*);
extern GTlsDatabase* g_tls_connection_get_database(GTlsConnection*);
extern GTlsInteraction* g_tls_connection_get_interaction(GTlsConnection*);
extern GTlsCertificate* g_tls_connection_get_peer_certificate(GTlsConnection*);
extern GTlsCertificateFlags g_tls_connection_get_peer_certificate_errors(GTlsConnection*);
extern GTlsRehandshakeMode g_tls_connection_get_rehandshake_mode(GTlsConnection*);
extern int g_tls_connection_get_require_close_notify(GTlsConnection*);
extern int g_tls_connection_get_use_system_certdb(GTlsConnection*);
extern int g_tls_connection_handshake(GTlsConnection*, GCancellable*, GError**);
extern void g_tls_connection_handshake_async(GTlsConnection*, int32_t, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_tls_connection_handshake_async(GTlsConnection* this, int32_t arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_tls_connection_handshake_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_tls_connection_handshake_async(this, arg0, arg1, 0, 0);
	}
}
extern int g_tls_connection_handshake_finish(GTlsConnection*, GAsyncResult*, GError**);
extern void g_tls_connection_set_certificate(GTlsConnection*, GTlsCertificate*);
extern void g_tls_connection_set_database(GTlsConnection*, GTlsDatabase*);
extern void g_tls_connection_set_interaction(GTlsConnection*, GTlsInteraction*);
extern void g_tls_connection_set_rehandshake_mode(GTlsConnection*, GTlsRehandshakeMode);
extern void g_tls_connection_set_require_close_notify(GTlsConnection*, int);
extern void g_tls_connection_set_use_system_certdb(GTlsConnection*, int);
extern GType g_tls_connection_get_type();
extern char* g_tls_database_create_certificate_handle(GTlsDatabase*, GTlsCertificate*);
extern GTlsCertificate* g_tls_database_lookup_certificate_for_handle(GTlsDatabase*, char*, GTlsInteraction*, GTlsDatabaseLookupFlags, GCancellable*, GError**);
extern void g_tls_database_lookup_certificate_for_handle_async(GTlsDatabase*, char*, GTlsInteraction*, GTlsDatabaseLookupFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_tls_database_lookup_certificate_for_handle_async(GTlsDatabase* this, char* arg0, GTlsInteraction* arg1, GTlsDatabaseLookupFlags arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_tls_database_lookup_certificate_for_handle_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_tls_database_lookup_certificate_for_handle_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern GTlsCertificate* g_tls_database_lookup_certificate_for_handle_finish(GTlsDatabase*, GAsyncResult*, GError**);
extern GTlsCertificate* g_tls_database_lookup_certificate_issuer(GTlsDatabase*, GTlsCertificate*, GTlsInteraction*, GTlsDatabaseLookupFlags, GCancellable*, GError**);
extern void g_tls_database_lookup_certificate_issuer_async(GTlsDatabase*, GTlsCertificate*, GTlsInteraction*, GTlsDatabaseLookupFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_tls_database_lookup_certificate_issuer_async(GTlsDatabase* this, GTlsCertificate* arg0, GTlsInteraction* arg1, GTlsDatabaseLookupFlags arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_tls_database_lookup_certificate_issuer_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_tls_database_lookup_certificate_issuer_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern GTlsCertificate* g_tls_database_lookup_certificate_issuer_finish(GTlsDatabase*, GAsyncResult*, GError**);
extern void g_tls_database_lookup_certificates_issued_by_async(GTlsDatabase*, GByteArray*, GTlsInteraction*, GTlsDatabaseLookupFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_tls_database_lookup_certificates_issued_by_async(GTlsDatabase* this, GByteArray* arg0, GTlsInteraction* arg1, GTlsDatabaseLookupFlags arg2, GCancellable* arg3, void* gofunc) {
	if (gofunc) {
		g_tls_database_lookup_certificates_issued_by_async(this, arg0, arg1, arg2, arg3, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_tls_database_lookup_certificates_issued_by_async(this, arg0, arg1, arg2, arg3, 0, 0);
	}
}
extern GTlsCertificateFlags g_tls_database_verify_chain(GTlsDatabase*, GTlsCertificate*, char*, GSocketConnectable*, GTlsInteraction*, GTlsDatabaseVerifyFlags, GCancellable*, GError**);
extern void g_tls_database_verify_chain_async(GTlsDatabase*, GTlsCertificate*, char*, GSocketConnectable*, GTlsInteraction*, GTlsDatabaseVerifyFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_tls_database_verify_chain_async(GTlsDatabase* this, GTlsCertificate* arg0, char* arg1, GSocketConnectable* arg2, GTlsInteraction* arg3, GTlsDatabaseVerifyFlags arg4, GCancellable* arg5, void* gofunc) {
	if (gofunc) {
		g_tls_database_verify_chain_async(this, arg0, arg1, arg2, arg3, arg4, arg5, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_tls_database_verify_chain_async(this, arg0, arg1, arg2, arg3, arg4, arg5, 0, 0);
	}
}
extern GTlsCertificateFlags g_tls_database_verify_chain_finish(GTlsDatabase*, GAsyncResult*, GError**);
extern GType g_tls_database_get_type();
extern GTlsDatabase* g_tls_file_database_new(char*, GError**);
extern GType g_tls_file_database_get_type();
extern GTlsInteractionResult g_tls_interaction_ask_password(GTlsInteraction*, GTlsPassword*, GCancellable*, GError**);
extern void g_tls_interaction_ask_password_async(GTlsInteraction*, GTlsPassword*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_tls_interaction_ask_password_async(GTlsInteraction* this, GTlsPassword* arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_tls_interaction_ask_password_async(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_tls_interaction_ask_password_async(this, arg0, arg1, 0, 0);
	}
}
extern GTlsInteractionResult g_tls_interaction_ask_password_finish(GTlsInteraction*, GAsyncResult*, GError**);
extern GTlsInteractionResult g_tls_interaction_invoke_ask_password(GTlsInteraction*, GTlsPassword*, GCancellable*, GError**);
extern GType g_tls_interaction_get_type();
extern GTlsPassword* g_tls_password_new(GTlsPasswordFlags, char*);
extern char* g_tls_password_get_description(GTlsPassword*);
extern GTlsPasswordFlags g_tls_password_get_flags(GTlsPassword*);
extern uint8_t* g_tls_password_get_value(GTlsPassword*, uint64_t*);
extern char* g_tls_password_get_warning(GTlsPassword*);
extern void g_tls_password_set_description(GTlsPassword*, char*);
extern void g_tls_password_set_flags(GTlsPassword*, GTlsPasswordFlags);
extern void g_tls_password_set_value(GTlsPassword*, uint8_t*, int64_t);
extern void g_tls_password_set_value_full(GTlsPassword*, uint8_t*, int64_t, GDestroyNotify);
extern void g_tls_password_set_warning(GTlsPassword*, char*);
extern GType g_tls_password_get_type();
extern GIOStream* g_tls_server_connection_new(GIOStream*, GTlsCertificate*, GError**);
extern GType g_tls_server_connection_get_type();
extern GCredentials* g_unix_connection_receive_credentials(GUnixConnection*, GCancellable*, GError**);
extern int32_t g_unix_connection_receive_fd(GUnixConnection*, GCancellable*, GError**);
extern int g_unix_connection_send_credentials(GUnixConnection*, GCancellable*, GError**);
extern int g_unix_connection_send_fd(GUnixConnection*, int32_t, GCancellable*, GError**);
extern GType g_unix_connection_get_type();
extern GSocketControlMessage* g_unix_credentials_message_new();
extern GSocketControlMessage* g_unix_credentials_message_new_with_credentials(GCredentials*);
extern int g_unix_credentials_message_is_supported();
extern GCredentials* g_unix_credentials_message_get_credentials(GUnixCredentialsMessage*);
extern GType g_unix_credentials_message_get_type();
extern GUnixFDList* g_unix_fd_list_new();
extern GUnixFDList* g_unix_fd_list_new_from_array(int32_t*, int32_t);
extern int32_t g_unix_fd_list_append(GUnixFDList*, int32_t, GError**);
extern int32_t g_unix_fd_list_get(GUnixFDList*, int32_t, GError**);
extern int32_t g_unix_fd_list_get_length(GUnixFDList*);
extern int32_t* g_unix_fd_list_peek_fds(GUnixFDList*, int32_t*);
extern int32_t* g_unix_fd_list_steal_fds(GUnixFDList*, int32_t*);
extern GType g_unix_fd_list_get_type();
extern GSocketControlMessage* g_unix_fd_message_new();
extern GSocketControlMessage* g_unix_fd_message_new_with_fd_list(GUnixFDList*);
extern int g_unix_fd_message_append_fd(GUnixFDMessage*, int32_t, GError**);
extern GUnixFDList* g_unix_fd_message_get_fd_list(GUnixFDMessage*);
extern int32_t* g_unix_fd_message_steal_fds(GUnixFDMessage*, int32_t*);
extern GType g_unix_fd_message_get_type();
extern GInputStream* g_unix_input_stream_new(int32_t, int);
extern int g_unix_input_stream_get_close_fd(GUnixInputStream*);
extern int32_t g_unix_input_stream_get_fd(GUnixInputStream*);
extern void g_unix_input_stream_set_close_fd(GUnixInputStream*, int);
extern GType g_unix_input_stream_get_type();
extern GUnixMountMonitor* g_unix_mount_monitor_new();
extern void g_unix_mount_monitor_set_rate_limit(GUnixMountMonitor*, int32_t);
extern GType g_unix_mount_monitor_get_type();
extern int32_t g_unix_mount_point_compare(GUnixMountPoint*, GUnixMountPoint*);
extern void g_unix_mount_point_free(GUnixMountPoint*);
extern char* g_unix_mount_point_get_device_path(GUnixMountPoint*);
extern char* g_unix_mount_point_get_fs_type(GUnixMountPoint*);
extern char* g_unix_mount_point_get_mount_path(GUnixMountPoint*);
extern int g_unix_mount_point_guess_can_eject(GUnixMountPoint*);
extern GIcon* g_unix_mount_point_guess_icon(GUnixMountPoint*);
extern char* g_unix_mount_point_guess_name(GUnixMountPoint*);
extern int g_unix_mount_point_is_loopback(GUnixMountPoint*);
extern int g_unix_mount_point_is_readonly(GUnixMountPoint*);
extern int g_unix_mount_point_is_user_mountable(GUnixMountPoint*);
extern GOutputStream* g_unix_output_stream_new(int32_t, int);
extern int g_unix_output_stream_get_close_fd(GUnixOutputStream*);
extern int32_t g_unix_output_stream_get_fd(GUnixOutputStream*);
extern void g_unix_output_stream_set_close_fd(GUnixOutputStream*, int);
extern GType g_unix_output_stream_get_type();
extern GSocketAddress* g_unix_socket_address_new(char*);
extern GSocketAddress* g_unix_socket_address_new_abstract(uint8_t*, int32_t);
extern GSocketAddress* g_unix_socket_address_new_with_type(uint8_t*, int32_t, GUnixSocketAddressType);
extern int g_unix_socket_address_abstract_names_supported();
extern GUnixSocketAddressType g_unix_socket_address_get_address_type(GUnixSocketAddress*);
extern int g_unix_socket_address_get_is_abstract(GUnixSocketAddress*);
extern char* g_unix_socket_address_get_path(GUnixSocketAddress*);
extern uint64_t g_unix_socket_address_get_path_len(GUnixSocketAddress*);
extern GType g_unix_socket_address_get_type();
extern GVfs* g_vfs_get_default();
extern GVfs* g_vfs_get_local();
extern GFile* g_vfs_get_file_for_path(GVfs*, char*);
extern GFile* g_vfs_get_file_for_uri(GVfs*, char*);
extern char** g_vfs_get_supported_uri_schemes(GVfs*);
extern int g_vfs_is_active(GVfs*);
extern GFile* g_vfs_parse_name(GVfs*, char*);
extern GType g_vfs_get_type();
extern int g_volume_can_eject(GVolume*);
extern int g_volume_can_mount(GVolume*);
extern void g_volume_eject(GVolume*, GMountUnmountFlags, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_volume_eject(GVolume* this, GMountUnmountFlags arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_volume_eject(this, arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_volume_eject(this, arg0, arg1, 0, 0);
	}
}
extern int g_volume_eject_finish(GVolume*, GAsyncResult*, GError**);
extern void g_volume_eject_with_operation(GVolume*, GMountUnmountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_volume_eject_with_operation(GVolume* this, GMountUnmountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_volume_eject_with_operation(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_volume_eject_with_operation(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_volume_eject_with_operation_finish(GVolume*, GAsyncResult*, GError**);
extern char** g_volume_enumerate_identifiers(GVolume*);
extern GFile* g_volume_get_activation_root(GVolume*);
extern GDrive* g_volume_get_drive(GVolume*);
extern GIcon* g_volume_get_icon(GVolume*);
extern char* g_volume_get_identifier(GVolume*, char*);
extern GMount* g_volume_get_mount(GVolume*);
extern char* g_volume_get_name(GVolume*);
extern char* g_volume_get_uuid(GVolume*);
extern void g_volume_mount(GVolume*, GMountMountFlags, GMountOperation*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_volume_mount(GVolume* this, GMountMountFlags arg0, GMountOperation* arg1, GCancellable* arg2, void* gofunc) {
	if (gofunc) {
		g_volume_mount(this, arg0, arg1, arg2, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_volume_mount(this, arg0, arg1, arg2, 0, 0);
	}
}
extern int g_volume_mount_finish(GVolume*, GAsyncResult*, GError**);
extern int g_volume_should_automount(GVolume*);
extern GType g_volume_get_type();
extern GVolume* g_volume_monitor_adopt_orphan_mount(GMount*);
extern GVolumeMonitor* g_volume_monitor_get();
extern GList* g_volume_monitor_get_connected_drives(GVolumeMonitor*);
extern GMount* g_volume_monitor_get_mount_for_uuid(GVolumeMonitor*, char*);
extern GList* g_volume_monitor_get_mounts(GVolumeMonitor*);
extern GVolume* g_volume_monitor_get_volume_for_uuid(GVolumeMonitor*, char*);
extern GList* g_volume_monitor_get_volumes(GVolumeMonitor*);
extern GType g_volume_monitor_get_type();
extern GZlibCompressor* g_zlib_compressor_new(GZlibCompressorFormat, int32_t);
extern GFileInfo* g_zlib_compressor_get_file_info(GZlibCompressor*);
extern void g_zlib_compressor_set_file_info(GZlibCompressor*, GFileInfo*);
extern GType g_zlib_compressor_get_type();
extern GZlibDecompressor* g_zlib_decompressor_new(GZlibCompressorFormat);
extern GFileInfo* g_zlib_decompressor_get_file_info(GZlibDecompressor*);
extern GType g_zlib_decompressor_get_type();
extern void g_bus_get(GBusType, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_bus_get(GBusType arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_bus_get(arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_bus_get(arg0, arg1, 0, 0);
	}
}
extern GDBusConnection* g_bus_get_finish(GAsyncResult*, GError**);
extern GDBusConnection* g_bus_get_sync(GBusType, GCancellable*, GError**);
extern uint32_t g_bus_own_name_on_connection_with_closures(GDBusConnection*, char*, GBusNameOwnerFlags, GClosure*, GClosure*);
extern uint32_t g_bus_own_name_with_closures(GBusType, char*, GBusNameOwnerFlags, GClosure*, GClosure*, GClosure*);
extern void g_bus_unown_name(uint32_t);
extern void g_bus_unwatch_name(uint32_t);
extern uint32_t g_bus_watch_name_on_connection_with_closures(GDBusConnection*, char*, GBusNameWatcherFlags, GClosure*, GClosure*);
extern uint32_t g_bus_watch_name_with_closures(GBusType, char*, GBusNameWatcherFlags, GClosure*, GClosure*);
extern int g_content_type_can_be_executable(char*);
extern int g_content_type_equals(char*, char*);
extern char* g_content_type_from_mime_type(char*);
extern char* g_content_type_get_description(char*);
extern GIcon* g_content_type_get_icon(char*);
extern char* g_content_type_get_mime_type(char*);
extern char* g_content_type_guess(char*, uint8_t*, uint64_t, int*);
extern char** g_content_type_guess_for_tree(GFile*);
extern int g_content_type_is_a(char*, char*);
extern int g_content_type_is_unknown(char*);
extern GList* g_content_types_get_registered();
extern char* g_dbus_address_get_for_bus_sync(GBusType, GCancellable*, GError**);
extern void g_dbus_address_get_stream(char*, GCancellable*, GAsyncReadyCallback, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_dbus_address_get_stream(char* arg0, GCancellable* arg1, void* gofunc) {
	if (gofunc) {
		g_dbus_address_get_stream(arg0, arg1, _GAsyncReadyCallback_c_wrapper_once, gofunc);
	} else {
		g_dbus_address_get_stream(arg0, arg1, 0, 0);
	}
}
extern GIOStream* g_dbus_address_get_stream_finish(GAsyncResult*, char*, GError**);
extern GIOStream* g_dbus_address_get_stream_sync(char*, char*, GCancellable*, GError**);
extern char* g_dbus_error_encode_gerror(GError*);
extern char* g_dbus_error_get_remote_error(GError*);
extern int g_dbus_error_is_remote_error(GError*);
extern GError* g_dbus_error_new_for_dbus_error(char*, char*);
extern uint32_t g_dbus_error_quark();
extern int g_dbus_error_register_error(uint32_t, int32_t, char*);
extern void g_dbus_error_register_error_domain(char*, uint64_t*, GDBusErrorEntry*, uint32_t);
extern int g_dbus_error_strip_remote_error(GError*);
extern int g_dbus_error_unregister_error(uint32_t, int32_t, char*);
extern char* g_dbus_generate_guid();
extern GVariant* g_dbus_gvalue_to_gvariant(GValue*, GVariantType*);
extern void g_dbus_gvariant_to_gvalue(GVariant*, GValue*);
extern int g_dbus_is_address(char*);
extern int g_dbus_is_guid(char*);
extern int g_dbus_is_interface_name(char*);
extern int g_dbus_is_member_name(char*);
extern int g_dbus_is_name(char*);
extern int g_dbus_is_supported_address(char*, GError**);
extern int g_dbus_is_unique_name(char*);
extern GIOErrorEnum g_io_error_from_errno(int32_t);
extern uint32_t g_io_error_quark();
extern GType g_io_extension_get_type(GIOExtension*);
extern GList* g_io_modules_load_all_in_directory(char*);
extern GList* g_io_modules_load_all_in_directory_with_scope(char*, GIOModuleScope*);
extern void g_io_modules_scan_all_in_directory(char*);
extern void g_io_modules_scan_all_in_directory_with_scope(char*, GIOModuleScope*);
extern void g_io_scheduler_cancel_all_jobs();
extern void g_io_scheduler_push_job(GIOSchedulerJobFunc, void*, GDestroyNotify, int32_t, GCancellable*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_io_scheduler_push_job(void* gofunc, int32_t arg3, GCancellable* arg4) {
	if (gofunc) {
		g_io_scheduler_push_job(_GIOSchedulerJobFunc_c_wrapper, gofunc, _c_callback_cleanup, arg3, arg4);
	} else {
		g_io_scheduler_push_job(0, 0, 0, arg3, arg4);
	}
}
extern uint32_t g_resolver_error_quark();
extern void g_simple_async_report_gerror_in_idle(GObject*, GAsyncReadyCallback, void*, GError*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _g_simple_async_report_gerror_in_idle(GObject* arg0, void* gofunc, GError* arg3) {
	if (gofunc) {
		g_simple_async_report_gerror_in_idle(arg0, _GAsyncReadyCallback_c_wrapper_once, gofunc, arg3);
	} else {
		g_simple_async_report_gerror_in_idle(arg0, 0, 0, arg3);
	}
}
extern uint32_t g_tls_error_quark();
extern int g_unix_is_mount_path_system_internal(char*);
extern int32_t g_unix_mount_compare(GUnixMountEntry*, GUnixMountEntry*);
extern void g_unix_mount_free(GUnixMountEntry*);
extern char* g_unix_mount_get_device_path(GUnixMountEntry*);
extern char* g_unix_mount_get_fs_type(GUnixMountEntry*);
extern char* g_unix_mount_get_mount_path(GUnixMountEntry*);
extern int g_unix_mount_guess_can_eject(GUnixMountEntry*);
extern GIcon* g_unix_mount_guess_icon(GUnixMountEntry*);
extern char* g_unix_mount_guess_name(GUnixMountEntry*);
extern int g_unix_mount_guess_should_display(GUnixMountEntry*);
extern int g_unix_mount_is_readonly(GUnixMountEntry*);
extern int g_unix_mount_is_system_internal(GUnixMountEntry*);
extern int g_unix_mount_points_changed_since(uint64_t);
extern int g_unix_mounts_changed_since(uint64_t);
struct _GActionEntry { uint8_t _data[64]; };
struct _GActionGroupInterface { uint8_t _data[120]; };
struct _GActionInterface { uint8_t _data[80]; };
struct _GAppInfoIface { uint8_t _data[192]; };
struct _GAppLaunchContextClass { uint8_t _data[200]; };
struct _GAppLaunchContextPrivate {};
struct _GApplicationClass { uint8_t _data[312]; };
struct _GApplicationCommandLineClass { uint8_t _data[248]; };
struct _GApplicationCommandLinePrivate {};
struct _GApplicationPrivate {};
struct _GAsyncInitableIface { uint8_t _data[32]; };
struct _GAsyncResultIface { uint8_t _data[32]; };
struct _GBufferedInputStreamClass { uint8_t _data[336]; };
struct _GBufferedInputStreamPrivate {};
struct _GBufferedOutputStreamClass { uint8_t _data[336]; };
struct _GBufferedOutputStreamPrivate {};
struct _GCancellableClass { uint8_t _data[184]; };
struct _GCancellablePrivate {};
struct _GCharsetConverterClass { uint8_t _data[136]; };
struct _GConverterIface { uint8_t _data[32]; };
struct _GConverterInputStreamClass { uint8_t _data[312]; };
struct _GConverterInputStreamPrivate {};
struct _GConverterOutputStreamClass { uint8_t _data[360]; };
struct _GConverterOutputStreamPrivate {};
struct _GCredentialsClass {};
struct _GDBusAnnotationInfo { uint8_t _data[32]; };
struct _GDBusArgInfo { uint8_t _data[32]; };
struct _GDBusErrorEntry { uint8_t _data[16]; };
struct _GDBusInterfaceIface { uint8_t _data[40]; };
struct _GDBusInterfaceInfo { uint8_t _data[48]; };
struct _GDBusInterfaceSkeletonClass { uint8_t _data[304]; };
struct _GDBusInterfaceSkeletonPrivate {};
struct _GDBusInterfaceVTable { uint8_t _data[88]; };
struct _GDBusMethodInfo { uint8_t _data[40]; };
struct _GDBusNodeInfo { uint8_t _data[40]; };
struct _GDBusObjectIface { uint8_t _data[56]; };
struct _GDBusObjectManagerClientClass { uint8_t _data[216]; };
struct _GDBusObjectManagerClientPrivate {};
struct _GDBusObjectManagerIface { uint8_t _data[80]; };
struct _GDBusObjectManagerServerClass { uint8_t _data[200]; };
struct _GDBusObjectManagerServerPrivate {};
struct _GDBusObjectProxyClass { uint8_t _data[200]; };
struct _GDBusObjectProxyPrivate {};
struct _GDBusObjectSkeletonClass { uint8_t _data[208]; };
struct _GDBusObjectSkeletonPrivate {};
struct _GDBusPropertyInfo { uint8_t _data[40]; };
struct _GDBusProxyClass { uint8_t _data[408]; };
struct _GDBusProxyPrivate {};
struct _GDBusSignalInfo { uint8_t _data[32]; };
struct _GDBusSubtreeVTable { uint8_t _data[88]; };
struct _GDataInputStreamClass { uint8_t _data[376]; };
struct _GDataInputStreamPrivate {};
struct _GDataOutputStreamClass { uint8_t _data[360]; };
struct _GDataOutputStreamPrivate {};
struct _GDesktopAppInfoClass { uint8_t _data[136]; };
struct _GDesktopAppInfoLookupIface { uint8_t _data[24]; };
struct _GDriveIface { uint8_t _data[248]; };
struct _GEmblemClass {};
struct _GEmblemedIconClass { uint8_t _data[136]; };
struct _GEmblemedIconPrivate {};
struct _GFileAttributeInfo { uint8_t _data[16]; };
struct _GFileAttributeInfoList { uint8_t _data[16]; };
struct _GFileAttributeMatcher {};
struct _GFileDescriptorBasedIface { uint8_t _data[24]; };
struct _GFileEnumeratorClass { uint8_t _data[240]; };
struct _GFileEnumeratorPrivate {};
struct _GFileIOStreamClass { uint8_t _data[368]; };
struct _GFileIOStreamPrivate {};
struct _GFileIconClass {};
struct _GFileIface { uint8_t _data[816]; };
struct _GFileInfoClass {};
struct _GFileInputStreamClass { uint8_t _data[336]; };
struct _GFileInputStreamPrivate {};
struct _GFileMonitorClass { uint8_t _data[192]; };
struct _GFileMonitorPrivate {};
struct _GFileOutputStreamClass { uint8_t _data[408]; };
struct _GFileOutputStreamPrivate {};
struct _GFilenameCompleterClass { uint8_t _data[168]; };
struct _GFilterInputStreamClass { uint8_t _data[272]; };
struct _GFilterOutputStreamClass { uint8_t _data[320]; };
struct _GIOExtension {};
struct _GIOExtensionPoint {};
struct _GIOModuleClass {};
struct _GIOModuleScope {};
struct _GIOSchedulerJob {};
struct _GIOStreamAdapter {};
struct _GIOStreamClass { uint8_t _data[256]; };
struct _GIOStreamPrivate {};
struct _GIconIface { uint8_t _data[48]; };
struct _GInetAddressClass { uint8_t _data[152]; };
struct _GInetAddressPrivate {};
struct _GInetSocketAddressClass { uint8_t _data[160]; };
struct _GInetSocketAddressPrivate {};
struct _GInitableIface { uint8_t _data[24]; };
struct _GInputStreamClass { uint8_t _data[248]; };
struct _GInputStreamPrivate {};
struct _GInputVector { uint8_t _data[16]; };
struct _GLoadableIconIface { uint8_t _data[40]; };
struct _GMemoryInputStreamClass { uint8_t _data[288]; };
struct _GMemoryInputStreamPrivate {};
struct _GMemoryOutputStreamClass { uint8_t _data[336]; };
struct _GMemoryOutputStreamPrivate {};
struct _GMountIface { uint8_t _data[216]; };
struct _GMountOperationClass { uint8_t _data[256]; };
struct _GMountOperationPrivate {};
struct _GNativeVolumeMonitorClass { uint8_t _data[344]; };
struct _GNetworkAddressClass { uint8_t _data[136]; };
struct _GNetworkAddressPrivate {};
struct _GNetworkServiceClass { uint8_t _data[136]; };
struct _GNetworkServicePrivate {};
struct _GOutputStreamClass { uint8_t _data[296]; };
struct _GOutputStreamPrivate {};
struct _GOutputVector { uint8_t _data[16]; };
struct _GPermissionClass { uint8_t _data[312]; };
struct _GPermissionPrivate {};
struct _GPollableInputStreamInterface { uint8_t _data[48]; };
struct _GPollableOutputStreamInterface { uint8_t _data[48]; };
struct _GProxyAddressClass { uint8_t _data[160]; };
struct _GProxyAddressEnumeratorClass { uint8_t _data[216]; };
struct _GProxyAddressEnumeratorPrivate {};
struct _GProxyAddressPrivate {};
struct _GProxyInterface { uint8_t _data[48]; };
struct _GProxyResolverInterface { uint8_t _data[48]; };
struct _GResolverClass { uint8_t _data[264]; };
struct _GResolverPrivate {};
struct _GSeekableIface { uint8_t _data[56]; };
struct _GSettingsBackend {};
struct _GSettingsClass { uint8_t _data[328]; };
struct _GSettingsPrivate {};
struct _GSimpleActionGroupClass { uint8_t _data[232]; };
struct _GSimpleActionGroupPrivate {};
struct _GSimpleAsyncResultClass {};
struct _GSocketAddressClass { uint8_t _data[160]; };
struct _GSocketAddressEnumeratorClass { uint8_t _data[160]; };
struct _GSocketClass { uint8_t _data[216]; };
struct _GSocketClientClass { uint8_t _data[176]; };
struct _GSocketClientPrivate {};
struct _GSocketConnectableIface { uint8_t _data[32]; };
struct _GSocketConnectionClass { uint8_t _data[304]; };
struct _GSocketConnectionPrivate {};
struct _GSocketControlMessageClass { uint8_t _data[216]; };
struct _GSocketControlMessagePrivate {};
struct _GSocketListenerClass { uint8_t _data[192]; };
struct _GSocketListenerPrivate {};
struct _GSocketPrivate {};
struct _GSocketServiceClass { uint8_t _data[248]; };
struct _GSocketServicePrivate {};
struct _GSrvTarget {};
struct _GTcpConnectionClass { uint8_t _data[304]; };
struct _GTcpConnectionPrivate {};
struct _GTcpWrapperConnectionClass { uint8_t _data[304]; };
struct _GTcpWrapperConnectionPrivate {};
struct _GThemedIconClass {};
struct _GThreadedSocketServiceClass { uint8_t _data[296]; };
struct _GThreadedSocketServicePrivate {};
struct _GTlsBackendInterface { uint8_t _data[64]; };
struct _GTlsCertificateClass { uint8_t _data[208]; };
struct _GTlsCertificatePrivate {};
struct _GTlsClientConnectionInterface { uint8_t _data[16]; };
struct _GTlsConnectionClass { uint8_t _data[352]; };
struct _GTlsConnectionPrivate {};
struct _GTlsDatabaseClass { uint8_t _data[368]; };
struct _GTlsDatabasePrivate {};
struct _GTlsFileDatabaseInterface { uint8_t _data[80]; };
struct _GTlsInteractionClass { uint8_t _data[352]; };
struct _GTlsInteractionPrivate {};
struct _GTlsPasswordClass { uint8_t _data[192]; };
struct _GTlsPasswordPrivate {};
struct _GTlsServerConnectionInterface { uint8_t _data[16]; };
struct _GUnixConnectionClass { uint8_t _data[304]; };
struct _GUnixConnectionPrivate {};
struct _GUnixCredentialsMessageClass { uint8_t _data[232]; };
struct _GUnixCredentialsMessagePrivate {};
struct _GUnixFDListClass { uint8_t _data[176]; };
struct _GUnixFDListPrivate {};
struct _GUnixFDMessageClass { uint8_t _data[232]; };
struct _GUnixFDMessagePrivate {};
struct _GUnixInputStreamClass { uint8_t _data[288]; };
struct _GUnixInputStreamPrivate {};
struct _GUnixMountEntry {};
struct _GUnixMountMonitorClass {};
struct _GUnixMountPoint {};
struct _GUnixOutputStreamClass { uint8_t _data[336]; };
struct _GUnixOutputStreamPrivate {};
struct _GUnixSocketAddressClass { uint8_t _data[160]; };
struct _GUnixSocketAddressPrivate {};
struct _GVfsClass { uint8_t _data[272]; };
struct _GVolumeIface { uint8_t _data[168]; };
struct _GVolumeMonitorClass { uint8_t _data[336]; };
struct _GZlibCompressorClass { uint8_t _data[136]; };
struct _GZlibDecompressorClass { uint8_t _data[136]; };


extern GObject *g_object_ref_sink(GObject*);
extern void g_object_unref(GObject*);
extern void g_error_free(GError*);
extern void g_free(void*);

#cgo pkg-config: gio-2.0
*/
import "C"
import "unsafe"
import "errors"

import (
	"github.com/bytbox/gogobject/gobject-2.0"
	"github.com/bytbox/gogobject/glib-2.0"
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


//export _Gio_go_callback_cleanup
func _Gio_go_callback_cleanup(gofunc unsafe.Pointer) {
	gobject.Holder.Release(gofunc)
}


// blacklisted: Action (interface)
// blacklisted: ActionEntry (struct)
type ActionGroupLike interface {
	ImplementsGActionGroup() *C.GActionGroup
}

type ActionGroup struct {
	gobject.Object
	ActionGroupImpl
}

type ActionGroupImpl struct {}

func ToActionGroup(objlike gobject.ObjectLike) *ActionGroup {
	t := (*ActionGroupImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ActionGroup)(obj)
	}
	panic("cannot cast to ActionGroup")
}

func (this0 *ActionGroupImpl) ImplementsGActionGroup() *C.GActionGroup {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GActionGroup)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *ActionGroupImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_action_group_get_type())
}

func ActionGroupGetType() gobject.Type {
	return (*ActionGroupImpl)(nil).GetStaticType()
}
func (this0 *ActionGroupImpl) ActionAdded(action_name0 string) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	C.g_action_group_action_added(this1, action_name1)
}
func (this0 *ActionGroupImpl) ActionEnabledChanged(action_name0 string, enabled0 bool) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	var enabled1 C.int
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	enabled1 = _GoBoolToCBool(enabled0)
	C.g_action_group_action_enabled_changed(this1, action_name1, enabled1)
}
func (this0 *ActionGroupImpl) ActionRemoved(action_name0 string) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	C.g_action_group_action_removed(this1, action_name1)
}
func (this0 *ActionGroupImpl) ActionStateChanged(action_name0 string, state0 *glib.Variant) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	var state1 *C.GVariant
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	state1 = (*C.GVariant)(unsafe.Pointer(state0))
	C.g_action_group_action_state_changed(this1, action_name1, state1)
}
func (this0 *ActionGroupImpl) ActivateAction(action_name0 string, parameter0 *glib.Variant) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	var parameter1 *C.GVariant
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	parameter1 = (*C.GVariant)(unsafe.Pointer(parameter0))
	C.g_action_group_activate_action(this1, action_name1, parameter1)
}
func (this0 *ActionGroupImpl) ChangeActionState(action_name0 string, value0 *glib.Variant) {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	var value1 *C.GVariant
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	value1 = (*C.GVariant)(unsafe.Pointer(value0))
	C.g_action_group_change_action_state(this1, action_name1, value1)
}
func (this0 *ActionGroupImpl) GetActionEnabled(action_name0 string) bool {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_get_action_enabled(this1, action_name1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *ActionGroupImpl) GetActionParameterType(action_name0 string) *glib.VariantType {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_get_action_parameter_type(this1, action_name1)
	var ret2 *glib.VariantType
	ret2 = (*glib.VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionGroupImpl) GetActionState(action_name0 string) *glib.Variant {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_get_action_state(this1, action_name1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionGroupImpl) GetActionStateHint(action_name0 string) *glib.Variant {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_get_action_state_hint(this1, action_name1)
	var ret2 *glib.Variant
	ret2 = (*glib.Variant)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionGroupImpl) GetActionStateType(action_name0 string) *glib.VariantType {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_get_action_state_type(this1, action_name1)
	var ret2 *glib.VariantType
	ret2 = (*glib.VariantType)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ActionGroupImpl) HasAction(action_name0 string) bool {
	var this1 *C.GActionGroup
	var action_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	action_name1 = _GoStringToGString(action_name0)
	defer C.free(unsafe.Pointer(action_name1))
	ret1 := C.g_action_group_has_action(this1, action_name1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *ActionGroupImpl) ListActions() []string {
	var this1 *C.GActionGroup
	if this0 != nil {
		this1 = this0.ImplementsGActionGroup()}
	ret1 := C.g_action_group_list_actions(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
// blacklisted: ActionGroupInterface (struct)
// blacklisted: ActionInterface (struct)
type AppInfoLike interface {
	ImplementsGAppInfo() *C.GAppInfo
}

type AppInfo struct {
	gobject.Object
	AppInfoImpl
}

type AppInfoImpl struct {}

func ToAppInfo(objlike gobject.ObjectLike) *AppInfo {
	t := (*AppInfoImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*AppInfo)(obj)
	}
	panic("cannot cast to AppInfo")
}

func (this0 *AppInfoImpl) ImplementsGAppInfo() *C.GAppInfo {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GAppInfo)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *AppInfoImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_app_info_get_type())
}

func AppInfoGetType() gobject.Type {
	return (*AppInfoImpl)(nil).GetStaticType()
}
// blacklisted: AppInfo.create_from_commandline (method)
// blacklisted: AppInfo.get_all (method)
// blacklisted: AppInfo.get_all_for_type (method)
// blacklisted: AppInfo.get_default_for_type (method)
// blacklisted: AppInfo.get_default_for_uri_scheme (method)
// blacklisted: AppInfo.get_fallback_for_type (method)
// blacklisted: AppInfo.get_recommended_for_type (method)
// blacklisted: AppInfo.launch_default_for_uri (method)
// blacklisted: AppInfo.reset_type_associations (method)
// blacklisted: AppInfo.add_supports_type (method)
// blacklisted: AppInfo.can_delete (method)
// blacklisted: AppInfo.can_remove_supports_type (method)
// blacklisted: AppInfo.delete (method)
// blacklisted: AppInfo.dup (method)
// blacklisted: AppInfo.equal (method)
// blacklisted: AppInfo.get_commandline (method)
// blacklisted: AppInfo.get_description (method)
// blacklisted: AppInfo.get_display_name (method)
// blacklisted: AppInfo.get_executable (method)
// blacklisted: AppInfo.get_icon (method)
// blacklisted: AppInfo.get_id (method)
// blacklisted: AppInfo.get_name (method)
// blacklisted: AppInfo.launch (method)
// blacklisted: AppInfo.launch_uris (method)
// blacklisted: AppInfo.remove_supports_type (method)
// blacklisted: AppInfo.set_as_default_for_extension (method)
// blacklisted: AppInfo.set_as_default_for_type (method)
// blacklisted: AppInfo.set_as_last_used_for_type (method)
// blacklisted: AppInfo.should_show (method)
// blacklisted: AppInfo.supports_files (method)
// blacklisted: AppInfo.supports_uris (method)
type AppInfoCreateFlags C.uint32_t
const (
	AppInfoCreateFlagsNone AppInfoCreateFlags = 0
	AppInfoCreateFlagsNeedsTerminal AppInfoCreateFlags = 1
	AppInfoCreateFlagsSupportsURIs AppInfoCreateFlags = 2
	AppInfoCreateFlagsSupportsStartupNotification AppInfoCreateFlags = 4
)
// blacklisted: AppInfoIface (struct)
// blacklisted: AppLaunchContext (object)
// blacklisted: AppLaunchContextClass (struct)
// blacklisted: AppLaunchContextPrivate (struct)
type ApplicationLike interface {
	gobject.ObjectLike
	InheritedFromGApplication() *C.GApplication
}

type Application struct {
	gobject.Object
	ActionGroupImpl
}

func ToApplication(objlike gobject.ObjectLike) *Application {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Application)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Application)(obj)
	}
	panic("cannot cast to Application")
}

func (this0 *Application) InheritedFromGApplication() *C.GApplication {
	if this0 == nil {
		return nil
	}
	return (*C.GApplication)(this0.C)
}

func (this0 *Application) GetStaticType() gobject.Type {
	return gobject.Type(C.g_application_get_type())
}

func ApplicationGetType() gobject.Type {
	return (*Application)(nil).GetStaticType()
}
func NewApplication(application_id0 string, flags0 ApplicationFlags) *Application {
	var application_id1 *C.char
	var flags1 C.GApplicationFlags
	application_id1 = _GoStringToGString(application_id0)
	defer C.free(unsafe.Pointer(application_id1))
	flags1 = C.GApplicationFlags(flags0)
	ret1 := C.g_application_new(application_id1, flags1)
	var ret2 *Application
	ret2 = (*Application)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func ApplicationIDIsValid(application_id0 string) bool {
	var application_id1 *C.char
	application_id1 = _GoStringToGString(application_id0)
	defer C.free(unsafe.Pointer(application_id1))
	ret1 := C.g_application_id_is_valid(application_id1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Application) Activate() {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	C.g_application_activate(this1)
}
func (this0 *Application) GetApplicationID() string {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_application_id(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Application) GetFlags() ApplicationFlags {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_flags(this1)
	var ret2 ApplicationFlags
	ret2 = ApplicationFlags(ret1)
	return ret2
}
func (this0 *Application) GetInactivityTimeout() int {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_inactivity_timeout(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Application) GetIsRegistered() bool {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_is_registered(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Application) GetIsRemote() bool {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	ret1 := C.g_application_get_is_remote(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Application) Hold() {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	C.g_application_hold(this1)
}
func (this0 *Application) Open(files0 []FileLike, hint0 string) {
	var this1 *C.GApplication
	var files1 **C.GFile
	var n_files1 C.int32_t
	var hint1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	files1 = (**C.GFile)(C.malloc(C.size_t(int(unsafe.Sizeof(*files1)) * len(files0))))
	defer C.free(unsafe.Pointer(files1))
	for i, e := range files0 {
		if e != nil {
			(*(*[999999]*C.GFile)(unsafe.Pointer(files1)))[i] = e.ImplementsGFile()}
	}
	n_files1 = C.int32_t(len(files0))
	hint1 = _GoStringToGString(hint0)
	defer C.free(unsafe.Pointer(hint1))
	C.g_application_open(this1, files1, n_files1, hint1)
}
func (this0 *Application) Register(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GApplication
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	if cancellable0 != nil {
		cancellable1 = cancellable0.InheritedFromGCancellable()
	}
	ret1 := C.g_application_register(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Application) Release() {
	var this1 *C.GApplication
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	C.g_application_release(this1)
}
func (this0 *Application) Run(argv0 []string) int {
	var this1 *C.GApplication
	var argv1 **C.char
	var argc1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	argv1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*argv1)) * len(argv0))))
	defer C.free(unsafe.Pointer(argv1))
	for i, e := range argv0 {
		(*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(argv1)))[i]))
	}
	argc1 = C.int32_t(len(argv0))
	ret1 := C.g_application_run(this1, argc1, argv1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Application) SetActionGroup(action_group0 ActionGroupLike) {
	var this1 *C.GApplication
	var action_group1 *C.GActionGroup
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	if action_group0 != nil {
		action_group1 = action_group0.ImplementsGActionGroup()}
	C.g_application_set_action_group(this1, action_group1)
}
func (this0 *Application) SetApplicationID(application_id0 string) {
	var this1 *C.GApplication
	var application_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	application_id1 = _GoStringToGString(application_id0)
	defer C.free(unsafe.Pointer(application_id1))
	C.g_application_set_application_id(this1, application_id1)
}
func (this0 *Application) SetFlags(flags0 ApplicationFlags) {
	var this1 *C.GApplication
	var flags1 C.GApplicationFlags
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	flags1 = C.GApplicationFlags(flags0)
	C.g_application_set_flags(this1, flags1)
}
func (this0 *Application) SetInactivityTimeout(inactivity_timeout0 int) {
	var this1 *C.GApplication
	var inactivity_timeout1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGApplication()
	}
	inactivity_timeout1 = C.uint32_t(inactivity_timeout0)
	C.g_application_set_inactivity_timeout(this1, inactivity_timeout1)
}
// blacklisted: ApplicationClass (struct)
// blacklisted: ApplicationCommandLine (object)
// blacklisted: ApplicationCommandLineClass (struct)
// blacklisted: ApplicationCommandLinePrivate (struct)
type ApplicationFlags C.uint32_t
const (
	ApplicationFlagsFlagsNone ApplicationFlags = 0
	ApplicationFlagsIsService ApplicationFlags = 1
	ApplicationFlagsIsLauncher ApplicationFlags = 2
	ApplicationFlagsHandlesOpen ApplicationFlags = 4
	ApplicationFlagsHandlesCommandLine ApplicationFlags = 8
	ApplicationFlagsSendEnvironment ApplicationFlags = 16
	ApplicationFlagsNonUnique ApplicationFlags = 32
)
// blacklisted: ApplicationPrivate (struct)
type AskPasswordFlags C.uint32_t
const (
	AskPasswordFlagsNeedPassword AskPasswordFlags = 1
	AskPasswordFlagsNeedUsername AskPasswordFlags = 2
	AskPasswordFlagsNeedDomain AskPasswordFlags = 4
	AskPasswordFlagsSavingSupported AskPasswordFlags = 8
	AskPasswordFlagsAnonymousSupported AskPasswordFlags = 16
)
// blacklisted: AsyncInitable (interface)
// blacklisted: AsyncInitableIface (struct)
type AsyncReadyCallback func(source_object *gobject.Object, res *AsyncResult)
//export _GAsyncReadyCallback_c_wrapper
func _GAsyncReadyCallback_c_wrapper(source_object0 unsafe.Pointer, res0 unsafe.Pointer, user_data0 unsafe.Pointer) {
	var source_object1 *gobject.Object
	var res1 *AsyncResult
	var user_data1 AsyncReadyCallback
	source_object1 = (*gobject.Object)(gobject.ObjectWrap(unsafe.Pointer((*C.GObject)(source_object0)), true))
	res1 = (*AsyncResult)(gobject.ObjectWrap(unsafe.Pointer((*C.GAsyncResult)(res0)), true))
	user_data1 = *(*AsyncReadyCallback)(user_data0)
	user_data1(source_object1, res1)
}
//export _GAsyncReadyCallback_c_wrapper_once
func _GAsyncReadyCallback_c_wrapper_once(source_object0 unsafe.Pointer, res0 unsafe.Pointer, user_data0 unsafe.Pointer) {
	_GAsyncReadyCallback_c_wrapper(source_object0, res0, user_data0)
	gobject.Holder.Release(user_data0)
}
type AsyncResultLike interface {
	ImplementsGAsyncResult() *C.GAsyncResult
}

type AsyncResult struct {
	gobject.Object
	AsyncResultImpl
}

type AsyncResultImpl struct {}

func ToAsyncResult(objlike gobject.ObjectLike) *AsyncResult {
	t := (*AsyncResultImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*AsyncResult)(obj)
	}
	panic("cannot cast to AsyncResult")
}

func (this0 *AsyncResultImpl) ImplementsGAsyncResult() *C.GAsyncResult {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GAsyncResult)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *AsyncResultImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_async_result_get_type())
}

func AsyncResultGetType() gobject.Type {
	return (*AsyncResultImpl)(nil).GetStaticType()
}
func (this0 *AsyncResultImpl) GetSourceObject() *gobject.Object {
	var this1 *C.GAsyncResult
	if this0 != nil {
		this1 = this0.ImplementsGAsyncResult()}
	ret1 := C.g_async_result_get_source_object(this1)
	var ret2 *gobject.Object
	ret2 = (*gobject.Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *AsyncResultImpl) GetUserData() {
	var this1 *C.GAsyncResult
	if this0 != nil {
		this1 = this0.ImplementsGAsyncResult()}
	C.g_async_result_get_user_data(this1)
}
// blacklisted: AsyncResultIface (struct)
// blacklisted: BufferedInputStream (object)
// blacklisted: BufferedInputStreamClass (struct)
// blacklisted: BufferedInputStreamPrivate (struct)
// blacklisted: BufferedOutputStream (object)
// blacklisted: BufferedOutputStreamClass (struct)
// blacklisted: BufferedOutputStreamPrivate (struct)
// blacklisted: BusAcquiredCallback (callback)
// blacklisted: BusNameAcquiredCallback (callback)
// blacklisted: BusNameAppearedCallback (callback)
// blacklisted: BusNameLostCallback (callback)
type BusNameOwnerFlags C.uint32_t
const (
	BusNameOwnerFlagsNone BusNameOwnerFlags = 0
	BusNameOwnerFlagsAllowReplacement BusNameOwnerFlags = 1
	BusNameOwnerFlagsReplace BusNameOwnerFlags = 2
)
// blacklisted: BusNameVanishedCallback (callback)
type BusNameWatcherFlags C.uint32_t
const (
	BusNameWatcherFlagsNone BusNameWatcherFlags = 0
	BusNameWatcherFlagsAutoStart BusNameWatcherFlags = 1
)
type BusType C.int32_t
const (
	BusTypeStarter BusType = -1
	BusTypeNone BusType = 0
	BusTypeSystem BusType = 1
	BusTypeSession BusType = 2
)
type CancellableLike interface {
	gobject.ObjectLike
	InheritedFromGCancellable() *C.GCancellable
}

type Cancellable struct {
	gobject.Object
	
}

func ToCancellable(objlike gobject.ObjectLike) *Cancellable {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Cancellable)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Cancellable)(obj)
	}
	panic("cannot cast to Cancellable")
}

func (this0 *Cancellable) InheritedFromGCancellable() *C.GCancellable {
	if this0 == nil {
		return nil
	}
	return (*C.GCancellable)(this0.C)
}

func (this0 *Cancellable) GetStaticType() gobject.Type {
	return gobject.Type(C.g_cancellable_get_type())
}

func CancellableGetType() gobject.Type {
	return (*Cancellable)(nil).GetStaticType()
}
func NewCancellable() *Cancellable {
	ret1 := C.g_cancellable_new()
	var ret2 *Cancellable
	ret2 = (*Cancellable)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func CancellableGetCurrent() *Cancellable {
	ret1 := C.g_cancellable_get_current()
	var ret2 *Cancellable
	ret2 = (*Cancellable)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Cancellable) Cancel() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	C.g_cancellable_cancel(this1)
}
// blacklisted: Cancellable.connect (method)
func (this0 *Cancellable) Disconnect(handler_id0 uint64) {
	var this1 *C.GCancellable
	var handler_id1 C.uint64_t
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	handler_id1 = C.uint64_t(handler_id0)
	C.g_cancellable_disconnect(this1, handler_id1)
}
func (this0 *Cancellable) GetFD() int {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	ret1 := C.g_cancellable_get_fd(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Cancellable) IsCancelled() bool {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	ret1 := C.g_cancellable_is_cancelled(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Cancellable) MakePollfd(pollfd0 *glib.PollFD) bool {
	var this1 *C.GCancellable
	var pollfd1 *C.GPollFD
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	pollfd1 = (*C.GPollFD)(unsafe.Pointer(pollfd0))
	ret1 := C.g_cancellable_make_pollfd(this1, pollfd1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Cancellable) PopCurrent() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	C.g_cancellable_pop_current(this1)
}
func (this0 *Cancellable) PushCurrent() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	C.g_cancellable_push_current(this1)
}
func (this0 *Cancellable) ReleaseFD() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	C.g_cancellable_release_fd(this1)
}
func (this0 *Cancellable) Reset() {
	var this1 *C.GCancellable
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	C.g_cancellable_reset(this1)
}
func (this0 *Cancellable) SetErrorIfCancelled() (bool, error) {
	var this1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGCancellable()
	}
	ret1 := C.g_cancellable_set_error_if_cancelled(this1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: CancellableClass (struct)
// blacklisted: CancellablePrivate (struct)
// blacklisted: CancellableSourceFunc (callback)
// blacklisted: CharsetConverter (object)
// blacklisted: CharsetConverterClass (struct)
// blacklisted: Converter (interface)
type ConverterFlags C.uint32_t
const (
	ConverterFlagsNone ConverterFlags = 0
	ConverterFlagsInputAtEnd ConverterFlags = 1
	ConverterFlagsFlush ConverterFlags = 2
)
// blacklisted: ConverterIface (struct)
// blacklisted: ConverterInputStream (object)
// blacklisted: ConverterInputStreamClass (struct)
// blacklisted: ConverterInputStreamPrivate (struct)
// blacklisted: ConverterOutputStream (object)
// blacklisted: ConverterOutputStreamClass (struct)
// blacklisted: ConverterOutputStreamPrivate (struct)
type ConverterResult C.uint32_t
const (
	ConverterResultError ConverterResult = 0
	ConverterResultConverted ConverterResult = 1
	ConverterResultFinished ConverterResult = 2
	ConverterResultFlushed ConverterResult = 3
)
// blacklisted: Credentials (object)
// blacklisted: CredentialsClass (struct)
type CredentialsType C.uint32_t
const (
	CredentialsTypeInvalid CredentialsType = 0
	CredentialsTypeLinuxUcred CredentialsType = 1
	CredentialsTypeFreebsdCmsgcred CredentialsType = 2
	CredentialsTypeOpenbsdSockpeercred CredentialsType = 3
)
// blacklisted: DBusAnnotationInfo (struct)
// blacklisted: DBusArgInfo (struct)
// blacklisted: DBusAuthObserver (object)
type DBusCallFlags C.uint32_t
const (
	DBusCallFlagsNone DBusCallFlags = 0
	DBusCallFlagsNoAutoStart DBusCallFlags = 1
)
type DBusCapabilityFlags C.uint32_t
const (
	DBusCapabilityFlagsNone DBusCapabilityFlags = 0
	DBusCapabilityFlagsUnixFDPassing DBusCapabilityFlags = 1
)
// blacklisted: DBusConnection (object)
type DBusConnectionFlags C.uint32_t
const (
	DBusConnectionFlagsNone DBusConnectionFlags = 0
	DBusConnectionFlagsAuthenticationClient DBusConnectionFlags = 1
	DBusConnectionFlagsAuthenticationServer DBusConnectionFlags = 2
	DBusConnectionFlagsAuthenticationAllowAnonymous DBusConnectionFlags = 4
	DBusConnectionFlagsMessageBusConnection DBusConnectionFlags = 8
	DBusConnectionFlagsDelayMessageProcessing DBusConnectionFlags = 16
)
type DBusError C.uint32_t
const (
	DBusErrorFailed DBusError = 0
	DBusErrorNoMemory DBusError = 1
	DBusErrorServiceUnknown DBusError = 2
	DBusErrorNameHasNoOwner DBusError = 3
	DBusErrorNoReply DBusError = 4
	DBusErrorIoError DBusError = 5
	DBusErrorBadAddress DBusError = 6
	DBusErrorNotSupported DBusError = 7
	DBusErrorLimitsExceeded DBusError = 8
	DBusErrorAccessDenied DBusError = 9
	DBusErrorAuthFailed DBusError = 10
	DBusErrorNoServer DBusError = 11
	DBusErrorTimeout DBusError = 12
	DBusErrorNoNetwork DBusError = 13
	DBusErrorAddressInUse DBusError = 14
	DBusErrorDisconnected DBusError = 15
	DBusErrorInvalidArgs DBusError = 16
	DBusErrorFileNotFound DBusError = 17
	DBusErrorFileExists DBusError = 18
	DBusErrorUnknownMethod DBusError = 19
	DBusErrorTimedOut DBusError = 20
	DBusErrorMatchRuleNotFound DBusError = 21
	DBusErrorMatchRuleInvalid DBusError = 22
	DBusErrorSpawnExecFailed DBusError = 23
	DBusErrorSpawnForkFailed DBusError = 24
	DBusErrorSpawnChildExited DBusError = 25
	DBusErrorSpawnChildSignaled DBusError = 26
	DBusErrorSpawnFailed DBusError = 27
	DBusErrorSpawnSetupFailed DBusError = 28
	DBusErrorSpawnConfigInvalid DBusError = 29
	DBusErrorSpawnServiceInvalid DBusError = 30
	DBusErrorSpawnServiceNotFound DBusError = 31
	DBusErrorSpawnPermissionsInvalid DBusError = 32
	DBusErrorSpawnFileInvalid DBusError = 33
	DBusErrorSpawnNoMemory DBusError = 34
	DBusErrorUnixProcessIDUnknown DBusError = 35
	DBusErrorInvalidSignature DBusError = 36
	DBusErrorInvalidFileContent DBusError = 37
	DBusErrorSelinuxSecurityContextUnknown DBusError = 38
	DBusErrorAdtAuditDataUnknown DBusError = 39
	DBusErrorObjectPathInUse DBusError = 40
)
// blacklisted: DBusErrorEntry (struct)
// blacklisted: DBusInterface (interface)
// blacklisted: DBusInterfaceGetPropertyFunc (callback)
// blacklisted: DBusInterfaceIface (struct)
// blacklisted: DBusInterfaceInfo (struct)
// blacklisted: DBusInterfaceMethodCallFunc (callback)
// blacklisted: DBusInterfaceSetPropertyFunc (callback)
// blacklisted: DBusInterfaceSkeleton (object)
// blacklisted: DBusInterfaceSkeletonClass (struct)
type DBusInterfaceSkeletonFlags C.uint32_t
const (
	DBusInterfaceSkeletonFlagsNone DBusInterfaceSkeletonFlags = 0
	DBusInterfaceSkeletonFlagsHandleMethodInvocationsInThread DBusInterfaceSkeletonFlags = 1
)
// blacklisted: DBusInterfaceSkeletonPrivate (struct)
// blacklisted: DBusInterfaceVTable (struct)
// blacklisted: DBusMessage (object)
type DBusMessageByteOrder C.uint32_t
const (
	DBusMessageByteOrderBigEndian DBusMessageByteOrder = 66
	DBusMessageByteOrderLittleEndian DBusMessageByteOrder = 108
)
// blacklisted: DBusMessageFilterFunction (callback)
type DBusMessageFlags C.uint32_t
const (
	DBusMessageFlagsNone DBusMessageFlags = 0
	DBusMessageFlagsNoReplyExpected DBusMessageFlags = 1
	DBusMessageFlagsNoAutoStart DBusMessageFlags = 2
)
type DBusMessageHeaderField C.uint32_t
const (
	DBusMessageHeaderFieldInvalid DBusMessageHeaderField = 0
	DBusMessageHeaderFieldPath DBusMessageHeaderField = 1
	DBusMessageHeaderFieldInterface DBusMessageHeaderField = 2
	DBusMessageHeaderFieldMember DBusMessageHeaderField = 3
	DBusMessageHeaderFieldErrorName DBusMessageHeaderField = 4
	DBusMessageHeaderFieldReplySerial DBusMessageHeaderField = 5
	DBusMessageHeaderFieldDestination DBusMessageHeaderField = 6
	DBusMessageHeaderFieldSender DBusMessageHeaderField = 7
	DBusMessageHeaderFieldSignature DBusMessageHeaderField = 8
	DBusMessageHeaderFieldNumUnixFds DBusMessageHeaderField = 9
)
type DBusMessageType C.uint32_t
const (
	DBusMessageTypeInvalid DBusMessageType = 0
	DBusMessageTypeMethodCall DBusMessageType = 1
	DBusMessageTypeMethodReturn DBusMessageType = 2
	DBusMessageTypeError DBusMessageType = 3
	DBusMessageTypeSignal DBusMessageType = 4
)
// blacklisted: DBusMethodInfo (struct)
// blacklisted: DBusMethodInvocation (object)
// blacklisted: DBusNodeInfo (struct)
// blacklisted: DBusObject (interface)
// blacklisted: DBusObjectIface (struct)
// blacklisted: DBusObjectManager (interface)
// blacklisted: DBusObjectManagerClient (object)
// blacklisted: DBusObjectManagerClientClass (struct)
type DBusObjectManagerClientFlags C.uint32_t
const (
	DBusObjectManagerClientFlagsNone DBusObjectManagerClientFlags = 0
	DBusObjectManagerClientFlagsDoNotAutoStart DBusObjectManagerClientFlags = 1
)
// blacklisted: DBusObjectManagerClientPrivate (struct)
// blacklisted: DBusObjectManagerIface (struct)
// blacklisted: DBusObjectManagerServer (object)
// blacklisted: DBusObjectManagerServerClass (struct)
// blacklisted: DBusObjectManagerServerPrivate (struct)
// blacklisted: DBusObjectProxy (object)
// blacklisted: DBusObjectProxyClass (struct)
// blacklisted: DBusObjectProxyPrivate (struct)
// blacklisted: DBusObjectSkeleton (object)
// blacklisted: DBusObjectSkeletonClass (struct)
// blacklisted: DBusObjectSkeletonPrivate (struct)
// blacklisted: DBusPropertyInfo (struct)
type DBusPropertyInfoFlags C.uint32_t
const (
	DBusPropertyInfoFlagsNone DBusPropertyInfoFlags = 0
	DBusPropertyInfoFlagsReadable DBusPropertyInfoFlags = 1
	DBusPropertyInfoFlagsWritable DBusPropertyInfoFlags = 2
)
// blacklisted: DBusProxy (object)
// blacklisted: DBusProxyClass (struct)
type DBusProxyFlags C.uint32_t
const (
	DBusProxyFlagsNone DBusProxyFlags = 0
	DBusProxyFlagsDoNotLoadProperties DBusProxyFlags = 1
	DBusProxyFlagsDoNotConnectSignals DBusProxyFlags = 2
	DBusProxyFlagsDoNotAutoStart DBusProxyFlags = 4
)
// blacklisted: DBusProxyPrivate (struct)
// blacklisted: DBusProxyTypeFunc (callback)
type DBusSendMessageFlags C.uint32_t
const (
	DBusSendMessageFlagsNone DBusSendMessageFlags = 0
	DBusSendMessageFlagsPreserveSerial DBusSendMessageFlags = 1
)
// blacklisted: DBusServer (object)
type DBusServerFlags C.uint32_t
const (
	DBusServerFlagsNone DBusServerFlags = 0
	DBusServerFlagsRunInThread DBusServerFlags = 1
	DBusServerFlagsAuthenticationAllowAnonymous DBusServerFlags = 2
)
// blacklisted: DBusSignalCallback (callback)
type DBusSignalFlags C.uint32_t
const (
	DBusSignalFlagsNone DBusSignalFlags = 0
	DBusSignalFlagsNoMatchRule DBusSignalFlags = 1
)
// blacklisted: DBusSignalInfo (struct)
// blacklisted: DBusSubtreeDispatchFunc (callback)
type DBusSubtreeFlags C.uint32_t
const (
	DBusSubtreeFlagsNone DBusSubtreeFlags = 0
	DBusSubtreeFlagsDispatchToUnenumeratedNodes DBusSubtreeFlags = 1
)
// blacklisted: DBusSubtreeIntrospectFunc (callback)
// blacklisted: DBusSubtreeVTable (struct)
const DesktopAppInfoLookupExtensionPointName = "gio-desktop-app-info-lookup"
// blacklisted: DataInputStream (object)
// blacklisted: DataInputStreamClass (struct)
// blacklisted: DataInputStreamPrivate (struct)
// blacklisted: DataOutputStream (object)
// blacklisted: DataOutputStreamClass (struct)
// blacklisted: DataOutputStreamPrivate (struct)
type DataStreamByteOrder C.uint32_t
const (
	DataStreamByteOrderBigEndian DataStreamByteOrder = 0
	DataStreamByteOrderLittleEndian DataStreamByteOrder = 1
	DataStreamByteOrderHostEndian DataStreamByteOrder = 2
)
type DataStreamNewlineType C.uint32_t
const (
	DataStreamNewlineTypeLf DataStreamNewlineType = 0
	DataStreamNewlineTypeCr DataStreamNewlineType = 1
	DataStreamNewlineTypeCrLf DataStreamNewlineType = 2
	DataStreamNewlineTypeAny DataStreamNewlineType = 3
)
// blacklisted: DesktopAppInfo (object)
// blacklisted: DesktopAppInfoClass (struct)
// blacklisted: DesktopAppInfoLookup (interface)
// blacklisted: DesktopAppInfoLookupIface (struct)
// blacklisted: DesktopAppLaunchCallback (callback)
// blacklisted: Drive (interface)
// blacklisted: DriveIface (struct)
type DriveStartFlags C.uint32_t
const (
	DriveStartFlagsNone DriveStartFlags = 0
)
type DriveStartStopType C.uint32_t
const (
	DriveStartStopTypeUnknown DriveStartStopType = 0
	DriveStartStopTypeShutdown DriveStartStopType = 1
	DriveStartStopTypeNetwork DriveStartStopType = 2
	DriveStartStopTypeMultidisk DriveStartStopType = 3
	DriveStartStopTypePassword DriveStartStopType = 4
)
type EmblemLike interface {
	gobject.ObjectLike
	InheritedFromGEmblem() *C.GEmblem
}

type Emblem struct {
	gobject.Object
	IconImpl
}

func ToEmblem(objlike gobject.ObjectLike) *Emblem {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Emblem)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Emblem)(obj)
	}
	panic("cannot cast to Emblem")
}

func (this0 *Emblem) InheritedFromGEmblem() *C.GEmblem {
	if this0 == nil {
		return nil
	}
	return (*C.GEmblem)(this0.C)
}

func (this0 *Emblem) GetStaticType() gobject.Type {
	return gobject.Type(C.g_emblem_get_type())
}

func EmblemGetType() gobject.Type {
	return (*Emblem)(nil).GetStaticType()
}
func NewEmblem(icon0 IconLike) *Emblem {
	var icon1 *C.GIcon
	if icon0 != nil {
		icon1 = icon0.ImplementsGIcon()}
	ret1 := C.g_emblem_new(icon1)
	var ret2 *Emblem
	ret2 = (*Emblem)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewEmblemWithOrigin(icon0 IconLike, origin0 EmblemOrigin) *Emblem {
	var icon1 *C.GIcon
	var origin1 C.GEmblemOrigin
	if icon0 != nil {
		icon1 = icon0.ImplementsGIcon()}
	origin1 = C.GEmblemOrigin(origin0)
	ret1 := C.g_emblem_new_with_origin(icon1, origin1)
	var ret2 *Emblem
	ret2 = (*Emblem)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Emblem) GetIcon() *Icon {
	var this1 *C.GEmblem
	if this0 != nil {
		this1 = this0.InheritedFromGEmblem()
	}
	ret1 := C.g_emblem_get_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Emblem) GetOrigin() EmblemOrigin {
	var this1 *C.GEmblem
	if this0 != nil {
		this1 = this0.InheritedFromGEmblem()
	}
	ret1 := C.g_emblem_get_origin(this1)
	var ret2 EmblemOrigin
	ret2 = EmblemOrigin(ret1)
	return ret2
}
// blacklisted: EmblemClass (struct)
type EmblemOrigin C.uint32_t
const (
	EmblemOriginUnknown EmblemOrigin = 0
	EmblemOriginDevice EmblemOrigin = 1
	EmblemOriginLivemetadata EmblemOrigin = 2
	EmblemOriginTag EmblemOrigin = 3
)
type EmblemedIconLike interface {
	gobject.ObjectLike
	InheritedFromGEmblemedIcon() *C.GEmblemedIcon
}

type EmblemedIcon struct {
	gobject.Object
	IconImpl
}

func ToEmblemedIcon(objlike gobject.ObjectLike) *EmblemedIcon {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*EmblemedIcon)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*EmblemedIcon)(obj)
	}
	panic("cannot cast to EmblemedIcon")
}

func (this0 *EmblemedIcon) InheritedFromGEmblemedIcon() *C.GEmblemedIcon {
	if this0 == nil {
		return nil
	}
	return (*C.GEmblemedIcon)(this0.C)
}

func (this0 *EmblemedIcon) GetStaticType() gobject.Type {
	return gobject.Type(C.g_emblemed_icon_get_type())
}

func EmblemedIconGetType() gobject.Type {
	return (*EmblemedIcon)(nil).GetStaticType()
}
func EmblemedIconNew(icon0 IconLike, emblem0 EmblemLike) *Icon {
	var icon1 *C.GIcon
	var emblem1 *C.GEmblem
	if icon0 != nil {
		icon1 = icon0.ImplementsGIcon()}
	if emblem0 != nil {
		emblem1 = emblem0.InheritedFromGEmblem()
	}
	ret1 := C.g_emblemed_icon_new(icon1, emblem1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *EmblemedIcon) AddEmblem(emblem0 EmblemLike) {
	var this1 *C.GEmblemedIcon
	var emblem1 *C.GEmblem
	if this0 != nil {
		this1 = this0.InheritedFromGEmblemedIcon()
	}
	if emblem0 != nil {
		emblem1 = emblem0.InheritedFromGEmblem()
	}
	C.g_emblemed_icon_add_emblem(this1, emblem1)
}
func (this0 *EmblemedIcon) ClearEmblems() {
	var this1 *C.GEmblemedIcon
	if this0 != nil {
		this1 = this0.InheritedFromGEmblemedIcon()
	}
	C.g_emblemed_icon_clear_emblems(this1)
}
func (this0 *EmblemedIcon) GetEmblems() []*Emblem {
	var this1 *C.GEmblemedIcon
	if this0 != nil {
		this1 = this0.InheritedFromGEmblemedIcon()
	}
	ret1 := C.g_emblemed_icon_get_emblems(this1)
	var ret2 []*Emblem
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Emblem
		elt = (*Emblem)(gobject.ObjectWrap(unsafe.Pointer((*C.GEmblem)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *EmblemedIcon) GetIcon() *Icon {
	var this1 *C.GEmblemedIcon
	if this0 != nil {
		this1 = this0.InheritedFromGEmblemedIcon()
	}
	ret1 := C.g_emblemed_icon_get_icon(this1)
	var ret2 *Icon
	ret2 = (*Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
// blacklisted: EmblemedIconClass (struct)
// blacklisted: EmblemedIconPrivate (struct)
const FileAttributeAccessCanDelete = "access::can-delete"
const FileAttributeAccessCanExecute = "access::can-execute"
const FileAttributeAccessCanRead = "access::can-read"
const FileAttributeAccessCanRename = "access::can-rename"
const FileAttributeAccessCanTrash = "access::can-trash"
const FileAttributeAccessCanWrite = "access::can-write"
const FileAttributeDosIsArchive = "dos::is-archive"
const FileAttributeDosIsSystem = "dos::is-system"
const FileAttributeEtagValue = "etag::value"
const FileAttributeFilesystemFree = "filesystem::free"
const FileAttributeFilesystemReadonly = "filesystem::readonly"
const FileAttributeFilesystemSize = "filesystem::size"
const FileAttributeFilesystemType = "filesystem::type"
const FileAttributeFilesystemUsePreview = "filesystem::use-preview"
const FileAttributeGvfsBackend = "gvfs::backend"
const FileAttributeIDFile = "id::file"
const FileAttributeIDFilesystem = "id::filesystem"
const FileAttributeMountableCanEject = "mountable::can-eject"
const FileAttributeMountableCanMount = "mountable::can-mount"
const FileAttributeMountableCanPoll = "mountable::can-poll"
const FileAttributeMountableCanStart = "mountable::can-start"
const FileAttributeMountableCanStartDegraded = "mountable::can-start-degraded"
const FileAttributeMountableCanStop = "mountable::can-stop"
const FileAttributeMountableCanUnmount = "mountable::can-unmount"
const FileAttributeMountableHalUdi = "mountable::hal-udi"
const FileAttributeMountableIsMediaCheckAutomatic = "mountable::is-media-check-automatic"
const FileAttributeMountableStartStopType = "mountable::start-stop-type"
const FileAttributeMountableUnixDevice = "mountable::unix-device"
const FileAttributeMountableUnixDeviceFile = "mountable::unix-device-file"
const FileAttributeOwnerGroup = "owner::group"
const FileAttributeOwnerUser = "owner::user"
const FileAttributeOwnerUserReal = "owner::user-real"
const FileAttributePreviewIcon = "preview::icon"
const FileAttributeSelinuxContext = "selinux::context"
const FileAttributeStandardAllocatedSize = "standard::allocated-size"
const FileAttributeStandardContentType = "standard::content-type"
const FileAttributeStandardCopyName = "standard::copy-name"
const FileAttributeStandardDescription = "standard::description"
const FileAttributeStandardDisplayName = "standard::display-name"
const FileAttributeStandardEditName = "standard::edit-name"
const FileAttributeStandardFastContentType = "standard::fast-content-type"
const FileAttributeStandardIcon = "standard::icon"
const FileAttributeStandardIsBackup = "standard::is-backup"
const FileAttributeStandardIsHidden = "standard::is-hidden"
const FileAttributeStandardIsSymlink = "standard::is-symlink"
const FileAttributeStandardIsVirtual = "standard::is-virtual"
const FileAttributeStandardName = "standard::name"
const FileAttributeStandardSize = "standard::size"
const FileAttributeStandardSortOrder = "standard::sort-order"
const FileAttributeStandardSymlinkTarget = "standard::symlink-target"
const FileAttributeStandardTargetURI = "standard::target-uri"
const FileAttributeStandardType = "standard::type"
const FileAttributeThumbnailingFailed = "thumbnail::failed"
const FileAttributeThumbnailPath = "thumbnail::path"
const FileAttributeTimeAccess = "time::access"
const FileAttributeTimeAccessUsec = "time::access-usec"
const FileAttributeTimeChanged = "time::changed"
const FileAttributeTimeChangedUsec = "time::changed-usec"
const FileAttributeTimeCreated = "time::created"
const FileAttributeTimeCreatedUsec = "time::created-usec"
const FileAttributeTimeModified = "time::modified"
const FileAttributeTimeModifiedUsec = "time::modified-usec"
const FileAttributeTrashDeletionDate = "trash::deletion-date"
const FileAttributeTrashItemCount = "trash::item-count"
const FileAttributeTrashOrigPath = "trash::orig-path"
const FileAttributeUnixBlocks = "unix::blocks"
const FileAttributeUnixBlockSize = "unix::block-size"
const FileAttributeUnixDevice = "unix::device"
const FileAttributeUnixGid = "unix::gid"
const FileAttributeUnixInode = "unix::inode"
const FileAttributeUnixIsMountpoint = "unix::is-mountpoint"
const FileAttributeUnixMode = "unix::mode"
const FileAttributeUnixNlink = "unix::nlink"
const FileAttributeUnixRdev = "unix::rdev"
const FileAttributeUnixUid = "unix::uid"
type FileLike interface {
	ImplementsGFile() *C.GFile
}

type File struct {
	gobject.Object
	FileImpl
}

type FileImpl struct {}

func ToFile(objlike gobject.ObjectLike) *File {
	t := (*FileImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*File)(obj)
	}
	panic("cannot cast to File")
}

func (this0 *FileImpl) ImplementsGFile() *C.GFile {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GFile)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *FileImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_file_get_type())
}

func FileGetType() gobject.Type {
	return (*FileImpl)(nil).GetStaticType()
}
// blacklisted: File.hash (method)
// blacklisted: File.new_for_commandline_arg (method)
// blacklisted: File.new_for_path (method)
// blacklisted: File.new_for_uri (method)
// blacklisted: File.parse_name (method)
// blacklisted: File.append_to (method)
// blacklisted: File.append_to_async (method)
// blacklisted: File.append_to_finish (method)
// blacklisted: File.copy (method)
// blacklisted: File.copy_attributes (method)
// blacklisted: File.copy_finish (method)
// blacklisted: File.create (method)
// blacklisted: File.create_async (method)
// blacklisted: File.create_finish (method)
// blacklisted: File.create_readwrite (method)
// blacklisted: File.create_readwrite_async (method)
// blacklisted: File.create_readwrite_finish (method)
// blacklisted: File.delete (method)
// blacklisted: File.dup (method)
// blacklisted: File.eject_mountable (method)
// blacklisted: File.eject_mountable_finish (method)
// blacklisted: File.eject_mountable_with_operation (method)
// blacklisted: File.eject_mountable_with_operation_finish (method)
// blacklisted: File.enumerate_children (method)
// blacklisted: File.enumerate_children_async (method)
// blacklisted: File.enumerate_children_finish (method)
// blacklisted: File.equal (method)
// blacklisted: File.find_enclosing_mount (method)
// blacklisted: File.find_enclosing_mount_async (method)
// blacklisted: File.find_enclosing_mount_finish (method)
// blacklisted: File.get_basename (method)
// blacklisted: File.get_child (method)
// blacklisted: File.get_child_for_display_name (method)
// blacklisted: File.get_parent (method)
// blacklisted: File.get_parse_name (method)
// blacklisted: File.get_path (method)
// blacklisted: File.get_relative_path (method)
// blacklisted: File.get_uri (method)
// blacklisted: File.get_uri_scheme (method)
// blacklisted: File.has_parent (method)
// blacklisted: File.has_prefix (method)
// blacklisted: File.has_uri_scheme (method)
// blacklisted: File.icon_new (method)
// blacklisted: File.is_native (method)
// blacklisted: File.load_contents (method)
// blacklisted: File.load_contents_async (method)
// blacklisted: File.load_contents_finish (method)
// blacklisted: File.load_partial_contents_finish (method)
// blacklisted: File.make_directory (method)
// blacklisted: File.make_directory_with_parents (method)
// blacklisted: File.make_symbolic_link (method)
// blacklisted: File.monitor (method)
// blacklisted: File.monitor_directory (method)
// blacklisted: File.monitor_file (method)
// blacklisted: File.mount_enclosing_volume (method)
// blacklisted: File.mount_enclosing_volume_finish (method)
// blacklisted: File.mount_mountable (method)
// blacklisted: File.mount_mountable_finish (method)
// blacklisted: File.move (method)
// blacklisted: File.open_readwrite (method)
// blacklisted: File.open_readwrite_async (method)
// blacklisted: File.open_readwrite_finish (method)
// blacklisted: File.poll_mountable (method)
// blacklisted: File.poll_mountable_finish (method)
// blacklisted: File.query_default_handler (method)
// blacklisted: File.query_exists (method)
// blacklisted: File.query_file_type (method)
// blacklisted: File.query_filesystem_info (method)
// blacklisted: File.query_filesystem_info_async (method)
// blacklisted: File.query_filesystem_info_finish (method)
// blacklisted: File.query_info (method)
// blacklisted: File.query_info_async (method)
// blacklisted: File.query_info_finish (method)
// blacklisted: File.query_settable_attributes (method)
// blacklisted: File.query_writable_namespaces (method)
// blacklisted: File.read (method)
// blacklisted: File.read_async (method)
// blacklisted: File.read_finish (method)
// blacklisted: File.replace (method)
// blacklisted: File.replace_async (method)
// blacklisted: File.replace_contents (method)
// blacklisted: File.replace_contents_async (method)
// blacklisted: File.replace_contents_finish (method)
// blacklisted: File.replace_finish (method)
// blacklisted: File.replace_readwrite (method)
// blacklisted: File.replace_readwrite_async (method)
// blacklisted: File.replace_readwrite_finish (method)
// blacklisted: File.resolve_relative_path (method)
// blacklisted: File.set_attribute (method)
// blacklisted: File.set_attribute_byte_string (method)
// blacklisted: File.set_attribute_int32 (method)
// blacklisted: File.set_attribute_int64 (method)
// blacklisted: File.set_attribute_string (method)
// blacklisted: File.set_attribute_uint32 (method)
// blacklisted: File.set_attribute_uint64 (method)
// blacklisted: File.set_attributes_async (method)
// blacklisted: File.set_attributes_finish (method)
// blacklisted: File.set_attributes_from_info (method)
// blacklisted: File.set_display_name (method)
// blacklisted: File.set_display_name_async (method)
// blacklisted: File.set_display_name_finish (method)
// blacklisted: File.start_mountable (method)
// blacklisted: File.start_mountable_finish (method)
// blacklisted: File.stop_mountable (method)
// blacklisted: File.stop_mountable_finish (method)
// blacklisted: File.supports_thread_contexts (method)
// blacklisted: File.trash (method)
// blacklisted: File.unmount_mountable (method)
// blacklisted: File.unmount_mountable_finish (method)
// blacklisted: File.unmount_mountable_with_operation (method)
// blacklisted: File.unmount_mountable_with_operation_finish (method)
// blacklisted: FileAttributeInfo (struct)
type FileAttributeInfoFlags C.uint32_t
const (
	FileAttributeInfoFlagsNone FileAttributeInfoFlags = 0
	FileAttributeInfoFlagsCopyWithFile FileAttributeInfoFlags = 1
	FileAttributeInfoFlagsCopyWhenMoved FileAttributeInfoFlags = 2
)
// blacklisted: FileAttributeInfoList (struct)
// blacklisted: FileAttributeMatcher (struct)
type FileAttributeStatus C.uint32_t
const (
	FileAttributeStatusUnset FileAttributeStatus = 0
	FileAttributeStatusSet FileAttributeStatus = 1
	FileAttributeStatusErrorSetting FileAttributeStatus = 2
)
type FileAttributeType C.uint32_t
const (
	FileAttributeTypeInvalid FileAttributeType = 0
	FileAttributeTypeString FileAttributeType = 1
	FileAttributeTypeByteString FileAttributeType = 2
	FileAttributeTypeBoolean FileAttributeType = 3
	FileAttributeTypeUint32 FileAttributeType = 4
	FileAttributeTypeInt32 FileAttributeType = 5
	FileAttributeTypeUint64 FileAttributeType = 6
	FileAttributeTypeInt64 FileAttributeType = 7
	FileAttributeTypeObject FileAttributeType = 8
	FileAttributeTypeStringv FileAttributeType = 9
)
type FileCopyFlags C.uint32_t
const (
	FileCopyFlagsNone FileCopyFlags = 0
	FileCopyFlagsOverwrite FileCopyFlags = 1
	FileCopyFlagsBackup FileCopyFlags = 2
	FileCopyFlagsNofollowSymlinks FileCopyFlags = 4
	FileCopyFlagsAllMetadata FileCopyFlags = 8
	FileCopyFlagsNoFallbackForMove FileCopyFlags = 16
	FileCopyFlagsTargetDefaultPerms FileCopyFlags = 32
)
type FileCreateFlags C.uint32_t
const (
	FileCreateFlagsNone FileCreateFlags = 0
	FileCreateFlagsPrivate FileCreateFlags = 1
	FileCreateFlagsReplaceDestination FileCreateFlags = 2
)
// blacklisted: FileDescriptorBased (interface)
// blacklisted: FileDescriptorBasedIface (struct)
// blacklisted: FileEnumerator (object)
// blacklisted: FileEnumeratorClass (struct)
// blacklisted: FileEnumeratorPrivate (struct)
// blacklisted: FileIOStream (object)
// blacklisted: FileIOStreamClass (struct)
// blacklisted: FileIOStreamPrivate (struct)
// blacklisted: FileIcon (object)
// blacklisted: FileIconClass (struct)
// blacklisted: FileIface (struct)
// blacklisted: FileInfo (object)
// blacklisted: FileInfoClass (struct)
// blacklisted: FileInputStream (object)
// blacklisted: FileInputStreamClass (struct)
// blacklisted: FileInputStreamPrivate (struct)
// blacklisted: FileMonitor (object)
// blacklisted: FileMonitorClass (struct)
type FileMonitorEvent C.uint32_t
const (
	FileMonitorEventChanged FileMonitorEvent = 0
	FileMonitorEventChangesDoneHint FileMonitorEvent = 1
	FileMonitorEventDeleted FileMonitorEvent = 2
	FileMonitorEventCreated FileMonitorEvent = 3
	FileMonitorEventAttributeChanged FileMonitorEvent = 4
	FileMonitorEventPreUnmount FileMonitorEvent = 5
	FileMonitorEventUnmounted FileMonitorEvent = 6
	FileMonitorEventMoved FileMonitorEvent = 7
)
type FileMonitorFlags C.uint32_t
const (
	FileMonitorFlagsNone FileMonitorFlags = 0
	FileMonitorFlagsWatchMounts FileMonitorFlags = 1
	FileMonitorFlagsSendMoved FileMonitorFlags = 2
)
// blacklisted: FileMonitorPrivate (struct)
// blacklisted: FileOutputStream (object)
// blacklisted: FileOutputStreamClass (struct)
// blacklisted: FileOutputStreamPrivate (struct)
// blacklisted: FileProgressCallback (callback)
type FileQueryInfoFlags C.uint32_t
const (
	FileQueryInfoFlagsNone FileQueryInfoFlags = 0
	FileQueryInfoFlagsNofollowSymlinks FileQueryInfoFlags = 1
)
// blacklisted: FileReadMoreCallback (callback)
type FileType C.uint32_t
const (
	FileTypeUnknown FileType = 0
	FileTypeRegular FileType = 1
	FileTypeDirectory FileType = 2
	FileTypeSymbolicLink FileType = 3
	FileTypeSpecial FileType = 4
	FileTypeShortcut FileType = 5
	FileTypeMountable FileType = 6
)
// blacklisted: FilenameCompleter (object)
// blacklisted: FilenameCompleterClass (struct)
type FilesystemPreviewType C.uint32_t
const (
	FilesystemPreviewTypeIfAlways FilesystemPreviewType = 0
	FilesystemPreviewTypeIfLocal FilesystemPreviewType = 1
	FilesystemPreviewTypeNever FilesystemPreviewType = 2
)
// blacklisted: FilterInputStream (object)
// blacklisted: FilterInputStreamClass (struct)
// blacklisted: FilterOutputStream (object)
// blacklisted: FilterOutputStreamClass (struct)
type IOErrorEnum C.uint32_t
const (
	IOErrorEnumFailed IOErrorEnum = 0
	IOErrorEnumNotFound IOErrorEnum = 1
	IOErrorEnumExists IOErrorEnum = 2
	IOErrorEnumIsDirectory IOErrorEnum = 3
	IOErrorEnumNotDirectory IOErrorEnum = 4
	IOErrorEnumNotEmpty IOErrorEnum = 5
	IOErrorEnumNotRegularFile IOErrorEnum = 6
	IOErrorEnumNotSymbolicLink IOErrorEnum = 7
	IOErrorEnumNotMountableFile IOErrorEnum = 8
	IOErrorEnumFilenameTooLong IOErrorEnum = 9
	IOErrorEnumInvalidFilename IOErrorEnum = 10
	IOErrorEnumTooManyLinks IOErrorEnum = 11
	IOErrorEnumNoSpace IOErrorEnum = 12
	IOErrorEnumInvalidArgument IOErrorEnum = 13
	IOErrorEnumPermissionDenied IOErrorEnum = 14
	IOErrorEnumNotSupported IOErrorEnum = 15
	IOErrorEnumNotMounted IOErrorEnum = 16
	IOErrorEnumAlreadyMounted IOErrorEnum = 17
	IOErrorEnumClosed IOErrorEnum = 18
	IOErrorEnumCancelled IOErrorEnum = 19
	IOErrorEnumPending IOErrorEnum = 20
	IOErrorEnumReadOnly IOErrorEnum = 21
	IOErrorEnumCantCreateBackup IOErrorEnum = 22
	IOErrorEnumWrongEtag IOErrorEnum = 23
	IOErrorEnumTimedOut IOErrorEnum = 24
	IOErrorEnumWouldRecurse IOErrorEnum = 25
	IOErrorEnumBusy IOErrorEnum = 26
	IOErrorEnumWouldBlock IOErrorEnum = 27
	IOErrorEnumHostNotFound IOErrorEnum = 28
	IOErrorEnumWouldMerge IOErrorEnum = 29
	IOErrorEnumFailedHandled IOErrorEnum = 30
	IOErrorEnumTooManyOpenFiles IOErrorEnum = 31
	IOErrorEnumNotInitialized IOErrorEnum = 32
	IOErrorEnumAddressInUse IOErrorEnum = 33
	IOErrorEnumPartialInput IOErrorEnum = 34
	IOErrorEnumInvalidData IOErrorEnum = 35
	IOErrorEnumDbusError IOErrorEnum = 36
	IOErrorEnumHostUnreachable IOErrorEnum = 37
	IOErrorEnumNetworkUnreachable IOErrorEnum = 38
	IOErrorEnumConnectionRefused IOErrorEnum = 39
	IOErrorEnumProxyFailed IOErrorEnum = 40
	IOErrorEnumProxyAuthFailed IOErrorEnum = 41
	IOErrorEnumProxyNeedAuth IOErrorEnum = 42
	IOErrorEnumProxyNotAllowed IOErrorEnum = 43
)
// blacklisted: IOExtension (struct)
// blacklisted: IOExtensionPoint (struct)
// blacklisted: IOModule (object)
// blacklisted: IOModuleClass (struct)
// blacklisted: IOModuleScope (struct)
type IOModuleScopeFlags C.uint32_t
const (
	IOModuleScopeFlagsNone IOModuleScopeFlags = 0
	IOModuleScopeFlagsBlockDuplicates IOModuleScopeFlags = 1
)
// blacklisted: IOSchedulerJob (struct)
// blacklisted: IOSchedulerJobFunc (callback)
// blacklisted: IOStream (object)
// blacklisted: IOStreamAdapter (struct)
// blacklisted: IOStreamClass (struct)
// blacklisted: IOStreamPrivate (struct)
type IOStreamSpliceFlags C.uint32_t
const (
	IOStreamSpliceFlagsNone IOStreamSpliceFlags = 0
	IOStreamSpliceFlagsCloseStream1 IOStreamSpliceFlags = 1
	IOStreamSpliceFlagsCloseStream2 IOStreamSpliceFlags = 2
	IOStreamSpliceFlagsWaitForBoth IOStreamSpliceFlags = 4
)
type IconLike interface {
	ImplementsGIcon() *C.GIcon
}

type Icon struct {
	gobject.Object
	IconImpl
}

type IconImpl struct {}

func ToIcon(objlike gobject.ObjectLike) *Icon {
	t := (*IconImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Icon)(obj)
	}
	panic("cannot cast to Icon")
}

func (this0 *IconImpl) ImplementsGIcon() *C.GIcon {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GIcon)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *IconImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.g_icon_get_type())
}

func IconGetType() gobject.Type {
	return (*IconImpl)(nil).GetStaticType()
}
// blacklisted: Icon.hash (method)
// blacklisted: Icon.new_for_string (method)
// blacklisted: Icon.equal (method)
// blacklisted: Icon.to_string (method)
// blacklisted: IconIface (struct)
// blacklisted: InetAddress (object)
// blacklisted: InetAddressClass (struct)
// blacklisted: InetAddressPrivate (struct)
// blacklisted: InetSocketAddress (object)
// blacklisted: InetSocketAddressClass (struct)
// blacklisted: InetSocketAddressPrivate (struct)
// blacklisted: Initable (interface)
// blacklisted: InitableIface (struct)
// blacklisted: InputStream (object)
// blacklisted: InputStreamClass (struct)
// blacklisted: InputStreamPrivate (struct)
// blacklisted: InputVector (struct)
// blacklisted: LoadableIcon (interface)
// blacklisted: LoadableIconIface (struct)
// blacklisted: MemoryInputStream (object)
// blacklisted: MemoryInputStreamClass (struct)
// blacklisted: MemoryInputStreamPrivate (struct)
// blacklisted: MemoryOutputStream (object)
// blacklisted: MemoryOutputStreamClass (struct)
// blacklisted: MemoryOutputStreamPrivate (struct)
// blacklisted: Mount (interface)
// blacklisted: MountIface (struct)
type MountMountFlags C.uint32_t
const (
	MountMountFlagsNone MountMountFlags = 0
)
type MountOperationLike interface {
	gobject.ObjectLike
	InheritedFromGMountOperation() *C.GMountOperation
}

type MountOperation struct {
	gobject.Object
	
}

func ToMountOperation(objlike gobject.ObjectLike) *MountOperation {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*MountOperation)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*MountOperation)(obj)
	}
	panic("cannot cast to MountOperation")
}

func (this0 *MountOperation) InheritedFromGMountOperation() *C.GMountOperation {
	if this0 == nil {
		return nil
	}
	return (*C.GMountOperation)(this0.C)
}

func (this0 *MountOperation) GetStaticType() gobject.Type {
	return gobject.Type(C.g_mount_operation_get_type())
}

func MountOperationGetType() gobject.Type {
	return (*MountOperation)(nil).GetStaticType()
}
func NewMountOperation() *MountOperation {
	ret1 := C.g_mount_operation_new()
	var ret2 *MountOperation
	ret2 = (*MountOperation)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MountOperation) GetAnonymous() bool {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_anonymous(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *MountOperation) GetChoice() int {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_choice(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *MountOperation) GetDomain() string {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_domain(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MountOperation) GetPassword() string {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_password(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MountOperation) GetPasswordSave() PasswordSave {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_password_save(this1)
	var ret2 PasswordSave
	ret2 = PasswordSave(ret1)
	return ret2
}
func (this0 *MountOperation) GetUsername() string {
	var this1 *C.GMountOperation
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	ret1 := C.g_mount_operation_get_username(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MountOperation) Reply(result0 MountOperationResult) {
	var this1 *C.GMountOperation
	var result1 C.GMountOperationResult
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	result1 = C.GMountOperationResult(result0)
	C.g_mount_operation_reply(this1, result1)
}
func (this0 *MountOperation) SetAnonymous(anonymous0 bool) {
	var this1 *C.GMountOperation
	var anonymous1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	anonymous1 = _GoBoolToCBool(anonymous0)
	C.g_mount_operation_set_anonymous(this1, anonymous1)
}
func (this0 *MountOperation) SetChoice(choice0 int) {
	var this1 *C.GMountOperation
	var choice1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	choice1 = C.int32_t(choice0)
	C.g_mount_operation_set_choice(this1, choice1)
}
func (this0 *MountOperation) SetDomain(domain0 string) {
	var this1 *C.GMountOperation
	var domain1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	domain1 = _GoStringToGString(domain0)
	defer C.free(unsafe.Pointer(domain1))
	C.g_mount_operation_set_domain(this1, domain1)
}
func (this0 *MountOperation) SetPassword(password0 string) {
	var this1 *C.GMountOperation
	var password1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	password1 = _GoStringToGString(password0)
	defer C.free(unsafe.Pointer(password1))
	C.g_mount_operation_set_password(this1, password1)
}
func (this0 *MountOperation) SetPasswordSave(save0 PasswordSave) {
	var this1 *C.GMountOperation
	var save1 C.GPasswordSave
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	save1 = C.GPasswordSave(save0)
	C.g_mount_operation_set_password_save(this1, save1)
}
func (this0 *MountOperation) SetUsername(username0 string) {
	var this1 *C.GMountOperation
	var username1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGMountOperation()
	}
	username1 = _GoStringToGString(username0)
	defer C.free(unsafe.Pointer(username1))
	C.g_mount_operation_set_username(this1, username1)
}
// blacklisted: MountOperationClass (struct)
// blacklisted: MountOperationPrivate (struct)
type MountOperationResult C.uint32_t
const (
	MountOperationResultHandled MountOperationResult = 0
	MountOperationResultAborted MountOperationResult = 1
	MountOperationResultUnhandled MountOperationResult = 2
)
type MountUnmountFlags C.uint32_t
const (
	MountUnmountFlagsNone MountUnmountFlags = 0
	MountUnmountFlagsForce MountUnmountFlags = 1
)
const NativeVolumeMonitorExtensionPointName = "gio-native-volume-monitor"
// blacklisted: NativeVolumeMonitor (object)
// blacklisted: NativeVolumeMonitorClass (struct)
// blacklisted: NetworkAddress (object)
// blacklisted: NetworkAddressClass (struct)
// blacklisted: NetworkAddressPrivate (struct)
// blacklisted: NetworkService (object)
// blacklisted: NetworkServiceClass (struct)
// blacklisted: NetworkServicePrivate (struct)
// blacklisted: OutputStream (object)
// blacklisted: OutputStreamClass (struct)
// blacklisted: OutputStreamPrivate (struct)
type OutputStreamSpliceFlags C.uint32_t
const (
	OutputStreamSpliceFlagsNone OutputStreamSpliceFlags = 0
	OutputStreamSpliceFlagsCloseSource OutputStreamSpliceFlags = 1
	OutputStreamSpliceFlagsCloseTarget OutputStreamSpliceFlags = 2
)
// blacklisted: OutputVector (struct)
const ProxyExtensionPointName = "gio-proxy"
const ProxyResolverExtensionPointName = "gio-proxy-resolver"
type PasswordSave C.uint32_t
const (
	PasswordSaveNever PasswordSave = 0
	PasswordSaveForSession PasswordSave = 1
	PasswordSavePermanently PasswordSave = 2
)
type PermissionLike interface {
	gobject.ObjectLike
	InheritedFromGPermission() *C.GPermission
}

type Permission struct {
	gobject.Object
	
}

func ToPermission(objlike gobject.ObjectLike) *Permission {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Permission)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Permission)(obj)
	}
	panic("cannot cast to Permission")
}

func (this0 *Permission) InheritedFromGPermission() *C.GPermission {
	if this0 == nil {
		return nil
	}
	return (*C.GPermission)(this0.C)
}

func (this0 *Permission) GetStaticType() gobject.Type {
	return gobject.Type(C.g_permission_get_type())
}

func PermissionGetType() gobject.Type {
	return (*Permission)(nil).GetStaticType()
}
func (this0 *Permission) Acquire(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GPermission
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if cancellable0 != nil {
		cancellable1 = cancellable0.InheritedFromGCancellable()
	}
	ret1 := C.g_permission_acquire(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Permission) AcquireAsync(cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GPermission
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if cancellable0 != nil {
		cancellable1 = cancellable0.InheritedFromGCancellable()
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_permission_acquire_async(this1, cancellable1, callback1)
}
func (this0 *Permission) AcquireFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GPermission
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()}
	ret1 := C.g_permission_acquire_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Permission) GetAllowed() bool {
	var this1 *C.GPermission
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	ret1 := C.g_permission_get_allowed(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Permission) GetCanAcquire() bool {
	var this1 *C.GPermission
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	ret1 := C.g_permission_get_can_acquire(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Permission) GetCanRelease() bool {
	var this1 *C.GPermission
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	ret1 := C.g_permission_get_can_release(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Permission) ImplUpdate(allowed0 bool, can_acquire0 bool, can_release0 bool) {
	var this1 *C.GPermission
	var allowed1 C.int
	var can_acquire1 C.int
	var can_release1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	allowed1 = _GoBoolToCBool(allowed0)
	can_acquire1 = _GoBoolToCBool(can_acquire0)
	can_release1 = _GoBoolToCBool(can_release0)
	C.g_permission_impl_update(this1, allowed1, can_acquire1, can_release1)
}
func (this0 *Permission) Release(cancellable0 CancellableLike) (bool, error) {
	var this1 *C.GPermission
	var cancellable1 *C.GCancellable
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if cancellable0 != nil {
		cancellable1 = cancellable0.InheritedFromGCancellable()
	}
	ret1 := C.g_permission_release(this1, cancellable1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Permission) ReleaseAsync(cancellable0 CancellableLike, callback0 AsyncReadyCallback) {
	var this1 *C.GPermission
	var cancellable1 *C.GCancellable
	var callback1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if cancellable0 != nil {
		cancellable1 = cancellable0.InheritedFromGCancellable()
	}
	if callback0 != nil {
		callback1 = unsafe.Pointer(&callback0)}
	gobject.Holder.Grab(callback1)
	C._g_permission_release_async(this1, cancellable1, callback1)
}
func (this0 *Permission) ReleaseFinish(result0 AsyncResultLike) (bool, error) {
	var this1 *C.GPermission
	var result1 *C.GAsyncResult
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGPermission()
	}
	if result0 != nil {
		result1 = result0.ImplementsGAsyncResult()}
	ret1 := C.g_permission_release_finish(this1, result1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
// blacklisted: PermissionClass (struct)
// blacklisted: PermissionPrivate (struct)
// blacklisted: PollableInputStream (interface)
// blacklisted: PollableInputStreamInterface (struct)
// blacklisted: PollableOutputStream (interface)
// blacklisted: PollableOutputStreamInterface (struct)
// blacklisted: PollableSourceFunc (callback)
// blacklisted: Proxy (interface)
// blacklisted: ProxyAddress (object)
// blacklisted: ProxyAddressClass (struct)
// blacklisted: ProxyAddressEnumerator (object)
// blacklisted: ProxyAddressEnumeratorClass (struct)
// blacklisted: ProxyAddressEnumeratorPrivate (struct)
// blacklisted: ProxyAddressPrivate (struct)
// blacklisted: ProxyInterface (struct)
// blacklisted: ProxyResolver (interface)
// blacklisted: ProxyResolverInterface (struct)
// blacklisted: Resolver (object)
// blacklisted: ResolverClass (struct)
type ResolverError C.uint32_t
const (
	ResolverErrorNotFound ResolverError = 0
	ResolverErrorTemporaryFailure ResolverError = 1
	ResolverErrorInternal ResolverError = 2
)
// blacklisted: ResolverPrivate (struct)
// blacklisted: Seekable (interface)
// blacklisted: SeekableIface (struct)
// blacklisted: Settings (object)
// blacklisted: SettingsBackend (struct)
type SettingsBindFlags C.uint32_t
const (
	SettingsBindFlagsDefault SettingsBindFlags = 0
	SettingsBindFlagsGet SettingsBindFlags = 1
	SettingsBindFlagsSet SettingsBindFlags = 2
	SettingsBindFlagsNoSensitivity SettingsBindFlags = 4
	SettingsBindFlagsGetNoChanges SettingsBindFlags = 8
	SettingsBindFlagsInvertBoolean SettingsBindFlags = 16
)
// blacklisted: SettingsBindGetMapping (callback)
// blacklisted: SettingsBindSetMapping (callback)
// blacklisted: SettingsClass (struct)
// blacklisted: SettingsGetMapping (callback)
// blacklisted: SettingsPrivate (struct)
// blacklisted: SimpleAction (object)
// blacklisted: SimpleActionGroup (object)
// blacklisted: SimpleActionGroupClass (struct)
// blacklisted: SimpleActionGroupPrivate (struct)
// blacklisted: SimpleAsyncResult (object)
// blacklisted: SimpleAsyncResultClass (struct)
// blacklisted: SimpleAsyncThreadFunc (callback)
// blacklisted: SimplePermission (object)
// blacklisted: Socket (object)
// blacklisted: SocketAddress (object)
// blacklisted: SocketAddressClass (struct)
// blacklisted: SocketAddressEnumerator (object)
// blacklisted: SocketAddressEnumeratorClass (struct)
// blacklisted: SocketClass (struct)
// blacklisted: SocketClient (object)
// blacklisted: SocketClientClass (struct)
// blacklisted: SocketClientPrivate (struct)
// blacklisted: SocketConnectable (interface)
// blacklisted: SocketConnectableIface (struct)
// blacklisted: SocketConnection (object)
// blacklisted: SocketConnectionClass (struct)
// blacklisted: SocketConnectionPrivate (struct)
// blacklisted: SocketControlMessage (object)
// blacklisted: SocketControlMessageClass (struct)
// blacklisted: SocketControlMessagePrivate (struct)
type SocketFamily C.uint32_t
const (
	SocketFamilyInvalid SocketFamily = 0
	SocketFamilyUnix SocketFamily = 1
	SocketFamilyIpv4 SocketFamily = 2
	SocketFamilyIpv6 SocketFamily = 10
)
// blacklisted: SocketListener (object)
// blacklisted: SocketListenerClass (struct)
// blacklisted: SocketListenerPrivate (struct)
type SocketMsgFlags C.uint32_t
const (
	SocketMsgFlagsNone SocketMsgFlags = 0
	SocketMsgFlagsOob SocketMsgFlags = 1
	SocketMsgFlagsPeek SocketMsgFlags = 2
	SocketMsgFlagsDontroute SocketMsgFlags = 4
)
// blacklisted: SocketPrivate (struct)
type SocketProtocol C.int32_t
const (
	SocketProtocolUnknown SocketProtocol = -1
	SocketProtocolDefault SocketProtocol = 0
	SocketProtocolTcp SocketProtocol = 6
	SocketProtocolUdp SocketProtocol = 17
	SocketProtocolSctp SocketProtocol = 132
)
// blacklisted: SocketService (object)
// blacklisted: SocketServiceClass (struct)
// blacklisted: SocketServicePrivate (struct)
// blacklisted: SocketSourceFunc (callback)
type SocketType C.uint32_t
const (
	SocketTypeInvalid SocketType = 0
	SocketTypeStream SocketType = 1
	SocketTypeDatagram SocketType = 2
	SocketTypeSeqpacket SocketType = 3
)
// blacklisted: SrvTarget (struct)
const TlsBackendExtensionPointName = "gio-tls-backend"
const TlsDatabasePurposeAuthenticateClient = "1.3.6.1.5.5.7.3.2"
const TlsDatabasePurposeAuthenticateServer = "1.3.6.1.5.5.7.3.1"
// blacklisted: TcpConnection (object)
// blacklisted: TcpConnectionClass (struct)
// blacklisted: TcpConnectionPrivate (struct)
// blacklisted: TcpWrapperConnection (object)
// blacklisted: TcpWrapperConnectionClass (struct)
// blacklisted: TcpWrapperConnectionPrivate (struct)
// blacklisted: ThemedIcon (object)
// blacklisted: ThemedIconClass (struct)
// blacklisted: ThreadedSocketService (object)
// blacklisted: ThreadedSocketServiceClass (struct)
// blacklisted: ThreadedSocketServicePrivate (struct)
type TlsAuthenticationMode C.uint32_t
const (
	TlsAuthenticationModeNone TlsAuthenticationMode = 0
	TlsAuthenticationModeRequested TlsAuthenticationMode = 1
	TlsAuthenticationModeRequired TlsAuthenticationMode = 2
)
// blacklisted: TlsBackend (interface)
// blacklisted: TlsBackendInterface (struct)
// blacklisted: TlsCertificate (object)
// blacklisted: TlsCertificateClass (struct)
type TlsCertificateFlags C.uint32_t
const (
	TlsCertificateFlagsUnknownCa TlsCertificateFlags = 1
	TlsCertificateFlagsBadIdentity TlsCertificateFlags = 2
	TlsCertificateFlagsNotActivated TlsCertificateFlags = 4
	TlsCertificateFlagsExpired TlsCertificateFlags = 8
	TlsCertificateFlagsRevoked TlsCertificateFlags = 16
	TlsCertificateFlagsInsecure TlsCertificateFlags = 32
	TlsCertificateFlagsGenericError TlsCertificateFlags = 64
	TlsCertificateFlagsValidateAll TlsCertificateFlags = 127
)
// blacklisted: TlsCertificatePrivate (struct)
// blacklisted: TlsClientConnection (interface)
// blacklisted: TlsClientConnectionInterface (struct)
// blacklisted: TlsConnection (object)
// blacklisted: TlsConnectionClass (struct)
// blacklisted: TlsConnectionPrivate (struct)
// blacklisted: TlsDatabase (object)
// blacklisted: TlsDatabaseClass (struct)
type TlsDatabaseLookupFlags C.uint32_t
const (
	TlsDatabaseLookupFlagsNone TlsDatabaseLookupFlags = 0
	TlsDatabaseLookupFlagsKeypair TlsDatabaseLookupFlags = 1
)
// blacklisted: TlsDatabasePrivate (struct)
type TlsDatabaseVerifyFlags C.uint32_t
const (
	TlsDatabaseVerifyFlagsNone TlsDatabaseVerifyFlags = 0
)
type TlsError C.uint32_t
const (
	TlsErrorUnavailable TlsError = 0
	TlsErrorMisc TlsError = 1
	TlsErrorBadCertificate TlsError = 2
	TlsErrorNotTls TlsError = 3
	TlsErrorHandshake TlsError = 4
	TlsErrorCertificateRequired TlsError = 5
	TlsErrorEof TlsError = 6
)
// blacklisted: TlsFileDatabase (interface)
// blacklisted: TlsFileDatabaseInterface (struct)
// blacklisted: TlsInteraction (object)
// blacklisted: TlsInteractionClass (struct)
// blacklisted: TlsInteractionPrivate (struct)
type TlsInteractionResult C.uint32_t
const (
	TlsInteractionResultUnhandled TlsInteractionResult = 0
	TlsInteractionResultHandled TlsInteractionResult = 1
	TlsInteractionResultFailed TlsInteractionResult = 2
)
// blacklisted: TlsPassword (object)
// blacklisted: TlsPasswordClass (struct)
type TlsPasswordFlags C.uint32_t
const (
	TlsPasswordFlagsNone TlsPasswordFlags = 0
	TlsPasswordFlagsRetry TlsPasswordFlags = 2
	TlsPasswordFlagsManyTries TlsPasswordFlags = 4
	TlsPasswordFlagsFinalTry TlsPasswordFlags = 8
)
// blacklisted: TlsPasswordPrivate (struct)
type TlsRehandshakeMode C.uint32_t
const (
	TlsRehandshakeModeNever TlsRehandshakeMode = 0
	TlsRehandshakeModeSafely TlsRehandshakeMode = 1
	TlsRehandshakeModeUnsafely TlsRehandshakeMode = 2
)
// blacklisted: TlsServerConnection (interface)
// blacklisted: TlsServerConnectionInterface (struct)
// blacklisted: UnixConnection (object)
// blacklisted: UnixConnectionClass (struct)
// blacklisted: UnixConnectionPrivate (struct)
// blacklisted: UnixCredentialsMessage (object)
// blacklisted: UnixCredentialsMessageClass (struct)
// blacklisted: UnixCredentialsMessagePrivate (struct)
// blacklisted: UnixFDList (object)
// blacklisted: UnixFDListClass (struct)
// blacklisted: UnixFDListPrivate (struct)
// blacklisted: UnixFDMessage (object)
// blacklisted: UnixFDMessageClass (struct)
// blacklisted: UnixFDMessagePrivate (struct)
// blacklisted: UnixInputStream (object)
// blacklisted: UnixInputStreamClass (struct)
// blacklisted: UnixInputStreamPrivate (struct)
// blacklisted: UnixMountEntry (struct)
// blacklisted: UnixMountMonitor (object)
// blacklisted: UnixMountMonitorClass (struct)
// blacklisted: UnixMountPoint (struct)
// blacklisted: UnixOutputStream (object)
// blacklisted: UnixOutputStreamClass (struct)
// blacklisted: UnixOutputStreamPrivate (struct)
// blacklisted: UnixSocketAddress (object)
// blacklisted: UnixSocketAddressClass (struct)
// blacklisted: UnixSocketAddressPrivate (struct)
type UnixSocketAddressType C.uint32_t
const (
	UnixSocketAddressTypeInvalid UnixSocketAddressType = 0
	UnixSocketAddressTypeAnonymous UnixSocketAddressType = 1
	UnixSocketAddressTypePath UnixSocketAddressType = 2
	UnixSocketAddressTypeAbstract UnixSocketAddressType = 3
	UnixSocketAddressTypeAbstractPadded UnixSocketAddressType = 4
)
const VfsExtensionPointName = "gio-vfs"
const VolumeIdentifierKindHalUdi = "hal-udi"
const VolumeIdentifierKindLabel = "label"
const VolumeIdentifierKindNfsMount = "nfs-mount"
const VolumeIdentifierKindUnixDevice = "unix-device"
const VolumeIdentifierKindUuid = "uuid"
const VolumeMonitorExtensionPointName = "gio-volume-monitor"
// blacklisted: Vfs (object)
// blacklisted: VfsClass (struct)
// blacklisted: Volume (interface)
// blacklisted: VolumeIface (struct)
// blacklisted: VolumeMonitor (object)
// blacklisted: VolumeMonitorClass (struct)
// blacklisted: ZlibCompressor (object)
// blacklisted: ZlibCompressorClass (struct)
type ZlibCompressorFormat C.uint32_t
const (
	ZlibCompressorFormatZlib ZlibCompressorFormat = 0
	ZlibCompressorFormatGzip ZlibCompressorFormat = 1
	ZlibCompressorFormatRaw ZlibCompressorFormat = 2
)
// blacklisted: ZlibDecompressor (object)
// blacklisted: ZlibDecompressorClass (struct)
// blacklisted: app_info_create_from_commandline (function)
// blacklisted: app_info_get_all (function)
// blacklisted: app_info_get_all_for_type (function)
// blacklisted: app_info_get_default_for_type (function)
// blacklisted: app_info_get_default_for_uri_scheme (function)
// blacklisted: app_info_get_fallback_for_type (function)
// blacklisted: app_info_get_recommended_for_type (function)
// blacklisted: app_info_launch_default_for_uri (function)
// blacklisted: app_info_reset_type_associations (function)
// blacklisted: async_initable_newv_async (function)
// blacklisted: bus_get (function)
// blacklisted: bus_get_finish (function)
// blacklisted: bus_get_sync (function)
// blacklisted: bus_own_name_on_connection (function)
// blacklisted: bus_own_name (function)
// blacklisted: bus_unown_name (function)
// blacklisted: bus_unwatch_name (function)
// blacklisted: bus_watch_name_on_connection (function)
// blacklisted: bus_watch_name (function)
// blacklisted: content_type_can_be_executable (function)
// blacklisted: content_type_equals (function)
// blacklisted: content_type_from_mime_type (function)
// blacklisted: content_type_get_description (function)
// blacklisted: content_type_get_icon (function)
// blacklisted: content_type_get_mime_type (function)
// blacklisted: content_type_guess (function)
// blacklisted: content_type_guess_for_tree (function)
// blacklisted: content_type_is_a (function)
// blacklisted: content_type_is_unknown (function)
// blacklisted: content_types_get_registered (function)
// blacklisted: dbus_address_get_for_bus_sync (function)
// blacklisted: dbus_address_get_stream (function)
// blacklisted: dbus_address_get_stream_finish (function)
// blacklisted: dbus_address_get_stream_sync (function)
// blacklisted: dbus_annotation_info_lookup (function)
// blacklisted: dbus_error_encode_gerror (function)
// blacklisted: dbus_error_get_remote_error (function)
// blacklisted: dbus_error_is_remote_error (function)
// blacklisted: dbus_error_new_for_dbus_error (function)
// blacklisted: dbus_error_quark (function)
// blacklisted: dbus_error_register_error (function)
// blacklisted: dbus_error_register_error_domain (function)
// blacklisted: dbus_error_strip_remote_error (function)
// blacklisted: dbus_error_unregister_error (function)
// blacklisted: dbus_generate_guid (function)
// blacklisted: dbus_gvalue_to_gvariant (function)
// blacklisted: dbus_gvariant_to_gvalue (function)
// blacklisted: dbus_is_address (function)
// blacklisted: dbus_is_guid (function)
// blacklisted: dbus_is_interface_name (function)
// blacklisted: dbus_is_member_name (function)
// blacklisted: dbus_is_name (function)
// blacklisted: dbus_is_supported_address (function)
// blacklisted: dbus_is_unique_name (function)
// blacklisted: file_hash (function)
// blacklisted: file_new_for_commandline_arg (function)
// blacklisted: file_new_for_path (function)
// blacklisted: file_new_for_uri (function)
// blacklisted: file_parse_name (function)
// blacklisted: icon_hash (function)
// blacklisted: icon_new_for_string (function)
// blacklisted: initable_newv (function)
// blacklisted: io_error_from_errno (function)
// blacklisted: io_error_quark (function)
// blacklisted: io_extension_get_type (function)
// blacklisted: io_extension_point_implement (function)
// blacklisted: io_extension_point_lookup (function)
// blacklisted: io_extension_point_register (function)
// blacklisted: io_modules_load_all_in_directory (function)
// blacklisted: io_modules_load_all_in_directory_with_scope (function)
// blacklisted: io_modules_scan_all_in_directory (function)
// blacklisted: io_modules_scan_all_in_directory_with_scope (function)
// blacklisted: io_scheduler_cancel_all_jobs (function)
// blacklisted: io_scheduler_push_job (function)
// blacklisted: proxy_get_default_for_protocol (function)
// blacklisted: proxy_resolver_get_default (function)
// blacklisted: resolver_error_quark (function)
// blacklisted: simple_async_report_gerror_in_idle (function)
// blacklisted: tls_backend_get_default (function)
// blacklisted: tls_client_connection_new (function)
// blacklisted: tls_error_quark (function)
// blacklisted: tls_file_database_new (function)
// blacklisted: tls_server_connection_new (function)
// blacklisted: unix_is_mount_path_system_internal (function)
// blacklisted: unix_mount_compare (function)
// blacklisted: unix_mount_free (function)
// blacklisted: unix_mount_get_device_path (function)
// blacklisted: unix_mount_get_fs_type (function)
// blacklisted: unix_mount_get_mount_path (function)
// blacklisted: unix_mount_guess_can_eject (function)
// blacklisted: unix_mount_guess_icon (function)
// blacklisted: unix_mount_guess_name (function)
// blacklisted: unix_mount_guess_should_display (function)
// blacklisted: unix_mount_is_readonly (function)
// blacklisted: unix_mount_is_system_internal (function)
// blacklisted: unix_mount_points_changed_since (function)
// blacklisted: unix_mounts_changed_since (function)
