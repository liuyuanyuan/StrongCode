package com.highgo.admin;

import java.util.List;

import org.eclipse.jface.action.ActionContributionItem;
import org.eclipse.jface.action.IAction;
import org.eclipse.jface.action.IContributionItem;
import org.eclipse.jface.viewers.ISelection;
import org.eclipse.jface.viewers.ISelectionProvider;
import org.eclipse.jface.viewers.IStructuredSelection;
import org.eclipse.ui.IWorkbenchPart;
import org.eclipse.ui.IWorkbenchWindow;
import org.jkiss.dbeaver.Log;
import org.jkiss.dbeaver.core.DBeaverUI;
import org.jkiss.dbeaver.model.struct.DBSObject;
import org.jkiss.dbeaver.ui.ActionUtils;
import org.jkiss.dbeaver.ui.DBeaverIcons;
import org.jkiss.dbeaver.ui.actions.common.EmptyListAction;
import org.jkiss.dbeaver.ui.actions.datasource.DataSourceMenuContributor;
import org.jkiss.dbeaver.ui.navigator.NavigatorUtils;
import org.jkiss.utils.CommonUtils;

public class AdvanceToolsContributor extends DataSourceMenuContributor 
{

	private static Log log = Log.getLog(AdvanceToolsContributor.class);
	
	@Override
	protected void fillContributionItems(List<IContributionItem> menuItems) 
	{
		log.debug("AdvanceToolsContributor.fillContributionItems");
		
		IWorkbenchPart activePart = DBeaverUI.getActiveWorkbenchWindow().getActivePage().getActivePart();
        if (activePart == null) 
        {
            return;
        }
        
        final ISelectionProvider selectionProvider = activePart.getSite().getSelectionProvider();
        if (selectionProvider == null) 
        {
            return;
        }
        
        ISelection selection = selectionProvider.getSelection();
        if (!(selection instanceof IStructuredSelection)) 
        {
            return;
        }
        
        DBSObject selectedObject = NavigatorUtils.getSelectedObject((IStructuredSelection) selection);
        if (selectedObject == null)
        {
        	return;
        } else
        {
            List<AdvanceToolDescriptor> tools = AdvanceToolsRegistry.getInstance().getTools();//
            log.debug("tools size=" + tools.size());
            fillToolsMenu(menuItems, tools, selection);
        }		
	}
	
	
	private static void fillToolsMenu(List<IContributionItem> menuItems, List<AdvanceToolDescriptor> tools, ISelection selection)
    {
		log.debug("AdvanceToolsContributor.fillToolsMenu");
		
        boolean hasTools = false;
        if (!CommonUtils.isEmpty(tools)) 
        {
            IWorkbenchWindow workbenchWindow = DBeaverUI.getActiveWorkbenchWindow();
            if (workbenchWindow.getActivePage() != null) 
            {
                IWorkbenchPart activePart = workbenchWindow.getActivePage().getActivePart();
                if (activePart != null) 
                {
                    //Map<ToolGroupDescriptor, IMenuManager> groupsMap = new HashMap<>();
                    for (AdvanceToolDescriptor tool : tools) //
                    {
                        hasTools = true;
                        //IMenuManager parentMenu = null;
                        //if (tool.getGroup() != null) {
                        //    parentMenu = getGroupMenu(menuItems, groupsMap, tool.getGroup());
                        //}
                        IAction action = ActionUtils.makeAction(
                            new NavigatorActionExecuteAdvanceTool(workbenchWindow, tool),//
                            activePart.getSite(),
                            selection,
                            tool.getLabel(),
                            tool.getIcon() == null ? null : DBeaverIcons.getImageDescriptor(tool.getIcon()),
                            tool.getDescription());
                        //if (parentMenu == null) {
                            menuItems.add(new ActionContributionItem(action));
                        //} else {
                        //    parentMenu.add(new ActionContributionItem(action));
                        //}
                    }
                }
            }
        }
        if (!hasTools) {
            menuItems.add(new ActionContributionItem(new EmptyListAction()));
        }
    }

}
