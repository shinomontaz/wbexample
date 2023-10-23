<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

use App\Http\Controllers\Spa\LoginController;
use App\Http\Controllers\Spa\CommandsController;
use App\Http\Controllers\Spa\DataController;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

Route::controller(AuthController::class)->group(function(){
    Route::post('login', 'login')->name('login');
});

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

Route::controller(CommandsController::class)->group(function(){
    Route::post('pause-fleet', 'PauseFleet')->name('pause-fleet');
    Route::post('add-point', 'AddPoint')->name('add-point');
    Route::post('generate-fleet', 'GenerateFleet')->name('generate-fleet');
});

Route::controller(DataController::class)->group(function(){
    Route::get('trucks', 'GetTrucks')->name('trucks');
    Route::get('points', 'GetPoints')->name('points');
    Route::get('positions', 'GetPositions')->name('positions');
});
