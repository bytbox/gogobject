package atk

/*
#include <stdlib.h>
#include <stdint.h>

typedef size_t GType;

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
typedef struct _AtkAction AtkAction;
typedef struct _AtkActionIface AtkActionIface;
typedef struct _AtkAttribute AtkAttribute;
typedef struct _AtkComponent AtkComponent;
typedef struct _AtkComponentIface AtkComponentIface;
typedef uint32_t AtkCoordType;
typedef struct _AtkDocument AtkDocument;
typedef struct _AtkDocumentIface AtkDocumentIface;
typedef struct _AtkEditableText AtkEditableText;
typedef struct _AtkEditableTextIface AtkEditableTextIface;
typedef void* AtkEventListener;
extern void _AtkEventListener_c_wrapper();
extern void _AtkEventListener_c_wrapper_once();
typedef void* AtkEventListenerInit;
extern void _AtkEventListenerInit_c_wrapper();
extern void _AtkEventListenerInit_c_wrapper_once();
typedef void* AtkFocusHandler;
extern void _AtkFocusHandler_c_wrapper();
extern void _AtkFocusHandler_c_wrapper_once();
typedef void* AtkFunction;
extern void _AtkFunction_c_wrapper();
extern void _AtkFunction_c_wrapper_once();
typedef struct _AtkGObjectAccessible AtkGObjectAccessible;
typedef struct _AtkGObjectAccessibleClass AtkGObjectAccessibleClass;
typedef struct _AtkHyperlink AtkHyperlink;
typedef struct _AtkHyperlinkClass AtkHyperlinkClass;
typedef struct _AtkHyperlinkImpl AtkHyperlinkImpl;
typedef struct _AtkHyperlinkImplIface AtkHyperlinkImplIface;
typedef uint32_t AtkHyperlinkStateFlags;
typedef struct _AtkHypertext AtkHypertext;
typedef struct _AtkHypertextIface AtkHypertextIface;
typedef struct _AtkImage AtkImage;
typedef struct _AtkImageIface AtkImageIface;
typedef struct _AtkImplementor AtkImplementor;
typedef struct _AtkImplementorIface AtkImplementorIface;
typedef struct _AtkKeyEventStruct AtkKeyEventStruct;
typedef uint32_t AtkKeyEventType;
typedef void* AtkKeySnoopFunc;
extern void _AtkKeySnoopFunc_c_wrapper();
extern void _AtkKeySnoopFunc_c_wrapper_once();
typedef uint32_t AtkLayer;
typedef struct _AtkMisc AtkMisc;
typedef struct _AtkMiscClass AtkMiscClass;
typedef struct _AtkNoOpObject AtkNoOpObject;
typedef struct _AtkNoOpObjectClass AtkNoOpObjectClass;
typedef struct _AtkNoOpObjectFactory AtkNoOpObjectFactory;
typedef struct _AtkNoOpObjectFactoryClass AtkNoOpObjectFactoryClass;
typedef struct _AtkObject AtkObject;
typedef struct _AtkObjectClass AtkObjectClass;
typedef struct _AtkObjectFactory AtkObjectFactory;
typedef struct _AtkObjectFactoryClass AtkObjectFactoryClass;
typedef struct _AtkPlug AtkPlug;
typedef struct _AtkPlugClass AtkPlugClass;
typedef struct _AtkRectangle AtkRectangle;
typedef struct _AtkRegistry AtkRegistry;
typedef struct _AtkRelation AtkRelation;
typedef struct _AtkRelationClass AtkRelationClass;
typedef struct _AtkRelationSet AtkRelationSet;
typedef struct _AtkRelationSetClass AtkRelationSetClass;
typedef uint32_t AtkRelationType;
typedef uint32_t AtkRole;
typedef struct _AtkSelection AtkSelection;
typedef struct _AtkSelectionIface AtkSelectionIface;
typedef struct _AtkSocket AtkSocket;
typedef struct _AtkSocketClass AtkSocketClass;
typedef struct _AtkStateSet AtkStateSet;
typedef struct _AtkStateSetClass AtkStateSetClass;
typedef uint32_t AtkStateType;
typedef struct _AtkStreamableContent AtkStreamableContent;
typedef struct _AtkStreamableContentIface AtkStreamableContentIface;
typedef struct _AtkTable AtkTable;
typedef struct _AtkTableIface AtkTableIface;
typedef struct _AtkText AtkText;
typedef uint32_t AtkTextAttribute;
typedef uint32_t AtkTextBoundary;
typedef uint32_t AtkTextClipType;
typedef struct _AtkTextIface AtkTextIface;
typedef struct _AtkTextRange AtkTextRange;
typedef struct _AtkTextRectangle AtkTextRectangle;
typedef struct _AtkUtil AtkUtil;
typedef struct _AtkUtilClass AtkUtilClass;
typedef struct _AtkValue AtkValue;
typedef struct _AtkValueIface AtkValueIface;
typedef struct _AtkWindow AtkWindow;
typedef struct _AtkWindowIface AtkWindowIface;
typedef struct _Atk_PropertyValues Atk_PropertyValues;
typedef struct _Atk_Registry Atk_Registry;
typedef struct _Atk_RegistryClass Atk_RegistryClass;
extern int atk_action_do_action(AtkAction*, int32_t);
extern char* atk_action_get_description(AtkAction*, int32_t);
extern char* atk_action_get_keybinding(AtkAction*, int32_t);
extern char* atk_action_get_localized_name(AtkAction*, int32_t);
extern int32_t atk_action_get_n_actions(AtkAction*);
extern char* atk_action_get_name(AtkAction*, int32_t);
extern int atk_action_set_description(AtkAction*, int32_t, char*);
extern GType atk_action_get_type();
extern void atk_attribute_set_free(GSList*);
extern int atk_component_contains(AtkComponent*, int32_t, int32_t, AtkCoordType);
extern double atk_component_get_alpha(AtkComponent*);
extern void atk_component_get_extents(AtkComponent*, int32_t*, int32_t*, int32_t*, int32_t*, AtkCoordType);
extern AtkLayer atk_component_get_layer(AtkComponent*);
extern int32_t atk_component_get_mdi_zorder(AtkComponent*);
extern void atk_component_get_position(AtkComponent*, int32_t*, int32_t*, AtkCoordType);
extern void atk_component_get_size(AtkComponent*, int32_t*, int32_t*);
extern int atk_component_grab_focus(AtkComponent*);
extern AtkObject* atk_component_ref_accessible_at_point(AtkComponent*, int32_t, int32_t, AtkCoordType);
extern void atk_component_remove_focus_handler(AtkComponent*, uint32_t);
extern int atk_component_set_extents(AtkComponent*, int32_t, int32_t, int32_t, int32_t, AtkCoordType);
extern int atk_component_set_position(AtkComponent*, int32_t, int32_t, AtkCoordType);
extern int atk_component_set_size(AtkComponent*, int32_t, int32_t);
extern GType atk_component_get_type();
extern char* atk_document_get_attribute_value(AtkDocument*, char*);
extern GSList* atk_document_get_attributes(AtkDocument*);
extern void* atk_document_get_document(AtkDocument*);
extern char* atk_document_get_document_type(AtkDocument*);
extern char* atk_document_get_locale(AtkDocument*);
extern int atk_document_set_attribute_value(AtkDocument*, char*, char*);
extern GType atk_document_get_type();
extern void atk_editable_text_copy_text(AtkEditableText*, int32_t, int32_t);
extern void atk_editable_text_cut_text(AtkEditableText*, int32_t, int32_t);
extern void atk_editable_text_delete_text(AtkEditableText*, int32_t, int32_t);
extern void atk_editable_text_insert_text(AtkEditableText*, char*, int32_t, int32_t*);
extern void atk_editable_text_paste_text(AtkEditableText*, int32_t);
extern int atk_editable_text_set_run_attributes(AtkEditableText*, GSList*, int32_t, int32_t);
extern void atk_editable_text_set_text_contents(AtkEditableText*, char*);
extern GType atk_editable_text_get_type();
extern AtkObject* atk_gobject_accessible_for_object(GObject*);
extern GObject* atk_gobject_accessible_get_object(AtkGObjectAccessible*);
extern GType atk_gobject_accessible_get_type();
extern int32_t atk_hyperlink_get_end_index(AtkHyperlink*);
extern int32_t atk_hyperlink_get_n_anchors(AtkHyperlink*);
extern AtkObject* atk_hyperlink_get_object(AtkHyperlink*, int32_t);
extern int32_t atk_hyperlink_get_start_index(AtkHyperlink*);
extern char* atk_hyperlink_get_uri(AtkHyperlink*, int32_t);
extern int atk_hyperlink_is_inline(AtkHyperlink*);
extern int atk_hyperlink_is_valid(AtkHyperlink*);
extern GType atk_hyperlink_get_type();
extern AtkHyperlink* atk_hyperlink_impl_get_hyperlink(AtkHyperlinkImpl*);
extern GType atk_hyperlink_impl_get_type();
extern AtkHyperlink* atk_hypertext_get_link(AtkHypertext*, int32_t);
extern int32_t atk_hypertext_get_link_index(AtkHypertext*, int32_t);
extern int32_t atk_hypertext_get_n_links(AtkHypertext*);
extern GType atk_hypertext_get_type();
extern char* atk_image_get_image_description(AtkImage*);
extern char* atk_image_get_image_locale(AtkImage*);
extern void atk_image_get_image_position(AtkImage*, int32_t*, int32_t*, AtkCoordType);
extern void atk_image_get_image_size(AtkImage*, int32_t*, int32_t*);
extern int atk_image_set_image_description(AtkImage*, char*);
extern GType atk_image_get_type();
extern AtkObject* atk_implementor_ref_accessible(AtkImplementor*);
extern GType atk_implementor_get_type();
extern AtkMisc* atk_misc_get_instance();
extern void atk_misc_threads_enter(AtkMisc*);
extern void atk_misc_threads_leave(AtkMisc*);
extern GType atk_misc_get_type();
extern AtkObject* atk_no_op_object_new(GObject*);
extern GType atk_no_op_object_get_type();
extern AtkObjectFactory* atk_no_op_object_factory_new();
extern GType atk_no_op_object_factory_get_type();
extern int atk_object_add_relationship(AtkObject*, AtkRelationType, AtkObject*);
extern GSList* atk_object_get_attributes(AtkObject*);
extern char* atk_object_get_description(AtkObject*);
extern int32_t atk_object_get_index_in_parent(AtkObject*);
extern int32_t atk_object_get_n_accessible_children(AtkObject*);
extern char* atk_object_get_name(AtkObject*);
extern AtkObject* atk_object_get_parent(AtkObject*);
extern AtkRole atk_object_get_role(AtkObject*);
extern void atk_object_initialize(AtkObject*, void*);
extern void atk_object_notify_state_change(AtkObject*, uint64_t, int);
extern AtkObject* atk_object_ref_accessible_child(AtkObject*, int32_t);
extern AtkRelationSet* atk_object_ref_relation_set(AtkObject*);
extern AtkStateSet* atk_object_ref_state_set(AtkObject*);
extern void atk_object_remove_property_change_handler(AtkObject*, uint32_t);
extern int atk_object_remove_relationship(AtkObject*, AtkRelationType, AtkObject*);
extern void atk_object_set_description(AtkObject*, char*);
extern void atk_object_set_name(AtkObject*, char*);
extern void atk_object_set_parent(AtkObject*, AtkObject*);
extern void atk_object_set_role(AtkObject*, AtkRole);
extern GType atk_object_get_type();
extern AtkObject* atk_object_factory_create_accessible(AtkObjectFactory*, GObject*);
extern GType atk_object_factory_get_accessible_type(AtkObjectFactory*);
extern void atk_object_factory_invalidate(AtkObjectFactory*);
extern GType atk_object_factory_get_type();
extern AtkObject* atk_plug_new();
extern char* atk_plug_get_id(AtkPlug*);
extern GType atk_plug_get_type();
extern AtkObjectFactory* atk_registry_get_factory(AtkRegistry*, GType);
extern GType atk_registry_get_factory_type(AtkRegistry*, GType);
extern void atk_registry_set_factory_type(AtkRegistry*, GType, GType);
extern GType atk_registry_get_type();
extern AtkRelation* atk_relation_new(AtkObject*, int32_t, AtkRelationType);
extern void atk_relation_add_target(AtkRelation*, AtkObject*);
extern AtkRelationType atk_relation_get_relation_type(AtkRelation*);
extern GPtrArray* atk_relation_get_target(AtkRelation*);
extern int atk_relation_remove_target(AtkRelation*, AtkObject*);
extern GType atk_relation_get_type();
extern AtkRelationSet* atk_relation_set_new();
extern void atk_relation_set_add(AtkRelationSet*, AtkRelation*);
extern void atk_relation_set_add_relation_by_type(AtkRelationSet*, AtkRelationType, AtkObject*);
extern int atk_relation_set_contains(AtkRelationSet*, AtkRelationType);
extern int32_t atk_relation_set_get_n_relations(AtkRelationSet*);
extern AtkRelation* atk_relation_set_get_relation(AtkRelationSet*, int32_t);
extern AtkRelation* atk_relation_set_get_relation_by_type(AtkRelationSet*, AtkRelationType);
extern void atk_relation_set_remove(AtkRelationSet*, AtkRelation*);
extern GType atk_relation_set_get_type();
extern int atk_selection_add_selection(AtkSelection*, int32_t);
extern int atk_selection_clear_selection(AtkSelection*);
extern int32_t atk_selection_get_selection_count(AtkSelection*);
extern int atk_selection_is_child_selected(AtkSelection*, int32_t);
extern AtkObject* atk_selection_ref_selection(AtkSelection*, int32_t);
extern int atk_selection_remove_selection(AtkSelection*, int32_t);
extern int atk_selection_select_all_selection(AtkSelection*);
extern GType atk_selection_get_type();
extern AtkObject* atk_socket_new();
extern void atk_socket_embed(AtkSocket*, char*);
extern int atk_socket_is_occupied(AtkSocket*);
extern GType atk_socket_get_type();
extern AtkStateSet* atk_state_set_new();
extern int atk_state_set_add_state(AtkStateSet*, AtkStateType);
extern void atk_state_set_add_states(AtkStateSet*, AtkStateType*, int32_t);
extern AtkStateSet* atk_state_set_and_sets(AtkStateSet*, AtkStateSet*);
extern void atk_state_set_clear_states(AtkStateSet*);
extern int atk_state_set_contains_state(AtkStateSet*, AtkStateType);
extern int atk_state_set_contains_states(AtkStateSet*, AtkStateType*, int32_t);
extern int atk_state_set_is_empty(AtkStateSet*);
extern AtkStateSet* atk_state_set_or_sets(AtkStateSet*, AtkStateSet*);
extern int atk_state_set_remove_state(AtkStateSet*, AtkStateType);
extern AtkStateSet* atk_state_set_xor_sets(AtkStateSet*, AtkStateSet*);
extern GType atk_state_set_get_type();
extern char* atk_streamable_content_get_mime_type(AtkStreamableContent*, int32_t);
extern int32_t atk_streamable_content_get_n_mime_types(AtkStreamableContent*);
extern GIOChannel* atk_streamable_content_get_stream(AtkStreamableContent*, char*);
extern char* atk_streamable_content_get_uri(AtkStreamableContent*, char*);
extern GType atk_streamable_content_get_type();
extern int atk_table_add_column_selection(AtkTable*, int32_t);
extern int atk_table_add_row_selection(AtkTable*, int32_t);
extern AtkObject* atk_table_get_caption(AtkTable*);
extern int32_t atk_table_get_column_at_index(AtkTable*, int32_t);
extern char* atk_table_get_column_description(AtkTable*, int32_t);
extern int32_t atk_table_get_column_extent_at(AtkTable*, int32_t, int32_t);
extern AtkObject* atk_table_get_column_header(AtkTable*, int32_t);
extern int32_t atk_table_get_index_at(AtkTable*, int32_t, int32_t);
extern int32_t atk_table_get_n_columns(AtkTable*);
extern int32_t atk_table_get_n_rows(AtkTable*);
extern int32_t atk_table_get_row_at_index(AtkTable*, int32_t);
extern char* atk_table_get_row_description(AtkTable*, int32_t);
extern int32_t atk_table_get_row_extent_at(AtkTable*, int32_t, int32_t);
extern AtkObject* atk_table_get_row_header(AtkTable*, int32_t);
extern int32_t atk_table_get_selected_columns(AtkTable*, int32_t*);
extern int32_t atk_table_get_selected_rows(AtkTable*, int32_t*);
extern AtkObject* atk_table_get_summary(AtkTable*);
extern int atk_table_is_column_selected(AtkTable*, int32_t);
extern int atk_table_is_row_selected(AtkTable*, int32_t);
extern int atk_table_is_selected(AtkTable*, int32_t, int32_t);
extern AtkObject* atk_table_ref_at(AtkTable*, int32_t, int32_t);
extern int atk_table_remove_column_selection(AtkTable*, int32_t);
extern int atk_table_remove_row_selection(AtkTable*, int32_t);
extern void atk_table_set_caption(AtkTable*, AtkObject*);
extern void atk_table_set_column_description(AtkTable*, int32_t, char*);
extern void atk_table_set_column_header(AtkTable*, int32_t, AtkObject*);
extern void atk_table_set_row_description(AtkTable*, int32_t, char*);
extern void atk_table_set_row_header(AtkTable*, int32_t, AtkObject*);
extern void atk_table_set_summary(AtkTable*, AtkObject*);
extern GType atk_table_get_type();
extern void atk_text_free_ranges(AtkTextRange*);
extern int atk_text_add_selection(AtkText*, int32_t, int32_t);
extern AtkTextRange** atk_text_get_bounded_ranges(AtkText*, AtkTextRectangle*, AtkCoordType, AtkTextClipType, AtkTextClipType);
extern int32_t atk_text_get_caret_offset(AtkText*);
extern uint32_t atk_text_get_character_at_offset(AtkText*, int32_t);
extern int32_t atk_text_get_character_count(AtkText*);
extern void atk_text_get_character_extents(AtkText*, int32_t, int32_t*, int32_t*, int32_t*, int32_t*, AtkCoordType);
extern GSList* atk_text_get_default_attributes(AtkText*);
extern int32_t atk_text_get_n_selections(AtkText*);
extern int32_t atk_text_get_offset_at_point(AtkText*, int32_t, int32_t, AtkCoordType);
extern void atk_text_get_range_extents(AtkText*, int32_t, int32_t, AtkCoordType, AtkTextRectangle*);
extern GSList* atk_text_get_run_attributes(AtkText*, int32_t, int32_t*, int32_t*);
extern char* atk_text_get_selection(AtkText*, int32_t, int32_t*, int32_t*);
extern char* atk_text_get_text(AtkText*, int32_t, int32_t);
extern char* atk_text_get_text_after_offset(AtkText*, int32_t, AtkTextBoundary, int32_t*, int32_t*);
extern char* atk_text_get_text_at_offset(AtkText*, int32_t, AtkTextBoundary, int32_t*, int32_t*);
extern char* atk_text_get_text_before_offset(AtkText*, int32_t, AtkTextBoundary, int32_t*, int32_t*);
extern int atk_text_remove_selection(AtkText*, int32_t);
extern int atk_text_set_caret_offset(AtkText*, int32_t);
extern int atk_text_set_selection(AtkText*, int32_t, int32_t, int32_t);
extern GType atk_text_get_type();
extern GType atk_util_get_type();
extern void atk_value_get_current_value(AtkValue*, GValue*);
extern void atk_value_get_maximum_value(AtkValue*, GValue*);
extern void atk_value_get_minimum_increment(AtkValue*, GValue*);
extern void atk_value_get_minimum_value(AtkValue*, GValue*);
extern int atk_value_set_current_value(AtkValue*, GValue*);
extern GType atk_value_get_type();
extern GType atk_window_get_type();
extern void atk_focus_tracker_notify(AtkObject*);
extern AtkRegistry* atk_get_default_registry();
extern AtkObject* atk_get_focus_object();
extern AtkObject* atk_get_root();
extern char* atk_get_toolkit_name();
extern char* atk_get_toolkit_version();
extern char* atk_get_version();
extern AtkRelationType atk_relation_type_for_name(char*);
extern char* atk_relation_type_get_name(AtkRelationType);
extern AtkRelationType atk_relation_type_register(char*);
extern void atk_remove_focus_tracker(uint32_t);
extern void atk_remove_global_event_listener(uint32_t);
extern void atk_remove_key_event_listener(uint32_t);
extern AtkRole atk_role_for_name(char*);
extern char* atk_role_get_localized_name(AtkRole);
extern char* atk_role_get_name(AtkRole);
extern AtkRole atk_role_register(char*);
extern AtkStateType atk_state_type_for_name(char*);
extern char* atk_state_type_get_name(AtkStateType);
extern AtkStateType atk_state_type_register(char*);
extern AtkTextAttribute atk_text_attribute_for_name(char*);
extern char* atk_text_attribute_get_name(AtkTextAttribute);
extern char* atk_text_attribute_get_value(AtkTextAttribute, int32_t);
extern AtkTextAttribute atk_text_attribute_register(char*);
struct _AtkActionIface { uint8_t _data[80]; };
struct _AtkAttribute { uint8_t _data[16]; };
struct _AtkComponentIface { uint8_t _data[136]; };
struct _AtkDocumentIface { uint8_t _data[96]; };
struct _AtkEditableTextIface { uint8_t _data[88]; };
struct _AtkGObjectAccessibleClass { uint8_t _data[368]; };
struct _AtkHyperlinkClass { uint8_t _data[216]; };
struct _AtkHyperlinkImplIface { uint8_t _data[32]; };
struct _AtkHypertextIface { uint8_t _data[72]; };
struct _AtkImageIface { uint8_t _data[64]; };
struct _AtkImplementor {};
struct _AtkKeyEventStruct { uint8_t _data[32]; };
struct _AtkMiscClass { uint8_t _data[408]; };
struct _AtkNoOpObjectClass { uint8_t _data[352]; };
struct _AtkNoOpObjectFactoryClass { uint8_t _data[176]; };
struct _AtkObjectClass { uint8_t _data[352]; };
struct _AtkObjectFactoryClass { uint8_t _data[176]; };
struct _AtkPlugClass { uint8_t _data[360]; };
struct _AtkRectangle { uint8_t _data[16]; };
struct _AtkRelationClass { uint8_t _data[136]; };
struct _AtkRelationSetClass { uint8_t _data[152]; };
struct _AtkSelectionIface { uint8_t _data[96]; };
struct _AtkSocketClass { uint8_t _data[360]; };
struct _AtkStateSetClass { uint8_t _data[136]; };
struct _AtkStreamableContentIface { uint8_t _data[72]; };
struct _AtkTableIface { uint8_t _data[336]; };
struct _AtkTextIface { uint8_t _data[208]; };
struct _AtkTextRange { uint8_t _data[32]; };
struct _AtkTextRectangle { uint8_t _data[16]; };
struct _AtkUtilClass { uint8_t _data[192]; };
struct _AtkValueIface { uint8_t _data[64]; };
struct _AtkWindowIface { uint8_t _data[144]; };
struct _Atk_PropertyValues { uint8_t _data[56]; };
struct _Atk_Registry { uint8_t _data[40]; };
struct _Atk_RegistryClass { uint8_t _data[136]; };


extern GObject *g_object_ref_sink(GObject*);
extern void g_object_unref(GObject*);
extern void g_error_free(GError*);
extern void g_free(void*);

#cgo pkg-config: atk
*/
import "C"
import "unsafe"

import (
	"github.com/bytbox/gogobject/gobject-2.0"
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


//export _Atk_go_callback_cleanup
func _Atk_go_callback_cleanup(gofunc unsafe.Pointer) {
	gobject.Holder.Release(gofunc)
}


type ActionLike interface {
	ImplementsAtkAction() *C.AtkAction
}

type Action struct {
	gobject.Object
	ActionImpl
}

type ActionImpl struct {}

func ToAction(objlike gobject.ObjectLike) *Action {
	t := (*ActionImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Action)(obj)
	}
	panic("cannot cast to Action")
}

func (this0 *ActionImpl) ImplementsAtkAction() *C.AtkAction {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkAction)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *ActionImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_action_get_type())
}

func ActionGetType() gobject.Type {
	return (*ActionImpl)(nil).GetStaticType()
}
func (this0 *ActionImpl) DoAction(i0 int) bool {
	var this1 *C.AtkAction
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkAction()}
	i1 = C.int32_t(i0)
	ret1 := C.atk_action_do_action(this1, i1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *ActionImpl) GetDescription(i0 int) string {
	var this1 *C.AtkAction
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkAction()}
	i1 = C.int32_t(i0)
	ret1 := C.atk_action_get_description(this1, i1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *ActionImpl) GetKeybinding(i0 int) string {
	var this1 *C.AtkAction
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkAction()}
	i1 = C.int32_t(i0)
	ret1 := C.atk_action_get_keybinding(this1, i1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *ActionImpl) GetLocalizedName(i0 int) string {
	var this1 *C.AtkAction
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkAction()}
	i1 = C.int32_t(i0)
	ret1 := C.atk_action_get_localized_name(this1, i1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *ActionImpl) GetNActions() int {
	var this1 *C.AtkAction
	if this0 != nil {
		this1 = this0.ImplementsAtkAction()}
	ret1 := C.atk_action_get_n_actions(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *ActionImpl) GetName(i0 int) string {
	var this1 *C.AtkAction
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkAction()}
	i1 = C.int32_t(i0)
	ret1 := C.atk_action_get_name(this1, i1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *ActionImpl) SetDescription(i0 int, desc0 string) bool {
	var this1 *C.AtkAction
	var i1 C.int32_t
	var desc1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsAtkAction()}
	i1 = C.int32_t(i0)
	desc1 = _GoStringToGString(desc0)
	defer C.free(unsafe.Pointer(desc1))
	ret1 := C.atk_action_set_description(this1, i1, desc1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type Attribute struct {
	name0 *C.char
	value0 *C.char
}
func (this0 *Attribute) Name() string {
	var name1 string
	name1 = C.GoString(this0.name0)
	return name1
}
func (this0 *Attribute) Value() string {
	var value1 string
	value1 = C.GoString(this0.value0)
	return value1
}
func AttributeSetFree(attrib_set0 []unsafe.Pointer) {
	var attrib_set1 *C.GSList
	C.atk_attribute_set_free(attrib_set1)
}
type ComponentLike interface {
	ImplementsAtkComponent() *C.AtkComponent
}

type Component struct {
	gobject.Object
	ComponentImpl
}

type ComponentImpl struct {}

func ToComponent(objlike gobject.ObjectLike) *Component {
	t := (*ComponentImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Component)(obj)
	}
	panic("cannot cast to Component")
}

func (this0 *ComponentImpl) ImplementsAtkComponent() *C.AtkComponent {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkComponent)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *ComponentImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_component_get_type())
}

func ComponentGetType() gobject.Type {
	return (*ComponentImpl)(nil).GetStaticType()
}
func (this0 *ComponentImpl) Contains(x0 int, y0 int, coord_type0 CoordType) bool {
	var this1 *C.AtkComponent
	var x1 C.int32_t
	var y1 C.int32_t
	var coord_type1 C.AtkCoordType
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	coord_type1 = C.AtkCoordType(coord_type0)
	ret1 := C.atk_component_contains(this1, x1, y1, coord_type1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *ComponentImpl) GetAlpha() float64 {
	var this1 *C.AtkComponent
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	ret1 := C.atk_component_get_alpha(this1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *ComponentImpl) GetExtents(x0 *int, y0 *int, width0 *int, height0 *int, coord_type0 CoordType) {
	var this1 *C.AtkComponent
	var x1 *C.int32_t
	var y1 *C.int32_t
	var width1 *C.int32_t
	var height1 *C.int32_t
	var coord_type1 C.AtkCoordType
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	x1 = (*C.int32_t)(unsafe.Pointer(x0))
	y1 = (*C.int32_t)(unsafe.Pointer(y0))
	width1 = (*C.int32_t)(unsafe.Pointer(width0))
	height1 = (*C.int32_t)(unsafe.Pointer(height0))
	coord_type1 = C.AtkCoordType(coord_type0)
	C.atk_component_get_extents(this1, x1, y1, width1, height1, coord_type1)
}
func (this0 *ComponentImpl) GetLayer() Layer {
	var this1 *C.AtkComponent
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	ret1 := C.atk_component_get_layer(this1)
	var ret2 Layer
	ret2 = Layer(ret1)
	return ret2
}
func (this0 *ComponentImpl) GetMdiZorder() int {
	var this1 *C.AtkComponent
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	ret1 := C.atk_component_get_mdi_zorder(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *ComponentImpl) GetPosition(x0 *int, y0 *int, coord_type0 CoordType) {
	var this1 *C.AtkComponent
	var x1 *C.int32_t
	var y1 *C.int32_t
	var coord_type1 C.AtkCoordType
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	x1 = (*C.int32_t)(unsafe.Pointer(x0))
	y1 = (*C.int32_t)(unsafe.Pointer(y0))
	coord_type1 = C.AtkCoordType(coord_type0)
	C.atk_component_get_position(this1, x1, y1, coord_type1)
}
func (this0 *ComponentImpl) GetSize(width0 *int, height0 *int) {
	var this1 *C.AtkComponent
	var width1 *C.int32_t
	var height1 *C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	width1 = (*C.int32_t)(unsafe.Pointer(width0))
	height1 = (*C.int32_t)(unsafe.Pointer(height0))
	C.atk_component_get_size(this1, width1, height1)
}
func (this0 *ComponentImpl) GrabFocus() bool {
	var this1 *C.AtkComponent
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	ret1 := C.atk_component_grab_focus(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *ComponentImpl) RefAccessibleAtPoint(x0 int, y0 int, coord_type0 CoordType) *Object {
	var this1 *C.AtkComponent
	var x1 C.int32_t
	var y1 C.int32_t
	var coord_type1 C.AtkCoordType
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	coord_type1 = C.AtkCoordType(coord_type0)
	ret1 := C.atk_component_ref_accessible_at_point(this1, x1, y1, coord_type1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *ComponentImpl) RemoveFocusHandler(handler_id0 int) {
	var this1 *C.AtkComponent
	var handler_id1 C.uint32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	handler_id1 = C.uint32_t(handler_id0)
	C.atk_component_remove_focus_handler(this1, handler_id1)
}
func (this0 *ComponentImpl) SetExtents(x0 int, y0 int, width0 int, height0 int, coord_type0 CoordType) bool {
	var this1 *C.AtkComponent
	var x1 C.int32_t
	var y1 C.int32_t
	var width1 C.int32_t
	var height1 C.int32_t
	var coord_type1 C.AtkCoordType
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	coord_type1 = C.AtkCoordType(coord_type0)
	ret1 := C.atk_component_set_extents(this1, x1, y1, width1, height1, coord_type1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *ComponentImpl) SetPosition(x0 int, y0 int, coord_type0 CoordType) bool {
	var this1 *C.AtkComponent
	var x1 C.int32_t
	var y1 C.int32_t
	var coord_type1 C.AtkCoordType
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	coord_type1 = C.AtkCoordType(coord_type0)
	ret1 := C.atk_component_set_position(this1, x1, y1, coord_type1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *ComponentImpl) SetSize(width0 int, height0 int) bool {
	var this1 *C.AtkComponent
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkComponent()}
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	ret1 := C.atk_component_set_size(this1, width1, height1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type CoordType C.uint32_t
const (
	CoordTypeScreen CoordType = 0
	CoordTypeWindow CoordType = 1
)
type DocumentLike interface {
	ImplementsAtkDocument() *C.AtkDocument
}

type Document struct {
	gobject.Object
	DocumentImpl
}

type DocumentImpl struct {}

func ToDocument(objlike gobject.ObjectLike) *Document {
	t := (*DocumentImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Document)(obj)
	}
	panic("cannot cast to Document")
}

func (this0 *DocumentImpl) ImplementsAtkDocument() *C.AtkDocument {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkDocument)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *DocumentImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_document_get_type())
}

func DocumentGetType() gobject.Type {
	return (*DocumentImpl)(nil).GetStaticType()
}
func (this0 *DocumentImpl) GetAttributeValue(attribute_name0 string) string {
	var this1 *C.AtkDocument
	var attribute_name1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsAtkDocument()}
	attribute_name1 = _GoStringToGString(attribute_name0)
	defer C.free(unsafe.Pointer(attribute_name1))
	ret1 := C.atk_document_get_attribute_value(this1, attribute_name1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
// blacklisted: Document.get_attributes (method)
func (this0 *DocumentImpl) GetDocument() {
	var this1 *C.AtkDocument
	if this0 != nil {
		this1 = this0.ImplementsAtkDocument()}
	C.atk_document_get_document(this1)
}
func (this0 *DocumentImpl) GetDocumentType() string {
	var this1 *C.AtkDocument
	if this0 != nil {
		this1 = this0.ImplementsAtkDocument()}
	ret1 := C.atk_document_get_document_type(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *DocumentImpl) GetLocale() string {
	var this1 *C.AtkDocument
	if this0 != nil {
		this1 = this0.ImplementsAtkDocument()}
	ret1 := C.atk_document_get_locale(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *DocumentImpl) SetAttributeValue(attribute_name0 string, attribute_value0 string) bool {
	var this1 *C.AtkDocument
	var attribute_name1 *C.char
	var attribute_value1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsAtkDocument()}
	attribute_name1 = _GoStringToGString(attribute_name0)
	defer C.free(unsafe.Pointer(attribute_name1))
	attribute_value1 = _GoStringToGString(attribute_value0)
	defer C.free(unsafe.Pointer(attribute_value1))
	ret1 := C.atk_document_set_attribute_value(this1, attribute_name1, attribute_value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type EditableTextLike interface {
	ImplementsAtkEditableText() *C.AtkEditableText
}

type EditableText struct {
	gobject.Object
	EditableTextImpl
}

type EditableTextImpl struct {}

func ToEditableText(objlike gobject.ObjectLike) *EditableText {
	t := (*EditableTextImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*EditableText)(obj)
	}
	panic("cannot cast to EditableText")
}

func (this0 *EditableTextImpl) ImplementsAtkEditableText() *C.AtkEditableText {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkEditableText)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *EditableTextImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_editable_text_get_type())
}

func EditableTextGetType() gobject.Type {
	return (*EditableTextImpl)(nil).GetStaticType()
}
func (this0 *EditableTextImpl) CopyText(start_pos0 int, end_pos0 int) {
	var this1 *C.AtkEditableText
	var start_pos1 C.int32_t
	var end_pos1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkEditableText()}
	start_pos1 = C.int32_t(start_pos0)
	end_pos1 = C.int32_t(end_pos0)
	C.atk_editable_text_copy_text(this1, start_pos1, end_pos1)
}
func (this0 *EditableTextImpl) CutText(start_pos0 int, end_pos0 int) {
	var this1 *C.AtkEditableText
	var start_pos1 C.int32_t
	var end_pos1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkEditableText()}
	start_pos1 = C.int32_t(start_pos0)
	end_pos1 = C.int32_t(end_pos0)
	C.atk_editable_text_cut_text(this1, start_pos1, end_pos1)
}
func (this0 *EditableTextImpl) DeleteText(start_pos0 int, end_pos0 int) {
	var this1 *C.AtkEditableText
	var start_pos1 C.int32_t
	var end_pos1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkEditableText()}
	start_pos1 = C.int32_t(start_pos0)
	end_pos1 = C.int32_t(end_pos0)
	C.atk_editable_text_delete_text(this1, start_pos1, end_pos1)
}
func (this0 *EditableTextImpl) InsertText(string0 string, length0 int, position0 *int) {
	var this1 *C.AtkEditableText
	var string1 *C.char
	var length1 C.int32_t
	var position1 *C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkEditableText()}
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	length1 = C.int32_t(length0)
	position1 = (*C.int32_t)(unsafe.Pointer(position0))
	C.atk_editable_text_insert_text(this1, string1, length1, position1)
}
func (this0 *EditableTextImpl) PasteText(position0 int) {
	var this1 *C.AtkEditableText
	var position1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkEditableText()}
	position1 = C.int32_t(position0)
	C.atk_editable_text_paste_text(this1, position1)
}
func (this0 *EditableTextImpl) SetRunAttributes(attrib_set0 []unsafe.Pointer, start_offset0 int, end_offset0 int) bool {
	var this1 *C.AtkEditableText
	var attrib_set1 *C.GSList
	var start_offset1 C.int32_t
	var end_offset1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkEditableText()}
	start_offset1 = C.int32_t(start_offset0)
	end_offset1 = C.int32_t(end_offset0)
	ret1 := C.atk_editable_text_set_run_attributes(this1, attrib_set1, start_offset1, end_offset1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *EditableTextImpl) SetTextContents(string0 string) {
	var this1 *C.AtkEditableText
	var string1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsAtkEditableText()}
	string1 = _GoStringToGString(string0)
	defer C.free(unsafe.Pointer(string1))
	C.atk_editable_text_set_text_contents(this1, string1)
}
// blacklisted (no userdata): type EventListener func(obj *Object)
// blacklisted (no userdata): type EventListenerInit func()
// blacklisted (no userdata): type FocusHandler func(unknown *Object, unknown bool)
type Function func() bool
//export _AtkFunction_c_wrapper
func _AtkFunction_c_wrapper(data0 unsafe.Pointer) int32 {
	var data1 Function
	data1 = *(*Function)(data0)
	ret1 := data1()
	var ret2 C.int
	ret2 = _GoBoolToCBool(ret1)
	return (int32)(ret2)
}
//export _AtkFunction_c_wrapper_once
func _AtkFunction_c_wrapper_once(data0 unsafe.Pointer) int32 {
	ret := _AtkFunction_c_wrapper(data0)
	gobject.Holder.Release(data0)
	return ret
}
type GObjectAccessibleLike interface {
	ObjectLike
	InheritedFromAtkGObjectAccessible() *C.AtkGObjectAccessible
}

type GObjectAccessible struct {
	Object
	
}

func ToGObjectAccessible(objlike gobject.ObjectLike) *GObjectAccessible {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*GObjectAccessible)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*GObjectAccessible)(obj)
	}
	panic("cannot cast to GObjectAccessible")
}

func (this0 *GObjectAccessible) InheritedFromAtkGObjectAccessible() *C.AtkGObjectAccessible {
	if this0 == nil {
		return nil
	}
	return (*C.AtkGObjectAccessible)(this0.C)
}

func (this0 *GObjectAccessible) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_gobject_accessible_get_type())
}

func GObjectAccessibleGetType() gobject.Type {
	return (*GObjectAccessible)(nil).GetStaticType()
}
func GObjectAccessibleForObject(obj0 gobject.ObjectLike) *Object {
	var obj1 *C.GObject
	if obj0 != nil {
		obj1 = obj0.InheritedFromGObject()
	}
	ret1 := C.atk_gobject_accessible_for_object(obj1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *GObjectAccessible) GetObject() *gobject.Object {
	var this1 *C.AtkGObjectAccessible
	if this0 != nil {
		this1 = this0.InheritedFromAtkGObjectAccessible()
	}
	ret1 := C.atk_gobject_accessible_get_object(this1)
	var ret2 *gobject.Object
	ret2 = (*gobject.Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
type HyperlinkLike interface {
	gobject.ObjectLike
	InheritedFromAtkHyperlink() *C.AtkHyperlink
}

type Hyperlink struct {
	gobject.Object
	ActionImpl
}

func ToHyperlink(objlike gobject.ObjectLike) *Hyperlink {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Hyperlink)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Hyperlink)(obj)
	}
	panic("cannot cast to Hyperlink")
}

func (this0 *Hyperlink) InheritedFromAtkHyperlink() *C.AtkHyperlink {
	if this0 == nil {
		return nil
	}
	return (*C.AtkHyperlink)(this0.C)
}

func (this0 *Hyperlink) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_hyperlink_get_type())
}

func HyperlinkGetType() gobject.Type {
	return (*Hyperlink)(nil).GetStaticType()
}
func (this0 *Hyperlink) GetEndIndex() int {
	var this1 *C.AtkHyperlink
	if this0 != nil {
		this1 = this0.InheritedFromAtkHyperlink()
	}
	ret1 := C.atk_hyperlink_get_end_index(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Hyperlink) GetNAnchors() int {
	var this1 *C.AtkHyperlink
	if this0 != nil {
		this1 = this0.InheritedFromAtkHyperlink()
	}
	ret1 := C.atk_hyperlink_get_n_anchors(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Hyperlink) GetObject(i0 int) *Object {
	var this1 *C.AtkHyperlink
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromAtkHyperlink()
	}
	i1 = C.int32_t(i0)
	ret1 := C.atk_hyperlink_get_object(this1, i1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Hyperlink) GetStartIndex() int {
	var this1 *C.AtkHyperlink
	if this0 != nil {
		this1 = this0.InheritedFromAtkHyperlink()
	}
	ret1 := C.atk_hyperlink_get_start_index(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Hyperlink) GetURI(i0 int) string {
	var this1 *C.AtkHyperlink
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromAtkHyperlink()
	}
	i1 = C.int32_t(i0)
	ret1 := C.atk_hyperlink_get_uri(this1, i1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Hyperlink) IsInline() bool {
	var this1 *C.AtkHyperlink
	if this0 != nil {
		this1 = this0.InheritedFromAtkHyperlink()
	}
	ret1 := C.atk_hyperlink_is_inline(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Hyperlink) IsValid() bool {
	var this1 *C.AtkHyperlink
	if this0 != nil {
		this1 = this0.InheritedFromAtkHyperlink()
	}
	ret1 := C.atk_hyperlink_is_valid(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type HyperlinkImplLike interface {
	ImplementsAtkHyperlinkImpl() *C.AtkHyperlinkImpl
}

type HyperlinkImpl struct {
	gobject.Object
	HyperlinkImplImpl
}

type HyperlinkImplImpl struct {}

func ToHyperlinkImpl(objlike gobject.ObjectLike) *HyperlinkImpl {
	t := (*HyperlinkImplImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*HyperlinkImpl)(obj)
	}
	panic("cannot cast to HyperlinkImpl")
}

func (this0 *HyperlinkImplImpl) ImplementsAtkHyperlinkImpl() *C.AtkHyperlinkImpl {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkHyperlinkImpl)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *HyperlinkImplImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_hyperlink_impl_get_type())
}

func HyperlinkImplGetType() gobject.Type {
	return (*HyperlinkImplImpl)(nil).GetStaticType()
}
func (this0 *HyperlinkImplImpl) GetHyperlink() *Hyperlink {
	var this1 *C.AtkHyperlinkImpl
	if this0 != nil {
		this1 = this0.ImplementsAtkHyperlinkImpl()}
	ret1 := C.atk_hyperlink_impl_get_hyperlink(this1)
	var ret2 *Hyperlink
	ret2 = (*Hyperlink)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
type HyperlinkStateFlags C.uint32_t
const (
	HyperlinkStateFlagsInline HyperlinkStateFlags = 1
)
type HypertextLike interface {
	ImplementsAtkHypertext() *C.AtkHypertext
}

type Hypertext struct {
	gobject.Object
	HypertextImpl
}

type HypertextImpl struct {}

func ToHypertext(objlike gobject.ObjectLike) *Hypertext {
	t := (*HypertextImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Hypertext)(obj)
	}
	panic("cannot cast to Hypertext")
}

func (this0 *HypertextImpl) ImplementsAtkHypertext() *C.AtkHypertext {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkHypertext)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *HypertextImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_hypertext_get_type())
}

func HypertextGetType() gobject.Type {
	return (*HypertextImpl)(nil).GetStaticType()
}
func (this0 *HypertextImpl) GetLink(link_index0 int) *Hyperlink {
	var this1 *C.AtkHypertext
	var link_index1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkHypertext()}
	link_index1 = C.int32_t(link_index0)
	ret1 := C.atk_hypertext_get_link(this1, link_index1)
	var ret2 *Hyperlink
	ret2 = (*Hyperlink)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *HypertextImpl) GetLinkIndex(char_index0 int) int {
	var this1 *C.AtkHypertext
	var char_index1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkHypertext()}
	char_index1 = C.int32_t(char_index0)
	ret1 := C.atk_hypertext_get_link_index(this1, char_index1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *HypertextImpl) GetNLinks() int {
	var this1 *C.AtkHypertext
	if this0 != nil {
		this1 = this0.ImplementsAtkHypertext()}
	ret1 := C.atk_hypertext_get_n_links(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
type ImageLike interface {
	ImplementsAtkImage() *C.AtkImage
}

type Image struct {
	gobject.Object
	ImageImpl
}

type ImageImpl struct {}

func ToImage(objlike gobject.ObjectLike) *Image {
	t := (*ImageImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Image)(obj)
	}
	panic("cannot cast to Image")
}

func (this0 *ImageImpl) ImplementsAtkImage() *C.AtkImage {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkImage)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *ImageImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_image_get_type())
}

func ImageGetType() gobject.Type {
	return (*ImageImpl)(nil).GetStaticType()
}
func (this0 *ImageImpl) GetImageDescription() string {
	var this1 *C.AtkImage
	if this0 != nil {
		this1 = this0.ImplementsAtkImage()}
	ret1 := C.atk_image_get_image_description(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *ImageImpl) GetImageLocale() string {
	var this1 *C.AtkImage
	if this0 != nil {
		this1 = this0.ImplementsAtkImage()}
	ret1 := C.atk_image_get_image_locale(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *ImageImpl) GetImagePosition(x0 *int, y0 *int, coord_type0 CoordType) {
	var this1 *C.AtkImage
	var x1 *C.int32_t
	var y1 *C.int32_t
	var coord_type1 C.AtkCoordType
	if this0 != nil {
		this1 = this0.ImplementsAtkImage()}
	x1 = (*C.int32_t)(unsafe.Pointer(x0))
	y1 = (*C.int32_t)(unsafe.Pointer(y0))
	coord_type1 = C.AtkCoordType(coord_type0)
	C.atk_image_get_image_position(this1, x1, y1, coord_type1)
}
func (this0 *ImageImpl) GetImageSize(width0 *int, height0 *int) {
	var this1 *C.AtkImage
	var width1 *C.int32_t
	var height1 *C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkImage()}
	width1 = (*C.int32_t)(unsafe.Pointer(width0))
	height1 = (*C.int32_t)(unsafe.Pointer(height0))
	C.atk_image_get_image_size(this1, width1, height1)
}
func (this0 *ImageImpl) SetImageDescription(description0 string) bool {
	var this1 *C.AtkImage
	var description1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsAtkImage()}
	description1 = _GoStringToGString(description0)
	defer C.free(unsafe.Pointer(description1))
	ret1 := C.atk_image_set_image_description(this1, description1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type Implementor struct {}
func (this0 *Implementor) RefAccessible() *Object {
	var this1 *C.AtkImplementor
	this1 = (*C.AtkImplementor)(unsafe.Pointer(this0))
	ret1 := C.atk_implementor_ref_accessible(this1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
type ImplementorIfaceLike interface {
	ImplementsAtkImplementorIface() *C.AtkImplementorIface
}

type ImplementorIface struct {
	gobject.Object
	ImplementorIfaceImpl
}

type ImplementorIfaceImpl struct {}

func ToImplementorIface(objlike gobject.ObjectLike) *ImplementorIface {
	t := (*ImplementorIfaceImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ImplementorIface)(obj)
	}
	panic("cannot cast to ImplementorIface")
}

func (this0 *ImplementorIfaceImpl) ImplementsAtkImplementorIface() *C.AtkImplementorIface {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkImplementorIface)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *ImplementorIfaceImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_implementor_get_type())
}

func ImplementorIfaceGetType() gobject.Type {
	return (*ImplementorIfaceImpl)(nil).GetStaticType()
}
type KeyEventStruct struct {
	Type int32
	State uint32
	Keyval uint32
	Length int32
	string0 *C.char
	Keycode uint16
	_ [2]byte
	Timestamp uint32
}
func (this0 *KeyEventStruct) String() string {
	var string1 string
	string1 = C.GoString(this0.string0)
	return string1
}
type KeyEventType C.uint32_t
const (
	KeyEventTypePress KeyEventType = 0
	KeyEventTypeRelease KeyEventType = 1
	KeyEventTypeLastDefined KeyEventType = 2
)
type KeySnoopFunc func(event *KeyEventStruct) int
//export _AtkKeySnoopFunc_c_wrapper
func _AtkKeySnoopFunc_c_wrapper(event0 unsafe.Pointer, func_data0 unsafe.Pointer) int32 {
	var event1 *KeyEventStruct
	var func_data1 KeySnoopFunc
	event1 = (*KeyEventStruct)(unsafe.Pointer((*C.AtkKeyEventStruct)(event0)))
	func_data1 = *(*KeySnoopFunc)(func_data0)
	ret1 := func_data1(event1)
	var ret2 C.int32_t
	ret2 = C.int32_t(ret1)
	return (int32)(ret2)
}
//export _AtkKeySnoopFunc_c_wrapper_once
func _AtkKeySnoopFunc_c_wrapper_once(event0 unsafe.Pointer, func_data0 unsafe.Pointer) int32 {
	ret := _AtkKeySnoopFunc_c_wrapper(event0, func_data0)
	gobject.Holder.Release(func_data0)
	return ret
}
type Layer C.uint32_t
const (
	LayerInvalid Layer = 0
	LayerBackground Layer = 1
	LayerCanvas Layer = 2
	LayerWidget Layer = 3
	LayerMdi Layer = 4
	LayerPopup Layer = 5
	LayerOverlay Layer = 6
	LayerWindow Layer = 7
)
type MiscLike interface {
	gobject.ObjectLike
	InheritedFromAtkMisc() *C.AtkMisc
}

type Misc struct {
	gobject.Object
	
}

func ToMisc(objlike gobject.ObjectLike) *Misc {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Misc)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Misc)(obj)
	}
	panic("cannot cast to Misc")
}

func (this0 *Misc) InheritedFromAtkMisc() *C.AtkMisc {
	if this0 == nil {
		return nil
	}
	return (*C.AtkMisc)(this0.C)
}

func (this0 *Misc) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_misc_get_type())
}

func MiscGetType() gobject.Type {
	return (*Misc)(nil).GetStaticType()
}
func MiscGetInstance() *Misc {
	ret1 := C.atk_misc_get_instance()
	var ret2 *Misc
	ret2 = (*Misc)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Misc) ThreadsEnter() {
	var this1 *C.AtkMisc
	if this0 != nil {
		this1 = this0.InheritedFromAtkMisc()
	}
	C.atk_misc_threads_enter(this1)
}
func (this0 *Misc) ThreadsLeave() {
	var this1 *C.AtkMisc
	if this0 != nil {
		this1 = this0.InheritedFromAtkMisc()
	}
	C.atk_misc_threads_leave(this1)
}
type NoOpObjectLike interface {
	ObjectLike
	InheritedFromAtkNoOpObject() *C.AtkNoOpObject
}

type NoOpObject struct {
	Object
	ActionImpl
	ComponentImpl
	DocumentImpl
	EditableTextImpl
	HypertextImpl
	ImageImpl
	SelectionImpl
	TableImpl
	TextImpl
	ValueImpl
	WindowImpl
}

func ToNoOpObject(objlike gobject.ObjectLike) *NoOpObject {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*NoOpObject)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*NoOpObject)(obj)
	}
	panic("cannot cast to NoOpObject")
}

func (this0 *NoOpObject) InheritedFromAtkNoOpObject() *C.AtkNoOpObject {
	if this0 == nil {
		return nil
	}
	return (*C.AtkNoOpObject)(this0.C)
}

func (this0 *NoOpObject) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_no_op_object_get_type())
}

func NoOpObjectGetType() gobject.Type {
	return (*NoOpObject)(nil).GetStaticType()
}
func NewNoOpObject(obj0 gobject.ObjectLike) *NoOpObject {
	var obj1 *C.GObject
	if obj0 != nil {
		obj1 = obj0.InheritedFromGObject()
	}
	ret1 := C.atk_no_op_object_new(obj1)
	var ret2 *NoOpObject
	ret2 = (*NoOpObject)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
type NoOpObjectFactoryLike interface {
	ObjectFactoryLike
	InheritedFromAtkNoOpObjectFactory() *C.AtkNoOpObjectFactory
}

type NoOpObjectFactory struct {
	ObjectFactory
	
}

func ToNoOpObjectFactory(objlike gobject.ObjectLike) *NoOpObjectFactory {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*NoOpObjectFactory)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*NoOpObjectFactory)(obj)
	}
	panic("cannot cast to NoOpObjectFactory")
}

func (this0 *NoOpObjectFactory) InheritedFromAtkNoOpObjectFactory() *C.AtkNoOpObjectFactory {
	if this0 == nil {
		return nil
	}
	return (*C.AtkNoOpObjectFactory)(this0.C)
}

func (this0 *NoOpObjectFactory) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_no_op_object_factory_get_type())
}

func NoOpObjectFactoryGetType() gobject.Type {
	return (*NoOpObjectFactory)(nil).GetStaticType()
}
func NewNoOpObjectFactory() *NoOpObjectFactory {
	ret1 := C.atk_no_op_object_factory_new()
	var ret2 *NoOpObjectFactory
	ret2 = (*NoOpObjectFactory)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
type ObjectLike interface {
	gobject.ObjectLike
	InheritedFromAtkObject() *C.AtkObject
}

type Object struct {
	gobject.Object
	
}

func ToObject(objlike gobject.ObjectLike) *Object {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Object)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Object)(obj)
	}
	panic("cannot cast to Object")
}

func (this0 *Object) InheritedFromAtkObject() *C.AtkObject {
	if this0 == nil {
		return nil
	}
	return (*C.AtkObject)(this0.C)
}

func (this0 *Object) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_object_get_type())
}

func ObjectGetType() gobject.Type {
	return (*Object)(nil).GetStaticType()
}
func (this0 *Object) AddRelationship(relationship0 RelationType, target0 ObjectLike) bool {
	var this1 *C.AtkObject
	var relationship1 C.AtkRelationType
	var target1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	relationship1 = C.AtkRelationType(relationship0)
	if target0 != nil {
		target1 = target0.InheritedFromAtkObject()
	}
	ret1 := C.atk_object_add_relationship(this1, relationship1, target1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: Object.get_attributes (method)
func (this0 *Object) GetDescription() string {
	var this1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	ret1 := C.atk_object_get_description(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Object) GetIndexInParent() int {
	var this1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	ret1 := C.atk_object_get_index_in_parent(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Object) GetNAccessibleChildren() int {
	var this1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	ret1 := C.atk_object_get_n_accessible_children(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Object) GetName() string {
	var this1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	ret1 := C.atk_object_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Object) GetParent() *Object {
	var this1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	ret1 := C.atk_object_get_parent(this1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Object) GetRole() Role {
	var this1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	ret1 := C.atk_object_get_role(this1)
	var ret2 Role
	ret2 = Role(ret1)
	return ret2
}
func (this0 *Object) Initialize(data0 unsafe.Pointer) {
	var this1 *C.AtkObject
	var data1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	data1 = unsafe.Pointer(data0)
	C.atk_object_initialize(this1, data1)
}
func (this0 *Object) NotifyStateChange(state0 uint64, value0 bool) {
	var this1 *C.AtkObject
	var state1 C.uint64_t
	var value1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	state1 = C.uint64_t(state0)
	value1 = _GoBoolToCBool(value0)
	C.atk_object_notify_state_change(this1, state1, value1)
}
func (this0 *Object) RefAccessibleChild(i0 int) *Object {
	var this1 *C.AtkObject
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	i1 = C.int32_t(i0)
	ret1 := C.atk_object_ref_accessible_child(this1, i1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Object) RefRelationSet() *RelationSet {
	var this1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	ret1 := C.atk_object_ref_relation_set(this1)
	var ret2 *RelationSet
	ret2 = (*RelationSet)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Object) RefStateSet() *StateSet {
	var this1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	ret1 := C.atk_object_ref_state_set(this1)
	var ret2 *StateSet
	ret2 = (*StateSet)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Object) RemovePropertyChangeHandler(handler_id0 int) {
	var this1 *C.AtkObject
	var handler_id1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	handler_id1 = C.uint32_t(handler_id0)
	C.atk_object_remove_property_change_handler(this1, handler_id1)
}
func (this0 *Object) RemoveRelationship(relationship0 RelationType, target0 ObjectLike) bool {
	var this1 *C.AtkObject
	var relationship1 C.AtkRelationType
	var target1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	relationship1 = C.AtkRelationType(relationship0)
	if target0 != nil {
		target1 = target0.InheritedFromAtkObject()
	}
	ret1 := C.atk_object_remove_relationship(this1, relationship1, target1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Object) SetDescription(description0 string) {
	var this1 *C.AtkObject
	var description1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	description1 = _GoStringToGString(description0)
	defer C.free(unsafe.Pointer(description1))
	C.atk_object_set_description(this1, description1)
}
func (this0 *Object) SetName(name0 string) {
	var this1 *C.AtkObject
	var name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	C.atk_object_set_name(this1, name1)
}
func (this0 *Object) SetParent(parent0 ObjectLike) {
	var this1 *C.AtkObject
	var parent1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	if parent0 != nil {
		parent1 = parent0.InheritedFromAtkObject()
	}
	C.atk_object_set_parent(this1, parent1)
}
func (this0 *Object) SetRole(role0 Role) {
	var this1 *C.AtkObject
	var role1 C.AtkRole
	if this0 != nil {
		this1 = this0.InheritedFromAtkObject()
	}
	role1 = C.AtkRole(role0)
	C.atk_object_set_role(this1, role1)
}
type ObjectFactoryLike interface {
	gobject.ObjectLike
	InheritedFromAtkObjectFactory() *C.AtkObjectFactory
}

type ObjectFactory struct {
	gobject.Object
	
}

func ToObjectFactory(objlike gobject.ObjectLike) *ObjectFactory {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*ObjectFactory)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*ObjectFactory)(obj)
	}
	panic("cannot cast to ObjectFactory")
}

func (this0 *ObjectFactory) InheritedFromAtkObjectFactory() *C.AtkObjectFactory {
	if this0 == nil {
		return nil
	}
	return (*C.AtkObjectFactory)(this0.C)
}

func (this0 *ObjectFactory) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_object_factory_get_type())
}

func ObjectFactoryGetType() gobject.Type {
	return (*ObjectFactory)(nil).GetStaticType()
}
func (this0 *ObjectFactory) CreateAccessible(obj0 gobject.ObjectLike) *Object {
	var this1 *C.AtkObjectFactory
	var obj1 *C.GObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkObjectFactory()
	}
	if obj0 != nil {
		obj1 = obj0.InheritedFromGObject()
	}
	ret1 := C.atk_object_factory_create_accessible(this1, obj1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *ObjectFactory) GetAccessibleType() gobject.Type {
	var this1 *C.AtkObjectFactory
	if this0 != nil {
		this1 = this0.InheritedFromAtkObjectFactory()
	}
	ret1 := C.atk_object_factory_get_accessible_type(this1)
	var ret2 gobject.Type
	ret2 = gobject.Type(ret1)
	return ret2
}
func (this0 *ObjectFactory) Invalidate() {
	var this1 *C.AtkObjectFactory
	if this0 != nil {
		this1 = this0.InheritedFromAtkObjectFactory()
	}
	C.atk_object_factory_invalidate(this1)
}
type PlugLike interface {
	ObjectLike
	InheritedFromAtkPlug() *C.AtkPlug
}

type Plug struct {
	Object
	ComponentImpl
}

func ToPlug(objlike gobject.ObjectLike) *Plug {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Plug)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Plug)(obj)
	}
	panic("cannot cast to Plug")
}

func (this0 *Plug) InheritedFromAtkPlug() *C.AtkPlug {
	if this0 == nil {
		return nil
	}
	return (*C.AtkPlug)(this0.C)
}

func (this0 *Plug) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_plug_get_type())
}

func PlugGetType() gobject.Type {
	return (*Plug)(nil).GetStaticType()
}
func NewPlug() *Plug {
	ret1 := C.atk_plug_new()
	var ret2 *Plug
	ret2 = (*Plug)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Plug) GetID() string {
	var this1 *C.AtkPlug
	if this0 != nil {
		this1 = this0.InheritedFromAtkPlug()
	}
	ret1 := C.atk_plug_get_id(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
type Rectangle struct {
	X int32
	Y int32
	Width int32
	Height int32
}
type RegistryLike interface {
	gobject.ObjectLike
	InheritedFromAtkRegistry() *C.AtkRegistry
}

type Registry struct {
	gobject.Object
	
}

func ToRegistry(objlike gobject.ObjectLike) *Registry {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Registry)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Registry)(obj)
	}
	panic("cannot cast to Registry")
}

func (this0 *Registry) InheritedFromAtkRegistry() *C.AtkRegistry {
	if this0 == nil {
		return nil
	}
	return (*C.AtkRegistry)(this0.C)
}

func (this0 *Registry) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_registry_get_type())
}

func RegistryGetType() gobject.Type {
	return (*Registry)(nil).GetStaticType()
}
func (this0 *Registry) GetFactory(type0 gobject.Type) *ObjectFactory {
	var this1 *C.AtkRegistry
	var type1 C.GType
	if this0 != nil {
		this1 = this0.InheritedFromAtkRegistry()
	}
	type1 = C.GType(type0)
	ret1 := C.atk_registry_get_factory(this1, type1)
	var ret2 *ObjectFactory
	ret2 = (*ObjectFactory)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Registry) GetFactoryType(type0 gobject.Type) gobject.Type {
	var this1 *C.AtkRegistry
	var type1 C.GType
	if this0 != nil {
		this1 = this0.InheritedFromAtkRegistry()
	}
	type1 = C.GType(type0)
	ret1 := C.atk_registry_get_factory_type(this1, type1)
	var ret2 gobject.Type
	ret2 = gobject.Type(ret1)
	return ret2
}
func (this0 *Registry) SetFactoryType(type0 gobject.Type, factory_type0 gobject.Type) {
	var this1 *C.AtkRegistry
	var type1 C.GType
	var factory_type1 C.GType
	if this0 != nil {
		this1 = this0.InheritedFromAtkRegistry()
	}
	type1 = C.GType(type0)
	factory_type1 = C.GType(factory_type0)
	C.atk_registry_set_factory_type(this1, type1, factory_type1)
}
type RelationLike interface {
	gobject.ObjectLike
	InheritedFromAtkRelation() *C.AtkRelation
}

type Relation struct {
	gobject.Object
	
}

func ToRelation(objlike gobject.ObjectLike) *Relation {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Relation)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Relation)(obj)
	}
	panic("cannot cast to Relation")
}

func (this0 *Relation) InheritedFromAtkRelation() *C.AtkRelation {
	if this0 == nil {
		return nil
	}
	return (*C.AtkRelation)(this0.C)
}

func (this0 *Relation) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_relation_get_type())
}

func RelationGetType() gobject.Type {
	return (*Relation)(nil).GetStaticType()
}
func NewRelation(targets0 ObjectLike, n_targets0 int, relationship0 RelationType) *Relation {
	var targets1 *C.AtkObject
	var n_targets1 C.int32_t
	var relationship1 C.AtkRelationType
	if targets0 != nil {
		targets1 = targets0.InheritedFromAtkObject()
	}
	n_targets1 = C.int32_t(n_targets0)
	relationship1 = C.AtkRelationType(relationship0)
	ret1 := C.atk_relation_new(targets1, n_targets1, relationship1)
	var ret2 *Relation
	ret2 = (*Relation)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Relation) AddTarget(target0 ObjectLike) {
	var this1 *C.AtkRelation
	var target1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkRelation()
	}
	if target0 != nil {
		target1 = target0.InheritedFromAtkObject()
	}
	C.atk_relation_add_target(this1, target1)
}
func (this0 *Relation) GetRelationType() RelationType {
	var this1 *C.AtkRelation
	if this0 != nil {
		this1 = this0.InheritedFromAtkRelation()
	}
	ret1 := C.atk_relation_get_relation_type(this1)
	var ret2 RelationType
	ret2 = RelationType(ret1)
	return ret2
}
// blacklisted: Relation.get_target (method)
func (this0 *Relation) RemoveTarget(target0 ObjectLike) bool {
	var this1 *C.AtkRelation
	var target1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkRelation()
	}
	if target0 != nil {
		target1 = target0.InheritedFromAtkObject()
	}
	ret1 := C.atk_relation_remove_target(this1, target1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type RelationSetLike interface {
	gobject.ObjectLike
	InheritedFromAtkRelationSet() *C.AtkRelationSet
}

type RelationSet struct {
	gobject.Object
	
}

func ToRelationSet(objlike gobject.ObjectLike) *RelationSet {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*RelationSet)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*RelationSet)(obj)
	}
	panic("cannot cast to RelationSet")
}

func (this0 *RelationSet) InheritedFromAtkRelationSet() *C.AtkRelationSet {
	if this0 == nil {
		return nil
	}
	return (*C.AtkRelationSet)(this0.C)
}

func (this0 *RelationSet) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_relation_set_get_type())
}

func RelationSetGetType() gobject.Type {
	return (*RelationSet)(nil).GetStaticType()
}
func NewRelationSet() *RelationSet {
	ret1 := C.atk_relation_set_new()
	var ret2 *RelationSet
	ret2 = (*RelationSet)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *RelationSet) Add(relation0 RelationLike) {
	var this1 *C.AtkRelationSet
	var relation1 *C.AtkRelation
	if this0 != nil {
		this1 = this0.InheritedFromAtkRelationSet()
	}
	if relation0 != nil {
		relation1 = relation0.InheritedFromAtkRelation()
	}
	C.atk_relation_set_add(this1, relation1)
}
func (this0 *RelationSet) AddRelationByType(relationship0 RelationType, target0 ObjectLike) {
	var this1 *C.AtkRelationSet
	var relationship1 C.AtkRelationType
	var target1 *C.AtkObject
	if this0 != nil {
		this1 = this0.InheritedFromAtkRelationSet()
	}
	relationship1 = C.AtkRelationType(relationship0)
	if target0 != nil {
		target1 = target0.InheritedFromAtkObject()
	}
	C.atk_relation_set_add_relation_by_type(this1, relationship1, target1)
}
func (this0 *RelationSet) Contains(relationship0 RelationType) bool {
	var this1 *C.AtkRelationSet
	var relationship1 C.AtkRelationType
	if this0 != nil {
		this1 = this0.InheritedFromAtkRelationSet()
	}
	relationship1 = C.AtkRelationType(relationship0)
	ret1 := C.atk_relation_set_contains(this1, relationship1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *RelationSet) GetNRelations() int {
	var this1 *C.AtkRelationSet
	if this0 != nil {
		this1 = this0.InheritedFromAtkRelationSet()
	}
	ret1 := C.atk_relation_set_get_n_relations(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *RelationSet) GetRelation(i0 int) *Relation {
	var this1 *C.AtkRelationSet
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromAtkRelationSet()
	}
	i1 = C.int32_t(i0)
	ret1 := C.atk_relation_set_get_relation(this1, i1)
	var ret2 *Relation
	ret2 = (*Relation)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *RelationSet) GetRelationByType(relationship0 RelationType) *Relation {
	var this1 *C.AtkRelationSet
	var relationship1 C.AtkRelationType
	if this0 != nil {
		this1 = this0.InheritedFromAtkRelationSet()
	}
	relationship1 = C.AtkRelationType(relationship0)
	ret1 := C.atk_relation_set_get_relation_by_type(this1, relationship1)
	var ret2 *Relation
	ret2 = (*Relation)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *RelationSet) Remove(relation0 RelationLike) {
	var this1 *C.AtkRelationSet
	var relation1 *C.AtkRelation
	if this0 != nil {
		this1 = this0.InheritedFromAtkRelationSet()
	}
	if relation0 != nil {
		relation1 = relation0.InheritedFromAtkRelation()
	}
	C.atk_relation_set_remove(this1, relation1)
}
type RelationType C.uint32_t
const (
	RelationTypeNull RelationType = 0
	RelationTypeControlledBy RelationType = 1
	RelationTypeControllerFor RelationType = 2
	RelationTypeLabelFor RelationType = 3
	RelationTypeLabelledBy RelationType = 4
	RelationTypeMemberOf RelationType = 5
	RelationTypeNodeChildOf RelationType = 6
	RelationTypeFlowsTo RelationType = 7
	RelationTypeFlowsFrom RelationType = 8
	RelationTypeSubwindowOf RelationType = 9
	RelationTypeEmbeds RelationType = 10
	RelationTypeEmbeddedBy RelationType = 11
	RelationTypePopupFor RelationType = 12
	RelationTypeParentWindowOf RelationType = 13
	RelationTypeDescribedBy RelationType = 14
	RelationTypeDescriptionFor RelationType = 15
	RelationTypeNodeParentOf RelationType = 16
	RelationTypeLastDefined RelationType = 17
)
type Role C.uint32_t
const (
	RoleInvalid Role = 0
	RoleAccelLabel Role = 1
	RoleAlert Role = 2
	RoleAnimation Role = 3
	RoleArrow Role = 4
	RoleCalendar Role = 5
	RoleCanvas Role = 6
	RoleCheckBox Role = 7
	RoleCheckMenuItem Role = 8
	RoleColorChooser Role = 9
	RoleColumnHeader Role = 10
	RoleComboBox Role = 11
	RoleDateEditor Role = 12
	RoleDesktopIcon Role = 13
	RoleDesktopFrame Role = 14
	RoleDial Role = 15
	RoleDialog Role = 16
	RoleDirectoryPane Role = 17
	RoleDrawingArea Role = 18
	RoleFileChooser Role = 19
	RoleFiller Role = 20
	RoleFontChooser Role = 21
	RoleFrame Role = 22
	RoleGlassPane Role = 23
	RoleHtmlContainer Role = 24
	RoleIcon Role = 25
	RoleImage Role = 26
	RoleInternalFrame Role = 27
	RoleLabel Role = 28
	RoleLayeredPane Role = 29
	RoleList Role = 30
	RoleListItem Role = 31
	RoleMenu Role = 32
	RoleMenuBar Role = 33
	RoleMenuItem Role = 34
	RoleOptionPane Role = 35
	RolePageTab Role = 36
	RolePageTabList Role = 37
	RolePanel Role = 38
	RolePasswordText Role = 39
	RolePopupMenu Role = 40
	RoleProgressBar Role = 41
	RolePushButton Role = 42
	RoleRadioButton Role = 43
	RoleRadioMenuItem Role = 44
	RoleRootPane Role = 45
	RoleRowHeader Role = 46
	RoleScrollBar Role = 47
	RoleScrollPane Role = 48
	RoleSeparator Role = 49
	RoleSlider Role = 50
	RoleSplitPane Role = 51
	RoleSpinButton Role = 52
	RoleStatusbar Role = 53
	RoleTable Role = 54
	RoleTableCell Role = 55
	RoleTableColumnHeader Role = 56
	RoleTableRowHeader Role = 57
	RoleTearOffMenuItem Role = 58
	RoleTerminal Role = 59
	RoleText Role = 60
	RoleToggleButton Role = 61
	RoleToolBar Role = 62
	RoleToolTip Role = 63
	RoleTree Role = 64
	RoleTreeTable Role = 65
	RoleUnknown Role = 66
	RoleViewport Role = 67
	RoleWindow Role = 68
	RoleHeader Role = 69
	RoleFooter Role = 70
	RoleParagraph Role = 71
	RoleRuler Role = 72
	RoleApplication Role = 73
	RoleAutocomplete Role = 74
	RoleEditbar Role = 75
	RoleEmbedded Role = 76
	RoleEntry Role = 77
	RoleChart Role = 78
	RoleCaption Role = 79
	RoleDocumentFrame Role = 80
	RoleHeading Role = 81
	RolePage Role = 82
	RoleSection Role = 83
	RoleRedundantObject Role = 84
	RoleForm Role = 85
	RoleLink Role = 86
	RoleInputMethodWindow Role = 87
	RoleTableRow Role = 88
	RoleTreeItem Role = 89
	RoleDocumentSpreadsheet Role = 90
	RoleDocumentPresentation Role = 91
	RoleDocumentText Role = 92
	RoleDocumentWeb Role = 93
	RoleDocumentEmail Role = 94
	RoleComment Role = 95
	RoleListBox Role = 96
	RoleGrouping Role = 97
	RoleImageMap Role = 98
	RoleNotification Role = 99
	RoleInfoBar Role = 100
	RoleLastDefined Role = 101
)
type SelectionLike interface {
	ImplementsAtkSelection() *C.AtkSelection
}

type Selection struct {
	gobject.Object
	SelectionImpl
}

type SelectionImpl struct {}

func ToSelection(objlike gobject.ObjectLike) *Selection {
	t := (*SelectionImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Selection)(obj)
	}
	panic("cannot cast to Selection")
}

func (this0 *SelectionImpl) ImplementsAtkSelection() *C.AtkSelection {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkSelection)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *SelectionImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_selection_get_type())
}

func SelectionGetType() gobject.Type {
	return (*SelectionImpl)(nil).GetStaticType()
}
func (this0 *SelectionImpl) AddSelection(i0 int) bool {
	var this1 *C.AtkSelection
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkSelection()}
	i1 = C.int32_t(i0)
	ret1 := C.atk_selection_add_selection(this1, i1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *SelectionImpl) ClearSelection() bool {
	var this1 *C.AtkSelection
	if this0 != nil {
		this1 = this0.ImplementsAtkSelection()}
	ret1 := C.atk_selection_clear_selection(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *SelectionImpl) GetSelectionCount() int {
	var this1 *C.AtkSelection
	if this0 != nil {
		this1 = this0.ImplementsAtkSelection()}
	ret1 := C.atk_selection_get_selection_count(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *SelectionImpl) IsChildSelected(i0 int) bool {
	var this1 *C.AtkSelection
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkSelection()}
	i1 = C.int32_t(i0)
	ret1 := C.atk_selection_is_child_selected(this1, i1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *SelectionImpl) RefSelection(i0 int) *Object {
	var this1 *C.AtkSelection
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkSelection()}
	i1 = C.int32_t(i0)
	ret1 := C.atk_selection_ref_selection(this1, i1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *SelectionImpl) RemoveSelection(i0 int) bool {
	var this1 *C.AtkSelection
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkSelection()}
	i1 = C.int32_t(i0)
	ret1 := C.atk_selection_remove_selection(this1, i1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *SelectionImpl) SelectAllSelection() bool {
	var this1 *C.AtkSelection
	if this0 != nil {
		this1 = this0.ImplementsAtkSelection()}
	ret1 := C.atk_selection_select_all_selection(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type SocketLike interface {
	ObjectLike
	InheritedFromAtkSocket() *C.AtkSocket
}

type Socket struct {
	Object
	ComponentImpl
}

func ToSocket(objlike gobject.ObjectLike) *Socket {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Socket)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Socket)(obj)
	}
	panic("cannot cast to Socket")
}

func (this0 *Socket) InheritedFromAtkSocket() *C.AtkSocket {
	if this0 == nil {
		return nil
	}
	return (*C.AtkSocket)(this0.C)
}

func (this0 *Socket) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_socket_get_type())
}

func SocketGetType() gobject.Type {
	return (*Socket)(nil).GetStaticType()
}
func NewSocket() *Socket {
	ret1 := C.atk_socket_new()
	var ret2 *Socket
	ret2 = (*Socket)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Socket) Embed(plug_id0 string) {
	var this1 *C.AtkSocket
	var plug_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromAtkSocket()
	}
	plug_id1 = _GoStringToGString(plug_id0)
	defer C.free(unsafe.Pointer(plug_id1))
	C.atk_socket_embed(this1, plug_id1)
}
func (this0 *Socket) IsOccupied() bool {
	var this1 *C.AtkSocket
	if this0 != nil {
		this1 = this0.InheritedFromAtkSocket()
	}
	ret1 := C.atk_socket_is_occupied(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type StateSetLike interface {
	gobject.ObjectLike
	InheritedFromAtkStateSet() *C.AtkStateSet
}

type StateSet struct {
	gobject.Object
	
}

func ToStateSet(objlike gobject.ObjectLike) *StateSet {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*StateSet)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*StateSet)(obj)
	}
	panic("cannot cast to StateSet")
}

func (this0 *StateSet) InheritedFromAtkStateSet() *C.AtkStateSet {
	if this0 == nil {
		return nil
	}
	return (*C.AtkStateSet)(this0.C)
}

func (this0 *StateSet) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_state_set_get_type())
}

func StateSetGetType() gobject.Type {
	return (*StateSet)(nil).GetStaticType()
}
func NewStateSet() *StateSet {
	ret1 := C.atk_state_set_new()
	var ret2 *StateSet
	ret2 = (*StateSet)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *StateSet) AddState(type0 StateType) bool {
	var this1 *C.AtkStateSet
	var type1 C.AtkStateType
	if this0 != nil {
		this1 = this0.InheritedFromAtkStateSet()
	}
	type1 = C.AtkStateType(type0)
	ret1 := C.atk_state_set_add_state(this1, type1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: StateSet.add_states (method)
func (this0 *StateSet) AndSets(compare_set0 StateSetLike) *StateSet {
	var this1 *C.AtkStateSet
	var compare_set1 *C.AtkStateSet
	if this0 != nil {
		this1 = this0.InheritedFromAtkStateSet()
	}
	if compare_set0 != nil {
		compare_set1 = compare_set0.InheritedFromAtkStateSet()
	}
	ret1 := C.atk_state_set_and_sets(this1, compare_set1)
	var ret2 *StateSet
	ret2 = (*StateSet)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *StateSet) ClearStates() {
	var this1 *C.AtkStateSet
	if this0 != nil {
		this1 = this0.InheritedFromAtkStateSet()
	}
	C.atk_state_set_clear_states(this1)
}
func (this0 *StateSet) ContainsState(type0 StateType) bool {
	var this1 *C.AtkStateSet
	var type1 C.AtkStateType
	if this0 != nil {
		this1 = this0.InheritedFromAtkStateSet()
	}
	type1 = C.AtkStateType(type0)
	ret1 := C.atk_state_set_contains_state(this1, type1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: StateSet.contains_states (method)
func (this0 *StateSet) IsEmpty() bool {
	var this1 *C.AtkStateSet
	if this0 != nil {
		this1 = this0.InheritedFromAtkStateSet()
	}
	ret1 := C.atk_state_set_is_empty(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *StateSet) OrSets(compare_set0 StateSetLike) *StateSet {
	var this1 *C.AtkStateSet
	var compare_set1 *C.AtkStateSet
	if this0 != nil {
		this1 = this0.InheritedFromAtkStateSet()
	}
	if compare_set0 != nil {
		compare_set1 = compare_set0.InheritedFromAtkStateSet()
	}
	ret1 := C.atk_state_set_or_sets(this1, compare_set1)
	var ret2 *StateSet
	ret2 = (*StateSet)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *StateSet) RemoveState(type0 StateType) bool {
	var this1 *C.AtkStateSet
	var type1 C.AtkStateType
	if this0 != nil {
		this1 = this0.InheritedFromAtkStateSet()
	}
	type1 = C.AtkStateType(type0)
	ret1 := C.atk_state_set_remove_state(this1, type1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *StateSet) XorSets(compare_set0 StateSetLike) *StateSet {
	var this1 *C.AtkStateSet
	var compare_set1 *C.AtkStateSet
	if this0 != nil {
		this1 = this0.InheritedFromAtkStateSet()
	}
	if compare_set0 != nil {
		compare_set1 = compare_set0.InheritedFromAtkStateSet()
	}
	ret1 := C.atk_state_set_xor_sets(this1, compare_set1)
	var ret2 *StateSet
	ret2 = (*StateSet)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
type StateType C.uint32_t
const (
	StateTypeInvalid StateType = 0
	StateTypeActive StateType = 1
	StateTypeArmed StateType = 2
	StateTypeBusy StateType = 3
	StateTypeChecked StateType = 4
	StateTypeDefunct StateType = 5
	StateTypeEditable StateType = 6
	StateTypeEnabled StateType = 7
	StateTypeExpandable StateType = 8
	StateTypeExpanded StateType = 9
	StateTypeFocusable StateType = 10
	StateTypeFocused StateType = 11
	StateTypeHorizontal StateType = 12
	StateTypeIconified StateType = 13
	StateTypeModal StateType = 14
	StateTypeMultiLine StateType = 15
	StateTypeMultiselectable StateType = 16
	StateTypeOpaque StateType = 17
	StateTypePressed StateType = 18
	StateTypeResizable StateType = 19
	StateTypeSelectable StateType = 20
	StateTypeSelected StateType = 21
	StateTypeSensitive StateType = 22
	StateTypeShowing StateType = 23
	StateTypeSingleLine StateType = 24
	StateTypeStale StateType = 25
	StateTypeTransient StateType = 26
	StateTypeVertical StateType = 27
	StateTypeVisible StateType = 28
	StateTypeManagesDescendants StateType = 29
	StateTypeIndeterminate StateType = 30
	StateTypeTruncated StateType = 31
	StateTypeRequired StateType = 32
	StateTypeInvalidEntry StateType = 33
	StateTypeSupportsAutocompletion StateType = 34
	StateTypeSelectableText StateType = 35
	StateTypeDefault StateType = 36
	StateTypeAnimated StateType = 37
	StateTypeVisited StateType = 38
	StateTypeLastDefined StateType = 39
)
type StreamableContentLike interface {
	ImplementsAtkStreamableContent() *C.AtkStreamableContent
}

type StreamableContent struct {
	gobject.Object
	StreamableContentImpl
}

type StreamableContentImpl struct {}

func ToStreamableContent(objlike gobject.ObjectLike) *StreamableContent {
	t := (*StreamableContentImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*StreamableContent)(obj)
	}
	panic("cannot cast to StreamableContent")
}

func (this0 *StreamableContentImpl) ImplementsAtkStreamableContent() *C.AtkStreamableContent {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkStreamableContent)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *StreamableContentImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_streamable_content_get_type())
}

func StreamableContentGetType() gobject.Type {
	return (*StreamableContentImpl)(nil).GetStaticType()
}
func (this0 *StreamableContentImpl) GetMIMEType(i0 int) string {
	var this1 *C.AtkStreamableContent
	var i1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkStreamableContent()}
	i1 = C.int32_t(i0)
	ret1 := C.atk_streamable_content_get_mime_type(this1, i1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *StreamableContentImpl) GetNMIMETypes() int {
	var this1 *C.AtkStreamableContent
	if this0 != nil {
		this1 = this0.ImplementsAtkStreamableContent()}
	ret1 := C.atk_streamable_content_get_n_mime_types(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
// blacklisted: StreamableContent.get_stream (method)
func (this0 *StreamableContentImpl) GetURI(mime_type0 string) string {
	var this1 *C.AtkStreamableContent
	var mime_type1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsAtkStreamableContent()}
	mime_type1 = _GoStringToGString(mime_type0)
	defer C.free(unsafe.Pointer(mime_type1))
	ret1 := C.atk_streamable_content_get_uri(this1, mime_type1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
type TableLike interface {
	ImplementsAtkTable() *C.AtkTable
}

type Table struct {
	gobject.Object
	TableImpl
}

type TableImpl struct {}

func ToTable(objlike gobject.ObjectLike) *Table {
	t := (*TableImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Table)(obj)
	}
	panic("cannot cast to Table")
}

func (this0 *TableImpl) ImplementsAtkTable() *C.AtkTable {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkTable)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *TableImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_table_get_type())
}

func TableGetType() gobject.Type {
	return (*TableImpl)(nil).GetStaticType()
}
func (this0 *TableImpl) AddColumnSelection(column0 int) bool {
	var this1 *C.AtkTable
	var column1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	column1 = C.int32_t(column0)
	ret1 := C.atk_table_add_column_selection(this1, column1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TableImpl) AddRowSelection(row0 int) bool {
	var this1 *C.AtkTable
	var row1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	ret1 := C.atk_table_add_row_selection(this1, row1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TableImpl) GetCaption() *Object {
	var this1 *C.AtkTable
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	ret1 := C.atk_table_get_caption(this1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *TableImpl) GetColumnAtIndex(index_0 int) int {
	var this1 *C.AtkTable
	var index_1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	index_1 = C.int32_t(index_0)
	ret1 := C.atk_table_get_column_at_index(this1, index_1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TableImpl) GetColumnDescription(column0 int) string {
	var this1 *C.AtkTable
	var column1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	column1 = C.int32_t(column0)
	ret1 := C.atk_table_get_column_description(this1, column1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *TableImpl) GetColumnExtentAt(row0 int, column0 int) int {
	var this1 *C.AtkTable
	var row1 C.int32_t
	var column1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	column1 = C.int32_t(column0)
	ret1 := C.atk_table_get_column_extent_at(this1, row1, column1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TableImpl) GetColumnHeader(column0 int) *Object {
	var this1 *C.AtkTable
	var column1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	column1 = C.int32_t(column0)
	ret1 := C.atk_table_get_column_header(this1, column1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *TableImpl) GetIndexAt(row0 int, column0 int) int {
	var this1 *C.AtkTable
	var row1 C.int32_t
	var column1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	column1 = C.int32_t(column0)
	ret1 := C.atk_table_get_index_at(this1, row1, column1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TableImpl) GetNColumns() int {
	var this1 *C.AtkTable
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	ret1 := C.atk_table_get_n_columns(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TableImpl) GetNRows() int {
	var this1 *C.AtkTable
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	ret1 := C.atk_table_get_n_rows(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TableImpl) GetRowAtIndex(index_0 int) int {
	var this1 *C.AtkTable
	var index_1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	index_1 = C.int32_t(index_0)
	ret1 := C.atk_table_get_row_at_index(this1, index_1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TableImpl) GetRowDescription(row0 int) string {
	var this1 *C.AtkTable
	var row1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	ret1 := C.atk_table_get_row_description(this1, row1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *TableImpl) GetRowExtentAt(row0 int, column0 int) int {
	var this1 *C.AtkTable
	var row1 C.int32_t
	var column1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	column1 = C.int32_t(column0)
	ret1 := C.atk_table_get_row_extent_at(this1, row1, column1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TableImpl) GetRowHeader(row0 int) *Object {
	var this1 *C.AtkTable
	var row1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	ret1 := C.atk_table_get_row_header(this1, row1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *TableImpl) GetSelectedColumns(selected0 *int) int {
	var this1 *C.AtkTable
	var selected1 *C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	selected1 = (*C.int32_t)(unsafe.Pointer(selected0))
	ret1 := C.atk_table_get_selected_columns(this1, selected1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TableImpl) GetSelectedRows(selected0 *int) int {
	var this1 *C.AtkTable
	var selected1 *C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	selected1 = (*C.int32_t)(unsafe.Pointer(selected0))
	ret1 := C.atk_table_get_selected_rows(this1, selected1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TableImpl) GetSummary() *Object {
	var this1 *C.AtkTable
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	ret1 := C.atk_table_get_summary(this1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *TableImpl) IsColumnSelected(column0 int) bool {
	var this1 *C.AtkTable
	var column1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	column1 = C.int32_t(column0)
	ret1 := C.atk_table_is_column_selected(this1, column1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TableImpl) IsRowSelected(row0 int) bool {
	var this1 *C.AtkTable
	var row1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	ret1 := C.atk_table_is_row_selected(this1, row1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TableImpl) IsSelected(row0 int, column0 int) bool {
	var this1 *C.AtkTable
	var row1 C.int32_t
	var column1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	column1 = C.int32_t(column0)
	ret1 := C.atk_table_is_selected(this1, row1, column1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TableImpl) RefAt(row0 int, column0 int) *Object {
	var this1 *C.AtkTable
	var row1 C.int32_t
	var column1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	column1 = C.int32_t(column0)
	ret1 := C.atk_table_ref_at(this1, row1, column1)
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *TableImpl) RemoveColumnSelection(column0 int) bool {
	var this1 *C.AtkTable
	var column1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	column1 = C.int32_t(column0)
	ret1 := C.atk_table_remove_column_selection(this1, column1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TableImpl) RemoveRowSelection(row0 int) bool {
	var this1 *C.AtkTable
	var row1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	ret1 := C.atk_table_remove_row_selection(this1, row1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TableImpl) SetCaption(caption0 ObjectLike) {
	var this1 *C.AtkTable
	var caption1 *C.AtkObject
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	if caption0 != nil {
		caption1 = caption0.InheritedFromAtkObject()
	}
	C.atk_table_set_caption(this1, caption1)
}
func (this0 *TableImpl) SetColumnDescription(column0 int, description0 string) {
	var this1 *C.AtkTable
	var column1 C.int32_t
	var description1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	column1 = C.int32_t(column0)
	description1 = _GoStringToGString(description0)
	defer C.free(unsafe.Pointer(description1))
	C.atk_table_set_column_description(this1, column1, description1)
}
func (this0 *TableImpl) SetColumnHeader(column0 int, header0 ObjectLike) {
	var this1 *C.AtkTable
	var column1 C.int32_t
	var header1 *C.AtkObject
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	column1 = C.int32_t(column0)
	if header0 != nil {
		header1 = header0.InheritedFromAtkObject()
	}
	C.atk_table_set_column_header(this1, column1, header1)
}
func (this0 *TableImpl) SetRowDescription(row0 int, description0 string) {
	var this1 *C.AtkTable
	var row1 C.int32_t
	var description1 *C.char
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	description1 = _GoStringToGString(description0)
	defer C.free(unsafe.Pointer(description1))
	C.atk_table_set_row_description(this1, row1, description1)
}
func (this0 *TableImpl) SetRowHeader(row0 int, header0 ObjectLike) {
	var this1 *C.AtkTable
	var row1 C.int32_t
	var header1 *C.AtkObject
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	row1 = C.int32_t(row0)
	if header0 != nil {
		header1 = header0.InheritedFromAtkObject()
	}
	C.atk_table_set_row_header(this1, row1, header1)
}
func (this0 *TableImpl) SetSummary(accessible0 ObjectLike) {
	var this1 *C.AtkTable
	var accessible1 *C.AtkObject
	if this0 != nil {
		this1 = this0.ImplementsAtkTable()}
	if accessible0 != nil {
		accessible1 = accessible0.InheritedFromAtkObject()
	}
	C.atk_table_set_summary(this1, accessible1)
}
type TextLike interface {
	ImplementsAtkText() *C.AtkText
}

type Text struct {
	gobject.Object
	TextImpl
}

type TextImpl struct {}

func ToText(objlike gobject.ObjectLike) *Text {
	t := (*TextImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Text)(obj)
	}
	panic("cannot cast to Text")
}

func (this0 *TextImpl) ImplementsAtkText() *C.AtkText {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkText)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *TextImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_text_get_type())
}

func TextGetType() gobject.Type {
	return (*TextImpl)(nil).GetStaticType()
}
func TextFreeRanges(ranges0 *TextRange) {
	var ranges1 *C.AtkTextRange
	ranges1 = (*C.AtkTextRange)(unsafe.Pointer(ranges0))
	C.atk_text_free_ranges(ranges1)
}
func (this0 *TextImpl) AddSelection(start_offset0 int, end_offset0 int) bool {
	var this1 *C.AtkText
	var start_offset1 C.int32_t
	var end_offset1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	start_offset1 = C.int32_t(start_offset0)
	end_offset1 = C.int32_t(end_offset0)
	ret1 := C.atk_text_add_selection(this1, start_offset1, end_offset1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TextImpl) GetBoundedRanges(rect0 *TextRectangle, coord_type0 CoordType, x_clip_type0 TextClipType, y_clip_type0 TextClipType) []*TextRange {
	var this1 *C.AtkText
	var rect1 *C.AtkTextRectangle
	var coord_type1 C.AtkCoordType
	var x_clip_type1 C.AtkTextClipType
	var y_clip_type1 C.AtkTextClipType
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	rect1 = (*C.AtkTextRectangle)(unsafe.Pointer(rect0))
	coord_type1 = C.AtkCoordType(coord_type0)
	x_clip_type1 = C.AtkTextClipType(x_clip_type0)
	y_clip_type1 = C.AtkTextClipType(y_clip_type0)
	ret1 := C.atk_text_get_bounded_ranges(this1, rect1, coord_type1, x_clip_type1, y_clip_type1)
	var ret2 []*TextRange
	for i := range ret2 {
		ret2[i] = (*TextRange)(unsafe.Pointer((*(*[999999]*C.AtkTextRange)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *TextImpl) GetCaretOffset() int {
	var this1 *C.AtkText
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	ret1 := C.atk_text_get_caret_offset(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TextImpl) GetCharacterAtOffset(offset0 int) rune {
	var this1 *C.AtkText
	var offset1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	offset1 = C.int32_t(offset0)
	ret1 := C.atk_text_get_character_at_offset(this1, offset1)
	var ret2 rune
	ret2 = rune(ret1)
	return ret2
}
func (this0 *TextImpl) GetCharacterCount() int {
	var this1 *C.AtkText
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	ret1 := C.atk_text_get_character_count(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TextImpl) GetCharacterExtents(offset0 int, x0 *int, y0 *int, width0 *int, height0 *int, coords0 CoordType) {
	var this1 *C.AtkText
	var offset1 C.int32_t
	var x1 *C.int32_t
	var y1 *C.int32_t
	var width1 *C.int32_t
	var height1 *C.int32_t
	var coords1 C.AtkCoordType
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	offset1 = C.int32_t(offset0)
	x1 = (*C.int32_t)(unsafe.Pointer(x0))
	y1 = (*C.int32_t)(unsafe.Pointer(y0))
	width1 = (*C.int32_t)(unsafe.Pointer(width0))
	height1 = (*C.int32_t)(unsafe.Pointer(height0))
	coords1 = C.AtkCoordType(coords0)
	C.atk_text_get_character_extents(this1, offset1, x1, y1, width1, height1, coords1)
}
// blacklisted: Text.get_default_attributes (method)
func (this0 *TextImpl) GetNSelections() int {
	var this1 *C.AtkText
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	ret1 := C.atk_text_get_n_selections(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TextImpl) GetOffsetAtPoint(x0 int, y0 int, coords0 CoordType) int {
	var this1 *C.AtkText
	var x1 C.int32_t
	var y1 C.int32_t
	var coords1 C.AtkCoordType
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	coords1 = C.AtkCoordType(coords0)
	ret1 := C.atk_text_get_offset_at_point(this1, x1, y1, coords1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TextImpl) GetRangeExtents(start_offset0 int, end_offset0 int, coord_type0 CoordType, rect0 *TextRectangle) {
	var this1 *C.AtkText
	var start_offset1 C.int32_t
	var end_offset1 C.int32_t
	var coord_type1 C.AtkCoordType
	var rect1 *C.AtkTextRectangle
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	start_offset1 = C.int32_t(start_offset0)
	end_offset1 = C.int32_t(end_offset0)
	coord_type1 = C.AtkCoordType(coord_type0)
	rect1 = (*C.AtkTextRectangle)(unsafe.Pointer(rect0))
	C.atk_text_get_range_extents(this1, start_offset1, end_offset1, coord_type1, rect1)
}
// blacklisted: Text.get_run_attributes (method)
func (this0 *TextImpl) GetSelection(selection_num0 int, start_offset0 *int, end_offset0 *int) string {
	var this1 *C.AtkText
	var selection_num1 C.int32_t
	var start_offset1 *C.int32_t
	var end_offset1 *C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	selection_num1 = C.int32_t(selection_num0)
	start_offset1 = (*C.int32_t)(unsafe.Pointer(start_offset0))
	end_offset1 = (*C.int32_t)(unsafe.Pointer(end_offset0))
	ret1 := C.atk_text_get_selection(this1, selection_num1, start_offset1, end_offset1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *TextImpl) GetText(start_offset0 int, end_offset0 int) string {
	var this1 *C.AtkText
	var start_offset1 C.int32_t
	var end_offset1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	start_offset1 = C.int32_t(start_offset0)
	end_offset1 = C.int32_t(end_offset0)
	ret1 := C.atk_text_get_text(this1, start_offset1, end_offset1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *TextImpl) GetTextAfterOffset(offset0 int, boundary_type0 TextBoundary, start_offset0 *int, end_offset0 *int) string {
	var this1 *C.AtkText
	var offset1 C.int32_t
	var boundary_type1 C.AtkTextBoundary
	var start_offset1 *C.int32_t
	var end_offset1 *C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	offset1 = C.int32_t(offset0)
	boundary_type1 = C.AtkTextBoundary(boundary_type0)
	start_offset1 = (*C.int32_t)(unsafe.Pointer(start_offset0))
	end_offset1 = (*C.int32_t)(unsafe.Pointer(end_offset0))
	ret1 := C.atk_text_get_text_after_offset(this1, offset1, boundary_type1, start_offset1, end_offset1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *TextImpl) GetTextAtOffset(offset0 int, boundary_type0 TextBoundary, start_offset0 *int, end_offset0 *int) string {
	var this1 *C.AtkText
	var offset1 C.int32_t
	var boundary_type1 C.AtkTextBoundary
	var start_offset1 *C.int32_t
	var end_offset1 *C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	offset1 = C.int32_t(offset0)
	boundary_type1 = C.AtkTextBoundary(boundary_type0)
	start_offset1 = (*C.int32_t)(unsafe.Pointer(start_offset0))
	end_offset1 = (*C.int32_t)(unsafe.Pointer(end_offset0))
	ret1 := C.atk_text_get_text_at_offset(this1, offset1, boundary_type1, start_offset1, end_offset1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *TextImpl) GetTextBeforeOffset(offset0 int, boundary_type0 TextBoundary, start_offset0 *int, end_offset0 *int) string {
	var this1 *C.AtkText
	var offset1 C.int32_t
	var boundary_type1 C.AtkTextBoundary
	var start_offset1 *C.int32_t
	var end_offset1 *C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	offset1 = C.int32_t(offset0)
	boundary_type1 = C.AtkTextBoundary(boundary_type0)
	start_offset1 = (*C.int32_t)(unsafe.Pointer(start_offset0))
	end_offset1 = (*C.int32_t)(unsafe.Pointer(end_offset0))
	ret1 := C.atk_text_get_text_before_offset(this1, offset1, boundary_type1, start_offset1, end_offset1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *TextImpl) RemoveSelection(selection_num0 int) bool {
	var this1 *C.AtkText
	var selection_num1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	selection_num1 = C.int32_t(selection_num0)
	ret1 := C.atk_text_remove_selection(this1, selection_num1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TextImpl) SetCaretOffset(offset0 int) bool {
	var this1 *C.AtkText
	var offset1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	offset1 = C.int32_t(offset0)
	ret1 := C.atk_text_set_caret_offset(this1, offset1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TextImpl) SetSelection(selection_num0 int, start_offset0 int, end_offset0 int) bool {
	var this1 *C.AtkText
	var selection_num1 C.int32_t
	var start_offset1 C.int32_t
	var end_offset1 C.int32_t
	if this0 != nil {
		this1 = this0.ImplementsAtkText()}
	selection_num1 = C.int32_t(selection_num0)
	start_offset1 = C.int32_t(start_offset0)
	end_offset1 = C.int32_t(end_offset0)
	ret1 := C.atk_text_set_selection(this1, selection_num1, start_offset1, end_offset1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type TextAttribute C.uint32_t
const (
	TextAttributeInvalid TextAttribute = 0
	TextAttributeLeftMargin TextAttribute = 1
	TextAttributeRightMargin TextAttribute = 2
	TextAttributeIndent TextAttribute = 3
	TextAttributeInvisible TextAttribute = 4
	TextAttributeEditable TextAttribute = 5
	TextAttributePixelsAboveLines TextAttribute = 6
	TextAttributePixelsBelowLines TextAttribute = 7
	TextAttributePixelsInsideWrap TextAttribute = 8
	TextAttributeBgFullHeight TextAttribute = 9
	TextAttributeRise TextAttribute = 10
	TextAttributeUnderline TextAttribute = 11
	TextAttributeStrikethrough TextAttribute = 12
	TextAttributeSize TextAttribute = 13
	TextAttributeScale TextAttribute = 14
	TextAttributeWeight TextAttribute = 15
	TextAttributeLanguage TextAttribute = 16
	TextAttributeFamilyName TextAttribute = 17
	TextAttributeBgColor TextAttribute = 18
	TextAttributeFgColor TextAttribute = 19
	TextAttributeBgStipple TextAttribute = 20
	TextAttributeFgStipple TextAttribute = 21
	TextAttributeWrapMode TextAttribute = 22
	TextAttributeDirection TextAttribute = 23
	TextAttributeJustification TextAttribute = 24
	TextAttributeStretch TextAttribute = 25
	TextAttributeVariant TextAttribute = 26
	TextAttributeStyle TextAttribute = 27
	TextAttributeLastDefined TextAttribute = 28
)
type TextBoundary C.uint32_t
const (
	TextBoundaryChar TextBoundary = 0
	TextBoundaryWordStart TextBoundary = 1
	TextBoundaryWordEnd TextBoundary = 2
	TextBoundarySentenceStart TextBoundary = 3
	TextBoundarySentenceEnd TextBoundary = 4
	TextBoundaryLineStart TextBoundary = 5
	TextBoundaryLineEnd TextBoundary = 6
)
type TextClipType C.uint32_t
const (
	TextClipTypeNone TextClipType = 0
	TextClipTypeMin TextClipType = 1
	TextClipTypeMax TextClipType = 2
	TextClipTypeBoth TextClipType = 3
)
type TextRange struct {
	Bounds TextRectangle
	StartOffset int32
	EndOffset int32
	content0 *C.char
}
func (this0 *TextRange) Content() string {
	var content1 string
	content1 = C.GoString(this0.content0)
	return content1
}
type TextRectangle struct {
	X int32
	Y int32
	Width int32
	Height int32
}
type UtilLike interface {
	gobject.ObjectLike
	InheritedFromAtkUtil() *C.AtkUtil
}

type Util struct {
	gobject.Object
	
}

func ToUtil(objlike gobject.ObjectLike) *Util {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Util)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Util)(obj)
	}
	panic("cannot cast to Util")
}

func (this0 *Util) InheritedFromAtkUtil() *C.AtkUtil {
	if this0 == nil {
		return nil
	}
	return (*C.AtkUtil)(this0.C)
}

func (this0 *Util) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_util_get_type())
}

func UtilGetType() gobject.Type {
	return (*Util)(nil).GetStaticType()
}
type ValueLike interface {
	ImplementsAtkValue() *C.AtkValue
}

type Value struct {
	gobject.Object
	ValueImpl
}

type ValueImpl struct {}

func ToValue(objlike gobject.ObjectLike) *Value {
	t := (*ValueImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Value)(obj)
	}
	panic("cannot cast to Value")
}

func (this0 *ValueImpl) ImplementsAtkValue() *C.AtkValue {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkValue)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *ValueImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_value_get_type())
}

func ValueGetType() gobject.Type {
	return (*ValueImpl)(nil).GetStaticType()
}
func (this0 *ValueImpl) GetCurrentValue(value0 *gobject.Value) {
	var this1 *C.AtkValue
	var value1 *C.GValue
	if this0 != nil {
		this1 = this0.ImplementsAtkValue()}
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	C.atk_value_get_current_value(this1, value1)
}
func (this0 *ValueImpl) GetMaximumValue(value0 *gobject.Value) {
	var this1 *C.AtkValue
	var value1 *C.GValue
	if this0 != nil {
		this1 = this0.ImplementsAtkValue()}
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	C.atk_value_get_maximum_value(this1, value1)
}
func (this0 *ValueImpl) GetMinimumIncrement(value0 *gobject.Value) {
	var this1 *C.AtkValue
	var value1 *C.GValue
	if this0 != nil {
		this1 = this0.ImplementsAtkValue()}
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	C.atk_value_get_minimum_increment(this1, value1)
}
func (this0 *ValueImpl) GetMinimumValue(value0 *gobject.Value) {
	var this1 *C.AtkValue
	var value1 *C.GValue
	if this0 != nil {
		this1 = this0.ImplementsAtkValue()}
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	C.atk_value_get_minimum_value(this1, value1)
}
func (this0 *ValueImpl) SetCurrentValue(value0 *gobject.Value) bool {
	var this1 *C.AtkValue
	var value1 *C.GValue
	if this0 != nil {
		this1 = this0.ImplementsAtkValue()}
	value1 = (*C.GValue)(unsafe.Pointer(value0))
	ret1 := C.atk_value_set_current_value(this1, value1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type WindowLike interface {
	ImplementsAtkWindow() *C.AtkWindow
}

type Window struct {
	gobject.Object
	WindowImpl
}

type WindowImpl struct {}

func ToWindow(objlike gobject.ObjectLike) *Window {
	t := (*WindowImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Window)(obj)
	}
	panic("cannot cast to Window")
}

func (this0 *WindowImpl) ImplementsAtkWindow() *C.AtkWindow {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.AtkWindow)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *WindowImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.atk_window_get_type())
}

func WindowGetType() gobject.Type {
	return (*WindowImpl)(nil).GetStaticType()
}
type _PropertyValues struct {
	property_name0 *C.char
	OldValue gobject.Value
	NewValue gobject.Value
}
func (this0 *_PropertyValues) PropertyName() string {
	var property_name1 string
	property_name1 = C.GoString(this0.property_name0)
	return property_name1
}
// blacklisted: _Registry (struct)
// blacklisted: _RegistryClass (struct)
// blacklisted: attribute_set_free (function)
func FocusTrackerNotify(object0 ObjectLike) {
	var object1 *C.AtkObject
	if object0 != nil {
		object1 = object0.InheritedFromAtkObject()
	}
	C.atk_focus_tracker_notify(object1)
}
func GetDefaultRegistry() *Registry {
	ret1 := C.atk_get_default_registry()
	var ret2 *Registry
	ret2 = (*Registry)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func GetFocusObject() *Object {
	ret1 := C.atk_get_focus_object()
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func GetRoot() *Object {
	ret1 := C.atk_get_root()
	var ret2 *Object
	ret2 = (*Object)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func GetToolkitName() string {
	ret1 := C.atk_get_toolkit_name()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func GetToolkitVersion() string {
	ret1 := C.atk_get_toolkit_version()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func GetVersion() string {
	ret1 := C.atk_get_version()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func RelationTypeForName(name0 string) RelationType {
	var name1 *C.char
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.atk_relation_type_for_name(name1)
	var ret2 RelationType
	ret2 = RelationType(ret1)
	return ret2
}
func RelationTypeGetName(type0 RelationType) string {
	var type1 C.AtkRelationType
	type1 = C.AtkRelationType(type0)
	ret1 := C.atk_relation_type_get_name(type1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func RelationTypeRegister(name0 string) RelationType {
	var name1 *C.char
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.atk_relation_type_register(name1)
	var ret2 RelationType
	ret2 = RelationType(ret1)
	return ret2
}
func RemoveFocusTracker(tracker_id0 int) {
	var tracker_id1 C.uint32_t
	tracker_id1 = C.uint32_t(tracker_id0)
	C.atk_remove_focus_tracker(tracker_id1)
}
func RemoveGlobalEventListener(listener_id0 int) {
	var listener_id1 C.uint32_t
	listener_id1 = C.uint32_t(listener_id0)
	C.atk_remove_global_event_listener(listener_id1)
}
func RemoveKeyEventListener(listener_id0 int) {
	var listener_id1 C.uint32_t
	listener_id1 = C.uint32_t(listener_id0)
	C.atk_remove_key_event_listener(listener_id1)
}
func RoleForName(name0 string) Role {
	var name1 *C.char
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.atk_role_for_name(name1)
	var ret2 Role
	ret2 = Role(ret1)
	return ret2
}
func RoleGetLocalizedName(role0 Role) string {
	var role1 C.AtkRole
	role1 = C.AtkRole(role0)
	ret1 := C.atk_role_get_localized_name(role1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func RoleGetName(role0 Role) string {
	var role1 C.AtkRole
	role1 = C.AtkRole(role0)
	ret1 := C.atk_role_get_name(role1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func RoleRegister(name0 string) Role {
	var name1 *C.char
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.atk_role_register(name1)
	var ret2 Role
	ret2 = Role(ret1)
	return ret2
}
func StateTypeForName(name0 string) StateType {
	var name1 *C.char
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.atk_state_type_for_name(name1)
	var ret2 StateType
	ret2 = StateType(ret1)
	return ret2
}
func StateTypeGetName(type0 StateType) string {
	var type1 C.AtkStateType
	type1 = C.AtkStateType(type0)
	ret1 := C.atk_state_type_get_name(type1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func StateTypeRegister(name0 string) StateType {
	var name1 *C.char
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.atk_state_type_register(name1)
	var ret2 StateType
	ret2 = StateType(ret1)
	return ret2
}
func TextAttributeForName(name0 string) TextAttribute {
	var name1 *C.char
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.atk_text_attribute_for_name(name1)
	var ret2 TextAttribute
	ret2 = TextAttribute(ret1)
	return ret2
}
func TextAttributeGetName(attr0 TextAttribute) string {
	var attr1 C.AtkTextAttribute
	attr1 = C.AtkTextAttribute(attr0)
	ret1 := C.atk_text_attribute_get_name(attr1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func TextAttributeGetValue(attr0 TextAttribute, index_0 int) string {
	var attr1 C.AtkTextAttribute
	var index_1 C.int32_t
	attr1 = C.AtkTextAttribute(attr0)
	index_1 = C.int32_t(index_0)
	ret1 := C.atk_text_attribute_get_value(attr1, index_1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func TextAttributeRegister(name0 string) TextAttribute {
	var name1 *C.char
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.atk_text_attribute_register(name1)
	var ret2 TextAttribute
	ret2 = TextAttribute(ret1)
	return ret2
}
// blacklisted: text_free_ranges (function)
