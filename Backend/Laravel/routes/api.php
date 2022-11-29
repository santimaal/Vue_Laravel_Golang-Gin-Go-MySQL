<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\TableController;
use App\Http\Controllers\ThematicController;
use App\Http\Controllers\UserController;
use App\Http\Controllers\ReserveController;

Route::get('/table', [TableController::class, 'index']);
Route::get('/table/{id}', [TableController::class, 'show']);
Route::post('/table', [TableController::class, 'store']);
Route::put('/table/{id}', [TableController::class, 'update']);
Route::delete('/table/{id}', [TableController::class, 'destroy']);
Route::get('/thematic', [ThematicController::class, 'index']);
Route::get('/thematic/{id}', [ThematicController::class, 'show']);
Route::post('/thematic', [ThematicController::class, 'store']);
Route::put('/thematic/{id}', [ThematicController::class, 'update']);
Route::delete('/thematic/{id}', [ThematicController::class, 'destroy']);
// Route::resource('table', TableController::class);
// Route::resource('thematic', ThematicController::class);
Route::resource('user', UserController::class);
Route::resource('reserve', ReserveController::class);