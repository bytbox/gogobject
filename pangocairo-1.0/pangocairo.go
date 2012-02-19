package pangocairo

/*
#include <stdlib.h>
#include <stdint.h>

typedef size_t GType;

extern void _PangoCairo_go_callback_cleanup(void *gofunc);
static void _c_callback_cleanup(void *userdata)
{
	_PangoCairo_go_callback_cleanup(userdata);
}
typedef struct _Bitmap Bitmap;
struct _Bitmap {};
typedef struct _Face Face;
struct _Face {};
typedef struct _Library Library;
struct _Library {};
typedef struct _Pattern Pattern;
struct _Pattern {};
typedef struct _CharSet CharSet;
struct _CharSet {};
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
typedef struct _PangoFT2FontMap PangoFT2FontMap;
typedef void* PangoFT2SubstituteFunc;
extern void _PangoFT2SubstituteFunc_c_wrapper();
extern void _PangoFT2SubstituteFunc_c_wrapper_once();
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
typedef struct _PangoCairoFcFontMap PangoCairoFcFontMap;
typedef struct _PangoCairoFont PangoCairoFont;
typedef struct _PangoCairoFontMap PangoCairoFontMap;
typedef void* PangoCairoShapeRendererFunc;
extern void _PangoCairoShapeRendererFunc_c_wrapper();
extern void _PangoCairoShapeRendererFunc_c_wrapper_once();
extern GType pango_cairo_fc_font_map_get_type();
extern cairoScaledFont* pango_cairo_font_get_scaled_font(PangoCairoFont*);
extern GType pango_cairo_font_get_type();
extern PangoFontMap* pango_cairo_font_map_get_default();
extern PangoFontMap* pango_cairo_font_map_new();
extern PangoFontMap* pango_cairo_font_map_new_for_font_type(cairoFontType);
extern cairoFontType pango_cairo_font_map_get_font_type(PangoCairoFontMap*);
extern double pango_cairo_font_map_get_resolution(PangoCairoFontMap*);
extern void pango_cairo_font_map_set_default(PangoCairoFontMap*);
extern void pango_cairo_font_map_set_resolution(PangoCairoFontMap*, double);
extern GType pango_cairo_font_map_get_type();
extern cairoFontOptions* pango_cairo_context_get_font_options(PangoContext*);
extern double pango_cairo_context_get_resolution(PangoContext*);
extern void pango_cairo_context_set_font_options(PangoContext*, cairoFontOptions*);
extern void pango_cairo_context_set_resolution(PangoContext*, double);
extern void pango_cairo_context_set_shape_renderer(PangoContext*, PangoCairoShapeRendererFunc, void*, GDestroyNotify);
#pragma GCC diagnostic ignored "-Wunused-function"
static void _pango_cairo_context_set_shape_renderer(PangoContext* arg0, void* gofunc) {
	if (gofunc) {
		pango_cairo_context_set_shape_renderer(arg0, _PangoCairoShapeRendererFunc_c_wrapper, gofunc, _c_callback_cleanup);
	} else {
		pango_cairo_context_set_shape_renderer(arg0, 0, 0, 0);
	}
}
extern PangoContext* pango_cairo_create_context(cairoContext*);
extern PangoLayout* pango_cairo_create_layout(cairoContext*);
extern void pango_cairo_error_underline_path(cairoContext*, double, double, double, double);
extern void pango_cairo_glyph_string_path(cairoContext*, PangoFont*, PangoGlyphString*);
extern void pango_cairo_layout_line_path(cairoContext*, PangoLayoutLine*);
extern void pango_cairo_layout_path(cairoContext*, PangoLayout*);
extern void pango_cairo_show_error_underline(cairoContext*, double, double, double, double);
extern void pango_cairo_show_glyph_item(cairoContext*, char*, PangoGlyphItem*);
extern void pango_cairo_show_glyph_string(cairoContext*, PangoFont*, PangoGlyphString*);
extern void pango_cairo_show_layout(cairoContext*, PangoLayout*);
extern void pango_cairo_show_layout_line(cairoContext*, PangoLayoutLine*);
extern void pango_cairo_update_context(cairoContext*, PangoContext*);
extern void pango_cairo_update_layout(cairoContext*, PangoLayout*);


extern GObject *g_object_ref_sink(GObject*);
extern void g_object_unref(GObject*);
extern void g_free(void*);
extern void g_error_free(GError*);

#cgo pkg-config: pangocairo
*/
import "C"
import "unsafe"

// package dependencies
import (
	"github.com/bytbox/gogobject/cairo-1.0"
	"github.com/bytbox/gogobject/gobject-2.0"
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


//export _PangoCairo_go_callback_cleanup
func _PangoCairo_go_callback_cleanup(gofunc unsafe.Pointer) {
	gobject.Holder.Release(gofunc)
}


type FcFontMapLike interface {
	pango.FontMapLike
	InheritedFromPangoCairoFcFontMap() *C.PangoCairoFcFontMap
}

type FcFontMap struct {
	pango.FontMap
	FontMapImpl
}

func ToFcFontMap(objlike gobject.ObjectLike) *FcFontMap {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*FcFontMap)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FcFontMap)(obj)
	}
	panic("cannot cast to FcFontMap")
}

func (this0 *FcFontMap) InheritedFromPangoCairoFcFontMap() *C.PangoCairoFcFontMap {
	if this0 == nil {
		return nil
	}
	return (*C.PangoCairoFcFontMap)(this0.C)
}

func (this0 *FcFontMap) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_cairo_fc_font_map_get_type())
}

func FcFontMapGetType() gobject.Type {
	return (*FcFontMap)(nil).GetStaticType()
}
type FontLike interface {
	ImplementsPangoCairoFont() *C.PangoCairoFont
}

type Font struct {
	gobject.Object
	FontImpl
}

type FontImpl struct {}

func ToFont(objlike gobject.ObjectLike) *Font {
	t := (*FontImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Font)(obj)
	}
	panic("cannot cast to Font")
}

func (this0 *FontImpl) ImplementsPangoCairoFont() *C.PangoCairoFont {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.PangoCairoFont)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *FontImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_cairo_font_get_type())
}

func FontGetType() gobject.Type {
	return (*FontImpl)(nil).GetStaticType()
}
func (this0 *FontImpl) GetScaledFont() *cairo.ScaledFont {
	var this1 *C.PangoCairoFont
	if this0 != nil {
		this1 = this0.ImplementsPangoCairoFont()}
	ret1 := C.pango_cairo_font_get_scaled_font(this1)
	var ret2 *cairo.ScaledFont
	ret2 = (*cairo.ScaledFont)(cairo.ScaledFontWrap(unsafe.Pointer(ret1), false))
	return ret2
}
type FontMapLike interface {
	ImplementsPangoCairoFontMap() *C.PangoCairoFontMap
}

type FontMap struct {
	gobject.Object
	FontMapImpl
}

type FontMapImpl struct {}

func ToFontMap(objlike gobject.ObjectLike) *FontMap {
	t := (*FontMapImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*FontMap)(obj)
	}
	panic("cannot cast to FontMap")
}

func (this0 *FontMapImpl) ImplementsPangoCairoFontMap() *C.PangoCairoFontMap {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.PangoCairoFontMap)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *FontMapImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.pango_cairo_font_map_get_type())
}

func FontMapGetType() gobject.Type {
	return (*FontMapImpl)(nil).GetStaticType()
}
func FontMapGetDefault() *pango.FontMap {
	ret1 := C.pango_cairo_font_map_get_default()
	var ret2 *pango.FontMap
	ret2 = (*pango.FontMap)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func FontMapNew() *pango.FontMap {
	ret1 := C.pango_cairo_font_map_new()
	var ret2 *pango.FontMap
	ret2 = (*pango.FontMap)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
// blacklisted: FontMap.new_for_font_type (method)
// blacklisted: FontMap.get_font_type (method)
func (this0 *FontMapImpl) GetResolution() float64 {
	var this1 *C.PangoCairoFontMap
	if this0 != nil {
		this1 = this0.ImplementsPangoCairoFontMap()}
	ret1 := C.pango_cairo_font_map_get_resolution(this1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *FontMapImpl) SetDefault() {
	var this1 *C.PangoCairoFontMap
	if this0 != nil {
		this1 = this0.ImplementsPangoCairoFontMap()}
	C.pango_cairo_font_map_set_default(this1)
}
func (this0 *FontMapImpl) SetResolution(dpi0 float64) {
	var this1 *C.PangoCairoFontMap
	var dpi1 C.double
	if this0 != nil {
		this1 = this0.ImplementsPangoCairoFontMap()}
	dpi1 = C.double(dpi0)
	C.pango_cairo_font_map_set_resolution(this1, dpi1)
}
type ShapeRendererFunc func(cr *cairo.Context, attr *pango.AttrShape, do_path bool)
//export _PangoCairoShapeRendererFunc_c_wrapper
func _PangoCairoShapeRendererFunc_c_wrapper(cr0 unsafe.Pointer, attr0 unsafe.Pointer, do_path0 int32, data0 unsafe.Pointer) {
	var cr1 *cairo.Context
	var attr1 *pango.AttrShape
	var do_path1 bool
	var data1 ShapeRendererFunc
	cr1 = (*cairo.Context)(cairo.ContextWrap(unsafe.Pointer((*C.cairoContext)(cr0)), true))
	attr1 = (*pango.AttrShape)(unsafe.Pointer((*C.PangoAttrShape)(attr0)))
	do_path1 = (C.int)(do_path0) != 0
	data1 = *(*ShapeRendererFunc)(data0)
	data1(cr1, attr1, do_path1)
}
//export _PangoCairoShapeRendererFunc_c_wrapper_once
func _PangoCairoShapeRendererFunc_c_wrapper_once(cr0 unsafe.Pointer, attr0 unsafe.Pointer, do_path0 int32, data0 unsafe.Pointer) {
	_PangoCairoShapeRendererFunc_c_wrapper(cr0, attr0, do_path0, data0)
	gobject.Holder.Release(data0)
}
func ContextGetFontOptions(context0 pango.ContextLike) *cairo.FontOptions {
	var context1 *C.PangoContext
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	ret1 := C.pango_cairo_context_get_font_options(context1)
	var ret2 *cairo.FontOptions
	ret2 = (*cairo.FontOptions)(cairo.FontOptionsWrap(unsafe.Pointer(ret1)))
	return ret2
}
func ContextGetResolution(context0 pango.ContextLike) float64 {
	var context1 *C.PangoContext
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	ret1 := C.pango_cairo_context_get_resolution(context1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func ContextSetFontOptions(context0 pango.ContextLike, options0 *cairo.FontOptions) {
	var context1 *C.PangoContext
	var options1 *C.cairoFontOptions
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	if options0 != nil {
		options1 = (*C.cairoFontOptions)(options0.C)
	}
	C.pango_cairo_context_set_font_options(context1, options1)
}
func ContextSetResolution(context0 pango.ContextLike, dpi0 float64) {
	var context1 *C.PangoContext
	var dpi1 C.double
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	dpi1 = C.double(dpi0)
	C.pango_cairo_context_set_resolution(context1, dpi1)
}
func ContextSetShapeRenderer(context0 pango.ContextLike, func0 ShapeRendererFunc) {
	var context1 *C.PangoContext
	var func1 unsafe.Pointer
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	if func0 != nil {
		func1 = unsafe.Pointer(&func0)}
	gobject.Holder.Grab(func1)
	C._pango_cairo_context_set_shape_renderer(context1, func1)
}
func CreateContext(cr0 *cairo.Context) *pango.Context {
	var cr1 *C.cairoContext
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	ret1 := C.pango_cairo_create_context(cr1)
	var ret2 *pango.Context
	ret2 = (*pango.Context)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func CreateLayout(cr0 *cairo.Context) *pango.Layout {
	var cr1 *C.cairoContext
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	ret1 := C.pango_cairo_create_layout(cr1)
	var ret2 *pango.Layout
	ret2 = (*pango.Layout)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func ErrorUnderlinePath(cr0 *cairo.Context, x0 float64, y0 float64, width0 float64, height0 float64) {
	var cr1 *C.cairoContext
	var x1 C.double
	var y1 C.double
	var width1 C.double
	var height1 C.double
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	x1 = C.double(x0)
	y1 = C.double(y0)
	width1 = C.double(width0)
	height1 = C.double(height0)
	C.pango_cairo_error_underline_path(cr1, x1, y1, width1, height1)
}
// blacklisted: font_map_get_default (function)
// blacklisted: font_map_new (function)
// blacklisted: font_map_new_for_font_type (function)
func GlyphStringPath(cr0 *cairo.Context, font0 pango.FontLike, glyphs0 *pango.GlyphString) {
	var cr1 *C.cairoContext
	var font1 *C.PangoFont
	var glyphs1 *C.PangoGlyphString
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	if font0 != nil {
		font1 = font0.InheritedFromPangoFont()
	}
	glyphs1 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs0))
	C.pango_cairo_glyph_string_path(cr1, font1, glyphs1)
}
func LayoutLinePath(cr0 *cairo.Context, line0 *pango.LayoutLine) {
	var cr1 *C.cairoContext
	var line1 *C.PangoLayoutLine
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	line1 = (*C.PangoLayoutLine)(unsafe.Pointer(line0))
	C.pango_cairo_layout_line_path(cr1, line1)
}
func LayoutPath(cr0 *cairo.Context, layout0 pango.LayoutLike) {
	var cr1 *C.cairoContext
	var layout1 *C.PangoLayout
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	if layout0 != nil {
		layout1 = layout0.InheritedFromPangoLayout()
	}
	C.pango_cairo_layout_path(cr1, layout1)
}
func ShowErrorUnderline(cr0 *cairo.Context, x0 float64, y0 float64, width0 float64, height0 float64) {
	var cr1 *C.cairoContext
	var x1 C.double
	var y1 C.double
	var width1 C.double
	var height1 C.double
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	x1 = C.double(x0)
	y1 = C.double(y0)
	width1 = C.double(width0)
	height1 = C.double(height0)
	C.pango_cairo_show_error_underline(cr1, x1, y1, width1, height1)
}
func ShowGlyphItem(cr0 *cairo.Context, text0 string, glyph_item0 *pango.GlyphItem) {
	var cr1 *C.cairoContext
	var text1 *C.char
	var glyph_item1 *C.PangoGlyphItem
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	glyph_item1 = (*C.PangoGlyphItem)(unsafe.Pointer(glyph_item0))
	C.pango_cairo_show_glyph_item(cr1, text1, glyph_item1)
}
func ShowGlyphString(cr0 *cairo.Context, font0 pango.FontLike, glyphs0 *pango.GlyphString) {
	var cr1 *C.cairoContext
	var font1 *C.PangoFont
	var glyphs1 *C.PangoGlyphString
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	if font0 != nil {
		font1 = font0.InheritedFromPangoFont()
	}
	glyphs1 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs0))
	C.pango_cairo_show_glyph_string(cr1, font1, glyphs1)
}
func ShowLayout(cr0 *cairo.Context, layout0 pango.LayoutLike) {
	var cr1 *C.cairoContext
	var layout1 *C.PangoLayout
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	if layout0 != nil {
		layout1 = layout0.InheritedFromPangoLayout()
	}
	C.pango_cairo_show_layout(cr1, layout1)
}
func ShowLayoutLine(cr0 *cairo.Context, line0 *pango.LayoutLine) {
	var cr1 *C.cairoContext
	var line1 *C.PangoLayoutLine
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	line1 = (*C.PangoLayoutLine)(unsafe.Pointer(line0))
	C.pango_cairo_show_layout_line(cr1, line1)
}
func UpdateContext(cr0 *cairo.Context, context0 pango.ContextLike) {
	var cr1 *C.cairoContext
	var context1 *C.PangoContext
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	if context0 != nil {
		context1 = context0.InheritedFromPangoContext()
	}
	C.pango_cairo_update_context(cr1, context1)
}
func UpdateLayout(cr0 *cairo.Context, layout0 pango.LayoutLike) {
	var cr1 *C.cairoContext
	var layout1 *C.PangoLayout
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	if layout0 != nil {
		layout1 = layout0.InheritedFromPangoLayout()
	}
	C.pango_cairo_update_layout(cr1, layout1)
}
