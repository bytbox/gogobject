package gtksource

/*
#include <stdlib.h>
#include <stdint.h>

typedef size_t GType;

typedef struct _Display Display;
struct _Display {};
typedef struct _Screen Screen;
struct _Screen {};
typedef struct _Visual Visual;
struct _Visual {};
typedef struct _XEvent XEvent;
struct _XEvent {};
typedef struct _XConfigureEvent XConfigureEvent;
struct _XConfigureEvent {};
typedef struct _XImage XImage;
struct _XImage {};
typedef struct _XFontStruct XFontStruct;
struct _XFontStruct {};
typedef struct _XTrapezoid XTrapezoid;
struct _XTrapezoid {};
typedef struct _XVisualInfo XVisualInfo;
struct _XVisualInfo {};
typedef struct _XWindowAttributes XWindowAttributes;
struct _XWindowAttributes {};
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
typedef struct _GtkAboutDialog GtkAboutDialog;
typedef struct _GtkAboutDialogClass GtkAboutDialogClass;
struct _GtkAboutDialogClass { uint8_t _data[1160]; };
typedef struct _GtkAboutDialogPrivate GtkAboutDialogPrivate;
struct _GtkAboutDialogPrivate {};
typedef uint32_t GtkAccelFlags;
typedef struct _GtkAccelGroup GtkAccelGroup;
typedef void* GtkAccelGroupActivate;
extern void _GtkAccelGroupActivate_c_wrapper();
extern void _GtkAccelGroupActivate_c_wrapper_once();
typedef struct _GtkAccelGroupClass GtkAccelGroupClass;
struct _GtkAccelGroupClass { uint8_t _data[176]; };
typedef struct _GtkAccelGroupEntry GtkAccelGroupEntry;
struct _GtkAccelGroupEntry { uint8_t _data[32]; };
typedef void* GtkAccelGroupFindFunc;
extern void _GtkAccelGroupFindFunc_c_wrapper();
extern void _GtkAccelGroupFindFunc_c_wrapper_once();
typedef struct _GtkAccelGroupPrivate GtkAccelGroupPrivate;
struct _GtkAccelGroupPrivate {};
typedef struct _GtkAccelKey GtkAccelKey;
struct _GtkAccelKey { uint8_t _data[12]; };
typedef struct _GtkAccelLabel GtkAccelLabel;
typedef struct _GtkAccelLabelClass GtkAccelLabelClass;
struct _GtkAccelLabelClass { uint8_t _data[1032]; };
typedef struct _GtkAccelLabelPrivate GtkAccelLabelPrivate;
struct _GtkAccelLabelPrivate {};
typedef struct _GtkAccelMap GtkAccelMap;
typedef struct _GtkAccelMapClass GtkAccelMapClass;
struct _GtkAccelMapClass {};
typedef void* GtkAccelMapForeach;
extern void _GtkAccelMapForeach_c_wrapper();
extern void _GtkAccelMapForeach_c_wrapper_once();
typedef struct _GtkAccessible GtkAccessible;
typedef struct _GtkAccessibleClass GtkAccessibleClass;
struct _GtkAccessibleClass { uint8_t _data[392]; };
typedef struct _GtkAccessiblePrivate GtkAccessiblePrivate;
struct _GtkAccessiblePrivate {};
typedef struct _GtkAction GtkAction;
typedef struct _GtkActionClass GtkActionClass;
struct _GtkActionClass { uint8_t _data[232]; };
typedef struct _GtkActionEntry GtkActionEntry;
struct _GtkActionEntry { uint8_t _data[48]; };
typedef struct _GtkActionGroup GtkActionGroup;
typedef struct _GtkActionGroupClass GtkActionGroupClass;
struct _GtkActionGroupClass { uint8_t _data[176]; };
typedef struct _GtkActionGroupPrivate GtkActionGroupPrivate;
struct _GtkActionGroupPrivate {};
typedef struct _GtkActionPrivate GtkActionPrivate;
struct _GtkActionPrivate {};
typedef struct _GtkActivatable GtkActivatable;
typedef struct _GtkActivatableIface GtkActivatableIface;
struct _GtkActivatableIface { uint8_t _data[32]; };
typedef struct _GtkAdjustment GtkAdjustment;
typedef struct _GtkAdjustmentClass GtkAdjustmentClass;
struct _GtkAdjustmentClass { uint8_t _data[184]; };
typedef struct _GtkAdjustmentPrivate GtkAdjustmentPrivate;
struct _GtkAdjustmentPrivate {};
typedef uint32_t GtkAlign;
typedef struct _GtkAlignment GtkAlignment;
typedef struct _GtkAlignmentClass GtkAlignmentClass;
struct _GtkAlignmentClass { uint8_t _data[1040]; };
typedef struct _GtkAlignmentPrivate GtkAlignmentPrivate;
struct _GtkAlignmentPrivate {};
typedef struct _GtkAppChooser GtkAppChooser;
typedef struct _GtkAppChooserButton GtkAppChooserButton;
typedef struct _GtkAppChooserButtonClass GtkAppChooserButtonClass;
struct _GtkAppChooserButtonClass { uint8_t _data[1184]; };
typedef struct _GtkAppChooserButtonPrivate GtkAppChooserButtonPrivate;
struct _GtkAppChooserButtonPrivate {};
typedef struct _GtkAppChooserDialog GtkAppChooserDialog;
typedef struct _GtkAppChooserDialogClass GtkAppChooserDialogClass;
struct _GtkAppChooserDialogClass { uint8_t _data[1248]; };
typedef struct _GtkAppChooserDialogPrivate GtkAppChooserDialogPrivate;
struct _GtkAppChooserDialogPrivate {};
typedef struct _GtkAppChooserWidget GtkAppChooserWidget;
typedef struct _GtkAppChooserWidgetClass GtkAppChooserWidgetClass;
struct _GtkAppChooserWidgetClass { uint8_t _data[1160]; };
typedef struct _GtkAppChooserWidgetPrivate GtkAppChooserWidgetPrivate;
struct _GtkAppChooserWidgetPrivate {};
typedef struct _GtkApplication GtkApplication;
typedef struct _GtkApplicationClass GtkApplicationClass;
struct _GtkApplicationClass { uint8_t _data[440]; };
typedef struct _GtkApplicationPrivate GtkApplicationPrivate;
struct _GtkApplicationPrivate {};
typedef struct _GtkArrow GtkArrow;
typedef struct _GtkArrowClass GtkArrowClass;
struct _GtkArrowClass { uint8_t _data[888]; };
typedef uint32_t GtkArrowPlacement;
typedef struct _GtkArrowPrivate GtkArrowPrivate;
struct _GtkArrowPrivate {};
typedef uint32_t GtkArrowType;
typedef struct _GtkAspectFrame GtkAspectFrame;
typedef struct _GtkAspectFrameClass GtkAspectFrameClass;
struct _GtkAspectFrameClass { uint8_t _data[1080]; };
typedef struct _GtkAspectFramePrivate GtkAspectFramePrivate;
struct _GtkAspectFramePrivate {};
typedef struct _GtkAssistant GtkAssistant;
typedef struct _GtkAssistantClass GtkAssistantClass;
struct _GtkAssistantClass { uint8_t _data[1144]; };
typedef void* GtkAssistantPageFunc;
extern void _GtkAssistantPageFunc_c_wrapper();
extern void _GtkAssistantPageFunc_c_wrapper_once();
typedef uint32_t GtkAssistantPageType;
typedef struct _GtkAssistantPrivate GtkAssistantPrivate;
struct _GtkAssistantPrivate {};
typedef uint32_t GtkAttachOptions;
typedef struct _GtkBin GtkBin;
typedef struct _GtkBinClass GtkBinClass;
struct _GtkBinClass { uint8_t _data[1008]; };
typedef struct _GtkBinPrivate GtkBinPrivate;
struct _GtkBinPrivate {};
typedef struct _GtkBindingArg GtkBindingArg;
struct _GtkBindingArg { uint8_t _data[8]; };
typedef struct _GtkBindingEntry GtkBindingEntry;
struct _GtkBindingEntry { uint8_t _data[56]; };
typedef struct _GtkBindingSet GtkBindingSet;
struct _GtkBindingSet { uint8_t _data[64]; };
typedef struct _GtkBindingSignal GtkBindingSignal;
struct _GtkBindingSignal { uint8_t _data[32]; };
typedef struct _GtkBorder GtkBorder;
struct _GtkBorder { uint8_t _data[8]; };
typedef uint32_t GtkBorderStyle;
typedef struct _GtkBox GtkBox;
typedef struct _GtkBoxClass GtkBoxClass;
struct _GtkBoxClass { uint8_t _data[1008]; };
typedef struct _GtkBoxPrivate GtkBoxPrivate;
struct _GtkBoxPrivate {};
typedef struct _GtkBuildable GtkBuildable;
typedef struct _GtkBuildableIface GtkBuildableIface;
struct _GtkBuildableIface { uint8_t _data[96]; };
typedef struct _GtkBuilder GtkBuilder;
typedef struct _GtkBuilderClass GtkBuilderClass;
struct _GtkBuilderClass { uint8_t _data[208]; };
typedef void* GtkBuilderConnectFunc;
extern void _GtkBuilderConnectFunc_c_wrapper();
extern void _GtkBuilderConnectFunc_c_wrapper_once();
typedef uint32_t GtkBuilderError;
typedef struct _GtkBuilderPrivate GtkBuilderPrivate;
struct _GtkBuilderPrivate {};
typedef struct _GtkButton GtkButton;
typedef struct _GtkButtonBox GtkButtonBox;
typedef struct _GtkButtonBoxClass GtkButtonBoxClass;
struct _GtkButtonBoxClass { uint8_t _data[1040]; };
typedef struct _GtkButtonBoxPrivate GtkButtonBoxPrivate;
struct _GtkButtonBoxPrivate {};
typedef uint32_t GtkButtonBoxStyle;
typedef struct _GtkButtonClass GtkButtonClass;
struct _GtkButtonClass { uint8_t _data[1088]; };
typedef struct _GtkButtonPrivate GtkButtonPrivate;
struct _GtkButtonPrivate {};
typedef uint32_t GtkButtonsType;
typedef struct _GtkCalendar GtkCalendar;
typedef struct _GtkCalendarClass GtkCalendarClass;
struct _GtkCalendarClass { uint8_t _data[912]; };
typedef void* GtkCalendarDetailFunc;
extern void _GtkCalendarDetailFunc_c_wrapper();
extern void _GtkCalendarDetailFunc_c_wrapper_once();
typedef uint32_t GtkCalendarDisplayOptions;
typedef struct _GtkCalendarPrivate GtkCalendarPrivate;
struct _GtkCalendarPrivate {};
typedef void* GtkCallback;
extern void _GtkCallback_c_wrapper();
extern void _GtkCallback_c_wrapper_once();
typedef void* GtkCellAllocCallback;
extern void _GtkCellAllocCallback_c_wrapper();
extern void _GtkCellAllocCallback_c_wrapper_once();
typedef struct _GtkCellArea GtkCellArea;
typedef struct _GtkCellAreaBox GtkCellAreaBox;
typedef struct _GtkCellAreaBoxClass GtkCellAreaBoxClass;
struct _GtkCellAreaBoxClass { uint8_t _data[384]; };
typedef struct _GtkCellAreaBoxPrivate GtkCellAreaBoxPrivate;
struct _GtkCellAreaBoxPrivate {};
typedef struct _GtkCellAreaClass GtkCellAreaClass;
struct _GtkCellAreaClass { uint8_t _data[352]; };
typedef struct _GtkCellAreaContext GtkCellAreaContext;
typedef struct _GtkCellAreaContextClass GtkCellAreaContextClass;
struct _GtkCellAreaContextClass { uint8_t _data[216]; };
typedef struct _GtkCellAreaContextPrivate GtkCellAreaContextPrivate;
struct _GtkCellAreaContextPrivate {};
typedef struct _GtkCellAreaPrivate GtkCellAreaPrivate;
struct _GtkCellAreaPrivate {};
typedef void* GtkCellCallback;
extern void _GtkCellCallback_c_wrapper();
extern void _GtkCellCallback_c_wrapper_once();
typedef struct _GtkCellEditable GtkCellEditable;
typedef struct _GtkCellEditableIface GtkCellEditableIface;
struct _GtkCellEditableIface { uint8_t _data[40]; };
typedef struct _GtkCellLayout GtkCellLayout;
typedef void* GtkCellLayoutDataFunc;
extern void _GtkCellLayoutDataFunc_c_wrapper();
extern void _GtkCellLayoutDataFunc_c_wrapper_once();
typedef struct _GtkCellLayoutIface GtkCellLayoutIface;
struct _GtkCellLayoutIface { uint8_t _data[88]; };
typedef struct _GtkCellRenderer GtkCellRenderer;
typedef struct _GtkCellRendererAccel GtkCellRendererAccel;
typedef struct _GtkCellRendererAccelClass GtkCellRendererAccelClass;
struct _GtkCellRendererAccelClass { uint8_t _data[360]; };
typedef uint32_t GtkCellRendererAccelMode;
typedef struct _GtkCellRendererAccelPrivate GtkCellRendererAccelPrivate;
struct _GtkCellRendererAccelPrivate {};
typedef struct _GtkCellRendererClass GtkCellRendererClass;
struct _GtkCellRendererClass { uint8_t _data[264]; };
typedef struct _GtkCellRendererCombo GtkCellRendererCombo;
typedef struct _GtkCellRendererComboClass GtkCellRendererComboClass;
struct _GtkCellRendererComboClass { uint8_t _data[336]; };
typedef struct _GtkCellRendererComboPrivate GtkCellRendererComboPrivate;
struct _GtkCellRendererComboPrivate {};
typedef uint32_t GtkCellRendererMode;
typedef struct _GtkCellRendererPixbuf GtkCellRendererPixbuf;
typedef struct _GtkCellRendererPixbufClass GtkCellRendererPixbufClass;
struct _GtkCellRendererPixbufClass { uint8_t _data[296]; };
typedef struct _GtkCellRendererPixbufPrivate GtkCellRendererPixbufPrivate;
struct _GtkCellRendererPixbufPrivate {};
typedef struct _GtkCellRendererPrivate GtkCellRendererPrivate;
struct _GtkCellRendererPrivate {};
typedef struct _GtkCellRendererProgress GtkCellRendererProgress;
typedef struct _GtkCellRendererProgressClass GtkCellRendererProgressClass;
struct _GtkCellRendererProgressClass { uint8_t _data[296]; };
typedef struct _GtkCellRendererProgressPrivate GtkCellRendererProgressPrivate;
struct _GtkCellRendererProgressPrivate {};
typedef struct _GtkCellRendererSpin GtkCellRendererSpin;
typedef struct _GtkCellRendererSpinClass GtkCellRendererSpinClass;
struct _GtkCellRendererSpinClass { uint8_t _data[336]; };
typedef struct _GtkCellRendererSpinPrivate GtkCellRendererSpinPrivate;
struct _GtkCellRendererSpinPrivate {};
typedef struct _GtkCellRendererSpinner GtkCellRendererSpinner;
typedef struct _GtkCellRendererSpinnerClass GtkCellRendererSpinnerClass;
struct _GtkCellRendererSpinnerClass { uint8_t _data[296]; };
typedef struct _GtkCellRendererSpinnerPrivate GtkCellRendererSpinnerPrivate;
struct _GtkCellRendererSpinnerPrivate {};
typedef uint32_t GtkCellRendererState;
typedef struct _GtkCellRendererText GtkCellRendererText;
typedef struct _GtkCellRendererTextClass GtkCellRendererTextClass;
struct _GtkCellRendererTextClass { uint8_t _data[304]; };
typedef struct _GtkCellRendererTextPrivate GtkCellRendererTextPrivate;
struct _GtkCellRendererTextPrivate {};
typedef struct _GtkCellRendererToggle GtkCellRendererToggle;
typedef struct _GtkCellRendererToggleClass GtkCellRendererToggleClass;
struct _GtkCellRendererToggleClass { uint8_t _data[304]; };
typedef struct _GtkCellRendererTogglePrivate GtkCellRendererTogglePrivate;
struct _GtkCellRendererTogglePrivate {};
typedef struct _GtkCellView GtkCellView;
typedef struct _GtkCellViewClass GtkCellViewClass;
struct _GtkCellViewClass { uint8_t _data[856]; };
typedef struct _GtkCellViewPrivate GtkCellViewPrivate;
struct _GtkCellViewPrivate {};
typedef struct _GtkCheckButton GtkCheckButton;
typedef struct _GtkCheckButtonClass GtkCheckButtonClass;
struct _GtkCheckButtonClass { uint8_t _data[1168]; };
typedef struct _GtkCheckMenuItem GtkCheckMenuItem;
typedef struct _GtkCheckMenuItemClass GtkCheckMenuItemClass;
struct _GtkCheckMenuItemClass { uint8_t _data[1160]; };
typedef struct _GtkCheckMenuItemPrivate GtkCheckMenuItemPrivate;
struct _GtkCheckMenuItemPrivate {};
typedef struct _GtkClipboard GtkClipboard;
typedef void* GtkClipboardClearFunc;
extern void _GtkClipboardClearFunc_c_wrapper();
extern void _GtkClipboardClearFunc_c_wrapper_once();
typedef void* GtkClipboardGetFunc;
extern void _GtkClipboardGetFunc_c_wrapper();
extern void _GtkClipboardGetFunc_c_wrapper_once();
typedef void* GtkClipboardImageReceivedFunc;
extern void _GtkClipboardImageReceivedFunc_c_wrapper();
extern void _GtkClipboardImageReceivedFunc_c_wrapper_once();
typedef void* GtkClipboardReceivedFunc;
extern void _GtkClipboardReceivedFunc_c_wrapper();
extern void _GtkClipboardReceivedFunc_c_wrapper_once();
typedef void* GtkClipboardRichTextReceivedFunc;
extern void _GtkClipboardRichTextReceivedFunc_c_wrapper();
extern void _GtkClipboardRichTextReceivedFunc_c_wrapper_once();
typedef void* GtkClipboardTargetsReceivedFunc;
extern void _GtkClipboardTargetsReceivedFunc_c_wrapper();
extern void _GtkClipboardTargetsReceivedFunc_c_wrapper_once();
typedef void* GtkClipboardTextReceivedFunc;
extern void _GtkClipboardTextReceivedFunc_c_wrapper();
extern void _GtkClipboardTextReceivedFunc_c_wrapper_once();
typedef void* GtkClipboardURIReceivedFunc;
extern void _GtkClipboardURIReceivedFunc_c_wrapper();
extern void _GtkClipboardURIReceivedFunc_c_wrapper_once();
typedef struct _GtkColorButton GtkColorButton;
typedef struct _GtkColorButtonClass GtkColorButtonClass;
struct _GtkColorButtonClass { uint8_t _data[1128]; };
typedef struct _GtkColorButtonPrivate GtkColorButtonPrivate;
struct _GtkColorButtonPrivate {};
typedef struct _GtkColorSelection GtkColorSelection;
typedef void* GtkColorSelectionChangePaletteFunc;
extern void _GtkColorSelectionChangePaletteFunc_c_wrapper();
extern void _GtkColorSelectionChangePaletteFunc_c_wrapper_once();
typedef void* GtkColorSelectionChangePaletteWithScreenFunc;
extern void _GtkColorSelectionChangePaletteWithScreenFunc_c_wrapper();
extern void _GtkColorSelectionChangePaletteWithScreenFunc_c_wrapper_once();
typedef struct _GtkColorSelectionClass GtkColorSelectionClass;
struct _GtkColorSelectionClass { uint8_t _data[1048]; };
typedef struct _GtkColorSelectionDialog GtkColorSelectionDialog;
typedef struct _GtkColorSelectionDialogClass GtkColorSelectionDialogClass;
struct _GtkColorSelectionDialogClass { uint8_t _data[1152]; };
typedef struct _GtkColorSelectionDialogPrivate GtkColorSelectionDialogPrivate;
struct _GtkColorSelectionDialogPrivate {};
typedef struct _GtkColorSelectionPrivate GtkColorSelectionPrivate;
struct _GtkColorSelectionPrivate {};
typedef struct _GtkComboBox GtkComboBox;
typedef struct _GtkComboBoxClass GtkComboBoxClass;
struct _GtkComboBoxClass { uint8_t _data[1048]; };
typedef struct _GtkComboBoxPrivate GtkComboBoxPrivate;
struct _GtkComboBoxPrivate {};
typedef struct _GtkComboBoxText GtkComboBoxText;
typedef struct _GtkComboBoxTextClass GtkComboBoxTextClass;
struct _GtkComboBoxTextClass { uint8_t _data[1080]; };
typedef struct _GtkComboBoxTextPrivate GtkComboBoxTextPrivate;
struct _GtkComboBoxTextPrivate {};
typedef struct _GtkContainer GtkContainer;
typedef struct _GtkContainerClass GtkContainerClass;
struct _GtkContainerClass { uint8_t _data[976]; };
typedef struct _GtkContainerPrivate GtkContainerPrivate;
struct _GtkContainerPrivate {};
typedef uint32_t GtkCornerType;
typedef struct _GtkCssProvider GtkCssProvider;
typedef struct _GtkCssProviderClass GtkCssProviderClass;
struct _GtkCssProviderClass { uint8_t _data[168]; };
typedef uint32_t GtkCssProviderError;
typedef struct _GtkCssProviderPrivate GtkCssProviderPrivate;
struct _GtkCssProviderPrivate {};
typedef struct _GtkCssSection GtkCssSection;
struct _GtkCssSection {};
typedef uint32_t GtkCssSectionType;
typedef uint32_t GtkDebugFlag;
typedef uint32_t GtkDeleteType;
typedef uint32_t GtkDestDefaults;
typedef struct _GtkDialog GtkDialog;
typedef struct _GtkDialogClass GtkDialogClass;
struct _GtkDialogClass { uint8_t _data[1120]; };
typedef uint32_t GtkDialogFlags;
typedef struct _GtkDialogPrivate GtkDialogPrivate;
struct _GtkDialogPrivate {};
typedef uint32_t GtkDirectionType;
typedef uint32_t GtkDragResult;
typedef struct _GtkDrawingArea GtkDrawingArea;
typedef struct _GtkDrawingAreaClass GtkDrawingAreaClass;
struct _GtkDrawingAreaClass { uint8_t _data[856]; };
typedef struct _GtkEditable GtkEditable;
typedef struct _GtkEditableInterface GtkEditableInterface;
struct _GtkEditableInterface { uint8_t _data[96]; };
typedef struct _GtkEntry GtkEntry;
typedef struct _GtkEntryBuffer GtkEntryBuffer;
typedef struct _GtkEntryBufferClass GtkEntryBufferClass;
struct _GtkEntryBufferClass { uint8_t _data[248]; };
typedef struct _GtkEntryBufferPrivate GtkEntryBufferPrivate;
struct _GtkEntryBufferPrivate {};
typedef struct _GtkEntryClass GtkEntryClass;
struct _GtkEntryClass { uint8_t _data[976]; };
typedef struct _GtkEntryCompletion GtkEntryCompletion;
typedef struct _GtkEntryCompletionClass GtkEntryCompletionClass;
struct _GtkEntryCompletionClass { uint8_t _data[200]; };
typedef void* GtkEntryCompletionMatchFunc;
extern void _GtkEntryCompletionMatchFunc_c_wrapper();
extern void _GtkEntryCompletionMatchFunc_c_wrapper_once();
typedef struct _GtkEntryCompletionPrivate GtkEntryCompletionPrivate;
struct _GtkEntryCompletionPrivate {};
typedef uint32_t GtkEntryIconPosition;
typedef struct _GtkEntryPrivate GtkEntryPrivate;
struct _GtkEntryPrivate {};
typedef struct _GtkEventBox GtkEventBox;
typedef struct _GtkEventBoxClass GtkEventBoxClass;
struct _GtkEventBoxClass { uint8_t _data[1040]; };
typedef struct _GtkEventBoxPrivate GtkEventBoxPrivate;
struct _GtkEventBoxPrivate {};
typedef struct _GtkExpander GtkExpander;
typedef struct _GtkExpanderClass GtkExpanderClass;
struct _GtkExpanderClass { uint8_t _data[1048]; };
typedef struct _GtkExpanderPrivate GtkExpanderPrivate;
struct _GtkExpanderPrivate {};
typedef uint32_t GtkExpanderStyle;
typedef struct _GtkFileChooser GtkFileChooser;
typedef uint32_t GtkFileChooserAction;
typedef struct _GtkFileChooserButton GtkFileChooserButton;
typedef struct _GtkFileChooserButtonClass GtkFileChooserButtonClass;
struct _GtkFileChooserButtonClass { uint8_t _data[1048]; };
typedef struct _GtkFileChooserButtonPrivate GtkFileChooserButtonPrivate;
struct _GtkFileChooserButtonPrivate {};
typedef uint32_t GtkFileChooserConfirmation;
typedef struct _GtkFileChooserDialog GtkFileChooserDialog;
typedef struct _GtkFileChooserDialogClass GtkFileChooserDialogClass;
struct _GtkFileChooserDialogClass { uint8_t _data[1152]; };
typedef struct _GtkFileChooserDialogPrivate GtkFileChooserDialogPrivate;
struct _GtkFileChooserDialogPrivate {};
typedef uint32_t GtkFileChooserError;
typedef struct _GtkFileChooserWidget GtkFileChooserWidget;
typedef struct _GtkFileChooserWidgetClass GtkFileChooserWidgetClass;
struct _GtkFileChooserWidgetClass { uint8_t _data[1040]; };
typedef struct _GtkFileChooserWidgetPrivate GtkFileChooserWidgetPrivate;
struct _GtkFileChooserWidgetPrivate {};
typedef struct _GtkFileFilter GtkFileFilter;
typedef uint32_t GtkFileFilterFlags;
typedef void* GtkFileFilterFunc;
extern void _GtkFileFilterFunc_c_wrapper();
extern void _GtkFileFilterFunc_c_wrapper_once();
typedef struct _GtkFileFilterInfo GtkFileFilterInfo;
struct _GtkFileFilterInfo { uint8_t _data[40]; };
typedef struct _GtkFixed GtkFixed;
typedef struct _GtkFixedChild GtkFixedChild;
struct _GtkFixedChild { uint8_t _data[16]; };
typedef struct _GtkFixedClass GtkFixedClass;
struct _GtkFixedClass { uint8_t _data[1008]; };
typedef struct _GtkFixedPrivate GtkFixedPrivate;
struct _GtkFixedPrivate {};
typedef struct _GtkFontButton GtkFontButton;
typedef struct _GtkFontButtonClass GtkFontButtonClass;
struct _GtkFontButtonClass { uint8_t _data[1128]; };
typedef struct _GtkFontButtonPrivate GtkFontButtonPrivate;
struct _GtkFontButtonPrivate {};
typedef struct _GtkFontChooser GtkFontChooser;
typedef struct _GtkFontChooserDialog GtkFontChooserDialog;
typedef struct _GtkFontChooserDialogClass GtkFontChooserDialogClass;
struct _GtkFontChooserDialogClass { uint8_t _data[1152]; };
typedef struct _GtkFontChooserDialogPrivate GtkFontChooserDialogPrivate;
struct _GtkFontChooserDialogPrivate {};
typedef struct _GtkFontChooserIface GtkFontChooserIface;
struct _GtkFontChooserIface { uint8_t _data[152]; };
typedef struct _GtkFontChooserWidget GtkFontChooserWidget;
typedef struct _GtkFontChooserWidgetClass GtkFontChooserWidgetClass;
struct _GtkFontChooserWidgetClass { uint8_t _data[1072]; };
typedef struct _GtkFontChooserWidgetPrivate GtkFontChooserWidgetPrivate;
struct _GtkFontChooserWidgetPrivate {};
typedef void* GtkFontFilterFunc;
extern void _GtkFontFilterFunc_c_wrapper();
extern void _GtkFontFilterFunc_c_wrapper_once();
typedef struct _GtkFontSelection GtkFontSelection;
typedef struct _GtkFontSelectionClass GtkFontSelectionClass;
struct _GtkFontSelectionClass { uint8_t _data[1040]; };
typedef struct _GtkFontSelectionDialog GtkFontSelectionDialog;
typedef struct _GtkFontSelectionDialogClass GtkFontSelectionDialogClass;
struct _GtkFontSelectionDialogClass { uint8_t _data[1152]; };
typedef struct _GtkFontSelectionDialogPrivate GtkFontSelectionDialogPrivate;
struct _GtkFontSelectionDialogPrivate {};
typedef struct _GtkFontSelectionPrivate GtkFontSelectionPrivate;
struct _GtkFontSelectionPrivate {};
typedef struct _GtkFrame GtkFrame;
typedef struct _GtkFrameClass GtkFrameClass;
struct _GtkFrameClass { uint8_t _data[1048]; };
typedef struct _GtkFramePrivate GtkFramePrivate;
struct _GtkFramePrivate {};
typedef struct _GtkGradient GtkGradient;
struct _GtkGradient {};
typedef struct _GtkGrid GtkGrid;
typedef struct _GtkGridClass GtkGridClass;
struct _GtkGridClass { uint8_t _data[1040]; };
typedef struct _GtkGridPrivate GtkGridPrivate;
struct _GtkGridPrivate {};
typedef struct _GtkHBox GtkHBox;
typedef struct _GtkHBoxClass GtkHBoxClass;
struct _GtkHBoxClass { uint8_t _data[1008]; };
typedef struct _GtkHButtonBox GtkHButtonBox;
typedef struct _GtkHButtonBoxClass GtkHButtonBoxClass;
struct _GtkHButtonBoxClass { uint8_t _data[1040]; };
typedef struct _GtkHPaned GtkHPaned;
typedef struct _GtkHPanedClass GtkHPanedClass;
struct _GtkHPanedClass { uint8_t _data[1056]; };
typedef struct _GtkHSV GtkHSV;
typedef struct _GtkHSVClass GtkHSVClass;
struct _GtkHSVClass { uint8_t _data[872]; };
typedef struct _GtkHSVPrivate GtkHSVPrivate;
struct _GtkHSVPrivate {};
typedef struct _GtkHScale GtkHScale;
typedef struct _GtkHScaleClass GtkHScaleClass;
struct _GtkHScaleClass { uint8_t _data[968]; };
typedef struct _GtkHScrollbar GtkHScrollbar;
typedef struct _GtkHScrollbarClass GtkHScrollbarClass;
struct _GtkHScrollbarClass { uint8_t _data[944]; };
typedef struct _GtkHSeparator GtkHSeparator;
typedef struct _GtkHSeparatorClass GtkHSeparatorClass;
struct _GtkHSeparatorClass { uint8_t _data[856]; };
typedef struct _GtkHandleBox GtkHandleBox;
typedef struct _GtkHandleBoxClass GtkHandleBoxClass;
struct _GtkHandleBoxClass { uint8_t _data[1056]; };
typedef struct _GtkHandleBoxPrivate GtkHandleBoxPrivate;
struct _GtkHandleBoxPrivate {};
typedef struct _GtkIMContext GtkIMContext;
typedef struct _GtkIMContextClass GtkIMContextClass;
struct _GtkIMContextClass { uint8_t _data[312]; };
typedef struct _GtkIMContextInfo GtkIMContextInfo;
struct _GtkIMContextInfo { uint8_t _data[40]; };
typedef struct _GtkIMContextSimple GtkIMContextSimple;
typedef struct _GtkIMContextSimpleClass GtkIMContextSimpleClass;
struct _GtkIMContextSimpleClass { uint8_t _data[312]; };
typedef struct _GtkIMContextSimplePrivate GtkIMContextSimplePrivate;
struct _GtkIMContextSimplePrivate {};
typedef struct _GtkIMMulticontext GtkIMMulticontext;
typedef struct _GtkIMMulticontextClass GtkIMMulticontextClass;
struct _GtkIMMulticontextClass { uint8_t _data[344]; };
typedef struct _GtkIMMulticontextPrivate GtkIMMulticontextPrivate;
struct _GtkIMMulticontextPrivate {};
typedef uint32_t GtkIMPreeditStyle;
typedef uint32_t GtkIMStatusStyle;
typedef struct _GtkIconFactory GtkIconFactory;
typedef struct _GtkIconFactoryClass GtkIconFactoryClass;
struct _GtkIconFactoryClass { uint8_t _data[168]; };
typedef struct _GtkIconFactoryPrivate GtkIconFactoryPrivate;
struct _GtkIconFactoryPrivate {};
typedef struct _GtkIconInfo GtkIconInfo;
struct _GtkIconInfo {};
typedef uint32_t GtkIconLookupFlags;
typedef struct _GtkIconSet GtkIconSet;
struct _GtkIconSet {};
typedef uint32_t GtkIconSize;
typedef struct _GtkIconSource GtkIconSource;
struct _GtkIconSource {};
typedef struct _GtkIconTheme GtkIconTheme;
typedef struct _GtkIconThemeClass GtkIconThemeClass;
struct _GtkIconThemeClass { uint8_t _data[176]; };
typedef uint32_t GtkIconThemeError;
typedef struct _GtkIconThemePrivate GtkIconThemePrivate;
struct _GtkIconThemePrivate {};
typedef struct _GtkIconView GtkIconView;
typedef struct _GtkIconViewClass GtkIconViewClass;
struct _GtkIconViewClass { uint8_t _data[1072]; };
typedef uint32_t GtkIconViewDropPosition;
typedef void* GtkIconViewForeachFunc;
extern void _GtkIconViewForeachFunc_c_wrapper();
extern void _GtkIconViewForeachFunc_c_wrapper_once();
typedef struct _GtkIconViewPrivate GtkIconViewPrivate;
struct _GtkIconViewPrivate {};
typedef struct _GtkImage GtkImage;
typedef struct _GtkImageClass GtkImageClass;
struct _GtkImageClass { uint8_t _data[888]; };
typedef struct _GtkImageMenuItem GtkImageMenuItem;
typedef struct _GtkImageMenuItemClass GtkImageMenuItemClass;
struct _GtkImageMenuItemClass { uint8_t _data[1144]; };
typedef struct _GtkImageMenuItemPrivate GtkImageMenuItemPrivate;
struct _GtkImageMenuItemPrivate {};
typedef struct _GtkImagePrivate GtkImagePrivate;
struct _GtkImagePrivate {};
typedef uint32_t GtkImageType;
typedef struct _GtkInfoBar GtkInfoBar;
typedef struct _GtkInfoBarClass GtkInfoBarClass;
struct _GtkInfoBarClass { uint8_t _data[1056]; };
typedef struct _GtkInfoBarPrivate GtkInfoBarPrivate;
struct _GtkInfoBarPrivate {};
typedef struct _GtkInvisible GtkInvisible;
typedef struct _GtkInvisibleClass GtkInvisibleClass;
struct _GtkInvisibleClass { uint8_t _data[856]; };
typedef struct _GtkInvisiblePrivate GtkInvisiblePrivate;
struct _GtkInvisiblePrivate {};
typedef uint32_t GtkJunctionSides;
typedef uint32_t GtkJustification;
typedef void* GtkKeySnoopFunc;
extern void _GtkKeySnoopFunc_c_wrapper();
extern void _GtkKeySnoopFunc_c_wrapper_once();
typedef struct _GtkLabel GtkLabel;
typedef struct _GtkLabelClass GtkLabelClass;
struct _GtkLabelClass { uint8_t _data[952]; };
typedef struct _GtkLabelPrivate GtkLabelPrivate;
struct _GtkLabelPrivate {};
typedef struct _GtkLabelSelectionInfo GtkLabelSelectionInfo;
struct _GtkLabelSelectionInfo {};
typedef struct _GtkLayout GtkLayout;
typedef struct _GtkLayoutClass GtkLayoutClass;
struct _GtkLayoutClass { uint8_t _data[1008]; };
typedef struct _GtkLayoutPrivate GtkLayoutPrivate;
struct _GtkLayoutPrivate {};
typedef uint32_t GtkLicense;
typedef struct _GtkLinkButton GtkLinkButton;
typedef struct _GtkLinkButtonClass GtkLinkButtonClass;
struct _GtkLinkButtonClass { uint8_t _data[1128]; };
typedef struct _GtkLinkButtonPrivate GtkLinkButtonPrivate;
struct _GtkLinkButtonPrivate {};
typedef struct _GtkListStore GtkListStore;
typedef struct _GtkListStoreClass GtkListStoreClass;
struct _GtkListStoreClass { uint8_t _data[168]; };
typedef struct _GtkListStorePrivate GtkListStorePrivate;
struct _GtkListStorePrivate {};
typedef struct _GtkLockButton GtkLockButton;
typedef struct _GtkLockButtonClass GtkLockButtonClass;
struct _GtkLockButtonClass { uint8_t _data[1152]; };
typedef struct _GtkLockButtonPrivate GtkLockButtonPrivate;
struct _GtkLockButtonPrivate {};
typedef struct _GtkMenu GtkMenu;
typedef struct _GtkMenuBar GtkMenuBar;
typedef struct _GtkMenuBarClass GtkMenuBarClass;
struct _GtkMenuBarClass { uint8_t _data[1120]; };
typedef struct _GtkMenuBarPrivate GtkMenuBarPrivate;
struct _GtkMenuBarPrivate {};
typedef struct _GtkMenuClass GtkMenuClass;
struct _GtkMenuClass { uint8_t _data[1120]; };
typedef void* GtkMenuDetachFunc;
extern void _GtkMenuDetachFunc_c_wrapper();
extern void _GtkMenuDetachFunc_c_wrapper_once();
typedef uint32_t GtkMenuDirectionType;
typedef struct _GtkMenuItem GtkMenuItem;
typedef struct _GtkMenuItemClass GtkMenuItemClass;
struct _GtkMenuItemClass { uint8_t _data[1112]; };
typedef struct _GtkMenuItemPrivate GtkMenuItemPrivate;
struct _GtkMenuItemPrivate {};
typedef void* GtkMenuPositionFunc;
extern void _GtkMenuPositionFunc_c_wrapper();
extern void _GtkMenuPositionFunc_c_wrapper_once();
typedef struct _GtkMenuPrivate GtkMenuPrivate;
struct _GtkMenuPrivate {};
typedef struct _GtkMenuShell GtkMenuShell;
typedef struct _GtkMenuShellClass GtkMenuShellClass;
struct _GtkMenuShellClass { uint8_t _data[1088]; };
typedef struct _GtkMenuShellPrivate GtkMenuShellPrivate;
struct _GtkMenuShellPrivate {};
typedef struct _GtkMenuToolButton GtkMenuToolButton;
typedef struct _GtkMenuToolButtonClass GtkMenuToolButtonClass;
struct _GtkMenuToolButtonClass { uint8_t _data[1144]; };
typedef struct _GtkMenuToolButtonPrivate GtkMenuToolButtonPrivate;
struct _GtkMenuToolButtonPrivate {};
typedef struct _GtkMessageDialog GtkMessageDialog;
typedef struct _GtkMessageDialogClass GtkMessageDialogClass;
struct _GtkMessageDialogClass { uint8_t _data[1152]; };
typedef struct _GtkMessageDialogPrivate GtkMessageDialogPrivate;
struct _GtkMessageDialogPrivate {};
typedef uint32_t GtkMessageType;
typedef struct _GtkMisc GtkMisc;
typedef struct _GtkMiscClass GtkMiscClass;
struct _GtkMiscClass { uint8_t _data[856]; };
typedef struct _GtkMiscPrivate GtkMiscPrivate;
struct _GtkMiscPrivate {};
typedef void* GtkModuleDisplayInitFunc;
extern void _GtkModuleDisplayInitFunc_c_wrapper();
extern void _GtkModuleDisplayInitFunc_c_wrapper_once();
typedef void* GtkModuleInitFunc;
extern void _GtkModuleInitFunc_c_wrapper();
extern void _GtkModuleInitFunc_c_wrapper_once();
typedef struct _GtkMountOperation GtkMountOperation;
typedef struct _GtkMountOperationClass GtkMountOperationClass;
struct _GtkMountOperationClass { uint8_t _data[288]; };
typedef struct _GtkMountOperationPrivate GtkMountOperationPrivate;
struct _GtkMountOperationPrivate {};
typedef uint32_t GtkMovementStep;
typedef struct _GtkNotebook GtkNotebook;
typedef struct _GtkNotebookClass GtkNotebookClass;
struct _GtkNotebookClass { uint8_t _data[1128]; };
typedef struct _GtkNotebookPrivate GtkNotebookPrivate;
struct _GtkNotebookPrivate {};
typedef uint32_t GtkNotebookTab;
typedef uint32_t GtkNumberUpLayout;
typedef struct _GtkNumerableIcon GtkNumerableIcon;
typedef struct _GtkNumerableIconClass GtkNumerableIconClass;
struct _GtkNumerableIconClass { uint8_t _data[264]; };
typedef struct _GtkNumerableIconPrivate GtkNumerableIconPrivate;
struct _GtkNumerableIconPrivate {};
typedef struct _GtkOffscreenWindow GtkOffscreenWindow;
typedef struct _GtkOffscreenWindowClass GtkOffscreenWindowClass;
struct _GtkOffscreenWindowClass { uint8_t _data[1104]; };
typedef struct _GtkOrientable GtkOrientable;
typedef struct _GtkOrientableIface GtkOrientableIface;
struct _GtkOrientableIface { uint8_t _data[16]; };
typedef uint32_t GtkOrientation;
typedef struct _GtkOverlay GtkOverlay;
typedef struct _GtkOverlayClass GtkOverlayClass;
struct _GtkOverlayClass { uint8_t _data[1080]; };
typedef struct _GtkOverlayPrivate GtkOverlayPrivate;
struct _GtkOverlayPrivate {};
typedef uint32_t GtkPackDirection;
typedef uint32_t GtkPackType;
typedef uint32_t GtkPageOrientation;
typedef struct _GtkPageRange GtkPageRange;
struct _GtkPageRange { uint8_t _data[8]; };
typedef uint32_t GtkPageSet;
typedef struct _GtkPageSetup GtkPageSetup;
typedef void* GtkPageSetupDoneFunc;
extern void _GtkPageSetupDoneFunc_c_wrapper();
extern void _GtkPageSetupDoneFunc_c_wrapper_once();
typedef struct _GtkPaned GtkPaned;
typedef struct _GtkPanedClass GtkPanedClass;
struct _GtkPanedClass { uint8_t _data[1056]; };
typedef struct _GtkPanedPrivate GtkPanedPrivate;
struct _GtkPanedPrivate {};
typedef struct _GtkPaperSize GtkPaperSize;
struct _GtkPaperSize {};
typedef uint32_t GtkPathPriorityType;
typedef uint32_t GtkPathType;
typedef struct _GtkPlug GtkPlug;
typedef struct _GtkPlugClass GtkPlugClass;
struct _GtkPlugClass { uint8_t _data[1112]; };
typedef struct _GtkPlugPrivate GtkPlugPrivate;
struct _GtkPlugPrivate {};
typedef uint32_t GtkPolicyType;
typedef uint32_t GtkPositionType;
typedef struct _GtkPrintContext GtkPrintContext;
typedef uint32_t GtkPrintDuplex;
typedef uint32_t GtkPrintError;
typedef struct _GtkPrintOperation GtkPrintOperation;
typedef uint32_t GtkPrintOperationAction;
typedef struct _GtkPrintOperationClass GtkPrintOperationClass;
struct _GtkPrintOperationClass { uint8_t _data[288]; };
typedef struct _GtkPrintOperationPreview GtkPrintOperationPreview;
typedef struct _GtkPrintOperationPreviewIface GtkPrintOperationPreviewIface;
struct _GtkPrintOperationPreviewIface { uint8_t _data[120]; };
typedef struct _GtkPrintOperationPrivate GtkPrintOperationPrivate;
struct _GtkPrintOperationPrivate {};
typedef uint32_t GtkPrintOperationResult;
typedef uint32_t GtkPrintPages;
typedef uint32_t GtkPrintQuality;
typedef struct _GtkPrintSettings GtkPrintSettings;
typedef void* GtkPrintSettingsFunc;
extern void _GtkPrintSettingsFunc_c_wrapper();
extern void _GtkPrintSettingsFunc_c_wrapper_once();
typedef uint32_t GtkPrintStatus;
typedef struct _GtkProgressBar GtkProgressBar;
typedef struct _GtkProgressBarClass GtkProgressBarClass;
struct _GtkProgressBarClass { uint8_t _data[856]; };
typedef struct _GtkProgressBarPrivate GtkProgressBarPrivate;
struct _GtkProgressBarPrivate {};
typedef struct _GtkRadioAction GtkRadioAction;
typedef struct _GtkRadioActionClass GtkRadioActionClass;
struct _GtkRadioActionClass { uint8_t _data[312]; };
typedef struct _GtkRadioActionEntry GtkRadioActionEntry;
struct _GtkRadioActionEntry { uint8_t _data[48]; };
typedef struct _GtkRadioActionPrivate GtkRadioActionPrivate;
struct _GtkRadioActionPrivate {};
typedef struct _GtkRadioButton GtkRadioButton;
typedef struct _GtkRadioButtonClass GtkRadioButtonClass;
struct _GtkRadioButtonClass { uint8_t _data[1208]; };
typedef struct _GtkRadioButtonPrivate GtkRadioButtonPrivate;
struct _GtkRadioButtonPrivate {};
typedef struct _GtkRadioMenuItem GtkRadioMenuItem;
typedef struct _GtkRadioMenuItemClass GtkRadioMenuItemClass;
struct _GtkRadioMenuItemClass { uint8_t _data[1200]; };
typedef struct _GtkRadioMenuItemPrivate GtkRadioMenuItemPrivate;
struct _GtkRadioMenuItemPrivate {};
typedef struct _GtkRadioToolButton GtkRadioToolButton;
typedef struct _GtkRadioToolButtonClass GtkRadioToolButtonClass;
struct _GtkRadioToolButtonClass { uint8_t _data[1176]; };
typedef struct _GtkRange GtkRange;
typedef struct _GtkRangeClass GtkRangeClass;
struct _GtkRangeClass { uint8_t _data[912]; };
typedef struct _GtkRangePrivate GtkRangePrivate;
struct _GtkRangePrivate {};
typedef struct _GtkRcContext GtkRcContext;
struct _GtkRcContext {};
typedef uint32_t GtkRcFlags;
typedef struct _GtkRcProperty GtkRcProperty;
struct _GtkRcProperty { uint8_t _data[40]; };
typedef void* GtkRcPropertyParser;
extern void _GtkRcPropertyParser_c_wrapper();
extern void _GtkRcPropertyParser_c_wrapper_once();
typedef struct _GtkRcStyle GtkRcStyle;
typedef struct _GtkRcStyleClass GtkRcStyleClass;
struct _GtkRcStyleClass { uint8_t _data[200]; };
typedef uint32_t GtkRcTokenType;
typedef struct _GtkRecentAction GtkRecentAction;
typedef struct _GtkRecentActionClass GtkRecentActionClass;
struct _GtkRecentActionClass { uint8_t _data[264]; };
typedef struct _GtkRecentActionPrivate GtkRecentActionPrivate;
struct _GtkRecentActionPrivate {};
typedef struct _GtkRecentChooser GtkRecentChooser;
typedef struct _GtkRecentChooserDialog GtkRecentChooserDialog;
typedef struct _GtkRecentChooserDialogClass GtkRecentChooserDialogClass;
struct _GtkRecentChooserDialogClass { uint8_t _data[1152]; };
typedef struct _GtkRecentChooserDialogPrivate GtkRecentChooserDialogPrivate;
struct _GtkRecentChooserDialogPrivate {};
typedef uint32_t GtkRecentChooserError;
typedef struct _GtkRecentChooserIface GtkRecentChooserIface;
struct _GtkRecentChooserIface { uint8_t _data[128]; };
typedef struct _GtkRecentChooserMenu GtkRecentChooserMenu;
typedef struct _GtkRecentChooserMenuClass GtkRecentChooserMenuClass;
struct _GtkRecentChooserMenuClass { uint8_t _data[1152]; };
typedef struct _GtkRecentChooserMenuPrivate GtkRecentChooserMenuPrivate;
struct _GtkRecentChooserMenuPrivate {};
typedef struct _GtkRecentChooserWidget GtkRecentChooserWidget;
typedef struct _GtkRecentChooserWidgetClass GtkRecentChooserWidgetClass;
struct _GtkRecentChooserWidgetClass { uint8_t _data[1040]; };
typedef struct _GtkRecentChooserWidgetPrivate GtkRecentChooserWidgetPrivate;
struct _GtkRecentChooserWidgetPrivate {};
typedef struct _GtkRecentData GtkRecentData;
struct _GtkRecentData { uint8_t _data[56]; };
typedef struct _GtkRecentFilter GtkRecentFilter;
typedef uint32_t GtkRecentFilterFlags;
typedef void* GtkRecentFilterFunc;
extern void _GtkRecentFilterFunc_c_wrapper();
extern void _GtkRecentFilterFunc_c_wrapper_once();
typedef struct _GtkRecentFilterInfo GtkRecentFilterInfo;
struct _GtkRecentFilterInfo { uint8_t _data[56]; };
typedef struct _GtkRecentInfo GtkRecentInfo;
struct _GtkRecentInfo {};
typedef struct _GtkRecentManager GtkRecentManager;
typedef struct _GtkRecentManagerClass GtkRecentManagerClass;
struct _GtkRecentManagerClass { uint8_t _data[176]; };
typedef uint32_t GtkRecentManagerError;
typedef struct _GtkRecentManagerPrivate GtkRecentManagerPrivate;
struct _GtkRecentManagerPrivate {};
typedef void* GtkRecentSortFunc;
extern void _GtkRecentSortFunc_c_wrapper();
extern void _GtkRecentSortFunc_c_wrapper_once();
typedef uint32_t GtkRecentSortType;
typedef uint32_t GtkRegionFlags;
typedef uint32_t GtkReliefStyle;
typedef struct _GtkRequestedSize GtkRequestedSize;
struct _GtkRequestedSize { uint8_t _data[16]; };
typedef struct _GtkRequisition GtkRequisition;
struct _GtkRequisition { uint8_t _data[8]; };
typedef uint32_t GtkResizeMode;
typedef int32_t GtkResponseType;
typedef struct _GtkScale GtkScale;
typedef struct _GtkScaleButton GtkScaleButton;
typedef struct _GtkScaleButtonClass GtkScaleButtonClass;
struct _GtkScaleButtonClass { uint8_t _data[1128]; };
typedef struct _GtkScaleButtonPrivate GtkScaleButtonPrivate;
struct _GtkScaleButtonPrivate {};
typedef struct _GtkScaleClass GtkScaleClass;
struct _GtkScaleClass { uint8_t _data[968]; };
typedef struct _GtkScalePrivate GtkScalePrivate;
struct _GtkScalePrivate {};
typedef uint32_t GtkScrollStep;
typedef uint32_t GtkScrollType;
typedef struct _GtkScrollable GtkScrollable;
typedef struct _GtkScrollableInterface GtkScrollableInterface;
struct _GtkScrollableInterface { uint8_t _data[16]; };
typedef uint32_t GtkScrollablePolicy;
typedef struct _GtkScrollbar GtkScrollbar;
typedef struct _GtkScrollbarClass GtkScrollbarClass;
struct _GtkScrollbarClass { uint8_t _data[944]; };
typedef struct _GtkScrolledWindow GtkScrolledWindow;
typedef struct _GtkScrolledWindowClass GtkScrolledWindowClass;
struct _GtkScrolledWindowClass { uint8_t _data[1064]; };
typedef struct _GtkScrolledWindowPrivate GtkScrolledWindowPrivate;
struct _GtkScrolledWindowPrivate {};
typedef struct _GtkSelectionData GtkSelectionData;
struct _GtkSelectionData {};
typedef uint32_t GtkSelectionMode;
typedef uint32_t GtkSensitivityType;
typedef struct _GtkSeparator GtkSeparator;
typedef struct _GtkSeparatorClass GtkSeparatorClass;
struct _GtkSeparatorClass { uint8_t _data[856]; };
typedef struct _GtkSeparatorMenuItem GtkSeparatorMenuItem;
typedef struct _GtkSeparatorMenuItemClass GtkSeparatorMenuItemClass;
struct _GtkSeparatorMenuItemClass { uint8_t _data[1144]; };
typedef struct _GtkSeparatorPrivate GtkSeparatorPrivate;
struct _GtkSeparatorPrivate {};
typedef struct _GtkSeparatorToolItem GtkSeparatorToolItem;
typedef struct _GtkSeparatorToolItemClass GtkSeparatorToolItemClass;
struct _GtkSeparatorToolItemClass { uint8_t _data[1088]; };
typedef struct _GtkSeparatorToolItemPrivate GtkSeparatorToolItemPrivate;
struct _GtkSeparatorToolItemPrivate {};
typedef struct _GtkSettings GtkSettings;
typedef struct _GtkSettingsClass GtkSettingsClass;
struct _GtkSettingsClass { uint8_t _data[168]; };
typedef struct _GtkSettingsPrivate GtkSettingsPrivate;
struct _GtkSettingsPrivate {};
typedef struct _GtkSettingsValue GtkSettingsValue;
struct _GtkSettingsValue { uint8_t _data[32]; };
typedef uint32_t GtkShadowType;
typedef struct _GtkSizeGroup GtkSizeGroup;
typedef struct _GtkSizeGroupClass GtkSizeGroupClass;
struct _GtkSizeGroupClass { uint8_t _data[168]; };
typedef uint32_t GtkSizeGroupMode;
typedef struct _GtkSizeGroupPrivate GtkSizeGroupPrivate;
struct _GtkSizeGroupPrivate {};
typedef uint32_t GtkSizeRequestMode;
typedef struct _GtkSocket GtkSocket;
typedef struct _GtkSocketClass GtkSocketClass;
struct _GtkSocketClass { uint8_t _data[1024]; };
typedef struct _GtkSocketPrivate GtkSocketPrivate;
struct _GtkSocketPrivate {};
typedef uint32_t GtkSortType;
typedef struct _GtkSpinButton GtkSpinButton;
typedef struct _GtkSpinButtonClass GtkSpinButtonClass;
struct _GtkSpinButtonClass { uint8_t _data[1048]; };
typedef struct _GtkSpinButtonPrivate GtkSpinButtonPrivate;
struct _GtkSpinButtonPrivate {};
typedef uint32_t GtkSpinButtonUpdatePolicy;
typedef uint32_t GtkSpinType;
typedef struct _GtkSpinner GtkSpinner;
typedef struct _GtkSpinnerClass GtkSpinnerClass;
struct _GtkSpinnerClass { uint8_t _data[856]; };
typedef struct _GtkSpinnerPrivate GtkSpinnerPrivate;
struct _GtkSpinnerPrivate {};
typedef uint32_t GtkStateFlags;
typedef uint32_t GtkStateType;
typedef struct _GtkStatusIcon GtkStatusIcon;
typedef struct _GtkStatusIconClass GtkStatusIconClass;
struct _GtkStatusIconClass { uint8_t _data[224]; };
typedef struct _GtkStatusIconPrivate GtkStatusIconPrivate;
struct _GtkStatusIconPrivate {};
typedef struct _GtkStatusbar GtkStatusbar;
typedef struct _GtkStatusbarClass GtkStatusbarClass;
struct _GtkStatusbarClass { uint8_t _data[1064]; };
typedef struct _GtkStatusbarPrivate GtkStatusbarPrivate;
struct _GtkStatusbarPrivate {};
typedef struct _GtkStockItem GtkStockItem;
struct _GtkStockItem { uint8_t _data[32]; };
typedef struct _GtkStyle GtkStyle;
typedef struct _GtkStyleClass GtkStyleClass;
struct _GtkStyleClass { uint8_t _data[440]; };
typedef struct _GtkStyleContext GtkStyleContext;
typedef struct _GtkStyleContextClass GtkStyleContextClass;
struct _GtkStyleContextClass { uint8_t _data[176]; };
typedef struct _GtkStyleContextPrivate GtkStyleContextPrivate;
struct _GtkStyleContextPrivate {};
typedef struct _GtkStyleProperties GtkStyleProperties;
typedef struct _GtkStylePropertiesClass GtkStylePropertiesClass;
struct _GtkStylePropertiesClass { uint8_t _data[168]; };
typedef void* GtkStylePropertyParser;
extern void _GtkStylePropertyParser_c_wrapper();
extern void _GtkStylePropertyParser_c_wrapper_once();
typedef struct _GtkStyleProvider GtkStyleProvider;
typedef struct _GtkStyleProviderIface GtkStyleProviderIface;
struct _GtkStyleProviderIface { uint8_t _data[40]; };
typedef struct _GtkSwitch GtkSwitch;
typedef struct _GtkSwitchClass GtkSwitchClass;
struct _GtkSwitchClass { uint8_t _data[880]; };
typedef struct _GtkSwitchPrivate GtkSwitchPrivate;
struct _GtkSwitchPrivate {};
typedef struct _GtkSymbolicColor GtkSymbolicColor;
struct _GtkSymbolicColor {};
typedef struct _GtkTable GtkTable;
typedef struct _GtkTableChild GtkTableChild;
struct _GtkTableChild { uint8_t _data[48]; };
typedef struct _GtkTableClass GtkTableClass;
struct _GtkTableClass { uint8_t _data[1008]; };
typedef struct _GtkTablePrivate GtkTablePrivate;
struct _GtkTablePrivate {};
typedef struct _GtkTableRowCol GtkTableRowCol;
struct _GtkTableRowCol { uint8_t _data[28]; };
typedef struct _GtkTargetEntry GtkTargetEntry;
struct _GtkTargetEntry { uint8_t _data[16]; };
typedef uint32_t GtkTargetFlags;
typedef struct _GtkTargetList GtkTargetList;
struct _GtkTargetList {};
typedef struct _GtkTearoffMenuItem GtkTearoffMenuItem;
typedef struct _GtkTearoffMenuItemClass GtkTearoffMenuItemClass;
struct _GtkTearoffMenuItemClass { uint8_t _data[1144]; };
typedef struct _GtkTearoffMenuItemPrivate GtkTearoffMenuItemPrivate;
struct _GtkTearoffMenuItemPrivate {};
typedef struct _GtkTextAppearance GtkTextAppearance;
struct _GtkTextAppearance { uint8_t _data[64]; };
typedef struct _GtkTextAttributes GtkTextAttributes;
struct _GtkTextAttributes { uint8_t _data[192]; };
typedef struct _GtkTextBTree GtkTextBTree;
struct _GtkTextBTree {};
typedef struct _GtkTextBuffer GtkTextBuffer;
typedef struct _GtkTextBufferClass GtkTextBufferClass;
struct _GtkTextBufferClass { uint8_t _data[272]; };
typedef void* GtkTextBufferDeserializeFunc;
extern void _GtkTextBufferDeserializeFunc_c_wrapper();
extern void _GtkTextBufferDeserializeFunc_c_wrapper_once();
typedef struct _GtkTextBufferPrivate GtkTextBufferPrivate;
struct _GtkTextBufferPrivate {};
typedef void* GtkTextBufferSerializeFunc;
extern void _GtkTextBufferSerializeFunc_c_wrapper();
extern void _GtkTextBufferSerializeFunc_c_wrapper_once();
typedef int32_t GtkTextBufferTargetInfo;
typedef void* GtkTextCharPredicate;
extern void _GtkTextCharPredicate_c_wrapper();
extern void _GtkTextCharPredicate_c_wrapper_once();
typedef struct _GtkTextChildAnchor GtkTextChildAnchor;
typedef struct _GtkTextChildAnchorClass GtkTextChildAnchorClass;
struct _GtkTextChildAnchorClass { uint8_t _data[168]; };
typedef uint32_t GtkTextDirection;
typedef struct _GtkTextIter GtkTextIter;
struct _GtkTextIter { uint8_t _data[80]; };
typedef struct _GtkTextMark GtkTextMark;
typedef struct _GtkTextMarkClass GtkTextMarkClass;
struct _GtkTextMarkClass { uint8_t _data[168]; };
typedef uint32_t GtkTextSearchFlags;
typedef struct _GtkTextTag GtkTextTag;
typedef struct _GtkTextTagClass GtkTextTagClass;
struct _GtkTextTagClass { uint8_t _data[176]; };
typedef struct _GtkTextTagPrivate GtkTextTagPrivate;
struct _GtkTextTagPrivate {};
typedef struct _GtkTextTagTable GtkTextTagTable;
typedef struct _GtkTextTagTableClass GtkTextTagTableClass;
struct _GtkTextTagTableClass { uint8_t _data[192]; };
typedef void* GtkTextTagTableForeach;
extern void _GtkTextTagTableForeach_c_wrapper();
extern void _GtkTextTagTableForeach_c_wrapper_once();
typedef struct _GtkTextTagTablePrivate GtkTextTagTablePrivate;
struct _GtkTextTagTablePrivate {};
typedef struct _GtkTextView GtkTextView;
typedef struct _GtkTextViewClass GtkTextViewClass;
struct _GtkTextViewClass { uint8_t _data[1120]; };
typedef struct _GtkTextViewPrivate GtkTextViewPrivate;
struct _GtkTextViewPrivate {};
typedef uint32_t GtkTextWindowType;
typedef struct _GtkThemeEngine GtkThemeEngine;
struct _GtkThemeEngine {};
typedef struct _GtkThemingEngine GtkThemingEngine;
typedef struct _GtkThemingEngineClass GtkThemingEngineClass;
struct _GtkThemingEngineClass { uint8_t _data[384]; };
typedef struct _GtkThemingEnginePrivate GtkThemingEnginePrivate;
struct _GtkThemingEnginePrivate {};
typedef struct _GtkToggleAction GtkToggleAction;
typedef struct _GtkToggleActionClass GtkToggleActionClass;
struct _GtkToggleActionClass { uint8_t _data[272]; };
typedef struct _GtkToggleActionEntry GtkToggleActionEntry;
struct _GtkToggleActionEntry { uint8_t _data[56]; };
typedef struct _GtkToggleActionPrivate GtkToggleActionPrivate;
struct _GtkToggleActionPrivate {};
typedef struct _GtkToggleButton GtkToggleButton;
typedef struct _GtkToggleButtonClass GtkToggleButtonClass;
struct _GtkToggleButtonClass { uint8_t _data[1128]; };
typedef struct _GtkToggleButtonPrivate GtkToggleButtonPrivate;
struct _GtkToggleButtonPrivate {};
typedef struct _GtkToggleToolButton GtkToggleToolButton;
typedef struct _GtkToggleToolButtonClass GtkToggleToolButtonClass;
struct _GtkToggleToolButtonClass { uint8_t _data[1144]; };
typedef struct _GtkToggleToolButtonPrivate GtkToggleToolButtonPrivate;
struct _GtkToggleToolButtonPrivate {};
typedef struct _GtkToolButton GtkToolButton;
typedef struct _GtkToolButtonClass GtkToolButtonClass;
struct _GtkToolButtonClass { uint8_t _data[1104]; };
typedef struct _GtkToolButtonPrivate GtkToolButtonPrivate;
struct _GtkToolButtonPrivate {};
typedef struct _GtkToolItem GtkToolItem;
typedef struct _GtkToolItemClass GtkToolItemClass;
struct _GtkToolItemClass { uint8_t _data[1056]; };
typedef struct _GtkToolItemGroup GtkToolItemGroup;
typedef struct _GtkToolItemGroupClass GtkToolItemGroupClass;
struct _GtkToolItemGroupClass { uint8_t _data[1008]; };
typedef struct _GtkToolItemGroupPrivate GtkToolItemGroupPrivate;
struct _GtkToolItemGroupPrivate {};
typedef struct _GtkToolItemPrivate GtkToolItemPrivate;
struct _GtkToolItemPrivate {};
typedef struct _GtkToolPalette GtkToolPalette;
typedef struct _GtkToolPaletteClass GtkToolPaletteClass;
struct _GtkToolPaletteClass { uint8_t _data[1008]; };
typedef uint32_t GtkToolPaletteDragTargets;
typedef struct _GtkToolPalettePrivate GtkToolPalettePrivate;
struct _GtkToolPalettePrivate {};
typedef struct _GtkToolShell GtkToolShell;
typedef struct _GtkToolShellIface GtkToolShellIface;
struct _GtkToolShellIface { uint8_t _data[88]; };
typedef struct _GtkToolbar GtkToolbar;
typedef struct _GtkToolbarClass GtkToolbarClass;
struct _GtkToolbarClass { uint8_t _data[1032]; };
typedef struct _GtkToolbarPrivate GtkToolbarPrivate;
struct _GtkToolbarPrivate {};
typedef uint32_t GtkToolbarSpaceStyle;
typedef uint32_t GtkToolbarStyle;
typedef struct _GtkTooltip GtkTooltip;
typedef void* GtkTranslateFunc;
extern void _GtkTranslateFunc_c_wrapper();
extern void _GtkTranslateFunc_c_wrapper_once();
typedef void* GtkTreeCellDataFunc;
extern void _GtkTreeCellDataFunc_c_wrapper();
extern void _GtkTreeCellDataFunc_c_wrapper_once();
typedef void* GtkTreeDestroyCountFunc;
extern void _GtkTreeDestroyCountFunc_c_wrapper();
extern void _GtkTreeDestroyCountFunc_c_wrapper_once();
typedef struct _GtkTreeDragDest GtkTreeDragDest;
typedef struct _GtkTreeDragDestIface GtkTreeDragDestIface;
struct _GtkTreeDragDestIface { uint8_t _data[32]; };
typedef struct _GtkTreeDragSource GtkTreeDragSource;
typedef struct _GtkTreeDragSourceIface GtkTreeDragSourceIface;
struct _GtkTreeDragSourceIface { uint8_t _data[40]; };
typedef struct _GtkTreeIter GtkTreeIter;
struct _GtkTreeIter { uint8_t _data[32]; };
typedef void* GtkTreeIterCompareFunc;
extern void _GtkTreeIterCompareFunc_c_wrapper();
extern void _GtkTreeIterCompareFunc_c_wrapper_once();
typedef struct _GtkTreeModel GtkTreeModel;
typedef struct _GtkTreeModelFilter GtkTreeModelFilter;
typedef struct _GtkTreeModelFilterClass GtkTreeModelFilterClass;
struct _GtkTreeModelFilterClass { uint8_t _data[184]; };
typedef void* GtkTreeModelFilterModifyFunc;
extern void _GtkTreeModelFilterModifyFunc_c_wrapper();
extern void _GtkTreeModelFilterModifyFunc_c_wrapper_once();
typedef struct _GtkTreeModelFilterPrivate GtkTreeModelFilterPrivate;
struct _GtkTreeModelFilterPrivate {};
typedef void* GtkTreeModelFilterVisibleFunc;
extern void _GtkTreeModelFilterVisibleFunc_c_wrapper();
extern void _GtkTreeModelFilterVisibleFunc_c_wrapper_once();
typedef uint32_t GtkTreeModelFlags;
typedef void* GtkTreeModelForeachFunc;
extern void _GtkTreeModelForeachFunc_c_wrapper();
extern void _GtkTreeModelForeachFunc_c_wrapper_once();
typedef struct _GtkTreeModelIface GtkTreeModelIface;
struct _GtkTreeModelIface { uint8_t _data[176]; };
typedef struct _GtkTreeModelSort GtkTreeModelSort;
typedef struct _GtkTreeModelSortClass GtkTreeModelSortClass;
struct _GtkTreeModelSortClass { uint8_t _data[168]; };
typedef struct _GtkTreeModelSortPrivate GtkTreeModelSortPrivate;
struct _GtkTreeModelSortPrivate {};
typedef struct _GtkTreePath GtkTreePath;
struct _GtkTreePath {};
typedef struct _GtkTreeRowReference GtkTreeRowReference;
struct _GtkTreeRowReference {};
typedef struct _GtkTreeSelection GtkTreeSelection;
typedef struct _GtkTreeSelectionClass GtkTreeSelectionClass;
struct _GtkTreeSelectionClass { uint8_t _data[176]; };
typedef void* GtkTreeSelectionForeachFunc;
extern void _GtkTreeSelectionForeachFunc_c_wrapper();
extern void _GtkTreeSelectionForeachFunc_c_wrapper_once();
typedef void* GtkTreeSelectionFunc;
extern void _GtkTreeSelectionFunc_c_wrapper();
extern void _GtkTreeSelectionFunc_c_wrapper_once();
typedef struct _GtkTreeSelectionPrivate GtkTreeSelectionPrivate;
struct _GtkTreeSelectionPrivate {};
typedef struct _GtkTreeSortable GtkTreeSortable;
typedef struct _GtkTreeSortableIface GtkTreeSortableIface;
struct _GtkTreeSortableIface { uint8_t _data[64]; };
typedef struct _GtkTreeStore GtkTreeStore;
typedef struct _GtkTreeStoreClass GtkTreeStoreClass;
struct _GtkTreeStoreClass { uint8_t _data[168]; };
typedef struct _GtkTreeStorePrivate GtkTreeStorePrivate;
struct _GtkTreeStorePrivate {};
typedef struct _GtkTreeView GtkTreeView;
typedef struct _GtkTreeViewClass GtkTreeViewClass;
struct _GtkTreeViewClass { uint8_t _data[1160]; };
typedef struct _GtkTreeViewColumn GtkTreeViewColumn;
typedef struct _GtkTreeViewColumnClass GtkTreeViewColumnClass;
struct _GtkTreeViewColumnClass { uint8_t _data[176]; };
typedef void* GtkTreeViewColumnDropFunc;
extern void _GtkTreeViewColumnDropFunc_c_wrapper();
extern void _GtkTreeViewColumnDropFunc_c_wrapper_once();
typedef struct _GtkTreeViewColumnPrivate GtkTreeViewColumnPrivate;
struct _GtkTreeViewColumnPrivate {};
typedef uint32_t GtkTreeViewColumnSizing;
typedef uint32_t GtkTreeViewDropPosition;
typedef uint32_t GtkTreeViewGridLines;
typedef void* GtkTreeViewMappingFunc;
extern void _GtkTreeViewMappingFunc_c_wrapper();
extern void _GtkTreeViewMappingFunc_c_wrapper_once();
typedef struct _GtkTreeViewPrivate GtkTreeViewPrivate;
struct _GtkTreeViewPrivate {};
typedef void* GtkTreeViewRowSeparatorFunc;
extern void _GtkTreeViewRowSeparatorFunc_c_wrapper();
extern void _GtkTreeViewRowSeparatorFunc_c_wrapper_once();
typedef void* GtkTreeViewSearchEqualFunc;
extern void _GtkTreeViewSearchEqualFunc_c_wrapper();
extern void _GtkTreeViewSearchEqualFunc_c_wrapper_once();
typedef void* GtkTreeViewSearchPositionFunc;
extern void _GtkTreeViewSearchPositionFunc_c_wrapper();
extern void _GtkTreeViewSearchPositionFunc_c_wrapper_once();
typedef struct _GtkUIManager GtkUIManager;
typedef struct _GtkUIManagerClass GtkUIManagerClass;
struct _GtkUIManagerClass { uint8_t _data[232]; };
typedef uint32_t GtkUIManagerItemType;
typedef struct _GtkUIManagerPrivate GtkUIManagerPrivate;
struct _GtkUIManagerPrivate {};
typedef uint32_t GtkUnit;
typedef struct _GtkVBox GtkVBox;
typedef struct _GtkVBoxClass GtkVBoxClass;
struct _GtkVBoxClass { uint8_t _data[1008]; };
typedef struct _GtkVButtonBox GtkVButtonBox;
typedef struct _GtkVButtonBoxClass GtkVButtonBoxClass;
struct _GtkVButtonBoxClass { uint8_t _data[1040]; };
typedef struct _GtkVPaned GtkVPaned;
typedef struct _GtkVPanedClass GtkVPanedClass;
struct _GtkVPanedClass { uint8_t _data[1056]; };
typedef struct _GtkVScale GtkVScale;
typedef struct _GtkVScaleClass GtkVScaleClass;
struct _GtkVScaleClass { uint8_t _data[968]; };
typedef struct _GtkVScrollbar GtkVScrollbar;
typedef struct _GtkVScrollbarClass GtkVScrollbarClass;
struct _GtkVScrollbarClass { uint8_t _data[944]; };
typedef struct _GtkVSeparator GtkVSeparator;
typedef struct _GtkVSeparatorClass GtkVSeparatorClass;
struct _GtkVSeparatorClass { uint8_t _data[856]; };
typedef struct _GtkViewport GtkViewport;
typedef struct _GtkViewportClass GtkViewportClass;
struct _GtkViewportClass { uint8_t _data[1040]; };
typedef struct _GtkViewportPrivate GtkViewportPrivate;
struct _GtkViewportPrivate {};
typedef struct _GtkVolumeButton GtkVolumeButton;
typedef struct _GtkVolumeButtonClass GtkVolumeButtonClass;
struct _GtkVolumeButtonClass { uint8_t _data[1160]; };
typedef struct _GtkWidget GtkWidget;
typedef struct _GtkWidgetAuxInfo GtkWidgetAuxInfo;
struct _GtkWidgetAuxInfo { uint8_t _data[24]; };
typedef struct _GtkWidgetClass GtkWidgetClass;
struct _GtkWidgetClass { uint8_t _data[824]; };
typedef struct _GtkWidgetClassPrivate GtkWidgetClassPrivate;
struct _GtkWidgetClassPrivate {};
typedef uint32_t GtkWidgetHelpType;
typedef struct _GtkWidgetPath GtkWidgetPath;
struct _GtkWidgetPath {};
typedef struct _GtkWidgetPrivate GtkWidgetPrivate;
struct _GtkWidgetPrivate {};
typedef struct _GtkWindow GtkWindow;
typedef struct _GtkWindowClass GtkWindowClass;
struct _GtkWindowClass { uint8_t _data[1072]; };
typedef struct _GtkWindowGeometryInfo GtkWindowGeometryInfo;
struct _GtkWindowGeometryInfo {};
typedef struct _GtkWindowGroup GtkWindowGroup;
typedef struct _GtkWindowGroupClass GtkWindowGroupClass;
struct _GtkWindowGroupClass { uint8_t _data[168]; };
typedef struct _GtkWindowGroupPrivate GtkWindowGroupPrivate;
struct _GtkWindowGroupPrivate {};
typedef uint32_t GtkWindowPosition;
typedef struct _GtkWindowPrivate GtkWindowPrivate;
struct _GtkWindowPrivate {};
typedef uint32_t GtkWindowType;
typedef uint32_t GtkWrapMode;
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
typedef struct _GdkAppLaunchContext GdkAppLaunchContext;
typedef void *GdkAtom;
typedef uint32_t GdkAxisUse;
typedef uint32_t GdkByteOrder;
typedef struct _GdkColor GdkColor;
struct _GdkColor { uint8_t _data[12]; };
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
struct _GdkEvent { uint8_t _data[88]; };
typedef struct _GdkEventAny GdkEventAny;
struct _GdkEventAny { uint8_t _data[24]; };
typedef struct _GdkEventButton GdkEventButton;
struct _GdkEventButton { uint8_t _data[80]; };
typedef struct _GdkEventConfigure GdkEventConfigure;
struct _GdkEventConfigure { uint8_t _data[40]; };
typedef struct _GdkEventCrossing GdkEventCrossing;
struct _GdkEventCrossing { uint8_t _data[88]; };
typedef struct _GdkEventDND GdkEventDND;
struct _GdkEventDND { uint8_t _data[40]; };
typedef struct _GdkEventExpose GdkEventExpose;
struct _GdkEventExpose { uint8_t _data[56]; };
typedef struct _GdkEventFocus GdkEventFocus;
struct _GdkEventFocus { uint8_t _data[24]; };
typedef void* GdkEventFunc;
extern void _GdkEventFunc_c_wrapper();
extern void _GdkEventFunc_c_wrapper_once();
typedef struct _GdkEventGrabBroken GdkEventGrabBroken;
struct _GdkEventGrabBroken { uint8_t _data[40]; };
typedef struct _GdkEventKey GdkEventKey;
struct _GdkEventKey { uint8_t _data[56]; };
typedef uint32_t GdkEventMask;
typedef struct _GdkEventMotion GdkEventMotion;
struct _GdkEventMotion { uint8_t _data[80]; };
typedef struct _GdkEventOwnerChange GdkEventOwnerChange;
struct _GdkEventOwnerChange { uint8_t _data[56]; };
typedef struct _GdkEventProperty GdkEventProperty;
struct _GdkEventProperty { uint8_t _data[40]; };
typedef struct _GdkEventProximity GdkEventProximity;
struct _GdkEventProximity { uint8_t _data[32]; };
typedef struct _GdkEventScroll GdkEventScroll;
struct _GdkEventScroll { uint8_t _data[72]; };
typedef struct _GdkEventSelection GdkEventSelection;
struct _GdkEventSelection { uint8_t _data[64]; };
typedef struct _GdkEventSetting GdkEventSetting;
struct _GdkEventSetting { uint8_t _data[32]; };
typedef int32_t GdkEventType;
typedef struct _GdkEventVisibility GdkEventVisibility;
struct _GdkEventVisibility { uint8_t _data[24]; };
typedef struct _GdkEventWindowState GdkEventWindowState;
struct _GdkEventWindowState { uint8_t _data[32]; };
typedef uint32_t GdkExtensionMode;
typedef void* GdkFilterFunc;
extern void _GdkFilterFunc_c_wrapper();
extern void _GdkFilterFunc_c_wrapper_once();
typedef uint32_t GdkFilterReturn;
typedef struct _GdkGeometry GdkGeometry;
struct _GdkGeometry { uint8_t _data[56]; };
typedef uint32_t GdkGrabOwnership;
typedef uint32_t GdkGrabStatus;
typedef uint32_t GdkGravity;
typedef uint32_t GdkInputMode;
typedef uint32_t GdkInputSource;
typedef struct _GdkKeymap GdkKeymap;
typedef struct _GdkKeymapKey GdkKeymapKey;
struct _GdkKeymapKey { uint8_t _data[12]; };
typedef uint32_t GdkModifierType;
typedef uint32_t GdkNotifyType;
typedef uint32_t GdkOwnerChange;
typedef struct _GdkPoint GdkPoint;
struct _GdkPoint { uint8_t _data[8]; };
typedef uint32_t GdkPropMode;
typedef uint32_t GdkPropertyState;
typedef struct _GdkRGBA GdkRGBA;
struct _GdkRGBA { uint8_t _data[32]; };
typedef struct _GdkScreen GdkScreen;
typedef uint32_t GdkScrollDirection;
typedef uint32_t GdkSettingAction;
typedef int32_t GdkStatus;
typedef struct _GdkTimeCoord GdkTimeCoord;
struct _GdkTimeCoord { uint8_t _data[1032]; };
typedef uint32_t GdkVisibilityState;
typedef struct _GdkVisual GdkVisual;
typedef uint32_t GdkVisualType;
typedef uint32_t GdkWMDecoration;
typedef uint32_t GdkWMFunction;
typedef struct _GdkWindow GdkWindow;
typedef struct _GdkWindowAttr GdkWindowAttr;
struct _GdkWindowAttr { uint8_t _data[80]; };
typedef uint32_t GdkWindowAttributesType;
typedef void* GdkWindowChildFunc;
extern void _GdkWindowChildFunc_c_wrapper();
extern void _GdkWindowChildFunc_c_wrapper_once();
typedef struct _GdkWindowClass GdkWindowClass;
struct _GdkWindowClass { uint8_t _data[232]; };
typedef uint32_t GdkWindowEdge;
typedef uint32_t GdkWindowHints;
typedef struct _GdkWindowRedirect GdkWindowRedirect;
struct _GdkWindowRedirect {};
typedef uint32_t GdkWindowState;
typedef uint32_t GdkWindowType;
typedef uint32_t GdkWindowTypeHint;
typedef uint32_t GdkWindowWindowClass;
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
typedef struct _AtkAction AtkAction;
typedef struct _AtkActionIface AtkActionIface;
struct _AtkActionIface { uint8_t _data[80]; };
typedef struct _AtkAttribute AtkAttribute;
struct _AtkAttribute { uint8_t _data[16]; };
typedef struct _AtkComponent AtkComponent;
typedef struct _AtkComponentIface AtkComponentIface;
struct _AtkComponentIface { uint8_t _data[136]; };
typedef uint32_t AtkCoordType;
typedef struct _AtkDocument AtkDocument;
typedef struct _AtkDocumentIface AtkDocumentIface;
struct _AtkDocumentIface { uint8_t _data[96]; };
typedef struct _AtkEditableText AtkEditableText;
typedef struct _AtkEditableTextIface AtkEditableTextIface;
struct _AtkEditableTextIface { uint8_t _data[88]; };
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
struct _AtkGObjectAccessibleClass { uint8_t _data[368]; };
typedef struct _AtkHyperlink AtkHyperlink;
typedef struct _AtkHyperlinkClass AtkHyperlinkClass;
struct _AtkHyperlinkClass { uint8_t _data[216]; };
typedef struct _AtkHyperlinkImpl AtkHyperlinkImpl;
typedef struct _AtkHyperlinkImplIface AtkHyperlinkImplIface;
struct _AtkHyperlinkImplIface { uint8_t _data[32]; };
typedef uint32_t AtkHyperlinkStateFlags;
typedef struct _AtkHypertext AtkHypertext;
typedef struct _AtkHypertextIface AtkHypertextIface;
struct _AtkHypertextIface { uint8_t _data[72]; };
typedef struct _AtkImage AtkImage;
typedef struct _AtkImageIface AtkImageIface;
struct _AtkImageIface { uint8_t _data[64]; };
typedef struct _AtkImplementor AtkImplementor;
struct _AtkImplementor {};
typedef struct _AtkImplementorIface AtkImplementorIface;
typedef struct _AtkKeyEventStruct AtkKeyEventStruct;
struct _AtkKeyEventStruct { uint8_t _data[32]; };
typedef uint32_t AtkKeyEventType;
typedef void* AtkKeySnoopFunc;
extern void _AtkKeySnoopFunc_c_wrapper();
extern void _AtkKeySnoopFunc_c_wrapper_once();
typedef uint32_t AtkLayer;
typedef struct _AtkMisc AtkMisc;
typedef struct _AtkMiscClass AtkMiscClass;
struct _AtkMiscClass { uint8_t _data[408]; };
typedef struct _AtkNoOpObject AtkNoOpObject;
typedef struct _AtkNoOpObjectClass AtkNoOpObjectClass;
struct _AtkNoOpObjectClass { uint8_t _data[352]; };
typedef struct _AtkNoOpObjectFactory AtkNoOpObjectFactory;
typedef struct _AtkNoOpObjectFactoryClass AtkNoOpObjectFactoryClass;
struct _AtkNoOpObjectFactoryClass { uint8_t _data[176]; };
typedef struct _AtkObject AtkObject;
typedef struct _AtkObjectClass AtkObjectClass;
struct _AtkObjectClass { uint8_t _data[352]; };
typedef struct _AtkObjectFactory AtkObjectFactory;
typedef struct _AtkObjectFactoryClass AtkObjectFactoryClass;
struct _AtkObjectFactoryClass { uint8_t _data[176]; };
typedef struct _AtkPlug AtkPlug;
typedef struct _AtkPlugClass AtkPlugClass;
struct _AtkPlugClass { uint8_t _data[360]; };
typedef struct _AtkRectangle AtkRectangle;
struct _AtkRectangle { uint8_t _data[16]; };
typedef struct _AtkRegistry AtkRegistry;
typedef struct _AtkRelation AtkRelation;
typedef struct _AtkRelationClass AtkRelationClass;
struct _AtkRelationClass { uint8_t _data[136]; };
typedef struct _AtkRelationSet AtkRelationSet;
typedef struct _AtkRelationSetClass AtkRelationSetClass;
struct _AtkRelationSetClass { uint8_t _data[152]; };
typedef uint32_t AtkRelationType;
typedef uint32_t AtkRole;
typedef struct _AtkSelection AtkSelection;
typedef struct _AtkSelectionIface AtkSelectionIface;
struct _AtkSelectionIface { uint8_t _data[96]; };
typedef struct _AtkSocket AtkSocket;
typedef struct _AtkSocketClass AtkSocketClass;
struct _AtkSocketClass { uint8_t _data[360]; };
typedef struct _AtkStateSet AtkStateSet;
typedef struct _AtkStateSetClass AtkStateSetClass;
struct _AtkStateSetClass { uint8_t _data[136]; };
typedef uint32_t AtkStateType;
typedef struct _AtkStreamableContent AtkStreamableContent;
typedef struct _AtkStreamableContentIface AtkStreamableContentIface;
struct _AtkStreamableContentIface { uint8_t _data[72]; };
typedef struct _AtkTable AtkTable;
typedef struct _AtkTableIface AtkTableIface;
struct _AtkTableIface { uint8_t _data[336]; };
typedef struct _AtkText AtkText;
typedef uint32_t AtkTextAttribute;
typedef uint32_t AtkTextBoundary;
typedef uint32_t AtkTextClipType;
typedef struct _AtkTextIface AtkTextIface;
struct _AtkTextIface { uint8_t _data[208]; };
typedef struct _AtkTextRange AtkTextRange;
struct _AtkTextRange { uint8_t _data[32]; };
typedef struct _AtkTextRectangle AtkTextRectangle;
struct _AtkTextRectangle { uint8_t _data[16]; };
typedef struct _AtkUtil AtkUtil;
typedef struct _AtkUtilClass AtkUtilClass;
struct _AtkUtilClass { uint8_t _data[192]; };
typedef struct _AtkValue AtkValue;
typedef struct _AtkValueIface AtkValueIface;
struct _AtkValueIface { uint8_t _data[64]; };
typedef struct _AtkWindow AtkWindow;
typedef struct _AtkWindowIface AtkWindowIface;
struct _AtkWindowIface { uint8_t _data[144]; };
typedef struct _Atk_PropertyValues Atk_PropertyValues;
struct _Atk_PropertyValues { uint8_t _data[56]; };
typedef struct _Atk_Registry Atk_Registry;
struct _Atk_Registry { uint8_t _data[40]; };
typedef struct _Atk_RegistryClass Atk_RegistryClass;
struct _Atk_RegistryClass { uint8_t _data[136]; };
typedef uint32_t GtkSourceBracketMatchType;
typedef struct _GtkSourceBuffer GtkSourceBuffer;
typedef struct _GtkSourceBufferClass GtkSourceBufferClass;
typedef struct _GtkSourceBufferPrivate GtkSourceBufferPrivate;
typedef struct _GtkSourceCompletion GtkSourceCompletion;
typedef uint32_t GtkSourceCompletionActivation;
typedef struct _GtkSourceCompletionClass GtkSourceCompletionClass;
typedef struct _GtkSourceCompletionContext GtkSourceCompletionContext;
typedef struct _GtkSourceCompletionContextClass GtkSourceCompletionContextClass;
typedef struct _GtkSourceCompletionContextPrivate GtkSourceCompletionContextPrivate;
typedef uint32_t GtkSourceCompletionError;
typedef struct _GtkSourceCompletionInfo GtkSourceCompletionInfo;
typedef struct _GtkSourceCompletionInfoClass GtkSourceCompletionInfoClass;
typedef struct _GtkSourceCompletionInfoPrivate GtkSourceCompletionInfoPrivate;
typedef struct _GtkSourceCompletionItem GtkSourceCompletionItem;
typedef struct _GtkSourceCompletionItemClass GtkSourceCompletionItemClass;
typedef struct _GtkSourceCompletionItemPrivate GtkSourceCompletionItemPrivate;
typedef struct _GtkSourceCompletionPrivate GtkSourceCompletionPrivate;
typedef struct _GtkSourceCompletionProposal GtkSourceCompletionProposal;
typedef struct _GtkSourceCompletionProposalIface GtkSourceCompletionProposalIface;
typedef struct _GtkSourceCompletionProvider GtkSourceCompletionProvider;
typedef struct _GtkSourceCompletionProviderIface GtkSourceCompletionProviderIface;
typedef struct _GtkSourceCompletionWords GtkSourceCompletionWords;
typedef struct _GtkSourceCompletionWordsClass GtkSourceCompletionWordsClass;
typedef struct _GtkSourceCompletionWordsPrivate GtkSourceCompletionWordsPrivate;
typedef uint32_t GtkSourceDrawSpacesFlags;
typedef struct _GtkSourceGutter GtkSourceGutter;
typedef struct _GtkSourceGutterClass GtkSourceGutterClass;
typedef struct _GtkSourceGutterPrivate GtkSourceGutterPrivate;
typedef struct _GtkSourceGutterRenderer GtkSourceGutterRenderer;
typedef uint32_t GtkSourceGutterRendererAlignmentMode;
typedef struct _GtkSourceGutterRendererClass GtkSourceGutterRendererClass;
typedef struct _GtkSourceGutterRendererPixbuf GtkSourceGutterRendererPixbuf;
typedef struct _GtkSourceGutterRendererPixbufClass GtkSourceGutterRendererPixbufClass;
typedef struct _GtkSourceGutterRendererPixbufPrivate GtkSourceGutterRendererPixbufPrivate;
typedef struct _GtkSourceGutterRendererPrivate GtkSourceGutterRendererPrivate;
typedef uint32_t GtkSourceGutterRendererState;
typedef struct _GtkSourceGutterRendererText GtkSourceGutterRendererText;
typedef struct _GtkSourceGutterRendererTextClass GtkSourceGutterRendererTextClass;
typedef struct _GtkSourceGutterRendererTextPrivate GtkSourceGutterRendererTextPrivate;
typedef struct _GtkSourceLanguage GtkSourceLanguage;
typedef struct _GtkSourceLanguageClass GtkSourceLanguageClass;
typedef struct _GtkSourceLanguageManager GtkSourceLanguageManager;
typedef struct _GtkSourceLanguageManagerClass GtkSourceLanguageManagerClass;
typedef struct _GtkSourceLanguageManagerPrivate GtkSourceLanguageManagerPrivate;
typedef struct _GtkSourceLanguagePrivate GtkSourceLanguagePrivate;
typedef struct _GtkSourceMark GtkSourceMark;
typedef struct _GtkSourceMarkAttributes GtkSourceMarkAttributes;
typedef struct _GtkSourceMarkAttributesClass GtkSourceMarkAttributesClass;
typedef struct _GtkSourceMarkAttributesPrivate GtkSourceMarkAttributesPrivate;
typedef struct _GtkSourceMarkClass GtkSourceMarkClass;
typedef struct _GtkSourceMarkPrivate GtkSourceMarkPrivate;
typedef struct _GtkSourcePrintCompositor GtkSourcePrintCompositor;
typedef struct _GtkSourcePrintCompositorClass GtkSourcePrintCompositorClass;
typedef struct _GtkSourcePrintCompositorPrivate GtkSourcePrintCompositorPrivate;
typedef uint32_t GtkSourceSmartHomeEndType;
typedef struct _GtkSourceStyle GtkSourceStyle;
typedef struct _GtkSourceStyleScheme GtkSourceStyleScheme;
typedef struct _GtkSourceStyleSchemeClass GtkSourceStyleSchemeClass;
typedef struct _GtkSourceStyleSchemeManager GtkSourceStyleSchemeManager;
typedef struct _GtkSourceStyleSchemeManagerClass GtkSourceStyleSchemeManagerClass;
typedef struct _GtkSourceStyleSchemeManagerPrivate GtkSourceStyleSchemeManagerPrivate;
typedef struct _GtkSourceStyleSchemePrivate GtkSourceStyleSchemePrivate;
typedef struct _GtkSourceUndoManager GtkSourceUndoManager;
typedef struct _GtkSourceUndoManagerIface GtkSourceUndoManagerIface;
typedef struct _GtkSourceView GtkSourceView;
typedef struct _GtkSourceViewClass GtkSourceViewClass;
typedef int32_t GtkSourceViewGutterPosition;
typedef struct _GtkSourceViewPrivate GtkSourceViewPrivate;
extern GtkSourceBuffer* gtk_source_buffer_new(GtkTextTagTable*);
extern GtkSourceBuffer* gtk_source_buffer_new_with_language(GtkSourceLanguage*);
extern int gtk_source_buffer_backward_iter_to_source_mark(GtkSourceBuffer*, GtkTextIter*, char*);
extern void gtk_source_buffer_begin_not_undoable_action(GtkSourceBuffer*);
extern int gtk_source_buffer_can_redo(GtkSourceBuffer*);
extern int gtk_source_buffer_can_undo(GtkSourceBuffer*);
extern GtkSourceMark* gtk_source_buffer_create_source_mark(GtkSourceBuffer*, char*, char*, GtkTextIter*);
extern void gtk_source_buffer_end_not_undoable_action(GtkSourceBuffer*);
extern void gtk_source_buffer_ensure_highlight(GtkSourceBuffer*, GtkTextIter*, GtkTextIter*);
extern int gtk_source_buffer_forward_iter_to_source_mark(GtkSourceBuffer*, GtkTextIter*, char*);
extern char** gtk_source_buffer_get_context_classes_at_iter(GtkSourceBuffer*, GtkTextIter*);
extern int gtk_source_buffer_get_highlight_matching_brackets(GtkSourceBuffer*);
extern int gtk_source_buffer_get_highlight_syntax(GtkSourceBuffer*);
extern GtkSourceLanguage* gtk_source_buffer_get_language(GtkSourceBuffer*);
extern int32_t gtk_source_buffer_get_max_undo_levels(GtkSourceBuffer*);
extern GSList* gtk_source_buffer_get_source_marks_at_iter(GtkSourceBuffer*, GtkTextIter*, char*);
extern GSList* gtk_source_buffer_get_source_marks_at_line(GtkSourceBuffer*, int32_t, char*);
extern GtkSourceStyleScheme* gtk_source_buffer_get_style_scheme(GtkSourceBuffer*);
extern GtkSourceUndoManager* gtk_source_buffer_get_undo_manager(GtkSourceBuffer*);
extern int gtk_source_buffer_iter_backward_to_context_class_toggle(GtkSourceBuffer*, GtkTextIter*, char*);
extern int gtk_source_buffer_iter_forward_to_context_class_toggle(GtkSourceBuffer*, GtkTextIter*, char*);
extern int gtk_source_buffer_iter_has_context_class(GtkSourceBuffer*, GtkTextIter*, char*);
extern void gtk_source_buffer_redo(GtkSourceBuffer*);
extern void gtk_source_buffer_remove_source_marks(GtkSourceBuffer*, GtkTextIter*, GtkTextIter*, char*);
extern void gtk_source_buffer_set_highlight_matching_brackets(GtkSourceBuffer*, int);
extern void gtk_source_buffer_set_highlight_syntax(GtkSourceBuffer*, int);
extern void gtk_source_buffer_set_language(GtkSourceBuffer*, GtkSourceLanguage*);
extern void gtk_source_buffer_set_max_undo_levels(GtkSourceBuffer*, int32_t);
extern void gtk_source_buffer_set_style_scheme(GtkSourceBuffer*, GtkSourceStyleScheme*);
extern void gtk_source_buffer_set_undo_manager(GtkSourceBuffer*, GtkSourceUndoManager*);
extern void gtk_source_buffer_undo(GtkSourceBuffer*);
extern GType gtk_source_buffer_get_type();
extern int gtk_source_completion_add_provider(GtkSourceCompletion*, GtkSourceCompletionProvider*, GError**);
extern void gtk_source_completion_block_interactive(GtkSourceCompletion*);
extern GtkSourceCompletionContext* gtk_source_completion_create_context(GtkSourceCompletion*, GtkTextIter*);
extern GtkSourceCompletionInfo* gtk_source_completion_get_info_window(GtkSourceCompletion*);
extern GList* gtk_source_completion_get_providers(GtkSourceCompletion*);
extern GtkSourceView* gtk_source_completion_get_view(GtkSourceCompletion*);
extern void gtk_source_completion_hide(GtkSourceCompletion*);
extern void gtk_source_completion_move_window(GtkSourceCompletion*, GtkTextIter*);
extern int gtk_source_completion_remove_provider(GtkSourceCompletion*, GtkSourceCompletionProvider*, GError**);
extern int gtk_source_completion_show(GtkSourceCompletion*, GList*, GtkSourceCompletionContext*);
extern void gtk_source_completion_unblock_interactive(GtkSourceCompletion*);
extern GType gtk_source_completion_get_type();
extern void gtk_source_completion_context_add_proposals(GtkSourceCompletionContext*, GtkSourceCompletionProvider*, GList*, int);
extern GtkSourceCompletionActivation gtk_source_completion_context_get_activation(GtkSourceCompletionContext*);
extern void gtk_source_completion_context_get_iter(GtkSourceCompletionContext*, GtkTextIter*);
extern GType gtk_source_completion_context_get_type();
extern GtkSourceCompletionInfo* gtk_source_completion_info_new();
extern GtkWidget* gtk_source_completion_info_get_widget(GtkSourceCompletionInfo*);
extern void gtk_source_completion_info_move_to_iter(GtkSourceCompletionInfo*, GtkTextView*, GtkTextIter*);
extern void gtk_source_completion_info_set_widget(GtkSourceCompletionInfo*, GtkWidget*);
extern GType gtk_source_completion_info_get_type();
extern GtkSourceCompletionItem* gtk_source_completion_item_new(char*, char*, GdkPixbuf*, char*);
extern GtkSourceCompletionItem* gtk_source_completion_item_new_from_stock(char*, char*, char*, char*);
extern GtkSourceCompletionItem* gtk_source_completion_item_new_with_markup(char*, char*, GdkPixbuf*, char*);
extern GType gtk_source_completion_item_get_type();
extern void gtk_source_completion_proposal_changed(GtkSourceCompletionProposal*);
extern int gtk_source_completion_proposal_equal(GtkSourceCompletionProposal*, GtkSourceCompletionProposal*);
extern GdkPixbuf* gtk_source_completion_proposal_get_icon(GtkSourceCompletionProposal*);
extern char* gtk_source_completion_proposal_get_info(GtkSourceCompletionProposal*);
extern char* gtk_source_completion_proposal_get_label(GtkSourceCompletionProposal*);
extern char* gtk_source_completion_proposal_get_markup(GtkSourceCompletionProposal*);
extern char* gtk_source_completion_proposal_get_text(GtkSourceCompletionProposal*);
extern uint32_t gtk_source_completion_proposal_hash(GtkSourceCompletionProposal*);
extern GType gtk_source_completion_proposal_get_type();
extern int gtk_source_completion_provider_activate_proposal(GtkSourceCompletionProvider*, GtkSourceCompletionProposal*, GtkTextIter*);
extern GtkSourceCompletionActivation gtk_source_completion_provider_get_activation(GtkSourceCompletionProvider*);
extern GdkPixbuf* gtk_source_completion_provider_get_icon(GtkSourceCompletionProvider*);
extern GtkWidget* gtk_source_completion_provider_get_info_widget(GtkSourceCompletionProvider*, GtkSourceCompletionProposal*);
extern int32_t gtk_source_completion_provider_get_interactive_delay(GtkSourceCompletionProvider*);
extern char* gtk_source_completion_provider_get_name(GtkSourceCompletionProvider*);
extern int32_t gtk_source_completion_provider_get_priority(GtkSourceCompletionProvider*);
extern int gtk_source_completion_provider_get_start_iter(GtkSourceCompletionProvider*, GtkSourceCompletionContext*, GtkSourceCompletionProposal*, GtkTextIter*);
extern int gtk_source_completion_provider_match(GtkSourceCompletionProvider*, GtkSourceCompletionContext*);
extern void gtk_source_completion_provider_populate(GtkSourceCompletionProvider*, GtkSourceCompletionContext*);
extern void gtk_source_completion_provider_update_info(GtkSourceCompletionProvider*, GtkSourceCompletionProposal*, GtkSourceCompletionInfo*);
extern GType gtk_source_completion_provider_get_type();
extern GtkSourceCompletionWords* gtk_source_completion_words_new(char*, GdkPixbuf*);
extern void gtk_source_completion_words_register(GtkSourceCompletionWords*, GtkTextBuffer*);
extern void gtk_source_completion_words_unregister(GtkSourceCompletionWords*, GtkTextBuffer*);
extern GType gtk_source_completion_words_get_type();
extern void gtk_source_gutter_get_padding(GtkSourceGutter*, int32_t*, int32_t*);
extern GtkSourceGutterRenderer* gtk_source_gutter_get_renderer_at_pos(GtkSourceGutter*, int32_t, int32_t);
extern GdkWindow* gtk_source_gutter_get_window(GtkSourceGutter*);
extern int gtk_source_gutter_insert(GtkSourceGutter*, GtkSourceGutterRenderer*, int32_t);
extern void gtk_source_gutter_queue_draw(GtkSourceGutter*);
extern void gtk_source_gutter_remove(GtkSourceGutter*, GtkSourceGutterRenderer*);
extern void gtk_source_gutter_reorder(GtkSourceGutter*, GtkSourceGutterRenderer*, int32_t);
extern void gtk_source_gutter_set_padding(GtkSourceGutter*, int32_t, int32_t);
extern GType gtk_source_gutter_get_type();
extern void gtk_source_gutter_renderer_activate(GtkSourceGutterRenderer*, GtkTextIter*, cairoRectangleInt*, GdkEvent*);
extern void gtk_source_gutter_renderer_begin(GtkSourceGutterRenderer*, cairoContext*, cairoRectangleInt*, cairoRectangleInt*, GtkTextIter*, GtkTextIter*);
extern void gtk_source_gutter_renderer_draw(GtkSourceGutterRenderer*, cairoContext*, cairoRectangleInt*, cairoRectangleInt*, GtkTextIter*, GtkTextIter*, GtkSourceGutterRendererState);
extern void gtk_source_gutter_renderer_end(GtkSourceGutterRenderer*);
extern void gtk_source_gutter_renderer_get_alignment(GtkSourceGutterRenderer*, float*, float*);
extern GtkSourceGutterRendererAlignmentMode gtk_source_gutter_renderer_get_alignment_mode(GtkSourceGutterRenderer*);
extern int gtk_source_gutter_renderer_get_background(GtkSourceGutterRenderer*, GdkRGBA*);
extern void gtk_source_gutter_renderer_get_padding(GtkSourceGutterRenderer*, int32_t*, int32_t*);
extern int32_t gtk_source_gutter_renderer_get_size(GtkSourceGutterRenderer*);
extern GtkTextView* gtk_source_gutter_renderer_get_view(GtkSourceGutterRenderer*);
extern int gtk_source_gutter_renderer_get_visible(GtkSourceGutterRenderer*);
extern GtkTextWindowType gtk_source_gutter_renderer_get_window_type(GtkSourceGutterRenderer*);
extern int gtk_source_gutter_renderer_query_activatable(GtkSourceGutterRenderer*, GtkTextIter*, cairoRectangleInt*, GdkEvent*);
extern void gtk_source_gutter_renderer_query_data(GtkSourceGutterRenderer*, GtkTextIter*, GtkTextIter*, GtkSourceGutterRendererState);
extern int gtk_source_gutter_renderer_query_tooltip(GtkSourceGutterRenderer*, GtkTextIter*, cairoRectangleInt*, int32_t, int32_t, GtkTooltip*);
extern void gtk_source_gutter_renderer_queue_draw(GtkSourceGutterRenderer*);
extern void gtk_source_gutter_renderer_set_alignment(GtkSourceGutterRenderer*, float, float);
extern void gtk_source_gutter_renderer_set_alignment_mode(GtkSourceGutterRenderer*, GtkSourceGutterRendererAlignmentMode);
extern void gtk_source_gutter_renderer_set_background(GtkSourceGutterRenderer*, GdkRGBA*);
extern void gtk_source_gutter_renderer_set_padding(GtkSourceGutterRenderer*, int32_t, int32_t);
extern void gtk_source_gutter_renderer_set_size(GtkSourceGutterRenderer*, int32_t);
extern void gtk_source_gutter_renderer_set_visible(GtkSourceGutterRenderer*, int);
extern GType gtk_source_gutter_renderer_get_type();
extern GtkSourceGutterRenderer* gtk_source_gutter_renderer_pixbuf_new();
extern GIcon* gtk_source_gutter_renderer_pixbuf_get_gicon(GtkSourceGutterRendererPixbuf*);
extern char* gtk_source_gutter_renderer_pixbuf_get_icon_name(GtkSourceGutterRendererPixbuf*);
extern GdkPixbuf* gtk_source_gutter_renderer_pixbuf_get_pixbuf(GtkSourceGutterRendererPixbuf*);
extern char* gtk_source_gutter_renderer_pixbuf_get_stock_id(GtkSourceGutterRendererPixbuf*);
extern void gtk_source_gutter_renderer_pixbuf_set_gicon(GtkSourceGutterRendererPixbuf*, GIcon*);
extern void gtk_source_gutter_renderer_pixbuf_set_icon_name(GtkSourceGutterRendererPixbuf*, char*);
extern void gtk_source_gutter_renderer_pixbuf_set_pixbuf(GtkSourceGutterRendererPixbuf*, GdkPixbuf*);
extern void gtk_source_gutter_renderer_pixbuf_set_stock_id(GtkSourceGutterRendererPixbuf*, char*);
extern GType gtk_source_gutter_renderer_pixbuf_get_type();
extern GtkSourceGutterRenderer* gtk_source_gutter_renderer_text_new();
extern void gtk_source_gutter_renderer_text_measure(GtkSourceGutterRendererText*, char*, int32_t*, int32_t*);
extern void gtk_source_gutter_renderer_text_measure_markup(GtkSourceGutterRendererText*, char*, int32_t*, int32_t*);
extern void gtk_source_gutter_renderer_text_set_markup(GtkSourceGutterRendererText*, char*, int32_t);
extern void gtk_source_gutter_renderer_text_set_text(GtkSourceGutterRendererText*, char*, int32_t);
extern GType gtk_source_gutter_renderer_text_get_type();
extern char** gtk_source_language_get_globs(GtkSourceLanguage*);
extern int gtk_source_language_get_hidden(GtkSourceLanguage*);
extern char* gtk_source_language_get_id(GtkSourceLanguage*);
extern char* gtk_source_language_get_metadata(GtkSourceLanguage*, char*);
extern char** gtk_source_language_get_mime_types(GtkSourceLanguage*);
extern char* gtk_source_language_get_name(GtkSourceLanguage*);
extern char* gtk_source_language_get_section(GtkSourceLanguage*);
extern char** gtk_source_language_get_style_ids(GtkSourceLanguage*);
extern char* gtk_source_language_get_style_name(GtkSourceLanguage*, char*);
extern GType gtk_source_language_get_type();
extern GtkSourceLanguageManager* gtk_source_language_manager_new();
extern GtkSourceLanguageManager* gtk_source_language_manager_get_default();
extern GtkSourceLanguage* gtk_source_language_manager_get_language(GtkSourceLanguageManager*, char*);
extern char** gtk_source_language_manager_get_language_ids(GtkSourceLanguageManager*);
extern char** gtk_source_language_manager_get_search_path(GtkSourceLanguageManager*);
extern GtkSourceLanguage* gtk_source_language_manager_guess_language(GtkSourceLanguageManager*, char*, char*);
extern void gtk_source_language_manager_set_search_path(GtkSourceLanguageManager*, char**);
extern GType gtk_source_language_manager_get_type();
extern GtkSourceMark* gtk_source_mark_new(char*, char*);
extern char* gtk_source_mark_get_category(GtkSourceMark*);
extern GtkSourceMark* gtk_source_mark_next(GtkSourceMark*, char*);
extern GtkSourceMark* gtk_source_mark_prev(GtkSourceMark*, char*);
extern GType gtk_source_mark_get_type();
extern GtkSourceMarkAttributes* gtk_source_mark_attributes_new();
extern int gtk_source_mark_attributes_get_background(GtkSourceMarkAttributes*, GdkRGBA*);
extern GIcon* gtk_source_mark_attributes_get_gicon(GtkSourceMarkAttributes*);
extern char* gtk_source_mark_attributes_get_icon_name(GtkSourceMarkAttributes*);
extern GdkPixbuf* gtk_source_mark_attributes_get_pixbuf(GtkSourceMarkAttributes*);
extern char* gtk_source_mark_attributes_get_stock_id(GtkSourceMarkAttributes*);
extern char* gtk_source_mark_attributes_get_tooltip_markup(GtkSourceMarkAttributes*, GtkSourceMark*);
extern char* gtk_source_mark_attributes_get_tooltip_text(GtkSourceMarkAttributes*, GtkSourceMark*);
extern GdkPixbuf* gtk_source_mark_attributes_render_icon(GtkSourceMarkAttributes*, GtkWidget*, int32_t);
extern void gtk_source_mark_attributes_set_background(GtkSourceMarkAttributes*, GdkRGBA*);
extern void gtk_source_mark_attributes_set_gicon(GtkSourceMarkAttributes*, GIcon*);
extern void gtk_source_mark_attributes_set_icon_name(GtkSourceMarkAttributes*, char*);
extern void gtk_source_mark_attributes_set_pixbuf(GtkSourceMarkAttributes*, GdkPixbuf*);
extern void gtk_source_mark_attributes_set_stock_id(GtkSourceMarkAttributes*, char*);
extern GType gtk_source_mark_attributes_get_type();
extern GtkSourcePrintCompositor* gtk_source_print_compositor_new(GtkSourceBuffer*);
extern GtkSourcePrintCompositor* gtk_source_print_compositor_new_from_view(GtkSourceView*);
extern void gtk_source_print_compositor_draw_page(GtkSourcePrintCompositor*, GtkPrintContext*, int32_t);
extern char* gtk_source_print_compositor_get_body_font_name(GtkSourcePrintCompositor*);
extern double gtk_source_print_compositor_get_bottom_margin(GtkSourcePrintCompositor*, GtkUnit);
extern GtkSourceBuffer* gtk_source_print_compositor_get_buffer(GtkSourcePrintCompositor*);
extern char* gtk_source_print_compositor_get_footer_font_name(GtkSourcePrintCompositor*);
extern char* gtk_source_print_compositor_get_header_font_name(GtkSourcePrintCompositor*);
extern int gtk_source_print_compositor_get_highlight_syntax(GtkSourcePrintCompositor*);
extern double gtk_source_print_compositor_get_left_margin(GtkSourcePrintCompositor*, GtkUnit);
extern char* gtk_source_print_compositor_get_line_numbers_font_name(GtkSourcePrintCompositor*);
extern int32_t gtk_source_print_compositor_get_n_pages(GtkSourcePrintCompositor*);
extern double gtk_source_print_compositor_get_pagination_progress(GtkSourcePrintCompositor*);
extern int gtk_source_print_compositor_get_print_footer(GtkSourcePrintCompositor*);
extern int gtk_source_print_compositor_get_print_header(GtkSourcePrintCompositor*);
extern uint32_t gtk_source_print_compositor_get_print_line_numbers(GtkSourcePrintCompositor*);
extern double gtk_source_print_compositor_get_right_margin(GtkSourcePrintCompositor*, GtkUnit);
extern uint32_t gtk_source_print_compositor_get_tab_width(GtkSourcePrintCompositor*);
extern double gtk_source_print_compositor_get_top_margin(GtkSourcePrintCompositor*, GtkUnit);
extern GtkWrapMode gtk_source_print_compositor_get_wrap_mode(GtkSourcePrintCompositor*);
extern int gtk_source_print_compositor_paginate(GtkSourcePrintCompositor*, GtkPrintContext*);
extern void gtk_source_print_compositor_set_body_font_name(GtkSourcePrintCompositor*, char*);
extern void gtk_source_print_compositor_set_bottom_margin(GtkSourcePrintCompositor*, double, GtkUnit);
extern void gtk_source_print_compositor_set_footer_font_name(GtkSourcePrintCompositor*, char*);
extern void gtk_source_print_compositor_set_footer_format(GtkSourcePrintCompositor*, int, char*, char*, char*);
extern void gtk_source_print_compositor_set_header_font_name(GtkSourcePrintCompositor*, char*);
extern void gtk_source_print_compositor_set_header_format(GtkSourcePrintCompositor*, int, char*, char*, char*);
extern void gtk_source_print_compositor_set_highlight_syntax(GtkSourcePrintCompositor*, int);
extern void gtk_source_print_compositor_set_left_margin(GtkSourcePrintCompositor*, double, GtkUnit);
extern void gtk_source_print_compositor_set_line_numbers_font_name(GtkSourcePrintCompositor*, char*);
extern void gtk_source_print_compositor_set_print_footer(GtkSourcePrintCompositor*, int);
extern void gtk_source_print_compositor_set_print_header(GtkSourcePrintCompositor*, int);
extern void gtk_source_print_compositor_set_print_line_numbers(GtkSourcePrintCompositor*, uint32_t);
extern void gtk_source_print_compositor_set_right_margin(GtkSourcePrintCompositor*, double, GtkUnit);
extern void gtk_source_print_compositor_set_tab_width(GtkSourcePrintCompositor*, uint32_t);
extern void gtk_source_print_compositor_set_top_margin(GtkSourcePrintCompositor*, double, GtkUnit);
extern void gtk_source_print_compositor_set_wrap_mode(GtkSourcePrintCompositor*, GtkWrapMode);
extern GType gtk_source_print_compositor_get_type();
extern GtkSourceStyle* gtk_source_style_copy(GtkSourceStyle*);
extern GType gtk_source_style_get_type();
extern char** gtk_source_style_scheme_get_authors(GtkSourceStyleScheme*);
extern char* gtk_source_style_scheme_get_description(GtkSourceStyleScheme*);
extern char* gtk_source_style_scheme_get_filename(GtkSourceStyleScheme*);
extern char* gtk_source_style_scheme_get_id(GtkSourceStyleScheme*);
extern char* gtk_source_style_scheme_get_name(GtkSourceStyleScheme*);
extern GtkSourceStyle* gtk_source_style_scheme_get_style(GtkSourceStyleScheme*, char*);
extern GType gtk_source_style_scheme_get_type();
extern GtkSourceStyleSchemeManager* gtk_source_style_scheme_manager_new();
extern GtkSourceStyleSchemeManager* gtk_source_style_scheme_manager_get_default();
extern void gtk_source_style_scheme_manager_append_search_path(GtkSourceStyleSchemeManager*, char*);
extern void gtk_source_style_scheme_manager_force_rescan(GtkSourceStyleSchemeManager*);
extern GtkSourceStyleScheme* gtk_source_style_scheme_manager_get_scheme(GtkSourceStyleSchemeManager*, char*);
extern char** gtk_source_style_scheme_manager_get_scheme_ids(GtkSourceStyleSchemeManager*);
extern char** gtk_source_style_scheme_manager_get_search_path(GtkSourceStyleSchemeManager*);
extern void gtk_source_style_scheme_manager_prepend_search_path(GtkSourceStyleSchemeManager*, char*);
extern void gtk_source_style_scheme_manager_set_search_path(GtkSourceStyleSchemeManager*, char**);
extern GType gtk_source_style_scheme_manager_get_type();
extern void gtk_source_undo_manager_begin_not_undoable_action(GtkSourceUndoManager*);
extern int gtk_source_undo_manager_can_redo(GtkSourceUndoManager*);
extern void gtk_source_undo_manager_can_redo_changed(GtkSourceUndoManager*);
extern int gtk_source_undo_manager_can_undo(GtkSourceUndoManager*);
extern void gtk_source_undo_manager_can_undo_changed(GtkSourceUndoManager*);
extern void gtk_source_undo_manager_end_not_undoable_action(GtkSourceUndoManager*);
extern void gtk_source_undo_manager_redo(GtkSourceUndoManager*);
extern void gtk_source_undo_manager_undo(GtkSourceUndoManager*);
extern GType gtk_source_undo_manager_get_type();
extern GtkWidget* gtk_source_view_new();
extern GtkWidget* gtk_source_view_new_with_buffer(GtkSourceBuffer*);
extern int gtk_source_view_get_auto_indent(GtkSourceView*);
extern GtkSourceCompletion* gtk_source_view_get_completion(GtkSourceView*);
extern GtkSourceDrawSpacesFlags gtk_source_view_get_draw_spaces(GtkSourceView*);
extern GtkSourceGutter* gtk_source_view_get_gutter(GtkSourceView*, GtkTextWindowType);
extern int gtk_source_view_get_highlight_current_line(GtkSourceView*);
extern int gtk_source_view_get_indent_on_tab(GtkSourceView*);
extern int32_t gtk_source_view_get_indent_width(GtkSourceView*);
extern int gtk_source_view_get_insert_spaces_instead_of_tabs(GtkSourceView*);
extern GtkSourceMarkAttributes* gtk_source_view_get_mark_attributes(GtkSourceView*, char*, int32_t*);
extern uint32_t gtk_source_view_get_right_margin_position(GtkSourceView*);
extern int gtk_source_view_get_show_line_marks(GtkSourceView*);
extern int gtk_source_view_get_show_line_numbers(GtkSourceView*);
extern int gtk_source_view_get_show_right_margin(GtkSourceView*);
extern GtkSourceSmartHomeEndType gtk_source_view_get_smart_home_end(GtkSourceView*);
extern uint32_t gtk_source_view_get_tab_width(GtkSourceView*);
extern uint32_t gtk_source_view_get_visual_column(GtkSourceView*, GtkTextIter*);
extern void gtk_source_view_set_auto_indent(GtkSourceView*, int);
extern void gtk_source_view_set_draw_spaces(GtkSourceView*, GtkSourceDrawSpacesFlags);
extern void gtk_source_view_set_highlight_current_line(GtkSourceView*, int);
extern void gtk_source_view_set_indent_on_tab(GtkSourceView*, int);
extern void gtk_source_view_set_indent_width(GtkSourceView*, int32_t);
extern void gtk_source_view_set_insert_spaces_instead_of_tabs(GtkSourceView*, int);
extern void gtk_source_view_set_mark_attributes(GtkSourceView*, char*, GtkSourceMarkAttributes*, int32_t);
extern void gtk_source_view_set_right_margin_position(GtkSourceView*, uint32_t);
extern void gtk_source_view_set_show_line_marks(GtkSourceView*, int);
extern void gtk_source_view_set_show_line_numbers(GtkSourceView*, int);
extern void gtk_source_view_set_show_right_margin(GtkSourceView*, int);
extern void gtk_source_view_set_smart_home_end(GtkSourceView*, GtkSourceSmartHomeEndType);
extern void gtk_source_view_set_tab_width(GtkSourceView*, uint32_t);
extern GType gtk_source_view_get_type();
extern uint32_t gtk_source_completion_error_quark();
struct _GtkSourceBufferClass { uint8_t _data[320]; };
struct _GtkSourceBufferPrivate {};
struct _GtkSourceCompletionClass { uint8_t _data[192]; };
struct _GtkSourceCompletionContextClass { uint8_t _data[168]; };
struct _GtkSourceCompletionContextPrivate {};
struct _GtkSourceCompletionInfoClass { uint8_t _data[1080]; };
struct _GtkSourceCompletionInfoPrivate {};
struct _GtkSourceCompletionItemClass { uint8_t _data[136]; };
struct _GtkSourceCompletionItemPrivate {};
struct _GtkSourceCompletionPrivate {};
struct _GtkSourceCompletionProposalIface { uint8_t _data[80]; };
struct _GtkSourceCompletionProviderIface { uint8_t _data[104]; };
struct _GtkSourceCompletionWordsClass { uint8_t _data[136]; };
struct _GtkSourceCompletionWordsPrivate {};
struct _GtkSourceGutterClass { uint8_t _data[136]; };
struct _GtkSourceGutterPrivate {};
struct _GtkSourceGutterRendererClass { uint8_t _data[216]; };
struct _GtkSourceGutterRendererPixbufClass { uint8_t _data[216]; };
struct _GtkSourceGutterRendererPixbufPrivate {};
struct _GtkSourceGutterRendererPrivate {};
struct _GtkSourceGutterRendererTextClass { uint8_t _data[216]; };
struct _GtkSourceGutterRendererTextPrivate {};
struct _GtkSourceLanguageClass { uint8_t _data[152]; };
struct _GtkSourceLanguageManagerClass { uint8_t _data[168]; };
struct _GtkSourceLanguageManagerPrivate {};
struct _GtkSourceLanguagePrivate {};
struct _GtkSourceMarkAttributesClass { uint8_t _data[136]; };
struct _GtkSourceMarkAttributesPrivate {};
struct _GtkSourceMarkClass { uint8_t _data[184]; };
struct _GtkSourceMarkPrivate {};
struct _GtkSourcePrintCompositorClass { uint8_t _data[152]; };
struct _GtkSourcePrintCompositorPrivate {};
struct _GtkSourceStyleSchemeClass { uint8_t _data[152]; };
struct _GtkSourceStyleSchemeManagerClass { uint8_t _data[168]; };
struct _GtkSourceStyleSchemeManagerPrivate {};
struct _GtkSourceStyleSchemePrivate {};
struct _GtkSourceUndoManagerIface { uint8_t _data[80]; };
struct _GtkSourceViewClass { uint8_t _data[1168]; };
struct _GtkSourceViewPrivate {};


extern GObject *g_object_ref_sink(GObject*);
extern void g_object_unref(GObject*);
extern void g_error_free(GError*);
extern void g_free(void*);

#cgo pkg-config: gtksourceview-3.0
*/
import "C"
import "unsafe"
import "errors"

// package dependencies
import (
	"github.com/bytbox/gogobject/gobject-2.0"
	"github.com/bytbox/gogobject/gdkpixbuf-2.0"
	"github.com/bytbox/gogobject/gdk-3.0"
	"github.com/bytbox/gogobject/gtk-3.0"
	"github.com/bytbox/gogobject/cairo-1.0"
	"github.com/bytbox/gogobject/gio-2.0"
	"github.com/bytbox/gogobject/atk-1.0"
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


//export _GtkSource_go_callback_cleanup
func _GtkSource_go_callback_cleanup(gofunc unsafe.Pointer) {
	gobject.Holder.Release(gofunc)
}


type BracketMatchType C.uint32_t
const (
	BracketMatchTypeNone BracketMatchType = 0
	BracketMatchTypeOutOfRange BracketMatchType = 1
	BracketMatchTypeNotFound BracketMatchType = 2
	BracketMatchTypeFound BracketMatchType = 3
)
type BufferLike interface {
	gtk.TextBufferLike
	InheritedFromGtkSourceBuffer() *C.GtkSourceBuffer
}

type Buffer struct {
	gtk.TextBuffer
	
}

func ToBuffer(objlike gobject.ObjectLike) *Buffer {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Buffer)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Buffer)(obj)
	}
	panic("cannot cast to Buffer")
}

func (this0 *Buffer) InheritedFromGtkSourceBuffer() *C.GtkSourceBuffer {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceBuffer)(this0.C)
}

func (this0 *Buffer) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_buffer_get_type())
}

func BufferGetType() gobject.Type {
	return (*Buffer)(nil).GetStaticType()
}
func NewBuffer(table0 gtk.TextTagTableLike) *Buffer {
	var table1 *C.GtkTextTagTable
	if table0 != nil {
		table1 = table0.InheritedFromGtkTextTagTable()
	}
	ret1 := C.gtk_source_buffer_new(table1)
	var ret2 *Buffer
	ret2 = (*Buffer)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewBufferWithLanguage(language0 LanguageLike) *Buffer {
	var language1 *C.GtkSourceLanguage
	if language0 != nil {
		language1 = language0.InheritedFromGtkSourceLanguage()
	}
	ret1 := C.gtk_source_buffer_new_with_language(language1)
	var ret2 *Buffer
	ret2 = (*Buffer)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Buffer) BackwardIterToSourceMark(iter0 *gtk.TextIter, category0 string) bool {
	var this1 *C.GtkSourceBuffer
	var iter1 *C.GtkTextIter
	var category1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	ret1 := C.gtk_source_buffer_backward_iter_to_source_mark(this1, iter1, category1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Buffer) BeginNotUndoableAction() {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	C.gtk_source_buffer_begin_not_undoable_action(this1)
}
func (this0 *Buffer) CanRedo() bool {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	ret1 := C.gtk_source_buffer_can_redo(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Buffer) CanUndo() bool {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	ret1 := C.gtk_source_buffer_can_undo(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Buffer) CreateSourceMark(name0 string, category0 string, where0 *gtk.TextIter) *Mark {
	var this1 *C.GtkSourceBuffer
	var name1 *C.char
	var category1 *C.char
	var where1 *C.GtkTextIter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	where1 = (*C.GtkTextIter)(unsafe.Pointer(where0))
	ret1 := C.gtk_source_buffer_create_source_mark(this1, name1, category1, where1)
	var ret2 *Mark
	ret2 = (*Mark)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Buffer) EndNotUndoableAction() {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	C.gtk_source_buffer_end_not_undoable_action(this1)
}
func (this0 *Buffer) EnsureHighlight(start0 *gtk.TextIter, end0 *gtk.TextIter) {
	var this1 *C.GtkSourceBuffer
	var start1 *C.GtkTextIter
	var end1 *C.GtkTextIter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	start1 = (*C.GtkTextIter)(unsafe.Pointer(start0))
	end1 = (*C.GtkTextIter)(unsafe.Pointer(end0))
	C.gtk_source_buffer_ensure_highlight(this1, start1, end1)
}
func (this0 *Buffer) ForwardIterToSourceMark(iter0 *gtk.TextIter, category0 string) bool {
	var this1 *C.GtkSourceBuffer
	var iter1 *C.GtkTextIter
	var category1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	ret1 := C.gtk_source_buffer_forward_iter_to_source_mark(this1, iter1, category1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Buffer) GetContextClassesAtIter(iter0 *gtk.TextIter) []string {
	var this1 *C.GtkSourceBuffer
	var iter1 *C.GtkTextIter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	ret1 := C.gtk_source_buffer_get_context_classes_at_iter(this1, iter1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *Buffer) GetHighlightMatchingBrackets() bool {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	ret1 := C.gtk_source_buffer_get_highlight_matching_brackets(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Buffer) GetHighlightSyntax() bool {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	ret1 := C.gtk_source_buffer_get_highlight_syntax(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Buffer) GetLanguage() *Language {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	ret1 := C.gtk_source_buffer_get_language(this1)
	var ret2 *Language
	ret2 = (*Language)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Buffer) GetMaxUndoLevels() int {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	ret1 := C.gtk_source_buffer_get_max_undo_levels(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *Buffer) GetSourceMarksAtIter(iter0 *gtk.TextIter, category0 string) []*Mark {
	var this1 *C.GtkSourceBuffer
	var iter1 *C.GtkTextIter
	var category1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	ret1 := C.gtk_source_buffer_get_source_marks_at_iter(this1, iter1, category1)
	var ret2 []*Mark
	for iter := (*_GSList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Mark
		elt = (*Mark)(gobject.ObjectWrap(unsafe.Pointer((*C.GtkSourceMark)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Buffer) GetSourceMarksAtLine(line0 int, category0 string) []*Mark {
	var this1 *C.GtkSourceBuffer
	var line1 C.int32_t
	var category1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	line1 = C.int32_t(line0)
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	ret1 := C.gtk_source_buffer_get_source_marks_at_line(this1, line1, category1)
	var ret2 []*Mark
	for iter := (*_GSList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *Mark
		elt = (*Mark)(gobject.ObjectWrap(unsafe.Pointer((*C.GtkSourceMark)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Buffer) GetStyleScheme() *StyleScheme {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	ret1 := C.gtk_source_buffer_get_style_scheme(this1)
	var ret2 *StyleScheme
	ret2 = (*StyleScheme)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Buffer) GetUndoManager() *UndoManager {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	ret1 := C.gtk_source_buffer_get_undo_manager(this1)
	var ret2 *UndoManager
	ret2 = (*UndoManager)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Buffer) IterBackwardToContextClassToggle(iter0 *gtk.TextIter, context_class0 string) bool {
	var this1 *C.GtkSourceBuffer
	var iter1 *C.GtkTextIter
	var context_class1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	context_class1 = _GoStringToGString(context_class0)
	defer C.free(unsafe.Pointer(context_class1))
	ret1 := C.gtk_source_buffer_iter_backward_to_context_class_toggle(this1, iter1, context_class1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Buffer) IterForwardToContextClassToggle(iter0 *gtk.TextIter, context_class0 string) bool {
	var this1 *C.GtkSourceBuffer
	var iter1 *C.GtkTextIter
	var context_class1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	context_class1 = _GoStringToGString(context_class0)
	defer C.free(unsafe.Pointer(context_class1))
	ret1 := C.gtk_source_buffer_iter_forward_to_context_class_toggle(this1, iter1, context_class1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Buffer) IterHasContextClass(iter0 *gtk.TextIter, context_class0 string) bool {
	var this1 *C.GtkSourceBuffer
	var iter1 *C.GtkTextIter
	var context_class1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	context_class1 = _GoStringToGString(context_class0)
	defer C.free(unsafe.Pointer(context_class1))
	ret1 := C.gtk_source_buffer_iter_has_context_class(this1, iter1, context_class1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Buffer) Redo() {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	C.gtk_source_buffer_redo(this1)
}
func (this0 *Buffer) RemoveSourceMarks(start0 *gtk.TextIter, end0 *gtk.TextIter, category0 string) {
	var this1 *C.GtkSourceBuffer
	var start1 *C.GtkTextIter
	var end1 *C.GtkTextIter
	var category1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	start1 = (*C.GtkTextIter)(unsafe.Pointer(start0))
	end1 = (*C.GtkTextIter)(unsafe.Pointer(end0))
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	C.gtk_source_buffer_remove_source_marks(this1, start1, end1, category1)
}
func (this0 *Buffer) SetHighlightMatchingBrackets(highlight0 bool) {
	var this1 *C.GtkSourceBuffer
	var highlight1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	highlight1 = _GoBoolToCBool(highlight0)
	C.gtk_source_buffer_set_highlight_matching_brackets(this1, highlight1)
}
func (this0 *Buffer) SetHighlightSyntax(highlight0 bool) {
	var this1 *C.GtkSourceBuffer
	var highlight1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	highlight1 = _GoBoolToCBool(highlight0)
	C.gtk_source_buffer_set_highlight_syntax(this1, highlight1)
}
func (this0 *Buffer) SetLanguage(language0 LanguageLike) {
	var this1 *C.GtkSourceBuffer
	var language1 *C.GtkSourceLanguage
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	if language0 != nil {
		language1 = language0.InheritedFromGtkSourceLanguage()
	}
	C.gtk_source_buffer_set_language(this1, language1)
}
func (this0 *Buffer) SetMaxUndoLevels(max_undo_levels0 int) {
	var this1 *C.GtkSourceBuffer
	var max_undo_levels1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	max_undo_levels1 = C.int32_t(max_undo_levels0)
	C.gtk_source_buffer_set_max_undo_levels(this1, max_undo_levels1)
}
func (this0 *Buffer) SetStyleScheme(scheme0 StyleSchemeLike) {
	var this1 *C.GtkSourceBuffer
	var scheme1 *C.GtkSourceStyleScheme
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	if scheme0 != nil {
		scheme1 = scheme0.InheritedFromGtkSourceStyleScheme()
	}
	C.gtk_source_buffer_set_style_scheme(this1, scheme1)
}
func (this0 *Buffer) SetUndoManager(manager0 UndoManagerLike) {
	var this1 *C.GtkSourceBuffer
	var manager1 *C.GtkSourceUndoManager
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	if manager0 != nil {
		manager1 = manager0.ImplementsGtkSourceUndoManager()}
	C.gtk_source_buffer_set_undo_manager(this1, manager1)
}
func (this0 *Buffer) Undo() {
	var this1 *C.GtkSourceBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceBuffer()
	}
	C.gtk_source_buffer_undo(this1)
}
const CompletionCapabilityAutomatic = "standard::automatic"
const CompletionCapabilityInteractive = "standard::interactive"
type CompletionLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourceCompletion() *C.GtkSourceCompletion
}

type Completion struct {
	gobject.Object
	
}

func ToCompletion(objlike gobject.ObjectLike) *Completion {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Completion)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Completion)(obj)
	}
	panic("cannot cast to Completion")
}

func (this0 *Completion) InheritedFromGtkSourceCompletion() *C.GtkSourceCompletion {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceCompletion)(this0.C)
}

func (this0 *Completion) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_completion_get_type())
}

func CompletionGetType() gobject.Type {
	return (*Completion)(nil).GetStaticType()
}
func (this0 *Completion) AddProvider(provider0 CompletionProviderLike) (bool, error) {
	var this1 *C.GtkSourceCompletion
	var provider1 *C.GtkSourceCompletionProvider
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	if provider0 != nil {
		provider1 = provider0.ImplementsGtkSourceCompletionProvider()}
	ret1 := C.gtk_source_completion_add_provider(this1, provider1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Completion) BlockInteractive() {
	var this1 *C.GtkSourceCompletion
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	C.gtk_source_completion_block_interactive(this1)
}
func (this0 *Completion) CreateContext(position0 *gtk.TextIter) *CompletionContext {
	var this1 *C.GtkSourceCompletion
	var position1 *C.GtkTextIter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	position1 = (*C.GtkTextIter)(unsafe.Pointer(position0))
	ret1 := C.gtk_source_completion_create_context(this1, position1)
	var ret2 *CompletionContext
	ret2 = (*CompletionContext)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Completion) GetInfoWindow() *CompletionInfo {
	var this1 *C.GtkSourceCompletion
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	ret1 := C.gtk_source_completion_get_info_window(this1)
	var ret2 *CompletionInfo
	ret2 = (*CompletionInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Completion) GetProviders() []*CompletionProvider {
	var this1 *C.GtkSourceCompletion
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	ret1 := C.gtk_source_completion_get_providers(this1)
	var ret2 []*CompletionProvider
	for iter := (*_GList)(unsafe.Pointer(ret1)); iter != nil; iter = iter.next {
		var elt *CompletionProvider
		elt = (*CompletionProvider)(gobject.ObjectWrap(unsafe.Pointer((*C.GtkSourceCompletionProvider)(iter.data)), true))
		ret2 = append(ret2, elt)
	}
	return ret2
}
func (this0 *Completion) GetView() *View {
	var this1 *C.GtkSourceCompletion
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	ret1 := C.gtk_source_completion_get_view(this1)
	var ret2 *View
	ret2 = (*View)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Completion) Hide() {
	var this1 *C.GtkSourceCompletion
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	C.gtk_source_completion_hide(this1)
}
func (this0 *Completion) MoveWindow(iter0 *gtk.TextIter) {
	var this1 *C.GtkSourceCompletion
	var iter1 *C.GtkTextIter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	C.gtk_source_completion_move_window(this1, iter1)
}
func (this0 *Completion) RemoveProvider(provider0 CompletionProviderLike) (bool, error) {
	var this1 *C.GtkSourceCompletion
	var provider1 *C.GtkSourceCompletionProvider
	var err1 *C.GError
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	if provider0 != nil {
		provider1 = provider0.ImplementsGtkSourceCompletionProvider()}
	ret1 := C.gtk_source_completion_remove_provider(this1, provider1, &err1)
	var ret2 bool
	var err2 error
	ret2 = ret1 != 0
	if err1 != nil {
		err2 = errors.New(C.GoString(((*_GError)(unsafe.Pointer(err1))).message))
		C.g_error_free(err1)
	}
	return ret2, err2
}
func (this0 *Completion) Show(providers0 []*CompletionProvider, context0 CompletionContextLike) bool {
	var this1 *C.GtkSourceCompletion
	var providers1 *C.GList
	var context1 *C.GtkSourceCompletionContext
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	if context0 != nil {
		context1 = context0.InheritedFromGtkSourceCompletionContext()
	}
	ret1 := C.gtk_source_completion_show(this1, providers1, context1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Completion) UnblockInteractive() {
	var this1 *C.GtkSourceCompletion
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletion()
	}
	C.gtk_source_completion_unblock_interactive(this1)
}
type CompletionActivation C.uint32_t
const (
	CompletionActivationNone CompletionActivation = 0
	CompletionActivationInteractive CompletionActivation = 1
	CompletionActivationUserRequested CompletionActivation = 2
)
type CompletionContextLike interface {
	gobject.InitiallyUnownedLike
	InheritedFromGtkSourceCompletionContext() *C.GtkSourceCompletionContext
}

type CompletionContext struct {
	gobject.InitiallyUnowned
	
}

func ToCompletionContext(objlike gobject.ObjectLike) *CompletionContext {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*CompletionContext)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*CompletionContext)(obj)
	}
	panic("cannot cast to CompletionContext")
}

func (this0 *CompletionContext) InheritedFromGtkSourceCompletionContext() *C.GtkSourceCompletionContext {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceCompletionContext)(this0.C)
}

func (this0 *CompletionContext) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_completion_context_get_type())
}

func CompletionContextGetType() gobject.Type {
	return (*CompletionContext)(nil).GetStaticType()
}
func (this0 *CompletionContext) AddProposals(provider0 CompletionProviderLike, proposals0 []*CompletionProposal, finished0 bool) {
	var this1 *C.GtkSourceCompletionContext
	var provider1 *C.GtkSourceCompletionProvider
	var proposals1 *C.GList
	var finished1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletionContext()
	}
	if provider0 != nil {
		provider1 = provider0.ImplementsGtkSourceCompletionProvider()}
	finished1 = _GoBoolToCBool(finished0)
	C.gtk_source_completion_context_add_proposals(this1, provider1, proposals1, finished1)
}
func (this0 *CompletionContext) GetActivation() CompletionActivation {
	var this1 *C.GtkSourceCompletionContext
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletionContext()
	}
	ret1 := C.gtk_source_completion_context_get_activation(this1)
	var ret2 CompletionActivation
	ret2 = CompletionActivation(ret1)
	return ret2
}
func (this0 *CompletionContext) GetIter() gtk.TextIter {
	var this1 *C.GtkSourceCompletionContext
	var iter1 C.GtkTextIter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletionContext()
	}
	C.gtk_source_completion_context_get_iter(this1, &iter1)
	var iter2 gtk.TextIter
	iter2 = *(*gtk.TextIter)(unsafe.Pointer(&iter1))
	return iter2
}
type CompletionError C.uint32_t
const (
	CompletionErrorAlreadyBound CompletionError = 0
	CompletionErrorNotBound CompletionError = 1
)
type CompletionInfoLike interface {
	gtk.WindowLike
	InheritedFromGtkSourceCompletionInfo() *C.GtkSourceCompletionInfo
}

type CompletionInfo struct {
	gtk.Window
	atk.ImplementorIfaceImpl
	gtk.BuildableImpl
}

func ToCompletionInfo(objlike gobject.ObjectLike) *CompletionInfo {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*CompletionInfo)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*CompletionInfo)(obj)
	}
	panic("cannot cast to CompletionInfo")
}

func (this0 *CompletionInfo) InheritedFromGtkSourceCompletionInfo() *C.GtkSourceCompletionInfo {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceCompletionInfo)(this0.C)
}

func (this0 *CompletionInfo) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_completion_info_get_type())
}

func CompletionInfoGetType() gobject.Type {
	return (*CompletionInfo)(nil).GetStaticType()
}
func NewCompletionInfo() *CompletionInfo {
	ret1 := C.gtk_source_completion_info_new()
	var ret2 *CompletionInfo
	ret2 = (*CompletionInfo)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *CompletionInfo) GetWidget() *gtk.Widget {
	var this1 *C.GtkSourceCompletionInfo
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletionInfo()
	}
	ret1 := C.gtk_source_completion_info_get_widget(this1)
	var ret2 *gtk.Widget
	ret2 = (*gtk.Widget)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *CompletionInfo) MoveToIter(view0 gtk.TextViewLike, iter0 *gtk.TextIter) {
	var this1 *C.GtkSourceCompletionInfo
	var view1 *C.GtkTextView
	var iter1 *C.GtkTextIter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletionInfo()
	}
	if view0 != nil {
		view1 = view0.InheritedFromGtkTextView()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	C.gtk_source_completion_info_move_to_iter(this1, view1, iter1)
}
func (this0 *CompletionInfo) SetWidget(widget0 gtk.WidgetLike) {
	var this1 *C.GtkSourceCompletionInfo
	var widget1 *C.GtkWidget
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletionInfo()
	}
	if widget0 != nil {
		widget1 = widget0.InheritedFromGtkWidget()
	}
	C.gtk_source_completion_info_set_widget(this1, widget1)
}
type CompletionItemLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourceCompletionItem() *C.GtkSourceCompletionItem
}

type CompletionItem struct {
	gobject.Object
	CompletionProposalImpl
}

func ToCompletionItem(objlike gobject.ObjectLike) *CompletionItem {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*CompletionItem)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*CompletionItem)(obj)
	}
	panic("cannot cast to CompletionItem")
}

func (this0 *CompletionItem) InheritedFromGtkSourceCompletionItem() *C.GtkSourceCompletionItem {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceCompletionItem)(this0.C)
}

func (this0 *CompletionItem) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_completion_item_get_type())
}

func CompletionItemGetType() gobject.Type {
	return (*CompletionItem)(nil).GetStaticType()
}
func NewCompletionItem(label0 string, text0 string, icon0 gdkpixbuf.PixbufLike, info0 string) *CompletionItem {
	var label1 *C.char
	var text1 *C.char
	var icon1 *C.GdkPixbuf
	var info1 *C.char
	label1 = _GoStringToGString(label0)
	defer C.free(unsafe.Pointer(label1))
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	if icon0 != nil {
		icon1 = icon0.InheritedFromGdkPixbuf()
	}
	info1 = _GoStringToGString(info0)
	defer C.free(unsafe.Pointer(info1))
	ret1 := C.gtk_source_completion_item_new(label1, text1, icon1, info1)
	var ret2 *CompletionItem
	ret2 = (*CompletionItem)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewCompletionItemFromStock(label0 string, text0 string, stock0 string, info0 string) *CompletionItem {
	var label1 *C.char
	var text1 *C.char
	var stock1 *C.char
	var info1 *C.char
	label1 = _GoStringToGString(label0)
	defer C.free(unsafe.Pointer(label1))
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	stock1 = _GoStringToGString(stock0)
	defer C.free(unsafe.Pointer(stock1))
	info1 = _GoStringToGString(info0)
	defer C.free(unsafe.Pointer(info1))
	ret1 := C.gtk_source_completion_item_new_from_stock(label1, text1, stock1, info1)
	var ret2 *CompletionItem
	ret2 = (*CompletionItem)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewCompletionItemWithMarkup(markup0 string, text0 string, icon0 gdkpixbuf.PixbufLike, info0 string) *CompletionItem {
	var markup1 *C.char
	var text1 *C.char
	var icon1 *C.GdkPixbuf
	var info1 *C.char
	markup1 = _GoStringToGString(markup0)
	defer C.free(unsafe.Pointer(markup1))
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	if icon0 != nil {
		icon1 = icon0.InheritedFromGdkPixbuf()
	}
	info1 = _GoStringToGString(info0)
	defer C.free(unsafe.Pointer(info1))
	ret1 := C.gtk_source_completion_item_new_with_markup(markup1, text1, icon1, info1)
	var ret2 *CompletionItem
	ret2 = (*CompletionItem)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
type CompletionProposalLike interface {
	ImplementsGtkSourceCompletionProposal() *C.GtkSourceCompletionProposal
}

type CompletionProposal struct {
	gobject.Object
	CompletionProposalImpl
}

type CompletionProposalImpl struct {}

func ToCompletionProposal(objlike gobject.ObjectLike) *CompletionProposal {
	t := (*CompletionProposalImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*CompletionProposal)(obj)
	}
	panic("cannot cast to CompletionProposal")
}

func (this0 *CompletionProposalImpl) ImplementsGtkSourceCompletionProposal() *C.GtkSourceCompletionProposal {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GtkSourceCompletionProposal)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *CompletionProposalImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_completion_proposal_get_type())
}

func CompletionProposalGetType() gobject.Type {
	return (*CompletionProposalImpl)(nil).GetStaticType()
}
func (this0 *CompletionProposalImpl) Changed() {
	var this1 *C.GtkSourceCompletionProposal
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProposal()}
	C.gtk_source_completion_proposal_changed(this1)
}
func (this0 *CompletionProposalImpl) Equal(other0 CompletionProposalLike) bool {
	var this1 *C.GtkSourceCompletionProposal
	var other1 *C.GtkSourceCompletionProposal
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProposal()}
	if other0 != nil {
		other1 = other0.ImplementsGtkSourceCompletionProposal()}
	ret1 := C.gtk_source_completion_proposal_equal(this1, other1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *CompletionProposalImpl) GetIcon() *gdkpixbuf.Pixbuf {
	var this1 *C.GtkSourceCompletionProposal
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProposal()}
	ret1 := C.gtk_source_completion_proposal_get_icon(this1)
	var ret2 *gdkpixbuf.Pixbuf
	ret2 = (*gdkpixbuf.Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *CompletionProposalImpl) GetInfo() string {
	var this1 *C.GtkSourceCompletionProposal
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProposal()}
	ret1 := C.gtk_source_completion_proposal_get_info(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *CompletionProposalImpl) GetLabel() string {
	var this1 *C.GtkSourceCompletionProposal
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProposal()}
	ret1 := C.gtk_source_completion_proposal_get_label(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *CompletionProposalImpl) GetMarkup() string {
	var this1 *C.GtkSourceCompletionProposal
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProposal()}
	ret1 := C.gtk_source_completion_proposal_get_markup(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *CompletionProposalImpl) GetText() string {
	var this1 *C.GtkSourceCompletionProposal
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProposal()}
	ret1 := C.gtk_source_completion_proposal_get_text(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *CompletionProposalImpl) Hash() int {
	var this1 *C.GtkSourceCompletionProposal
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProposal()}
	ret1 := C.gtk_source_completion_proposal_hash(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
type CompletionProviderLike interface {
	ImplementsGtkSourceCompletionProvider() *C.GtkSourceCompletionProvider
}

type CompletionProvider struct {
	gobject.Object
	CompletionProviderImpl
}

type CompletionProviderImpl struct {}

func ToCompletionProvider(objlike gobject.ObjectLike) *CompletionProvider {
	t := (*CompletionProviderImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*CompletionProvider)(obj)
	}
	panic("cannot cast to CompletionProvider")
}

func (this0 *CompletionProviderImpl) ImplementsGtkSourceCompletionProvider() *C.GtkSourceCompletionProvider {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GtkSourceCompletionProvider)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *CompletionProviderImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_completion_provider_get_type())
}

func CompletionProviderGetType() gobject.Type {
	return (*CompletionProviderImpl)(nil).GetStaticType()
}
func (this0 *CompletionProviderImpl) ActivateProposal(proposal0 CompletionProposalLike, iter0 *gtk.TextIter) bool {
	var this1 *C.GtkSourceCompletionProvider
	var proposal1 *C.GtkSourceCompletionProposal
	var iter1 *C.GtkTextIter
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	if proposal0 != nil {
		proposal1 = proposal0.ImplementsGtkSourceCompletionProposal()}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	ret1 := C.gtk_source_completion_provider_activate_proposal(this1, proposal1, iter1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *CompletionProviderImpl) GetActivation() CompletionActivation {
	var this1 *C.GtkSourceCompletionProvider
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	ret1 := C.gtk_source_completion_provider_get_activation(this1)
	var ret2 CompletionActivation
	ret2 = CompletionActivation(ret1)
	return ret2
}
func (this0 *CompletionProviderImpl) GetIcon() *gdkpixbuf.Pixbuf {
	var this1 *C.GtkSourceCompletionProvider
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	ret1 := C.gtk_source_completion_provider_get_icon(this1)
	var ret2 *gdkpixbuf.Pixbuf
	ret2 = (*gdkpixbuf.Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *CompletionProviderImpl) GetInfoWidget(proposal0 CompletionProposalLike) *gtk.Widget {
	var this1 *C.GtkSourceCompletionProvider
	var proposal1 *C.GtkSourceCompletionProposal
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	if proposal0 != nil {
		proposal1 = proposal0.ImplementsGtkSourceCompletionProposal()}
	ret1 := C.gtk_source_completion_provider_get_info_widget(this1, proposal1)
	var ret2 *gtk.Widget
	ret2 = (*gtk.Widget)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *CompletionProviderImpl) GetInteractiveDelay() int {
	var this1 *C.GtkSourceCompletionProvider
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	ret1 := C.gtk_source_completion_provider_get_interactive_delay(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *CompletionProviderImpl) GetName() string {
	var this1 *C.GtkSourceCompletionProvider
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	ret1 := C.gtk_source_completion_provider_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *CompletionProviderImpl) GetPriority() int {
	var this1 *C.GtkSourceCompletionProvider
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	ret1 := C.gtk_source_completion_provider_get_priority(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *CompletionProviderImpl) GetStartIter(context0 CompletionContextLike, proposal0 CompletionProposalLike, iter0 *gtk.TextIter) bool {
	var this1 *C.GtkSourceCompletionProvider
	var context1 *C.GtkSourceCompletionContext
	var proposal1 *C.GtkSourceCompletionProposal
	var iter1 *C.GtkTextIter
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	if context0 != nil {
		context1 = context0.InheritedFromGtkSourceCompletionContext()
	}
	if proposal0 != nil {
		proposal1 = proposal0.ImplementsGtkSourceCompletionProposal()}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	ret1 := C.gtk_source_completion_provider_get_start_iter(this1, context1, proposal1, iter1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *CompletionProviderImpl) Match(context0 CompletionContextLike) bool {
	var this1 *C.GtkSourceCompletionProvider
	var context1 *C.GtkSourceCompletionContext
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	if context0 != nil {
		context1 = context0.InheritedFromGtkSourceCompletionContext()
	}
	ret1 := C.gtk_source_completion_provider_match(this1, context1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *CompletionProviderImpl) Populate(context0 CompletionContextLike) {
	var this1 *C.GtkSourceCompletionProvider
	var context1 *C.GtkSourceCompletionContext
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	if context0 != nil {
		context1 = context0.InheritedFromGtkSourceCompletionContext()
	}
	C.gtk_source_completion_provider_populate(this1, context1)
}
func (this0 *CompletionProviderImpl) UpdateInfo(proposal0 CompletionProposalLike, info0 CompletionInfoLike) {
	var this1 *C.GtkSourceCompletionProvider
	var proposal1 *C.GtkSourceCompletionProposal
	var info1 *C.GtkSourceCompletionInfo
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceCompletionProvider()}
	if proposal0 != nil {
		proposal1 = proposal0.ImplementsGtkSourceCompletionProposal()}
	if info0 != nil {
		info1 = info0.InheritedFromGtkSourceCompletionInfo()
	}
	C.gtk_source_completion_provider_update_info(this1, proposal1, info1)
}
type CompletionWordsLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourceCompletionWords() *C.GtkSourceCompletionWords
}

type CompletionWords struct {
	gobject.Object
	CompletionProviderImpl
}

func ToCompletionWords(objlike gobject.ObjectLike) *CompletionWords {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*CompletionWords)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*CompletionWords)(obj)
	}
	panic("cannot cast to CompletionWords")
}

func (this0 *CompletionWords) InheritedFromGtkSourceCompletionWords() *C.GtkSourceCompletionWords {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceCompletionWords)(this0.C)
}

func (this0 *CompletionWords) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_completion_words_get_type())
}

func CompletionWordsGetType() gobject.Type {
	return (*CompletionWords)(nil).GetStaticType()
}
func NewCompletionWords(name0 string, icon0 gdkpixbuf.PixbufLike) *CompletionWords {
	var name1 *C.char
	var icon1 *C.GdkPixbuf
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	if icon0 != nil {
		icon1 = icon0.InheritedFromGdkPixbuf()
	}
	ret1 := C.gtk_source_completion_words_new(name1, icon1)
	var ret2 *CompletionWords
	ret2 = (*CompletionWords)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *CompletionWords) Register(buffer0 gtk.TextBufferLike) {
	var this1 *C.GtkSourceCompletionWords
	var buffer1 *C.GtkTextBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletionWords()
	}
	if buffer0 != nil {
		buffer1 = buffer0.InheritedFromGtkTextBuffer()
	}
	C.gtk_source_completion_words_register(this1, buffer1)
}
func (this0 *CompletionWords) Unregister(buffer0 gtk.TextBufferLike) {
	var this1 *C.GtkSourceCompletionWords
	var buffer1 *C.GtkTextBuffer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceCompletionWords()
	}
	if buffer0 != nil {
		buffer1 = buffer0.InheritedFromGtkTextBuffer()
	}
	C.gtk_source_completion_words_unregister(this1, buffer1)
}
type DrawSpacesFlags C.uint32_t
const (
	DrawSpacesFlagsSpace DrawSpacesFlags = 1
	DrawSpacesFlagsTab DrawSpacesFlags = 2
	DrawSpacesFlagsNewline DrawSpacesFlags = 4
	DrawSpacesFlagsNbsp DrawSpacesFlags = 8
	DrawSpacesFlagsLeading DrawSpacesFlags = 16
	DrawSpacesFlagsText DrawSpacesFlags = 32
	DrawSpacesFlagsTrailing DrawSpacesFlags = 64
	DrawSpacesFlagsAll DrawSpacesFlags = 127
)
type GutterLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourceGutter() *C.GtkSourceGutter
}

type Gutter struct {
	gobject.Object
	
}

func ToGutter(objlike gobject.ObjectLike) *Gutter {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Gutter)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Gutter)(obj)
	}
	panic("cannot cast to Gutter")
}

func (this0 *Gutter) InheritedFromGtkSourceGutter() *C.GtkSourceGutter {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceGutter)(this0.C)
}

func (this0 *Gutter) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_gutter_get_type())
}

func GutterGetType() gobject.Type {
	return (*Gutter)(nil).GetStaticType()
}
func (this0 *Gutter) GetPadding(xpad0 *int, ypad0 *int) {
	var this1 *C.GtkSourceGutter
	var xpad1 *C.int32_t
	var ypad1 *C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutter()
	}
	xpad1 = (*C.int32_t)(unsafe.Pointer(xpad0))
	ypad1 = (*C.int32_t)(unsafe.Pointer(ypad0))
	C.gtk_source_gutter_get_padding(this1, xpad1, ypad1)
}
func (this0 *Gutter) GetRendererAtPos(x0 int, y0 int) *GutterRenderer {
	var this1 *C.GtkSourceGutter
	var x1 C.int32_t
	var y1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutter()
	}
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	ret1 := C.gtk_source_gutter_get_renderer_at_pos(this1, x1, y1)
	var ret2 *GutterRenderer
	ret2 = (*GutterRenderer)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Gutter) GetWindow() *gdk.Window {
	var this1 *C.GtkSourceGutter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutter()
	}
	ret1 := C.gtk_source_gutter_get_window(this1)
	var ret2 *gdk.Window
	ret2 = (*gdk.Window)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Gutter) Insert(renderer0 GutterRendererLike, position0 int) bool {
	var this1 *C.GtkSourceGutter
	var renderer1 *C.GtkSourceGutterRenderer
	var position1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutter()
	}
	if renderer0 != nil {
		renderer1 = renderer0.InheritedFromGtkSourceGutterRenderer()
	}
	position1 = C.int32_t(position0)
	ret1 := C.gtk_source_gutter_insert(this1, renderer1, position1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Gutter) QueueDraw() {
	var this1 *C.GtkSourceGutter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutter()
	}
	C.gtk_source_gutter_queue_draw(this1)
}
func (this0 *Gutter) Remove(renderer0 GutterRendererLike) {
	var this1 *C.GtkSourceGutter
	var renderer1 *C.GtkSourceGutterRenderer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutter()
	}
	if renderer0 != nil {
		renderer1 = renderer0.InheritedFromGtkSourceGutterRenderer()
	}
	C.gtk_source_gutter_remove(this1, renderer1)
}
func (this0 *Gutter) Reorder(renderer0 GutterRendererLike, position0 int) {
	var this1 *C.GtkSourceGutter
	var renderer1 *C.GtkSourceGutterRenderer
	var position1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutter()
	}
	if renderer0 != nil {
		renderer1 = renderer0.InheritedFromGtkSourceGutterRenderer()
	}
	position1 = C.int32_t(position0)
	C.gtk_source_gutter_reorder(this1, renderer1, position1)
}
func (this0 *Gutter) SetPadding(xpad0 int, ypad0 int) {
	var this1 *C.GtkSourceGutter
	var xpad1 C.int32_t
	var ypad1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutter()
	}
	xpad1 = C.int32_t(xpad0)
	ypad1 = C.int32_t(ypad0)
	C.gtk_source_gutter_set_padding(this1, xpad1, ypad1)
}
type GutterRendererLike interface {
	gobject.InitiallyUnownedLike
	InheritedFromGtkSourceGutterRenderer() *C.GtkSourceGutterRenderer
}

type GutterRenderer struct {
	gobject.InitiallyUnowned
	
}

func ToGutterRenderer(objlike gobject.ObjectLike) *GutterRenderer {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*GutterRenderer)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*GutterRenderer)(obj)
	}
	panic("cannot cast to GutterRenderer")
}

func (this0 *GutterRenderer) InheritedFromGtkSourceGutterRenderer() *C.GtkSourceGutterRenderer {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceGutterRenderer)(this0.C)
}

func (this0 *GutterRenderer) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_gutter_renderer_get_type())
}

func GutterRendererGetType() gobject.Type {
	return (*GutterRenderer)(nil).GetStaticType()
}
func (this0 *GutterRenderer) Activate(iter0 *gtk.TextIter, area0 *cairo.RectangleInt, event0 *gdk.Event) {
	var this1 *C.GtkSourceGutterRenderer
	var iter1 *C.GtkTextIter
	var area1 *C.cairoRectangleInt
	var event1 *C.GdkEvent
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	area1 = (*C.cairoRectangleInt)(unsafe.Pointer(area0))
	C.gtk_source_gutter_renderer_activate(this1, iter1, area1, event1)
}
func (this0 *GutterRenderer) Begin(cr0 *cairo.Context, background_area0 *cairo.RectangleInt, cell_area0 *cairo.RectangleInt, start0 *gtk.TextIter, end0 *gtk.TextIter) {
	var this1 *C.GtkSourceGutterRenderer
	var cr1 *C.cairoContext
	var background_area1 *C.cairoRectangleInt
	var cell_area1 *C.cairoRectangleInt
	var start1 *C.GtkTextIter
	var end1 *C.GtkTextIter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	background_area1 = (*C.cairoRectangleInt)(unsafe.Pointer(background_area0))
	cell_area1 = (*C.cairoRectangleInt)(unsafe.Pointer(cell_area0))
	start1 = (*C.GtkTextIter)(unsafe.Pointer(start0))
	end1 = (*C.GtkTextIter)(unsafe.Pointer(end0))
	C.gtk_source_gutter_renderer_begin(this1, cr1, background_area1, cell_area1, start1, end1)
}
func (this0 *GutterRenderer) Draw(cr0 *cairo.Context, background_area0 *cairo.RectangleInt, cell_area0 *cairo.RectangleInt, start0 *gtk.TextIter, end0 *gtk.TextIter, state0 GutterRendererState) {
	var this1 *C.GtkSourceGutterRenderer
	var cr1 *C.cairoContext
	var background_area1 *C.cairoRectangleInt
	var cell_area1 *C.cairoRectangleInt
	var start1 *C.GtkTextIter
	var end1 *C.GtkTextIter
	var state1 C.GtkSourceGutterRendererState
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	if cr0 != nil {
		cr1 = (*C.cairoContext)(cr0.C)
	}
	background_area1 = (*C.cairoRectangleInt)(unsafe.Pointer(background_area0))
	cell_area1 = (*C.cairoRectangleInt)(unsafe.Pointer(cell_area0))
	start1 = (*C.GtkTextIter)(unsafe.Pointer(start0))
	end1 = (*C.GtkTextIter)(unsafe.Pointer(end0))
	state1 = C.GtkSourceGutterRendererState(state0)
	C.gtk_source_gutter_renderer_draw(this1, cr1, background_area1, cell_area1, start1, end1, state1)
}
func (this0 *GutterRenderer) End() {
	var this1 *C.GtkSourceGutterRenderer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	C.gtk_source_gutter_renderer_end(this1)
}
func (this0 *GutterRenderer) GetAlignment() (float64, float64) {
	var this1 *C.GtkSourceGutterRenderer
	var xalign1 C.float
	var yalign1 C.float
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	C.gtk_source_gutter_renderer_get_alignment(this1, &xalign1, &yalign1)
	var xalign2 float64
	var yalign2 float64
	xalign2 = float64(xalign1)
	yalign2 = float64(yalign1)
	return xalign2, yalign2
}
func (this0 *GutterRenderer) GetAlignmentMode() GutterRendererAlignmentMode {
	var this1 *C.GtkSourceGutterRenderer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	ret1 := C.gtk_source_gutter_renderer_get_alignment_mode(this1)
	var ret2 GutterRendererAlignmentMode
	ret2 = GutterRendererAlignmentMode(ret1)
	return ret2
}
func (this0 *GutterRenderer) GetBackground() (gdk.RGBA, bool) {
	var this1 *C.GtkSourceGutterRenderer
	var color1 C.GdkRGBA
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	ret1 := C.gtk_source_gutter_renderer_get_background(this1, &color1)
	var color2 gdk.RGBA
	var ret2 bool
	color2 = *(*gdk.RGBA)(unsafe.Pointer(&color1))
	ret2 = ret1 != 0
	return color2, ret2
}
func (this0 *GutterRenderer) GetPadding() (int, int) {
	var this1 *C.GtkSourceGutterRenderer
	var xpad1 C.int32_t
	var ypad1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	C.gtk_source_gutter_renderer_get_padding(this1, &xpad1, &ypad1)
	var xpad2 int
	var ypad2 int
	xpad2 = int(xpad1)
	ypad2 = int(ypad1)
	return xpad2, ypad2
}
func (this0 *GutterRenderer) GetSize() int {
	var this1 *C.GtkSourceGutterRenderer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	ret1 := C.gtk_source_gutter_renderer_get_size(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *GutterRenderer) GetView() *gtk.TextView {
	var this1 *C.GtkSourceGutterRenderer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	ret1 := C.gtk_source_gutter_renderer_get_view(this1)
	var ret2 *gtk.TextView
	ret2 = (*gtk.TextView)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *GutterRenderer) GetVisible() bool {
	var this1 *C.GtkSourceGutterRenderer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	ret1 := C.gtk_source_gutter_renderer_get_visible(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *GutterRenderer) GetWindowType() gtk.TextWindowType {
	var this1 *C.GtkSourceGutterRenderer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	ret1 := C.gtk_source_gutter_renderer_get_window_type(this1)
	var ret2 gtk.TextWindowType
	ret2 = gtk.TextWindowType(ret1)
	return ret2
}
func (this0 *GutterRenderer) QueryActivatable(iter0 *gtk.TextIter, area0 *cairo.RectangleInt, event0 *gdk.Event) bool {
	var this1 *C.GtkSourceGutterRenderer
	var iter1 *C.GtkTextIter
	var area1 *C.cairoRectangleInt
	var event1 *C.GdkEvent
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	area1 = (*C.cairoRectangleInt)(unsafe.Pointer(area0))
	ret1 := C.gtk_source_gutter_renderer_query_activatable(this1, iter1, area1, event1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *GutterRenderer) QueryData(start0 *gtk.TextIter, end0 *gtk.TextIter, state0 GutterRendererState) {
	var this1 *C.GtkSourceGutterRenderer
	var start1 *C.GtkTextIter
	var end1 *C.GtkTextIter
	var state1 C.GtkSourceGutterRendererState
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	start1 = (*C.GtkTextIter)(unsafe.Pointer(start0))
	end1 = (*C.GtkTextIter)(unsafe.Pointer(end0))
	state1 = C.GtkSourceGutterRendererState(state0)
	C.gtk_source_gutter_renderer_query_data(this1, start1, end1, state1)
}
func (this0 *GutterRenderer) QueryTooltip(iter0 *gtk.TextIter, area0 *cairo.RectangleInt, x0 int, y0 int, tooltip0 gtk.TooltipLike) bool {
	var this1 *C.GtkSourceGutterRenderer
	var iter1 *C.GtkTextIter
	var area1 *C.cairoRectangleInt
	var x1 C.int32_t
	var y1 C.int32_t
	var tooltip1 *C.GtkTooltip
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	area1 = (*C.cairoRectangleInt)(unsafe.Pointer(area0))
	x1 = C.int32_t(x0)
	y1 = C.int32_t(y0)
	if tooltip0 != nil {
		tooltip1 = tooltip0.InheritedFromGtkTooltip()
	}
	ret1 := C.gtk_source_gutter_renderer_query_tooltip(this1, iter1, area1, x1, y1, tooltip1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *GutterRenderer) QueueDraw() {
	var this1 *C.GtkSourceGutterRenderer
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	C.gtk_source_gutter_renderer_queue_draw(this1)
}
func (this0 *GutterRenderer) SetAlignment(xalign0 float64, yalign0 float64) {
	var this1 *C.GtkSourceGutterRenderer
	var xalign1 C.float
	var yalign1 C.float
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	xalign1 = C.float(xalign0)
	yalign1 = C.float(yalign0)
	C.gtk_source_gutter_renderer_set_alignment(this1, xalign1, yalign1)
}
func (this0 *GutterRenderer) SetAlignmentMode(mode0 GutterRendererAlignmentMode) {
	var this1 *C.GtkSourceGutterRenderer
	var mode1 C.GtkSourceGutterRendererAlignmentMode
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	mode1 = C.GtkSourceGutterRendererAlignmentMode(mode0)
	C.gtk_source_gutter_renderer_set_alignment_mode(this1, mode1)
}
func (this0 *GutterRenderer) SetBackground(color0 *gdk.RGBA) {
	var this1 *C.GtkSourceGutterRenderer
	var color1 *C.GdkRGBA
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	color1 = (*C.GdkRGBA)(unsafe.Pointer(color0))
	C.gtk_source_gutter_renderer_set_background(this1, color1)
}
func (this0 *GutterRenderer) SetPadding(xpad0 int, ypad0 int) {
	var this1 *C.GtkSourceGutterRenderer
	var xpad1 C.int32_t
	var ypad1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	xpad1 = C.int32_t(xpad0)
	ypad1 = C.int32_t(ypad0)
	C.gtk_source_gutter_renderer_set_padding(this1, xpad1, ypad1)
}
func (this0 *GutterRenderer) SetSize(size0 int) {
	var this1 *C.GtkSourceGutterRenderer
	var size1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	size1 = C.int32_t(size0)
	C.gtk_source_gutter_renderer_set_size(this1, size1)
}
func (this0 *GutterRenderer) SetVisible(visible0 bool) {
	var this1 *C.GtkSourceGutterRenderer
	var visible1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRenderer()
	}
	visible1 = _GoBoolToCBool(visible0)
	C.gtk_source_gutter_renderer_set_visible(this1, visible1)
}
type GutterRendererAlignmentMode C.uint32_t
const (
	GutterRendererAlignmentModeCell GutterRendererAlignmentMode = 0
	GutterRendererAlignmentModeFirst GutterRendererAlignmentMode = 1
	GutterRendererAlignmentModeLast GutterRendererAlignmentMode = 2
)
type GutterRendererPixbufLike interface {
	GutterRendererLike
	InheritedFromGtkSourceGutterRendererPixbuf() *C.GtkSourceGutterRendererPixbuf
}

type GutterRendererPixbuf struct {
	GutterRenderer
	
}

func ToGutterRendererPixbuf(objlike gobject.ObjectLike) *GutterRendererPixbuf {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*GutterRendererPixbuf)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*GutterRendererPixbuf)(obj)
	}
	panic("cannot cast to GutterRendererPixbuf")
}

func (this0 *GutterRendererPixbuf) InheritedFromGtkSourceGutterRendererPixbuf() *C.GtkSourceGutterRendererPixbuf {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceGutterRendererPixbuf)(this0.C)
}

func (this0 *GutterRendererPixbuf) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_gutter_renderer_pixbuf_get_type())
}

func GutterRendererPixbufGetType() gobject.Type {
	return (*GutterRendererPixbuf)(nil).GetStaticType()
}
func NewGutterRendererPixbuf() *GutterRendererPixbuf {
	ret1 := C.gtk_source_gutter_renderer_pixbuf_new()
	var ret2 *GutterRendererPixbuf
	ret2 = (*GutterRendererPixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *GutterRendererPixbuf) GetGIcon() *gio.Icon {
	var this1 *C.GtkSourceGutterRendererPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererPixbuf()
	}
	ret1 := C.gtk_source_gutter_renderer_pixbuf_get_gicon(this1)
	var ret2 *gio.Icon
	ret2 = (*gio.Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *GutterRendererPixbuf) GetIconName() string {
	var this1 *C.GtkSourceGutterRendererPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererPixbuf()
	}
	ret1 := C.gtk_source_gutter_renderer_pixbuf_get_icon_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *GutterRendererPixbuf) GetPixbuf() *gdkpixbuf.Pixbuf {
	var this1 *C.GtkSourceGutterRendererPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererPixbuf()
	}
	ret1 := C.gtk_source_gutter_renderer_pixbuf_get_pixbuf(this1)
	var ret2 *gdkpixbuf.Pixbuf
	ret2 = (*gdkpixbuf.Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *GutterRendererPixbuf) GetStockID() string {
	var this1 *C.GtkSourceGutterRendererPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererPixbuf()
	}
	ret1 := C.gtk_source_gutter_renderer_pixbuf_get_stock_id(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *GutterRendererPixbuf) SetGIcon(icon0 gio.IconLike) {
	var this1 *C.GtkSourceGutterRendererPixbuf
	var icon1 *C.GIcon
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererPixbuf()
	}
	if icon0 != nil {
		icon1 = icon0.ImplementsGIcon()}
	C.gtk_source_gutter_renderer_pixbuf_set_gicon(this1, icon1)
}
func (this0 *GutterRendererPixbuf) SetIconName(icon_name0 string) {
	var this1 *C.GtkSourceGutterRendererPixbuf
	var icon_name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererPixbuf()
	}
	icon_name1 = _GoStringToGString(icon_name0)
	defer C.free(unsafe.Pointer(icon_name1))
	C.gtk_source_gutter_renderer_pixbuf_set_icon_name(this1, icon_name1)
}
func (this0 *GutterRendererPixbuf) SetPixbuf(pixbuf0 gdkpixbuf.PixbufLike) {
	var this1 *C.GtkSourceGutterRendererPixbuf
	var pixbuf1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererPixbuf()
	}
	if pixbuf0 != nil {
		pixbuf1 = pixbuf0.InheritedFromGdkPixbuf()
	}
	C.gtk_source_gutter_renderer_pixbuf_set_pixbuf(this1, pixbuf1)
}
func (this0 *GutterRendererPixbuf) SetStockID(stock_id0 string) {
	var this1 *C.GtkSourceGutterRendererPixbuf
	var stock_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererPixbuf()
	}
	stock_id1 = _GoStringToGString(stock_id0)
	defer C.free(unsafe.Pointer(stock_id1))
	C.gtk_source_gutter_renderer_pixbuf_set_stock_id(this1, stock_id1)
}
type GutterRendererState C.uint32_t
const (
	GutterRendererStateNormal GutterRendererState = 0
	GutterRendererStateCursor GutterRendererState = 1
	GutterRendererStatePrelit GutterRendererState = 2
	GutterRendererStateSelected GutterRendererState = 4
)
type GutterRendererTextLike interface {
	GutterRendererLike
	InheritedFromGtkSourceGutterRendererText() *C.GtkSourceGutterRendererText
}

type GutterRendererText struct {
	GutterRenderer
	
}

func ToGutterRendererText(objlike gobject.ObjectLike) *GutterRendererText {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*GutterRendererText)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*GutterRendererText)(obj)
	}
	panic("cannot cast to GutterRendererText")
}

func (this0 *GutterRendererText) InheritedFromGtkSourceGutterRendererText() *C.GtkSourceGutterRendererText {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceGutterRendererText)(this0.C)
}

func (this0 *GutterRendererText) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_gutter_renderer_text_get_type())
}

func GutterRendererTextGetType() gobject.Type {
	return (*GutterRendererText)(nil).GetStaticType()
}
func NewGutterRendererText() *GutterRendererText {
	ret1 := C.gtk_source_gutter_renderer_text_new()
	var ret2 *GutterRendererText
	ret2 = (*GutterRendererText)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *GutterRendererText) Measure(text0 string, width0 *int, height0 *int) {
	var this1 *C.GtkSourceGutterRendererText
	var text1 *C.char
	var width1 *C.int32_t
	var height1 *C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererText()
	}
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	width1 = (*C.int32_t)(unsafe.Pointer(width0))
	height1 = (*C.int32_t)(unsafe.Pointer(height0))
	C.gtk_source_gutter_renderer_text_measure(this1, text1, width1, height1)
}
func (this0 *GutterRendererText) MeasureMarkup(markup0 string, width0 *int, height0 *int) {
	var this1 *C.GtkSourceGutterRendererText
	var markup1 *C.char
	var width1 *C.int32_t
	var height1 *C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererText()
	}
	markup1 = _GoStringToGString(markup0)
	defer C.free(unsafe.Pointer(markup1))
	width1 = (*C.int32_t)(unsafe.Pointer(width0))
	height1 = (*C.int32_t)(unsafe.Pointer(height0))
	C.gtk_source_gutter_renderer_text_measure_markup(this1, markup1, width1, height1)
}
func (this0 *GutterRendererText) SetMarkup(markup0 string, length0 int) {
	var this1 *C.GtkSourceGutterRendererText
	var markup1 *C.char
	var length1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererText()
	}
	markup1 = _GoStringToGString(markup0)
	defer C.free(unsafe.Pointer(markup1))
	length1 = C.int32_t(length0)
	C.gtk_source_gutter_renderer_text_set_markup(this1, markup1, length1)
}
func (this0 *GutterRendererText) SetText(text0 string, length0 int) {
	var this1 *C.GtkSourceGutterRendererText
	var text1 *C.char
	var length1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceGutterRendererText()
	}
	text1 = _GoStringToGString(text0)
	defer C.free(unsafe.Pointer(text1))
	length1 = C.int32_t(length0)
	C.gtk_source_gutter_renderer_text_set_text(this1, text1, length1)
}
type LanguageLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourceLanguage() *C.GtkSourceLanguage
}

type Language struct {
	gobject.Object
	
}

func ToLanguage(objlike gobject.ObjectLike) *Language {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Language)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Language)(obj)
	}
	panic("cannot cast to Language")
}

func (this0 *Language) InheritedFromGtkSourceLanguage() *C.GtkSourceLanguage {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceLanguage)(this0.C)
}

func (this0 *Language) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_language_get_type())
}

func LanguageGetType() gobject.Type {
	return (*Language)(nil).GetStaticType()
}
func (this0 *Language) GetGlobs() []string {
	var this1 *C.GtkSourceLanguage
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguage()
	}
	ret1 := C.gtk_source_language_get_globs(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *Language) GetHidden() bool {
	var this1 *C.GtkSourceLanguage
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguage()
	}
	ret1 := C.gtk_source_language_get_hidden(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *Language) GetID() string {
	var this1 *C.GtkSourceLanguage
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguage()
	}
	ret1 := C.gtk_source_language_get_id(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Language) GetMetadata(name0 string) string {
	var this1 *C.GtkSourceLanguage
	var name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguage()
	}
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	ret1 := C.gtk_source_language_get_metadata(this1, name1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Language) GetMIMETypes() []string {
	var this1 *C.GtkSourceLanguage
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguage()
	}
	ret1 := C.gtk_source_language_get_mime_types(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *Language) GetName() string {
	var this1 *C.GtkSourceLanguage
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguage()
	}
	ret1 := C.gtk_source_language_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Language) GetSection() string {
	var this1 *C.GtkSourceLanguage
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguage()
	}
	ret1 := C.gtk_source_language_get_section(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Language) GetStyleIDs() []string {
	var this1 *C.GtkSourceLanguage
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguage()
	}
	ret1 := C.gtk_source_language_get_style_ids(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
		C.g_free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i]))
	}
	return ret2
}
func (this0 *Language) GetStyleName(style_id0 string) string {
	var this1 *C.GtkSourceLanguage
	var style_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguage()
	}
	style_id1 = _GoStringToGString(style_id0)
	defer C.free(unsafe.Pointer(style_id1))
	ret1 := C.gtk_source_language_get_style_name(this1, style_id1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
type LanguageManagerLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourceLanguageManager() *C.GtkSourceLanguageManager
}

type LanguageManager struct {
	gobject.Object
	
}

func ToLanguageManager(objlike gobject.ObjectLike) *LanguageManager {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*LanguageManager)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*LanguageManager)(obj)
	}
	panic("cannot cast to LanguageManager")
}

func (this0 *LanguageManager) InheritedFromGtkSourceLanguageManager() *C.GtkSourceLanguageManager {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceLanguageManager)(this0.C)
}

func (this0 *LanguageManager) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_language_manager_get_type())
}

func LanguageManagerGetType() gobject.Type {
	return (*LanguageManager)(nil).GetStaticType()
}
func NewLanguageManager() *LanguageManager {
	ret1 := C.gtk_source_language_manager_new()
	var ret2 *LanguageManager
	ret2 = (*LanguageManager)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func LanguageManagerGetDefault() *LanguageManager {
	ret1 := C.gtk_source_language_manager_get_default()
	var ret2 *LanguageManager
	ret2 = (*LanguageManager)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *LanguageManager) GetLanguage(id0 string) *Language {
	var this1 *C.GtkSourceLanguageManager
	var id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguageManager()
	}
	id1 = _GoStringToGString(id0)
	defer C.free(unsafe.Pointer(id1))
	ret1 := C.gtk_source_language_manager_get_language(this1, id1)
	var ret2 *Language
	ret2 = (*Language)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *LanguageManager) GetLanguageIDs() []string {
	var this1 *C.GtkSourceLanguageManager
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguageManager()
	}
	ret1 := C.gtk_source_language_manager_get_language_ids(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func (this0 *LanguageManager) GetSearchPath() []string {
	var this1 *C.GtkSourceLanguageManager
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguageManager()
	}
	ret1 := C.gtk_source_language_manager_get_search_path(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func (this0 *LanguageManager) GuessLanguage(filename0 string, content_type0 string) *Language {
	var this1 *C.GtkSourceLanguageManager
	var filename1 *C.char
	var content_type1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguageManager()
	}
	filename1 = _GoStringToGString(filename0)
	defer C.free(unsafe.Pointer(filename1))
	content_type1 = _GoStringToGString(content_type0)
	defer C.free(unsafe.Pointer(content_type1))
	ret1 := C.gtk_source_language_manager_guess_language(this1, filename1, content_type1)
	var ret2 *Language
	ret2 = (*Language)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *LanguageManager) SetSearchPath(dirs0 []string) {
	var this1 *C.GtkSourceLanguageManager
	var dirs1 **C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceLanguageManager()
	}
	dirs1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*dirs1)) * (len(dirs0) + 1))))
	defer C.free(unsafe.Pointer(dirs1))
	for i, e := range dirs0 {
		(*(*[999999]*C.char)(unsafe.Pointer(dirs1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(dirs1)))[i]))
	}
	(*(*[999999]*C.char)(unsafe.Pointer(dirs1)))[len(dirs0)] = nil
	C.gtk_source_language_manager_set_search_path(this1, dirs1)
}
type MarkLike interface {
	gtk.TextMarkLike
	InheritedFromGtkSourceMark() *C.GtkSourceMark
}

type Mark struct {
	gtk.TextMark
	
}

func ToMark(objlike gobject.ObjectLike) *Mark {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Mark)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Mark)(obj)
	}
	panic("cannot cast to Mark")
}

func (this0 *Mark) InheritedFromGtkSourceMark() *C.GtkSourceMark {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceMark)(this0.C)
}

func (this0 *Mark) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_mark_get_type())
}

func MarkGetType() gobject.Type {
	return (*Mark)(nil).GetStaticType()
}
func NewMark(name0 string, category0 string) *Mark {
	var name1 *C.char
	var category1 *C.char
	name1 = _GoStringToGString(name0)
	defer C.free(unsafe.Pointer(name1))
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	ret1 := C.gtk_source_mark_new(name1, category1)
	var ret2 *Mark
	ret2 = (*Mark)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *Mark) GetCategory() string {
	var this1 *C.GtkSourceMark
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMark()
	}
	ret1 := C.gtk_source_mark_get_category(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *Mark) Next(category0 string) *Mark {
	var this1 *C.GtkSourceMark
	var category1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMark()
	}
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	ret1 := C.gtk_source_mark_next(this1, category1)
	var ret2 *Mark
	ret2 = (*Mark)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *Mark) Prev(category0 string) *Mark {
	var this1 *C.GtkSourceMark
	var category1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMark()
	}
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	ret1 := C.gtk_source_mark_prev(this1, category1)
	var ret2 *Mark
	ret2 = (*Mark)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
type MarkAttributesLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourceMarkAttributes() *C.GtkSourceMarkAttributes
}

type MarkAttributes struct {
	gobject.Object
	
}

func ToMarkAttributes(objlike gobject.ObjectLike) *MarkAttributes {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*MarkAttributes)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*MarkAttributes)(obj)
	}
	panic("cannot cast to MarkAttributes")
}

func (this0 *MarkAttributes) InheritedFromGtkSourceMarkAttributes() *C.GtkSourceMarkAttributes {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceMarkAttributes)(this0.C)
}

func (this0 *MarkAttributes) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_mark_attributes_get_type())
}

func MarkAttributesGetType() gobject.Type {
	return (*MarkAttributes)(nil).GetStaticType()
}
func NewMarkAttributes() *MarkAttributes {
	ret1 := C.gtk_source_mark_attributes_new()
	var ret2 *MarkAttributes
	ret2 = (*MarkAttributes)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *MarkAttributes) GetBackground() (gdk.RGBA, bool) {
	var this1 *C.GtkSourceMarkAttributes
	var background1 C.GdkRGBA
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	ret1 := C.gtk_source_mark_attributes_get_background(this1, &background1)
	var background2 gdk.RGBA
	var ret2 bool
	background2 = *(*gdk.RGBA)(unsafe.Pointer(&background1))
	ret2 = ret1 != 0
	return background2, ret2
}
func (this0 *MarkAttributes) GetGIcon() *gio.Icon {
	var this1 *C.GtkSourceMarkAttributes
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	ret1 := C.gtk_source_mark_attributes_get_gicon(this1)
	var ret2 *gio.Icon
	ret2 = (*gio.Icon)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *MarkAttributes) GetIconName() string {
	var this1 *C.GtkSourceMarkAttributes
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	ret1 := C.gtk_source_mark_attributes_get_icon_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MarkAttributes) GetPixbuf() *gdkpixbuf.Pixbuf {
	var this1 *C.GtkSourceMarkAttributes
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	ret1 := C.gtk_source_mark_attributes_get_pixbuf(this1)
	var ret2 *gdkpixbuf.Pixbuf
	ret2 = (*gdkpixbuf.Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *MarkAttributes) GetStockID() string {
	var this1 *C.GtkSourceMarkAttributes
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	ret1 := C.gtk_source_mark_attributes_get_stock_id(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *MarkAttributes) GetTooltipMarkup(mark0 MarkLike) string {
	var this1 *C.GtkSourceMarkAttributes
	var mark1 *C.GtkSourceMark
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	if mark0 != nil {
		mark1 = mark0.InheritedFromGtkSourceMark()
	}
	ret1 := C.gtk_source_mark_attributes_get_tooltip_markup(this1, mark1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *MarkAttributes) GetTooltipText(mark0 MarkLike) string {
	var this1 *C.GtkSourceMarkAttributes
	var mark1 *C.GtkSourceMark
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	if mark0 != nil {
		mark1 = mark0.InheritedFromGtkSourceMark()
	}
	ret1 := C.gtk_source_mark_attributes_get_tooltip_text(this1, mark1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *MarkAttributes) RenderIcon(widget0 gtk.WidgetLike, size0 int) *gdkpixbuf.Pixbuf {
	var this1 *C.GtkSourceMarkAttributes
	var widget1 *C.GtkWidget
	var size1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	if widget0 != nil {
		widget1 = widget0.InheritedFromGtkWidget()
	}
	size1 = C.int32_t(size0)
	ret1 := C.gtk_source_mark_attributes_render_icon(this1, widget1, size1)
	var ret2 *gdkpixbuf.Pixbuf
	ret2 = (*gdkpixbuf.Pixbuf)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *MarkAttributes) SetBackground(background0 *gdk.RGBA) {
	var this1 *C.GtkSourceMarkAttributes
	var background1 *C.GdkRGBA
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	background1 = (*C.GdkRGBA)(unsafe.Pointer(background0))
	C.gtk_source_mark_attributes_set_background(this1, background1)
}
func (this0 *MarkAttributes) SetGIcon(gicon0 gio.IconLike) {
	var this1 *C.GtkSourceMarkAttributes
	var gicon1 *C.GIcon
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	if gicon0 != nil {
		gicon1 = gicon0.ImplementsGIcon()}
	C.gtk_source_mark_attributes_set_gicon(this1, gicon1)
}
func (this0 *MarkAttributes) SetIconName(icon_name0 string) {
	var this1 *C.GtkSourceMarkAttributes
	var icon_name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	icon_name1 = _GoStringToGString(icon_name0)
	defer C.free(unsafe.Pointer(icon_name1))
	C.gtk_source_mark_attributes_set_icon_name(this1, icon_name1)
}
func (this0 *MarkAttributes) SetPixbuf(pixbuf0 gdkpixbuf.PixbufLike) {
	var this1 *C.GtkSourceMarkAttributes
	var pixbuf1 *C.GdkPixbuf
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	if pixbuf0 != nil {
		pixbuf1 = pixbuf0.InheritedFromGdkPixbuf()
	}
	C.gtk_source_mark_attributes_set_pixbuf(this1, pixbuf1)
}
func (this0 *MarkAttributes) SetStockID(stock_id0 string) {
	var this1 *C.GtkSourceMarkAttributes
	var stock_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceMarkAttributes()
	}
	stock_id1 = _GoStringToGString(stock_id0)
	defer C.free(unsafe.Pointer(stock_id1))
	C.gtk_source_mark_attributes_set_stock_id(this1, stock_id1)
}
type PrintCompositorLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourcePrintCompositor() *C.GtkSourcePrintCompositor
}

type PrintCompositor struct {
	gobject.Object
	
}

func ToPrintCompositor(objlike gobject.ObjectLike) *PrintCompositor {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*PrintCompositor)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*PrintCompositor)(obj)
	}
	panic("cannot cast to PrintCompositor")
}

func (this0 *PrintCompositor) InheritedFromGtkSourcePrintCompositor() *C.GtkSourcePrintCompositor {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourcePrintCompositor)(this0.C)
}

func (this0 *PrintCompositor) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_print_compositor_get_type())
}

func PrintCompositorGetType() gobject.Type {
	return (*PrintCompositor)(nil).GetStaticType()
}
func NewPrintCompositor(buffer0 BufferLike) *PrintCompositor {
	var buffer1 *C.GtkSourceBuffer
	if buffer0 != nil {
		buffer1 = buffer0.InheritedFromGtkSourceBuffer()
	}
	ret1 := C.gtk_source_print_compositor_new(buffer1)
	var ret2 *PrintCompositor
	ret2 = (*PrintCompositor)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func NewPrintCompositorFromView(view0 ViewLike) *PrintCompositor {
	var view1 *C.GtkSourceView
	if view0 != nil {
		view1 = view0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_print_compositor_new_from_view(view1)
	var ret2 *PrintCompositor
	ret2 = (*PrintCompositor)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func (this0 *PrintCompositor) DrawPage(context0 gtk.PrintContextLike, page_nr0 int) {
	var this1 *C.GtkSourcePrintCompositor
	var context1 *C.GtkPrintContext
	var page_nr1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	if context0 != nil {
		context1 = context0.InheritedFromGtkPrintContext()
	}
	page_nr1 = C.int32_t(page_nr0)
	C.gtk_source_print_compositor_draw_page(this1, context1, page_nr1)
}
func (this0 *PrintCompositor) GetBodyFontName() string {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_body_font_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *PrintCompositor) GetBottomMargin(unit0 gtk.Unit) float64 {
	var this1 *C.GtkSourcePrintCompositor
	var unit1 C.GtkUnit
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	unit1 = C.GtkUnit(unit0)
	ret1 := C.gtk_source_print_compositor_get_bottom_margin(this1, unit1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *PrintCompositor) GetBuffer() *Buffer {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_buffer(this1)
	var ret2 *Buffer
	ret2 = (*Buffer)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *PrintCompositor) GetFooterFontName() string {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_footer_font_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *PrintCompositor) GetHeaderFontName() string {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_header_font_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *PrintCompositor) GetHighlightSyntax() bool {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_highlight_syntax(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *PrintCompositor) GetLeftMargin(unit0 gtk.Unit) float64 {
	var this1 *C.GtkSourcePrintCompositor
	var unit1 C.GtkUnit
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	unit1 = C.GtkUnit(unit0)
	ret1 := C.gtk_source_print_compositor_get_left_margin(this1, unit1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *PrintCompositor) GetLineNumbersFontName() string {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_line_numbers_font_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	C.g_free(unsafe.Pointer(ret1))
	return ret2
}
func (this0 *PrintCompositor) GetNPages() int {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_n_pages(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *PrintCompositor) GetPaginationProgress() float64 {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_pagination_progress(this1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *PrintCompositor) GetPrintFooter() bool {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_print_footer(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *PrintCompositor) GetPrintHeader() bool {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_print_header(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *PrintCompositor) GetPrintLineNumbers() int {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_print_line_numbers(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *PrintCompositor) GetRightMargin(unit0 gtk.Unit) float64 {
	var this1 *C.GtkSourcePrintCompositor
	var unit1 C.GtkUnit
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	unit1 = C.GtkUnit(unit0)
	ret1 := C.gtk_source_print_compositor_get_right_margin(this1, unit1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *PrintCompositor) GetTabWidth() int {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_tab_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *PrintCompositor) GetTopMargin(unit0 gtk.Unit) float64 {
	var this1 *C.GtkSourcePrintCompositor
	var unit1 C.GtkUnit
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	unit1 = C.GtkUnit(unit0)
	ret1 := C.gtk_source_print_compositor_get_top_margin(this1, unit1)
	var ret2 float64
	ret2 = float64(ret1)
	return ret2
}
func (this0 *PrintCompositor) GetWrapMode() gtk.WrapMode {
	var this1 *C.GtkSourcePrintCompositor
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	ret1 := C.gtk_source_print_compositor_get_wrap_mode(this1)
	var ret2 gtk.WrapMode
	ret2 = gtk.WrapMode(ret1)
	return ret2
}
func (this0 *PrintCompositor) Paginate(context0 gtk.PrintContextLike) bool {
	var this1 *C.GtkSourcePrintCompositor
	var context1 *C.GtkPrintContext
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	if context0 != nil {
		context1 = context0.InheritedFromGtkPrintContext()
	}
	ret1 := C.gtk_source_print_compositor_paginate(this1, context1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *PrintCompositor) SetBodyFontName(font_name0 string) {
	var this1 *C.GtkSourcePrintCompositor
	var font_name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	font_name1 = _GoStringToGString(font_name0)
	defer C.free(unsafe.Pointer(font_name1))
	C.gtk_source_print_compositor_set_body_font_name(this1, font_name1)
}
func (this0 *PrintCompositor) SetBottomMargin(margin0 float64, unit0 gtk.Unit) {
	var this1 *C.GtkSourcePrintCompositor
	var margin1 C.double
	var unit1 C.GtkUnit
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	margin1 = C.double(margin0)
	unit1 = C.GtkUnit(unit0)
	C.gtk_source_print_compositor_set_bottom_margin(this1, margin1, unit1)
}
func (this0 *PrintCompositor) SetFooterFontName(font_name0 string) {
	var this1 *C.GtkSourcePrintCompositor
	var font_name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	font_name1 = _GoStringToGString(font_name0)
	defer C.free(unsafe.Pointer(font_name1))
	C.gtk_source_print_compositor_set_footer_font_name(this1, font_name1)
}
func (this0 *PrintCompositor) SetFooterFormat(separator0 bool, left0 string, center0 string, right0 string) {
	var this1 *C.GtkSourcePrintCompositor
	var separator1 C.int
	var left1 *C.char
	var center1 *C.char
	var right1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	separator1 = _GoBoolToCBool(separator0)
	left1 = _GoStringToGString(left0)
	defer C.free(unsafe.Pointer(left1))
	center1 = _GoStringToGString(center0)
	defer C.free(unsafe.Pointer(center1))
	right1 = _GoStringToGString(right0)
	defer C.free(unsafe.Pointer(right1))
	C.gtk_source_print_compositor_set_footer_format(this1, separator1, left1, center1, right1)
}
func (this0 *PrintCompositor) SetHeaderFontName(font_name0 string) {
	var this1 *C.GtkSourcePrintCompositor
	var font_name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	font_name1 = _GoStringToGString(font_name0)
	defer C.free(unsafe.Pointer(font_name1))
	C.gtk_source_print_compositor_set_header_font_name(this1, font_name1)
}
func (this0 *PrintCompositor) SetHeaderFormat(separator0 bool, left0 string, center0 string, right0 string) {
	var this1 *C.GtkSourcePrintCompositor
	var separator1 C.int
	var left1 *C.char
	var center1 *C.char
	var right1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	separator1 = _GoBoolToCBool(separator0)
	left1 = _GoStringToGString(left0)
	defer C.free(unsafe.Pointer(left1))
	center1 = _GoStringToGString(center0)
	defer C.free(unsafe.Pointer(center1))
	right1 = _GoStringToGString(right0)
	defer C.free(unsafe.Pointer(right1))
	C.gtk_source_print_compositor_set_header_format(this1, separator1, left1, center1, right1)
}
func (this0 *PrintCompositor) SetHighlightSyntax(highlight0 bool) {
	var this1 *C.GtkSourcePrintCompositor
	var highlight1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	highlight1 = _GoBoolToCBool(highlight0)
	C.gtk_source_print_compositor_set_highlight_syntax(this1, highlight1)
}
func (this0 *PrintCompositor) SetLeftMargin(margin0 float64, unit0 gtk.Unit) {
	var this1 *C.GtkSourcePrintCompositor
	var margin1 C.double
	var unit1 C.GtkUnit
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	margin1 = C.double(margin0)
	unit1 = C.GtkUnit(unit0)
	C.gtk_source_print_compositor_set_left_margin(this1, margin1, unit1)
}
func (this0 *PrintCompositor) SetLineNumbersFontName(font_name0 string) {
	var this1 *C.GtkSourcePrintCompositor
	var font_name1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	font_name1 = _GoStringToGString(font_name0)
	defer C.free(unsafe.Pointer(font_name1))
	C.gtk_source_print_compositor_set_line_numbers_font_name(this1, font_name1)
}
func (this0 *PrintCompositor) SetPrintFooter(print0 bool) {
	var this1 *C.GtkSourcePrintCompositor
	var print1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	print1 = _GoBoolToCBool(print0)
	C.gtk_source_print_compositor_set_print_footer(this1, print1)
}
func (this0 *PrintCompositor) SetPrintHeader(print0 bool) {
	var this1 *C.GtkSourcePrintCompositor
	var print1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	print1 = _GoBoolToCBool(print0)
	C.gtk_source_print_compositor_set_print_header(this1, print1)
}
func (this0 *PrintCompositor) SetPrintLineNumbers(interval0 int) {
	var this1 *C.GtkSourcePrintCompositor
	var interval1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	interval1 = C.uint32_t(interval0)
	C.gtk_source_print_compositor_set_print_line_numbers(this1, interval1)
}
func (this0 *PrintCompositor) SetRightMargin(margin0 float64, unit0 gtk.Unit) {
	var this1 *C.GtkSourcePrintCompositor
	var margin1 C.double
	var unit1 C.GtkUnit
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	margin1 = C.double(margin0)
	unit1 = C.GtkUnit(unit0)
	C.gtk_source_print_compositor_set_right_margin(this1, margin1, unit1)
}
func (this0 *PrintCompositor) SetTabWidth(width0 int) {
	var this1 *C.GtkSourcePrintCompositor
	var width1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	width1 = C.uint32_t(width0)
	C.gtk_source_print_compositor_set_tab_width(this1, width1)
}
func (this0 *PrintCompositor) SetTopMargin(margin0 float64, unit0 gtk.Unit) {
	var this1 *C.GtkSourcePrintCompositor
	var margin1 C.double
	var unit1 C.GtkUnit
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	margin1 = C.double(margin0)
	unit1 = C.GtkUnit(unit0)
	C.gtk_source_print_compositor_set_top_margin(this1, margin1, unit1)
}
func (this0 *PrintCompositor) SetWrapMode(wrap_mode0 gtk.WrapMode) {
	var this1 *C.GtkSourcePrintCompositor
	var wrap_mode1 C.GtkWrapMode
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourcePrintCompositor()
	}
	wrap_mode1 = C.GtkWrapMode(wrap_mode0)
	C.gtk_source_print_compositor_set_wrap_mode(this1, wrap_mode1)
}
type SmartHomeEndType C.uint32_t
const (
	SmartHomeEndTypeDisabled SmartHomeEndType = 0
	SmartHomeEndTypeBefore SmartHomeEndType = 1
	SmartHomeEndTypeAfter SmartHomeEndType = 2
	SmartHomeEndTypeAlways SmartHomeEndType = 3
)
type StyleLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourceStyle() *C.GtkSourceStyle
}

type Style struct {
	gobject.Object
	
}

func ToStyle(objlike gobject.ObjectLike) *Style {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*Style)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*Style)(obj)
	}
	panic("cannot cast to Style")
}

func (this0 *Style) InheritedFromGtkSourceStyle() *C.GtkSourceStyle {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceStyle)(this0.C)
}

func (this0 *Style) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_style_get_type())
}

func StyleGetType() gobject.Type {
	return (*Style)(nil).GetStaticType()
}
func (this0 *Style) Copy() *Style {
	var this1 *C.GtkSourceStyle
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyle()
	}
	ret1 := C.gtk_source_style_copy(this1)
	var ret2 *Style
	ret2 = (*Style)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
type StyleSchemeLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourceStyleScheme() *C.GtkSourceStyleScheme
}

type StyleScheme struct {
	gobject.Object
	
}

func ToStyleScheme(objlike gobject.ObjectLike) *StyleScheme {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*StyleScheme)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*StyleScheme)(obj)
	}
	panic("cannot cast to StyleScheme")
}

func (this0 *StyleScheme) InheritedFromGtkSourceStyleScheme() *C.GtkSourceStyleScheme {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceStyleScheme)(this0.C)
}

func (this0 *StyleScheme) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_style_scheme_get_type())
}

func StyleSchemeGetType() gobject.Type {
	return (*StyleScheme)(nil).GetStaticType()
}
func (this0 *StyleScheme) GetAuthors() []string {
	var this1 *C.GtkSourceStyleScheme
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleScheme()
	}
	ret1 := C.gtk_source_style_scheme_get_authors(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func (this0 *StyleScheme) GetDescription() string {
	var this1 *C.GtkSourceStyleScheme
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleScheme()
	}
	ret1 := C.gtk_source_style_scheme_get_description(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *StyleScheme) GetFilename() string {
	var this1 *C.GtkSourceStyleScheme
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleScheme()
	}
	ret1 := C.gtk_source_style_scheme_get_filename(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *StyleScheme) GetID() string {
	var this1 *C.GtkSourceStyleScheme
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleScheme()
	}
	ret1 := C.gtk_source_style_scheme_get_id(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *StyleScheme) GetName() string {
	var this1 *C.GtkSourceStyleScheme
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleScheme()
	}
	ret1 := C.gtk_source_style_scheme_get_name(this1)
	var ret2 string
	ret2 = C.GoString(ret1)
	return ret2
}
func (this0 *StyleScheme) GetStyle(style_id0 string) *Style {
	var this1 *C.GtkSourceStyleScheme
	var style_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleScheme()
	}
	style_id1 = _GoStringToGString(style_id0)
	defer C.free(unsafe.Pointer(style_id1))
	ret1 := C.gtk_source_style_scheme_get_style(this1, style_id1)
	var ret2 *Style
	ret2 = (*Style)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
type StyleSchemeManagerLike interface {
	gobject.ObjectLike
	InheritedFromGtkSourceStyleSchemeManager() *C.GtkSourceStyleSchemeManager
}

type StyleSchemeManager struct {
	gobject.Object
	
}

func ToStyleSchemeManager(objlike gobject.ObjectLike) *StyleSchemeManager {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*StyleSchemeManager)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*StyleSchemeManager)(obj)
	}
	panic("cannot cast to StyleSchemeManager")
}

func (this0 *StyleSchemeManager) InheritedFromGtkSourceStyleSchemeManager() *C.GtkSourceStyleSchemeManager {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceStyleSchemeManager)(this0.C)
}

func (this0 *StyleSchemeManager) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_style_scheme_manager_get_type())
}

func StyleSchemeManagerGetType() gobject.Type {
	return (*StyleSchemeManager)(nil).GetStaticType()
}
func NewStyleSchemeManager() *StyleSchemeManager {
	ret1 := C.gtk_source_style_scheme_manager_new()
	var ret2 *StyleSchemeManager
	ret2 = (*StyleSchemeManager)(gobject.ObjectWrap(unsafe.Pointer(ret1), false))
	return ret2
}
func StyleSchemeManagerGetDefault() *StyleSchemeManager {
	ret1 := C.gtk_source_style_scheme_manager_get_default()
	var ret2 *StyleSchemeManager
	ret2 = (*StyleSchemeManager)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *StyleSchemeManager) AppendSearchPath(path0 string) {
	var this1 *C.GtkSourceStyleSchemeManager
	var path1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleSchemeManager()
	}
	path1 = _GoStringToGString(path0)
	defer C.free(unsafe.Pointer(path1))
	C.gtk_source_style_scheme_manager_append_search_path(this1, path1)
}
func (this0 *StyleSchemeManager) ForceRescan() {
	var this1 *C.GtkSourceStyleSchemeManager
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleSchemeManager()
	}
	C.gtk_source_style_scheme_manager_force_rescan(this1)
}
func (this0 *StyleSchemeManager) GetScheme(scheme_id0 string) *StyleScheme {
	var this1 *C.GtkSourceStyleSchemeManager
	var scheme_id1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleSchemeManager()
	}
	scheme_id1 = _GoStringToGString(scheme_id0)
	defer C.free(unsafe.Pointer(scheme_id1))
	ret1 := C.gtk_source_style_scheme_manager_get_scheme(this1, scheme_id1)
	var ret2 *StyleScheme
	ret2 = (*StyleScheme)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *StyleSchemeManager) GetSchemeIDs() []string {
	var this1 *C.GtkSourceStyleSchemeManager
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleSchemeManager()
	}
	ret1 := C.gtk_source_style_scheme_manager_get_scheme_ids(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func (this0 *StyleSchemeManager) GetSearchPath() []string {
	var this1 *C.GtkSourceStyleSchemeManager
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleSchemeManager()
	}
	ret1 := C.gtk_source_style_scheme_manager_get_search_path(this1)
	var ret2 []string
	for i := range ret2 {
		ret2[i] = C.GoString((*(*[999999]*C.char)(unsafe.Pointer(ret1)))[i])
	}
	return ret2
}
func (this0 *StyleSchemeManager) PrependSearchPath(path0 string) {
	var this1 *C.GtkSourceStyleSchemeManager
	var path1 *C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleSchemeManager()
	}
	path1 = _GoStringToGString(path0)
	defer C.free(unsafe.Pointer(path1))
	C.gtk_source_style_scheme_manager_prepend_search_path(this1, path1)
}
func (this0 *StyleSchemeManager) SetSearchPath(path0 []string) {
	var this1 *C.GtkSourceStyleSchemeManager
	var path1 **C.char
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceStyleSchemeManager()
	}
	path1 = (**C.char)(C.malloc(C.size_t(int(unsafe.Sizeof(*path1)) * (len(path0) + 1))))
	defer C.free(unsafe.Pointer(path1))
	for i, e := range path0 {
		(*(*[999999]*C.char)(unsafe.Pointer(path1)))[i] = _GoStringToGString(e)
		defer C.free(unsafe.Pointer((*(*[999999]*C.char)(unsafe.Pointer(path1)))[i]))
	}
	(*(*[999999]*C.char)(unsafe.Pointer(path1)))[len(path0)] = nil
	C.gtk_source_style_scheme_manager_set_search_path(this1, path1)
}
type UndoManagerLike interface {
	ImplementsGtkSourceUndoManager() *C.GtkSourceUndoManager
}

type UndoManager struct {
	gobject.Object
	UndoManagerImpl
}

type UndoManagerImpl struct {}

func ToUndoManager(objlike gobject.ObjectLike) *UndoManager {
	t := (*UndoManagerImpl)(nil).GetStaticType()
	c := objlike.InheritedFromGObject()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*UndoManager)(obj)
	}
	panic("cannot cast to UndoManager")
}

func (this0 *UndoManagerImpl) ImplementsGtkSourceUndoManager() *C.GtkSourceUndoManager {
	obj := uintptr(unsafe.Pointer(this0)) - unsafe.Sizeof(uintptr(0))
	return (*C.GtkSourceUndoManager)((*gobject.Object)(unsafe.Pointer(obj)).C)
}

func (this0 *UndoManagerImpl) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_undo_manager_get_type())
}

func UndoManagerGetType() gobject.Type {
	return (*UndoManagerImpl)(nil).GetStaticType()
}
func (this0 *UndoManagerImpl) BeginNotUndoableAction() {
	var this1 *C.GtkSourceUndoManager
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceUndoManager()}
	C.gtk_source_undo_manager_begin_not_undoable_action(this1)
}
func (this0 *UndoManagerImpl) CanRedo() bool {
	var this1 *C.GtkSourceUndoManager
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceUndoManager()}
	ret1 := C.gtk_source_undo_manager_can_redo(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *UndoManagerImpl) CanRedoChanged() {
	var this1 *C.GtkSourceUndoManager
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceUndoManager()}
	C.gtk_source_undo_manager_can_redo_changed(this1)
}
func (this0 *UndoManagerImpl) CanUndo() bool {
	var this1 *C.GtkSourceUndoManager
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceUndoManager()}
	ret1 := C.gtk_source_undo_manager_can_undo(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *UndoManagerImpl) CanUndoChanged() {
	var this1 *C.GtkSourceUndoManager
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceUndoManager()}
	C.gtk_source_undo_manager_can_undo_changed(this1)
}
func (this0 *UndoManagerImpl) EndNotUndoableAction() {
	var this1 *C.GtkSourceUndoManager
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceUndoManager()}
	C.gtk_source_undo_manager_end_not_undoable_action(this1)
}
func (this0 *UndoManagerImpl) Redo() {
	var this1 *C.GtkSourceUndoManager
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceUndoManager()}
	C.gtk_source_undo_manager_redo(this1)
}
func (this0 *UndoManagerImpl) Undo() {
	var this1 *C.GtkSourceUndoManager
	if this0 != nil {
		this1 = this0.ImplementsGtkSourceUndoManager()}
	C.gtk_source_undo_manager_undo(this1)
}
type ViewLike interface {
	gtk.TextViewLike
	InheritedFromGtkSourceView() *C.GtkSourceView
}

type View struct {
	gtk.TextView
	atk.ImplementorIfaceImpl
	gtk.BuildableImpl
	gtk.ScrollableImpl
}

func ToView(objlike gobject.ObjectLike) *View {
	c := objlike.InheritedFromGObject()
	if c == nil {
		return nil
	}
	t := (*View)(nil).GetStaticType()
	obj := gobject.ObjectGrabIfType(unsafe.Pointer(c), t)
	if obj != nil {
		return (*View)(obj)
	}
	panic("cannot cast to View")
}

func (this0 *View) InheritedFromGtkSourceView() *C.GtkSourceView {
	if this0 == nil {
		return nil
	}
	return (*C.GtkSourceView)(this0.C)
}

func (this0 *View) GetStaticType() gobject.Type {
	return gobject.Type(C.gtk_source_view_get_type())
}

func ViewGetType() gobject.Type {
	return (*View)(nil).GetStaticType()
}
func NewView() *View {
	ret1 := C.gtk_source_view_new()
	var ret2 *View
	ret2 = (*View)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func NewViewWithBuffer(buffer0 BufferLike) *View {
	var buffer1 *C.GtkSourceBuffer
	if buffer0 != nil {
		buffer1 = buffer0.InheritedFromGtkSourceBuffer()
	}
	ret1 := C.gtk_source_view_new_with_buffer(buffer1)
	var ret2 *View
	ret2 = (*View)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *View) GetAutoIndent() bool {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_auto_indent(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *View) GetCompletion() *Completion {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_completion(this1)
	var ret2 *Completion
	ret2 = (*Completion)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *View) GetDrawSpaces() DrawSpacesFlags {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_draw_spaces(this1)
	var ret2 DrawSpacesFlags
	ret2 = DrawSpacesFlags(ret1)
	return ret2
}
func (this0 *View) GetGutter(window_type0 gtk.TextWindowType) *Gutter {
	var this1 *C.GtkSourceView
	var window_type1 C.GtkTextWindowType
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	window_type1 = C.GtkTextWindowType(window_type0)
	ret1 := C.gtk_source_view_get_gutter(this1, window_type1)
	var ret2 *Gutter
	ret2 = (*Gutter)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *View) GetHighlightCurrentLine() bool {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_highlight_current_line(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *View) GetIndentOnTab() bool {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_indent_on_tab(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *View) GetIndentWidth() int {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_indent_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *View) GetInsertSpacesInsteadOfTabs() bool {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_insert_spaces_instead_of_tabs(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *View) GetMarkAttributes(category0 string, priority0 *int) *MarkAttributes {
	var this1 *C.GtkSourceView
	var category1 *C.char
	var priority1 *C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	priority1 = (*C.int32_t)(unsafe.Pointer(priority0))
	ret1 := C.gtk_source_view_get_mark_attributes(this1, category1, priority1)
	var ret2 *MarkAttributes
	ret2 = (*MarkAttributes)(gobject.ObjectWrap(unsafe.Pointer(ret1), true))
	return ret2
}
func (this0 *View) GetRightMarginPosition() int {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_right_margin_position(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *View) GetShowLineMarks() bool {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_show_line_marks(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *View) GetShowLineNumbers() bool {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_show_line_numbers(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *View) GetShowRightMargin() bool {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_show_right_margin(this1)
	var ret2 bool
	ret2 = ret1 != 0
	return ret2
}
func (this0 *View) GetSmartHomeEnd() SmartHomeEndType {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_smart_home_end(this1)
	var ret2 SmartHomeEndType
	ret2 = SmartHomeEndType(ret1)
	return ret2
}
func (this0 *View) GetTabWidth() int {
	var this1 *C.GtkSourceView
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	ret1 := C.gtk_source_view_get_tab_width(this1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *View) GetVisualColumn(iter0 *gtk.TextIter) int {
	var this1 *C.GtkSourceView
	var iter1 *C.GtkTextIter
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	iter1 = (*C.GtkTextIter)(unsafe.Pointer(iter0))
	ret1 := C.gtk_source_view_get_visual_column(this1, iter1)
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
func (this0 *View) SetAutoIndent(enable0 bool) {
	var this1 *C.GtkSourceView
	var enable1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	enable1 = _GoBoolToCBool(enable0)
	C.gtk_source_view_set_auto_indent(this1, enable1)
}
func (this0 *View) SetDrawSpaces(flags0 DrawSpacesFlags) {
	var this1 *C.GtkSourceView
	var flags1 C.GtkSourceDrawSpacesFlags
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	flags1 = C.GtkSourceDrawSpacesFlags(flags0)
	C.gtk_source_view_set_draw_spaces(this1, flags1)
}
func (this0 *View) SetHighlightCurrentLine(hl0 bool) {
	var this1 *C.GtkSourceView
	var hl1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	hl1 = _GoBoolToCBool(hl0)
	C.gtk_source_view_set_highlight_current_line(this1, hl1)
}
func (this0 *View) SetIndentOnTab(enable0 bool) {
	var this1 *C.GtkSourceView
	var enable1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	enable1 = _GoBoolToCBool(enable0)
	C.gtk_source_view_set_indent_on_tab(this1, enable1)
}
func (this0 *View) SetIndentWidth(width0 int) {
	var this1 *C.GtkSourceView
	var width1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	width1 = C.int32_t(width0)
	C.gtk_source_view_set_indent_width(this1, width1)
}
func (this0 *View) SetInsertSpacesInsteadOfTabs(enable0 bool) {
	var this1 *C.GtkSourceView
	var enable1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	enable1 = _GoBoolToCBool(enable0)
	C.gtk_source_view_set_insert_spaces_instead_of_tabs(this1, enable1)
}
func (this0 *View) SetMarkAttributes(category0 string, attributes0 MarkAttributesLike, priority0 int) {
	var this1 *C.GtkSourceView
	var category1 *C.char
	var attributes1 *C.GtkSourceMarkAttributes
	var priority1 C.int32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	category1 = _GoStringToGString(category0)
	defer C.free(unsafe.Pointer(category1))
	if attributes0 != nil {
		attributes1 = attributes0.InheritedFromGtkSourceMarkAttributes()
	}
	priority1 = C.int32_t(priority0)
	C.gtk_source_view_set_mark_attributes(this1, category1, attributes1, priority1)
}
func (this0 *View) SetRightMarginPosition(pos0 int) {
	var this1 *C.GtkSourceView
	var pos1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	pos1 = C.uint32_t(pos0)
	C.gtk_source_view_set_right_margin_position(this1, pos1)
}
func (this0 *View) SetShowLineMarks(show0 bool) {
	var this1 *C.GtkSourceView
	var show1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	show1 = _GoBoolToCBool(show0)
	C.gtk_source_view_set_show_line_marks(this1, show1)
}
func (this0 *View) SetShowLineNumbers(show0 bool) {
	var this1 *C.GtkSourceView
	var show1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	show1 = _GoBoolToCBool(show0)
	C.gtk_source_view_set_show_line_numbers(this1, show1)
}
func (this0 *View) SetShowRightMargin(show0 bool) {
	var this1 *C.GtkSourceView
	var show1 C.int
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	show1 = _GoBoolToCBool(show0)
	C.gtk_source_view_set_show_right_margin(this1, show1)
}
func (this0 *View) SetSmartHomeEnd(smart_he0 SmartHomeEndType) {
	var this1 *C.GtkSourceView
	var smart_he1 C.GtkSourceSmartHomeEndType
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	smart_he1 = C.GtkSourceSmartHomeEndType(smart_he0)
	C.gtk_source_view_set_smart_home_end(this1, smart_he1)
}
func (this0 *View) SetTabWidth(width0 int) {
	var this1 *C.GtkSourceView
	var width1 C.uint32_t
	if this0 != nil {
		this1 = this0.InheritedFromGtkSourceView()
	}
	width1 = C.uint32_t(width0)
	C.gtk_source_view_set_tab_width(this1, width1)
}
type ViewGutterPosition C.int32_t
const (
	ViewGutterPositionLines ViewGutterPosition = -30
	ViewGutterPositionMarks ViewGutterPosition = -20
)
func CompletionErrorQuark() int {
	ret1 := C.gtk_source_completion_error_quark()
	var ret2 int
	ret2 = int(ret1)
	return ret2
}
