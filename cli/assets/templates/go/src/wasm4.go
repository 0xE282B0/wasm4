//
// WASM-4: https://wasm4.org/docs

package main

import "unsafe"

func main () {}

// ┌───────────────────────────────────────────────────────────────────────────┐
// │                                                                           │
// │ Platform Constants                                                        │
// │                                                                           │
// └───────────────────────────────────────────────────────────────────────────┘

const SCREEN_SIZE int = 160;

// ┌───────────────────────────────────────────────────────────────────────────┐
// │                                                                           │
// │ Memory Addresses                                                          │
// │                                                                           │
// └───────────────────────────────────────────────────────────────────────────┘

var PALETTE = (*[4]uint32)(unsafe.Pointer(uintptr(0x04)));
var DRAW_COLORS = (*uint16)(unsafe.Pointer(uintptr(0x14)));
var GAMEPAD1 = (*uint8)(unsafe.Pointer(uintptr(0x16)));
var GAMEPAD2 = (*uint8)(unsafe.Pointer(uintptr(0x17)));
var GAMEPAD3 = (*uint8)(unsafe.Pointer(uintptr(0x18)));
var GAMEPAD4 = (*uint8)(unsafe.Pointer(uintptr(0x19)));
var MOUSE_X = (*int16)(unsafe.Pointer(uintptr(0x1a)));
var MOUSE_Y = (*int16)(unsafe.Pointer(uintptr(0x1c)));
var MOUSE_BUTTONS = (*uint8)(unsafe.Pointer(uintptr(0x1e)));
var FRAMEBUFFER = (*[6400]uint8)(unsafe.Pointer(uintptr(0xa0)));

const BUTTON_1 byte = 1;
const BUTTON_2 byte = 2;
const BUTTON_LEFT byte = 16;
const BUTTON_RIGHT byte = 32;
const BUTTON_UP byte = 64;
const BUTTON_DOWN byte = 128;

// ┌───────────────────────────────────────────────────────────────────────────┐
// │                                                                           │
// │ Drawing Functions                                                         │
// │                                                                           │
// └───────────────────────────────────────────────────────────────────────────┘

/** Copies pixels to the framebuffer. */
//go:export blit
func blit (sprite *byte, x int, y int, width uint, height uint, flags uint);

/** Copies a subregion within a larger sprite atlas to the framebuffer. */
//go:export blitSub
func blitSub (sprite *byte, x int, y int, width uint, height uint,
    srcX uint, srcY uint, stride int, flags uint);

const BLIT_2BPP = 1;
const BLIT_1BPP = 0;
const BLIT_FLIP_X = 2;
const BLIT_FLIP_Y = 4;
const BLIT_ROTATE = 8;

/** Draws a line between two points. */
//go:export line
func line (x1 int, y1 int, x2 int, y2 int);

/** Draws an oval (or circle). */
//go:export oval
func oval (x int, y int, width uint, height uint);

/** Draws a rectangle. */
//go:export rect
func rect (x int, y int, width uint, height uint);

/** Draws text using the built-in system font. */
//go:export textUtf8
func text (text string, x int, y int);

// ┌───────────────────────────────────────────────────────────────────────────┐
// │                                                                           │
// │ Sound Functions                                                           │
// │                                                                           │
// └───────────────────────────────────────────────────────────────────────────┘

/** Plays a sound tone. */
//go:export tone
func tone (frequency uint, duration uint, volume uint, flags uint);

const TONE_PULSE1 = 0;
const TONE_PULSE2 = 1;
const TONE_TRIANGLE = 2;
const TONE_NOISE = 3;
const TONE_MODE1 = 0;
const TONE_MODE2 = 4;
const TONE_MODE3 = 8;
const TONE_MODE4 = 12;

// ┌───────────────────────────────────────────────────────────────────────────┐
// │                                                                           │
// │ Storage Functions                                                         │
// │                                                                           │
// └───────────────────────────────────────────────────────────────────────────┘

/** Reads up to `size` bytes from persistent storage into the pointer `destPtr`. */
//go:export diskr
func diskr (ptr unsafe.Pointer, count uint) uint;

/** Writes up to `size` bytes from the pointer `srcPtr` into persistent storage. */
//go:export diskw
func diskw (src unsafe.Pointer, count uint) uint;

// ┌───────────────────────────────────────────────────────────────────────────┐
// │                                                                           │
// │ Other Functions                                                           │
// │                                                                           │
// └───────────────────────────────────────────────────────────────────────────┘

/** Copies `size` bytes from `srcPtr` into `destPtr`. */
//go:export memcpy
func memcpy (dest *byte, src *byte, size uint) uint;

/** Fills memory at `destPtr` with `size` bytes of the fixed value `value`. */
//go:export memset
func memset (dest *byte, b byte, size uint) uint;

/** Prints a message to the debug console. */
//go:export traceUtf8
func trace (str string, byteLength int);
