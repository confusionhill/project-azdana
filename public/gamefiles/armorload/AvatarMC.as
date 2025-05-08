package 
{
    import flash.display.*;
    import flash.events.*;
    import flash.geom.*;
    import flash.net.*;
    import flash.system.*;
    import flash.utils.*;

    public class AvatarMC extends MovieClip
    {
        var ldr:Loader;
        public var mcChar:MovieClip;
        var serverFilePath:String = "http://game.aqworlds.com/game01/gamefiles/";
        public var shadow:MovieClip;
        public var strGender:String;
        var defaultCT:ColorTransform;
        var strSkinLinkage:String;
        public var pAV:Object;

        public function AvatarMC()
        {
            ldr = new Loader();
            defaultCT = MovieClip(this).transform.colorTransform;
            serverFilePath = "http://game.aqworlds.com/game01/gamefiles/";
            pAV = new Object();
            visible = false;
            return;
        }// end function

        public function updateColor(param1:Object = null)
        {
            scanColor(this, param1);
            return;
        }// end function

        private function hideOptionalParts() : void
        {
            var _loc_1:*;
            var _loc_2:*;
            _loc_1 = ["cape", "backhair", "robe", "backrobe"];
            for (_loc_2 in _loc_1)
            {
                // label
                if (typeof(mcChar[_loc_1[_loc_2]]) != undefined)
                {
                    mcChar[_loc_1[_loc_2]].visible = false;
                }// end if
            }// end of for ... in
            return;
        }// end function

        public function setColor(param1:MovieClip, param2:String, param3:String) : void
        {
            var _loc_4:Number;
            _loc_4 = Number(pAV.objData["intColor" + param2]);
            param1.isColored = true;
            param1.intColor = _loc_4;
            param1.strLocation = param2;
            param1.strShade = param3;
            changeColor(param1, _loc_4, param3);
            return;
        }// end function

        private function ioErrorHandler(param1:IOErrorEvent) : void
        {
            trace("ioErrorHandler: " + param1);
            return;
        }// end function

        public function changeColor(param1:MovieClip, param2:Number, param3:String) : void
        {
            var _loc_4:ColorTransform;
            _loc_4 = new ColorTransform();
            _loc_4.color = param2;
            switch(param3.toUpperCase())
            {
                case "LIGHT":
                {
                    _loc_4.redOffset = _loc_4.redOffset + 100;
                    _loc_4.greenOffset = _loc_4.greenOffset + 100;
                    _loc_4.blueOffset = _loc_4.blueOffset + 100;
                    break;
                }// end case
                case "DARK":
                {
                    _loc_4.redOffset = _loc_4.redOffset - 25;
                    _loc_4.greenOffset = _loc_4.greenOffset - 50;
                    _loc_4.blueOffset = _loc_4.blueOffset - 50;
                    break;
                }// end case
                case "DARKER":
                {
                    _loc_4.redOffset = _loc_4.redOffset - 125;
                    _loc_4.greenOffset = _loc_4.greenOffset - 125;
                    _loc_4.blueOffset = _loc_4.blueOffset - 125;
                    break;
                }// end case
                default:
                {
                    break;
                }// end default
            }// end switch
            if (param1.transform.colorTransform.color != _loc_4.color || param1.transform.colorTransform.redOffset != _loc_4.redOffset)
            {
                param1.transform.colorTransform = _loc_4;
            }// end if
            return;
        }// end function

        private function onHairLoadComplete(param1:Event) : void
        {
            var AssetClass:Class;
            var event:* = param1;
            AssetClass = event.target.loader.contentLoaderInfo.applicationDomain.getDefinition(pAV.objData.strHairName + pAV.objData.strGender + "Hair");
            mcChar.head.hair.removeChildAt(0);
            mcChar.head.hair.addChild(new AssetClass);
            try
            {
                AssetClass = getDefinitionByName(pAV.objData.strHairName + pAV.objData.strGender + "HairBack") as Class;
                mcChar.backhair.removeChildAt(0);
                mcChar.backhair.addChild(new AssetClass);
                mcChar.backhair.visible = true;
            }// end try
            catch (err:Error)
            {
                mcChar.backhair.visible = false;
            }// end catch
            return;
        }// end function

        private function onLoadSkinComplete(param1:Event)
        {
            var AssetClass:Class;
            var evt:* = param1;
            strGender = pAV.objData.strGender;
            hideOptionalParts();
            try
            {
                AssetClass = getDefinitionByName(strSkinLinkage + strGender + "Head") as Class;
                mcChar.head.removeChildAt(0);
                mcChar.head.addChildAt(new AssetClass, 0);
            }// end try
            catch (err:Error)
            {
                AssetClass = getDefinitionByName("mcHead" + strGender) as Class;
                mcChar.head.removeChildAt(0);
                mcChar.head.addChildAt(new AssetClass, 0);
            }// end catch
            AssetClass = getDefinitionByName(strSkinLinkage + strGender + "Chest") as Class;
            mcChar.chest.removeChildAt(0);
            mcChar.chest.addChild(new AssetClass);
            AssetClass = getDefinitionByName(strSkinLinkage + strGender + "Hip") as Class;
            mcChar.hip.removeChildAt(0);
            mcChar.hip.addChild(new AssetClass);
            AssetClass = getDefinitionByName(strSkinLinkage + strGender + "FootIdle") as Class;
            mcChar.idlefoot.removeChildAt(0);
            mcChar.idlefoot.addChild(new AssetClass);
            AssetClass = getDefinitionByName(strSkinLinkage + strGender + "Foot") as Class;
            mcChar.backfoot.removeChildAt(0);
            mcChar.backfoot.addChild(new AssetClass);
            AssetClass = getDefinitionByName(strSkinLinkage + strGender + "Shoulder") as Class;
            mcChar.frontshoulder.removeChildAt(0);
            mcChar.frontshoulder.addChild(new AssetClass);
            mcChar.backshoulder.removeChildAt(0);
            mcChar.backshoulder.addChild(new AssetClass);
            AssetClass = getDefinitionByName(strSkinLinkage + strGender + "Hand") as Class;
            mcChar.fronthand.removeChildAt(0);
            mcChar.fronthand.addChild(new AssetClass);
            mcChar.backhand.removeChildAt(0);
            mcChar.backhand.addChild(new AssetClass);
            AssetClass = getDefinitionByName(strSkinLinkage + strGender + "Thigh") as Class;
            mcChar.frontthigh.removeChildAt(0);
            mcChar.frontthigh.addChild(new AssetClass);
            mcChar.backthigh.removeChildAt(0);
            mcChar.backthigh.addChild(new AssetClass);
            AssetClass = getDefinitionByName(strSkinLinkage + strGender + "Shin") as Class;
            mcChar.frontshin.removeChildAt(0);
            mcChar.frontshin.addChild(new AssetClass);
            mcChar.backshin.removeChildAt(0);
            mcChar.backshin.addChild(new AssetClass);
            try
            {
                AssetClass = getDefinitionByName(strSkinLinkage + strGender + "Robe") as Class;
                mcChar.robe.removeChildAt(0);
                mcChar.robe.addChild(new AssetClass);
                mcChar.robe.visible = true;
            }// end try
            catch (err:Error)
            {
                try
                {
                AssetClass = getDefinitionByName(strSkinLinkage + strGender + "RobeBack") as Class;
                mcChar.backrobe.removeChildAt(0);
                mcChar.backrobe.addChild(new AssetClass);
                mcChar.backrobe.visible = true;
				}// end catch
				catch (err:Error)
            	{
            	}// end catch
            }// end try

            visible = true;
            return;
        }// end function

        private function scanColor(param1:MovieClip, param2) : void
        {
            var _loc_3:int;
            var _loc_4:DisplayObject;
            if ("isColored" in param1)
            {
                changeColor(param1, Number(param2["intColor" + param1.strLocation]), param1.strShade);
            }// end if
            _loc_3 = 0;
            while (_loc_3 < param1.numChildren)
            {
                // label
                _loc_4 = param1.getChildAt(_loc_3);
                if (_loc_4 is MovieClip)
                {
                    scanColor(MovieClip(_loc_4), param2);
                }// end if
                _loc_3++;
            }// end while
            return;
        }// end function

        public function loadHair() : void
        {
            var _loc_1:*;
            _loc_1 = new Loader();
            _loc_1.load(new URLRequest(serverFilePath + pAV.objData.strHairFileName), new LoaderContext(false, ApplicationDomain.currentDomain));
            _loc_1.contentLoaderInfo.addEventListener(Event.COMPLETE, onHairLoadComplete);
            return;
        }// end function

        public function loadArmor(param1:String, param2:String)
        {
            strSkinLinkage = param2;
            ldr.load(new URLRequest(serverFilePath + "classes/" + pAV.objData.strGender + "/" + param1), new LoaderContext(false, ApplicationDomain.currentDomain));
            ldr.contentLoaderInfo.addEventListener(Event.COMPLETE, onLoadSkinComplete);
            ldr.contentLoaderInfo.addEventListener(IOErrorEvent.IO_ERROR, ioErrorHandler);
            return;
        }// end function

    }
}
