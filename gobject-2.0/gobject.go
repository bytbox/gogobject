package gobject

/*
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

typedef size_t GType;


extern void _GObject_go_callback_cleanup(void *gofunc);
static void _c_callback_cleanup(void *userdata)
{
	_GObject_go_callback_cleanup(userdata);
}
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
typedef void* GClosureMarshal;
extern void _GClosureMarshal_c_wrapper();
extern void _GClosureMarshal_c_wrapper_once();
typedef void* GClosureNotify;
extern void _GClosureNotify_c_wrapper();
extern void _GClosureNotify_c_wrapper_once();
typedef struct _GClosureNotifyData GClosureNotifyData;
typedef uint32_t GConnectFlags;
typedef struct _GEnumClass GEnumClass;
typedef struct _GEnumValue GEnumValue;
typedef struct _GFlagsClass GFlagsClass;
typedef struct _GFlagsValue GFlagsValue;
typedef struct _GInitiallyUnowned GInitiallyUnowned;
typedef struct _GInitiallyUnownedClass GInitiallyUnownedClass;
typedef void* GInstanceInitFunc;
extern void _GInstanceInitFunc_c_wrapper();
extern void _GInstanceInitFunc_c_wrapper_once();
typedef void* GInterfaceFinalizeFunc;
extern void _GInterfaceFinalizeFunc_c_wrapper();
extern void _GInterfaceFinalizeFunc_c_wrapper_once();
typedef struct _GInterfaceInfo GInterfaceInfo;
typedef void* GInterfaceInitFunc;
extern void _GInterfaceInitFunc_c_wrapper();
extern void _GInterfaceInitFunc_c_wrapper_once();
typedef struct _GObject GObject;
typedef struct _GObjectClass GObjectClass;
typedef struct _GObjectConstructParam GObjectConstructParam;
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
typedef struct _GParamSpecString GParamSpecString;
typedef struct _GParamSpecTypeInfo GParamSpecTypeInfo;
typedef struct _GParamSpecUChar GParamSpecUChar;
typedef struct _GParamSpecUInt GParamSpecUInt;
typedef struct _GParamSpecUInt64 GParamSpecUInt64;
typedef struct _GParamSpecULong GParamSpecULong;
typedef struct _GParamSpecUnichar GParamSpecUnichar;
typedef struct _GParamSpecValueArray GParamSpecValueArray;
typedef struct _GParamSpecVariant GParamSpecVariant;
typedef struct _GParameter GParameter;
typedef void* GSignalAccumulator;
extern void _GSignalAccumulator_c_wrapper();
extern void _GSignalAccumulator_c_wrapper_once();
typedef void* GSignalEmissionHook;
extern void _GSignalEmissionHook_c_wrapper();
extern void _GSignalEmissionHook_c_wrapper_once();
typedef uint32_t GSignalFlags;
typedef struct _GSignalInvocationHint GSignalInvocationHint;
typedef uint32_t GSignalMatchType;
typedef struct _GSignalQuery GSignalQuery;
typedef void* GToggleNotify;
extern void _GToggleNotify_c_wrapper();
extern void _GToggleNotify_c_wrapper_once();
typedef struct _GTypeCValue GTypeCValue;
typedef struct _GTypeClass GTypeClass;
typedef void* GTypeClassCacheFunc;
extern void _GTypeClassCacheFunc_c_wrapper();
extern void _GTypeClassCacheFunc_c_wrapper_once();
typedef uint32_t GTypeDebugFlags;
typedef uint32_t GTypeFlags;
typedef uint32_t GTypeFundamentalFlags;
typedef struct _GTypeFundamentalInfo GTypeFundamentalInfo;
typedef struct _GTypeInfo GTypeInfo;
typedef struct _GTypeInstance GTypeInstance;
typedef struct _GTypeInterface GTypeInterface;
typedef void* GTypeInterfaceCheckFunc;
extern void _GTypeInterfaceCheckFunc_c_wrapper();
extern void _GTypeInterfaceCheckFunc_c_wrapper_once();
typedef struct _GTypeModule GTypeModule;
typedef struct _GTypeModuleClass GTypeModuleClass;
typedef struct _GTypePlugin GTypePlugin;
typedef struct _GTypePluginClass GTypePluginClass;
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
typedef struct _GTypeValueTable GTypeValueTable;
typedef struct _GValue GValue;
typedef struct _GValueArray GValueArray;
typedef void* GValueTransform;
extern void _GValueTransform_c_wrapper();
extern void _GValueTransform_c_wrapper_once();
typedef void* GWeakNotify;
extern void _GWeakNotify_c_wrapper();
extern void _GWeakNotify_c_wrapper_once();
typedef struct _G_Value__data__union G_Value__data__union;
extern GBindingFlags g_binding_get_flags(GBinding*);
extern GObject* g_binding_get_source(GBinding*);
extern char* g_binding_get_source_property(GBinding*);
extern GObject* g_binding_get_target(GBinding*);
extern char* g_binding_get_target_property(GBinding*);
extern GType g_binding_get_type();
extern void g_cclosure_marshal_BOOLEAN__BOXED_BOXED(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_BOOLEAN__FLAGS(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_STRING__OBJECT_POINTER(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__BOOLEAN(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__BOXED(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__CHAR(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__DOUBLE(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__ENUM(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__FLAGS(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__FLOAT(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__INT(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__LONG(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__OBJECT(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__PARAM(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__POINTER(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__STRING(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__UCHAR(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__UINT(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__UINT_POINTER(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__ULONG(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__VARIANT(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_VOID__VOID(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern void g_cclosure_marshal_generic(GClosure*, GValue*, uint32_t, GValue*, void*, void*);
extern GClosure* g_closure_new_object(uint32_t, GObject*);
extern GClosure* g_closure_new_simple(uint32_t, void*);
extern void g_closure_invalidate(GClosure*);
extern void g_closure_invoke(GClosure*, GValue*, uint32_t, GValue*, void*);
extern GClosure* g_closure_ref(GClosure*);
extern void g_closure_sink(GClosure*);
extern void g_closure_unref(GClosure*);
extern GType g_initially_unowned_get_type();
extern GObject* g_object_newv(GType, uint32_t, GParameter*);
extern GBinding* g_object_bind_property(void*, char*, void*, char*, GBindingFlags);
extern GBinding* g_object_bind_property_with_closures(void*, char*, void*, char*, GBindingFlags, GClosure*, GClosure*);
extern uint64_t g_object_compat_control(uint64_t, void*);
extern GParamSpec* g_object_interface_find_property(void*, char*);
extern void g_object_interface_install_property(void*, GParamSpec*);
extern GParamSpec** g_object_interface_list_properties(void*, uint32_t*);
extern void g_object_force_floating(GObject*);
extern void g_object_freeze_notify(GObject*);
extern void* g_object_get_data(GObject*, char*);
extern void g_object_get_property(GObject*, char*, GValue*);
extern void* g_object_get_qdata(GObject*, uint32_t);
extern int g_object_is_floating(GObject*);
extern void g_object_notify(GObject*, char*);
extern void g_object_notify_by_pspec(GObject*, GParamSpec*);
extern GObject* g_object_ref(GObject*);
extern GObject* g_object_ref_sink(GObject*);
extern void g_object_run_dispose(GObject*);
extern void g_object_set_data(GObject*, char*, void*);
extern void g_object_set_property(GObject*, char*, GValue*);
extern void* g_object_steal_data(GObject*, char*);
extern void* g_object_steal_qdata(GObject*, uint32_t);
extern void g_object_thaw_notify(GObject*);
extern void g_object_unref(GObject*);
extern void g_object_watch_closure(GObject*, GClosure*);
extern GType g_object_get_type();
extern GParamSpec* g_object_class_find_property(GObjectClass*, char*);
extern void g_object_class_install_properties(GObjectClass*, uint32_t, GParamSpec**);
extern void g_object_class_install_property(GObjectClass*, uint32_t, GParamSpec*);
extern GParamSpec** g_object_class_list_properties(GObjectClass*, uint32_t*);
extern void g_object_class_override_property(GObjectClass*, uint32_t, char*);
extern char* g_param_spec_get_blurb(GParamSpec*);
extern char* g_param_spec_get_name(GParamSpec*);
extern char* g_param_spec_get_nick(GParamSpec*);
extern void* g_param_spec_get_qdata(GParamSpec*, uint32_t);
extern GParamSpec* g_param_spec_get_redirect_target(GParamSpec*);
extern void g_param_spec_set_qdata(GParamSpec*, uint32_t, void*);
extern void g_param_spec_sink(GParamSpec*);
extern void* g_param_spec_steal_qdata(GParamSpec*, uint32_t);
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern void g_param_spec_pool_insert(GParamSpecPool*, GParamSpec*, GType);
extern GParamSpec** g_param_spec_pool_list(GParamSpecPool*, GType, uint32_t*);
extern GList* g_param_spec_pool_list_owned(GParamSpecPool*, GType);
extern GParamSpec* g_param_spec_pool_lookup(GParamSpecPool*, char*, GType, int);
extern void g_param_spec_pool_remove(GParamSpecPool*, GParamSpec*);
extern GParamSpecPool* g_param_spec_pool_new(int);
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GType intern();
extern GTypeClass* g_type_class_peek_parent(GTypeClass*);
extern void g_type_class_unref(GTypeClass*);
extern void g_type_class_add_private(void*, uint64_t);
extern GTypeClass* g_type_class_peek(GType);
extern GTypeClass* g_type_class_peek_static(GType);
extern GTypeClass* g_type_class_ref(GType);
extern GTypeInterface* g_type_interface_peek_parent(GTypeInterface*);
extern void g_type_interface_add_prerequisite(GType, GType);
extern GTypePlugin* g_type_interface_get_plugin(GType, GType);
extern GTypeInterface* g_type_interface_peek(GTypeClass*, GType);
extern GType* g_type_interface_prerequisites(GType, uint32_t*);
extern void g_type_module_add_interface(GTypeModule*, GType, GType, GInterfaceInfo*);
extern GType g_type_module_register_enum(GTypeModule*, char*, GEnumValue*);
extern GType g_type_module_register_flags(GTypeModule*, char*, GFlagsValue*);
extern GType g_type_module_register_type(GTypeModule*, GType, char*, GTypeInfo*, GTypeFlags);
extern void g_type_module_set_name(GTypeModule*, char*);
extern void g_type_module_unuse(GTypeModule*);
extern int g_type_module_use(GTypeModule*);
extern GType g_type_module_get_type();
extern void g_type_plugin_complete_interface_info(GTypePlugin*, GType, GType, GInterfaceInfo*);
extern void g_type_plugin_complete_type_info(GTypePlugin*, GType, GTypeInfo*, GTypeValueTable*);
extern void g_type_plugin_unuse(GTypePlugin*);
extern void g_type_plugin_use(GTypePlugin*);
extern GType g_type_plugin_get_type();
extern void g_value_copy(GValue*, GValue*);
extern GObject* g_value_dup_object(GValue*);
extern char* g_value_dup_string(GValue*);
extern GVariant* g_value_dup_variant(GValue*);
extern int g_value_fits_pointer(GValue*);
extern int g_value_get_boolean(GValue*);
extern void* g_value_get_boxed(GValue*);
extern uint8_t g_value_get_char(GValue*);
extern double g_value_get_double(GValue*);
extern int32_t g_value_get_enum(GValue*);
extern uint32_t g_value_get_flags(GValue*);
extern float g_value_get_float(GValue*);
extern int32_t g_value_get_int(GValue*);
extern int64_t g_value_get_int64(GValue*);
extern int64_t g_value_get_long(GValue*);
extern GObject* g_value_get_object(GValue*);
extern GParamSpec* g_value_get_param(GValue*);
extern void* g_value_get_pointer(GValue*);
extern char* g_value_get_string(GValue*);
extern uint8_t g_value_get_uchar(GValue*);
extern uint32_t g_value_get_uint(GValue*);
extern uint64_t g_value_get_uint64(GValue*);
extern uint64_t g_value_get_ulong(GValue*);
extern GVariant* g_value_get_variant(GValue*);
extern GValue* g_value_init(GValue*, GType);
extern void* g_value_peek_pointer(GValue*);
extern GValue* g_value_reset(GValue*);
extern void g_value_set_boolean(GValue*, int);
extern void g_value_set_boxed(GValue*, void*);
extern void g_value_set_boxed_take_ownership(GValue*, void*);
extern void g_value_set_char(GValue*, uint8_t);
extern void g_value_set_double(GValue*, double);
extern void g_value_set_enum(GValue*, int32_t);
extern void g_value_set_flags(GValue*, uint32_t);
extern void g_value_set_float(GValue*, float);
extern void g_value_set_gtype(GValue*, GType);
extern void g_value_set_instance(GValue*, void*);
extern void g_value_set_int(GValue*, int32_t);
extern void g_value_set_int64(GValue*, int64_t);
extern void g_value_set_long(GValue*, int64_t);
extern void g_value_set_object(GValue*, GObject*);
extern void g_value_set_param(GValue*, GParamSpec*);
extern void g_value_set_pointer(GValue*, void*);
extern void g_value_set_static_boxed(GValue*, void*);
extern void g_value_set_static_string(GValue*, char*);
extern void g_value_set_string(GValue*, char*);
extern void g_value_set_string_take_ownership(GValue*, char*);
extern void g_value_set_uchar(GValue*, uint8_t);
extern void g_value_set_uint(GValue*, uint32_t);
extern void g_value_set_uint64(GValue*, uint64_t);
extern void g_value_set_ulong(GValue*, uint64_t);
extern void g_value_set_variant(GValue*, GVariant*);
extern void g_value_take_boxed(GValue*, void*);
extern void g_value_take_string(GValue*, char*);
extern void g_value_take_variant(GValue*, GVariant*);
extern int g_value_transform(GValue*, GValue*);
extern void g_value_unset(GValue*);
extern int g_value_type_compatible(GType, GType);
extern int g_value_type_transformable(GType, GType);
extern GValueArray* g_value_array_new(uint32_t);
extern GValueArray* g_value_array_append(GValueArray*, GValue*);
extern GValueArray* g_value_array_copy(GValueArray*);
extern void g_value_array_free(GValueArray*);
extern GValue* g_value_array_get_nth(GValueArray*, uint32_t);
extern GValueArray* g_value_array_insert(GValueArray*, uint32_t, GValue*);
extern GValueArray* g_value_array_prepend(GValueArray*, GValue*);
extern GValueArray* g_value_array_remove(GValueArray*, uint32_t);
extern GValueArray* g_value_array_sort_with_data(GValueArray*, GCompareDataFunc, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static GValueArray* _g_value_array_sort_with_data(GValueArray* this, void* gofunc) {
	if (gofunc) {
		return g_value_array_sort_with_data(this, _GCompareDataFunc_c_wrapper, gofunc);
	} else {
		return g_value_array_sort_with_data(this, 0, 0);
	}
}
extern void g_boxed_free(GType, void*);
extern void g_enum_complete_type_info(GType, GTypeInfo*, GEnumValue*);
extern GType g_enum_register_static(char*, GEnumValue*);
extern void g_flags_complete_type_info(GType, GTypeInfo*, GFlagsValue*);
extern GType g_flags_register_static(char*, GFlagsValue*);
extern GType g_gtype_get_type();
extern GType g_param_type_register_static(char*, GParamSpecTypeInfo*);
extern int g_param_value_convert(GParamSpec*, GValue*, GValue*, int);
extern int g_param_value_defaults(GParamSpec*, GValue*);
extern void g_param_value_set_default(GParamSpec*, GValue*);
extern int g_param_value_validate(GParamSpec*, GValue*);
extern int32_t g_param_values_cmp(GParamSpec*, GValue*, GValue*);
extern GType g_pointer_type_register_static(char*);
extern int g_signal_accumulator_first_wins(GSignalInvocationHint*, GValue*, GValue*, void*);
extern int g_signal_accumulator_true_handled(GSignalInvocationHint*, GValue*, GValue*, void*);
extern uint64_t g_signal_add_emission_hook(uint32_t, uint32_t, GSignalEmissionHook, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static uint64_t _g_signal_add_emission_hook(uint32_t arg0, uint32_t arg1, void* gofunc) {
	if (gofunc) {
		return g_signal_add_emission_hook(arg0, arg1, _GSignalEmissionHook_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		return g_signal_add_emission_hook(arg0, arg1, 0, 0, 0);
	}
}
extern void g_signal_chain_from_overridden(GValue*, GValue*);
extern uint64_t g_signal_connect_closure(void*, char*, GClosure*, int);
extern uint64_t g_signal_connect_closure_by_id(void*, uint32_t, uint32_t, GClosure*, int);
extern void g_signal_emitv(GValue*, uint32_t, uint32_t, GValue*);
extern GSignalInvocationHint* g_signal_get_invocation_hint(void*);
extern void g_signal_handler_block(void*, uint64_t);
extern void g_signal_handler_disconnect(void*, uint64_t);
extern uint64_t g_signal_handler_find(void*, GSignalMatchType, uint32_t, uint32_t, GClosure*, void*, void*);
extern int g_signal_handler_is_connected(void*, uint64_t);
extern void g_signal_handler_unblock(void*, uint64_t);
extern uint32_t g_signal_handlers_block_matched(void*, GSignalMatchType, uint32_t, uint32_t, GClosure*, void*, void*);
extern void g_signal_handlers_destroy(void*);
extern uint32_t g_signal_handlers_disconnect_matched(void*, GSignalMatchType, uint32_t, uint32_t, GClosure*, void*, void*);
extern uint32_t g_signal_handlers_unblock_matched(void*, GSignalMatchType, uint32_t, uint32_t, GClosure*, void*, void*);
extern int g_signal_has_handler_pending(void*, uint32_t, uint32_t, int);
extern uint32_t* g_signal_list_ids(GType, uint32_t*);
extern uint32_t g_signal_lookup(char*, GType);
extern char* g_signal_name(uint32_t);
extern void g_signal_override_class_closure(uint32_t, GType, GClosure*);
extern int g_signal_parse_name(char*, GType, uint32_t*, uint32_t*, int);
extern void g_signal_query(uint32_t, GSignalQuery*);
extern void g_signal_remove_emission_hook(uint32_t, uint64_t);
extern void g_signal_stop_emission(void*, uint32_t, uint32_t);
extern void g_signal_stop_emission_by_name(void*, char*);
extern GClosure* g_signal_type_cclosure_new(GType, uint32_t);
extern void g_source_set_closure(GSource*, GClosure*);
extern void g_source_set_dummy_callback(GSource*);
extern char* g_strdup_value_contents(GValue*);
extern void g_type_add_class_private(GType, uint64_t);
extern void g_type_add_interface_dynamic(GType, GType, GTypePlugin*);
extern void g_type_add_interface_static(GType, GType, GInterfaceInfo*);
extern int g_type_check_class_is_a(GTypeClass*, GType);
extern int g_type_check_instance(GTypeInstance*);
extern int g_type_check_instance_is_a(GTypeInstance*, GType);
extern int g_type_check_is_value_type(GType);
extern int g_type_check_value(GValue*);
extern int g_type_check_value_holds(GValue*, GType);
extern GType* g_type_children(GType, uint32_t*);
extern GTypeInterface* g_type_default_interface_peek(GType);
extern GTypeInterface* g_type_default_interface_ref(GType);
extern void g_type_default_interface_unref(GTypeInterface*);
extern uint32_t g_type_depth(GType);
extern void g_type_free_instance(GTypeInstance*);
extern GType g_type_from_name(char*);
extern GType g_type_fundamental(GType);
extern GType g_type_fundamental_next();
extern GTypePlugin* g_type_get_plugin(GType);
extern void* g_type_get_qdata(GType, uint32_t);
extern void g_type_init();
extern void g_type_init_with_debug_flags(GTypeDebugFlags);
extern GType* g_type_interfaces(GType, uint32_t*);
extern int g_type_is_a(GType, GType);
extern char* g_type_name(GType);
extern char* g_type_name_from_class(GTypeClass*);
extern char* g_type_name_from_instance(GTypeInstance*);
extern GType g_type_next_base(GType, GType);
extern GType g_type_parent(GType);
extern uint32_t g_type_qname(GType);
extern void g_type_query(GType, GTypeQuery*);
extern GType g_type_register_dynamic(GType, char*, GTypePlugin*, GTypeFlags);
extern GType g_type_register_fundamental(GType, char*, GTypeInfo*, GTypeFundamentalInfo*, GTypeFlags);
extern GType g_type_register_static(GType, char*, GTypeInfo*, GTypeFlags);
extern void g_type_set_qdata(GType, uint32_t, void*);
extern int g_type_test_flags(GType, uint32_t);
extern GType g_value_get_gtype(GValue*);
struct _GCClosure { uint8_t _data[72]; };
struct _GClosure { uint8_t _data[64]; };
struct _GClosureNotifyData { uint8_t _data[16]; };
struct _GEnumClass { uint8_t _data[32]; };
struct _GEnumValue { uint8_t _data[24]; };
struct _GFlagsClass { uint8_t _data[24]; };
struct _GFlagsValue { uint8_t _data[24]; };
struct _GInitiallyUnownedClass { uint8_t _data[136]; };
struct _GInterfaceInfo { uint8_t _data[24]; };
struct _GObjectClass { uint8_t _data[136]; };
struct _GObjectConstructParam { uint8_t _data[16]; };
struct _GParamSpecClass { uint8_t _data[80]; };
struct _GParamSpecPool {};
struct _GParamSpecTypeInfo { uint8_t _data[56]; };
struct _GParameter { uint8_t _data[32]; };
struct _GSignalInvocationHint { uint8_t _data[12]; };
struct _GSignalQuery { uint8_t _data[56]; };
struct _GTypeCValue { uint8_t _data[8]; };
struct _GTypeClass { uint8_t _data[8]; };
struct _GTypeFundamentalInfo { uint8_t _data[4]; };
struct _GTypeInfo { uint8_t _data[72]; };
struct _GTypeInstance { uint8_t _data[8]; };
struct _GTypeInterface { uint8_t _data[16]; };
struct _GTypeModuleClass { uint8_t _data[184]; };
struct _GTypePluginClass { uint8_t _data[48]; };
struct _GTypeQuery { uint8_t _data[24]; };
struct _GTypeValueTable { uint8_t _data[64]; };
struct _GValue { uint8_t _data[24]; };
struct _GValueArray { uint8_t _data[24]; };
struct _G_Value__data__union { uint8_t _data[8]; };


extern void g_free(void*);

#include "gobject.h"

extern uint32_t g_quark_from_string(const char*);
extern void g_object_set_qdata(GObject*, uint32_t, void*);

extern void g_type_init();

extern GParamSpec *g_param_spec_ref_sink(GParamSpec*);
extern void g_param_spec_unref(GParamSpec*);

typedef int32_t (*_GSourceFunc)(void*);
extern uint32_t g_timeout_add(uint32_t, _GSourceFunc, void*);
extern int32_t g_source_remove(uint32_t);

extern int32_t fqueue_dispatcher(void*);
static uint32_t _g_timeout_add_fqueue(uint32_t time) {
	return g_timeout_add(time, fqueue_dispatcher, 0);
}

#cgo pkg-config: gobject-2.0
*/
import "C"
import "unsafe"
import "runtime"
import "reflect"
import "sync"

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


//export _GObject_go_callback_cleanup
func _GObject_go_callback_cleanup(gofunc unsafe.Pointer) {
	Holder.Release(gofunc)
}


// blacklisted: BaseFinalizeFunc (callback)
// blacklisted: BaseInitFunc (callback)
type BindingLike interface {
	ObjectLike
	InheritedFromGBinding() *C.GBinding
}

type Binding struct {
	Object
	
}

func ToBinding(objlike ObjectLike) *Binding {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Binding)(nil).GetStaticType()
	obj := ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Binding)(obj)
	}
	panic("cannot cast to Binding")
}

func (this0 *Binding) InheritedFromGBinding() *C.GBinding {
	if this0 == nil {
		return nil
	}
	return (*C.GBinding)(this0.C)
}

func (this0 *Binding) GetStaticType() Type {
	return Type(C.g_binding_get_type())
}

func BindingGetType() Type {
	return (*Binding)(nil).GetStaticType()
}
func (this0 *Binding) GetFlags() BindingFlags {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = this0.InheritedFromGBinding()
	}
	ret1 := C.g_binding_get_flags(this1)
	var ret2 BindingFlags
	ret2 = BindingFlags(ret1)
	return ret2
}
func (this0 *Binding) GetSource() *Object {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = this0.InheritedFromGBinding()
	}
	ret1 := C.g_binding_get_source(this1)
	var ret2 *Object
	ret2 = (*Object)(ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Binding) GetSourceProperty() string {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = this0.InheritedFromGBinding()
	}
	ret1 := C.g_binding_get_source_property(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Binding) GetTarget() *Object {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = this0.InheritedFromGBinding()
	}
	ret1 := C.g_binding_get_target(this1)
	var ret2 *Object
	ret2 = (*Object)(ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Binding) GetTargetProperty() string {
	var this1 *C.GBinding
	if this0 != nil {
		this1 = this0.InheritedFromGBinding()
	}
	ret1 := C.g_binding_get_target_property(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
type BindingFlags C.uint32_t
const (
	BindingFlagsDefault BindingFlags = 0
	BindingFlagsBidirectional BindingFlags = 1
	BindingFlagsSyncCreate BindingFlags = 2
	BindingFlagsInvertBoolean BindingFlags = 4
)
// blacklisted: BindingTransformFunc (callback)
// blacklisted: BoxedFreeFunc (callback)
type CClosure struct {
	Closure Closure
	Callback unsafe.Pointer
}
func CClosureMarshalBoolean_BoxedBoxed(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_BOOLEAN__BOXED_BOXED(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalBoolean_Flags(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_BOOLEAN__FLAGS(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalString_ObjectPointer(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_STRING__OBJECT_POINTER(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Boolean(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__BOOLEAN(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Boxed(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__BOXED(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Char(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__CHAR(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Double(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__DOUBLE(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Enum(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__ENUM(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Flags(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__FLAGS(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Float(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__FLOAT(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Int(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__INT(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Long(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__LONG(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Object(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__OBJECT(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Param(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__PARAM(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Pointer(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__POINTER(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_String(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__STRING(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Uchar(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__UCHAR(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Uint(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__UINT(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_UintPointer(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__UINT_POINTER(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Ulong(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__ULONG(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Variant(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__VARIANT(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalVoid_Void(closure0 *Closure, return_value0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_value1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_VOID__VOID(closure1, return_value1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
func CClosureMarshalGeneric(closure0 *Closure, return_gvalue0 *Value, n_param_values0 int, param_values0 *Value, invocation_hint0 unsafe.Pointer, marshal_data0 unsafe.Pointer) {
	var closure1 *C.GClosure
	var return_gvalue1 *C.GValue
	var n_param_values1 C.uint32_t
	var param_values1 *C.GValue
	var invocation_hint1 unsafe.Pointer
	var marshal_data1 unsafe.Pointer
	closure1 = (*C.GClosure)(unsafe.Pointer(closure0))
	return_gvalue1 = (*C.GValue)(unsafe.Pointer(return_gvalue0))
	n_param_values1 = C.uint32_t(n_param_values0)
	param_values1 = (*C.GValue)(unsafe.Pointer(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	marshal_data1 = unsafe.Pointer(marshal_data0)
	C.g_cclosure_marshal_generic(closure1, return_gvalue1, n_param_values1, param_values1, invocation_hint1, marshal_data1)
}
// blacklisted: Callback (callback)
// blacklisted: ClassFinalizeFunc (callback)
// blacklisted: ClassInitFunc (callback)
type Closure struct {
	RefCount uint32
	MetaMarshal uint32
	NGuards uint32
	NFnotifiers uint32
	NInotifiers uint32
	InInotify uint32
	Floating uint32
	DerivativeFlag uint32
	InMarshal uint32
	IsInvalid uint32
	Marshal unsafe.Pointer
	Data unsafe.Pointer
	Notifiers *ClosureNotifyData
}
func NewClosureObject(sizeof_closure0 int, object0 ObjectLike) *Closure {
	var sizeof_closure1 C.uint32_t
	var object1 *C.GObject
	sizeof_closure1 = C.uint32_t(sizeof_closure0)
	if object0 != nil {
		object1 = object0.InheritedFromGObject()
	}
	ret1 := C.g_closure_new_object(sizeof_closure1, object1)
	var ret2 *Closure
	ret2 = (*Closure)(unsafe.Pointer(ret1))
	return ret2
}
func NewClosureSimple(sizeof_closure0 int, data0 unsafe.Pointer) *Closure {
	var sizeof_closure1 C.uint32_t
	var data1 unsafe.Pointer
	sizeof_closure1 = C.uint32_t(sizeof_closure0)
	data1 = unsafe.Pointer(data0)
	ret1 := C.g_closure_new_simple(sizeof_closure1, data1)
	var ret2 *Closure
	ret2 = (*Closure)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Closure) Invalidate() {
	var this1 *C.GClosure
	this1 = (*C.GClosure)(unsafe.Pointer(this0))
	C.g_closure_invalidate(this1)
}
func (this0 *Closure) Invoke(return_value0 *Value, param_values0 []Value, invocation_hint0 unsafe.Pointer) {
	var this1 *C.GClosure
	var return_value1 *C.GValue
	var param_values1 *C.GValue
	var n_param_values1 C.uint32_t
	var invocation_hint1 unsafe.Pointer
	this1 = (*C.GClosure)(unsafe.Pointer(this0))
	return_value1 = (*C.GValue)(unsafe.Pointer(return_value0))
	param_values1 = (*C.GValue)(C.malloc(C.size_t(int(unsafe.Sizeof(*param_values1)) * len(param_values0))))
	defer C.free(unsafe.Pointer(param_values1))
	for i, e := range param_values0 {
		(*(*[999999]C.GValue)(unsafe.Pointer(param_values1)))[i] = *(*C.GValue)(unsafe.Pointer(&e))
	}
	n_param_values1 = C.uint32_t(len(param_values0))
	invocation_hint1 = unsafe.Pointer(invocation_hint0)
	C.g_closure_invoke(this1, return_value1, n_param_values1, param_values1, invocation_hint1)
}
func (this0 *Closure) Sink() {
	var this1 *C.GClosure
	this1 = (*C.GClosure)(unsafe.Pointer(this0))
	C.g_closure_sink(this1)
}
// blacklisted: ClosureMarshal (callback)
// blacklisted: ClosureNotify (callback)
type ClosureNotifyData struct {
	Data unsafe.Pointer
	Notify unsafe.Pointer
}
type ConnectFlags C.uint32_t
const (
	ConnectFlagsAfter ConnectFlags = 1
	ConnectFlagsSwapped ConnectFlags = 2
)
type EnumClass struct {
	GTypeClass TypeClass
	Minimum int32
	Maximum int32
	NValues uint32
	_ [4]byte
	Values *EnumValue
}
type EnumValue struct {
	Value int32
	_ [4]byte
	value_name0 *C.char
	value_nick0 *C.char
}
func (this0 *EnumValue) ValueName() string {
	var value_name1 string
	value_name1 = C.GoString(this0.value_name0)
	return value_name1
}
func (this0 *EnumValue) ValueNick() string {
	var value_nick1 string
	value_nick1 = C.GoString(this0.value_nick0)
	return value_nick1
}
type FlagsClass struct {
	GTypeClass TypeClass
	Mask uint32
	NValues uint32
	Values *FlagsValue
}
type FlagsValue struct {
	Value uint32
	_ [4]byte
	value_name0 *C.char
	value_nick0 *C.char
}
func (this0 *FlagsValue) ValueName() string {
	var value_name1 string
	value_name1 = C.GoString(this0.value_name0)
	return value_name1
}
func (this0 *FlagsValue) ValueNick() string {
	var value_nick1 string
	value_nick1 = C.GoString(this0.value_nick0)
	return value_nick1
}
type InitiallyUnownedLike interface {
	ObjectLike
	InheritedFromGInitiallyUnowned() *C.GInitiallyUnowned
}

type InitiallyUnowned struct {
	Object
	
}

func ToInitiallyUnowned(objlike ObjectLike) *InitiallyUnowned {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*InitiallyUnowned)(nil).GetStaticType()
	obj := ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*InitiallyUnowned)(obj)
	}
	panic("cannot cast to InitiallyUnowned")
}

func (this0 *InitiallyUnowned) InheritedFromGInitiallyUnowned() *C.GInitiallyUnowned {
	if this0 == nil {
		return nil
	}
	return (*C.GInitiallyUnowned)(this0.C)
}

func (this0 *InitiallyUnowned) GetStaticType() Type {
	return Type(C.g_initially_unowned_get_type())
}

func InitiallyUnownedGetType() Type {
	return (*InitiallyUnowned)(nil).GetStaticType()
}
// blacklisted: InstanceInitFunc (callback)
// blacklisted: InterfaceFinalizeFunc (callback)
type InterfaceInfo struct {
	InterfaceInit unsafe.Pointer
	InterfaceFinalize unsafe.Pointer
	InterfaceData unsafe.Pointer
}
// blacklisted: InterfaceInitFunc (callback)
type ObjectLike interface {
	
	InheritedFromGObject() *C.GObject
}

type Object struct {
	C unsafe.Pointer
	
}

func ToObject(objlike ObjectLike) *Object {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Object)(nil).GetStaticType()
	obj := ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Object)(obj)
	}
	panic("cannot cast to Object")
}

func (this0 *Object) InheritedFromGObject() *C.GObject {
	if this0 == nil {
		return nil
	}
	return (*C.GObject)(this0.C)
}

func (this0 *Object) GetStaticType() Type {
	return Type(C.g_object_get_type())
}

func ObjectGetType() Type {
	return (*Object)(nil).GetStaticType()
}
// blacklisted: Object.new (method)
// blacklisted: Object.bind_property (method)
// blacklisted: Object.bind_property_full (method)
// blacklisted: Object.compat_control (method)
// blacklisted: Object.interface_find_property (method)
// blacklisted: Object.interface_install_property (method)
// blacklisted: Object.interface_list_properties (method)
// blacklisted: Object.force_floating (method)
// blacklisted: Object.freeze_notify (method)
// blacklisted: Object.get_data (method)
// blacklisted: Object.get_property (method)
// blacklisted: Object.get_qdata (method)
// blacklisted: Object.is_floating (method)
// blacklisted: Object.notify (method)
// blacklisted: Object.notify_by_pspec (method)
// blacklisted: Object.ref (method)
// blacklisted: Object.ref_sink (method)
// blacklisted: Object.run_dispose (method)
// blacklisted: Object.set_data (method)
// blacklisted: Object.set_property (method)
// blacklisted: Object.steal_data (method)
// blacklisted: Object.steal_qdata (method)
// blacklisted: Object.thaw_notify (method)
// blacklisted: Object.unref (method)
// blacklisted: Object.watch_closure (method)
type ObjectConstructParam struct {
	pspec0 *C.GParamSpec
	Value *Value
}
func (this0 *ObjectConstructParam) Pspec() *ParamSpec {
	var pspec1 *ParamSpec
	pspec1 = (*ParamSpec)(ObjectWrap(unsafe.Pointer(this0.pspec0), true))
	return pspec1
}
// blacklisted: ObjectFinalizeFunc (callback)
// blacklisted: ObjectGetPropertyFunc (callback)
// blacklisted: ObjectSetPropertyFunc (callback)
const ParamMask = 255
const ParamReadwrite = 0
const ParamStaticStrings = 0
const ParamUserShift = 8
type ParamFlags C.uint32_t
const (
	ParamFlagsReadable ParamFlags = 1
	ParamFlagsWritable ParamFlags = 2
	ParamFlagsConstruct ParamFlags = 4
	ParamFlagsConstructOnly ParamFlags = 8
	ParamFlagsLaxValidation ParamFlags = 16
	ParamFlagsStaticName ParamFlags = 32
	ParamFlagsPrivate ParamFlags = 32
	ParamFlagsStaticNick ParamFlags = 64
	ParamFlagsStaticBlurb ParamFlags = 128
	ParamFlagsDeprecated ParamFlags = 2147483648
)
// blacklisted: ParamSpec (object)
// blacklisted: ParamSpecBoolean (object)
// blacklisted: ParamSpecBoxed (object)
// blacklisted: ParamSpecChar (object)
// blacklisted: ParamSpecDouble (object)
// blacklisted: ParamSpecEnum (object)
// blacklisted: ParamSpecFlags (object)
// blacklisted: ParamSpecFloat (object)
// blacklisted: ParamSpecGType (object)
// blacklisted: ParamSpecInt (object)
// blacklisted: ParamSpecInt64 (object)
// blacklisted: ParamSpecLong (object)
// blacklisted: ParamSpecObject (object)
// blacklisted: ParamSpecOverride (object)
// blacklisted: ParamSpecParam (object)
// blacklisted: ParamSpecPointer (object)
// blacklisted: ParamSpecPool (struct)
// blacklisted: ParamSpecString (object)
type ParamSpecTypeInfo struct {
	InstanceSize uint16
	NPreallocs uint16
	_ [4]byte
	InstanceInit unsafe.Pointer
	ValueType Type
	Finalize unsafe.Pointer
	ValueSetDefault unsafe.Pointer
	ValueValidate unsafe.Pointer
	ValuesCmp unsafe.Pointer
}
// blacklisted: ParamSpecUChar (object)
// blacklisted: ParamSpecUInt (object)
// blacklisted: ParamSpecUInt64 (object)
// blacklisted: ParamSpecULong (object)
// blacklisted: ParamSpecUnichar (object)
// blacklisted: ParamSpecValueArray (object)
// blacklisted: ParamSpecVariant (object)
type Parameter struct {
	name0 *C.char
	Value Value
}
func (this0 *Parameter) Name() string {
	var name1 string
	name1 = C.GoString(this0.name0)
	return name1
}
const SignalFlagsMask = 255
const SignalMatchMask = 63
// blacklisted: SignalAccumulator (callback)
// blacklisted: SignalEmissionHook (callback)
type SignalFlags C.uint32_t
const (
	SignalFlagsRunFirst SignalFlags = 1
	SignalFlagsRunLast SignalFlags = 2
	SignalFlagsRunCleanup SignalFlags = 4
	SignalFlagsNoRecurse SignalFlags = 8
	SignalFlagsDetailed SignalFlags = 16
	SignalFlagsAction SignalFlags = 32
	SignalFlagsNoHooks SignalFlags = 64
	SignalFlagsMustCollect SignalFlags = 128
)
type SignalInvocationHint struct {
	SignalID uint32
	Detail uint32
	RunType SignalFlags
}
type SignalMatchType C.uint32_t
const (
	SignalMatchTypeID SignalMatchType = 1
	SignalMatchTypeDetail SignalMatchType = 2
	SignalMatchTypeClosure SignalMatchType = 4
	SignalMatchTypeFunc SignalMatchType = 8
	SignalMatchTypeData SignalMatchType = 16
	SignalMatchTypeUnblocked SignalMatchType = 32
)
type SignalQuery struct {
	SignalID uint32
	_ [4]byte
	signal_name0 *C.char
	Itype Type
	SignalFlags SignalFlags
	_ [4]byte
	ReturnType Type
	NParams uint32
	_ [4]byte
	ParamTypes *Type
}
func (this0 *SignalQuery) SignalName() string {
	var signal_name1 string
	signal_name1 = C.GoString(this0.signal_name0)
	return signal_name1
}
const TypeFundamentalMax = 255
const TypeFundamentalShift = 2
const TypeReservedBseFirst = 32
const TypeReservedBseLast = 48
const TypeReservedGlibFirst = 22
const TypeReservedGlibLast = 31
const TypeReservedUserFirst = 49
// blacklisted: ToggleNotify (callback)
type TypeCValue struct {
	_data [8]byte
}
type TypeClass struct {
	GType Type
}
func (this0 *TypeClass) PeekParent() *TypeClass {
	var this1 *C.GTypeClass
	this1 = (*C.GTypeClass)(unsafe.Pointer(this0))
	ret1 := C.g_type_class_peek_parent(this1)
	var ret2 *TypeClass
	ret2 = (*TypeClass)(unsafe.Pointer(ret1))
	return ret2
}
func TypeClassAddPrivate(g_class0 unsafe.Pointer, private_size0 uint64) {
	var g_class1 unsafe.Pointer
	var private_size1 C.uint64_t
	g_class1 = unsafe.Pointer(g_class0)
	private_size1 = C.uint64_t(private_size0)
	C.g_type_class_add_private(g_class1, private_size1)
}
func TypeClassPeek(type0 Type) *TypeClass {
	var type1 C.GType
	type1 = C.GType(type0)
	ret1 := C.g_type_class_peek(type1)
	var ret2 *TypeClass
	ret2 = (*TypeClass)(unsafe.Pointer(ret1))
	return ret2
}
func TypeClassPeekStatic(type0 Type) *TypeClass {
	var type1 C.GType
	type1 = C.GType(type0)
	ret1 := C.g_type_class_peek_static(type1)
	var ret2 *TypeClass
	ret2 = (*TypeClass)(unsafe.Pointer(ret1))
	return ret2
}
// blacklisted: TypeClassCacheFunc (callback)
type TypeDebugFlags C.uint32_t
const (
	TypeDebugFlagsNone TypeDebugFlags = 0
	TypeDebugFlagsObjects TypeDebugFlags = 1
	TypeDebugFlagsSignals TypeDebugFlags = 2
	TypeDebugFlagsMask TypeDebugFlags = 3
)
type TypeFlags C.uint32_t
const (
	TypeFlagsAbstract TypeFlags = 16
	TypeFlagsValueAbstract TypeFlags = 32
)
type TypeFundamentalFlags C.uint32_t
const (
	TypeFundamentalFlagsClassed TypeFundamentalFlags = 1
	TypeFundamentalFlagsInstantiatable TypeFundamentalFlags = 2
	TypeFundamentalFlagsDerivable TypeFundamentalFlags = 4
	TypeFundamentalFlagsDeepDerivable TypeFundamentalFlags = 8
)
type TypeFundamentalInfo struct {
	TypeFlags TypeFundamentalFlags
}
type TypeInfo struct {
	ClassSize uint16
	_ [6]byte
	BaseInit unsafe.Pointer
	BaseFinalize unsafe.Pointer
	ClassInit unsafe.Pointer
	ClassFinalize unsafe.Pointer
	ClassData unsafe.Pointer
	InstanceSize uint16
	NPreallocs uint16
	_ [4]byte
	InstanceInit unsafe.Pointer
	ValueTable *TypeValueTable
}
type TypeInstance struct {
	GClass *TypeClass
}
type TypeInterface struct {
	GType Type
	GInstanceType Type
}
func (this0 *TypeInterface) PeekParent() *TypeInterface {
	var this1 *C.GTypeInterface
	this1 = (*C.GTypeInterface)(unsafe.Pointer(this0))
	ret1 := C.g_type_interface_peek_parent(this1)
	var ret2 *TypeInterface
	ret2 = (*TypeInterface)(unsafe.Pointer(ret1))
	return ret2
}
func TypeInterfaceAddPrerequisite(interface_type0 Type, prerequisite_type0 Type) {
	var interface_type1 C.GType
	var prerequisite_type1 C.GType
	interface_type1 = C.GType(interface_type0)
	prerequisite_type1 = C.GType(prerequisite_type0)
	C.g_type_interface_add_prerequisite(interface_type1, prerequisite_type1)
}
func TypeInterfaceGetPlugin(instance_type0 Type, interface_type0 Type) *TypePlugin {
	var instance_type1 C.GType
	var interface_type1 C.GType
	instance_type1 = C.GType(instance_type0)
	interface_type1 = C.GType(interface_type0)
	ret1 := C.g_type_interface_get_plugin(instance_type1, interface_type1)
	var ret2 *TypePlugin
	ret2 = (*TypePlugin)(ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func TypeInterfacePeek(instance_class0 *TypeClass, iface_type0 Type) *TypeInterface {
	var instance_class1 *C.GTypeClass
	var iface_type1 C.GType
	instance_class1 = (*C.GTypeClass)(unsafe.Pointer(instance_class0))
	iface_type1 = C.GType(iface_type0)
	ret1 := C.g_type_interface_peek(instance_class1, iface_type1)
	var ret2 *TypeInterface
	ret2 = (*TypeInterface)(unsafe.Pointer(ret1))
	return ret2
}
func TypeInterfacePrerequisites(interface_type0 Type) (int, []Type) {
	var interface_type1 C.GType
	var n_prerequisites1 C.uint32_t
	interface_type1 = C.GType(interface_type0)
	ret1 := C.g_type_interface_prerequisites(interface_type1, &n_prerequisites1)
	var n_prerequisites2 int
	var ret2 []Type
	n_prerequisites2 = int(n_prerequisites1)
	ret2 = make([]Type, n_prerequisites1)
	for i := range ret2 {
		ret2[i] = Type((*(*[999999]C.GType)(unsafe.Pointer(ret1)))[i])
	}
	return n_prerequisites2, ret2
}
// blacklisted: TypeInterfaceCheckFunc (callback)
// blacklisted: TypeModule (object)
type TypePluginLike interface {
	ImplementsGTypePlugin() *C.GTypePlugin
}

type TypePlugin struct {
	Object
	TypePluginImpl
}

type TypePluginImpl struct {}

func ToTypePlugin(objlike ObjectLike) *TypePlugin {
	t := (*TypePluginImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*TypePlugin)(obj)
	}
	panic("cannot cast to TypePlugin")
}

func (this0 *TypePluginImpl) ImplementsGTypePlugin() *C.GTypePlugin {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GTypePlugin)((*Object)(unsafe.Pointer(obj)).C)
}

func (this0 *TypePluginImpl) GetStaticType() Type {
	return Type(C.g_type_plugin_get_type())
}

func TypePluginGetType() Type {
	return (*TypePluginImpl)(nil).GetStaticType()
}
func (this0 *TypePluginImpl) CompleteInterfaceInfo(instance_type0 Type, interface_type0 Type, info0 *InterfaceInfo) {
	var this1 *C.GTypePlugin
	var instance_type1 C.GType
	var interface_type1 C.GType
	var info1 *C.GInterfaceInfo
	if this0 != nil {
		this1 = this0.ImplementsGTypePlugin()}
	instance_type1 = C.GType(instance_type0)
	interface_type1 = C.GType(interface_type0)
	info1 = (*C.GInterfaceInfo)(unsafe.Pointer(info0))
	C.g_type_plugin_complete_interface_info(this1, instance_type1, interface_type1, info1)
}
func (this0 *TypePluginImpl) CompleteTypeInfo(g_type0 Type, info0 *TypeInfo, value_table0 *TypeValueTable) {
	var this1 *C.GTypePlugin
	var g_type1 C.GType
	var info1 *C.GTypeInfo
	var value_table1 *C.GTypeValueTable
	if this0 != nil {
		this1 = this0.ImplementsGTypePlugin()}
	g_type1 = C.GType(g_type0)
	info1 = (*C.GTypeInfo)(unsafe.Pointer(info0))
	value_table1 = (*C.GTypeValueTable)(unsafe.Pointer(value_table0))
	C.g_type_plugin_complete_type_info(this1, g_type1, info1, value_table1)
}
func (this0 *TypePluginImpl) Unuse() {
	var this1 *C.GTypePlugin
	if this0 != nil {
		this1 = this0.ImplementsGTypePlugin()}
	C.g_type_plugin_unuse(this1)
}
func (this0 *TypePluginImpl) Use() {
	var this1 *C.GTypePlugin
	if this0 != nil {
		this1 = this0.ImplementsGTypePlugin()}
	C.g_type_plugin_use(this1)
}
type TypePluginClass struct {
	BaseIface TypeInterface
	UsePlugin unsafe.Pointer
	UnusePlugin unsafe.Pointer
	CompleteTypeInfo unsafe.Pointer
	CompleteInterfaceInfo unsafe.Pointer
}
// blacklisted: TypePluginCompleteInterfaceInfo (callback)
// blacklisted: TypePluginCompleteTypeInfo (callback)
// blacklisted: TypePluginUnuse (callback)
// blacklisted: TypePluginUse (callback)
type TypeQuery struct {
	Type Type
	type_name0 *C.char
	ClassSize uint32
	InstanceSize uint32
}
func (this0 *TypeQuery) TypeName() string {
	var type_name1 string
	type_name1 = C.GoString(this0.type_name0)
	return type_name1
}
type TypeValueTable struct {
	ValueInit unsafe.Pointer
	ValueFree unsafe.Pointer
	ValueCopy unsafe.Pointer
	ValuePeekPointer unsafe.Pointer
	collect_format0 *C.char
	CollectValue unsafe.Pointer
	lcopy_format0 *C.char
	LcopyValue unsafe.Pointer
}
func (this0 *TypeValueTable) CollectFormat() string {
	var collect_format1 string
	collect_format1 = C.GoString(this0.collect_format0)
	return collect_format1
}
func (this0 *TypeValueTable) LcopyFormat() string {
	var lcopy_format1 string
	lcopy_format1 = C.GoString(this0.lcopy_format0)
	return lcopy_format1
}
const ValueCollectFormatMaxLength = 8
const ValueNocopyContents = 134217728
type Value struct {
	GType Type
	Data [2]_Value__data__union
}
type ValueArray struct {
	NValues uint32
	_ [4]byte
	Values *Value
	NPrealloced uint32
	_ [4]byte
}
func NewValueArray(n_prealloced0 int) *ValueArray {
	var n_prealloced1 C.uint32_t
	n_prealloced1 = C.uint32_t(n_prealloced0)
	ret1 := C.g_value_array_new(n_prealloced1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Append(value0 *Value) *ValueArray {
	var this1 *C.GValueArray
	var value1 *C.GValue
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	ret1 := C.g_value_array_append(this1, value1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Copy() *ValueArray {
	var this1 *C.GValueArray
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	ret1 := C.g_value_array_copy(this1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Free() {
	var this1 *C.GValueArray
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	C.g_value_array_free(this1)
}
func (this0 *ValueArray) GetNth(index_0 int) *Value {
	var this1 *C.GValueArray
	var index_1 C.uint32_t
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	index_1 = C.uint32_t(index_0)
	ret1 := C.g_value_array_get_nth(this1, index_1)
	var ret2 *Value
	ret2 = (*Value)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Insert(index_0 int, value0 *Value) *ValueArray {
	var this1 *C.GValueArray
	var index_1 C.uint32_t
	var value1 *C.GValue
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	index_1 = C.uint32_t(index_0)
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	ret1 := C.g_value_array_insert(this1, index_1, value1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Prepend(value0 *Value) *ValueArray {
	var this1 *C.GValueArray
	var value1 *C.GValue
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	ret1 := C.g_value_array_prepend(this1, value1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *ValueArray) Remove(index_0 int) *ValueArray {
	var this1 *C.GValueArray
	var index_1 C.uint32_t
	this1 = (*C.GValueArray)(unsafe.Pointer(this0))
	index_1 = C.uint32_t(index_0)
	ret1 := C.g_value_array_remove(this1, index_1)
	var ret2 *ValueArray
	ret2 = (*ValueArray)(unsafe.Pointer(ret1))
	return ret2
}
// blacklisted: ValueTransform (callback)
// blacklisted: WeakNotify (callback)
type _Value__data__union struct {
	_data [8]byte
}
// blacklisted: boxed_free (function)
// blacklisted: cclosure_marshal_BOOLEAN__BOXED_BOXED (function)
// blacklisted: cclosure_marshal_BOOLEAN__FLAGS (function)
// blacklisted: cclosure_marshal_STRING__OBJECT_POINTER (function)
// blacklisted: cclosure_marshal_VOID__BOOLEAN (function)
// blacklisted: cclosure_marshal_VOID__BOXED (function)
// blacklisted: cclosure_marshal_VOID__CHAR (function)
// blacklisted: cclosure_marshal_VOID__DOUBLE (function)
// blacklisted: cclosure_marshal_VOID__ENUM (function)
// blacklisted: cclosure_marshal_VOID__FLAGS (function)
// blacklisted: cclosure_marshal_VOID__FLOAT (function)
// blacklisted: cclosure_marshal_VOID__INT (function)
// blacklisted: cclosure_marshal_VOID__LONG (function)
// blacklisted: cclosure_marshal_VOID__OBJECT (function)
// blacklisted: cclosure_marshal_VOID__PARAM (function)
// blacklisted: cclosure_marshal_VOID__POINTER (function)
// blacklisted: cclosure_marshal_VOID__STRING (function)
// blacklisted: cclosure_marshal_VOID__UCHAR (function)
// blacklisted: cclosure_marshal_VOID__UINT (function)
// blacklisted: cclosure_marshal_VOID__UINT_POINTER (function)
// blacklisted: cclosure_marshal_VOID__ULONG (function)
// blacklisted: cclosure_marshal_VOID__VARIANT (function)
// blacklisted: cclosure_marshal_VOID__VOID (function)
// blacklisted: cclosure_marshal_generic (function)
// blacklisted: enum_complete_type_info (function)
// blacklisted: enum_register_static (function)
// blacklisted: flags_complete_type_info (function)
// blacklisted: flags_register_static (function)
// blacklisted: gtype_get_type (function)
// blacklisted: param_spec_pool_new (function)
// blacklisted: param_type_register_static (function)
// blacklisted: param_value_convert (function)
// blacklisted: param_value_defaults (function)
// blacklisted: param_value_set_default (function)
// blacklisted: param_value_validate (function)
// blacklisted: param_values_cmp (function)
// blacklisted: pointer_type_register_static (function)
// blacklisted: signal_accumulator_first_wins (function)
// blacklisted: signal_accumulator_true_handled (function)
// blacklisted: signal_add_emission_hook (function)
// blacklisted: signal_chain_from_overridden (function)
// blacklisted: signal_connect_closure (function)
// blacklisted: signal_connect_closure_by_id (function)
// blacklisted: signal_emitv (function)
// blacklisted: signal_get_invocation_hint (function)
// blacklisted: signal_handler_block (function)
// blacklisted: signal_handler_disconnect (function)
// blacklisted: signal_handler_find (function)
// blacklisted: signal_handler_is_connected (function)
// blacklisted: signal_handler_unblock (function)
// blacklisted: signal_handlers_block_matched (function)
// blacklisted: signal_handlers_destroy (function)
// blacklisted: signal_handlers_disconnect_matched (function)
// blacklisted: signal_handlers_unblock_matched (function)
// blacklisted: signal_has_handler_pending (function)
// blacklisted: signal_list_ids (function)
// blacklisted: signal_lookup (function)
// blacklisted: signal_name (function)
// blacklisted: signal_override_class_closure (function)
// blacklisted: signal_parse_name (function)
// blacklisted: signal_query (function)
// blacklisted: signal_remove_emission_hook (function)
// blacklisted: signal_stop_emission (function)
// blacklisted: signal_stop_emission_by_name (function)
// blacklisted: signal_type_cclosure_new (function)
// blacklisted: source_set_closure (function)
// blacklisted: source_set_dummy_callback (function)
// blacklisted: strdup_value_contents (function)
// blacklisted: type_add_class_private (function)
// blacklisted: type_add_interface_dynamic (function)
// blacklisted: type_add_interface_static (function)
// blacklisted: type_check_class_is_a (function)
// blacklisted: type_check_instance (function)
// blacklisted: type_check_instance_is_a (function)
// blacklisted: type_check_is_value_type (function)
// blacklisted: type_check_value (function)
// blacklisted: type_check_value_holds (function)
// blacklisted: type_children (function)
// blacklisted: type_class_add_private (function)
// blacklisted: type_class_peek (function)
// blacklisted: type_class_peek_static (function)
// blacklisted: type_class_ref (function)
// blacklisted: type_default_interface_peek (function)
// blacklisted: type_default_interface_ref (function)
// blacklisted: type_default_interface_unref (function)
// blacklisted: type_depth (function)
// blacklisted: type_free_instance (function)
// blacklisted: type_from_name (function)
// blacklisted: type_fundamental (function)
// blacklisted: type_fundamental_next (function)
// blacklisted: type_get_plugin (function)
// blacklisted: type_get_qdata (function)
// blacklisted: type_init (function)
// blacklisted: type_init_with_debug_flags (function)
// blacklisted: type_interface_add_prerequisite (function)
// blacklisted: type_interface_get_plugin (function)
// blacklisted: type_interface_peek (function)
// blacklisted: type_interface_prerequisites (function)
// blacklisted: type_interfaces (function)
// blacklisted: type_is_a (function)
// blacklisted: type_name (function)
// blacklisted: type_name_from_class (function)
// blacklisted: type_name_from_instance (function)
// blacklisted: type_next_base (function)
// blacklisted: type_parent (function)
// blacklisted: type_qname (function)
// blacklisted: type_query (function)
// blacklisted: type_register_dynamic (function)
// blacklisted: type_register_fundamental (function)
// blacklisted: type_register_static (function)
// blacklisted: type_set_qdata (function)
// blacklisted: type_test_flags (function)
// blacklisted: value_get_gtype (function)
// blacklisted: value_type_compatible (function)
// blacklisted: value_type_transformable (function)


//--------------------------------------------------------------
// Holder
//--------------------------------------------------------------
// holy crap, what am I doing here...

type holder_key [2]unsafe.Pointer
type holder_type map[holder_key]int

var Holder = holder_type(make(map[holder_key]int))

func (this holder_type) Grab(x interface{}) {
	if x == nil {
		return
	}

	key := *(*holder_key)(unsafe.Pointer(&x))
	count := this[key]
	this[key] = count + 1
}

func (this holder_type) Release(x interface{}) {
	if x == nil {
		return
	}

	key := *(*holder_key)(unsafe.Pointer(&x))
	count := this[key]
	if count <= 1 {
		delete(this, key)
	} else {
		this[key] = count - 1
	}
}

//--------------------------------------------------------------
// FinalizerQueue
//--------------------------------------------------------------

type finalizer_item struct {
	ptr unsafe.Pointer
	finalizer func(unsafe.Pointer)
}

type fqueue_type struct {
	sync.Mutex
	queue []finalizer_item
	exec_queue []finalizer_item
	tid uint32
}

var FQueue fqueue_type

func (this *fqueue_type) Start(interval int) {
	this.Lock()
	this.queue = make([]finalizer_item, 0, 50)
	this.exec_queue = make([]finalizer_item, 50)
	this.tid = uint32(C._g_timeout_add_fqueue(C.uint32_t(interval)))
	this.Unlock()
}

func (this *fqueue_type) Stop() {
	this.Lock()
	// TODO: we'll discard few items here at Stop, is it ok?
	this.queue = nil
	C.g_source_remove(C.uint32_t(this.tid))
	this.Unlock()
}

// returns true if the item was enqueued, thread safe
func (this *fqueue_type) Push(ptr unsafe.Pointer, finalizer func(unsafe.Pointer)) bool {
	this.Lock()
	if this.queue != nil {
		this.queue = append(this.queue, finalizer_item{ptr, finalizer})
		this.Unlock()
		return true
	}
	this.Unlock()
	return false
}

// exec is only thread safe if executed by a single thread
func (this *fqueue_type) exec() {
	// exec_queue is used for not holding the lock a lot
	this.Lock()
	// common case
	if len(this.queue) == 0 {
		this.Unlock()
		return
	}

	// non-empty queue, copy everything to exec_queue
	if len(this.queue) > len(this.exec_queue) {
		this.exec_queue = make([]finalizer_item, len(this.queue))
	}
	nitems := copy(this.exec_queue, this.queue)
	this.queue = this.queue[:0]
	this.Unlock()

	// then do our work
	for i := 0; i < nitems; i++ {
		this.exec_queue[i].finalizer(this.exec_queue[i].ptr)
		this.exec_queue[i] = finalizer_item{}
	}
}

//export fqueue_dispatcher
func fqueue_dispatcher(unused unsafe.Pointer) int32 {
	FQueue.exec()
	return 1
}

//--------------------------------------------------------------
// NilString
//--------------------------------------------------------------

// its value will stay the same forever, use the value directly if you like
const NilString = "\x00"

//--------------------------------------------------------------
// Quark
//
// TODO: probably it's a temporary place for this, quarks are
// from glib
//--------------------------------------------------------------

type Quark uint32

func NewQuarkFromString(s string) Quark {
	cs := C.CString(s)
	quark := C.g_quark_from_string(cs)
	C.free(unsafe.Pointer(cs))
	return Quark(quark)
}

// we use this one to store Go's representation of the GObject
// as user data in that GObject once it was allocated. For the
// sake of avoiding allocations.
var go_repr Quark

func init() {
	go_repr = NewQuarkFromString("go-representation")
}

//--------------------------------------------------------------
// ParamSpec utils
//--------------------------------------------------------------

// Let's implement these manually (not Object based and small amount of things
// to implement).

// First some utils
func param_spec_finalizer(pspec *ParamSpec) {
	if FQueue.Push(unsafe.Pointer(pspec), param_spec_finalizer2) {
		return
	}
	C.g_param_spec_unref((*C.GParamSpec)(pspec.C))
}

func param_spec_finalizer2(pspec_un unsafe.Pointer) {
	pspec := (*ParamSpec)(pspec_un)
	C.g_param_spec_unref((*C.GParamSpec)(pspec.C))
}

func set_param_spec_finalizer(pspec *ParamSpec) {
	runtime.SetFinalizer(pspec, param_spec_finalizer)
}

func ParamSpecGrabIfType(c unsafe.Pointer, t Type) unsafe.Pointer {
	if c == nil {
		return nil
	}
	obj := &ParamSpec{c}
	if obj.GetType().IsA(t) {
		C.g_param_spec_ref_sink((*C.GParamSpec)(obj.C))
		set_param_spec_finalizer(obj)
		return unsafe.Pointer(obj)
	}
	return nil
}

func ParamSpecWrap(c unsafe.Pointer, grab bool) unsafe.Pointer {
	if c == nil {
		return nil
	}
	obj := &ParamSpec{c}
	if grab {
		C.g_param_spec_ref_sink((*C.GParamSpec)(obj.C))
	}
	set_param_spec_finalizer(obj)
	return unsafe.Pointer(obj)
}

//--------------------------------------------------------------
// ParamSpec
//--------------------------------------------------------------

type ParamSpecLike interface {
	InheritedFromGParamSpec() *C.GParamSpec
}

type ParamSpec struct {
	C unsafe.Pointer
}

func ToParamSpec(pspeclike ParamSpecLike) *ParamSpec {
	t := (*ParamSpec)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpec()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpec)(obj)
	}
	panic("cannot cast to ParamSpec")
}

func (this *ParamSpec) InheritedFromGParamSpec() *C.GParamSpec {
	return (*C.GParamSpec)(this.C)
}

func (this *ParamSpec) GetStaticType() Type {
	return Type(C._g_type_param())
}

func (this *ParamSpec) GetType() Type {
	return Type(C._g_param_spec_type(this.InheritedFromGParamSpec()))
}

func (this *ParamSpec) GetValueType() Type {
	return Type(C._g_param_spec_value_type(this.InheritedFromGParamSpec()))
}

//--------------------------------------------------------------
// ParamSpecBoolean
//--------------------------------------------------------------

type ParamSpecBooleanLike interface {
	InheritedFromGParamSpecBoolean() *C.GParamSpecBoolean
}

type ParamSpecBoolean struct {
	ParamSpec
}

func ToParamSpecBoolean(pspeclike ParamSpecBooleanLike) *ParamSpecBoolean {
	t := (*ParamSpecBoolean)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecBoolean()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecBoolean)(obj)
	}
	panic("cannot cast to ParamSpecBoolean")
}

func (this *ParamSpecBoolean) InheritedFromGParamSpecBoolean() *C.GParamSpecBoolean {
	return (*C.GParamSpecBoolean)(this.C)
}

func (this *ParamSpecBoolean) GetStaticType() Type {
	return Type(C._g_type_param_boolean())
}

//--------------------------------------------------------------
// ParamSpecBoxed
//--------------------------------------------------------------

type ParamSpecBoxedLike interface {
	InheritedFromGParamSpecBoxed() *C.GParamSpecBoxed
}

type ParamSpecBoxed struct {
	ParamSpec
}

func ToParamSpecBoxed(pspeclike ParamSpecBoxedLike) *ParamSpecBoxed {
	t := (*ParamSpecBoxed)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecBoxed()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecBoxed)(obj)
	}
	panic("cannot cast to ParamSpecBoxed")
}

func (this *ParamSpecBoxed) InheritedFromGParamSpecBoxed() *C.GParamSpecBoxed {
	return (*C.GParamSpecBoxed)(this.C)
}

func (this *ParamSpecBoxed) GetStaticType() Type {
	return Type(C._g_type_param_boxed())
}

//--------------------------------------------------------------
// ParamSpecChar
//--------------------------------------------------------------

type ParamSpecCharLike interface {
	InheritedFromGParamSpecChar() *C.GParamSpecChar
}

type ParamSpecChar struct {
	ParamSpec
}

func ToParamSpecChar(pspeclike ParamSpecCharLike) *ParamSpecChar {
	t := (*ParamSpecChar)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecChar()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecChar)(obj)
	}
	panic("cannot cast to ParamSpecChar")
}

func (this *ParamSpecChar) InheritedFromGParamSpecChar() *C.GParamSpecChar {
	return (*C.GParamSpecChar)(this.C)
}

func (this *ParamSpecChar) GetStaticType() Type {
	return Type(C._g_type_param_char())
}

//--------------------------------------------------------------
// ParamSpecDouble
//--------------------------------------------------------------

type ParamSpecDoubleLike interface {
	InheritedFromGParamSpecDouble() *C.GParamSpecDouble
}

type ParamSpecDouble struct {
	ParamSpec
}

func ToParamSpecDouble(pspeclike ParamSpecDoubleLike) *ParamSpecDouble {
	t := (*ParamSpecDouble)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecDouble()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecDouble)(obj)
	}
	panic("cannot cast to ParamSpecDouble")
}

func (this *ParamSpecDouble) InheritedFromGParamSpecDouble() *C.GParamSpecDouble {
	return (*C.GParamSpecDouble)(this.C)
}

func (this *ParamSpecDouble) GetStaticType() Type {
	return Type(C._g_type_param_double())
}

//--------------------------------------------------------------
// ParamSpecEnum
//--------------------------------------------------------------

type ParamSpecEnumLike interface {
	InheritedFromGParamSpecEnum() *C.GParamSpecEnum
}

type ParamSpecEnum struct {
	ParamSpec
}

func ToParamSpecEnum(pspeclike ParamSpecEnumLike) *ParamSpecEnum {
	t := (*ParamSpecEnum)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecEnum()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecEnum)(obj)
	}
	panic("cannot cast to ParamSpecEnum")
}

func (this *ParamSpecEnum) InheritedFromGParamSpecEnum() *C.GParamSpecEnum {
	return (*C.GParamSpecEnum)(this.C)
}

func (this *ParamSpecEnum) GetStaticType() Type {
	return Type(C._g_type_param_enum())
}

//--------------------------------------------------------------
// ParamSpecFlags
//--------------------------------------------------------------

type ParamSpecFlagsLike interface {
	InheritedFromGParamSpecFlags() *C.GParamSpecFlags
}

type ParamSpecFlags struct {
	ParamSpec
}

func ToParamSpecFlags(pspeclike ParamSpecFlagsLike) *ParamSpecFlags {
	t := (*ParamSpecFlags)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecFlags()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecFlags)(obj)
	}
	panic("cannot cast to ParamSpecFlags")
}

func (this *ParamSpecFlags) InheritedFromGParamSpecFlags() *C.GParamSpecFlags {
	return (*C.GParamSpecFlags)(this.C)
}

func (this *ParamSpecFlags) GetStaticType() Type {
	return Type(C._g_type_param_flags())
}

//--------------------------------------------------------------
// ParamSpecFloat
//--------------------------------------------------------------

type ParamSpecFloatLike interface {
	InheritedFromGParamSpecFloat() *C.GParamSpecFloat
}

type ParamSpecFloat struct {
	ParamSpec
}

func ToParamSpecFloat(pspeclike ParamSpecFloatLike) *ParamSpecFloat {
	t := (*ParamSpecFloat)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecFloat()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecFloat)(obj)
	}
	panic("cannot cast to ParamSpecFloat")
}

func (this *ParamSpecFloat) InheritedFromGParamSpecFloat() *C.GParamSpecFloat {
	return (*C.GParamSpecFloat)(this.C)
}

func (this *ParamSpecFloat) GetStaticType() Type {
	return Type(C._g_type_param_float())
}

//--------------------------------------------------------------
// ParamSpecGType
//--------------------------------------------------------------

type ParamSpecGTypeLike interface {
	InheritedFromGParamSpecGType() *C.GParamSpecGType
}

type ParamSpecGType struct {
	ParamSpec
}

func ToParamSpecGType(pspeclike ParamSpecGTypeLike) *ParamSpecGType {
	t := (*ParamSpecGType)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecGType()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecGType)(obj)
	}
	panic("cannot cast to ParamSpecGType")
}

func (this *ParamSpecGType) InheritedFromGParamSpecGType() *C.GParamSpecGType {
	return (*C.GParamSpecGType)(this.C)
}

func (this *ParamSpecGType) GetStaticType() Type {
	return Type(C._g_type_param_gtype())
}

//--------------------------------------------------------------
// ParamSpecInt
//--------------------------------------------------------------

type ParamSpecIntLike interface {
	InheritedFromGParamSpecInt() *C.GParamSpecInt
}

type ParamSpecInt struct {
	ParamSpec
}

func ToParamSpecInt(pspeclike ParamSpecIntLike) *ParamSpecInt {
	t := (*ParamSpecInt)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecInt()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecInt)(obj)
	}
	panic("cannot cast to ParamSpecInt")
}

func (this *ParamSpecInt) InheritedFromGParamSpecInt() *C.GParamSpecInt {
	return (*C.GParamSpecInt)(this.C)
}

func (this *ParamSpecInt) GetStaticType() Type {
	return Type(C._g_type_param_int())
}

//--------------------------------------------------------------
// ParamSpecInt64
//--------------------------------------------------------------

type ParamSpecInt64Like interface {
	InheritedFromGParamSpecInt64() *C.GParamSpecInt64
}

type ParamSpecInt64 struct {
	ParamSpec
}

func ToParamSpecInt64(pspeclike ParamSpecInt64Like) *ParamSpecInt64 {
	t := (*ParamSpecInt64)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecInt64()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecInt64)(obj)
	}
	panic("cannot cast to ParamSpecInt64")
}

func (this *ParamSpecInt64) InheritedFromGParamSpecInt64() *C.GParamSpecInt64 {
	return (*C.GParamSpecInt64)(this.C)
}

func (this *ParamSpecInt64) GetStaticType() Type {
	return Type(C._g_type_param_int64())
}

//--------------------------------------------------------------
// ParamSpecLong
//--------------------------------------------------------------

type ParamSpecLongLike interface {
	InheritedFromGParamSpecLong() *C.GParamSpecLong
}

type ParamSpecLong struct {
	ParamSpec
}

func ToParamSpecLong(pspeclike ParamSpecLongLike) *ParamSpecLong {
	t := (*ParamSpecLong)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecLong()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecLong)(obj)
	}
	panic("cannot cast to ParamSpecLong")
}

func (this *ParamSpecLong) InheritedFromGParamSpecLong() *C.GParamSpecLong {
	return (*C.GParamSpecLong)(this.C)
}

func (this *ParamSpecLong) GetStaticType() Type {
	return Type(C._g_type_param_long())
}

//--------------------------------------------------------------
// ParamSpecObject
//--------------------------------------------------------------

type ParamSpecObjectLike interface {
	InheritedFromGParamSpecObject() *C.GParamSpecObject
}

type ParamSpecObject struct {
	ParamSpec
}

func ToParamSpecObject(pspeclike ParamSpecObjectLike) *ParamSpecObject {
	t := (*ParamSpecObject)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecObject()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecObject)(obj)
	}
	panic("cannot cast to ParamSpecObject")
}

func (this *ParamSpecObject) InheritedFromGParamSpecObject() *C.GParamSpecObject {
	return (*C.GParamSpecObject)(this.C)
}

func (this *ParamSpecObject) GetStaticType() Type {
	return Type(C._g_type_param_object())
}

//--------------------------------------------------------------
// ParamSpecOverride
//--------------------------------------------------------------

type ParamSpecOverrideLike interface {
	InheritedFromGParamSpecOverride() *C.GParamSpecOverride
}

type ParamSpecOverride struct {
	ParamSpec
}

func ToParamSpecOverride(pspeclike ParamSpecOverrideLike) *ParamSpecOverride {
	t := (*ParamSpecOverride)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecOverride()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecOverride)(obj)
	}
	panic("cannot cast to ParamSpecOverride")
}

func (this *ParamSpecOverride) InheritedFromGParamSpecOverride() *C.GParamSpecOverride {
	return (*C.GParamSpecOverride)(this.C)
}

func (this *ParamSpecOverride) GetStaticType() Type {
	return Type(C._g_type_param_override())
}

//--------------------------------------------------------------
// ParamSpecParam
//--------------------------------------------------------------

type ParamSpecParamLike interface {
	InheritedFromGParamSpecParam() *C.GParamSpecParam
}

type ParamSpecParam struct {
	ParamSpec
}

func ToParamSpecParam(pspeclike ParamSpecParamLike) *ParamSpecParam {
	t := (*ParamSpecParam)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecParam()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecParam)(obj)
	}
	panic("cannot cast to ParamSpecParam")
}

func (this *ParamSpecParam) InheritedFromGParamSpecParam() *C.GParamSpecParam {
	return (*C.GParamSpecParam)(this.C)
}

func (this *ParamSpecParam) GetStaticType() Type {
	return Type(C._g_type_param_param())
}

//--------------------------------------------------------------
// ParamSpecPointer
//--------------------------------------------------------------

type ParamSpecPointerLike interface {
	InheritedFromGParamSpecPointer() *C.GParamSpecPointer
}

type ParamSpecPointer struct {
	ParamSpec
}

func ToParamSpecPointer(pspeclike ParamSpecPointerLike) *ParamSpecPointer {
	t := (*ParamSpecPointer)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecPointer()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecPointer)(obj)
	}
	panic("cannot cast to ParamSpecPointer")
}

func (this *ParamSpecPointer) InheritedFromGParamSpecPointer() *C.GParamSpecPointer {
	return (*C.GParamSpecPointer)(this.C)
}

func (this *ParamSpecPointer) GetStaticType() Type {
	return Type(C._g_type_param_pointer())
}

//--------------------------------------------------------------
// ParamSpecString
//--------------------------------------------------------------

type ParamSpecStringLike interface {
	InheritedFromGParamSpecString() *C.GParamSpecString
}

type ParamSpecString struct {
	ParamSpec
}

func ToParamSpecString(pspeclike ParamSpecStringLike) *ParamSpecString {
	t := (*ParamSpecString)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecString()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecString)(obj)
	}
	panic("cannot cast to ParamSpecString")
}

func (this *ParamSpecString) InheritedFromGParamSpecString() *C.GParamSpecString {
	return (*C.GParamSpecString)(this.C)
}

func (this *ParamSpecString) GetStaticType() Type {
	return Type(C._g_type_param_string())
}

//--------------------------------------------------------------
// ParamSpecUChar
//--------------------------------------------------------------

type ParamSpecUCharLike interface {
	InheritedFromGParamSpecUChar() *C.GParamSpecUChar
}

type ParamSpecUChar struct {
	ParamSpec
}

func ToParamSpecUChar(pspeclike ParamSpecUCharLike) *ParamSpecUChar {
	t := (*ParamSpecUChar)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecUChar()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecUChar)(obj)
	}
	panic("cannot cast to ParamSpecUChar")
}

func (this *ParamSpecUChar) InheritedFromGParamSpecUChar() *C.GParamSpecUChar {
	return (*C.GParamSpecUChar)(this.C)
}

func (this *ParamSpecUChar) GetStaticType() Type {
	return Type(C._g_type_param_uchar())
}

//--------------------------------------------------------------
// ParamSpecUInt
//--------------------------------------------------------------

type ParamSpecUIntLike interface {
	InheritedFromGParamSpecUInt() *C.GParamSpecUInt
}

type ParamSpecUInt struct {
	ParamSpec
}

func ToParamSpecUInt(pspeclike ParamSpecUIntLike) *ParamSpecUInt {
	t := (*ParamSpecUInt)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecUInt()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecUInt)(obj)
	}
	panic("cannot cast to ParamSpecUInt")
}

func (this *ParamSpecUInt) InheritedFromGParamSpecUInt() *C.GParamSpecUInt {
	return (*C.GParamSpecUInt)(this.C)
}

func (this *ParamSpecUInt) GetStaticType() Type {
	return Type(C._g_type_param_uint())
}

//--------------------------------------------------------------
// ParamSpecUInt64
//--------------------------------------------------------------

type ParamSpecUInt64Like interface {
	InheritedFromGParamSpecUInt64() *C.GParamSpecUInt64
}

type ParamSpecUInt64 struct {
	ParamSpec
}

func ToParamSpecUInt64(pspeclike ParamSpecUInt64Like) *ParamSpecUInt64 {
	t := (*ParamSpecUInt64)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecUInt64()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecUInt64)(obj)
	}
	panic("cannot cast to ParamSpecUInt64")
}

func (this *ParamSpecUInt64) InheritedFromGParamSpecUInt64() *C.GParamSpecUInt64 {
	return (*C.GParamSpecUInt64)(this.C)
}

func (this *ParamSpecUInt64) GetStaticType() Type {
	return Type(C._g_type_param_uint64())
}

//--------------------------------------------------------------
// ParamSpecULong
//--------------------------------------------------------------

type ParamSpecULongLike interface {
	InheritedFromGParamSpecULong() *C.GParamSpecULong
}

type ParamSpecULong struct {
	ParamSpec
}

func ToParamSpecULong(pspeclike ParamSpecULongLike) *ParamSpecULong {
	t := (*ParamSpecULong)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecULong()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecULong)(obj)
	}
	panic("cannot cast to ParamSpecULong")
}

func (this *ParamSpecULong) InheritedFromGParamSpecULong() *C.GParamSpecULong {
	return (*C.GParamSpecULong)(this.C)
}

func (this *ParamSpecULong) GetStaticType() Type {
	return Type(C._g_type_param_ulong())
}

//--------------------------------------------------------------
// ParamSpecUnichar
//--------------------------------------------------------------

type ParamSpecUnicharLike interface {
	InheritedFromGParamSpecUnichar() *C.GParamSpecUnichar
}

type ParamSpecUnichar struct {
	ParamSpec
}

func ToParamSpecUnichar(pspeclike ParamSpecUnicharLike) *ParamSpecUnichar {
	t := (*ParamSpecUnichar)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecUnichar()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecUnichar)(obj)
	}
	panic("cannot cast to ParamSpecUnichar")
}

func (this *ParamSpecUnichar) InheritedFromGParamSpecUnichar() *C.GParamSpecUnichar {
	return (*C.GParamSpecUnichar)(this.C)
}

func (this *ParamSpecUnichar) GetStaticType() Type {
	return Type(C._g_type_param_unichar())
}

//--------------------------------------------------------------
// ParamSpecValueArray
//--------------------------------------------------------------

type ParamSpecValueArrayLike interface {
	InheritedFromGParamSpecValueArray() *C.GParamSpecValueArray
}

type ParamSpecValueArray struct {
	ParamSpec
}

func ToParamSpecValueArray(pspeclike ParamSpecValueArrayLike) *ParamSpecValueArray {
	t := (*ParamSpecValueArray)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecValueArray()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecValueArray)(obj)
	}
	panic("cannot cast to ParamSpecValueArray")
}

func (this *ParamSpecValueArray) InheritedFromGParamSpecValueArray() *C.GParamSpecValueArray {
	return (*C.GParamSpecValueArray)(this.C)
}

func (this *ParamSpecValueArray) GetStaticType() Type {
	return Type(C._g_type_param_value_array())
}

//--------------------------------------------------------------
// ParamSpecVariant
//--------------------------------------------------------------

type ParamSpecVariantLike interface {
	InheritedFromGParamSpecVariant() *C.GParamSpecVariant
}

type ParamSpecVariant struct {
	ParamSpec
}

func ToParamSpecVariant(pspeclike ParamSpecVariantLike) *ParamSpecVariant {
	t := (*ParamSpecVariant)(nil).GetStaticType()
	c := pspeclike.InheritedFromGParamSpecVariant()
	obj := ParamSpecGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ParamSpecVariant)(obj)
	}
	panic("cannot cast to ParamSpecVariant")
}

func (this *ParamSpecVariant) InheritedFromGParamSpecVariant() *C.GParamSpecVariant {
	return (*C.GParamSpecVariant)(this.C)
}

func (this *ParamSpecVariant) GetStaticType() Type {
	return Type(C._g_type_param_variant())
}

//--------------------------------------------------------------
// Object
//--------------------------------------------------------------

func object_finalizer(obj *Object) {
	if FQueue.Push(unsafe.Pointer(obj), object_finalizer2) {
		return
	}
	C.g_object_set_qdata((*C.GObject)(obj.C), C.uint32_t(go_repr), nil)
	C.g_object_unref((*C.GObject)(obj.C))
}

func object_finalizer2(obj_un unsafe.Pointer) {
	obj := (*Object)(obj_un)
	C.g_object_set_qdata((*C.GObject)(obj.C), C.uint32_t(go_repr), nil)
	C.g_object_unref((*C.GObject)(obj.C))
}

func set_object_finalizer(obj *Object) {
	runtime.SetFinalizer(obj, object_finalizer)
}

func ObjectWrap(c unsafe.Pointer, grab bool) unsafe.Pointer {
	if c == nil {
		return nil
	}
	obj := (*Object)(C.g_object_get_qdata((*C.GObject)(c), C.uint32_t(go_repr)))
	if obj != nil {
		return unsafe.Pointer(obj)
	}
	obj = &Object{c}
	if grab {
		C.g_object_ref_sink((*C.GObject)(obj.C))
	}
	set_object_finalizer(obj)
	C.g_object_set_qdata((*C.GObject)(obj.C),
		C.uint32_t(go_repr), unsafe.Pointer(obj))
	return unsafe.Pointer(obj)
}

func ObjectGrabIfType(c unsafe.Pointer, t Type) unsafe.Pointer {
	if c == nil {
		return nil
	}
	hasrepr := true
	obj := (*Object)(C.g_object_get_qdata((*C.GObject)(c), C.uint32_t(go_repr)))
	if obj == nil {
		obj = &Object{c}
		hasrepr = false
	}
	if obj.GetType().IsA(t) {
		if !hasrepr {
			C.g_object_ref_sink((*C.GObject)(obj.C))
			set_object_finalizer(obj)
			C.g_object_set_qdata((*C.GObject)(obj.C),
				C.uint32_t(go_repr), unsafe.Pointer(obj))
		}
		return unsafe.Pointer(obj)
	}
	return nil
}

func (this *Object) GetType() Type {
	return Type(C._g_object_type((*C.GObject)(this.C)))
}

func (this *Object) Connect(signal string, clo interface{}) {
	csignal := C.CString(signal)
	Holder.Grab(clo)
	goclosure := C.g_goclosure_new(unsafe.Pointer(&clo), nil)
	C.g_signal_connect_closure(this.C, csignal, (*C.GClosure)(unsafe.Pointer(goclosure)), 0)
	C.free(unsafe.Pointer(csignal))
}

func (this *Object) ConnectMethod(signal string, clo interface{}, recv interface{}) {
	csignal := C.CString(signal)
	Holder.Grab(clo)
	Holder.Grab(recv)
	goclosure := C.g_goclosure_new(unsafe.Pointer(&clo), unsafe.Pointer(&recv))
	C.g_signal_connect_closure(this.C, csignal, (*C.GClosure)(unsafe.Pointer(goclosure)), 0)
	C.free(unsafe.Pointer(csignal))

}

func (this *Object) FindProperty(name string) *ParamSpec {
	cname := C.CString(name)
	ret := C._g_object_find_property(this.InheritedFromGObject(), cname)
	C.free(unsafe.Pointer(cname))
	return (*ParamSpec)(ParamSpecWrap(unsafe.Pointer(ret), true))
}

func (this *Object) SetProperty(name string, value interface{}) {
	cname := C.CString(name)
	pspec := this.FindProperty(name)
	if pspec == nil {
		panic("Object has no property with that name: " + name)
	}
	var gvalue Value
	gvalue.Init(pspec.GetValueType())
	gvalue.SetGoInterface(value)
	C.g_object_set_property(this.InheritedFromGObject(), cname,
		(*C.GValue)(unsafe.Pointer(&gvalue)))
	gvalue.Unset()
	C.free(unsafe.Pointer(cname))
}

func (this *Object) GetProperty(name string, value interface{}) {
	cname := C.CString(name)
	pspec := this.FindProperty(name)
	if pspec == nil {
		panic("Object has no property with that name: " + name)
	}
	var gvalue Value
	gvalue.Init(pspec.GetValueType())
	C.g_object_get_property(this.InheritedFromGObject(), cname,
		(*C.GValue)(unsafe.Pointer(&gvalue)))
	gvalue.GetGoInterface(value)
	gvalue.Unset()
	C.free(unsafe.Pointer(cname))
}

func ObjectBindProperty(source ObjectLike, source_property string, target ObjectLike, target_property string, flags BindingFlags) *Binding {
	csource_property := C.CString(source_property)
	ctarget_property := C.CString(target_property)
	obj := C.g_object_bind_property(
		unsafe.Pointer(source.InheritedFromGObject()), csource_property,
		unsafe.Pointer(target.InheritedFromGObject()), ctarget_property,
		C.GBindingFlags(flags))
	C.free(unsafe.Pointer(csource_property))
	C.free(unsafe.Pointer(ctarget_property))
	return (*Binding)(ObjectWrap(unsafe.Pointer(obj), true))
}

func (this *Object) Unref() {
	runtime.SetFinalizer(this, nil)
	C.g_object_set_qdata((*C.GObject)(this.C), C.uint32_t(go_repr), nil)
	C.g_object_unref((*C.GObject)(this.C))
	this.C = nil
}

//--------------------------------------------------------------
// Closures
//--------------------------------------------------------------

//export g_goclosure_finalize_go
func g_goclosure_finalize_go(goclosure_up unsafe.Pointer) {
	goclosure := (*C.GGoClosure)(goclosure_up)
	clo := *(*interface{})(C.g_goclosure_get_func(goclosure))
	recv := *(*interface{})(C.g_goclosure_get_recv(goclosure))
	Holder.Release(clo)
	Holder.Release(recv)
}

//export g_goclosure_marshal_go
func g_goclosure_marshal_go(goclosure_up, ret_up unsafe.Pointer, nargs int32, args_up unsafe.Pointer) {
	var callargs [20]reflect.Value
	var recv reflect.Value
	goclosure := (*C.GGoClosure)(goclosure_up)
	ret := (*Value)(ret_up)
	args := (*(*[alot]Value)(args_up))[:nargs]
	f := reflect.ValueOf(*(*interface{})(C.g_goclosure_get_func(goclosure)))
	ft := f.Type()
	callargsn := ft.NumIn()

	recvi := *(*interface{})(C.g_goclosure_get_recv(goclosure))
	if recvi != nil {
		recv = reflect.ValueOf(recvi)
	}

	if callargsn >= 20 {
		panic("too many arguments in a closure")
	}

	for i, n := 0, callargsn; i < n; i++ {
		idx := i
		if recvi != nil {
			idx--
			if i == 0 {
				callargs[i] = recv
				continue
			}
		}

		in := ft.In(i)

		// use default value, if there is not enough args
		if len(args) <= idx {
			callargs[i] = reflect.New(in).Elem()
			continue
		}

		v := args[idx].GetGoValue(in)
		callargs[i] = v
	}

	out := f.Call(callargs[:callargsn])
	if len(out) == 1 {
		ret.SetGoValue(out[0])
	}
}

//--------------------------------------------------------------
// Go Interface boxed type
//--------------------------------------------------------------

//export g_go_interface_copy_go
func g_go_interface_copy_go(boxed unsafe.Pointer) unsafe.Pointer {
	Holder.Grab(*(*interface{})(boxed))
	newboxed := C.malloc(C.size_t(unsafe.Sizeof([2]unsafe.Pointer{})))
	C.memcpy(newboxed, boxed, C.size_t(unsafe.Sizeof([2]unsafe.Pointer{})))
	return newboxed
}

//export g_go_interface_free_go
func g_go_interface_free_go(boxed unsafe.Pointer) {
	Holder.Release(*(*interface{})(boxed))
	C.free(boxed)
}

//--------------------------------------------------------------
// Type
//--------------------------------------------------------------

type Type C.GType

func (this Type) IsA(other Type) bool {
	return C.g_type_is_a(C.GType(this), C.GType(other)) != 0
}

func (this Type) String() string {
	cname := C.g_type_name(C.GType(this))
	if cname == nil {
		return ""
	}
	return C.GoString(cname)
}

func (this Type) asC() C.GType {
	return C.GType(this)
}

var (
	Interface Type
	Char Type
	UChar Type
	Boolean Type
	Int Type
	UInt Type
	Long Type
	ULong Type
	Int64 Type
	UInt64 Type
	Enum Type
	Flags Type
	Float Type
	Double Type
	String Type
	Pointer Type
	Boxed Type
	Param Type
	GObject Type
	GType Type
	Variant Type
	GoInterface Type
)

func init() {
	C.g_type_init()

	Interface = Type(C._g_type_interface())
	Char = Type(C._g_type_char())
	UChar = Type(C._g_type_uchar())
	Boolean = Type(C._g_type_boolean())
	Int = Type(C._g_type_int())
	UInt = Type(C._g_type_uint())
	Long = Type(C._g_type_long())
	ULong = Type(C._g_type_ulong())
	Int64 = Type(C._g_type_int64())
	UInt64 = Type(C._g_type_uint64())
	Enum = Type(C._g_type_enum())
	Flags = Type(C._g_type_flags())
	Float = Type(C._g_type_float())
	Double = Type(C._g_type_double())
	String = Type(C._g_type_string())
	Pointer = Type(C._g_type_pointer())
	Boxed = Type(C._g_type_boxed())
	Param = Type(C._g_type_param())
	GObject = Type(C._g_type_object())
	GType = Type(C._g_type_gtype())
	Variant = Type(C._g_type_variant())
	GoInterface = Type(C._g_type_go_interface())
}

// Every GObject generated by this generator implements this interface
// and it must work even if the receiver is a nil value
type StaticTyper interface {
	GetStaticType() Type
}

//--------------------------------------------------------------
// Value
//--------------------------------------------------------------

func (this *Value) asC() *C.GValue {
	return (*C.GValue)(unsafe.Pointer(this))
}

// g_value_init
func (this *Value) Init(t Type) {
	C.g_value_init(this.asC(), t.asC())
}

// g_value_copy
func (this *Value) Set(src *Value) {
	C.g_value_copy(src.asC(), this.asC())
}

// g_value_reset
func (this *Value) Reset() {
	C.g_value_reset(this.asC())
}

// g_value_unset
func (this *Value) Unset() {
	C.g_value_unset(this.asC())
}

// G_VALUE_TYPE
func (this *Value) GetType() Type {
	return Type(C._g_value_type(this.asC()))
}

// g_value_type_compatible
func ValueTypeCompatible(src, dst Type) bool {
	return C.g_value_type_compatible(src.asC(), dst.asC()) != 0
}

// g_value_type_transformable
func ValueTypeTransformable(src, dst Type) bool {
	return C.g_value_type_transformable(src.asC(), dst.asC()) != 0
}

// g_value_transform
func (this *Value) Transform(src *Value) bool {
	return C.g_value_transform(src.asC(), this.asC()) != 0
}

// g_value_get_boolean
func (this *Value) GetBool() bool {
	return C.g_value_get_boolean(this.asC()) != 0
}

// g_value_set_boolean
func (this *Value) SetBool(v bool) {
	C.g_value_set_boolean(this.asC(), _GoBoolToCBool(v))
}

// g_value_get_int64
func (this *Value) GetInt() int64 {
	return int64(C.g_value_get_int64(this.asC()))
}

// g_value_set_int64
func (this *Value) SetInt(v int64) {
	C.g_value_set_int64(this.asC(), C.int64_t(v))
}

// g_value_get_uint64
func (this *Value) GetUint() uint64 {
	return uint64(C.g_value_get_uint64(this.asC()))
}

// g_value_set_uint64
func (this *Value) SetUint(v uint64) {
	C.g_value_set_uint64(this.asC(), C.uint64_t(v))
}

// g_value_get_double
func (this *Value) GetFloat() float64 {
	return float64(C.g_value_get_double(this.asC()))
}

// g_value_set_double
func (this *Value) SetFloat(v float64) {
	C.g_value_set_double(this.asC(), C.double(v))
}

// g_value_get_string
func (this *Value) GetString() string {
	return C.GoString(C.g_value_get_string(this.asC()))
}

// g_value_take_string
func (this *Value) SetString(v string) {
	cstr := C.CString(v)
	C.g_value_take_string(this.asC(), cstr)
	// not freeing, because GValue takes the ownership
}

// g_value_get_object
func (this *Value) GetObject() unsafe.Pointer {
	return unsafe.Pointer(C.g_value_get_object(this.asC()))
}

// g_value_set_object
func (this *Value) SetObject(x unsafe.Pointer) {
	C.g_value_set_object(this.asC(), (*C.GObject)(x))
}

// g_value_get_boxed
func (this *Value) GetBoxed() unsafe.Pointer {
	return C.g_value_get_boxed(this.asC())
}

// g_value_take_boxed
func (this *Value) SetBoxed(x unsafe.Pointer) {
	C.g_value_take_boxed(this.asC(), x)
}

func (this *Value) GetBoxedInterface() interface{} {
	return *(*interface{})(C.g_value_get_boxed(this.asC()))
}

func (this *Value) SetBoxedInterface(x interface{}) {
	Holder.Grab(x)
	newboxed := C.malloc(C.size_t(unsafe.Sizeof([2]unsafe.Pointer{})))
	C.memcpy(newboxed, unsafe.Pointer(&x), C.size_t(unsafe.Sizeof([2]unsafe.Pointer{})))
	C.g_value_take_boxed(this.asC(), newboxed)
}

//--------------------------------------------------------------
// A giant glue for connecting GType and Go's reflection
//--------------------------------------------------------------

var statictyper = reflect.TypeOf((*StaticTyper)(nil)).Elem()
var objectlike = reflect.TypeOf((*ObjectLike)(nil)).Elem()

func (this *Value) SetGoValue(v reflect.Value) {
	valuetype := this.GetType()
	var src Value

	if valuetype == GoInterface {
		// special case
		this.SetBoxedInterface(v.Interface())
		return
	}

	transform := func() {
		ok := this.Transform(&src)
		if !ok {
			panic("Go value (" + v.Type().String() + ") is not transformable to " + valuetype.String())
		}
	}

	switch v.Kind() {
	case reflect.Bool:
		src.Init(Boolean)
		src.SetBool(v.Bool())
		transform()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		src.Init(Int64)
		src.SetInt(v.Int())
		transform()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		src.Init(UInt64)
		src.SetUint(v.Uint())
		transform()
	case reflect.Float32, reflect.Float64:
		src.Init(Double)
		src.SetFloat(v.Float())
		transform()
	case reflect.String:
		src.Init(String)
		src.SetString(v.String())
		transform()
		src.Unset()
	case reflect.Ptr:
		gotype := v.Type()
		src.Init(GObject)
		if gotype.Implements(objectlike) {
			obj, ok := v.Interface().(ObjectLike)
			if !ok {
				panic(gotype.String() + " is not transformable to GValue")
			}

			src.SetObject(unsafe.Pointer(obj.InheritedFromGObject()))
			transform()
		}
		src.Unset()
	}
}

var CairoMarshaler func(*Value, reflect.Type) (reflect.Value, bool)

func (this *Value) GetGoValue(t reflect.Type) reflect.Value {
	var out reflect.Value
	var dst Value

	if (this.GetType() == GoInterface) {
		return reflect.ValueOf(this.GetBoxedInterface())
	}

	transform := func() {
		ok := dst.Transform(this)
		if !ok {
			panic("GValue is not transformable to " + t.String())
		}
	}

	switch t.Kind() {
	case reflect.Bool:
		dst.Init(Boolean)
		transform()
		out = reflect.New(t).Elem()
		out.SetBool(dst.GetBool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst.Init(Int64)
		transform()
		out = reflect.New(t).Elem()
		out.SetInt(dst.GetInt())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		dst.Init(UInt64)
		transform()
		out = reflect.New(t).Elem()
		out.SetUint(dst.GetUint())
	case reflect.Float32, reflect.Float64:
		dst.Init(Double)
		transform()
		out = reflect.New(t).Elem()
		out.SetFloat(dst.GetFloat())
	case reflect.String:
		dst.Init(String)
		transform()
		out = reflect.New(t).Elem()
		out.SetString(dst.GetString())
		dst.Unset() // need to clean up in this case
	case reflect.Ptr:
		if t.Implements(objectlike) {
			// at this point we're sure that this is a pointer to the ObjectLike
			out = reflect.New(t)
			st, ok := out.Elem().Interface().(StaticTyper)
			if !ok {
				panic("ObjectLike type must implement StaticTyper as well")
			}
			dst.Init(st.GetStaticType())
			transform()
			*(*unsafe.Pointer)(unsafe.Pointer(out.Pointer())) = ObjectWrap(dst.GetObject(), true)
			dst.Unset()
			out = out.Elem()
		} else {
			// cairo marshaler hook
			if CairoMarshaler != nil {
				var ok bool
				out, ok = CairoMarshaler(this, t)
				if ok {
					break
				}
			}

			// must be a struct then
			out = reflect.New(t)
			*(*unsafe.Pointer)(unsafe.Pointer(out.Pointer())) = this.GetBoxed()
			out = out.Elem()
		}
	}
	return out
}

func (this *Value) SetGoInterface(v interface{}) {
	this.SetGoValue(reflect.ValueOf(v))
}

func (this *Value) GetGoInterface(v interface{}) {
	vp := reflect.ValueOf(v)
	if vp.Kind() != reflect.Ptr {
		panic("a pointer to value is expected for Value.GetGoInterface")
	}
	vp.Elem().Set(this.GetGoValue(vp.Type().Elem()))
}