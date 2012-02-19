package pango

/*
#include <stdlib.h>
#include <stdint.h>

typedef size_t GType;

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
typedef uint32_t PangoAlignment;
typedef struct _PangoAnalysis PangoAnalysis;
typedef struct _PangoAttrClass PangoAttrClass;
typedef struct _PangoAttrColor PangoAttrColor;
typedef void* PangoAttrFilterFunc;
extern void _PangoAttrFilterFunc_c_wrapper();
extern void _PangoAttrFilterFunc_c_wrapper_once();
typedef struct _PangoAttrFloat PangoAttrFloat;
typedef struct _PangoAttrFontDesc PangoAttrFontDesc;
typedef struct _PangoAttrInt PangoAttrInt;
typedef struct _PangoAttrIterator PangoAttrIterator;
typedef struct _PangoAttrLanguage PangoAttrLanguage;
typedef struct _PangoAttrList PangoAttrList;
typedef struct _PangoAttrShape PangoAttrShape;
typedef struct _PangoAttrSize PangoAttrSize;
typedef struct _PangoAttrString PangoAttrString;
typedef uint32_t PangoAttrType;
typedef struct _PangoAttribute PangoAttribute;
typedef uint32_t PangoBidiType;
typedef struct _PangoColor PangoColor;
typedef struct _PangoContext PangoContext;
typedef struct _PangoContextClass PangoContextClass;
typedef struct _PangoCoverage PangoCoverage;
typedef uint32_t PangoCoverageLevel;
typedef uint32_t PangoDirection;
typedef uint32_t PangoEllipsizeMode;
typedef struct _PangoEngineLang PangoEngineLang;
typedef struct _PangoEngineShape PangoEngineShape;
typedef struct _PangoFont PangoFont;
typedef struct _PangoFontDescription PangoFontDescription;
typedef struct _PangoFontFace PangoFontFace;
typedef struct _PangoFontFamily PangoFontFamily;
typedef struct _PangoFontMap PangoFontMap;
typedef uint32_t PangoFontMask;
typedef struct _PangoFontMetrics PangoFontMetrics;
typedef struct _PangoFontset PangoFontset;
typedef void* PangoFontsetForeachFunc;
extern void _PangoFontsetForeachFunc_c_wrapper();
extern void _PangoFontsetForeachFunc_c_wrapper_once();
typedef struct _PangoGlyphGeometry PangoGlyphGeometry;
typedef struct _PangoGlyphInfo PangoGlyphInfo;
typedef struct _PangoGlyphItem PangoGlyphItem;
typedef struct _PangoGlyphItemIter PangoGlyphItemIter;
typedef struct _PangoGlyphString PangoGlyphString;
typedef struct _PangoGlyphVisAttr PangoGlyphVisAttr;
typedef uint32_t PangoGravity;
typedef uint32_t PangoGravityHint;
typedef struct _PangoItem PangoItem;
typedef struct _PangoLanguage PangoLanguage;
typedef struct _PangoLayout PangoLayout;
typedef struct _PangoLayoutClass PangoLayoutClass;
typedef struct _PangoLayoutIter PangoLayoutIter;
typedef struct _PangoLayoutLine PangoLayoutLine;
typedef struct _PangoLogAttr PangoLogAttr;
typedef struct _PangoMatrix PangoMatrix;
typedef struct _PangoRectangle PangoRectangle;
typedef uint32_t PangoRenderPart;
typedef struct _PangoRenderer PangoRenderer;
typedef struct _PangoRendererClass PangoRendererClass;
typedef struct _PangoRendererPrivate PangoRendererPrivate;
typedef int32_t PangoScript;
typedef struct _PangoScriptIter PangoScriptIter;
typedef uint32_t PangoStretch;
typedef uint32_t PangoStyle;
typedef uint32_t PangoTabAlign;
typedef struct _PangoTabArray PangoTabArray;
typedef uint32_t PangoUnderline;
typedef uint32_t PangoVariant;
typedef uint32_t PangoWeight;
typedef uint32_t PangoWrapMode;
typedef struct _Pango_ScriptForLang Pango_ScriptForLang;
extern void pango_attr_iterator_destroy(PangoAttrIterator*);
extern GSList* pango_attr_iterator_get_attrs(PangoAttrIterator*);
extern void pango_attr_iterator_get_font(PangoAttrIterator*, PangoFontDescription*, PangoLanguage*, GSList*);
extern int pango_attr_iterator_next(PangoAttrIterator*);
extern void pango_attr_iterator_range(PangoAttrIterator*, int32_t*, int32_t*);
extern PangoAttrList* pango_attr_list_new();
extern void pango_attr_list_change(PangoAttrList*, PangoAttribute*);
extern PangoAttrList* pango_attr_list_copy(PangoAttrList*);
extern PangoAttrList* pango_attr_list_filter(PangoAttrList*, PangoAttrFilterFunc, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static PangoAttrList* _pango_attr_list_filter(PangoAttrList* this, void* gofunc) {
	if (gofunc) {
		return pango_attr_list_filter(this, _PangoAttrFilterFunc_c_wrapper, gofunc);
	} else {
		return pango_attr_list_filter(this, 0, 0);
	}
}
extern void pango_attr_list_insert(PangoAttrList*, PangoAttribute*);
extern void pango_attr_list_insert_before(PangoAttrList*, PangoAttribute*);
extern PangoAttrList* pango_attr_list_ref(PangoAttrList*);
extern void pango_attr_list_splice(PangoAttrList*, PangoAttrList*, int32_t, int32_t);
extern void pango_attr_list_unref(PangoAttrList*);
extern void pango_attribute_destroy(PangoAttribute*);
extern int pango_attribute_equal(PangoAttribute*, PangoAttribute*);
extern void pango_attribute_init(PangoAttribute*, PangoAttrClass*);
extern PangoColor* pango_color_copy(PangoColor*);
extern void pango_color_free(PangoColor*);
extern int pango_color_parse(PangoColor*, char*);
extern char* pango_color_to_string(PangoColor*);
extern PangoContext* pango_context_new();
extern PangoDirection pango_context_get_base_dir(PangoContext*);
extern PangoGravity pango_context_get_base_gravity(PangoContext*);
extern PangoFontDescription* pango_context_get_font_description(PangoContext*);
extern PangoFontMap* pango_context_get_font_map(PangoContext*);
extern PangoGravity pango_context_get_gravity(PangoContext*);
extern PangoGravityHint pango_context_get_gravity_hint(PangoContext*);
extern PangoLanguage* pango_context_get_language(PangoContext*);
extern PangoMatrix* pango_context_get_matrix(PangoContext*);
extern PangoFontMetrics* pango_context_get_metrics(PangoContext*, PangoFontDescription*, PangoLanguage*);
extern void pango_context_list_families(PangoContext*, PangoFontFamily***, int32_t*);
extern PangoFont* pango_context_load_font(PangoContext*, PangoFontDescription*);
extern PangoFontset* pango_context_load_fontset(PangoContext*, PangoFontDescription*, PangoLanguage*);
extern void pango_context_set_base_dir(PangoContext*, PangoDirection);
extern void pango_context_set_base_gravity(PangoContext*, PangoGravity);
extern void pango_context_set_font_description(PangoContext*, PangoFontDescription*);
extern void pango_context_set_font_map(PangoContext*, PangoFontMap*);
extern void pango_context_set_gravity_hint(PangoContext*, PangoGravityHint);
extern void pango_context_set_language(PangoContext*, PangoLanguage*);
extern void pango_context_set_matrix(PangoContext*, PangoMatrix*);
extern GType pango_context_get_type();
extern PangoCoverageLevel pango_coverage_get(PangoCoverage*, int32_t);
extern void pango_coverage_max(PangoCoverage*, PangoCoverage*);
extern void pango_coverage_set(PangoCoverage*, int32_t, PangoCoverageLevel);
extern void pango_coverage_to_bytes(PangoCoverage*, uint8_t*, int32_t*);
extern void pango_coverage_unref(PangoCoverage*);
extern void pango_font_descriptions_free(PangoFontDescription**, int32_t);
extern PangoFontDescription* pango_font_describe(PangoFont*);
extern PangoFontDescription* pango_font_describe_with_absolute_size(PangoFont*);
extern PangoFontMap* pango_font_get_font_map(PangoFont*);
extern void pango_font_get_glyph_extents(PangoFont*, uint32_t, PangoRectangle*, PangoRectangle*);
extern PangoFontMetrics* pango_font_get_metrics(PangoFont*, PangoLanguage*);
extern GType pango_font_get_type();
extern PangoFontDescription* pango_font_description_new();
extern int pango_font_description_better_match(PangoFontDescription*, PangoFontDescription*, PangoFontDescription*);
extern PangoFontDescription* pango_font_description_copy(PangoFontDescription*);
extern PangoFontDescription* pango_font_description_copy_static(PangoFontDescription*);
extern int pango_font_description_equal(PangoFontDescription*, PangoFontDescription*);
extern void pango_font_description_free(PangoFontDescription*);
extern char* pango_font_description_get_family(PangoFontDescription*);
extern PangoGravity pango_font_description_get_gravity(PangoFontDescription*);
extern PangoFontMask pango_font_description_get_set_fields(PangoFontDescription*);
extern int32_t pango_font_description_get_size(PangoFontDescription*);
extern int pango_font_description_get_size_is_absolute(PangoFontDescription*);
extern PangoStretch pango_font_description_get_stretch(PangoFontDescription*);
extern PangoStyle pango_font_description_get_style(PangoFontDescription*);
extern PangoVariant pango_font_description_get_variant(PangoFontDescription*);
extern PangoWeight pango_font_description_get_weight(PangoFontDescription*);
extern uint32_t pango_font_description_hash(PangoFontDescription*);
extern void pango_font_description_merge(PangoFontDescription*, PangoFontDescription*, int);
extern void pango_font_description_merge_static(PangoFontDescription*, PangoFontDescription*, int);
extern void pango_font_description_set_absolute_size(PangoFontDescription*, double);
extern void pango_font_description_set_family(PangoFontDescription*, char*);
extern void pango_font_description_set_family_static(PangoFontDescription*, char*);
extern void pango_font_description_set_gravity(PangoFontDescription*, PangoGravity);
extern void pango_font_description_set_size(PangoFontDescription*, int32_t);
extern void pango_font_description_set_stretch(PangoFontDescription*, PangoStretch);
extern void pango_font_description_set_style(PangoFontDescription*, PangoStyle);
extern void pango_font_description_set_variant(PangoFontDescription*, PangoVariant);
extern void pango_font_description_set_weight(PangoFontDescription*, PangoWeight);
extern char* pango_font_description_to_filename(PangoFontDescription*);
extern char* pango_font_description_to_string(PangoFontDescription*);
extern void pango_font_description_unset_fields(PangoFontDescription*, PangoFontMask);
extern PangoFontDescription* pango_font_description_from_string(char*);
extern PangoFontDescription* pango_font_face_describe(PangoFontFace*);
extern char* pango_font_face_get_face_name(PangoFontFace*);
extern int pango_font_face_is_synthesized(PangoFontFace*);
extern void pango_font_face_list_sizes(PangoFontFace*, int32_t*, int32_t*);
extern GType pango_font_face_get_type();
extern char* pango_font_family_get_name(PangoFontFamily*);
extern int pango_font_family_is_monospace(PangoFontFamily*);
extern void pango_font_family_list_faces(PangoFontFamily*, PangoFontFace***, int32_t*);
extern GType pango_font_family_get_type();
extern PangoContext* pango_font_map_create_context(PangoFontMap*);
extern void pango_font_map_list_families(PangoFontMap*, PangoFontFamily***, int32_t*);
extern PangoFont* pango_font_map_load_font(PangoFontMap*, PangoContext*, PangoFontDescription*);
extern PangoFontset* pango_font_map_load_fontset(PangoFontMap*, PangoContext*, PangoFontDescription*, PangoLanguage*);
extern GType pango_font_map_get_type();
extern int32_t pango_font_metrics_get_approximate_char_width(PangoFontMetrics*);
extern int32_t pango_font_metrics_get_approximate_digit_width(PangoFontMetrics*);
extern int32_t pango_font_metrics_get_ascent(PangoFontMetrics*);
extern int32_t pango_font_metrics_get_descent(PangoFontMetrics*);
extern int32_t pango_font_metrics_get_strikethrough_position(PangoFontMetrics*);
extern int32_t pango_font_metrics_get_strikethrough_thickness(PangoFontMetrics*);
extern int32_t pango_font_metrics_get_underline_position(PangoFontMetrics*);
extern int32_t pango_font_metrics_get_underline_thickness(PangoFontMetrics*);
extern PangoFontMetrics* pango_font_metrics_ref(PangoFontMetrics*);
extern void pango_font_metrics_unref(PangoFontMetrics*);
extern void pango_fontset_foreach(PangoFontset*, PangoFontsetForeachFunc, void*);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _pango_fontset_foreach(PangoFontset* this, void* gofunc) {
	if (gofunc) {
		pango_fontset_foreach(this, _PangoFontsetForeachFunc_c_wrapper, gofunc);
	} else {
		pango_fontset_foreach(this, 0, 0);
	}
}
extern PangoFont* pango_fontset_get_font(PangoFontset*, uint32_t);
extern PangoFontMetrics* pango_fontset_get_metrics(PangoFontset*);
extern GType pango_fontset_get_type();
extern PangoGlyphItem* pango_glyph_item_copy(PangoGlyphItem*);
extern void pango_glyph_item_free(PangoGlyphItem*);
extern void pango_glyph_item_get_logical_widths(PangoGlyphItem*, char*, int32_t*);
extern void pango_glyph_item_letter_space(PangoGlyphItem*, char*, PangoLogAttr*, int32_t);
extern PangoGlyphItem* pango_glyph_item_split(PangoGlyphItem*, char*, int32_t);
extern PangoGlyphItemIter* pango_glyph_item_iter_copy(PangoGlyphItemIter*);
extern void pango_glyph_item_iter_free(PangoGlyphItemIter*);
extern int pango_glyph_item_iter_init_end(PangoGlyphItemIter*, PangoGlyphItem*, char*);
extern int pango_glyph_item_iter_init_start(PangoGlyphItemIter*, PangoGlyphItem*, char*);
extern int pango_glyph_item_iter_next_cluster(PangoGlyphItemIter*);
extern int pango_glyph_item_iter_prev_cluster(PangoGlyphItemIter*);
extern PangoGlyphString* pango_glyph_string_new();
extern PangoGlyphString* pango_glyph_string_copy(PangoGlyphString*);
extern void pango_glyph_string_extents(PangoGlyphString*, PangoFont*, PangoRectangle*, PangoRectangle*);
extern void pango_glyph_string_extents_range(PangoGlyphString*, int32_t, int32_t, PangoFont*, PangoRectangle*, PangoRectangle*);
extern void pango_glyph_string_free(PangoGlyphString*);
extern void pango_glyph_string_get_logical_widths(PangoGlyphString*, char*, int32_t, int32_t, int32_t*);
extern int32_t pango_glyph_string_get_width(PangoGlyphString*);
extern void pango_glyph_string_index_to_x(PangoGlyphString*, char*, int32_t, PangoAnalysis*, int32_t, int, int32_t*);
extern void pango_glyph_string_set_size(PangoGlyphString*, int32_t);
extern void pango_glyph_string_x_to_index(PangoGlyphString*, char*, int32_t, PangoAnalysis*, int32_t, int32_t*, int32_t*);
extern PangoItem* pango_item_new();
extern PangoItem* pango_item_copy(PangoItem*);
extern void pango_item_free(PangoItem*);
extern PangoItem* pango_item_split(PangoItem*, int32_t, int32_t);
extern char* pango_language_get_sample_string(PangoLanguage*);
extern PangoScript* pango_language_get_scripts(PangoLanguage*, int32_t*);
extern int pango_language_includes_script(PangoLanguage*, PangoScript);
extern int pango_language_matches(PangoLanguage*, char*);
extern char* pango_language_to_string(PangoLanguage*);
extern PangoLanguage* pango_language_from_string(char*);
extern PangoLanguage* pango_language_get_default();
extern PangoLayout* pango_layout_new(PangoContext*);
extern void pango_layout_context_changed(PangoLayout*);
extern PangoLayout* pango_layout_copy(PangoLayout*);
extern PangoAlignment pango_layout_get_alignment(PangoLayout*);
extern PangoAttrList* pango_layout_get_attributes(PangoLayout*);
extern int pango_layout_get_auto_dir(PangoLayout*);
extern int32_t pango_layout_get_baseline(PangoLayout*);
extern int32_t pango_layout_get_character_count(PangoLayout*);
extern PangoContext* pango_layout_get_context(PangoLayout*);
extern void pango_layout_get_cursor_pos(PangoLayout*, int32_t, PangoRectangle*, PangoRectangle*);
extern PangoEllipsizeMode pango_layout_get_ellipsize(PangoLayout*);
extern void pango_layout_get_extents(PangoLayout*, PangoRectangle*, PangoRectangle*);
extern PangoFontDescription* pango_layout_get_font_description(PangoLayout*);
extern int32_t pango_layout_get_height(PangoLayout*);
extern int32_t pango_layout_get_indent(PangoLayout*);
extern PangoLayoutIter* pango_layout_get_iter(PangoLayout*);
extern int pango_layout_get_justify(PangoLayout*);
extern PangoLayoutLine* pango_layout_get_line(PangoLayout*, int32_t);
extern int32_t pango_layout_get_line_count(PangoLayout*);
extern PangoLayoutLine* pango_layout_get_line_readonly(PangoLayout*, int32_t);
extern GSList* pango_layout_get_lines(PangoLayout*);
extern GSList* pango_layout_get_lines_readonly(PangoLayout*);
extern void pango_layout_get_log_attrs(PangoLayout*, PangoLogAttr**, int32_t*);
extern PangoLogAttr* pango_layout_get_log_attrs_readonly(PangoLayout*, int32_t*);
extern void pango_layout_get_pixel_extents(PangoLayout*, PangoRectangle*, PangoRectangle*);
extern void pango_layout_get_pixel_size(PangoLayout*, int32_t*, int32_t*);
extern int pango_layout_get_single_paragraph_mode(PangoLayout*);
extern void pango_layout_get_size(PangoLayout*, int32_t*, int32_t*);
extern int32_t pango_layout_get_spacing(PangoLayout*);
extern PangoTabArray* pango_layout_get_tabs(PangoLayout*);
extern char* pango_layout_get_text(PangoLayout*);
extern int32_t pango_layout_get_unknown_glyphs_count(PangoLayout*);
extern int32_t pango_layout_get_width(PangoLayout*);
extern PangoWrapMode pango_layout_get_wrap(PangoLayout*);
extern void pango_layout_index_to_line_x(PangoLayout*, int32_t, int, int32_t*, int32_t*);
extern void pango_layout_index_to_pos(PangoLayout*, int32_t, PangoRectangle*);
extern int pango_layout_is_ellipsized(PangoLayout*);
extern int pango_layout_is_wrapped(PangoLayout*);
extern void pango_layout_move_cursor_visually(PangoLayout*, int, int32_t, int32_t, int32_t, int32_t*, int32_t*);
extern void pango_layout_set_alignment(PangoLayout*, PangoAlignment);
extern void pango_layout_set_attributes(PangoLayout*, PangoAttrList*);
extern void pango_layout_set_auto_dir(PangoLayout*, int);
extern void pango_layout_set_ellipsize(PangoLayout*, PangoEllipsizeMode);
extern void pango_layout_set_font_description(PangoLayout*, PangoFontDescription*);
extern void pango_layout_set_height(PangoLayout*, int32_t);
extern void pango_layout_set_indent(PangoLayout*, int32_t);
extern void pango_layout_set_justify(PangoLayout*, int);
extern void pango_layout_set_markup(PangoLayout*, char*, int32_t);
extern void pango_layout_set_markup_with_accel(PangoLayout*, char*, int32_t, uint32_t, uint32_t*);
extern void pango_layout_set_single_paragraph_mode(PangoLayout*, int);
extern void pango_layout_set_spacing(PangoLayout*, int32_t);
extern void pango_layout_set_tabs(PangoLayout*, PangoTabArray*);
extern void pango_layout_set_text(PangoLayout*, char*, int32_t);
extern void pango_layout_set_width(PangoLayout*, int32_t);
extern void pango_layout_set_wrap(PangoLayout*, PangoWrapMode);
extern int pango_layout_xy_to_index(PangoLayout*, int32_t, int32_t, int32_t*, int32_t*);
extern GType pango_layout_get_type();
extern int pango_layout_iter_at_last_line(PangoLayoutIter*);
extern PangoLayoutIter* pango_layout_iter_copy(PangoLayoutIter*);
extern void pango_layout_iter_free(PangoLayoutIter*);
extern int32_t pango_layout_iter_get_baseline(PangoLayoutIter*);
extern void pango_layout_iter_get_char_extents(PangoLayoutIter*, PangoRectangle*);
extern void pango_layout_iter_get_cluster_extents(PangoLayoutIter*, PangoRectangle*, PangoRectangle*);
extern int32_t pango_layout_iter_get_index(PangoLayoutIter*);
extern PangoLayout* pango_layout_iter_get_layout(PangoLayoutIter*);
extern void pango_layout_iter_get_layout_extents(PangoLayoutIter*, PangoRectangle*, PangoRectangle*);
extern PangoLayoutLine* pango_layout_iter_get_line(PangoLayoutIter*);
extern void pango_layout_iter_get_line_extents(PangoLayoutIter*, PangoRectangle*, PangoRectangle*);
extern PangoLayoutLine* pango_layout_iter_get_line_readonly(PangoLayoutIter*);
extern void pango_layout_iter_get_line_yrange(PangoLayoutIter*, int32_t*, int32_t*);
extern PangoGlyphItem* pango_layout_iter_get_run(PangoLayoutIter*);
extern void pango_layout_iter_get_run_extents(PangoLayoutIter*, PangoRectangle*, PangoRectangle*);
extern PangoGlyphItem* pango_layout_iter_get_run_readonly(PangoLayoutIter*);
extern int pango_layout_iter_next_char(PangoLayoutIter*);
extern int pango_layout_iter_next_cluster(PangoLayoutIter*);
extern int pango_layout_iter_next_line(PangoLayoutIter*);
extern int pango_layout_iter_next_run(PangoLayoutIter*);
extern void pango_layout_line_get_extents(PangoLayoutLine*, PangoRectangle*, PangoRectangle*);
extern void pango_layout_line_get_pixel_extents(PangoLayoutLine*, PangoRectangle*, PangoRectangle*);
extern void pango_layout_line_get_x_ranges(PangoLayoutLine*, int32_t, int32_t, int32_t**, int32_t*);
extern void pango_layout_line_index_to_x(PangoLayoutLine*, int32_t, int, int32_t*);
extern PangoLayoutLine* pango_layout_line_ref(PangoLayoutLine*);
extern void pango_layout_line_unref(PangoLayoutLine*);
extern int pango_layout_line_x_to_index(PangoLayoutLine*, int32_t, int32_t*, int32_t*);
extern void pango_matrix_concat(PangoMatrix*, PangoMatrix*);
extern PangoMatrix* pango_matrix_copy(PangoMatrix*);
extern void pango_matrix_free(PangoMatrix*);
extern double pango_matrix_get_font_scale_factor(PangoMatrix*);
extern void pango_matrix_rotate(PangoMatrix*, double);
extern void pango_matrix_scale(PangoMatrix*, double, double);
extern void pango_matrix_transform_distance(PangoMatrix*, double*, double*);
extern void pango_matrix_transform_pixel_rectangle(PangoMatrix*, PangoRectangle*);
extern void pango_matrix_transform_point(PangoMatrix*, double*, double*);
extern void pango_matrix_transform_rectangle(PangoMatrix*, PangoRectangle*);
extern void pango_matrix_translate(PangoMatrix*, double, double);
extern void pango_renderer_activate(PangoRenderer*);
extern void pango_renderer_deactivate(PangoRenderer*);
extern void pango_renderer_draw_error_underline(PangoRenderer*, int32_t, int32_t, int32_t, int32_t);
extern void pango_renderer_draw_glyph(PangoRenderer*, PangoFont*, uint32_t, double, double);
extern void pango_renderer_draw_glyph_item(PangoRenderer*, char*, PangoGlyphItem*, int32_t, int32_t);
extern void pango_renderer_draw_glyphs(PangoRenderer*, PangoFont*, PangoGlyphString*, int32_t, int32_t);
extern void pango_renderer_draw_layout(PangoRenderer*, PangoLayout*, int32_t, int32_t);
extern void pango_renderer_draw_layout_line(PangoRenderer*, PangoLayoutLine*, int32_t, int32_t);
extern void pango_renderer_draw_rectangle(PangoRenderer*, PangoRenderPart, int32_t, int32_t, int32_t, int32_t);
extern void pango_renderer_draw_trapezoid(PangoRenderer*, PangoRenderPart, double, double, double, double, double, double);
extern PangoColor* pango_renderer_get_color(PangoRenderer*, PangoRenderPart);
extern PangoLayout* pango_renderer_get_layout(PangoRenderer*);
extern PangoLayoutLine* pango_renderer_get_layout_line(PangoRenderer*);
extern PangoMatrix* pango_renderer_get_matrix(PangoRenderer*);
extern void pango_renderer_part_changed(PangoRenderer*, PangoRenderPart);
extern void pango_renderer_set_color(PangoRenderer*, PangoRenderPart, PangoColor*);
extern void pango_renderer_set_matrix(PangoRenderer*, PangoMatrix*);
extern GType pango_renderer_get_type();
extern void pango_script_iter_free(PangoScriptIter*);
extern void pango_script_iter_get_range(PangoScriptIter*, char**, char**, PangoScript*);
extern int pango_script_iter_next(PangoScriptIter*);
extern PangoTabArray* pango_tab_array_new(int32_t, int);
extern PangoTabArray* pango_tab_array_copy(PangoTabArray*);
extern void pango_tab_array_free(PangoTabArray*);
extern int pango_tab_array_get_positions_in_pixels(PangoTabArray*);
extern int32_t pango_tab_array_get_size(PangoTabArray*);
extern void pango_tab_array_get_tab(PangoTabArray*, int32_t, PangoTabAlign*, int32_t*);
extern void pango_tab_array_get_tabs(PangoTabArray*, PangoTabAlign**, int32_t**);
extern void pango_tab_array_resize(PangoTabArray*, int32_t);
extern void pango_tab_array_set_tab(PangoTabArray*, int32_t, PangoTabAlign, int32_t);
extern char* pango_attr_type_get_name(PangoAttrType);
extern PangoAttrType pango_attr_type_register(char*);
extern PangoBidiType pango_bidi_type_for_unichar(uint32_t);
extern void pango_break(char*, int32_t, PangoAnalysis*, PangoLogAttr*, int32_t);
extern void pango_extents_to_pixels(PangoRectangle*, PangoRectangle*);
extern PangoDirection pango_find_base_dir(char*, int32_t);
extern void pango_find_paragraph_boundary(char*, int32_t, int32_t*, int32_t*);
extern void pango_get_log_attrs(char*, int32_t, int32_t, PangoLanguage*, PangoLogAttr*, int32_t);
extern int pango_get_mirror_char(uint32_t, uint32_t*);
extern PangoGravity pango_gravity_get_for_matrix(PangoMatrix*);
extern PangoGravity pango_gravity_get_for_script(PangoScript, PangoGravity, PangoGravityHint);
extern PangoGravity pango_gravity_get_for_script_and_width(PangoScript, int, PangoGravity, PangoGravityHint);
extern double pango_gravity_to_rotation(PangoGravity);
extern int pango_is_zero_width(uint32_t);
extern GList* pango_itemize(PangoContext*, char*, int32_t, int32_t, PangoAttrList*, PangoAttrIterator*);
extern GList* pango_itemize_with_base_dir(PangoContext*, PangoDirection, char*, int32_t, int32_t, PangoAttrList*, PangoAttrIterator*);
extern uint8_t* pango_log2vis_get_embedding_levels(char*, int32_t, PangoDirection*);
extern int pango_parse_enum(GType, char*, int32_t*, int, char**);
extern int pango_parse_markup(char*, int32_t, uint32_t, PangoAttrList**, char**, uint32_t*, GError**);
extern int pango_parse_stretch(char*, PangoStretch*, int);
extern int pango_parse_style(char*, PangoStyle*, int);
extern int pango_parse_variant(char*, PangoVariant*, int);
extern int pango_parse_weight(char*, PangoWeight*, int);
extern void pango_quantize_line_geometry(int32_t*, int32_t*);
extern int32_t pango_read_line(void*, GString*);
extern int pango_scan_int(char**, int32_t*);
extern int pango_scan_string(char**, GString*);
extern int pango_scan_word(char**, GString*);
extern PangoScript pango_script_for_unichar(uint32_t);
extern PangoLanguage* pango_script_get_sample_language(PangoScript);
extern void pango_shape(char*, int32_t, PangoAnalysis*, PangoGlyphString*);
extern int pango_skip_space(char**);
extern char** pango_split_file_list(char*);
extern char* pango_trim_string(char*);
extern PangoDirection pango_unichar_direction(uint32_t);
extern int32_t pango_units_from_double(double);
extern double pango_units_to_double(int32_t);
extern int32_t pango_version();
extern char* pango_version_check(int32_t, int32_t, int32_t);
extern char* pango_version_string();
struct _PangoAnalysis { uint8_t _data[48]; };
struct _PangoAttrClass { uint8_t _data[32]; };
struct _PangoAttrColor { uint8_t _data[24]; };
struct _PangoAttrFloat { uint8_t _data[24]; };
struct _PangoAttrFontDesc { uint8_t _data[24]; };
struct _PangoAttrInt { uint8_t _data[24]; };
struct _PangoAttrIterator {};
struct _PangoAttrLanguage { uint8_t _data[24]; };
struct _PangoAttrList {};
struct _PangoAttrShape { uint8_t _data[72]; };
struct _PangoAttrSize { uint8_t _data[24]; };
struct _PangoAttrString { uint8_t _data[24]; };
struct _PangoAttribute { uint8_t _data[16]; };
struct _PangoColor { uint8_t _data[6]; };
struct _PangoContextClass {};
struct _PangoCoverage {};
struct _PangoEngineLang {};
struct _PangoEngineShape {};
struct _PangoFontDescription {};
struct _PangoFontMetrics {};
struct _PangoGlyphGeometry { uint8_t _data[12]; };
struct _PangoGlyphInfo { uint8_t _data[20]; };
struct _PangoGlyphItem { uint8_t _data[16]; };
struct _PangoGlyphItemIter { uint8_t _data[40]; };
struct _PangoGlyphString { uint8_t _data[32]; };
struct _PangoGlyphVisAttr { uint8_t _data[4]; };
struct _PangoItem { uint8_t _data[64]; };
struct _PangoLanguage {};
struct _PangoLayoutClass {};
struct _PangoLayoutIter {};
struct _PangoLayoutLine { uint8_t _data[32]; };
struct _PangoLogAttr { uint8_t _data[52]; };
struct _PangoMatrix { uint8_t _data[48]; };
struct _PangoRectangle { uint8_t _data[16]; };
struct _PangoRendererClass { uint8_t _data[248]; };
struct _PangoRendererPrivate {};
struct _PangoScriptIter {};
struct _PangoTabArray {};
struct _Pango_ScriptForLang { uint8_t _data[20]; };


extern GObject *g_object_ref_sink(GObject*);
extern void g_object_unref(GObject*);
extern void g_free(void*);
extern void g_error_free(GError*);

#cgo pkg-config: pango
*/
import "C"
import "unsafe"
import "errors"

// package dependencies
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


//export _Pango_go_callback_cleanup
func _Pango_go_callback_cleanup(gofunc unsafe.Pointer) {
	gobject.Holder.Release(gofunc)
}


const AnalysisFlagCenteredBaseline = 1
const AttrIndexFromTextBeginning = 0
type Alignment C.uint32_t
const (
	AlignmentLeft Alignment = 0
	AlignmentCenter Alignment = 1
	AlignmentRight Alignment = 2
)
type Analysis struct {
	ShapeEngine *EngineShape
	LangEngine *EngineLang
	font0 *C.PangoFont
	Level uint8
	Gravity uint8
	Flags uint8
	Script uint8
	_ [4]byte
	Language *Language
	extra_attrs0 *C.GSList
}
func (this0 *Analysis) Font() *Font {
	var font1 *Font
	font1 = (*Font)(gobject.ObjectWrap(unsafe.Pointer(this0.font0), true))
	return font1
}
func (this0 *Analysis) ExtraAttrs() []unsafe.Pointer {
	var extra_attrs1 []unsafe.Pointer
	for iter := (*_GSList)(unsafe.Pointer(this0.extra_attrs0)); iter != nil; iter = iter.next {
		var elt unsafe.Pointer
		elt = (unsafe.Pointer)(iter.data)
		extra_attrs1 = append(extra_attrs1, elt)
	}
	return extra_attrs1
}
type AttrClass struct {
	Type AttrType
	_ [4]byte
	Copy unsafe.Pointer
	Destroy unsafe.Pointer
	Equal unsafe.Pointer
}
type AttrColor struct {
	Attr Attribute
	Color Color
	_ [2]byte
}
type AttrFilterFunc func(attribute *Attribute) bool
//export _PangoAttrFilterFunc_c_wrapper
func _PangoAttrFilterFunc_c_wrapper(attribute0 unsafe.Pointer, data0 unsafe.Pointer) int32 {
	var attribute1 *Attribute
	var data1 AttrFilterFunc
	attribute1 = (*Attribute)(unsafe.Pointer((*C.PangoAttribute)(attribute0)))
	data1 = *(*AttrFilterFunc)(data0)
	ret1 := data1(attribute1)
	var ret2 C.int
	ret2 = _GoBoolToCBool(ret1)
	return (int32)(ret2)
}
//export _PangoAttrFilterFunc_c_wrapper_once
func _PangoAttrFilterFunc_c_wrapper_once(attribute0 unsafe.Pointer, data0 unsafe.Pointer) int32 {
	ret := _PangoAttrFilterFunc_c_wrapper(attribute0, data0)
	gobject.Holder.Release(data0)
	return ret
}
type AttrFloat struct {
	Attr Attribute
	Value float64
}
type AttrFontDesc struct {
	Attr Attribute
	Desc *FontDescription
}
type AttrInt struct {
	Attr Attribute
	Value int32
	_ [4]byte
}
type AttrIterator struct {}
func (this0 *AttrIterator) Destroy() {
	var this1 *C.PangoAttrIterator
	this1 = (*C.PangoAttrIterator)(unsafe.Pointer(this0))
	C.pango_attr_iterator_destroy(this1)
}
func (this0 *AttrIterator) GetAttrs() []Attribute {
	var this1 *C.PangoAttrIterator
	this1 = (*C.PangoAttrIterator)(unsafe.Pointer(this0))
	ret1 := C.pango_attr_iterator_get_attrs(this1)
	var ret2 []Attribute
	for iter := (*_GSList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt Attribute
		elt = *(*Attribute)(unsafe.Pointer((*C.PangoAttribute)(iter.data)))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *AttrIterator) GetFont(desc0 *FontDescription, language0 *Language, extra_attrs0 []Attribute) {
	var this1 *C.PangoAttrIterator
	var desc1 *C.PangoFontDescription
	var language1 *C.PangoLanguage
	var extra_attrs1 *C.GSList
	this1 = (*C.PangoAttrIterator)(unsafe.Pointer(this0))
	desc1 = (*C.PangoFontDescription)(unsafe.Pointer(desc0))
	language1 = (*C.PangoLanguage)(unsafe.Pointer(language0))
	C.pango_attr_iterator_get_font(this1, desc1, language1, extra_attrs1)
}
func (this0 *AttrIterator) Next() bool {
	var this1 *C.PangoAttrIterator
	this1 = (*C.PangoAttrIterator)(unsafe.Pointer(this0))
	ret1 := C.pango_attr_iterator_next(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type AttrLanguage struct {
	Attr Attribute
	Value *Language
}
type AttrList struct {}
func NewAttrList() *AttrList {
	ret1 := C.pango_attr_list_new()
	var ret2 *AttrList
	ret2 = (*AttrList)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *AttrList) Change(attr0 *Attribute) {
	var this1 *C.PangoAttrList
	var attr1 *C.PangoAttribute
	this1 = (*C.PangoAttrList)(unsafe.Pointer(this0))
	attr1 = (*C.PangoAttribute)(unsafe.Pointer(attr0))
	C.pango_attr_list_change(this1, attr1)
}
func (this0 *AttrList) Copy() *AttrList {
	var this1 *C.PangoAttrList
	this1 = (*C.PangoAttrList)(unsafe.Pointer(this0))
	ret1 := C.pango_attr_list_copy(this1)
	var ret2 *AttrList
	ret2 = (*AttrList)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *AttrList) Filter(func0 AttrFilterFunc) *AttrList {
	var this1 *C.PangoAttrList
	var func1 unsafe.Pointer
	this1 = (*C.PangoAttrList)(unsafe.Pointer(this0))
	if func0 != nil {
		func1 = unsafe.Pointer(&func0)}
	ret1 := C._pango_attr_list_filter(this1, func1)
	var ret2 *AttrList
	ret2 = (*AttrList)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *AttrList) Insert(attr0 *Attribute) {
	var this1 *C.PangoAttrList
	var attr1 *C.PangoAttribute
	this1 = (*C.PangoAttrList)(unsafe.Pointer(this0))
	attr1 = (*C.PangoAttribute)(unsafe.Pointer(attr0))
	C.pango_attr_list_insert(this1, attr1)
}
func (this0 *AttrList) InsertBefore(attr0 *Attribute) {
	var this1 *C.PangoAttrList
	var attr1 *C.PangoAttribute
	this1 = (*C.PangoAttrList)(unsafe.Pointer(this0))
	attr1 = (*C.PangoAttribute)(unsafe.Pointer(attr0))
	C.pango_attr_list_insert_before(this1, attr1)
}
func (this0 *AttrList) Splice(other0 *AttrList, pos0 int, len0 int) {
	var this1 *C.PangoAttrList
	var other1 *C.PangoAttrList
	var pos1 C.int32_t
	var len1 C.int32_t
	this1 = (*C.PangoAttrList)(unsafe.Pointer(this0))
	other1 = (*C.PangoAttrList)(unsafe.Pointer(other0))
	pos1 = C.int32_t(pos0)
	len1 = C.int32_t(len0)
	C.pango_attr_list_splice(this1, other1, pos1, len1)
}
type AttrShape struct {
	Attr Attribute
	InkRect Rectangle
	LogicalRect Rectangle
	Data unsafe.Pointer
	CopyFunc unsafe.Pointer
	DestroyFunc unsafe.Pointer
}
type AttrSize struct {
	Attr Attribute
	Size int32
	Absolute uint32
}
type AttrString struct {
	Attr Attribute
	value0 *C.char
}
func (this0 *AttrString) Value() string {
	var value1 string
	value1 = C.GoString(this0.value0)
	return value1
}
type AttrType C.uint32_t
const (
	AttrTypeInvalid AttrType = 0
	AttrTypeLanguage AttrType = 1
	AttrTypeFamily AttrType = 2
	AttrTypeStyle AttrType = 3
	AttrTypeWeight AttrType = 4
	AttrTypeVariant AttrType = 5
	AttrTypeStretch AttrType = 6
	AttrTypeSize AttrType = 7
	AttrTypeFontDesc AttrType = 8
	AttrTypeForeground AttrType = 9
	AttrTypeBackground AttrType = 10
	AttrTypeUnderline AttrType = 11
	AttrTypeStrikethrough AttrType = 12
	AttrTypeRise AttrType = 13
	AttrTypeShape AttrType = 14
	AttrTypeScale AttrType = 15
	AttrTypeFallback AttrType = 16
	AttrTypeLetterSpacing AttrType = 17
	AttrTypeUnderlineColor AttrType = 18
	AttrTypeStrikethroughColor AttrType = 19
	AttrTypeAbsoluteSize AttrType = 20
	AttrTypeGravity AttrType = 21
	AttrTypeGravityHint AttrType = 22
)
type Attribute struct {
	Klass *AttrClass
	StartIndex uint32
	EndIndex uint32
}
func (this0 *Attribute) Destroy() {
	var this1 *C.PangoAttribute
	this1 = (*C.PangoAttribute)(unsafe.Pointer(this0))
	C.pango_attribute_destroy(this1)
}
func (this0 *Attribute) Equal(attr20 *Attribute) bool {
	var this1 *C.PangoAttribute
	var attr21 *C.PangoAttribute
	this1 = (*C.PangoAttribute)(unsafe.Pointer(this0))
	attr21 = (*C.PangoAttribute)(unsafe.Pointer(attr20))
	ret1 := C.pango_attribute_equal(this1, attr21)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type BidiType C.uint32_t
const (
	BidiTypeL BidiType = 0
	BidiTypeLre BidiType = 1
	BidiTypeLro BidiType = 2
	BidiTypeR BidiType = 3
	BidiTypeAl BidiType = 4
	BidiTypeRle BidiType = 5
	BidiTypeRlo BidiType = 6
	BidiTypePdf BidiType = 7
	BidiTypeEn BidiType = 8
	BidiTypeEs BidiType = 9
	BidiTypeEt BidiType = 10
	BidiTypeAn BidiType = 11
	BidiTypeCs BidiType = 12
	BidiTypeNsm BidiType = 13
	BidiTypeBn BidiType = 14
	BidiTypeB BidiType = 15
	BidiTypeS BidiType = 16
	BidiTypeWs BidiType = 17
	BidiTypeOn BidiType = 18
)
type Color struct {
	Red uint16
	Green uint16
	Blue uint16
}
func (this0 *Color) Copy() *Color {
	var this1 *C.PangoColor
	this1 = (*C.PangoColor)(unsafe.Pointer(this0))
	ret1 := C.pango_color_copy(this1)
	var ret2 *Color
	ret2 = (*Color)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Color) Free() {
	var this1 *C.PangoColor
	this1 = (*C.PangoColor)(unsafe.Pointer(this0))
	C.pango_color_free(this1)
}
func (this0 *Color) Parse(spec0 string) bool {
	var this1 *C.PangoColor
	var spec1 *C.char
	this1 = (*C.PangoColor)(unsafe.Pointer(this0))
	spec1 = _GoStringToGString(spec0)
	defer C.free(unsafe.Pointer(spec1))
	ret1 := C.pango_color_parse(this1, spec1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Color) ToString() string {
	var this1 *C.PangoColor
	this1 = (*C.PangoColor)(unsafe.Pointer(this0))
	ret1 := C.pango_color_to_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
type ContextLike interface {
	gobject.ObjectLike
	InheritedFromPangoContext() *C.PangoContext
}

type Context struct {
	gobject.Object
	
}

func ToContext(objlike gobject.ObjectLike) *Context {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Context)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Context)(obj)
	}
	panic("cannot cast to Context")
}

func (this0 *Context) InheritedFromPangoContext() *C.PangoContext {
	if this0 == nil {
		return nil
	}
	return (*C.PangoContext)(this0.C)
}

func (this0 *Context) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_context_get_type())
}

func ContextGetType() gobject.Type {
	return (*Context)(nil).GetStaticType()
}
func NewContext() *Context {
	ret1 := C.pango_context_new()
	var ret2 *Context
	ret2 = (*Context)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Context) GetBaseDir() Direction {
	var this1 *C.PangoContext
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	ret1 := C.pango_context_get_base_dir(this1)
	var ret2 Direction
	ret2 = Direction(ret1)
	return ret2
}
func (this0 *Context) GetBaseGravity() Gravity {
	var this1 *C.PangoContext
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	ret1 := C.pango_context_get_base_gravity(this1)
	var ret2 Gravity
	ret2 = Gravity(ret1)
	return ret2
}
func (this0 *Context) GetFontDescription() *FontDescription {
	var this1 *C.PangoContext
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	ret1 := C.pango_context_get_font_description(this1)
	var ret2 *FontDescription
	ret2 = (*FontDescription)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Context) GetFontMap() *FontMap {
	var this1 *C.PangoContext
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	ret1 := C.pango_context_get_font_map(this1)
	var ret2 *FontMap
	ret2 = (*FontMap)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Context) GetGravity() Gravity {
	var this1 *C.PangoContext
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	ret1 := C.pango_context_get_gravity(this1)
	var ret2 Gravity
	ret2 = Gravity(ret1)
	return ret2
}
func (this0 *Context) GetGravityHint() GravityHint {
	var this1 *C.PangoContext
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	ret1 := C.pango_context_get_gravity_hint(this1)
	var ret2 GravityHint
	ret2 = GravityHint(ret1)
	return ret2
}
func (this0 *Context) GetLanguage() *Language {
	var this1 *C.PangoContext
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	ret1 := C.pango_context_get_language(this1)
	var ret2 *Language
	ret2 = (*Language)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Context) GetMatrix() *Matrix {
	var this1 *C.PangoContext
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	ret1 := C.pango_context_get_matrix(this1)
	var ret2 *Matrix
	ret2 = (*Matrix)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Context) GetMetrics(desc0 *FontDescription, language0 *Language) *FontMetrics {
	var this1 *C.PangoContext
	var desc1 *C.PangoFontDescription
	var language1 *C.PangoLanguage
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	desc1 = (*C.PangoFontDescription)(unsafe.Pointer(desc0))
	language1 = (*C.PangoLanguage)(unsafe.Pointer(language0))
	ret1 := C.pango_context_get_metrics(this1, desc1, language1)
	var ret2 *FontMetrics
	ret2 = (*FontMetrics)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Context) ListFamilies() []*FontFamily {
	var this1 *C.PangoContext
	var families1 **C.PangoFontFamily
	var n_families1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	C.pango_context_list_families(this1, &families1, &n_families1)
	var families2 []*FontFamily
	families2 = make([]*FontFamily, n_families1)
	for i := range families2 {
		families2[i] = (*FontFamily)(gobject.ObjectWrap(unsafe.Pointer((*(*[999999]*C.PangoFontFamily)(unsafe.Pointer(families1)))[i]), false))
	}
	return families2
}
func (this0 *Context) LoadFont(desc0 *FontDescription) *Font {
	var this1 *C.PangoContext
	var desc1 *C.PangoFontDescription
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	desc1 = (*C.PangoFontDescription)(unsafe.Pointer(desc0))
	ret1 := C.pango_context_load_font(this1, desc1)
	var ret2 *Font
	ret2 = (*Font)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Context) LoadFontset(desc0 *FontDescription, language0 *Language) *Fontset {
	var this1 *C.PangoContext
	var desc1 *C.PangoFontDescription
	var language1 *C.PangoLanguage
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	desc1 = (*C.PangoFontDescription)(unsafe.Pointer(desc0))
	language1 = (*C.PangoLanguage)(unsafe.Pointer(language0))
	ret1 := C.pango_context_load_fontset(this1, desc1, language1)
	var ret2 *Fontset
	ret2 = (*Fontset)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Context) SetBaseDir(direction0 Direction) {
	var this1 *C.PangoContext
	var direction1 C.PangoDirection
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	direction1 = C.PangoDirection(direction0)
	C.pango_context_set_base_dir(this1, direction1)
}
func (this0 *Context) SetBaseGravity(gravity0 Gravity) {
	var this1 *C.PangoContext
	var gravity1 C.PangoGravity
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	gravity1 = C.PangoGravity(gravity0)
	C.pango_context_set_base_gravity(this1, gravity1)
}
func (this0 *Context) SetFontDescription(desc0 *FontDescription) {
	var this1 *C.PangoContext
	var desc1 *C.PangoFontDescription
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	desc1 = (*C.PangoFontDescription)(unsafe.Pointer(desc0))
	C.pango_context_set_font_description(this1, desc1)
}
func (this0 *Context) SetFontMap(font_map0 FontMapLike) {
	var this1 *C.PangoContext
	var font_map1 *C.PangoFontMap
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	if font_map0 != nil {
		font_map1 = font_map0.InheritedFromPangoFontMap()
	}
	C.pango_context_set_font_map(this1, font_map1)
}
func (this0 *Context) SetGravityHint(hint0 GravityHint) {
	var this1 *C.PangoContext
	var hint1 C.PangoGravityHint
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	hint1 = C.PangoGravityHint(hint0)
	C.pango_context_set_gravity_hint(this1, hint1)
}
func (this0 *Context) SetLanguage(language0 *Language) {
	var this1 *C.PangoContext
	var language1 *C.PangoLanguage
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	language1 = (*C.PangoLanguage)(unsafe.Pointer(language0))
	C.pango_context_set_language(this1, language1)
}
func (this0 *Context) SetMatrix(matrix0 *Matrix) {
	var this1 *C.PangoContext
	var matrix1 *C.PangoMatrix
	if this0 != nil {
		this1 = this0.InheritedFromPangoContext()
	}
	matrix1 = (*C.PangoMatrix)(unsafe.Pointer(matrix0))
	C.pango_context_set_matrix(this1, matrix1)
}
type Coverage struct {}
func (this0 *Coverage) Get(index_0 int) CoverageLevel {
	var this1 *C.PangoCoverage
	var index_1 C.int32_t
	this1 = (*C.PangoCoverage)(unsafe.Pointer(this0))
	index_1 = C.int32_t(index_0)
	ret1 := C.pango_coverage_get(this1, index_1)
	var ret2 CoverageLevel
	ret2 = CoverageLevel(ret1)
	return ret2
}
func (this0 *Coverage) Max(other0 *Coverage) {
	var this1 *C.PangoCoverage
	var other1 *C.PangoCoverage
	this1 = (*C.PangoCoverage)(unsafe.Pointer(this0))
	other1 = (*C.PangoCoverage)(unsafe.Pointer(other0))
	C.pango_coverage_max(this1, other1)
}
func (this0 *Coverage) Set(index_0 int, level0 CoverageLevel) {
	var this1 *C.PangoCoverage
	var index_1 C.int32_t
	var level1 C.PangoCoverageLevel
	this1 = (*C.PangoCoverage)(unsafe.Pointer(this0))
	index_1 = C.int32_t(index_0)
	level1 = C.PangoCoverageLevel(level0)
	C.pango_coverage_set(this1, index_1, level1)
}
type CoverageLevel C.uint32_t
const (
	CoverageLevelNone CoverageLevel = 0
	CoverageLevelFallback CoverageLevel = 1
	CoverageLevelApproximate CoverageLevel = 2
	CoverageLevelExact CoverageLevel = 3
)
type Direction C.uint32_t
const (
	DirectionLtr Direction = 0
	DirectionRtl Direction = 1
	DirectionTtbLtr Direction = 2
	DirectionTtbRtl Direction = 3
	DirectionWeakLtr Direction = 4
	DirectionWeakRtl Direction = 5
	DirectionNeutral Direction = 6
)
const EngineTypeLang = "PangoEngineLang"
const EngineTypeShape = "PangoEngineShape"
type EllipsizeMode C.uint32_t
const (
	EllipsizeModeNone EllipsizeMode = 0
	EllipsizeModeStart EllipsizeMode = 1
	EllipsizeModeMiddle EllipsizeMode = 2
	EllipsizeModeEnd EllipsizeMode = 3
)
type EngineLang struct {}
type EngineShape struct {}
type FontLike interface {
	gobject.ObjectLike
	InheritedFromPangoFont() *C.PangoFont
}

type Font struct {
	gobject.Object
	
}

func ToFont(objlike gobject.ObjectLike) *Font {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Font)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Font)(obj)
	}
	panic("cannot cast to Font")
}

func (this0 *Font) InheritedFromPangoFont() *C.PangoFont {
	if this0 == nil {
		return nil
	}
	return (*C.PangoFont)(this0.C)
}

func (this0 *Font) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_font_get_type())
}

func FontGetType() gobject.Type {
	return (*Font)(nil).GetStaticType()
}
func FontDescriptionsFree(descs0 []*FontDescription) {
	var descs1 **C.PangoFontDescription
	var n_descs1 C.int32_t
	descs1 = (**C.PangoFontDescription)(C.malloc(C.size_t(int(unsafe.Sizeof(*descs1)) * len(descs0))))
	defer C.free(unsafe.Pointer(descs1))
	for i, e := range descs0 {
		(*(*[999999]*C.PangoFontDescription)(unsafe.Pointer(descs1)))[i] = (*C.PangoFontDescription)(unsafe.Pointer(e))
	}
	n_descs1 = C.int32_t(len(descs0))
	C.pango_font_descriptions_free(descs1, n_descs1)
}
func (this0 *Font) Describe() *FontDescription {
	var this1 *C.PangoFont
	if this0 != nil {
		this1 = this0.InheritedFromPangoFont()
	}
	ret1 := C.pango_font_describe(this1)
	var ret2 *FontDescription
	ret2 = (*FontDescription)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Font) DescribeWithAbsoluteSize() *FontDescription {
	var this1 *C.PangoFont
	if this0 != nil {
		this1 = this0.InheritedFromPangoFont()
	}
	ret1 := C.pango_font_describe_with_absolute_size(this1)
	var ret2 *FontDescription
	ret2 = (*FontDescription)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Font) GetFontMap() *FontMap {
	var this1 *C.PangoFont
	if this0 != nil {
		this1 = this0.InheritedFromPangoFont()
	}
	ret1 := C.pango_font_get_font_map(this1)
	var ret2 *FontMap
	ret2 = (*FontMap)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Font) GetGlyphExtents(glyph0 int) (Rectangle, Rectangle) {
	var this1 *C.PangoFont
	var glyph1 C.uint32_t
	var ink_rect1 C.PangoRectangle
	var logical_rect1 C.PangoRectangle
	if this0 != nil {
		this1 = this0.InheritedFromPangoFont()
	}
	glyph1 = C.uint32_t(glyph0)
	C.pango_font_get_glyph_extents(this1, glyph1, &ink_rect1, &logical_rect1)
	var ink_rect2 Rectangle
	var logical_rect2 Rectangle
	ink_rect2 = *(*Rectangle)(unsafe.Pointer(&ink_rect1))
	logical_rect2 = *(*Rectangle)(unsafe.Pointer(&logical_rect1))
	return ink_rect2, logical_rect2
}
func (this0 *Font) GetMetrics(language0 *Language) *FontMetrics {
	var this1 *C.PangoFont
	var language1 *C.PangoLanguage
	if this0 != nil {
		this1 = this0.InheritedFromPangoFont()
	}
	language1 = (*C.PangoLanguage)(unsafe.Pointer(language0))
	ret1 := C.pango_font_get_metrics(this1, language1)
	var ret2 *FontMetrics
	ret2 = (*FontMetrics)(unsafe.Pointer(ret1))
	return ret2
}
type FontDescription struct {}
func NewFontDescription() *FontDescription {
	ret1 := C.pango_font_description_new()
	var ret2 *FontDescription
	ret2 = (*FontDescription)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FontDescription) BetterMatch(old_match0 *FontDescription, new_match0 *FontDescription) bool {
	var this1 *C.PangoFontDescription
	var old_match1 *C.PangoFontDescription
	var new_match1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	old_match1 = (*C.PangoFontDescription)(unsafe.Pointer(old_match0))
	new_match1 = (*C.PangoFontDescription)(unsafe.Pointer(new_match0))
	ret1 := C.pango_font_description_better_match(this1, old_match1, new_match1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FontDescription) Copy() *FontDescription {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_copy(this1)
	var ret2 *FontDescription
	ret2 = (*FontDescription)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FontDescription) CopyStatic() *FontDescription {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_copy_static(this1)
	var ret2 *FontDescription
	ret2 = (*FontDescription)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FontDescription) Equal(desc20 *FontDescription) bool {
	var this1 *C.PangoFontDescription
	var desc21 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	desc21 = (*C.PangoFontDescription)(unsafe.Pointer(desc20))
	ret1 := C.pango_font_description_equal(this1, desc21)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FontDescription) Free() {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	C.pango_font_description_free(this1)
}
func (this0 *FontDescription) GetFamily() string {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_get_family(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FontDescription) GetGravity() Gravity {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_get_gravity(this1)
	var ret2 Gravity
	ret2 = Gravity(ret1)
	return ret2
}
func (this0 *FontDescription) GetSetFields() FontMask {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_get_set_fields(this1)
	var ret2 FontMask
	ret2 = FontMask(ret1)
	return ret2
}
func (this0 *FontDescription) GetSize() int {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_get_size(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *FontDescription) GetSizeIsAbsolute() bool {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_get_size_is_absolute(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FontDescription) GetStretch() Stretch {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_get_stretch(this1)
	var ret2 Stretch
	ret2 = Stretch(ret1)
	return ret2
}
func (this0 *FontDescription) GetStyle() Style {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_get_style(this1)
	var ret2 Style
	ret2 = Style(ret1)
	return ret2
}
func (this0 *FontDescription) GetVariant() Variant {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_get_variant(this1)
	var ret2 Variant
	ret2 = Variant(ret1)
	return ret2
}
func (this0 *FontDescription) GetWeight() Weight {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_get_weight(this1)
	var ret2 Weight
	ret2 = Weight(ret1)
	return ret2
}
func (this0 *FontDescription) Hash() int {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_hash(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *FontDescription) Merge(desc_to_merge0 *FontDescription, replace_existing0 bool) {
	var this1 *C.PangoFontDescription
	var desc_to_merge1 *C.PangoFontDescription
	var replace_existing1 C.int
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	desc_to_merge1 = (*C.PangoFontDescription)(unsafe.Pointer(desc_to_merge0))
	replace_existing1 = _GoBoolToCBool(replace_existing0)
	C.pango_font_description_merge(this1, desc_to_merge1, replace_existing1)
}
func (this0 *FontDescription) MergeStatic(desc_to_merge0 *FontDescription, replace_existing0 bool) {
	var this1 *C.PangoFontDescription
	var desc_to_merge1 *C.PangoFontDescription
	var replace_existing1 C.int
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	desc_to_merge1 = (*C.PangoFontDescription)(unsafe.Pointer(desc_to_merge0))
	replace_existing1 = _GoBoolToCBool(replace_existing0)
	C.pango_font_description_merge_static(this1, desc_to_merge1, replace_existing1)
}
func (this0 *FontDescription) SetAbsoluteSize(size0 float64) {
	var this1 *C.PangoFontDescription
	var size1 C.double
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	size1 = C.double(size0)
	C.pango_font_description_set_absolute_size(this1, size1)
}
func (this0 *FontDescription) SetFamily(family0 string) {
	var this1 *C.PangoFontDescription
	var family1 *C.char
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	family1 = _GoStringToGString(family0)
	defer C.free(unsafe.Pointer(family1))
	C.pango_font_description_set_family(this1, family1)
}
func (this0 *FontDescription) SetFamilyStatic(family0 string) {
	var this1 *C.PangoFontDescription
	var family1 *C.char
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	family1 = _GoStringToGString(family0)
	defer C.free(unsafe.Pointer(family1))
	C.pango_font_description_set_family_static(this1, family1)
}
func (this0 *FontDescription) SetGravity(gravity0 Gravity) {
	var this1 *C.PangoFontDescription
	var gravity1 C.PangoGravity
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	gravity1 = C.PangoGravity(gravity0)
	C.pango_font_description_set_gravity(this1, gravity1)
}
func (this0 *FontDescription) SetSize(size0 int) {
	var this1 *C.PangoFontDescription
	var size1 C.int32_t
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	size1 = C.int32_t(size0)
	C.pango_font_description_set_size(this1, size1)
}
func (this0 *FontDescription) SetStretch(stretch0 Stretch) {
	var this1 *C.PangoFontDescription
	var stretch1 C.PangoStretch
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	stretch1 = C.PangoStretch(stretch0)
	C.pango_font_description_set_stretch(this1, stretch1)
}
func (this0 *FontDescription) SetStyle(style0 Style) {
	var this1 *C.PangoFontDescription
	var style1 C.PangoStyle
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	style1 = C.PangoStyle(style0)
	C.pango_font_description_set_style(this1, style1)
}
func (this0 *FontDescription) SetVariant(variant0 Variant) {
	var this1 *C.PangoFontDescription
	var variant1 C.PangoVariant
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	variant1 = C.PangoVariant(variant0)
	C.pango_font_description_set_variant(this1, variant1)
}
func (this0 *FontDescription) SetWeight(weight0 Weight) {
	var this1 *C.PangoFontDescription
	var weight1 C.PangoWeight
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	weight1 = C.PangoWeight(weight0)
	C.pango_font_description_set_weight(this1, weight1)
}
func (this0 *FontDescription) ToFilename() string {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_to_filename(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FontDescription) ToString() string {
	var this1 *C.PangoFontDescription
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	ret1 := C.pango_font_description_to_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FontDescription) UnsetFields(to_unset0 FontMask) {
	var this1 *C.PangoFontDescription
	var to_unset1 C.PangoFontMask
	this1 = (*C.PangoFontDescription)(unsafe.Pointer(this0))
	to_unset1 = C.PangoFontMask(to_unset0)
	C.pango_font_description_unset_fields(this1, to_unset1)
}
func FontDescriptionFromString(str0 string) *FontDescription {
	var str1 *C.char
	str1 = _GoStringToGString(str0)
	defer C.free(unsafe.Pointer(str1))
	ret1 := C.pango_font_description_from_string(str1)
	var ret2 *FontDescription
	ret2 = (*FontDescription)(unsafe.Pointer(ret1))
	return ret2
}
type FontFaceLike interface {
	gobject.ObjectLike
	InheritedFromPangoFontFace() *C.PangoFontFace
}

type FontFace struct {
	gobject.Object
	
}

func ToFontFace(objlike gobject.ObjectLike) *FontFace {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*FontFace)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FontFace)(obj)
	}
	panic("cannot cast to FontFace")
}

func (this0 *FontFace) InheritedFromPangoFontFace() *C.PangoFontFace {
	if this0 == nil {
		return nil
	}
	return (*C.PangoFontFace)(this0.C)
}

func (this0 *FontFace) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_font_face_get_type())
}

func FontFaceGetType() gobject.Type {
	return (*FontFace)(nil).GetStaticType()
}
func (this0 *FontFace) Describe() *FontDescription {
	var this1 *C.PangoFontFace
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontFace()
	}
	ret1 := C.pango_font_face_describe(this1)
	var ret2 *FontDescription
	ret2 = (*FontDescription)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *FontFace) GetFaceName() string {
	var this1 *C.PangoFontFace
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontFace()
	}
	ret1 := C.pango_font_face_get_face_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FontFace) IsSynthesized() bool {
	var this1 *C.PangoFontFace
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontFace()
	}
	ret1 := C.pango_font_face_is_synthesized(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: FontFace.list_sizes (method)
type FontFamilyLike interface {
	gobject.ObjectLike
	InheritedFromPangoFontFamily() *C.PangoFontFamily
}

type FontFamily struct {
	gobject.Object
	
}

func ToFontFamily(objlike gobject.ObjectLike) *FontFamily {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*FontFamily)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FontFamily)(obj)
	}
	panic("cannot cast to FontFamily")
}

func (this0 *FontFamily) InheritedFromPangoFontFamily() *C.PangoFontFamily {
	if this0 == nil {
		return nil
	}
	return (*C.PangoFontFamily)(this0.C)
}

func (this0 *FontFamily) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_font_family_get_type())
}

func FontFamilyGetType() gobject.Type {
	return (*FontFamily)(nil).GetStaticType()
}
func (this0 *FontFamily) GetName() string {
	var this1 *C.PangoFontFamily
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontFamily()
	}
	ret1 := C.pango_font_family_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *FontFamily) IsMonospace() bool {
	var this1 *C.PangoFontFamily
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontFamily()
	}
	ret1 := C.pango_font_family_is_monospace(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *FontFamily) ListFaces() []*FontFace {
	var this1 *C.PangoFontFamily
	var faces1 **C.PangoFontFace
	var n_faces1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontFamily()
	}
	C.pango_font_family_list_faces(this1, &faces1, &n_faces1)
	var faces2 []*FontFace
	faces2 = make([]*FontFace, n_faces1)
	for i := range faces2 {
		faces2[i] = (*FontFace)(gobject.ObjectWrap(unsafe.Pointer((*(*[999999]*C.PangoFontFace)(unsafe.Pointer(faces1)))[i]), false))
	}
	return faces2
}
type FontMapLike interface {
	gobject.ObjectLike
	InheritedFromPangoFontMap() *C.PangoFontMap
}

type FontMap struct {
	gobject.Object
	
}

func ToFontMap(objlike gobject.ObjectLike) *FontMap {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*FontMap)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FontMap)(obj)
	}
	panic("cannot cast to FontMap")
}

func (this0 *FontMap) InheritedFromPangoFontMap() *C.PangoFontMap {
	if this0 == nil {
		return nil
	}
	return (*C.PangoFontMap)(this0.C)
}

func (this0 *FontMap) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_font_map_get_type())
}

func FontMapGetType() gobject.Type {
	return (*FontMap)(nil).GetStaticType()
}
func (this0 *FontMap) CreateContext() *Context {
	var this1 *C.PangoFontMap
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontMap()
	}
	ret1 := C.pango_font_map_create_context(this1)
	var ret2 *Context
	ret2 = (*Context)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *FontMap) ListFamilies() []*FontFamily {
	var this1 *C.PangoFontMap
	var families1 **C.PangoFontFamily
	var n_families1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontMap()
	}
	C.pango_font_map_list_families(this1, &families1, &n_families1)
	var families2 []*FontFamily
	families2 = make([]*FontFamily, n_families1)
	for i := range families2 {
		families2[i] = (*FontFamily)(gobject.ObjectWrap(unsafe.Pointer((*(*[999999]*C.PangoFontFamily)(unsafe.Pointer(families1)))[i]), false))
	}
	return families2
}
func (this0 *FontMap) LoadFont(context0 ContextLike, desc0 *FontDescription) *Font {
	var this1 *C.PangoFontMap
	var context1 *C.PangoContext
	var desc1 *C.PangoFontDescription
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontMap()
	}
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	desc1 = (*C.PangoFontDescription)(unsafe.Pointer(desc0))
	ret1 := C.pango_font_map_load_font(this1, context1, desc1)
	var ret2 *Font
	ret2 = (*Font)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *FontMap) LoadFontset(context0 ContextLike, desc0 *FontDescription, language0 *Language) *Fontset {
	var this1 *C.PangoFontMap
	var context1 *C.PangoContext
	var desc1 *C.PangoFontDescription
	var language1 *C.PangoLanguage
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontMap()
	}
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	desc1 = (*C.PangoFontDescription)(unsafe.Pointer(desc0))
	language1 = (*C.PangoLanguage)(unsafe.Pointer(language0))
	ret1 := C.pango_font_map_load_fontset(this1, context1, desc1, language1)
	var ret2 *Fontset
	ret2 = (*Fontset)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
type FontMask C.uint32_t
const (
	FontMaskFamily FontMask = 1
	FontMaskStyle FontMask = 2
	FontMaskVariant FontMask = 4
	FontMaskWeight FontMask = 8
	FontMaskStretch FontMask = 16
	FontMaskSize FontMask = 32
	FontMaskGravity FontMask = 64
)
type FontMetrics struct {}
func (this0 *FontMetrics) GetApproximateCharWidth() int {
	var this1 *C.PangoFontMetrics
	this1 = (*C.PangoFontMetrics)(unsafe.Pointer(this0))
	ret1 := C.pango_font_metrics_get_approximate_char_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *FontMetrics) GetApproximateDigitWidth() int {
	var this1 *C.PangoFontMetrics
	this1 = (*C.PangoFontMetrics)(unsafe.Pointer(this0))
	ret1 := C.pango_font_metrics_get_approximate_digit_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *FontMetrics) GetAscent() int {
	var this1 *C.PangoFontMetrics
	this1 = (*C.PangoFontMetrics)(unsafe.Pointer(this0))
	ret1 := C.pango_font_metrics_get_ascent(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *FontMetrics) GetDescent() int {
	var this1 *C.PangoFontMetrics
	this1 = (*C.PangoFontMetrics)(unsafe.Pointer(this0))
	ret1 := C.pango_font_metrics_get_descent(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *FontMetrics) GetStrikethroughPosition() int {
	var this1 *C.PangoFontMetrics
	this1 = (*C.PangoFontMetrics)(unsafe.Pointer(this0))
	ret1 := C.pango_font_metrics_get_strikethrough_position(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *FontMetrics) GetStrikethroughThickness() int {
	var this1 *C.PangoFontMetrics
	this1 = (*C.PangoFontMetrics)(unsafe.Pointer(this0))
	ret1 := C.pango_font_metrics_get_strikethrough_thickness(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *FontMetrics) GetUnderlinePosition() int {
	var this1 *C.PangoFontMetrics
	this1 = (*C.PangoFontMetrics)(unsafe.Pointer(this0))
	ret1 := C.pango_font_metrics_get_underline_position(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *FontMetrics) GetUnderlineThickness() int {
	var this1 *C.PangoFontMetrics
	this1 = (*C.PangoFontMetrics)(unsafe.Pointer(this0))
	ret1 := C.pango_font_metrics_get_underline_thickness(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
type FontsetLike interface {
	gobject.ObjectLike
	InheritedFromPangoFontset() *C.PangoFontset
}

type Fontset struct {
	gobject.Object
	
}

func ToFontset(objlike gobject.ObjectLike) *Fontset {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Fontset)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Fontset)(obj)
	}
	panic("cannot cast to Fontset")
}

func (this0 *Fontset) InheritedFromPangoFontset() *C.PangoFontset {
	if this0 == nil {
		return nil
	}
	return (*C.PangoFontset)(this0.C)
}

func (this0 *Fontset) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_fontset_get_type())
}

func FontsetGetType() gobject.Type {
	return (*Fontset)(nil).GetStaticType()
}
func (this0 *Fontset) Foreach(func0 FontsetForeachFunc) {
	var this1 *C.PangoFontset
	var func1 unsafe.Pointer
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontset()
	}
	if func0 != nil {
		func1 = unsafe.Pointer(&func0)}
	C._pango_fontset_foreach(this1, func1)
}
func (this0 *Fontset) GetFont(wc0 int) *Font {
	var this1 *C.PangoFontset
	var wc1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontset()
	}
	wc1 = C.uint32_t(wc0)
	ret1 := C.pango_fontset_get_font(this1, wc1)
	var ret2 *Font
	ret2 = (*Font)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Fontset) GetMetrics() *FontMetrics {
	var this1 *C.PangoFontset
	if this0 != nil {
		this1 = this0.InheritedFromPangoFontset()
	}
	ret1 := C.pango_fontset_get_metrics(this1)
	var ret2 *FontMetrics
	ret2 = (*FontMetrics)(unsafe.Pointer(ret1))
	return ret2
}
type FontsetForeachFunc func(fontset *Fontset, font *Font) bool
//export _PangoFontsetForeachFunc_c_wrapper
func _PangoFontsetForeachFunc_c_wrapper(fontset0 unsafe.Pointer, font0 unsafe.Pointer, data0 unsafe.Pointer) int32 {
	var fontset1 *Fontset
	var font1 *Font
	var data1 FontsetForeachFunc
	fontset1 = (*Fontset)(gobject.ObjectWrap(unsafe.Pointer((*C.PangoFontset)(fontset0)), true))
	font1 = (*Font)(gobject.ObjectWrap(unsafe.Pointer((*C.PangoFont)(font0)), true))
	data1 = *(*FontsetForeachFunc)(data0)
	ret1 := data1(fontset1, font1)
	var ret2 C.int
	ret2 = _GoBoolToCBool(ret1)
	return (int32)(ret2)
}
//export _PangoFontsetForeachFunc_c_wrapper_once
func _PangoFontsetForeachFunc_c_wrapper_once(fontset0 unsafe.Pointer, font0 unsafe.Pointer, data0 unsafe.Pointer) int32 {
	ret := _PangoFontsetForeachFunc_c_wrapper(fontset0, font0, data0)
	gobject.Holder.Release(data0)
	return ret
}
type GlyphGeometry struct {
	Width int32
	XOffset int32
	YOffset int32
}
type GlyphInfo struct {
	Glyph uint32
	Geometry GlyphGeometry
	Attr GlyphVisAttr
}
type GlyphItem struct {
	Item *Item
	Glyphs *GlyphString
}
func (this0 *GlyphItem) Copy() *GlyphItem {
	var this1 *C.PangoGlyphItem
	this1 = (*C.PangoGlyphItem)(unsafe.Pointer(this0))
	ret1 := C.pango_glyph_item_copy(this1)
	var ret2 *GlyphItem
	ret2 = (*GlyphItem)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *GlyphItem) Free() {
	var this1 *C.PangoGlyphItem
	this1 = (*C.PangoGlyphItem)(unsafe.Pointer(this0))
	C.pango_glyph_item_free(this1)
}
func (this0 *GlyphItem) LetterSpace(text0 string, log_attrs0 *LogAttr, letter_spacing0 int) {
	var this1 *C.PangoGlyphItem
	var text1 *C.char
	var log_attrs1 *C.PangoLogAttr
	var letter_spacing1 C.int32_t
	this1 = (*C.PangoGlyphItem)(unsafe.Pointer(this0))
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	log_attrs1 = (*C.PangoLogAttr)(unsafe.Pointer(log_attrs0))
	letter_spacing1 = C.int32_t(letter_spacing0)
	C.pango_glyph_item_letter_space(this1, text1, log_attrs1, letter_spacing1)
}
func (this0 *GlyphItem) Split(text0 string, split_index0 int) *GlyphItem {
	var this1 *C.PangoGlyphItem
	var text1 *C.char
	var split_index1 C.int32_t
	this1 = (*C.PangoGlyphItem)(unsafe.Pointer(this0))
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	split_index1 = C.int32_t(split_index0)
	ret1 := C.pango_glyph_item_split(this1, text1, split_index1)
	var ret2 *GlyphItem
	ret2 = (*GlyphItem)(unsafe.Pointer(ret1))
	return ret2
}
type GlyphItemIter struct {
	GlyphItem *GlyphItem
	text0 *C.char
	StartGlyph int32
	StartIndex int32
	StartChar int32
	EndGlyph int32
	EndIndex int32
	EndChar int32
}
func (this0 *GlyphItemIter) Text() string {
	var text1 string
	text1 = C.GoString(this0.text0)
	return text1
}
func (this0 *GlyphItemIter) Copy() *GlyphItemIter {
	var this1 *C.PangoGlyphItemIter
	this1 = (*C.PangoGlyphItemIter)(unsafe.Pointer(this0))
	ret1 := C.pango_glyph_item_iter_copy(this1)
	var ret2 *GlyphItemIter
	ret2 = (*GlyphItemIter)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *GlyphItemIter) Free() {
	var this1 *C.PangoGlyphItemIter
	this1 = (*C.PangoGlyphItemIter)(unsafe.Pointer(this0))
	C.pango_glyph_item_iter_free(this1)
}
func (this0 *GlyphItemIter) InitEnd(glyph_item0 *GlyphItem, text0 string) bool {
	var this1 *C.PangoGlyphItemIter
	var glyph_item1 *C.PangoGlyphItem
	var text1 *C.char
	this1 = (*C.PangoGlyphItemIter)(unsafe.Pointer(this0))
	glyph_item1 = (*C.PangoGlyphItem)(unsafe.Pointer(glyph_item0))
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	ret1 := C.pango_glyph_item_iter_init_end(this1, glyph_item1, text1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *GlyphItemIter) InitStart(glyph_item0 *GlyphItem, text0 string) bool {
	var this1 *C.PangoGlyphItemIter
	var glyph_item1 *C.PangoGlyphItem
	var text1 *C.char
	this1 = (*C.PangoGlyphItemIter)(unsafe.Pointer(this0))
	glyph_item1 = (*C.PangoGlyphItem)(unsafe.Pointer(glyph_item0))
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	ret1 := C.pango_glyph_item_iter_init_start(this1, glyph_item1, text1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *GlyphItemIter) NextCluster() bool {
	var this1 *C.PangoGlyphItemIter
	this1 = (*C.PangoGlyphItemIter)(unsafe.Pointer(this0))
	ret1 := C.pango_glyph_item_iter_next_cluster(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *GlyphItemIter) PrevCluster() bool {
	var this1 *C.PangoGlyphItemIter
	this1 = (*C.PangoGlyphItemIter)(unsafe.Pointer(this0))
	ret1 := C.pango_glyph_item_iter_prev_cluster(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type GlyphString struct {
	NumGlyphs int32
	_ [4]byte
	Glyphs *GlyphInfo
	LogClusters *int32
	Space int32
	_ [4]byte
}
func NewGlyphString() *GlyphString {
	ret1 := C.pango_glyph_string_new()
	var ret2 *GlyphString
	ret2 = (*GlyphString)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *GlyphString) Copy() *GlyphString {
	var this1 *C.PangoGlyphString
	this1 = (*C.PangoGlyphString)(unsafe.Pointer(this0))
	ret1 := C.pango_glyph_string_copy(this1)
	var ret2 *GlyphString
	ret2 = (*GlyphString)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *GlyphString) Extents(font0 FontLike) (Rectangle, Rectangle) {
	var this1 *C.PangoGlyphString
	var font1 *C.PangoFont
	var ink_rect1 C.PangoRectangle
	var logical_rect1 C.PangoRectangle
	this1 = (*C.PangoGlyphString)(unsafe.Pointer(this0))
	if font0 != nil {
		font1 = font0.InheritedFromPangoFont()
	}
	C.pango_glyph_string_extents(this1, font1, &ink_rect1, &logical_rect1)
	var ink_rect2 Rectangle
	var logical_rect2 Rectangle
	ink_rect2 = *(*Rectangle)(unsafe.Pointer(&ink_rect1))
	logical_rect2 = *(*Rectangle)(unsafe.Pointer(&logical_rect1))
	return ink_rect2, logical_rect2
}
func (this0 *GlyphString) ExtentsRange(start0 int, end0 int, font0 FontLike, ink_rect0 *Rectangle, logical_rect0 *Rectangle) {
	var this1 *C.PangoGlyphString
	var start1 C.int32_t
	var end1 C.int32_t
	var font1 *C.PangoFont
	var ink_rect1 *C.PangoRectangle
	var logical_rect1 *C.PangoRectangle
	this1 = (*C.PangoGlyphString)(unsafe.Pointer(this0))
	start1 = C.int32_t(start0)
	end1 = C.int32_t(end0)
	if font0 != nil {
		font1 = font0.InheritedFromPangoFont()
	}
	ink_rect1 = (*C.PangoRectangle)(unsafe.Pointer(ink_rect0))
	logical_rect1 = (*C.PangoRectangle)(unsafe.Pointer(logical_rect0))
	C.pango_glyph_string_extents_range(this1, start1, end1, font1, ink_rect1, logical_rect1)
}
func (this0 *GlyphString) Free() {
	var this1 *C.PangoGlyphString
	this1 = (*C.PangoGlyphString)(unsafe.Pointer(this0))
	C.pango_glyph_string_free(this1)
}
func (this0 *GlyphString) GetWidth() int {
	var this1 *C.PangoGlyphString
	this1 = (*C.PangoGlyphString)(unsafe.Pointer(this0))
	ret1 := C.pango_glyph_string_get_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *GlyphString) SetSize(new_len0 int) {
	var this1 *C.PangoGlyphString
	var new_len1 C.int32_t
	this1 = (*C.PangoGlyphString)(unsafe.Pointer(this0))
	new_len1 = C.int32_t(new_len0)
	C.pango_glyph_string_set_size(this1, new_len1)
}
type GlyphVisAttr struct {
	IsClusterStart uint32
}
type Gravity C.uint32_t
const (
	GravitySouth Gravity = 0
	GravityEast Gravity = 1
	GravityNorth Gravity = 2
	GravityWest Gravity = 3
	GravityAuto Gravity = 4
)
type GravityHint C.uint32_t
const (
	GravityHintNatural GravityHint = 0
	GravityHintStrong GravityHint = 1
	GravityHintLine GravityHint = 2
)
type Item struct {
	Offset int32
	Length int32
	NumChars int32
	_ [4]byte
	Analysis Analysis
}
func NewItem() *Item {
	ret1 := C.pango_item_new()
	var ret2 *Item
	ret2 = (*Item)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Item) Copy() *Item {
	var this1 *C.PangoItem
	this1 = (*C.PangoItem)(unsafe.Pointer(this0))
	ret1 := C.pango_item_copy(this1)
	var ret2 *Item
	ret2 = (*Item)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Item) Free() {
	var this1 *C.PangoItem
	this1 = (*C.PangoItem)(unsafe.Pointer(this0))
	C.pango_item_free(this1)
}
func (this0 *Item) Split(split_index0 int, split_offset0 int) *Item {
	var this1 *C.PangoItem
	var split_index1 C.int32_t
	var split_offset1 C.int32_t
	this1 = (*C.PangoItem)(unsafe.Pointer(this0))
	split_index1 = C.int32_t(split_index0)
	split_offset1 = C.int32_t(split_offset0)
	ret1 := C.pango_item_split(this1, split_index1, split_offset1)
	var ret2 *Item
	ret2 = (*Item)(unsafe.Pointer(ret1))
	return ret2
}
type Language struct {}
func (this0 *Language) GetSampleString() string {
	var this1 *C.PangoLanguage
	this1 = (*C.PangoLanguage)(unsafe.Pointer(this0))
	ret1 := C.pango_language_get_sample_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Language) IncludesScript(script0 Script) bool {
	var this1 *C.PangoLanguage
	var script1 C.PangoScript
	this1 = (*C.PangoLanguage)(unsafe.Pointer(this0))
	script1 = C.PangoScript(script0)
	ret1 := C.pango_language_includes_script(this1, script1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Language) Matches(range_list0 string) bool {
	var this1 *C.PangoLanguage
	var range_list1 *C.char
	this1 = (*C.PangoLanguage)(unsafe.Pointer(this0))
	range_list1 = _GoStringToGString(range_list0)
	defer C.free(unsafe.Pointer(range_list1))
	ret1 := C.pango_language_matches(this1, range_list1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Language) ToString() string {
	var this1 *C.PangoLanguage
	this1 = (*C.PangoLanguage)(unsafe.Pointer(this0))
	ret1 := C.pango_language_to_string(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func LanguageFromString(language0 string) *Language {
	var language1 *C.char
	language1 = _GoStringToGString(language0)
	defer C.free(unsafe.Pointer(language1))
	ret1 := C.pango_language_from_string(language1)
	var ret2 *Language
	ret2 = (*Language)(unsafe.Pointer(ret1))
	return ret2
}
func LanguageGetDefault() *Language {
	ret1 := C.pango_language_get_default()
	var ret2 *Language
	ret2 = (*Language)(unsafe.Pointer(ret1))
	return ret2
}
type LayoutLike interface {
	gobject.ObjectLike
	InheritedFromPangoLayout() *C.PangoLayout
}

type Layout struct {
	gobject.Object
	
}

func ToLayout(objlike gobject.ObjectLike) *Layout {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Layout)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Layout)(obj)
	}
	panic("cannot cast to Layout")
}

func (this0 *Layout) InheritedFromPangoLayout() *C.PangoLayout {
	if this0 == nil {
		return nil
	}
	return (*C.PangoLayout)(this0.C)
}

func (this0 *Layout) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_layout_get_type())
}

func LayoutGetType() gobject.Type {
	return (*Layout)(nil).GetStaticType()
}
func NewLayout(context0 ContextLike) *Layout {
	var context1 *C.PangoContext
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	ret1 := C.pango_layout_new(context1)
	var ret2 *Layout
	ret2 = (*Layout)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Layout) ContextChanged() {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	C.pango_layout_context_changed(this1)
}
func (this0 *Layout) Copy() *Layout {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_copy(this1)
	var ret2 *Layout
	ret2 = (*Layout)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Layout) GetAlignment() Alignment {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_alignment(this1)
	var ret2 Alignment
	ret2 = Alignment(ret1)
	return ret2
}
func (this0 *Layout) GetAttributes() *AttrList {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_attributes(this1)
	var ret2 *AttrList
	ret2 = (*AttrList)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Layout) GetAutoDir() bool {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_auto_dir(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Layout) GetBaseline() int {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_baseline(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Layout) GetCharacterCount() int {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_character_count(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Layout) GetContext() *Context {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_context(this1)
	var ret2 *Context
	ret2 = (*Context)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Layout) GetCursorPos(index_0 int) (Rectangle, Rectangle) {
	var this1 *C.PangoLayout
	var index_1 C.int32_t
	var strong_pos1 C.PangoRectangle
	var weak_pos1 C.PangoRectangle
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	index_1 = C.int32_t(index_0)
	C.pango_layout_get_cursor_pos(this1, index_1, &strong_pos1, &weak_pos1)
	var strong_pos2 Rectangle
	var weak_pos2 Rectangle
	strong_pos2 = *(*Rectangle)(unsafe.Pointer(&strong_pos1))
	weak_pos2 = *(*Rectangle)(unsafe.Pointer(&weak_pos1))
	return strong_pos2, weak_pos2
}
func (this0 *Layout) GetEllipsize() EllipsizeMode {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_ellipsize(this1)
	var ret2 EllipsizeMode
	ret2 = EllipsizeMode(ret1)
	return ret2
}
func (this0 *Layout) GetExtents() (Rectangle, Rectangle) {
	var this1 *C.PangoLayout
	var ink_rect1 C.PangoRectangle
	var logical_rect1 C.PangoRectangle
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	C.pango_layout_get_extents(this1, &ink_rect1, &logical_rect1)
	var ink_rect2 Rectangle
	var logical_rect2 Rectangle
	ink_rect2 = *(*Rectangle)(unsafe.Pointer(&ink_rect1))
	logical_rect2 = *(*Rectangle)(unsafe.Pointer(&logical_rect1))
	return ink_rect2, logical_rect2
}
func (this0 *Layout) GetFontDescription() *FontDescription {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_font_description(this1)
	var ret2 *FontDescription
	ret2 = (*FontDescription)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Layout) GetHeight() int {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_height(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Layout) GetIndent() int {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_indent(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Layout) GetIter() *LayoutIter {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_iter(this1)
	var ret2 *LayoutIter
	ret2 = (*LayoutIter)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Layout) GetJustify() bool {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_justify(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Layout) GetLine(line0 int) *LayoutLine {
	var this1 *C.PangoLayout
	var line1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	line1 = C.int32_t(line0)
	ret1 := C.pango_layout_get_line(this1, line1)
	var ret2 *LayoutLine
	ret2 = (*LayoutLine)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Layout) GetLineCount() int {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_line_count(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Layout) GetLineReadonly(line0 int) *LayoutLine {
	var this1 *C.PangoLayout
	var line1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	line1 = C.int32_t(line0)
	ret1 := C.pango_layout_get_line_readonly(this1, line1)
	var ret2 *LayoutLine
	ret2 = (*LayoutLine)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Layout) GetLines() []LayoutLine {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_lines(this1)
	var ret2 []LayoutLine
	for iter := (*_GSList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt LayoutLine
		elt = *(*LayoutLine)(unsafe.Pointer((*C.PangoLayoutLine)(iter.data)))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Layout) GetLinesReadonly() []LayoutLine {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_lines_readonly(this1)
	var ret2 []LayoutLine
	for iter := (*_GSList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt LayoutLine
		elt = *(*LayoutLine)(unsafe.Pointer((*C.PangoLayoutLine)(iter.data)))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Layout) GetLogAttrs() []LogAttr {
	var this1 *C.PangoLayout
	var attrs1 *C.PangoLogAttr
	var n_attrs1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	C.pango_layout_get_log_attrs(this1, &attrs1, &n_attrs1)
	var attrs2 []LogAttr
	attrs2 = make([]LogAttr, n_attrs1)
	for i := range attrs2 {
		attrs2[i] = *(*LogAttr)(unsafe.Pointer(&(*(*[999999]C.PangoLogAttr)(unsafe.Pointer(attrs1)))[i]))
	}
	return attrs2
}
// blacklisted: Layout.get_log_attrs_readonly (method)
func (this0 *Layout) GetPixelExtents() (Rectangle, Rectangle) {
	var this1 *C.PangoLayout
	var ink_rect1 C.PangoRectangle
	var logical_rect1 C.PangoRectangle
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	C.pango_layout_get_pixel_extents(this1, &ink_rect1, &logical_rect1)
	var ink_rect2 Rectangle
	var logical_rect2 Rectangle
	ink_rect2 = *(*Rectangle)(unsafe.Pointer(&ink_rect1))
	logical_rect2 = *(*Rectangle)(unsafe.Pointer(&logical_rect1))
	return ink_rect2, logical_rect2
}
func (this0 *Layout) GetPixelSize() (int, int) {
	var this1 *C.PangoLayout
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	C.pango_layout_get_pixel_size(this1, &width1, &height1)
	var width2 int
	var height2 int
	width2 = int(width1)
	height2 = int(height1)
	return width2, height2
}
func (this0 *Layout) GetSingleParagraphMode() bool {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_single_paragraph_mode(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Layout) GetSize() (int, int) {
	var this1 *C.PangoLayout
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	C.pango_layout_get_size(this1, &width1, &height1)
	var width2 int
	var height2 int
	width2 = int(width1)
	height2 = int(height1)
	return width2, height2
}
func (this0 *Layout) GetSpacing() int {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_spacing(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Layout) GetTabs() *TabArray {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_tabs(this1)
	var ret2 *TabArray
	ret2 = (*TabArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Layout) GetText() string {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_text(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Layout) GetUnknownGlyphsCount() int {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_unknown_glyphs_count(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Layout) GetWidth() int {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Layout) GetWrap() WrapMode {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_get_wrap(this1)
	var ret2 WrapMode
	ret2 = WrapMode(ret1)
	return ret2
}
func (this0 *Layout) IndexToLineX(index_0 int, trailing0 bool) (int, int) {
	var this1 *C.PangoLayout
	var index_1 C.int32_t
	var trailing1 C.int
	var line1 C.int32_t
	var x_pos1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	index_1 = C.int32_t(index_0)
	trailing1 = _GoBoolToCBool(trailing0)
	C.pango_layout_index_to_line_x(this1, index_1, trailing1, &line1, &x_pos1)
	var line2 int
	var x_pos2 int
	line2 = int(line1)
	x_pos2 = int(x_pos1)
	return line2, x_pos2
}
func (this0 *Layout) IndexToPos(index_0 int) Rectangle {
	var this1 *C.PangoLayout
	var index_1 C.int32_t
	var pos1 C.PangoRectangle
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	index_1 = C.int32_t(index_0)
	C.pango_layout_index_to_pos(this1, index_1, &pos1)
	var pos2 Rectangle
	pos2 = *(*Rectangle)(unsafe.Pointer(&pos1))
	return pos2
}
func (this0 *Layout) IsEllipsized() bool {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_is_ellipsized(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Layout) IsWrapped() bool {
	var this1 *C.PangoLayout
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ret1 := C.pango_layout_is_wrapped(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
// blacklisted: Layout.move_cursor_visually (method)
func (this0 *Layout) SetAlignment(alignment0 Alignment) {
	var this1 *C.PangoLayout
	var alignment1 C.PangoAlignment
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	alignment1 = C.PangoAlignment(alignment0)
	C.pango_layout_set_alignment(this1, alignment1)
}
func (this0 *Layout) SetAttributes(attrs0 *AttrList) {
	var this1 *C.PangoLayout
	var attrs1 *C.PangoAttrList
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	attrs1 = (*C.PangoAttrList)(unsafe.Pointer(attrs0))
	C.pango_layout_set_attributes(this1, attrs1)
}
func (this0 *Layout) SetAutoDir(auto_dir0 bool) {
	var this1 *C.PangoLayout
	var auto_dir1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	auto_dir1 = _GoBoolToCBool(auto_dir0)
	C.pango_layout_set_auto_dir(this1, auto_dir1)
}
func (this0 *Layout) SetEllipsize(ellipsize0 EllipsizeMode) {
	var this1 *C.PangoLayout
	var ellipsize1 C.PangoEllipsizeMode
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	ellipsize1 = C.PangoEllipsizeMode(ellipsize0)
	C.pango_layout_set_ellipsize(this1, ellipsize1)
}
func (this0 *Layout) SetFontDescription(desc0 *FontDescription) {
	var this1 *C.PangoLayout
	var desc1 *C.PangoFontDescription
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	desc1 = (*C.PangoFontDescription)(unsafe.Pointer(desc0))
	C.pango_layout_set_font_description(this1, desc1)
}
func (this0 *Layout) SetHeight(height0 int) {
	var this1 *C.PangoLayout
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	height1 = C.int32_t(height0)
	C.pango_layout_set_height(this1, height1)
}
func (this0 *Layout) SetIndent(indent0 int) {
	var this1 *C.PangoLayout
	var indent1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	indent1 = C.int32_t(indent0)
	C.pango_layout_set_indent(this1, indent1)
}
func (this0 *Layout) SetJustify(justify0 bool) {
	var this1 *C.PangoLayout
	var justify1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	justify1 = _GoBoolToCBool(justify0)
	C.pango_layout_set_justify(this1, justify1)
}
func (this0 *Layout) SetMarkup(markup0 string, length0 int) {
	var this1 *C.PangoLayout
	var markup1 *C.char
	var length1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	markup1 = _GoStringToGString(markup0)
	defer C.free(unsafe.Pointer(markup1))
	length1 = C.int32_t(length0)
	C.pango_layout_set_markup(this1, markup1, length1)
}
func (this0 *Layout) SetMarkupWithAccel(markup0 string, length0 int, accel_marker0 rune) rune {
	var this1 *C.PangoLayout
	var markup1 *C.char
	var length1 C.int32_t
	var accel_marker1 C.uint32_t
	var accel_char1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	markup1 = _GoStringToGString(markup0)
	defer C.free(unsafe.Pointer(markup1))
	length1 = C.int32_t(length0)
	accel_marker1 = C.uint32_t(accel_marker0)
	C.pango_layout_set_markup_with_accel(this1, markup1, length1, accel_marker1, &accel_char1)
	var accel_char2 rune
	accel_char2 = rune(accel_char1)
	return accel_char2
}
func (this0 *Layout) SetSingleParagraphMode(setting0 bool) {
	var this1 *C.PangoLayout
	var setting1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	setting1 = _GoBoolToCBool(setting0)
	C.pango_layout_set_single_paragraph_mode(this1, setting1)
}
func (this0 *Layout) SetSpacing(spacing0 int) {
	var this1 *C.PangoLayout
	var spacing1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	spacing1 = C.int32_t(spacing0)
	C.pango_layout_set_spacing(this1, spacing1)
}
func (this0 *Layout) SetTabs(tabs0 *TabArray) {
	var this1 *C.PangoLayout
	var tabs1 *C.PangoTabArray
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	tabs1 = (*C.PangoTabArray)(unsafe.Pointer(tabs0))
	C.pango_layout_set_tabs(this1, tabs1)
}
func (this0 *Layout) SetText(text0 string, length0 int) {
	var this1 *C.PangoLayout
	var text1 *C.char
	var length1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	length1 = C.int32_t(length0)
	C.pango_layout_set_text(this1, text1, length1)
}
func (this0 *Layout) SetWidth(width0 int) {
	var this1 *C.PangoLayout
	var width1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	width1 = C.int32_t(width0)
	C.pango_layout_set_width(this1, width1)
}
func (this0 *Layout) SetWrap(wrap0 WrapMode) {
	var this1 *C.PangoLayout
	var wrap1 C.PangoWrapMode
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	wrap1 = C.PangoWrapMode(wrap0)
	C.pango_layout_set_wrap(this1, wrap1)
}
func (this0 *Layout) XYToIndex(x0 int, y0 int) (int, int, bool) {
	var this1 *C.PangoLayout
	var x1 C.int32_t
	var y1 C.int32_t
	var index_1 C.int32_t
	var trailing1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoLayout()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	ret1 := C.pango_layout_xy_to_index(this1, x1, y1, &index_1, &trailing1)
	var index_2 int
	var trailing2 int
	var ret2 bool
	index_2 = int(index_1)
	trailing2 = int(trailing1)
	ret2 = ret1 != 0
	return index_2, trailing2, ret2
}
type LayoutIter struct {}
func (this0 *LayoutIter) AtLastLine() bool {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_at_last_line(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *LayoutIter) Copy() *LayoutIter {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_copy(this1)
	var ret2 *LayoutIter
	ret2 = (*LayoutIter)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *LayoutIter) Free() {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	C.pango_layout_iter_free(this1)
}
func (this0 *LayoutIter) GetBaseline() int {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_get_baseline(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *LayoutIter) GetCharExtents(logical_rect0 *Rectangle) {
	var this1 *C.PangoLayoutIter
	var logical_rect1 *C.PangoRectangle
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	logical_rect1 = (*C.PangoRectangle)(unsafe.Pointer(logical_rect0))
	C.pango_layout_iter_get_char_extents(this1, logical_rect1)
}
func (this0 *LayoutIter) GetClusterExtents() (Rectangle, Rectangle) {
	var this1 *C.PangoLayoutIter
	var ink_rect1 C.PangoRectangle
	var logical_rect1 C.PangoRectangle
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	C.pango_layout_iter_get_cluster_extents(this1, &ink_rect1, &logical_rect1)
	var ink_rect2 Rectangle
	var logical_rect2 Rectangle
	ink_rect2 = *(*Rectangle)(unsafe.Pointer(&ink_rect1))
	logical_rect2 = *(*Rectangle)(unsafe.Pointer(&logical_rect1))
	return ink_rect2, logical_rect2
}
func (this0 *LayoutIter) GetIndex() int {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_get_index(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *LayoutIter) GetLayout() *Layout {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_get_layout(this1)
	var ret2 *Layout
	ret2 = (*Layout)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *LayoutIter) GetLayoutExtents() (Rectangle, Rectangle) {
	var this1 *C.PangoLayoutIter
	var ink_rect1 C.PangoRectangle
	var logical_rect1 C.PangoRectangle
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	C.pango_layout_iter_get_layout_extents(this1, &ink_rect1, &logical_rect1)
	var ink_rect2 Rectangle
	var logical_rect2 Rectangle
	ink_rect2 = *(*Rectangle)(unsafe.Pointer(&ink_rect1))
	logical_rect2 = *(*Rectangle)(unsafe.Pointer(&logical_rect1))
	return ink_rect2, logical_rect2
}
func (this0 *LayoutIter) GetLine() *LayoutLine {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_get_line(this1)
	var ret2 *LayoutLine
	ret2 = (*LayoutLine)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *LayoutIter) GetLineExtents() (Rectangle, Rectangle) {
	var this1 *C.PangoLayoutIter
	var ink_rect1 C.PangoRectangle
	var logical_rect1 C.PangoRectangle
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	C.pango_layout_iter_get_line_extents(this1, &ink_rect1, &logical_rect1)
	var ink_rect2 Rectangle
	var logical_rect2 Rectangle
	ink_rect2 = *(*Rectangle)(unsafe.Pointer(&ink_rect1))
	logical_rect2 = *(*Rectangle)(unsafe.Pointer(&logical_rect1))
	return ink_rect2, logical_rect2
}
func (this0 *LayoutIter) GetLineReadonly() *LayoutLine {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_get_line_readonly(this1)
	var ret2 *LayoutLine
	ret2 = (*LayoutLine)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *LayoutIter) GetLineYRange() (int, int) {
	var this1 *C.PangoLayoutIter
	var y0_1 C.int32_t
	var y1_1 C.int32_t
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	C.pango_layout_iter_get_line_yrange(this1, &y0_1, &y1_1)
	var y0_2 int
	var y1_2 int
	y0_2 = int(y0_1)
	y1_2 = int(y1_1)
	return y0_2, y1_2
}
func (this0 *LayoutIter) GetRun() *GlyphItem {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_get_run(this1)
	var ret2 *GlyphItem
	ret2 = (*GlyphItem)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *LayoutIter) GetRunExtents() (Rectangle, Rectangle) {
	var this1 *C.PangoLayoutIter
	var ink_rect1 C.PangoRectangle
	var logical_rect1 C.PangoRectangle
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	C.pango_layout_iter_get_run_extents(this1, &ink_rect1, &logical_rect1)
	var ink_rect2 Rectangle
	var logical_rect2 Rectangle
	ink_rect2 = *(*Rectangle)(unsafe.Pointer(&ink_rect1))
	logical_rect2 = *(*Rectangle)(unsafe.Pointer(&logical_rect1))
	return ink_rect2, logical_rect2
}
func (this0 *LayoutIter) GetRunReadonly() *GlyphItem {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_get_run_readonly(this1)
	var ret2 *GlyphItem
	ret2 = (*GlyphItem)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *LayoutIter) NextChar() bool {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_next_char(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *LayoutIter) NextCluster() bool {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_next_cluster(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *LayoutIter) NextLine() bool {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_next_line(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *LayoutIter) NextRun() bool {
	var this1 *C.PangoLayoutIter
	this1 = (*C.PangoLayoutIter)(unsafe.Pointer(this0))
	ret1 := C.pango_layout_iter_next_run(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type LayoutLine struct {
	layout0 *C.PangoLayout
	StartIndex int32
	Length int32
	runs0 *C.GSList
	IsParagraphStart uint32
	ResolvedDir uint32
}
func (this0 *LayoutLine) Layout() *Layout {
	var layout1 *Layout
	layout1 = (*Layout)(gobject.ObjectWrap(unsafe.Pointer(this0.layout0), true))
	return layout1
}
func (this0 *LayoutLine) Runs() []unsafe.Pointer {
	var runs1 []unsafe.Pointer
	for iter := (*_GSList)(unsafe.Pointer(this0.runs0)); iter != nil; iter = iter.next {
		var elt unsafe.Pointer
		elt = (unsafe.Pointer)(iter.data)
		runs1 = append(runs1, elt)
	}
	return runs1
}
func (this0 *LayoutLine) GetExtents() (Rectangle, Rectangle) {
	var this1 *C.PangoLayoutLine
	var ink_rect1 C.PangoRectangle
	var logical_rect1 C.PangoRectangle
	this1 = (*C.PangoLayoutLine)(unsafe.Pointer(this0))
	C.pango_layout_line_get_extents(this1, &ink_rect1, &logical_rect1)
	var ink_rect2 Rectangle
	var logical_rect2 Rectangle
	ink_rect2 = *(*Rectangle)(unsafe.Pointer(&ink_rect1))
	logical_rect2 = *(*Rectangle)(unsafe.Pointer(&logical_rect1))
	return ink_rect2, logical_rect2
}
func (this0 *LayoutLine) GetPixelExtents() (Rectangle, Rectangle) {
	var this1 *C.PangoLayoutLine
	var ink_rect1 C.PangoRectangle
	var logical_rect1 C.PangoRectangle
	this1 = (*C.PangoLayoutLine)(unsafe.Pointer(this0))
	C.pango_layout_line_get_pixel_extents(this1, &ink_rect1, &logical_rect1)
	var ink_rect2 Rectangle
	var logical_rect2 Rectangle
	ink_rect2 = *(*Rectangle)(unsafe.Pointer(&ink_rect1))
	logical_rect2 = *(*Rectangle)(unsafe.Pointer(&logical_rect1))
	return ink_rect2, logical_rect2
}
func (this0 *LayoutLine) GetXRanges(start_index0 int, end_index0 int) []int {
	var this1 *C.PangoLayoutLine
	var start_index1 C.int32_t
	var end_index1 C.int32_t
	var ranges1 *C.int32_t
	var n_ranges1 C.int32_t
	this1 = (*C.PangoLayoutLine)(unsafe.Pointer(this0))
	start_index1 = C.int32_t(start_index0)
	end_index1 = C.int32_t(end_index0)
	C.pango_layout_line_get_x_ranges(this1, start_index1, end_index1, &ranges1, &n_ranges1)
	var ranges2 []int
	ranges2 = make([]int, n_ranges1)
	for i := range ranges2 {
		ranges2[i] = int((*(*[999999]C.int32_t)(unsafe.Pointer(ranges1)))[i])
	}
	return ranges2
}
func (this0 *LayoutLine) IndexToX(index_0 int, trailing0 bool) int {
	var this1 *C.PangoLayoutLine
	var index_1 C.int32_t
	var trailing1 C.int
	var x_pos1 C.int32_t
	this1 = (*C.PangoLayoutLine)(unsafe.Pointer(this0))
	index_1 = C.int32_t(index_0)
	trailing1 = _GoBoolToCBool(trailing0)
	C.pango_layout_line_index_to_x(this1, index_1, trailing1, &x_pos1)
	var x_pos2 int
	x_pos2 = int(x_pos1)
	return x_pos2
}
func (this0 *LayoutLine) XToIndex(x_pos0 int) (int, int, bool) {
	var this1 *C.PangoLayoutLine
	var x_pos1 C.int32_t
	var index_1 C.int32_t
	var trailing1 C.int32_t
	this1 = (*C.PangoLayoutLine)(unsafe.Pointer(this0))
	x_pos1 = C.int32_t(x_pos0)
	ret1 := C.pango_layout_line_x_to_index(this1, x_pos1, &index_1, &trailing1)
	var index_2 int
	var trailing2 int
	var ret2 bool
	index_2 = int(index_1)
	trailing2 = int(trailing1)
	ret2 = ret1 != 0
	return index_2, trailing2, ret2
}
type LogAttr struct {
	IsLineBreak uint32
	IsMandatoryBreak uint32
	IsCharBreak uint32
	IsWhite uint32
	IsCursorPosition uint32
	IsWordStart uint32
	IsWordEnd uint32
	IsSentenceBoundary uint32
	IsSentenceStart uint32
	IsSentenceEnd uint32
	BackspaceDeletesCharacter uint32
	IsExpandableSpace uint32
	IsWordBoundary uint32
}
type Matrix struct {
	Xx float64
	XY float64
	Yx float64
	Yy float64
	X0 float64
	Y0 float64
}
func (this0 *Matrix) Concat(new_matrix0 *Matrix) {
	var this1 *C.PangoMatrix
	var new_matrix1 *C.PangoMatrix
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	new_matrix1 = (*C.PangoMatrix)(unsafe.Pointer(new_matrix0))
	C.pango_matrix_concat(this1, new_matrix1)
}
func (this0 *Matrix) Copy() *Matrix {
	var this1 *C.PangoMatrix
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	ret1 := C.pango_matrix_copy(this1)
	var ret2 *Matrix
	ret2 = (*Matrix)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Matrix) Free() {
	var this1 *C.PangoMatrix
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	C.pango_matrix_free(this1)
}
func (this0 *Matrix) GetFontScaleFactor() float64 {
	var this1 *C.PangoMatrix
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	ret1 := C.pango_matrix_get_font_scale_factor(this1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *Matrix) Rotate(degrees0 float64) {
	var this1 *C.PangoMatrix
	var degrees1 C.double
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	degrees1 = C.double(degrees0)
	C.pango_matrix_rotate(this1, degrees1)
}
func (this0 *Matrix) Scale(scale_x0 float64, scale_y0 float64) {
	var this1 *C.PangoMatrix
	var scale_x1 C.double
	var scale_y1 C.double
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	scale_x1 = C.double(scale_x0)
	scale_y1 = C.double(scale_y0)
	C.pango_matrix_scale(this1, scale_x1, scale_y1)
}
func (this0 *Matrix) TransformDistance(dx0 float64, dy0 float64) (float64, float64) {
	var this1 *C.PangoMatrix
	var dx1 C.double
	var dy1 C.double
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	dx1 = C.double(dx0)
	dy1 = C.double(dy0)
	C.pango_matrix_transform_distance(this1, &dx1, &dy1)
	var dx2 float64
	var dy2 float64
	dx2 = float64(dx1)
	dy2 = float64(dy1)
	return dx2, dy2
}
func (this0 *Matrix) TransformPixelRectangle(rect0 Rectangle) Rectangle {
	var this1 *C.PangoMatrix
	var rect1 C.PangoRectangle
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	rect1 = *(*C.PangoRectangle)(unsafe.Pointer(&rect0))
	C.pango_matrix_transform_pixel_rectangle(this1, &rect1)
	var rect2 Rectangle
	rect2 = *(*Rectangle)(unsafe.Pointer(&rect1))
	return rect2
}
func (this0 *Matrix) TransformPoint(x0 float64, y0 float64) (float64, float64) {
	var this1 *C.PangoMatrix
	var x1 C.double
	var y1 C.double
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	x1 = C.double(x0)
	y1 = C.double(y0)
	C.pango_matrix_transform_point(this1, &x1, &y1)
	var x2 float64
	var y2 float64
	x2 = float64(x1)
	y2 = float64(y1)
	return x2, y2
}
func (this0 *Matrix) TransformRectangle(rect0 Rectangle) Rectangle {
	var this1 *C.PangoMatrix
	var rect1 C.PangoRectangle
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	rect1 = *(*C.PangoRectangle)(unsafe.Pointer(&rect0))
	C.pango_matrix_transform_rectangle(this1, &rect1)
	var rect2 Rectangle
	rect2 = *(*Rectangle)(unsafe.Pointer(&rect1))
	return rect2
}
func (this0 *Matrix) Translate(tx0 float64, ty0 float64) {
	var this1 *C.PangoMatrix
	var tx1 C.double
	var ty1 C.double
	this1 = (*C.PangoMatrix)(unsafe.Pointer(this0))
	tx1 = C.double(tx0)
	ty1 = C.double(ty0)
	C.pango_matrix_translate(this1, tx1, ty1)
}
const RenderTypeNone = "PangoRenderNone"
type Rectangle struct {
	X int32
	Y int32
	Width int32
	Height int32
}
type RenderPart C.uint32_t
const (
	RenderPartForeground RenderPart = 0
	RenderPartBackground RenderPart = 1
	RenderPartUnderline RenderPart = 2
	RenderPartStrikethrough RenderPart = 3
)
type RendererLike interface {
	gobject.ObjectLike
	InheritedFromPangoRenderer() *C.PangoRenderer
}

type Renderer struct {
	gobject.Object
	
}

func ToRenderer(objlike gobject.ObjectLike) *Renderer {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Renderer)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Renderer)(obj)
	}
	panic("cannot cast to Renderer")
}

func (this0 *Renderer) InheritedFromPangoRenderer() *C.PangoRenderer {
	if this0 == nil {
		return nil
	}
	return (*C.PangoRenderer)(this0.C)
}

func (this0 *Renderer) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_renderer_get_type())
}

func RendererGetType() gobject.Type {
	return (*Renderer)(nil).GetStaticType()
}
func (this0 *Renderer) Activate() {
	var this1 *C.PangoRenderer
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	C.pango_renderer_activate(this1)
}
func (this0 *Renderer) Deactivate() {
	var this1 *C.PangoRenderer
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	C.pango_renderer_deactivate(this1)
}
func (this0 *Renderer) DrawErrorUnderline(x0 int, y0 int, width0 int, height0 int) {
	var this1 *C.PangoRenderer
	var x1 C.int32_t
	var y1 C.int32_t
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	C.pango_renderer_draw_error_underline(this1, x1, y1, width1, height1)
}
func (this0 *Renderer) DrawGlyph(font0 FontLike, glyph0 int, x0 float64, y0 float64) {
	var this1 *C.PangoRenderer
	var font1 *C.PangoFont
	var glyph1 C.uint32_t
	var x1 C.double
	var y1 C.double
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	if font0 != nil {
		font1 = font0.InheritedFromPangoFont()
	}
	glyph1 = C.uint32_t(glyph0)
	x1 = C.double(x0)
	y1 = C.double(y0)
	C.pango_renderer_draw_glyph(this1, font1, glyph1, x1, y1)
}
func (this0 *Renderer) DrawGlyphItem(text0 string, glyph_item0 *GlyphItem, x0 int, y0 int) {
	var this1 *C.PangoRenderer
	var text1 *C.char
	var glyph_item1 *C.PangoGlyphItem
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	glyph_item1 = (*C.PangoGlyphItem)(unsafe.Pointer(glyph_item0))
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	C.pango_renderer_draw_glyph_item(this1, text1, glyph_item1, x1, y1)
}
func (this0 *Renderer) DrawGlyphs(font0 FontLike, glyphs0 *GlyphString, x0 int, y0 int) {
	var this1 *C.PangoRenderer
	var font1 *C.PangoFont
	var glyphs1 *C.PangoGlyphString
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	if font0 != nil {
		font1 = font0.InheritedFromPangoFont()
	}
	glyphs1 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs0))
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	C.pango_renderer_draw_glyphs(this1, font1, glyphs1, x1, y1)
}
func (this0 *Renderer) DrawLayout(layout0 LayoutLike, x0 int, y0 int) {
	var this1 *C.PangoRenderer
	var layout1 *C.PangoLayout
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	if layout0 != nil {
		layout1 = layout0.InheritedFromPangoLayout()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	C.pango_renderer_draw_layout(this1, layout1, x1, y1)
}
func (this0 *Renderer) DrawLayoutLine(line0 *LayoutLine, x0 int, y0 int) {
	var this1 *C.PangoRenderer
	var line1 *C.PangoLayoutLine
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	line1 = (*C.PangoLayoutLine)(unsafe.Pointer(line0))
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	C.pango_renderer_draw_layout_line(this1, line1, x1, y1)
}
func (this0 *Renderer) DrawRectangle(part0 RenderPart, x0 int, y0 int, width0 int, height0 int) {
	var this1 *C.PangoRenderer
	var part1 C.PangoRenderPart
	var x1 C.int32_t
	var y1 C.int32_t
	var width1 C.int32_t
	var height1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	part1 = C.PangoRenderPart(part0)
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	width1 = C.int32_t(width0)
	height1 = C.int32_t(height0)
	C.pango_renderer_draw_rectangle(this1, part1, x1, y1, width1, height1)
}
func (this0 *Renderer) DrawTrapezoid(part0 RenderPart, y1_0 float64, x110 float64, x210 float64, y20 float64, x120 float64, x220 float64) {
	var this1 *C.PangoRenderer
	var part1 C.PangoRenderPart
	var y1_1 C.double
	var x111 C.double
	var x211 C.double
	var y21 C.double
	var x121 C.double
	var x221 C.double
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	part1 = C.PangoRenderPart(part0)
	y1_1 = C.double(y1_0)
	x111 = C.double(x110)
	x211 = C.double(x210)
	y21 = C.double(y20)
	x121 = C.double(x120)
	x221 = C.double(x220)
	C.pango_renderer_draw_trapezoid(this1, part1, y1_1, x111, x211, y21, x121, x221)
}
func (this0 *Renderer) GetColor(part0 RenderPart) *Color {
	var this1 *C.PangoRenderer
	var part1 C.PangoRenderPart
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	part1 = C.PangoRenderPart(part0)
	ret1 := C.pango_renderer_get_color(this1, part1)
	var ret2 *Color
	ret2 = (*Color)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Renderer) GetLayout() *Layout {
	var this1 *C.PangoRenderer
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	ret1 := C.pango_renderer_get_layout(this1)
	var ret2 *Layout
	ret2 = (*Layout)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Renderer) GetLayoutLine() *LayoutLine {
	var this1 *C.PangoRenderer
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	ret1 := C.pango_renderer_get_layout_line(this1)
	var ret2 *LayoutLine
	ret2 = (*LayoutLine)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Renderer) GetMatrix() *Matrix {
	var this1 *C.PangoRenderer
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	ret1 := C.pango_renderer_get_matrix(this1)
	var ret2 *Matrix
	ret2 = (*Matrix)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *Renderer) PartChanged(part0 RenderPart) {
	var this1 *C.PangoRenderer
	var part1 C.PangoRenderPart
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	part1 = C.PangoRenderPart(part0)
	C.pango_renderer_part_changed(this1, part1)
}
func (this0 *Renderer) SetColor(part0 RenderPart, color0 *Color) {
	var this1 *C.PangoRenderer
	var part1 C.PangoRenderPart
	var color1 *C.PangoColor
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	part1 = C.PangoRenderPart(part0)
	color1 = (*C.PangoColor)(unsafe.Pointer(color0))
	C.pango_renderer_set_color(this1, part1, color1)
}
func (this0 *Renderer) SetMatrix(matrix0 *Matrix) {
	var this1 *C.PangoRenderer
	var matrix1 *C.PangoMatrix
	if this0 != nil {
		this1 = this0.InheritedFromPangoRenderer()
	}
	matrix1 = (*C.PangoMatrix)(unsafe.Pointer(matrix0))
	C.pango_renderer_set_matrix(this1, matrix1)
}
const Scale = 1024
type Script C.int32_t
const (
	ScriptInvalidCode Script = -1
	ScriptCommon Script = 0
	ScriptInherited Script = 1
	ScriptArabic Script = 2
	ScriptArmenian Script = 3
	ScriptBengali Script = 4
	ScriptBopomofo Script = 5
	ScriptCherokee Script = 6
	ScriptCoptic Script = 7
	ScriptCyrillic Script = 8
	ScriptDeseret Script = 9
	ScriptDevanagari Script = 10
	ScriptEthiopic Script = 11
	ScriptGeorgian Script = 12
	ScriptGothic Script = 13
	ScriptGreek Script = 14
	ScriptGujarati Script = 15
	ScriptGurmukhi Script = 16
	ScriptHan Script = 17
	ScriptHangul Script = 18
	ScriptHebrew Script = 19
	ScriptHiragana Script = 20
	ScriptKannada Script = 21
	ScriptKatakana Script = 22
	ScriptKhmer Script = 23
	ScriptLao Script = 24
	ScriptLatin Script = 25
	ScriptMalayalam Script = 26
	ScriptMongolian Script = 27
	ScriptMyanmar Script = 28
	ScriptOgham Script = 29
	ScriptOldItalic Script = 30
	ScriptOriya Script = 31
	ScriptRunic Script = 32
	ScriptSinhala Script = 33
	ScriptSyriac Script = 34
	ScriptTamil Script = 35
	ScriptTelugu Script = 36
	ScriptThaana Script = 37
	ScriptThai Script = 38
	ScriptTibetan Script = 39
	ScriptCanadianAboriginal Script = 40
	ScriptYi Script = 41
	ScriptTagalog Script = 42
	ScriptHanunoo Script = 43
	ScriptBuhid Script = 44
	ScriptTagbanwa Script = 45
	ScriptBraille Script = 46
	ScriptCypriot Script = 47
	ScriptLimbu Script = 48
	ScriptOsmanya Script = 49
	ScriptShavian Script = 50
	ScriptLinearB Script = 51
	ScriptTaiLe Script = 52
	ScriptUgaritic Script = 53
	ScriptNewTaiLue Script = 54
	ScriptBuginese Script = 55
	ScriptGlagolitic Script = 56
	ScriptTifinagh Script = 57
	ScriptSylotiNagri Script = 58
	ScriptOldPersian Script = 59
	ScriptKharoshthi Script = 60
	ScriptUnknown Script = 61
	ScriptBalinese Script = 62
	ScriptCuneiform Script = 63
	ScriptPhoenician Script = 64
	ScriptPhagsPa Script = 65
	ScriptNko Script = 66
	ScriptKayahLi Script = 67
	ScriptLepcha Script = 68
	ScriptRejang Script = 69
	ScriptSundanese Script = 70
	ScriptSaurashtra Script = 71
	ScriptCham Script = 72
	ScriptOlChiki Script = 73
	ScriptVai Script = 74
	ScriptCarian Script = 75
	ScriptLycian Script = 76
	ScriptLydian Script = 77
)
type ScriptIter struct {}
func (this0 *ScriptIter) Free() {
	var this1 *C.PangoScriptIter
	this1 = (*C.PangoScriptIter)(unsafe.Pointer(this0))
	C.pango_script_iter_free(this1)
}
func (this0 *ScriptIter) GetRange() (string, string, Script) {
	var this1 *C.PangoScriptIter
	var start1 *C.char
	var end1 *C.char
	var script1 C.PangoScript
	this1 = (*C.PangoScriptIter)(unsafe.Pointer(this0))
	C.pango_script_iter_get_range(this1, &start1, &end1, &script1)
	var start2 string
	var end2 string
	var script2 Script
	start2 = C.GoString(start1)
	C.g_free(unsafe.Pointer(start1))
	end2 = C.GoString(end1)
	C.g_free(unsafe.Pointer(end1))
	script2 = Script(script1)
	return start2, end2, script2
}
func (this0 *ScriptIter) Next() bool {
	var this1 *C.PangoScriptIter
	this1 = (*C.PangoScriptIter)(unsafe.Pointer(this0))
	ret1 := C.pango_script_iter_next(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
type Stretch C.uint32_t
const (
	StretchUltraCondensed Stretch = 0
	StretchExtraCondensed Stretch = 1
	StretchCondensed Stretch = 2
	StretchSemiCondensed Stretch = 3
	StretchNormal Stretch = 4
	StretchSemiExpanded Stretch = 5
	StretchExpanded Stretch = 6
	StretchExtraExpanded Stretch = 7
	StretchUltraExpanded Stretch = 8
)
type Style C.uint32_t
const (
	StyleNormal Style = 0
	StyleOblique Style = 1
	StyleItalic Style = 2
)
type TabAlign C.uint32_t
const (
	TabAlignLeft TabAlign = 0
)
type TabArray struct {}
func NewTabArray(initial_size0 int, positions_in_pixels0 bool) *TabArray {
	var initial_size1 C.int32_t
	var positions_in_pixels1 C.int
	initial_size1 = C.int32_t(initial_size0)
	positions_in_pixels1 = _GoBoolToCBool(positions_in_pixels0)
	ret1 := C.pango_tab_array_new(initial_size1, positions_in_pixels1)
	var ret2 *TabArray
	ret2 = (*TabArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *TabArray) Copy() *TabArray {
	var this1 *C.PangoTabArray
	this1 = (*C.PangoTabArray)(unsafe.Pointer(this0))
	ret1 := C.pango_tab_array_copy(this1)
	var ret2 *TabArray
	ret2 = (*TabArray)(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *TabArray) Free() {
	var this1 *C.PangoTabArray
	this1 = (*C.PangoTabArray)(unsafe.Pointer(this0))
	C.pango_tab_array_free(this1)
}
func (this0 *TabArray) GetPositionsInPixels() bool {
	var this1 *C.PangoTabArray
	this1 = (*C.PangoTabArray)(unsafe.Pointer(this0))
	ret1 := C.pango_tab_array_get_positions_in_pixels(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *TabArray) GetSize() int {
	var this1 *C.PangoTabArray
	this1 = (*C.PangoTabArray)(unsafe.Pointer(this0))
	ret1 := C.pango_tab_array_get_size(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *TabArray) GetTab(tab_index0 int) (TabAlign, int) {
	var this1 *C.PangoTabArray
	var tab_index1 C.int32_t
	var alignment1 C.PangoTabAlign
	var location1 C.int32_t
	this1 = (*C.PangoTabArray)(unsafe.Pointer(this0))
	tab_index1 = C.int32_t(tab_index0)
	C.pango_tab_array_get_tab(this1, tab_index1, &alignment1, &location1)
	var alignment2 TabAlign
	var location2 int
	alignment2 = TabAlign(alignment1)
	location2 = int(location1)
	return alignment2, location2
}
func (this0 *TabArray) Resize(new_size0 int) {
	var this1 *C.PangoTabArray
	var new_size1 C.int32_t
	this1 = (*C.PangoTabArray)(unsafe.Pointer(this0))
	new_size1 = C.int32_t(new_size0)
	C.pango_tab_array_resize(this1, new_size1)
}
func (this0 *TabArray) SetTab(tab_index0 int, alignment0 TabAlign, location0 int) {
	var this1 *C.PangoTabArray
	var tab_index1 C.int32_t
	var alignment1 C.PangoTabAlign
	var location1 C.int32_t
	this1 = (*C.PangoTabArray)(unsafe.Pointer(this0))
	tab_index1 = C.int32_t(tab_index0)
	alignment1 = C.PangoTabAlign(alignment0)
	location1 = C.int32_t(location0)
	C.pango_tab_array_set_tab(this1, tab_index1, alignment1, location1)
}
const UnknownGlyphHeight = 14
const UnknownGlyphWidth = 10
type Underline C.uint32_t
const (
	UnderlineNone Underline = 0
	UnderlineSingle Underline = 1
	UnderlineDouble Underline = 2
	UnderlineLow Underline = 3
	UnderlineError Underline = 4
)
type Variant C.uint32_t
const (
	VariantNormal Variant = 0
	VariantSmallCaps Variant = 1
)
type Weight C.uint32_t
const (
	WeightThin Weight = 100
	WeightUltralight Weight = 200
	WeightLight Weight = 300
	WeightBook Weight = 380
	WeightNormal Weight = 400
	WeightMedium Weight = 500
	WeightSemibold Weight = 600
	WeightBold Weight = 700
	WeightUltrabold Weight = 800
	WeightHeavy Weight = 900
	WeightUltraheavy Weight = 1000
)
type WrapMode C.uint32_t
const (
	WrapModeWord WrapMode = 0
	WrapModeChar WrapMode = 1
	WrapModeWordChar WrapMode = 2
)
type _ScriptForLang struct {
	Lang [7]uint8
	_ [1]byte
	Scripts [3]Script
}
func AttrTypeGetName(type0 AttrType) string {
	var type1 C.PangoAttrType
	type1 = C.PangoAttrType(type0)
	ret1 := C.pango_attr_type_get_name(type1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func AttrTypeRegister(name0 string) AttrType {
	var name1 *C.char
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.pango_attr_type_register(name1)
	var ret2 AttrType
	ret2 = AttrType(ret1)
	return ret2
}
func BidiTypeForUnichar(ch0 rune) BidiType {
	var ch1 C.uint32_t
	ch1 = C.uint32_t(ch0)
	ret1 := C.pango_bidi_type_for_unichar(ch1)
	var ret2 BidiType
	ret2 = BidiType(ret1)
	return ret2
}
func Break(text0 string, length0 int, analysis0 *Analysis, attrs0 *LogAttr, attrs_len0 int) {
	var text1 *C.char
	var length1 C.int32_t
	var analysis1 *C.PangoAnalysis
	var attrs1 *C.PangoLogAttr
	var attrs_len1 C.int32_t
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	length1 = C.int32_t(length0)
	analysis1 = (*C.PangoAnalysis)(unsafe.Pointer(analysis0))
	attrs1 = (*C.PangoLogAttr)(unsafe.Pointer(attrs0))
	attrs_len1 = C.int32_t(attrs_len0)
	C.pango_break(text1, length1, analysis1, attrs1, attrs_len1)
}
func ExtentsToPixels(inclusive0 *Rectangle, nearest0 *Rectangle) {
	var inclusive1 *C.PangoRectangle
	var nearest1 *C.PangoRectangle
	inclusive1 = (*C.PangoRectangle)(unsafe.Pointer(inclusive0))
	nearest1 = (*C.PangoRectangle)(unsafe.Pointer(nearest0))
	C.pango_extents_to_pixels(inclusive1, nearest1)
}
func FindBaseDir(text0 string, length0 int) Direction {
	var text1 *C.char
	var length1 C.int32_t
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	length1 = C.int32_t(length0)
	ret1 := C.pango_find_base_dir(text1, length1)
	var ret2 Direction
	ret2 = Direction(ret1)
	return ret2
}
// blacklisted: find_paragraph_boundary (function)
// blacklisted: font_description_from_string (function)
func GetLogAttrs(text0 string, length0 int, level0 int, language0 *Language, log_attrs0 *LogAttr, attrs_len0 int) {
	var text1 *C.char
	var length1 C.int32_t
	var level1 C.int32_t
	var language1 *C.PangoLanguage
	var log_attrs1 *C.PangoLogAttr
	var attrs_len1 C.int32_t
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	length1 = C.int32_t(length0)
	level1 = C.int32_t(level0)
	language1 = (*C.PangoLanguage)(unsafe.Pointer(language0))
	log_attrs1 = (*C.PangoLogAttr)(unsafe.Pointer(log_attrs0))
	attrs_len1 = C.int32_t(attrs_len0)
	C.pango_get_log_attrs(text1, length1, level1, language1, log_attrs1, attrs_len1)
}
// blacklisted: get_mirror_char (function)
func GravityGetForMatrix(matrix0 *Matrix) Gravity {
	var matrix1 *C.PangoMatrix
	matrix1 = (*C.PangoMatrix)(unsafe.Pointer(matrix0))
	ret1 := C.pango_gravity_get_for_matrix(matrix1)
	var ret2 Gravity
	ret2 = Gravity(ret1)
	return ret2
}
func GravityGetForScript(script0 Script, base_gravity0 Gravity, hint0 GravityHint) Gravity {
	var script1 C.PangoScript
	var base_gravity1 C.PangoGravity
	var hint1 C.PangoGravityHint
	script1 = C.PangoScript(script0)
	base_gravity1 = C.PangoGravity(base_gravity0)
	hint1 = C.PangoGravityHint(hint0)
	ret1 := C.pango_gravity_get_for_script(script1, base_gravity1, hint1)
	var ret2 Gravity
	ret2 = Gravity(ret1)
	return ret2
}
func GravityGetForScriptAndWidth(script0 Script, wide0 bool, base_gravity0 Gravity, hint0 GravityHint) Gravity {
	var script1 C.PangoScript
	var wide1 C.int
	var base_gravity1 C.PangoGravity
	var hint1 C.PangoGravityHint
	script1 = C.PangoScript(script0)
	wide1 = _GoBoolToCBool(wide0)
	base_gravity1 = C.PangoGravity(base_gravity0)
	hint1 = C.PangoGravityHint(hint0)
	ret1 := C.pango_gravity_get_for_script_and_width(script1, wide1, base_gravity1, hint1)
	var ret2 Gravity
	ret2 = Gravity(ret1)
	return ret2
}
func GravityToRotation(gravity0 Gravity) float64 {
	var gravity1 C.PangoGravity
	gravity1 = C.PangoGravity(gravity0)
	ret1 := C.pango_gravity_to_rotation(gravity1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func IsZeroWidth(ch0 rune) bool {
	var ch1 C.uint32_t
	ch1 = C.uint32_t(ch0)
	ret1 := C.pango_is_zero_width(ch1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func Itemize(context0 ContextLike, text0 string, start_index0 int, length0 int, attrs0 *AttrList, cached_iter0 *AttrIterator) []Item {
	var context1 *C.PangoContext
	var text1 *C.char
	var start_index1 C.int32_t
	var length1 C.int32_t
	var attrs1 *C.PangoAttrList
	var cached_iter1 *C.PangoAttrIterator
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	start_index1 = C.int32_t(start_index0)
	length1 = C.int32_t(length0)
	attrs1 = (*C.PangoAttrList)(unsafe.Pointer(attrs0))
	cached_iter1 = (*C.PangoAttrIterator)(unsafe.Pointer(cached_iter0))
	ret1 := C.pango_itemize(context1, text1, start_index1, length1, attrs1, cached_iter1)
	var ret2 []Item
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt Item
		elt = *(*Item)(unsafe.Pointer((*C.PangoItem)(iter.data)))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func ItemizeWithBaseDir(context0 ContextLike, base_dir0 Direction, text0 string, start_index0 int, length0 int, attrs0 *AttrList, cached_iter0 *AttrIterator) []Item {
	var context1 *C.PangoContext
	var base_dir1 C.PangoDirection
	var text1 *C.char
	var start_index1 C.int32_t
	var length1 C.int32_t
	var attrs1 *C.PangoAttrList
	var cached_iter1 *C.PangoAttrIterator
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	base_dir1 = C.PangoDirection(base_dir0)
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	start_index1 = C.int32_t(start_index0)
	length1 = C.int32_t(length0)
	attrs1 = (*C.PangoAttrList)(unsafe.Pointer(attrs0))
	cached_iter1 = (*C.PangoAttrIterator)(unsafe.Pointer(cached_iter0))
	ret1 := C.pango_itemize_with_base_dir(context1, base_dir1, text1, start_index1, length1, attrs1, cached_iter1)
	var ret2 []Item
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt Item
		elt = *(*Item)(unsafe.Pointer((*C.PangoItem)(iter.data)))
		ret2 = append(ret2, elt)
	}
	return ret2
}
// blacklisted: language_from_string (function)
// blacklisted: language_get_default (function)
// blacklisted: log2vis_get_embedding_levels (function)
func ParseEnum(type0 gobject.Type, str0 string, warn0 bool) (int, string, bool) {
	var type1 C.GType
	var str1 *C.char
	var warn1 C.int
	var value1 C.int32_t
	var possible_values1 *C.char
	type1 = C.GType(type0)
	str1 = _GoStringToGString(str0)
	defer C.free(unsafe.Pointer(str1))
	warn1 = _GoBoolToCBool(warn0)
	ret1 := C.pango_parse_enum(type1, str1, &value1, warn1, &possible_values1)
	var value2 int
	var possible_values2 string
	var ret2 bool
	value2 = int(value1)
	possible_values2 = C.GoString(possible_values1)
	C.g_free(unsafe.Pointer(possible_values1))
	ret2 = ret1 != 0
	return value2, possible_values2, ret2
}
func ParseMarkup(markup_text0 string, length0 int, accel_marker0 rune) (*AttrList, string, rune, bool, error) {
	var markup_text1 *C.char
	var length1 C.int32_t
	var accel_marker1 C.uint32_t
	var attr_list1 *C.PangoAttrList
	var text1 *C.char
	var accel_char1 C.uint32_t
	var err1 *C.GError
	markup_text1 = _GoStringToGString(markup_text0)
	defer C.free(unsafe.Pointer(markup_text1))
	length1 = C.int32_t(length0)
	accel_marker1 = C.uint32_t(accel_marker0)
	ret1 := C.pango_parse_markup(markup_text1, length1, accel_marker1, &attr_list1, &text1, &accel_char1, &err1)
	var attr_list2 *AttrList
	var text2 string
	var accel_char2 rune
	var ret2 bool
	var err2 error
	attr_list2 = (*AttrList)(unsafe.Pointer(attr_list1))
	text2 = C.GoString(text1)
	C.g_free(unsafe.Pointer(text1))
	accel_char2 = rune(accel_char1)
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return attr_list2, text2, accel_char2, ret2, err2
}
// blacklisted: parse_stretch (function)
// blacklisted: parse_style (function)
// blacklisted: parse_variant (function)
// blacklisted: parse_weight (function)
func QuantizeLineGeometry(thickness0 int, position0 int) (int, int) {
	var thickness1 C.int32_t
	var position1 C.int32_t
	thickness1 = C.int32_t(thickness0)
	position1 = C.int32_t(position0)
	C.pango_quantize_line_geometry(&thickness1, &position1)
	var thickness2 int
	var position2 int
	thickness2 = int(thickness1)
	position2 = int(position1)
	return thickness2, position2
}
// blacklisted: read_line (function)
func ScanInt(pos0 string) (string, int, bool) {
	var pos1 *C.char
	var out1 C.int32_t
	pos1 = _GoStringToGString(pos0)
	defer C.free(unsafe.Pointer(pos1))
	ret1 := C.pango_scan_int(&pos1, &out1)
	var pos2 string
	var out2 int
	var ret2 bool
	pos2 = C.GoString(pos1)
	out2 = int(out1)
	ret2 = ret1 != 0
	return pos2, out2, ret2
}
// blacklisted: scan_string (function)
// blacklisted: scan_word (function)
func ScriptForUnichar(ch0 rune) Script {
	var ch1 C.uint32_t
	ch1 = C.uint32_t(ch0)
	ret1 := C.pango_script_for_unichar(ch1)
	var ret2 Script
	ret2 = Script(ret1)
	return ret2
}
func ScriptGetSampleLanguage(script0 Script) *Language {
	var script1 C.PangoScript
	script1 = C.PangoScript(script0)
	ret1 := C.pango_script_get_sample_language(script1)
	var ret2 *Language
	ret2 = (*Language)(unsafe.Pointer(ret1))
	return ret2
}
func Shape(text0 string, length0 int, analysis0 *Analysis, glyphs0 *GlyphString) {
	var text1 *C.char
	var length1 C.int32_t
	var analysis1 *C.PangoAnalysis
	var glyphs1 *C.PangoGlyphString
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	length1 = C.int32_t(length0)
	analysis1 = (*C.PangoAnalysis)(unsafe.Pointer(analysis0))
	glyphs1 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs0))
	C.pango_shape(text1, length1, analysis1, glyphs1)
}
func SkipSpace(pos0 string) (string, bool) {
	var pos1 *C.char
	pos1 = _GoStringToGString(pos0)
	defer C.free(unsafe.Pointer(pos1))
	ret1 := C.pango_skip_space(&pos1)
	var pos2 string
	var ret2 bool
	pos2 = C.GoString(pos1)
	ret2 = ret1 != 0
	return pos2, ret2
}
func SplitFileList(str0 string) []string {
	var str1 *C.char
	str1 = _GoStringToGString(str0)
	defer C.free(unsafe.Pointer(str1))
	ret1 := C.pango_split_file_list(str1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func TrimString(str0 string) string {
	var str1 *C.char
	str1 = _GoStringToGString(str0)
	defer C.free(unsafe.Pointer(str1))
	ret1 := C.pango_trim_string(str1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func UnicharDirection(ch0 rune) Direction {
	var ch1 C.uint32_t
	ch1 = C.uint32_t(ch0)
	ret1 := C.pango_unichar_direction(ch1)
	var ret2 Direction
	ret2 = Direction(ret1)
	return ret2
}
func UnitsFromDouble(d0 float64) int {
	var d1 C.double
	d1 = C.double(d0)
	ret1 := C.pango_units_from_double(d1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func UnitsToDouble(i0 int) float64 {
	var i1 C.int32_t
	i1 = C.int32_t(i0)
	ret1 := C.pango_units_to_double(i1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func Version() int {
	ret1 := C.pango_version()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func VersionCheck(required_major0 int, required_minor0 int, required_micro0 int) string {
	var required_major1 C.int32_t
	var required_minor1 C.int32_t
	var required_micro1 C.int32_t
	required_major1 = C.int32_t(required_major0)
	required_minor1 = C.int32_t(required_minor0)
	required_micro1 = C.int32_t(required_micro0)
	ret1 := C.pango_version_check(required_major1, required_minor1, required_micro1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func VersionString() string {
	ret1 := C.pango_version_string()
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
